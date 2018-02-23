package marketplace

import (
	"github.com/vendasta/gosdks/basesdk"
	"github.com/vendasta/gosdks/config"
	"golang.org/x/net/context"
)

type AccountClientInterface interface {
	GetSessionTransferURL(ctx context.Context, accountID string) (string, error)
}

// accountClient is a client which handles calls to marketplace's account apis
type accountClient struct {
	basesdk.SDKClient
}

// NewAccountClient creates an account client to allow calls to be made to marketplace.
func NewAccountClient(ctx context.Context, appID string, privateKey []byte, env config.Env, rootURLOverride string) AccountClientInterface {
	var rootURL string
	if rootURLOverride != "" {
		rootURL = rootURLOverride
	} else {
		rootURL = rootURLFromEnv(env)
	}
	oauthClient := NewOAuthClient(env, rootURL)
	return accountClient{
		basesdk.BaseClient{
			Authorization: NewMarketplaceAuthorization(ctx, appID, privateKey, oauthClient, nil),
			RootURL:       rootURL,
		},
	}
}
