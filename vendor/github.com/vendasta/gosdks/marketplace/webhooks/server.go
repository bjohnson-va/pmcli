package marketplacewebhooks

import (
	"context"
	"crypto/rsa"
	"net/http"
	"net/url"

	"github.com/dgrijalva/jwt-go"
	"github.com/vendasta/gosdks/config"
	"github.com/vendasta/gosdks/marketplace/keys"
	"github.com/vendasta/gosdks/util"
)

type purchaseWebhookHandler func(payload *PurchaseWebhookPayload, url *url.URL) error
type accountWebhookHandler func(payload *AccountWebhookPayload, url *url.URL) error

// sessionWebhookHandler should return a code string that can later be redeemed for the given session
type sessionWebhookHandler func(payload *SessionPayload) (code string, err error)

type server struct {
	marketplacePublicKey *rsa.PublicKey
	ctx                  context.Context
}

func getPublicKeyForEnv(env config.Env) []byte {
	switch env {
	case config.Test:
		return []byte(marketplacekeys.TestPublicKey)
	case config.Demo:
		return []byte(marketplacekeys.DemoPublicKey)
	case config.Prod:
		return []byte(marketplacekeys.ProdPublicKey)
	}
	return []byte(marketplacekeys.LocalPublicKey)
}

// NewServer returns a new webhook handling server
func NewServer(ctx context.Context, env config.Env) (*server, error) {
	pemKey := getPublicKeyForEnv(env)

	rsaKey, err := jwt.ParseRSAPublicKeyFromPEM(pemKey)
	if err != nil {
		return nil, err
	}

	return &server{
		ctx:                  ctx,
		marketplacePublicKey: rsaKey,
	}, nil
}

// GetPurchaseWebhookHandler returns a mux compatible handler for handling purchase webhooks
func (s *server) GetPurchaseWebhookHandler(handler purchaseWebhookHandler) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := s.handlePurchaseWebhook(r, handler)
		if err != nil {
			if util.IsError(util.Unauthenticated, err) {
				http.Error(w, err.Error(), http.StatusUnauthorized)
			} else {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}

		w.Write([]byte("ok"))
	}
}

// GetAccountWebhookHandler returns a mux compatible handler for handling account webhooks
func (s *server) GetAccountWebhookHandler(handler accountWebhookHandler) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := s.handleAccountWebhook(r, handler)
		if err != nil {
			if util.IsError(util.Unauthenticated, err) {
				http.Error(w, err.Error(), http.StatusUnauthorized)
			} else {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}

		w.Write([]byte("ok"))
	}
}

// GetSessionWebhookHandler returns a mux compatible handler for handling session webhooks
func (s *server) GetSessionWebhookHandler(handler sessionWebhookHandler) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		resp, err := s.handleSessionWebhook(r, handler)
		if err != nil {
			if util.IsError(util.Unauthenticated, err) {
				http.Error(w, err.Error(), http.StatusUnauthorized)
			} else {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}

		w.Write(resp)
	}
}
