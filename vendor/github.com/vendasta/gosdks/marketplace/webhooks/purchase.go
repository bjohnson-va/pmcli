package marketplacewebhooks

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/vendasta/gosdks/logging"
	"github.com/vendasta/gosdks/util"
)

// PurchaseWebhookPayload represents the payload of a purchase webhook
type PurchaseWebhookPayload struct {
	Account       MarketplaceAccount `json:"account"`
	OrderFormJson json.RawMessage    `json:"order_form"`
	MarketID      string             `json:"market_id"`
	WebhookID     string             `json:"webhook_id"`
	Action        string             `json:"action"`
	PartnerID     string             `json:"partner_id"`
}

type purchaseWebhookClaims struct {
	jwt.StandardClaims
	WebhookPayload PurchaseWebhookPayload `json:"vendasta.com/marketplace/webhook"`
}

// handles a provision webhook by wrapping the passed handler
func (s *server) handlePurchaseWebhook(r *http.Request, handler purchaseWebhookHandler) error {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logging.Warningf(s.ctx, "Error reading request body: %s", err.Error())
		return util.Error(util.Internal, "Error reading request body: %s", err.Error())
	}

	payload, err := s.parsePurchasePayload(string(reqBody))
	logging.Infof(s.ctx, "Purchase webhook payload is %+v", payload)
	if err != nil {
		logging.Warningf(s.ctx, "Jwt auth failed: %s. Payload body: %s", err.Error(), string(reqBody))
		return util.Error(util.Unauthenticated, "Jwt auth failed: %s", err.Error())
	}

	err = handler(payload, r.URL)
	if err != nil {
		logging.Warningf(s.ctx, "Error handling purchase webhook payload: %s", err.Error())
		return util.Error(util.Internal, err.Error())
	}

	return nil
}

func (s *server) parsePurchasePayload(jwtPayload string) (*PurchaseWebhookPayload, error) {
	jwtClaims := &purchaseWebhookClaims{}
	_, err := jwt.ParseWithClaims(jwtPayload, jwtClaims,
		func(tok *jwt.Token) (interface{}, error) {
			return s.marketplacePublicKey, nil
		},
	)
	if err != nil {
		return nil, err
	}
	payload := &jwtClaims.WebhookPayload
	return payload, nil
}
