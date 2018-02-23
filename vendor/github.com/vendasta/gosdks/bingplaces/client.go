package bingplaces

import (
	"golang.org/x/net/context"
	"github.com/vendasta/gosdks/config"
	"google.golang.org/grpc"
	"github.com/vendasta/gosdks/pb/listing_search/v1"
	"github.com/vendasta/gosdks/vax"
	"google.golang.org/grpc/codes"
	"time"
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
	config.Test:  "bing-places-api-test.vendasta-internal.com:443",
	config.Demo:  "bing-places-api-demo.vendasta-internal.com:443",
	config.Prod:  "bing-places-api-prod.vendasta-internal.com:443",
}

var scopes = map[config.Env]string{
	config.Local: "",
	config.Test:  "https://bing-places-api-test.vendasta-internal.com",
	config.Demo:  "https://bing-places-api-demo.vendasta-internal.com",
	config.Prod:  "https://bing-places-api-prod.vendasta-internal.com",
}

// NewClient returns a Bing Places search client.
func NewClient(ctx context.Context, e config.Env, dialOptions ...grpc.DialOption) (Interface, error) {
	address := addresses[e]
	scope := scopes[e]
	conn, err := vax.NewGRPCConnection(ctx, address, true, scope, true, dialOptions...)
	if err != nil {
		return nil, err
	}
	return &bingClient{
		client: listingsearch_v1.NewSearchAdapterClient(conn),
	}, nil
}

// Interface defines all of the API methods available from Bing Places
type Interface interface {
	Search(ctx context.Context, companyName string, address string, locality string, state string) ([]*Location, error)
	Get(ctx context.Context, bingID string) (*Location, error)
}

type bingClient struct {
	client listingsearch_v1.SearchAdapterClient
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

// Search returns the found locations by given search terms
func (bc *bingClient) Search(ctx context.Context, companyName string, address string, locality string, state string) ([]*Location, error) {
	var resp *listingsearch_v1.SearchResponse
	err := vax.Invoke(ctx, func(ctx context.Context, cs vax.CallSettings) error {
		var err error
		resp, err = bc.client.Search(ctx,
			&listingsearch_v1.SearchRequest{
				CompanyName: companyName,
				Address:     address,
				Locality:    locality,
				State:       state,
			},
			grpc.FailFast(false),
		)
		return err
	}, defaultRetryCallOptions)
	if err != nil {
		return nil, err
	}
	respLocations := resp.GetLocations()
	locations := make([]*Location, len(respLocations))
	for i, location := range respLocations {
		locations[i] = toLocation(location)
	}
	return locations, nil
}

// Get returns the found location by given id
func (bc *bingClient) Get(ctx context.Context, bingID string) (*Location, error) {
	var resp *listingsearch_v1.GetResponse
	err := vax.Invoke(ctx, func(ctx context.Context, cs vax.CallSettings) error {
		var err error
		resp, err = bc.client.Get(ctx,
			&listingsearch_v1.GetRequest{
				Id: bingID,
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
