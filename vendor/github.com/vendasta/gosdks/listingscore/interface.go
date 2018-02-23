package listingscore

import (
	"github.com/vendasta/gosdks/pb/listing_score/v1"
	"golang.org/x/net/context"
)

// Interface to listing-score microservice API
type Interface interface {
	GetScore(ctx context.Context, accountGroupID string) (*listingscore_v1.GetScoreResponse, error)
	GetIndustryStats(ctx context.Context, taxonomy string, country string, state string, city string) (*listingscore_v1.GetIndustryStatsResponse, error)
	GetHistoricalScoreData(ctx context.Context, accountGroupID string) (*listingscore_v1.GetHistoricalScoreDataResponse, error)
	GetListingPresence(ctx context.Context, accountGroupID string) (*listingscore_v1.GetListingPresenceResponse, error)
	GetListingAccuracy(ctx context.Context, accountGroupID string) (*listingscore_v1.GetListingAccuracyResponse, error)
}
