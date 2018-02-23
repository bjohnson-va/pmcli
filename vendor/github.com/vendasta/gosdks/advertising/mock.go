package advertising

import (
	"github.com/stretchr/testify/mock"
	"github.com/vendasta/gosdks/pb/advertising/v1"
	"golang.org/x/net/context"
)

// MockAdvertisingClient mocks campaign
type MockAdvertisingClient struct {
	mock.Mock
}

// CreateOrderEvent mocks CreateOrderEvent() in Advertising
func (m *MockAdvertisingClient) CreateOrderEvent(ctx context.Context, order *OrderEvent) error {
	args := m.Called(ctx, order)
	return args.Error(0)
}

// CreateLifelineUpdateEvent mocks CreateLifelineUpdateEvent() in Advertising
func (m *MockAdvertisingClient) CreateLifelineUpdateEvent(ctx context.Context, request *advertising_v1.CreateLifelineUpdateEventRequest) error {
	args := m.Called(ctx, request)
	return args.Error(0)
}

// StoreCredentialsForBusiness mocks StoreCredentialsForBusiness() in Advertising
func (m *MockAdvertisingClient) StoreCredentialsForBusiness(ctx context.Context, oauthRefreshToken string, businessID string) error {
	args := m.Called(ctx, oauthRefreshToken, businessID)
	return args.Error(0)
}
