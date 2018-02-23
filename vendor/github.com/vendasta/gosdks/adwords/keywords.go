package adwords

import (
	"github.com/vendasta/gosdks/pb/adwords_service/v1"
	"github.com/vendasta/gosdks/vax"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// GetKeywordSuggestions returns a set of possible keywords related to the business
func (c *client) GetKeywordSuggestions(
	ctx context.Context,
	request *adwords_v1.GetKeywordSuggestionsRequest,
) (*adwords_v1.GetKeywordSuggestionsResponse, error) {
	ctx = metadata.NewOutgoingContext(ctx, nil)
	var response *adwords_v1.GetKeywordSuggestionsResponse
	err := vax.Invoke(ctx, func(ctx context.Context, cs vax.CallSettings) error {
		var err error
		response, err = c.KeywordsClient.GetKeywordSuggestions(ctx,
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

// GetCampaignPerformanceEstimation returns campaign performance estimations for a set of keywords
func (c *client) GetCampaignPerformanceEstimation(
	ctx context.Context,
	request *adwords_v1.GetCampaignPerformanceEstimationRequest,
) (*adwords_v1.GetCampaignPerformanceEstimationResponse, error) {
	ctx = metadata.NewOutgoingContext(ctx, nil)
	var response *adwords_v1.GetCampaignPerformanceEstimationResponse
	err := vax.Invoke(ctx, func(ctx context.Context, cs vax.CallSettings) error {
		var err error
		response, err = c.KeywordsClient.GetCampaignPerformanceEstimation(ctx,
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
