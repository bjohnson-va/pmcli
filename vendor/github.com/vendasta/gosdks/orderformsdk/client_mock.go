package orderformsdk

import (
	"golang.org/x/net/context"
)

type OrderFormMock struct {
	LockFunc func(ctx context.Context, orderFormSubmissionID string) error
	SetActivationIDFunc func(ctx context.Context, orderFormSubmissionID string, activationID string) error
	CreateFunc func(ctx context.Context, partnerID string, accountGroupID string, productID string, commonFields []OrderField, customFields []OrderField) (string, error)
	GetFunc func(ctx context.Context, orderFormSubmissionId string) (*OrderFormSubmission, error)
}

func (m *OrderFormMock) Lock(ctx context.Context, orderFormSubmissionID string) error {
	if m.LockFunc != nil {
		return m.LockFunc(ctx, orderFormSubmissionID)
	}
	return nil
}

func (m *OrderFormMock) SetActivationID(ctx context.Context, orderFormSubmissionID string, activationID string) error {
	if m.SetActivationIDFunc != nil {
		return m.SetActivationIDFunc(ctx, orderFormSubmissionID, activationID)
	}
	return nil
}

func (m *OrderFormMock) Create(ctx context.Context, partnerID string, accountGroupID string, productID string, commonFields []OrderField, customFields []OrderField) (string, error) {
	if m.CreateFunc != nil {
		return m.CreateFunc(ctx, partnerID, accountGroupID, productID, commonFields, customFields)
	}
	return "OFS-MOCK", nil
}

func (m *OrderFormMock) GetOrderFormSubmission(ctx context.Context, orderFormSubmissionId string) (*OrderFormSubmission, error) {
	if m.GetFunc != nil {
		return m.GetFunc(ctx, orderFormSubmissionId)
	}
	return &OrderFormSubmission{OrderFormSubmissionId: orderFormSubmissionId}, nil
}
