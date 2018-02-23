package reserveidsdk

import (
	"golang.org/x/net/context"
)

type ReserveIDMock struct {
	ReserveIDFunc func(ctx context.Context, partnerID string, businessID string, customerID string, appID string) (*ReserveIDResponse, error)
}

func (m *ReserveIDMock) ReserveID(ctx context.Context, partnerID string, businessID string, customerID string, appID string) (*ReserveIDResponse, error) {
	if m.ReserveIDFunc != nil {
		return m.ReserveIDFunc(ctx, partnerID, businessID, customerID, appID)
	}
	return nil, nil
}
