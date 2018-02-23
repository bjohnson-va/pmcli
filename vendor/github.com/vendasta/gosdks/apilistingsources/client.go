package apilistingsources

import (

	"github.com/vendasta/gosdks/config"
	"google.golang.org/grpc"
	"github.com/vendasta/gosdks/pb/listing_search/v1"
	"github.com/vendasta/gosdks/vax"
	"time"
	"google.golang.org/grpc/codes"
	"golang.org/x/net/context"
)

// Location is a holder for search results from the corresponding api that a search is for
type Location struct {
	// The id that represents the location within the underlying service
	ID string
	// The name of the location
	CompanyName string
	// The postal address of the location
	Address string
	// This typically refers to a city, but may refer to a suburb or a neighborhood in certain countries
	Locality string
	// The state or territory abbreviation
	State string
	// The country code
	Country string
	// A string specifying the zip code or postal code
	ZipCode string
	// A location on the Earth specified by a latitude and longitude
	Point *Geo
	// The website of the location
	Website string
	// The primary phone number of the location
	Phone string
	// The url location of where you can find the location on the internet
	URL string
}

// Geographical location
type Geo struct {
	Latitude  float64
	Longitude float64
}

var addresses = map[config.Env]string{
	config.Local: "domain:11000",
	config.Test:  "api-listing-sources-api-test.vendasta-internal.com:443",
	config.Demo:  "api-listing-sources-api-demo.vendasta-internal.com:443",
	config.Prod:  "api-listing-sources-api-prod.vendasta-internal.com:443",
}

var scopes = map[config.Env]string{
	config.Local: "",
	config.Test:  "https://api-listing-sources-api-test.vendasta-internal.com",
	config.Demo:  "https://api-listing-sources-api-demo.vendasta-internal.com",
	config.Prod:  "https://api-listing-sources-api-prod.vendasta-internal.com",
}

type Interface interface {
	Search(ctx context.Context, companyName string, address string, locality string, state string, longitude float64, latitude float64) ([]*Location, error)
	Get(ctx context.Context, ID string) (*Location, error)
}

type apiListingSourcesClient struct {
	client listingsearch_v1.SearchAdapterClient
}

// NewClient returns a new api-listing-source client
func NewClient(ctx context.Context, e config.Env, dialOptions ...grpc.DialOption) (Interface, error){
	address := addresses[e]
	scope := scopes[e]
	conn, err := vax.NewGRPCConnection(ctx, address,true, scope, true, dialOptions...)
	if err != nil {
		return nil, err
	}
	return &apiListingSourcesClient{
		client: listingsearch_v1.NewSearchAdapterClient(conn),
	}, nil
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

func toLocation(location *listingsearch_v1.Location) *Location {
	l := &Location{
		ID:          location.Id,
		CompanyName: location.CompanyName,
		Address:     location.Address,
		Locality:    location.Locality,
		State:       location.State,
		Country:     location.Country,
		ZipCode:     location.ZipCode,
		Website:     location.Website,
		Phone:       location.Phone,
		URL:         location.Url,
	}
	if location.Point != nil {
		l.Point = &Geo{
			Latitude:  location.Point.Latitude,
			Longitude: location.Point.Longitude,
		}
	}
	return l
}

func (lc *apiListingSourcesClient) Search(ctx context.Context, companyName string, address string, locality string, state string, longitude float64, latitude float64) ([]*Location, error) {
	var resp *listingsearch_v1.SearchResponse
	err := vax.Invoke(ctx, func(ctx context.Context, cs vax.CallSettings) error {
		var err error
		resp, err = lc.client.Search(ctx,
			&listingsearch_v1.SearchRequest{
				CompanyName: companyName,
				Address:     address,
				Locality:    locality,
				State:       state,
				Location: &listingsearch_v1.Geo{
					Longitude: longitude,
					Latitude: latitude,
				},
			},
			grpc.FailFast(false),
		)
		return err
	}, defaultRetryCallOptions)
	if err != nil {
		return nil, err
	}
	respLocation := resp.GetLocations()
	locations := make([]*Location, len(respLocation))
	for i, location := range respLocation {
		locations[i] = toLocation(location)

	}
	return locations, nil

}
func (lc *apiListingSourcesClient) Get(ctx context.Context, ID string) (*Location, error) {
	var resp *listingsearch_v1.GetResponse
	err := vax.Invoke(ctx, func(ctx context.Context, cs vax.CallSettings) error {
		var err error
		resp, err = lc.client.Get(ctx,
			&listingsearch_v1.GetRequest{
				Id: ID,
			},
			grpc.FailFast(false),
		)
		return err
	}, defaultRetryCallOptions)
	if err != nil {
		return nil, err
	}
	l := toLocation(resp.GetLocation())
	return l, nil
}