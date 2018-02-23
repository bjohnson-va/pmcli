package cssdk

import (
	"github.com/vendasta/gosdks/basesdk"
	"github.com/vendasta/gosdks/config"
	"golang.org/x/net/context"
)

const (
	taxonomyListPath = "/api/v1/vtax/get/"
)

// TaxonomyClientInterface exposes methods for interacting with taxonomy definitions from Core Services
type TaxonomyClientInterface interface {
	List(ctx context.Context) ([]*Taxonomy, error)
}

// TaxonomyClient is a client which handles calls to core services's Taxonomy apis.
// Implements the TaxonomyClientInterface
type TaxonomyClient struct {
	basesdk.SDKClient
}

// BuildTaxonomyClient creates a taxonomy client.
func BuildTaxonomyClient(apiUser string, apiKey string, env config.Env) TaxonomyClientInterface {
	auth := basesdk.UserKey{APIUser: apiUser, APIKey: apiKey}
	baseClient := basesdk.BaseClient{Authorization: auth, RootURL: rootURLFromEnv(env)}
	return &TaxonomyClient{baseClient}
}

//List returns a list of all the taxonomies that Core Services knows about
func (c *TaxonomyClient) List(ctx context.Context) ([]*Taxonomy, error) {
	params := map[string]interface{}{}
	response, err := c.Get(ctx, taxonomyListPath, params, basesdk.Idempotent())
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	return taxonomiesFromResponse(response.Body)
}
