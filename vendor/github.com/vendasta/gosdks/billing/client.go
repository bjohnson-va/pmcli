package billing

import (
	"fmt"
	"time"

	"github.com/golang/protobuf/ptypes"

	"github.com/vendasta/gosdks/pb/billing/v1"

	"github.com/vendasta/gosdks/config"
	"github.com/vendasta/gosdks/vax"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type client struct {
	billing_v1.BillingClient
}

// NewClient returns a new billing API client object
func NewClient(ctx context.Context, e config.Env, dialOptions ...grpc.DialOption) (Interface, error) {
	address := addresses[e]
	if address == "" {
		return nil, fmt.Errorf("Unable to create client with environment %d", e)
	}
	connection, err := vax.NewGRPCConnection(ctx, address, e != config.Local, scopes[e], true, dialOptions...)
	if err != nil {
		return nil, err
	}
	return &client{billing_v1.NewBillingClient(connection)}, nil
}

// CreateBillableItem creates a billable item with an optional expiry date and billing start date
func (c *client) CreateBillableItem(ctx context.Context, merchantID string, sku string, customerID string, orderID string, expiry time.Time, billingStart time.Time) error {
	ctx = metadata.NewOutgoingContext(ctx, nil)

	request := &billing_v1.CreateBillableItemRequest{
		MerchantId: merchantID,
		Sku:        sku,
		CustomerId: customerID,
		OrderId:    orderID,
	}

	if !expiry.IsZero() {
		expiryTimestamp, err := ptypes.TimestampProto(expiry)
		if err != nil {
			return err
		}
		request.Expiry = expiryTimestamp
	}

	if !billingStart.IsZero() {
		billingStartTimestamp, err := ptypes.TimestampProto(billingStart)
		if err != nil {
			return err
		}
		request.BillingStart = billingStartTimestamp
	}

	err := vax.Invoke(ctx, func(ctx context.Context, cs vax.CallSettings) error {
		var err error
		_, err = c.BillingClient.CreateBillableItem(ctx,
			request,
			grpc.FailFast(false),
		)
		return err
	}, defaultRetryCallOptions)

	if err != nil {
		return err
	}
	return nil
}

// ExpireBillableItem expires an existing billable item with an optionally given expiry date
func (c *client) ExpireBillableItem(ctx context.Context, merchantID string, sku string, customerID string, orderID string, expiry time.Time) error {
	ctx = metadata.NewOutgoingContext(ctx, nil)

	request := &billing_v1.ExpireBillableItemRequest{
		MerchantId: merchantID,
		Sku:        sku,
		CustomerId: customerID,
		OrderId:    orderID,
	}

	if !expiry.IsZero() {
		expiryTimestamp, err := ptypes.TimestampProto(expiry)
		if err != nil {
			return err
		}
		request.Expiry = expiryTimestamp
	}

	err := vax.Invoke(ctx, func(ctx context.Context, cs vax.CallSettings) error {
		var err error
		_, err = c.BillingClient.ExpireBillableItem(ctx,
			request,
			grpc.FailFast(false),
		)
		return err
	}, defaultRetryCallOptions)

	if err != nil {
		return err
	}
	return nil
}
