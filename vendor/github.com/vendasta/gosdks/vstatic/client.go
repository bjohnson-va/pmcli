package vstatic

import (
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/vendasta/gosdks/config"
	"github.com/vendasta/gosdks/logging"
	"github.com/vendasta/gosdks/pb/vstatic/v1"
	"github.com/vendasta/gosdks/vax"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

const cookieKeyDeploymentID = "vstatic-deploymentId"
const queryParamDeploymentID = "v"

var addresses = map[config.Env]string{
	config.Local: "vstatic:11000",
	config.Test:  "vstatic-api-test.apigateway.co:443",
	config.Demo:  "vstatic-api-demo.apigateway.co:443",
	config.Prod:  "vstatic-api-prod.apigateway.co:443",
}

var scopes = map[config.Env]string{
	config.Local: "",
	config.Test:  "https://vstatic-api-test.vendasta-internal.com",
	config.Demo:  "https://vstatic-api-demo.vendasta-internal.com",
	config.Prod:  "https://vstatic-api-prod.vendasta-internal.com",
}

func toProtoEnv(env config.Env) vstatic_v1.Environment {
	if env == config.Test {
		return vstatic_v1.Environment_Test
	} else if env == config.Demo {
		return vstatic_v1.Environment_Demo
	} else if env == config.Prod {
		return vstatic_v1.Environment_Prod
	}
	return vstatic_v1.Environment_All
}

// GetIAMToken will be called when serving the get index call so you can determine which iamToken to use, "" is ok
type GetIAMToken func(r *http.Request) (string, error)

// NewClient returns an vStatic client.
func NewClient(ctx context.Context, appID string, env config.Env, dialOptions ...grpc.DialOption) (Interface, error) {
	address := addresses[config.Prod]
	scope := scopes[config.Prod]
	conn, err := vax.NewGRPCConnection(ctx, address, true, scope, true, dialOptions...)
	if err != nil {
		return nil, err
	}
	return &vstaticClient{
		client: &grpcClient{client: vstatic_v1.NewVStaticClient(conn)},
		env:    toProtoEnv(env),
		appID:  appID,
	}, nil
}

// Interface defines all of the API methods available from vStatic.
type Interface interface {
	GetIndexHTMLHandler(ctx context.Context, iamTokenFunc GetIAMToken) func(http.ResponseWriter, *http.Request)
	GetAssetHandler(ctx context.Context, serveDirect bool) func(http.ResponseWriter, *http.Request)
}

type vstaticClient struct {
	client grpcGetter
	env    vstatic_v1.Environment
	appID  string
}

func (ic *vstaticClient) GetIndexHTMLHandler(ctx context.Context, iamTokenFunc GetIAMToken) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ic.serveIndex(ctx, w, r, iamTokenFunc)
	}
}

func (ic *vstaticClient) GetAssetHandler(ctx context.Context, serveDirect bool) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ic.serveAsset(ctx, w, r, serveDirect)
	}
}

func (ic *vstaticClient) serveIndex(ctx context.Context, w http.ResponseWriter, r *http.Request, iamTokenFunc GetIAMToken) {
	r.ParseForm()
	t, err := iamTokenFunc(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	d := getDeploymentID(r)
	html, deploymentID, err := ic.client.getIndex(ctx, ic.appID, ic.env, t, d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	c := &http.Cookie{
		Name:  cookieKeyDeploymentID,
		Value: deploymentID,
		Path:  "/",
	}
	http.SetCookie(w, c)
	w.WriteHeader(http.StatusOK)
	w.Write(html)
}

func getDeploymentID(r *http.Request) string {
	p := r.URL.Query()
	id, ok := p[queryParamDeploymentID]
	if ok {
		return id[0]
	}
	return ""
}

func (ic *vstaticClient) serveAsset(ctx context.Context, w http.ResponseWriter, r *http.Request, serveDirect bool) {
	deploymentID := getCurrentDeployment(ctx, r)

	servingURL, err := ic.client.getAssetURL(ctx, ic.appID, ic.env, r.URL.Path, deploymentID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if serveDirect || strings.HasSuffix(r.URL.Path, ".svg") {
		serveFile(ctx, w, servingURL)
		return
	}
	logging.Infof(ctx, "Redirecting to serve file: %s", r.URL.Path)
	http.Redirect(w, r, servingURL, 307)
}

func getCurrentDeployment(ctx context.Context, r *http.Request) string {
	deploymentID := ""
	c, err := r.Cookie(cookieKeyDeploymentID)
	if err != nil {
		logging.Warningf(ctx, "Could not find deployment id in cookie")
	} else {
		deploymentID = c.Value
	}
	return deploymentID
}

func serveFile(ctx context.Context, w http.ResponseWriter, URL string) {
	logging.Infof(ctx, "Serving file directly,: %s", URL)
	resp, err := http.Get(URL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	for k, h := range resp.Header {
		w.Header()[k] = h
	}
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}

type grpcGetter interface {
	getIndex(ctx context.Context, appID string, env vstatic_v1.Environment, iamToken string, deploymentID string) (HTML, string, error)
	getAssetURL(ctx context.Context, appID string, env vstatic_v1.Environment, assetPath string, deploymentID string) (string, error)
}

type grpcClient struct {
	client vstatic_v1.VStaticClient
}

// getAssetURL gets the asset url from vstatic
func (ic *grpcClient) getAssetURL(ctx context.Context, appID string, env vstatic_v1.Environment, assetPath string, deploymentID string) (string, error) {
	var resp *vstatic_v1.GetAssetURLResponse
	err := vax.Invoke(ctx, func(context.Context, vax.CallSettings) error {
		var err error
		resp, err = ic.client.GetAssetURL(ctx, &vstatic_v1.GetAssetURLRequest{
			AppId:        appID,
			Environment:  env,
			DeploymentId: deploymentID,
			AssetPath:    assetPath,
		}, grpc.FailFast(false))
		return err
	}, defaultRetryCallOptions)
	if err != nil {
		return "", err
	}
	return resp.ServingUrl, nil
}

// HTML represents content of your index.html
type HTML []byte

// getIndex implements interface
func (ic *grpcClient) getIndex(ctx context.Context, appID string, env vstatic_v1.Environment, iamToken string, deploymentID string) (HTML, string, error) {
	var resp *vstatic_v1.GetIndexResponse
	err := vax.Invoke(ctx, func(context.Context, vax.CallSettings) error {
		var err error
		resp, err = ic.client.GetIndex(ctx, &vstatic_v1.GetIndexRequest{
			AppId:        appID,
			Environment:  env,
			IamToken:     iamToken,
			DeploymentId: deploymentID,
		}, grpc.FailFast(false))
		return err
	}, defaultRetryCallOptions)
	if err != nil {
		return nil, "", err
	}
	return resp.Html, resp.DeploymentId, nil
}

var defaultRetryCallOptions = vax.WithRetry(func() vax.Retryer {
	return vax.OnCodes([]codes.Code{
		codes.DeadlineExceeded,
		codes.Unavailable,
		codes.Unknown,
	}, vax.Backoff{
		Initial:    10 * time.Millisecond,
		Max:        300 * time.Millisecond,
		Multiplier: 3,
	})
})
