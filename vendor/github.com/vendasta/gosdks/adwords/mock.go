package adwords

import (
	"github.com/stretchr/testify/mock"
	"github.com/vendasta/gosdks/pb/adwords_service/v1"
	"golang.org/x/net/context"
)

type MockAdwordsClient struct {
	mock.Mock
}

func (m *MockAdwordsClient) GetAccountPerformance(
	ctx context.Context,
	oauthRefreshToken string,
	customerID int64,
) (Stats, error) {
	args := m.Called(ctx, oauthRefreshToken, customerID)
	return args.Get(0).(Stats), args.Error(1)
}

func (m *MockAdwordsClient) ListAllAccounts(
	ctx context.Context,
	oauthRefreshToken string,
) ([]Account, error) {
	args := m.Called(ctx, oauthRefreshToken)
	return args.Get(0).([]Account), args.Error(1)
}

func (m *MockAdwordsClient) GetKeywordSuggestions(
	ctx context.Context,
	request *adwords_v1.GetKeywordSuggestionsRequest,
) (*adwords_v1.GetKeywordSuggestionsResponse, error) {
	args := m.Called(ctx, request)
	return args.Get(0).(*adwords_v1.GetKeywordSuggestionsResponse), args.Error(1)
}

func (m *MockAdwordsClient) GetCampaignPerformanceEstimation(
	ctx context.Context,
	request *adwords_v1.GetCampaignPerformanceEstimationRequest,
) (*adwords_v1.GetCampaignPerformanceEstimationResponse, error) {
	args := m.Called(ctx, request)
	return args.Get(0).(*adwords_v1.GetCampaignPerformanceEstimationResponse), args.Error(1)
}
