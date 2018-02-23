package marketplaceapi

import (
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/vendasta/gosdks/pb/marketplace/v1"
	"github.com/vendasta/gosdks/vax"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

// GRPCClient a GRPC client for the marketplace api microservice.
type GRPCClient struct {
	client marketplace_v1.MarketplaceClient
}

// NewGRPCClient returns an new GRPC client for the marketplace api microservice.
func NewGRPCClient(ctx context.Context, address string, useTLS bool, scope string, dialOptions ...grpc.DialOption) (Client, error) {
	conn, err := vax.NewGRPCConnection(ctx, address, useTLS, scope, false, dialOptions...)
	if err != nil {
		return nil, err
	}
	return &GRPCClient{client: marketplace_v1.NewMarketplaceClient(conn)}, nil
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

// GetOAuthToken get's an oauth token for an application.
func (c *GRPCClient) GetOAuthToken(ctx context.Context, grantType string, assertion string) (*Token, error) {
	ctx = metadata.NewOutgoingContext(ctx, nil)

	var res *marketplace_v1.GetOAuthTokenResponse
	err := vax.Invoke(ctx, func(ctx context.Context, settings vax.CallSettings) error {
		var err error
		ctx, _ = context.WithTimeout(ctx, 3*time.Second)
		res, err = c.client.GetOAuthToken(ctx, &marketplace_v1.GetOAuthTokenRequest{GrantType: grantType, Assertion: assertion}, grpc.FailFast(false))
		return err
	}, defaultRetryCallOptions)
	if err != nil {
		return nil, err
	}

	exp, err := ptypes.Timestamp(res.Expires)
	if err != nil {
		return nil, err
	}

	return &Token{
		AccessToken: res.AccessToken,
		TokenType:   res.TokenType,
		Expires:     exp,
	}, nil
}

// GetSessionFromToken is for getting the session information from a token
func (c *GRPCClient) GetSessionFromToken(ctx context.Context, accessToken string) (*Session, error) {
	var res *marketplace_v1.GetSessionFromTokenResponse
	err := vax.Invoke(ctx, func(ctx context.Context, settings vax.CallSettings) error {
		var err error
		ctx, _ = context.WithTimeout(ctx, 3*time.Second)
		res, err = c.client.GetSessionFromToken(ctx, &marketplace_v1.GetSessionFromTokenRequest{Token: accessToken}, grpc.FailFast(false))
		return err
	}, defaultRetryCallOptions)
	if err != nil {
		return nil, err
	}

	issued, err := ptypes.Timestamp(res.Session.Issued)
	if err != nil {
		return nil, err
	}

	expires, err := ptypes.Timestamp(res.Session.Expires)
	if err != nil {
		return nil, err
	}

	return &Session{
		AppID:   res.Session.AppId,
		Issued:  issued,
		Expires: expires,
	}, nil
}

// ListActiveAddons returns all the addons that have been activated for the specified business
func (c *GRPCClient) ListActiveAddons(ctx context.Context, businessID string) ([]*AddonActivation, error) {
	var res *marketplace_v1.ListActiveAddonsResponse
	err := vax.Invoke(ctx, func(ctx context.Context, settings vax.CallSettings) error {
		var err error
		ctx, _ = context.WithTimeout(ctx, 3*time.Second)
		res, err = c.client.ListActiveAddons(ctx, &marketplace_v1.ListActiveAddonsRequest{BusinessId: businessID}, grpc.FailFast(false))
		return err
	}, defaultRetryCallOptions)
	if err != nil {
		return nil, err
	}

	var activeAddons = make([]*AddonActivation, len(res.Addons))
	for i, v := range res.Addons {
		addon := AddonActivation{AddonID: v.AddonId, ActivationID: v.ActivationId}
		activeAddons[i] = &addon
	}

	return activeAddons, nil
}

// ListAssociatedUsers returns a list of all users associated with the application
func (c *GRPCClient) ListAssociatedUsers(ctx context.Context, businessID string) ([]*User, error) {
	var res *marketplace_v1.ListAssociatedUsersResponse
	err := vax.Invoke(ctx, func(ctx context.Context, settings vax.CallSettings) error {
		var err error
		ctx, _ = context.WithTimeout(ctx, 3*time.Second)
		res, err = c.client.ListAssociatedUsers(ctx, &marketplace_v1.ListAssociatedUsersRequest{BusinessId: businessID}, grpc.FailFast(false))
		return err
	}, defaultRetryCallOptions)
	if err != nil {
		return nil, err
	}

	var associatedUsers = make([]*User, len(res.Users))
	for i, v := range res.Users {
		user := User{UserID: v.UserId}
		associatedUsers[i] = &user
	}

	return associatedUsers, nil
}

// GetSessionTransferURL is for requesting the session transfer URL for a business.
func (c *GRPCClient) GetSessionTransferURL(ctx context.Context, businessID string) (sessionTransferURL string, err error) {
	var res *marketplace_v1.GetSessionTransferURLResponse
	err = vax.Invoke(ctx, func(ctx context.Context, settings vax.CallSettings) error {
		var err error
		ctx, _ = context.WithTimeout(ctx, 3*time.Second)
		res, err = c.client.GetSessionTransferURL(ctx, &marketplace_v1.GetSessionTransferURLRequest{AccountId: businessID}, grpc.FailFast(false))
		return err
	}, defaultRetryCallOptions)
	if err != nil {
		return "", err
	}

	return res.GetSessionTransferUrl(), nil
}
