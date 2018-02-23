package billingsdk

import (
	"time"

	"golang.org/x/net/context"
)

// BillingClientMock exposes methods for interacting with products in billing
type BillingClientMock struct {
	ProductCreateFunc     func(ctx context.Context, product Product, userID string, version int64) error
	ProductUpdateFunc     func(ctx context.Context, product Product, userID string, version int64) error
	ActivateProductFunc   func(ctx context.Context, accountGroupID string, productID string, userID string, opts PurchaseOptions) error
	DeactivateProductFunc func(ctx context.Context, accountGroupID string, productID string, userID string, opts PurchaseOptions) error
	GetLineItemsFunc      func(ctx context.Context, partnerID string, start time.Time, end time.Time, offset int, limit int) (*LineItemResponse, error)
	GetContractInfoFunc   func(ctx context.Context, partnerID string) (*ContractInfo, error)
}

// ProductCreate calls ProductCreateFunc if specified, return nil otherwise
func (m *BillingClientMock) ProductCreate(ctx context.Context, product Product, userID string, version int64) error {
	if m.ProductCreateFunc != nil {
		return m.ProductCreateFunc(ctx, product, userID, version)
	}
	return nil
}

// ProductUpdate calls ProductUpdateFunc if specified, return nil otherwise
func (m *BillingClientMock) ProductUpdate(ctx context.Context, product Product, userID string, version int64) error {
	if m.ProductUpdateFunc != nil {
		return m.ProductUpdateFunc(ctx, product, userID, version)
	}
	return nil
}

// ActivateProduct calls ActivateProductFunc if specified, return nil otherwise
func (m *BillingClientMock) ActivateProduct(ctx context.Context, accountGroupID string, productID string, userID string, opts PurchaseOptions) error {
	if m.ActivateProductFunc != nil {
		return m.ActivateProductFunc(ctx, accountGroupID, productID, userID, opts)
	}
	return nil
}

// DeactivateProduct calls DeactivateProductFunc if specified, return nil otherwise
func (m *BillingClientMock) DeactivateProduct(ctx context.Context, accountGroupID string, productID string, userID string, opts PurchaseOptions) error {
	if m.DeactivateProductFunc != nil {
		return m.DeactivateProductFunc(ctx, accountGroupID, productID, userID, opts)
	}
	return nil
}

// GetLineItems calls GetLineItemsFunc if specified, return nil otherwise
func (m *BillingClientMock) GetLineItems(ctx context.Context, partnerID string, start time.Time, end time.Time, offset int, limit int) (*LineItemResponse, error) {
	if m.GetLineItemsFunc != nil {
		return m.GetLineItemsFunc(ctx, partnerID, start, end, offset, limit)
	}
	return nil, nil
}

// GetContractInfo calls GetContractInfoFunc if specified, return nil otherwise
func (m *BillingClientMock) GetContractInfo(ctx context.Context, partnerID string) (*ContractInfo, error) {
	if m.GetContractInfoFunc != nil {
		return m.GetContractInfoFunc(ctx, partnerID)
	}
	return nil, nil
}
