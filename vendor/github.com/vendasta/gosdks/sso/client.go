package sso

import (
	"fmt"

	"github.com/vendasta/gosdks/config"
	"golang.org/x/net/context"
)

// Client is an interface for the sso microservice.
type Client interface {
	GetSessionTransferURL(ctx context.Context, serviceProviderID string, serviceContext ServiceContext) (sessionTransferURL string, err error)
}

// ServiceContext is a context provided by a service provider, used to determine the identity provider.
type ServiceContext interface {
}

// AccountContext the context for a session transfer for an account.
type AccountContext struct {
	AccountID string
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
	config.Local: "sso-api-test.vendasta-internal.com:443",
	config.Test:  "sso-api-test.vendasta-internal.com:443",
	config.Demo:  "sso-api-demo.vendasta-internal.com:443",
	config.Prod:  "sso-api-prod.vendasta-internal.com:443",
}

var scopes = map[config.Env]string{
	config.Local: "https://sso-api-test.vendasta-internal.com",
	config.Test:  "https://sso-api-test.vendasta-internal.com",
	config.Demo:  "https://sso-api-demo.vendasta-internal.com",
	config.Prod:  "https://sso-api-prod.vendasta-internal.com",
}
