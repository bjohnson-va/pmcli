package cssdk

import (
	"errors"

	"github.com/vendasta/gosdks/basesdk"
	"github.com/vendasta/gosdks/config"
	"golang.org/x/net/context"
)

const (
	socialServiceListPath = "/internalApi/v2/socialService/list/"
)

// SocialServicesClientInterface defines the interface of a social services client
type SocialServicesClientInterface interface {
	ListSocialServices(ctx context.Context, accountGroupID string) ([]*SocialService, error)
}

// SocialServicesClient is a client which handles calls to core services's social services apis, implements the interface
type SocialServicesClient struct {
	basesdk.SDKClient
}

// BuildSocialServicesClient creates a social profile client.
func BuildSocialServicesClient(apiUser string, apiKey string, env config.Env) SocialServicesClient {
	auth := basesdk.UserKey{APIUser: apiUser, APIKey: apiKey}
	baseClient := basesdk.BaseClient{Authorization: auth, RootURL: rootURLFromEnv(env)}
	return SocialServicesClient{baseClient}
}

//ListSocialServices Returns a list of the social services connected to the account group
func (c SocialServicesClient) ListSocialServices(ctx context.Context, accountGroupID string) ([]*SocialService, error) {
	if accountGroupID == "" {
		return nil, errors.New("account group id is required")
	}
	params := map[string]interface{}{"agid": accountGroupID}
	response, err := c.Get(ctx, socialServiceListPath, params, basesdk.Idempotent())
	if err != nil {
		return nil, err
	}
	socialServices, err := socialServicesFromResponse(response)
	if err != nil {
		return nil, err
	}
	return socialServices, nil
}
