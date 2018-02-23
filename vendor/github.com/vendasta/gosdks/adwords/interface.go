package adwords

import (
	"github.com/vendasta/gosdks/pb/adwords_service/v1"
	"golang.org/x/net/context"
)

// Interface for the Adwords SDK
type Interface interface {
	KeywordsService
	ReportsService
	AccountsService
}

type KeywordsService interface {
	// Request keyword suggestions for a business
	GetKeywordSuggestions(
		ctx context.Context,
		request *adwords_v1.GetKeywordSuggestionsRequest,
	) (*adwords_v1.GetKeywordSuggestionsResponse, error)

	// Request performance estimations For a Campaign of Keywords
	GetCampaignPerformanceEstimation(
		ctx context.Context,
		request *adwords_v1.GetCampaignPerformanceEstimationRequest,
	) (*adwords_v1.GetCampaignPerformanceEstimationResponse, error)
}

type ReportsService interface {
	GetAccountPerformance(
		ctx context.Context,
		oauthRefreshToken string,
		customerID int64,
	) (Stats, error)
}

type AccountsService interface {
	ListAllAccounts(
		ctx context.Context,
		oauthRefreshToken string,
	) ([]Account, error)
}
