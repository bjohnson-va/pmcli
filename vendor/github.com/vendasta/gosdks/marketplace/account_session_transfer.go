package marketplace

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/vendasta/gosdks/basesdk"
	"golang.org/x/net/context"
	"net/http"
)

// SessionTransferURLFromHTTPResponse implements SessionTransferResponseHandler
type SessionTransferURLFromHTTPResponse struct {
	SessionTransferResponseHandler
}

// SessionURLResponse represents the response from the session transfer URL endpoint
type SessionURLResponse struct {
	SessionTransferURL string `json:"session_transfer_url"`
}

// ProcessResponse extracts the session transfer URL from the given Response
func (h SessionTransferURLFromHTTPResponse) ProcessResponse(r *http.Response) (string, error) {
	defer r.Body.Close()
	type Response struct {
		SessionURLResponse *SessionURLResponse `json:"data"`
	}
	res := Response{}
	if err := json.NewDecoder(r.Body).Decode(&res); err != nil {
		reason := "failed to convert response to URL: " + err.Error()
		return "", errors.New(reason)
	}
	return res.SessionURLResponse.SessionTransferURL, nil
}

// SessionTransferResponseHandler processes the http.Response objects
// that are generated from the session transfer URL API.
type SessionTransferResponseHandler interface {
	ProcessResponse(response *http.Response) (string, error)
}

// GetSessionTransferURL fetches and returns the session transfer URL string
func (c accountClient) GetSessionTransferURL(ctx context.Context, accountID string) (string, error) {
	response, err := c.getSessionTransferURL(ctx, accountID)
	if err != nil {
		return "", err
	}
	return SessionTransferURLFromHTTPResponse{}.ProcessResponse(response)
}

func (c accountClient) getSessionTransferURL(ctx context.Context,
	accountID string) (*http.Response, error) {
	if accountID == "" {
		return nil, errors.New("accountID is required")
	}
	path := c.buildSessionTransferAPIURL(accountID)
	params := map[string]interface{}{}

	response, err := c.Get(ctx, path, params, basesdk.Idempotent())
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c accountClient) buildSessionTransferAPIURL(accountID string) string {
	return fmt.Sprintf("/api/v1/account/%s/session-transfer", accountID)
}
