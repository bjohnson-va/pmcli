package marketplacewebhooks

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	"github.com/vendasta/gosdks/config"
	"github.com/vendasta/gosdks/util"
)

func buildTestAccountWebhookPayload() []byte {
	payload, err := json.Marshal(&accountWebhookClaims{
		jwt.StandardClaims{
			Issuer: "Vendasta Marketplace",
		},
		AccountWebhookPayload{},
	})
	if err != nil {
		panic(err)
	}
	return payload
}

func Test_handleAccountWebhook(t *testing.T) {
	type testCase struct {
		name       string
		bodyString string
		handler    accountWebhookHandler

		expectedError error
	}

	cases := []*testCase{
		{
			name:       "should give unauthenticated error for payload signed with wrong key",
			bodyString: signJWT([]byte(invalidPrivateKey), buildTestAccountWebhookPayload(), nil),
			handler: func(payload *AccountWebhookPayload, url *url.URL) error {
				return nil
			},
			expectedError: util.Error(util.Unauthenticated, "Jwt auth failed: crypto/rsa: verification error"),
		},
		{
			name:       "should give internal error if provided handler returns error",
			bodyString: signJWT([]byte(localPrivateKey), buildTestAccountWebhookPayload(), nil),
			handler: func(payload *AccountWebhookPayload, url *url.URL) error {
				return util.Error(util.InvalidArgument, "some error")
			},
			expectedError: util.Error(util.Internal, "some error"),
		},
		{
			name:       "should give no error for correctly formatted JWT with right keys",
			bodyString: signJWT([]byte(localPrivateKey), buildTestAccountWebhookPayload(), nil),
			handler: func(payload *AccountWebhookPayload, url *url.URL) error {
				return nil
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ctx := context.Background()
			server, _ := NewServer(ctx, config.Local)
			r := &http.Request{
				Body: &bodyStub{Reader: strings.NewReader(c.bodyString)},
			}

			actualErr := server.handleAccountWebhook(r, c.handler)

			assert.Equal(t, c.expectedError, actualErr)
		})
	}
}
