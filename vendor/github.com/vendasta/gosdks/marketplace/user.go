package marketplace

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/vendasta/gosdks/basesdk"
	"github.com/vendasta/gosdks/config"
	"github.com/vendasta/gosdks/logging"
	"github.com/vendasta/gosdks/util"
	"github.com/vendasta/gosdks/validation"
	"golang.org/x/net/context"
)

// GetUserResponse implements the response struct get from marketplace
type GetUserResponse struct {
	Took int   `json:"took"`
	User *User `json:"data"`
}

// User implements the user struct get from marketplace
type User struct {
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	WorkPhone string   `json:"work_phone"`
	Email     string   `json:"email"`
	Accounts  []string `json:"accounts"`
	PartnerID string   `json:"string"`
	ID        string   `json:"id"`
}

// UserClient is a client which handles calls to marketplace's User apis
type userClient struct {
	basesdk.SDKClient
}

// UserInterface is the interface of user client
type UserInterface interface {
	GetUser(ctx context.Context, userID string) (*User, error)
	GetUserAccountPermission(ctx context.Context, userID string, accountID string) (bool, error)
}

// NewUserClient creats user client with marketplace auth
func NewUserClient(ctx context.Context, appID string, privateKey []byte, env config.Env, rootURLOverride string) UserInterface {
	var rootURL string
	if rootURLOverride != "" {
		rootURL = rootURLOverride
	} else {
		rootURL = rootURLFromEnv(env)
	}
	oauthClient := NewOAuthClient(env, rootURL)
	return &userClient{
		basesdk.BaseClient{
			Authorization: NewMarketplaceAuthorization(ctx, appID, privateKey, oauthClient, nil),
			RootURL:       rootURL,
		},
	}
}

// GetUser fetches and returns user
func (u *userClient) GetUser(ctx context.Context, userID string) (*User, error) {
	err := validation.NewValidator().Rule(validation.StringNotEmpty(userID, util.InvalidArgument, "User ID cannot be empty")).Validate()
	if err != nil {
		return nil, err
	}
	path := u.buildGetUserAPIURL(userID)
	params := map[string]interface{}{}
	response, err := u.Get(ctx, path, params, basesdk.Idempotent())
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	reqBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("Malformed request body: %s", err.Error())
	}

	getUserResponse := GetUserResponse{}
	if err := json.Unmarshal(reqBody, &getUserResponse); err != nil {
		reason := "failed to convert response to GetUserResponse: " + err.Error()
		return nil, errors.New(reason)
	}
	return getUserResponse.User, nil

}

func (u *userClient) GetUserAccountPermission(ctx context.Context, userID string, accountID string) (bool, error) {
	err := validation.NewValidator().
		Rule(validation.StringNotEmpty(userID, util.InvalidArgument, "User ID cannot be empty")).
		Rule(validation.StringNotEmpty(accountID, util.InvalidArgument, "Account ID cannot be empty")).
		Validate()
	if err != nil {
		return false, err
	}
	path := u.buildGetUserAccountPermissionAPIURL(userID, accountID)
	err = u.Head(ctx, path, basesdk.Idempotent())
	if err != nil {
		logging.Debugf(ctx, "Error getting user %s account permission for %s: %s", userID, accountID, err.Error())
		return false, err
	}

	return true, nil
}

func (u *userClient) buildGetUserAPIURL(userID string) string {
	return fmt.Sprintf("/api/v1/user/%s", userID)
}

func (u *userClient) buildGetUserAccountPermissionAPIURL(userID string, accountID string) string {
	return fmt.Sprintf("/api/v1/user/%s/permissions/%s", userID, accountID)
}
