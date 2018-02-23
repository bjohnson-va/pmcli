package grpc

// ServerTemplateData data for generating the server
type ServerTemplateData struct {
	Name           string
	VerifyIdentity bool
}

const (
	//GRPCServerRegisterTag is used to mark where the server definitions begin (makes it easier to update later)
	GRPCServerRegisterTag = "//REGISTER_GRPC_SERVERS_HERE"

	//ServerTemplate Template used for the whole GRPC Server
	ServerTemplate = `package main

import (
    "net/http"
	"os"
	"strings"

    "context"
    "time"

	"github.com/vendasta/gosdks/logging"
    "github.com/vendasta/gosdks/config"
    "github.com/vendasta/gosdks/config/elastic"
    "github.com/vendasta/gosdks/statsd"
    "github.com/vendasta/gosdks/vax"
    "github.com/vendasta/gosdks/serverconfig"
    "github.com/vendasta/gosdks/iam"
    "github.com/vendasta/gosdks/vstore"
    "google.golang.org/grpc"
)

const (
    APP_NAME = "{{.Name}}"
    GRPC_PORT = 11000
    HTTP_PORT = 11001
)

var (
    {{if .VerifyIdentity}}
	// AuthorizedServiceAccounts are service accounts with access to this microservice.
	// FILL THIS OUT
	AuthorizedServiceAccounts = map[config.Env]iam.ServiceAccountToScopes{
		config.Prod: {
			"account@appspot.gserviceaccount.com": {iam.READ, iam.WRITE, iam.DELETE},
			"otheraccount@appspot.gserviceaccount.com": {iam.READ},
		},
		config.Demo: {
			"account@appspot.gserviceaccount.com": {iam.READ, iam.WRITE, iam.DELETE},
			"otheraccount@appspot.gserviceaccount.com": {iam.READ},
		},
		config.Test: {
			"account@appspot.gserviceaccount.com": {iam.READ, iam.WRITE, iam.DELETE},
			"otheraccount@appspot.gserviceaccount.com": {iam.READ},
		},
		config.Local: {
			"account@appspot.gserviceaccount.com": {iam.READ, iam.WRITE, iam.DELETE},
			"otheraccount@appspot.gserviceaccount.com": {iam.READ},
		},
	}
    {{else}}
    AuthorizedTokens = map[string][]serverconfig.ApiKeyUser{
		"local": []serverconfig.ApiKeyUser{
			serverconfig.ApiKeyUser{Key: "local-api-key", UID: "UID-LL25G8C8N"},
		},
		"test": []serverconfig.ApiKeyUser{
			serverconfig.ApiKeyUser{Key: "test-api-key", UID: "UID-LL25G8C8N"},
		},
	}
    {{end}}
)

func main() {
    var err error
    ctx := context.Background()

    //Setup Application logging and switch the logger
    if !config.IsLocal() {
        namespace := config.GetGkeNamespace()
        podName := config.GetGkePodName()
        if err = logging.Initialize(namespace, podName, APP_NAME); err != nil {
            logging.Criticalf(ctx, "Error initializing logger: %s", err.Error())
            os.Exit(-1)
        }
    }

    //Setup ElasticSearch Client
    if err = elasticclient.Initialize(); err != nil {
        logging.Criticalf(ctx, "Error initilizing Elastic Client: %s", err.Error())
        os.Exit(-1)
    }

    //Setup StatsD Client
    if err = statsd.Initialize(APP_NAME, nil); err != nil {
        logging.Criticalf(ctx, "Error initilizing statsd client: %s", err.Error())
        os.Exit(-1)
    }
	iamClient, err := iam.NewClient(context.Background(), config.CurEnv(), grpc.WithUnaryInterceptor(logging.ClientInterceptor()))
 	if err != nil {
		logging.Criticalf(ctx, "Error initializing iam client %s", err.Error())
		os.Exit(-1)
	}

    {{if .VerifyIdentity}}
    //Create Auth Interceptor, pass the list of authorized service accounts
    var iamAuthService = iam.NewAuthService(iamClient, AuthorizedServiceAccounts[config.CurEnv()])
    iamAuthService.AllowPublicMethods(strings.Split(os.Getenv("PUBLIC_ROUTES"), ",")...)
    var ii = iamAuthService.Interceptor()
    {{else}}
    var ii = serverconfig.NewApiKeyAuthInterceptor(AuthorizedTokens).Interceptor()
    {{end}}
    //Create Logging Interceptor
    var loggingInterceptor = logging.Interceptor()

    //Create Timeout Interceptor
    var timeoutInterceptor = vax.TimeoutInterceptor(20 * time.Second)

    //Create a GRPC Server
    logging.Infof(ctx, "Creating GRPC server...")
    grpcServer := serverconfig.CreateGrpcServer(loggingInterceptor, ii, timeoutInterceptor)

    //Create a vStore client
    logging.Infof(ctx, "Creating vStore Client...")
	vstoreClient, err := vstore.New()
	if err != nil {
		logging.Criticalf(ctx, "Error initializing vstore client %s", err.Error())
		os.Exit(-1)
	}
	logging.Infof(ctx, "Using vStore Client: %#v", vstoreClient)

    //--------- INSERT YOUR CODE HERE ------------
	` + GRPCServerRegisterTag + `

    //Start GRPC API Server
    go func() {
        logging.Infof(ctx, "Running GRPC server...")
        if err = serverconfig.StartGrpcServer(grpcServer, GRPC_PORT); err != nil {
            logging.Criticalf(ctx, "Error starting GRPC Server: %s", err.Error())
            os.Exit(-1)
        }
    }()

    //Start Healthz and Debug HTTP API Server
    healthz := func(w http.ResponseWriter, _ *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        //TODO: (optional) INSERT YOUR CODE HERE
        return
    }

    logging.Infof(ctx, "Running HTTP server...")
    mux := http.NewServeMux()
	if err = serverconfig.StartHTTPServer(healthz, HTTP_PORT, mux); err != nil {
		logging.Criticalf(ctx, "Error starting Healthz & Debug server: %s", err.Error())
		os.Exit(-1)
	}
}
`

	//ServerTemplate Template used for the whole GRPC Server
	GrpcFuncTemplate = `
type {{.FunctionName}}TypedHandler interface {
	AuthenticateRequest(ctx context.Context, req {{.GoRequestNamespace}}.{{.RequestType}}) error
	ValidateRequest(req {{.GoRequestNamespace}}.{{.RequestType}}) error
	RequestFromProto(req {{.GoRequestNamespace}}.{{.RequestType}}) ({{.DomainNamespace}}.{{.RequestType}}, error)
	Command(ctx context.Context, req {{.DomainNamespace}}.{{.RequestType}}) ({{.DomainNamespace}}.{{.ResponseType}}, error)
	ResponseToProto(res {{.DomainNamespace}}.{{.ResponseType}}) ({{.GoResponseNamespace}}.{{.ResponseType}}, error)
}

type {{.FunctionName}}Handler struct {
	th {{.FunctionName}}TypedHandler
}

func (h *{{.FunctionName}}Handler) AuthenticateRequest(ctx context.Context, req grpchandler.GrpcRequest) error {
	treq, ok := req.({{.GoRequestNamespace}}.{{.RequestType}})
	if !ok {
		return util.Error(util.InvalidArgument, "Invalid type")
	}
	return h.th.AuthenticateRequest(ctx, treq)
}

func (h *{{.FunctionName}}Handler) ValidateRequest(req grpchandler.GrpcRequest) error {
	treq, ok := req.({{.GoRequestNamespace}}.{{.RequestType}})
	if !ok {
		return util.Error(util.InvalidArgument, "Invalid type")
	}
	return h.th.ValidateRequest(treq)
}

func (h *{{.FunctionName}}Handler) RequestFromProto(req grpchandler.GrpcRequest) (grpchandler.DomainRequest, error) {
	treq, ok := req.({{.GoRequestNamespace}}.{{.RequestType}})
	if !ok {
		return nil, util.Error(util.InvalidArgument, "Invalid type")
	}
	return h.th.RequestFromProto(treq)
}

func (h* {{.FunctionName}}Handler) Command(ctx context.Context, req grpchandler.DomainRequest) (grpchandler.DomainResponse, error) {
	treq, ok := req.({{.DomainNamespace}}.{{.RequestType}})
	if !ok {
		return nil, util.Error(util.InvalidArgument, "Invalid type")
	}
	return h.th.Command(ctx, treq)
}

func (h *{{.FunctionName}}Handler) ResponseToProto(res grpchandler.DomainResponse) (grpchandler.GrpcResponse, error) {
	tres, ok := res.({{.DomainNamespace}}.{{.ResponseType}})
	if !ok {
		return nil, util.Error(util.InvalidArgument, "Invalid type")
	}
	return h.th.ResponseToProto(tres)
}

func (grpc *{{.GrpcServiceName}}) {{.FunctionName}}(ctx context.Context, grpcReq *{{.GoRequestNamespace}}.{{.RequestType}}) (*{{.GoResponseNamespace}}.{{.ResponseType}}, error) {
	res, err := grpchandler.Handle(ctx, *grpcReq, &{{.FunctionName}}Handler{grpc.{{.FunctionName}}TypedHandler})
	if err != nil {
		return nil, err
	}
	tres, ok := res.({{.GoResponseNamespace}}.{{.ResponseType}})
	if !ok {
		return nil, util.Error(util.InvalidArgument, "Invalid type")
	}
	return &tres, nil
}

`
)
