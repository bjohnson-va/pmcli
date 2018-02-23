package adwords

import (
	"github.com/vendasta/gosdks/pb/adwords_service/v1"
	"github.com/vendasta/gosdks/util"
	"github.com/vendasta/gosdks/validation"
	"github.com/vendasta/gosdks/vax"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// GetAccountPerformance gets a performance report for the given account
func (c *client) GetAccountPerformance(
	ctx context.Context,
	oauthRefreshToken string,
	customerID int64,
) (Stats, error) {
	ctx = metadata.NewOutgoingContext(ctx, nil)

	err := validation.NewValidator().Rule(
		validation.StringNotEmpty(oauthRefreshToken, util.InvalidArgument, "oauthRefreshToken is required"),
		validation.IntGreaterThan(customerID, 0, util.InvalidArgument, "customerID is required"),
	).ValidateAndJoinErrors()
	if err != nil {
		return Stats{}, err
	}

	request := &adwords_v1.GetAccountPerformanceRequest{
		CustomerId:        customerID,
		OauthRefreshToken: oauthRefreshToken,
	}

	var response *adwords_v1.GetAccountPerformanceResponse
	err = vax.Invoke(ctx, func(ctx context.Context, cs vax.CallSettings) error {
		var err error
		response, err = c.ReportsClient.GetAccountPerformance(ctx,
			request,
			grpc.FailFast(false),
		)
		return err
	}, defaultRetryCallOptions)

	if err != nil {
		return Stats{}, err
	}
	stats := fromStatsProto(response)
	return stats, nil
}

// Stats are statistics pertaining to AdWords campaigns or an account
type Stats struct {
	Clicks           int64
	CostMicroDollars int64
	Impressions      int64
	Conversions      float64
	AllConversions   float64
}

func fromStatsProto(response *adwords_v1.GetAccountPerformanceResponse) Stats {
	return Stats{
		Clicks:           response.Clicks,
		AllConversions:   response.AllConversions,
		Conversions:      response.Conversions,
		CostMicroDollars: response.Cost,
		Impressions:      response.Impressions,
	}
}
