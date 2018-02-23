package cssdk

import (
	"errors"
	"net/http"

	"github.com/vendasta/gosdks/basesdk"
	"github.com/vendasta/gosdks/config"
	"golang.org/x/net/context"
)

const (
	registerPath = "/api/v1/socialProfile/register/"
	createPath   = "/api/v1/socialProfile/create/"
	lookupPath   = "/api/v1/socialProfile/lookupRegistrations/"
)

// SocialProfileClientInterface defines the interface of a social profile client
type SocialProfileClientInterface interface {
	CreateSocialProfile(ctx context.Context, accountGroupID string, socialProfileGroupID string) (*SocialProfile, error)
	RegisterSocialProfile(ctx context.Context, accountGroupID string, socialProfileID string) (*SocialProfileRegistration, error)
	LookupSocialProfileRegistrations(ctx context.Context, socialProfileID string) ([]*SocialProfileRegistration, error)
}

// SocialProfileClient is a client which handles calls to core services's social profile apis, implements the interface
type SocialProfileClient struct {
	basesdk.SDKClient
}

// BuildSocialProfileClient creates a social profile client.
func BuildSocialProfileClient(apiUser string, apiKey string, env config.Env) SocialProfileClient {
	auth := basesdk.UserKey{APIUser: apiUser, APIKey: apiKey}
	baseClient := basesdk.BaseClient{Authorization: auth, RootURL: rootURLFromEnv(env)}
	return SocialProfileClient{baseClient}
}

// CreateSocialProfileResponseHandler is a function which handles the http response from the create social profile api
type CreateSocialProfileResponseHandler func(r *http.Response) (*SocialProfile, error)

// CreateSocialProfile returns the social profile data
func (c SocialProfileClient) CreateSocialProfile(ctx context.Context, accountGroupID string, socialProfileGroupID string) (*SocialProfile, error) {
	return c.createSocialProfile(ctx, socialProfileFromResponse, accountGroupID, socialProfileGroupID)
}

func (c SocialProfileClient) createSocialProfile(ctx context.Context, responseHandler CreateSocialProfileResponseHandler, accountGroupID string, socialProfileGroupID string) (*SocialProfile, error) {
	if accountGroupID == "" {
		return nil, errors.New("account group id is required")
	}
	params := map[string]interface{}{"accountId": accountGroupID}
	if socialProfileGroupID != "" {
		params["spgid"] = socialProfileGroupID
	}
	response, err := c.Post(ctx, createPath, params)
	if err != nil {
		return nil, err
	}
	socialProfile, err := responseHandler(response)
	if err != nil {
		return nil, err
	}
	return socialProfile, nil
}

// RegisterSocialProfileResponseHandler is a function which handles the http response from the register social profile api
type RegisterSocialProfileResponseHandler func(r *http.Response) (*SocialProfileRegistration, error)

// RegisterSocialProfile returns the social profile registration data
func (c SocialProfileClient) RegisterSocialProfile(ctx context.Context, accountGroupID string, socialProfileID string) (*SocialProfileRegistration, error) {
	return c.registerSocialProfile(ctx, socialProfileRegistrationFromResponse, accountGroupID, socialProfileID)
}

func (c SocialProfileClient) registerSocialProfile(ctx context.Context, responseHandler RegisterSocialProfileResponseHandler, accountGroupID string, socialProfileID string) (*SocialProfileRegistration, error) {
	if accountGroupID == "" {
		return nil, errors.New("account group id is required")
	}
	if socialProfileID == "" {
		return nil, errors.New("social profile id is required")
	}
	params := map[string]interface{}{"accountId": accountGroupID, "spid": socialProfileID}
	response, err := c.Get(ctx, registerPath, params)
	if err != nil {
		return nil, err
	}
	socialProfileRegistration, err := responseHandler(response)
	if err != nil {
		return nil, err
	}
	return socialProfileRegistration, nil
}

// LookupSocialProfileRegistrationsResponseHandler is a function which handles the http response from the lookup social profile registrations api
type LookupSocialProfileRegistrationsResponseHandler func(r *http.Response) ([]*SocialProfileRegistration, error)

// LookupSocialProfileRegistrations returns the list of social profile registration data
func (c SocialProfileClient) LookupSocialProfileRegistrations(ctx context.Context, socialProfileID string) ([]*SocialProfileRegistration, error) {
	return c.lookupSocialProfilesRegistrations(ctx, socialProfileRegistrationsFromResponse, socialProfileID)
}

func (c SocialProfileClient) lookupSocialProfilesRegistrations(ctx context.Context, responseHandler LookupSocialProfileRegistrationsResponseHandler, socialProfileID string) ([]*SocialProfileRegistration, error) {
	if socialProfileID == "" {
		return nil, errors.New("social profile id is required")
	}
	params := map[string]interface{}{"spid": socialProfileID}
	response, err := c.Post(ctx, lookupPath, params)
	if err != nil {
		return nil, err
	}
	socialProfileRegistrations, err := responseHandler(response)
	if err != nil {
		return nil, err
	}
	return socialProfileRegistrations, nil
}
