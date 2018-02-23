package marketplace

import (
	"github.com/stretchr/testify/mock"
	"golang.org/x/net/context"
)

// MockActivityClient mocks activity client
type MockActivityClient struct {
	mock.Mock
}

// CreateActivity implements CreateActivity in activity client interface
func (a *MockActivityClient) CreateActivity(ctx context.Context, activity Activity) (*CreateActivityResponse, error) {
	call := a.Called(ctx, activity)
	return call.Get(0).(*CreateActivityResponse), call.Error(1)
}
