package listingsearch

import (
	"github.com/vendasta/gosdks/config"
	"github.com/vendasta/gosdks/pb/listing_search/v1"
	"github.com/vendasta/gosdks/vax"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"time"
	"sync"
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

type Interface interface {
	Search(ctx context.Context, companyName string, address string, locality string, state string, longitude float64, latitude float64, sourceID int64) ([]*Location, error)
	Get(ctx context.Context, ID string, sourceID int64) (*Location, error)
}

type clients struct {
	clientMap map[int64]listingsearch_v1.SearchAdapterClient
	mutex *sync.Mutex
}

func (lc *clients) getClient(ctx context.Context, e config.Env, sourceID int64) (listingsearch_v1.SearchAdapterClient, error) {
	lc.mutex.Lock()
	defer lc.mutex.Unlock()
	client, ok := lc.clientMap[sourceID]
	if ok {
		return client, nil
	}

	conn, err := vax.NewGRPCConnection(ctx, addresses[Source(sourceID)][e], true, scopes[Source(sourceID)][e], true)
	if err != nil {
		return nil, err
	}
	client = listingsearch_v1.NewSearchAdapterClient(conn)

	lc.clientMap[sourceID] = client
	return  client, nil

}

func NewClients(ctx context.Context, e config.Env, dialOptions ...grpc.DialOption) Interface {

	return &clients {
		clientMap: make(map[int64]listingsearch_v1.SearchAdapterClient),
		mutex: &sync.Mutex{},
	}

}

func (lc *clients) Search(ctx context.Context, companyName string, address string, locality string, state string, longitude float64, latitude float64, sourceID int64) ([]*Location, error) {

	client, err := lc.getClient(ctx, config.CurEnv(), sourceID)
	if err != nil {
		return nil, err
	}

	var resp *listingsearch_v1.SearchResponse
	err = vax.Invoke(ctx, func(ctx context.Context, cs vax.CallSettings) error {
		var err error
		resp, err = client.Search(ctx,
			&listingsearch_v1.SearchRequest{
				CompanyName: companyName,
				Address:     address,
				Locality:    locality,
				State:       state,
				Location: &listingsearch_v1.Geo{
					Longitude: longitude,
					Latitude:  latitude,
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

func (lc *clients) Get(ctx context.Context, ID string, sourceID int64) (*Location, error) {
	client, err := lc.getClient(ctx, config.CurEnv(), sourceID)
	if err != nil {
		return nil, err
	}
	var resp *listingsearch_v1.GetResponse
	err = vax.Invoke(ctx, func(ctx context.Context, cs vax.CallSettings) error {
		var err error
		resp, err = client.Get(ctx,
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
