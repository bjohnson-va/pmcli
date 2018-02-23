package listingscore

import (
	"fmt"

	"github.com/vendasta/gosdks/config"
	"github.com/vendasta/gosdks/pb/listing_score/v1"
	"github.com/vendasta/gosdks/vax"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// NewClient returns a new listing-score API client object
func NewClient(ctx context.Context, e config.Env, dialOptions ...grpc.DialOption) (Interface, error) {
	address := addresses[e]
	if address == "" {
		return nil, fmt.Errorf("Unable to create client with environment %d", e)
	}
	connection, err := vax.NewGRPCConnection(ctx, address, e != config.Local, scopes[e], true, dialOptions...)
	if err != nil {
		return nil, err
	}
	return &client{listingscore_v1.NewListingScoreClient(connection)}, nil
}

type client struct {
	listingscore_v1.ListingScoreClient
}

func (c *client) GetScore(ctx context.Context, accountGroupID string) (*listingscore_v1.GetScoreResponse, error) {
	ctx = metadata.NewOutgoingContext(ctx, nil)

	var response *listingscore_v1.GetScoreResponse
	request := &listingscore_v1.GetScoreRequest{AccountGroupId: accountGroupID}
	err := vax.Invoke(ctx, func(ctx context.Context, cs vax.CallSettings) error {
		var err error
		response, err = c.ListingScoreClient.GetScore(ctx,
			request,
			grpc.FailFast(false),
		)
		return err
	}, defaultRetryCallOptions)

	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *client) GetIndustryStats(ctx context.Context, taxonomy string, country string, state string, city string) (*listingscore_v1.GetIndustryStatsResponse, error) {
	ctx = metadata.NewOutgoingContext(ctx, nil)

	var response *listingscore_v1.GetIndustryStatsResponse
	request := &listingscore_v1.GetIndustryStatsRequest{TaxonomyId: taxonomy, Country: country, State: state, City: city}
	err := vax.Invoke(ctx, func(ctx context.Context, cs vax.CallSettings) error {
		var err error
		response, err = c.ListingScoreClient.GetIndustryStats(ctx,
			request,
			grpc.FailFast(false),
		)
		return err
	}, defaultRetryCallOptions)

	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *client) GetHistoricalScoreData(ctx context.Context, accountGroupID string) (*listingscore_v1.GetHistoricalScoreDataResponse, error) {
	ctx = metadata.NewOutgoingContext(ctx, nil)

	var response *listingscore_v1.GetHistoricalScoreDataResponse
	request := &listingscore_v1.GetHistoricalScoreDataRequest{AccountGroupId: accountGroupID}
	err := vax.Invoke(ctx, func(ctx context.Context, cs vax.CallSettings) error {
		var err error
		response, err = c.ListingScoreClient.GetHistoricalScoreData(ctx,
			request,
			grpc.FailFast(false),
		)
		return err
	}, defaultRetryCallOptions)

	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *client) GetListingPresence(ctx context.Context, accountGroupID string) (*listingscore_v1.GetListingPresenceResponse, error) {
	ctx = metadata.NewOutgoingContext(ctx, nil)

	var response *listingscore_v1.GetListingPresenceResponse
	request := &listingscore_v1.GetListingPresenceRequest{AccountGroupId: accountGroupID}
	err := vax.Invoke(ctx, func(ctx context.Context, cs vax.CallSettings) error {
		var err error
		response, err = c.ListingScoreClient.GetListingPresence(ctx,
			request,
			grpc.FailFast(false),
		)
		return err
	}, defaultRetryCallOptions)

	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *client) GetListingAccuracy(ctx context.Context, accountGroupID string) (*listingscore_v1.GetListingAccuracyResponse, error) {
	ctx = metadata.NewOutgoingContext(ctx, nil)

	var response *listingscore_v1.GetListingAccuracyResponse
	request := &listingscore_v1.GetListingAccuracyRequest{AccountGroupId: accountGroupID}
	err := vax.Invoke(ctx, func(ctx context.Context, cs vax.CallSettings) error {
		var err error
		response, err = c.ListingScoreClient.GetListingAccuracy(ctx,
			request,
			grpc.FailFast(false),
		)
		return err
	}, defaultRetryCallOptions)

	if err != nil {
		return nil, err
	}
	return response, nil
}
