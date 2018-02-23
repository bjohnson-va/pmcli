package smsdk

import (
	"context"

	"github.com/vendasta/gosdks/basesdk"
	"github.com/vendasta/gosdks/config"
	"github.com/vendasta/gosdks/util"
	"github.com/vendasta/gosdks/validation"
)

const (
	accountCreatePath     = "/internalApi/v2/account/create/"
	accountActivatePath   = "/internalApi/v2/account/activate/"
	accountDeactivatePath = "/internalApi/v2/account/deactivate/"
)

// AccountClientInterface exposes methods for interacting with account definitions from Social Marketing
type AccountClientInterface interface {
	Create(ctx context.Context, accountGroupID string, partnerID string) (*Account, error)
	Activate(ctx context.Context, accountGroupID string, partnerID string, ssoToken string) (string, error)
	Deactivate(ctx context.Context, accountGroupID string, partnerID string) error
}

// BuildAccountClient creates an account client.
func BuildAccountClient(apiUser string, apiKey string, env config.Env) AccountClientInterface {
	auth := basesdk.UserKey{APIUser: apiUser, APIKey: apiKey}
	baseClient := basesdk.BaseClient{Authorization: auth, RootURL: rootURLFromEnv(env)}
	return &AccountClient{baseClient}
}

// AccountClient is a client which handles calls to Social Marketing's account apis.
// Implements the AccountClientInterface
type AccountClient struct {
	basesdk.SDKClient
}

//Create activates a SM account for the account group - Should be deprecated
func (c *AccountClient) Create(ctx context.Context, accountGroupID string, partnerID string) (*Account, error) {
	err := validation.NewValidator().
		Rule(validation.StringNotEmpty(accountGroupID, util.InvalidArgument, "accountGroupID is required")).
		Rule(validation.StringNotEmpty(partnerID, util.InvalidArgument, "partnerID is required")).
		Validate()
	if err != nil {
		return nil, err
	}
	//The dummy params are straight from the AA implementation of calling SM account create
	params := map[string]interface{}{
		"address":        "dummy",
		"city":           "dummy",
		"state":          "sk",
		"country":        "ca",
		"zip":            "dummy",
		"companyName":    "dummy",
		"accountGroupId": accountGroupID,
		"pid":            partnerID,
	}
	response, err := c.Post(ctx, accountCreatePath, params)
	if err != nil {
		return nil, err
	}
	return accountFromResponse(response)
}

//Activate activates a SM account for the account group
func (c *AccountClient) Activate(ctx context.Context, accountGroupID string, partnerID string, ssoToken string) (string, error) {
	err := validation.NewValidator().
		Rule(validation.StringNotEmpty(accountGroupID, util.InvalidArgument, "accountGroupID is required")).
		Rule(validation.StringNotEmpty(partnerID, util.InvalidArgument, "partnerID is required")).
		Validate()
	if err != nil {
		return "", err
	}
	//The dummy params are straight from the AA implementation of calling SM account create
	params := map[string]interface{}{
		"ssoToken":       ssoToken,
		"accountGroupId": accountGroupID,
		"partnerId":      partnerID,
	}
	response, err := c.Post(ctx, accountActivatePath, params)
	if err != nil {
		return "", err
	}
	return accountIDFromResponse(response)
}

//Deactivate deactivates a SM account for the account group
func (c *AccountClient) Deactivate(ctx context.Context, accountGroupID string, partnerID string) error {
	err := validation.NewValidator().
		Rule(validation.StringNotEmpty(accountGroupID, util.InvalidArgument, "accountGroupID is required")).
		Rule(validation.StringNotEmpty(partnerID, util.InvalidArgument, "partnerID is required")).
		Validate()
	if err != nil {
		return err
	}
	params := map[string]interface{}{
		"accountGroupId": accountGroupID,
		"partnerId":      partnerID,
	}
	_, err = c.Post(ctx, accountDeactivatePath, params)
	if err != nil {
		return err
	}
	return nil
}
