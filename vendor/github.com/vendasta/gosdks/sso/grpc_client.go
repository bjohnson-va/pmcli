package sso

import (
	"time"

	"github.com/vendasta/gosdks/pb/sso/v1"
	"github.com/vendasta/gosdks/util"
	"github.com/vendasta/gosdks/vax"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

// GRPCClient a GRPC client for the marketplace api microservice.
type GRPCClient struct {
	client sso_v1.ServiceProviderClient
}

// NewGRPCClient returns an new GRPC client for the marketplace api microservice.
func NewGRPCClient(ctx context.Context, address string, useTLS bool, scope string, dialOptions ...grpc.DialOption) (Client, error) {
	conn, err := vax.NewGRPCConnection(ctx, address, useTLS, scope, true, dialOptions...)
	if err != nil {
		return nil, err
	}
	return &GRPCClient{client: sso_v1.NewServiceProviderClient(conn)}, nil
}

// defaultRetryCallOptions controls the errors that we will automatically retry on. This is due to the case where the
// server has given us an error that is deemed retry-able.
var defaultRetryCallOptions = vax.WithRetry(func() vax.Retryer {
	return vax.OnCodes([]codes.Code{
		codes.DeadlineExceeded,
		codes.Unavailable,
		codes.Unknown,
	}, vax.Backoff{
		Initial:    10 * time.Millisecond,
		Max:        300 * time.Millisecond,
		Multiplier: 3,
	})
})

// GetSessionTransferURL is for requesting the session transfer URL for a business.
func (c *GRPCClient) GetSessionTransferURL(ctx context.Context, serviceProviderID string, serviceContext ServiceContext) (sessionTransferURL string, err error) {
	req := &sso_v1.GetSessionTransferURLRequest{
		ServiceProviderId: serviceProviderID,
		Context:           &sso_v1.ServiceContext{},
	}
	switch v := serviceContext.(type) {
	case AccountContext:
		req.Context.Context = &sso_v1.ServiceContext_Account_{
			Account: &sso_v1.ServiceContext_Account{
				AccountId: v.AccountID,
			},
		}
	default:
		return "", util.Error(util.InvalidArgument, "unknown service context type provided")
	}

	var res *sso_v1.GetSessionTransferURLResponse
	err = vax.Invoke(ctx, func(ctx context.Context, settings vax.CallSettings) error {
		var err error
		ctx, _ = context.WithTimeout(ctx, 3*time.Second)
		res, err = c.client.GetSessionTransferURL(ctx, req, grpc.FailFast(false))
		return err
	}, defaultRetryCallOptions)
	if err != nil {
		return "", err
	}

	return res.GetSessionTransferUrl(), nil
}
