package marketplace

import (
	"time"

	"github.com/vendasta/gosdks/basesdk"
	"github.com/vendasta/gosdks/config"
	"golang.org/x/net/context"
)

type OAuthClientInterface interface {
	CreateJWTAssertion(appID string, appPrivateKey []byte, iat time.Time, exp time.Time) (string, error)
	GetOAuthToken(ctx context.Context, grantType string, assertion string) (*OAuthDetails, error)
}

// oAuthClient is a client which handles calls to marketplace's OAuth apis
type oAuthClient struct {
	basesdk.SDKClient
}

// NewOAuthClient constructs a new oAuthClient with no authorization headers
func NewOAuthClient(env config.Env, rootURLOverride string) OAuthClientInterface {
	var rootURL string
	if rootURLOverride != "" {
		rootURL = rootURLOverride
	} else {
		rootURL = rootURLFromEnv(env)
	}

	auth := basesdk.NoAuth{}
	baseClient := basesdk.BaseClient{
		Authorization: auth,
		RootURL:       rootURL,
	}
	return &oAuthClient{baseClient}
}
