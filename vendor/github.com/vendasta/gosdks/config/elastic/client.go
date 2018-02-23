package elasticclient

import (
	"net/http"

	"github.com/vendasta/gosdks/config"
	"github.com/vendasta/gosdks/logging"
	"golang.org/x/net/context"
	"gopkg.in/olivere/elastic.v3"
)

var elasticClient *elastic.Client

// Get returns the elastic Client for this instance.
func Get() *elastic.Client {
	return elasticClient
}

// ElasticLogger is an elastic search-conformant logger capable of utilizing our own context-based logger
type ElasticLogger struct {
	log func(context.Context, string, ...interface{})
}

// Printf invokes the elastic logger
func (e ElasticLogger) Printf(format string, v ...interface{}) {
	e.log(context.Background(), format, v...)
}

// Initialize the elastic client, must be called before attempting to use the client
func Initialize() error {
	logging.Debugf(context.Background(), "Initializing elastic Client...")
	var err error

	transport := &elasticTransport{
		APIKey:    GetElasticsearchAPIKey(),
		Transport: http.DefaultTransport,
	}

	elasticClient, err = elastic.NewClient(
		elastic.SetURL(GetElasticsearchURL()),
		elastic.SetSniff(false),
		elastic.SetScheme("https"),
		elastic.SetHttpClient(&http.Client{Transport: transport}),
		elastic.SetHealthcheck(false),
		elastic.SetMaxRetries(2),
		elastic.SetHealthcheckTimeoutStartup(0),
		elastic.SetSendGetBodyAs("POST"),
		elastic.SetErrorLog(ElasticLogger{log: logging.Errorf}),
		elastic.SetInfoLog(ElasticLogger{log: logging.Infof}),
		elastic.SetTraceLog(ElasticLogger{log: logging.Debugf}),
	)
	return err
}

type elasticTransport struct {
	Transport http.RoundTripper
	APIKey    string
}

func (e *elasticTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	rt := e.Transport
	req.Header.Add("x-api-key", e.APIKey)
	return rt.RoundTrip(req)
}

// GetElasticsearchURL returns the elasticesearch url based on the environment.
func GetElasticsearchURL() string {
	switch config.Getenvironment() {
	case "prod":
		return "https://elasticsearch-prod.vendasta-internal.com"
	case "demo":
		return "https://elasticsearch-demo.vendasta-internal.com"
	case "test":
		return "https://elasticsearch-test.vendasta-internal.com"
	}
	return "https://elasticsearch-test.vendasta-internal.com"
}

// GetElasticsearchAPIKey returns the elasticsearch api key based on the environment.
func GetElasticsearchAPIKey() string {
	switch config.Getenvironment() {
	case "prod":
		return "AIzaSyDAfizKwYN_c4YeqaZIO-T0RWcgOzA_d2k"
	case "demo":
		return "AIzaSyCtuSyDpIx5TZ3ivs46SnQlDKL1IPTcFZI"
	case "test":
		return "AIzaSyDyJbP9406qmnmC2WRW1UStMVUCsLtpYBY"
	}
	return "AIzaSyDyJbP9406qmnmC2WRW1UStMVUCsLtpYBY"
}
