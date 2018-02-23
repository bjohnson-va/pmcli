package marketplace

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/vendasta/gosdks/basesdk"
	"github.com/vendasta/gosdks/util"
	"github.com/vendasta/gosdks/validation"
	"golang.org/x/net/context"
	"time"
)

// OAuthDetails represents the OAuthToken response from the marketplace API
type OAuthDetails struct {
	AccessToken string `json:"access_token"`
	Expires     int    `json:"expires"`
	TokenType   string `json:"token_type"`
}

// JWTGrantType is the grant type expected by the marketplace API when
// using JWT for authorization.
const JWTGrantType = "urn:ietf:params:oauth:grant-type:jwt-bearer"

// CreateJWTAssertion creates the JWT assertion string required for getting
// an OAuth token from the marketplace API.
func (c *oAuthClient) CreateJWTAssertion(appID string, appPrivateKey []byte, iat time.Time, exp time.Time,
) (string, error) {
	if appID == "" {
		return "", fmt.Errorf("appID is required")
	}

	key, err := jwt.ParseRSAPrivateKeyFromPEM(appPrivateKey)
	if err != nil {
		return "", fmt.Errorf("unparsable private key: %s", err.Error())
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"iss": appID,
		"iat": iat.Unix(),
		"exp": exp.Unix(),
	})
	jot, err := token.SignedString(key)
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %s", err.Error())
	}
	return jot, nil
}

// GetOAuthToken requests an OAuth token from marketplace using the given
// grantType and assertion.
func (c *oAuthClient) GetOAuthToken(ctx context.Context, grantType string,
	assertion string) (*OAuthDetails, error) {
	err := validation.NewValidator().
		Rule(validation.StringNotEmpty(grantType, util.InvalidArgument, "grantType is required")).
		Rule(validation.StringNotEmpty(assertion, util.InvalidArgument, "assertion is required")).
		Validate()
	if err != nil {
		return nil, err
	}
	path := "/api/v1/oauth/token"
	params := map[string]interface{}{
		"grant_type": grantType,
		"assertion":  assertion,
	}

	r, err := c.Post(ctx, path, params, basesdk.Idempotent())
	if err != nil {
		return nil, fmt.Errorf("failed to post: %s", err)
	}
	defer r.Body.Close()
	type Response struct {
		OAuthDetails *OAuthDetails `json:"data"`
	}
	res := Response{}
	if err := json.NewDecoder(r.Body).Decode(&res); err != nil {
		return nil, fmt.Errorf("failed to decode: %s", err.Error())
	}
	if err != nil {
		return nil, err
	}
	return res.OAuthDetails, nil
}
