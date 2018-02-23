package marketplaceapi

import (
	"fmt"
	"time"

	"github.com/vendasta/gosdks/config"
	"golang.org/x/net/context"
)

// Token holds the response data from a GetOAuthToken call.
type Token struct {
	AccessToken string
	TokenType   string
	Expires     time.Time
}

// AddonActivation holds the information about active add-ons
type AddonActivation struct {
	AddonID      string
	ActivationID string
}

// User holds information about the user
type User struct {
	UserID string
}

// Session holds the response data from a GetSessionFromToken call.
type Session struct {
	AppID   string
	Issued  time.Time
	Expires time.Time
}

// Client is an interface for the marketplace api microservice.
type Client interface {
	// GetOAuthToken get's an OAuth access token for an application.
	GetOAuthToken(ctx context.Context, grantType string, assertion string) (*Token, error)
	// ListActiveAddons returns all the addons that have been activated for the specified business
	ListActiveAddons(ctx context.Context, businessID string) ([]*AddonActivation, error)
	// ListAssociatedUsers returns all the users connected to the specified business
	ListAssociatedUsers(ctx context.Context, businessID string) ([]*User, error)
	// GetSessionTransferURL is for requesting the session transfer URL for a business.
	GetSessionTransferURL(ctx context.Context, businessID string) (sessionTransferURL string, err error)
	// GetSessionFromToken is for requesting the session information from a token
	GetSessionFromToken(ctx context.Context, accessToken string) (*Session, error)
}

// NewClient returns the default concrete implementation of the client, pre-configured given an environment.
func NewClient(ctx context.Context, e config.Env) (Client, error) {
	address := addresses[e]
	if address == "" {
		return nil, fmt.Errorf("unable to create client with environment %d", e)
	}
	useTLS := e != config.Local
	return NewGRPCClient(ctx, address, useTLS, scopes[e])
}

var addresses = map[config.Env]string{
	config.Local: "marketplace-api-test.vendasta-internal.com:443",
	config.Test:  "marketplace-api-test.vendasta-internal.com:443",
	config.Demo:  "marketplace-api-demo.vendasta-internal.com:443",
	config.Prod:  "marketplace-api-prod.vendasta-internal.com:443",
}

var scopes = map[config.Env]string{
	config.Local: "https://marketplace-api-test.vendasta-internal.com",
	config.Test:  "https://marketplace-api-test.vendasta-internal.com",
	config.Demo:  "https://marketplace-api-demo.vendasta-internal.com",
	config.Prod:  "https://marketplace-api-prod.vendasta-internal.com",
}
