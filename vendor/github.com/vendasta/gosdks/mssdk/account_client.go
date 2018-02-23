package mssdk

import (
	"context"

	"github.com/vendasta/gosdks/basesdk"
	"github.com/vendasta/gosdks/config"
	"github.com/vendasta/gosdks/util"
	"github.com/vendasta/gosdks/validation"
)

const (
	accountCreatePath     = "/internalApi/v3/site/create/"
	accountDeactivatePath = "/internalApi/v3/site/delete/"
)

// AccountClientInterface exposes methods for interacting with account definitions from Presence Builder
type AccountClientInterface interface {
	Activate(ctx context.Context, accountGroupID string, partnerID string, ssoToken string) (string, error)
	Deactivate(ctx context.Context, accountGroupID string, partnerID string) error
}

// BuildAccountClient creates an account client.
func BuildAccountClient(apiUser string, apiKey string, env config.Env) AccountClientInterface {
	auth := basesdk.UserKey{APIUser: apiUser, APIKey: apiKey}
	baseClient := basesdk.BaseClient{Authorization: auth, RootURL: rootURLFromEnv(env)}
	return &AccountClient{baseClient}
}

// AccountClient is a client which handles calls to Presence Builder's account apis.
// Implements the AccountClientInterface
type AccountClient struct {
	basesdk.SDKClient
}

//Activate activates a MS account for the account group
func (c *AccountClient) Activate(ctx context.Context, accountGroupID string, partnerID string, ssoToken string) (string, error) {
	err := validation.NewValidator().
		Rule(validation.StringNotEmpty(accountGroupID, util.InvalidArgument, "accountGroupID is required")).
		Rule(validation.StringNotEmpty(partnerID, util.InvalidArgument, "partnerID is required")).
		Validate()
	if err != nil {
		return "", err
	}
	//The dummy params are straight from the AA implementation of calling MS account create
	params := map[string]interface{}{
		"ssoToken": ssoToken,
		"agid":     accountGroupID,
		"pid":      partnerID,
	}
	response, err := c.Post(ctx, accountCreatePath, params)
	if err != nil {
		return "", err
	}
	return accountIDFromResponse(response)
}

//Deactivate deactivates a MS account for the account group
func (c *AccountClient) Deactivate(ctx context.Context, accountID string, partnerID string) error {
	err := validation.NewValidator().
		Rule(validation.StringNotEmpty(accountID, util.InvalidArgument, "accountID is required")).
		Rule(validation.StringNotEmpty(partnerID, util.InvalidArgument, "partnerID is required")).
		Validate()
	if err != nil {
		return err
	}
	params := map[string]interface{}{
		"msid": accountID,
		"pid":  partnerID,
	}
	_, err = c.Post(ctx, accountDeactivatePath, params)
	if err != nil {
		return err
	}
	return nil
}
