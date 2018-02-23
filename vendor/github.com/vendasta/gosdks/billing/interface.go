package billing

import (
	"time"

	"golang.org/x/net/context"
)

// Interface to billing microservice API
type Interface interface {
	CreateBillableItem(ctx context.Context, merchantID string, sku string, customerID string, orderID string, expiry time.Time, billingStart time.Time) error
	ExpireBillableItem(ctx context.Context, merchantID string, sku string, customerID string, orderID string, expiry time.Time) error
}
