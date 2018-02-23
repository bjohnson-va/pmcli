package accountgroup

import (
	"github.com/stretchr/testify/mock"
	"golang.org/x/net/context"
)

type MockAccountGroupClient struct {
	mock.Mock
}

func (m *MockAccountGroupClient) GetMulti(ctx context.Context, accountGroupIDs []string, ro ...Option) ([]*AccountGroup, error) {
	args := m.Called(ctx, accountGroupIDs, ro)
	return args.Get(0).([]*AccountGroup), args.Error(1)
}

func (m *MockAccountGroupClient) Get(ctx context.Context, accountGroupID string, ro ...Option) (*AccountGroup, error) {
	args := m.Called(ctx, accountGroupID, ro)
	return args.Get(0).(*AccountGroup), args.Error(1)
}

func (m *MockAccountGroupClient) Lookup(ctx context.Context, lo ...Option) (*PagedResult, error) {
	args := m.Called(ctx, lo)
	return args.Get(0).(*PagedResult), args.Error(1)
}

func (m *MockAccountGroupClient) Create(ctx context.Context, socialProfileGroupID string, napData *NAPData, externalIds *ExternalIdentifiers, opts ...MutateOption) (string, error) {
	args := m.Called(ctx, socialProfileGroupID, napData, externalIds, opts)
	return args.String(0), args.Error(1)
}

func (m *MockAccountGroupClient) Update(ctx context.Context, accountGroupID string, opts ...MutateOption) error {
	args := m.Called(ctx, accountGroupID, opts)
	return args.Error(0)
}

func (m *MockAccountGroupClient) Delete(ctx context.Context, accountGroupID string) error {
	args := m.Called(ctx, accountGroupID)
	return args.Error(0)
}
