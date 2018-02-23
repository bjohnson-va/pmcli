package marketplace

import (
	"fmt"
	"net/http"
	"time"

	"github.com/vendasta/gosdks/basesdk"
	"github.com/vendasta/gosdks/logging"
	"golang.org/x/net/context"
)

type marketplaceToken struct {
	expiry time.Time
	value  string
}

// marketplaceAuthorization handles signing requests for marketplace SDK calls
type marketplaceAuthorization struct {
	ctx         context.Context
	privateKey  []byte
	appID       string
	oauthClient OAuthClientInterface
	token       *marketplaceToken
}

// NewMarketplaceAuthorization returns a new marketplaceAuthorization that meets the RequestAuthentication interface
func NewMarketplaceAuthorization(ctx context.Context, appID string, privateKey []byte, oauthClient OAuthClientInterface, token *marketplaceToken) basesdk.RequestAuthorization {
	return &marketplaceAuthorization{
		ctx:         ctx,
		privateKey:  privateKey,
		appID:       appID,
		oauthClient: oauthClient,
		token:       token,
	}
}

// SignRequest will add an Authorization header to the given request by requesting an oauth token from marketplace
func (m *marketplaceAuthorization) SignRequest(r *http.Request) {
	if m.token == nil || m.token.expiry.Before(time.Now().UTC()) {
		err := m.refreshToken()
		if err != nil {
			logging.Errorf(m.ctx, "Error refreshing token: %s", err.Error())
			return
		}
	}

	r.Header.Add("Authorization", fmt.Sprintf("Bearer %s", m.token.value))
}

func (m *marketplaceAuthorization) refreshToken() error {
	now := time.Now().UTC()
	expiry := now.Add(time.Hour)

	assertion, err := m.oauthClient.CreateJWTAssertion(m.appID, m.privateKey, now, expiry)
	if err != nil {
		logging.Warningf(m.ctx, "Error creating JWT assertion: %s", err.Error())
		return err
	}

	details, err := m.oauthClient.GetOAuthToken(m.ctx, JWTGrantType, assertion)
	if err != nil {
		logging.Warningf(m.ctx, "Error getting OAuth token: %s", err.Error())
		return err
	}

	m.token = &marketplaceToken{
		expiry: expiry,
		value:  details.AccessToken,
	}

	return nil
}
