package marketplace

import (
	"github.com/stretchr/testify/assert"
	"github.com/vendasta/gosdks/util"
	"golang.org/x/net/context"
	"net/http"
	"testing"
	"time"
)

type oauthClientStub struct {
	methods                    []string
	params                     map[string]interface{}
	createJWTResponse          string
	getOAuthTokenResponse      *OAuthDetails
	createJWTErrorResponse     error
	getOAuthTokenErrorResponse error
}

func (o *oauthClientStub) CreateJWTAssertion(appID string, appPrivateKey []byte, iat time.Time, exp time.Time) (string, error) {
	o.methods = append(o.methods, "CreateJWTAssertion")
	o.params = make(map[string]interface{})
	o.params["appID"] = appID
	o.params["appPrivateKey"] = appPrivateKey
	o.params["iat"] = iat
	o.params["exp"] = exp
	return o.createJWTResponse, o.createJWTErrorResponse
}

func (o *oauthClientStub) GetOAuthToken(ctx context.Context, grantType string, assertion string) (*OAuthDetails, error) {
	o.methods = append(o.methods, "CreateJWTAssertion")
	o.params["grantType"] = grantType
	o.params["assertion"] = assertion
	return o.getOAuthTokenResponse, o.getOAuthTokenErrorResponse
}

func Test_SignRequest(t *testing.T) {
	type testCase struct {
		name             string
		initialToken     *marketplaceToken
		createError      error
		getTokenError    error
		getTokenResponse *OAuthDetails

		expectedAuthHeader string
	}

	testCases := []*testCase{
		{
			name:               "sets no auth header if CreateJWTAssertion returns error",
			createError:        util.Error(util.Internal, "error"),
			expectedAuthHeader: "",
		},
		{
			name:               "sets no auth header if GetOAuthToken returns error",
			getTokenError:      util.Error(util.Internal, "error"),
			expectedAuthHeader: "",
		},
		{
			name:               "sets auth header to result of GetOAuthToken",
			getTokenResponse:   &OAuthDetails{AccessToken: "footoken"},
			expectedAuthHeader: "Bearer footoken",
		},
		{
			name:               "won't refresh token if current token is not expired",
			initialToken:       &marketplaceToken{value: "bartoken", expiry: time.Now().UTC().Add(time.Hour)},
			getTokenResponse:   &OAuthDetails{AccessToken: "footoken"},
			expectedAuthHeader: "Bearer bartoken",
		},
		{
			name:               "will refresh token if current token is expired",
			initialToken:       &marketplaceToken{value: "bartoken", expiry: time.Now().UTC().Add(-time.Hour)},
			getTokenResponse:   &OAuthDetails{AccessToken: "footoken"},
			expectedAuthHeader: "Bearer footoken",
		},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			ctx := context.Background()
			ocs := &oauthClientStub{
				createJWTErrorResponse:     c.createError,
				getOAuthTokenResponse:      c.getTokenResponse,
				getOAuthTokenErrorResponse: c.getTokenError,
			}
			r := &http.Request{
				Header: map[string][]string{},
			}
			ma := NewMarketplaceAuthorization(ctx, "app-id", []byte("private-key"), ocs, c.initialToken)

			ma.SignRequest(r)

			assert.Equal(t, r.Header.Get("Authorization"), c.expectedAuthHeader)
		})
	}
}
