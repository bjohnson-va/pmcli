package marketplacewebhooks

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/vendasta/gosdks/logging"
	"github.com/vendasta/gosdks/util"
)

// SessionPayload represents the payload of an account webhook
type SessionPayload struct {
	SessionID string `json:"vendasta.com/marketplace/session_id"`
	Language  string `json:"vendasta.com/marketplace/lang"`
	UserID    string `json:"vendasta.com/marketplace/user_id"`
	PartnerID string `json:"vendasta.com/marketplace/partner_id"`
	NavBarURL string `json:"vendasta.com/marketplace/product_navbar_data_url"`
}

type sessionClaims struct {
	jwt.StandardClaims
	SessionPayload
}

// handles a session webhook by wrapping the passed handler
func (s *server) handleSessionWebhook(r *http.Request, handler sessionWebhookHandler) ([]byte, error) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logging.Warningf(s.ctx, "Error reading request body: %s", err.Error())
		return nil, util.Error(util.Internal, "Error reading request body: %s", err.Error())
	}

	payload, err := s.parseSessionPayload(string(reqBody))
	logging.Infof(s.ctx, "session webhook payload is %+v", payload)
	if err != nil {
		logging.Errorf(s.ctx, "Jwt auth failed: %s. Payload body: %s", err.Error(), string(reqBody))
		return nil, util.Error(util.Unauthenticated, "Jwt auth failed: %s", err.Error())
	}

	code, err := handler(payload)
	if err != nil {
		logging.Warningf(s.ctx, "Error handling session webhook payload: %s", err.Error())
		return nil, util.Error(util.Internal, err.Error())
	}

	resp := map[string]string{
		"code": code,
	}
	return json.Marshal(resp)
}

func (s *server) parseSessionPayload(jwtPayload string) (*SessionPayload, error) {
	jwtClaims := &sessionClaims{}
	_, err := jwt.ParseWithClaims(jwtPayload, jwtClaims,
		func(tok *jwt.Token) (interface{}, error) {
			return s.marketplacePublicKey, nil
		},
	)
	if err != nil {
		return nil, err
	}
	payload := &jwtClaims.SessionPayload
	return payload, nil
}
