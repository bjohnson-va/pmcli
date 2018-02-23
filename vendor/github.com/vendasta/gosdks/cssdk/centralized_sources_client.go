package cssdk

import (
	"github.com/vendasta/gosdks/basesdk"
	"github.com/vendasta/gosdks/config"
	"golang.org/x/net/context"
)

const (
	lookupGlobalPath = "/api/v1/source/lookupGlobal/"
)

// CentralizedSourcesInterface defines the interface of a CentralizedSources client
type CentralizedSourcesInterface interface {
	Lookup(context.Context) ([]*Source, error)
}

type centralizedSourcesClient struct {
	basesdk.SDKClient
}

// BuildCentralizedSourcesClient creates a review client.
func BuildCentralizedSourcesClient(apiUser string, apiKey string, env config.Env) CentralizedSourcesInterface {
	auth := basesdk.UserKey{APIUser: apiUser, APIKey: apiKey}
	baseClient := basesdk.BaseClient{Authorization: auth, RootURL: rootURLFromEnv(env)}
	return centralizedSourcesClient{baseClient}
}

//Lookup fetches a list of sources from CS
func (c centralizedSourcesClient) Lookup(ctx context.Context) ([]*Source, error) {
	response, err := c.Get(ctx, lookupGlobalPath, nil, basesdk.Idempotent())
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	return sourcesFromResponse(response.Body)
}
