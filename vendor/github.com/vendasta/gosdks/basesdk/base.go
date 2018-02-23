package basesdk

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"

	"time"

	"github.com/vendasta/gosdks/logging"
	"github.com/vendasta/gosdks/util"
	"github.com/vendasta/gosdks/vax"
	"golang.org/x/net/context"
)

// SDKClient is the base client which allows calls to a server to be performed over http
type SDKClient interface {
	// Get performs a get request to the server
	Get(ctx context.Context, path string, params map[string]interface{}, ro ...RequestOpts) (*http.Response, error)

	// Post perform a post request to the server
	Post(ctx context.Context, path string, params map[string]interface{}, ro ...RequestOpts) (*http.Response, error)

	// Head performs a head request to the server
	Head(ctx context.Context, path string, ro ...RequestOpts) error
}

// BaseClient allows requests to be made to a server. This includes the request authorization and the different
// environments that can be called.
type BaseClient struct {
	Authorization RequestAuthorization
	RootURL       string
	CallOption    vax.CallOption
}

// HTTPClient provides an interface required to perform an http request
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type requestOptions struct {
	// indicates that the request is idempotent and errors >= 500 should be retried on.
	idempotent bool
}

// RequestOpts configures the requestOpts for making an HTTP request.
type RequestOpts func(*requestOptions)

// Idempotent indicates that the outgoing request is idempotent and errors >= 500 can and should be retried on.
func Idempotent() RequestOpts {
	return func(ro *requestOptions) {
		ro.idempotent = true
	}
}

// Get performs an http get request
func (client BaseClient) Get(ctx context.Context, path string, params map[string]interface{}, ro ...RequestOpts) (*http.Response, error) {
	return client.performRequest(ctx, &http.Client{}, http.MethodGet, path, params, ro...)
}

// Post performs an http post request
func (client BaseClient) Post(ctx context.Context, path string, params map[string]interface{}, ro ...RequestOpts) (*http.Response, error) {
	return client.performRequest(ctx, &http.Client{}, http.MethodPost, path, params, ro...)
}

// Head performs an http head request
func (client BaseClient) Head(ctx context.Context, path string, ro ...RequestOpts) error {
	_, err := client.performRequest(ctx, &http.Client{}, http.MethodHead, path, nil, ro...)
	return err
}

func (client BaseClient) performRequest(ctx context.Context, httpClient HTTPClient, method string, path string, params map[string]interface{}, requestOpts ...RequestOpts) (*http.Response, error) {
	if client.CallOption == nil {
		client.CallOption = defaultCallOption
	}
	validMethods := []string{http.MethodGet, http.MethodPost, http.MethodHead}
	if !util.StringInSlice(method, validMethods) {
		return nil, util.Error(util.InvalidArgument, "method %s not valid, must be one of %s", method, validMethods)
	}

	r := requestOptions{}
	for _, opt := range requestOpts {
		opt(&r)
	}

	var response *http.Response
	err := vax.Invoke(ctx, func(ctx context.Context, cs vax.CallSettings) error {
		request, err := client.buildRequest(ctx, method, path, params)
		if err != nil {
			return err
		}
		client.Authorization.SignRequest(request)

		trace := logging.FromContext(ctx)
		span := trace.NewHTTPRemoteChild(request)
		response, err = httpClient.Do(request)
		if err != nil {
			span.Finish()
			return err
		}
		span.Finish(logging.WithResponse(response))
		if r.idempotent && response.StatusCode >= 500 {
			return vax.RetryOnResponse(response)
		}
		if response.StatusCode > 299 {
			return parseError(response)
		}
		return nil
	}, client.CallOption)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (client BaseClient) buildRequest(ctx context.Context, method string, path string, params map[string]interface{}) (*http.Request, error) {
	url := client.RootURL + path
	var body io.Reader
	if method == http.MethodPost {
		jsonData, err := json.Marshal(params)
		if err != nil {
			return nil, err
		}
		body = bytes.NewBuffer(jsonData)
		params = nil
	}
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	request = request.WithContext(ctx)

	if method == http.MethodPost {
		request.Header.Set("Content-Type", "application/json")
	}

	if method == http.MethodGet {
		request = client.buildQueryString(request, params)
	}
	return request, nil
}

func (client BaseClient) buildQueryString(request *http.Request, params map[string]interface{}) *http.Request {
	q := request.URL.Query()
	for k, v := range params {
		list, isList := v.([]string)
		if isList {
			for _, lv := range list {
				q.Add(k, string(lv))
			}
		} else {
			item, isString := v.(string)
			if isString {
				q.Add(k, string(item))
			}
		}
	}
	request.URL.RawQuery = q.Encode()
	return request
}

var defaultCallOption = vax.WithRetry(func() vax.Retryer {
	return vax.OnCodes(nil, vax.Backoff{
		Initial:    time.Millisecond * 100,
		Max:        time.Second,
		Multiplier: 2,
	})
})

// ConvertTimeToVAPITimestamp converts a time.Time to the string representation that vapi expects
func ConvertTimeToVAPITimestamp(in time.Time) string {
	if !in.IsZero() {
		return in.Format(time.RFC3339)
	}
	return ""
}

// ParseCursorFromVAPINextQueryString returns the cursor from a vapi paged response nextQueryString
func ParseCursorFromVAPINextQueryString(qs string) (string, error) {
	m, err := url.ParseQuery(qs)
	cursors, ok := m["cursor"]
	if err != nil {
		return "", errors.New("Error parsing cursor")
	}
	if !ok {
		return "", nil
	}
	return cursors[0], nil
}
