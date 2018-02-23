package salestool

import (
	"errors"

	"encoding/json"

	"github.com/vendasta/gosdks/basesdk"
	"github.com/vendasta/gosdks/util"
	"golang.org/x/net/context"
)

const (
	getSalespersonPath           = "/internalApi/v2/salesperson/get/"
	getRoundRobinSalespersonPath = "/internalApi/v3/salesperson/roundRobin/"
)

// SalespersonClient is a client which handles calls to salestool's snapshot apis
type SalespersonClient struct {
	basesdk.SDKClient
}

// GetRequest defines structure of request for getting a salesperson
type GetSalespersonRequest struct {
	PartnerID        string
	SalespersonID    string
	SalespersonEmail string
	AccountGroupID   string
}

type RoundRobinRequest struct {
	PartnerID         string
	MarketID          string
	LastSalespersonID string
}

// BuildSalespersonClient creates a salesperson client
func BuildSalespersonClient(apiUser string, apiKey string, rootURL string) SalespersonClient {
	auth := basesdk.UserKey{APIUser: apiUser, APIKey: apiKey}
	baseClient := basesdk.BaseClient{Authorization: auth, RootURL: rootURL}
	return SalespersonClient{baseClient}
}

//Get gets a salesperson from ST
func (c SalespersonClient) GetSalesperson(ctx context.Context, r *GetSalespersonRequest) (*Salesperson, error) {

	if r.PartnerID == "" {
		return nil, errors.New("partnerId is required")
	}
	searchIds := []string{r.SalespersonID, r.SalespersonEmail, r.AccountGroupID}
	if err := util.RequireExclusiveProperties(searchIds); err != nil {
		return nil, err
	}

	params := map[string]interface{}{
		"salesPersonId":    r.SalespersonID,
		"pid":              r.PartnerID,
		"accountGroupId":   r.AccountGroupID,
		"salesPersonEmail": r.SalespersonEmail,
	}

	response, err := c.Get(ctx, getSalespersonPath, params, basesdk.Idempotent())
	if err != nil {
		return nil, err
	}

	sp, err := salespersonFromResponse(response)
	if err != nil {
		return nil, err
	}

	return sp, nil
}

//RoundRobin gets the next salesperson in a round robin ordering for a partner/market
func (c SalespersonClient) RoundRobin(ctx context.Context, r *RoundRobinRequest) (string, error) {
	if r.PartnerID == "" {
		return "", errors.New("partnerId is required")
	}

	params := map[string]interface{}{
		"partnerId":         r.PartnerID,
		"marketId":          r.MarketID,
		"lastSalespersonId": r.LastSalespersonID,
	}

	response, err := c.Get(ctx, getRoundRobinSalespersonPath, params)
	if err != nil {
		return "", err
	}

	type Data struct {
		Data map[string]string `json:"data"`
	}

	data := Data{}
	if err = json.NewDecoder(response.Body).Decode(&data); err != nil {
		return "", errors.New("Failed to decode JSON response: " + err.Error())
	}

	return data.Data["salespersonId"], nil
}
