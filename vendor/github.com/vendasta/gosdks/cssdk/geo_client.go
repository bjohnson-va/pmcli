package cssdk

import (
	"github.com/vendasta/gosdks/basesdk"
	"github.com/vendasta/gosdks/config"
	"golang.org/x/net/context"
)

const (
	geoInferPath = "/api/v1/geo/get/"
)

// GeoClientInterface exposes methods for inferring Geopoints from Core Services
type GeoClientInterface interface {
	InferGeolocation(ctx context.Context, country string, opts ...InferGeolocationOption) (*GeoPoint, error)
}

// GeoClient is a client which handles calls to core services's Geo apis.
// Implements the GeoClientInterface
type GeoClient struct {
	basesdk.SDKClient
}

type inferenceOptions struct {
	country string
	state   string
	city    string
	address string
}

// InferGeolocationOption allows clients to optionally apply more specific address information to their request to make the inferred location more accurate.
type InferGeolocationOption func(o *inferenceOptions)

// WithState uses the specified state as an input to an inference API
func WithState(state string) InferGeolocationOption {
	return func(o *inferenceOptions) {
		o.state = state
	}
}

// WithCity uses the specified city as an input to an inference API
func WithCity(city string) InferGeolocationOption {
	return func(o *inferenceOptions) {
		o.city = city
	}
}

// WithAddress uses the specified address as an input to an inference API
func WithAddress(address string) InferGeolocationOption {
	return func(o *inferenceOptions) {
		o.address = address
	}
}

// BuildGeoClient creates a geo client.
func BuildGeoClient(apiUser string, apiKey string, env config.Env) *GeoClient {
	auth := basesdk.UserKey{APIUser: apiUser, APIKey: apiKey}
	baseClient := basesdk.BaseClient{Authorization: auth, RootURL: rootURLFromEnv(env)}
	return &GeoClient{baseClient}
}

//List returns a list of all the taxonomies that Core Services knows about
func (c *GeoClient) InferGeolocation(ctx context.Context, country string, opts ...InferGeolocationOption) (*GeoPoint, error) {
	o := &inferenceOptions{country: country}
	for _, f := range opts {
		f(o)
	}
	params := map[string]interface{}{
		"country": o.country,
	}
	if o.state != "" {
		params["state"] = o.state
	}
	if o.city != "" {
		params["city"] = o.city
	}
	if o.address != "" {
		params["address"] = o.address
	}
	response, err := c.Get(ctx, geoInferPath, params, basesdk.Idempotent())
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	return geopointFromResponse(response.Body)
}
