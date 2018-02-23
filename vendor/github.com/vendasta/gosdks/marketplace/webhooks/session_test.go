package marketplacewebhooks

import (
	"encoding/json"
	"net/http"
	"strings"
	"testing"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	"github.com/vendasta/gosdks/config"
	"github.com/vendasta/gosdks/util"
	"golang.org/x/net/context"
)

func buildTestSessionPayload() []byte {
	payload, err := json.Marshal(&sessionClaims{
		jwt.StandardClaims{
			Issuer: "Vendasta Marketplace",
		},
		SessionPayload{},
	})
	if err != nil {
		panic(err)
	}
	return payload
}

func Test_handleSessionWebhook(t *testing.T) {
	type testCase struct {
		name       string
		bodyString string
		handler    sessionWebhookHandler

		expectedCodeJson []byte
		expectedError    error
	}

	cases := []*testCase{
		{
			name:       "should give unauthenticated error for payload signed with wrong key",
			bodyString: signJWT([]byte(invalidPrivateKey), buildTestSessionPayload(), nil),
			handler: func(payload *SessionPayload) (string, error) {
				return "", nil
			},
			expectedError: util.Error(util.Unauthenticated, "Jwt auth failed: crypto/rsa: verification error"),
		},
		{
			name:       "should give internal error if provided handler returns error",
			bodyString: signJWT([]byte(localPrivateKey), buildTestSessionPayload(), nil),
			handler: func(payload *SessionPayload) (string, error) {
				return "", util.Error(util.InvalidArgument, "some error")
			},
			expectedError: util.Error(util.Internal, "some error"),
		},
		{
			name:       "should return json containing code returned by provided handler",
			bodyString: signJWT([]byte(localPrivateKey), buildTestSessionPayload(), nil),
			handler: func(payload *SessionPayload) (string, error) {
				return "super-sekret", nil
			},
			expectedCodeJson: []byte("{\"code\":\"super-sekret\"}"),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ctx := context.Background()
			server, _ := NewServer(ctx, config.Local)
			r := &http.Request{
				Body: &bodyStub{Reader: strings.NewReader(c.bodyString)},
			}

			codeJson, actualErr := server.handleSessionWebhook(r, c.handler)

			assert.Equal(t, c.expectedCodeJson, codeJson)
			assert.Equal(t, c.expectedError, actualErr)
		})
	}
}
