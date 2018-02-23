package marketplace

import (
	"encoding/json"
	"io/ioutil"

	"github.com/vendasta/gosdks/basesdk"
	"github.com/vendasta/gosdks/config"
	"github.com/vendasta/gosdks/util"
	"github.com/vendasta/gosdks/validation"
	"golang.org/x/net/context"
)

const (
	newActivityPath = "/api/v1/activity/"
)

//CreateActivityResponse implements create activity response
type CreateActivityResponse struct {
	ActivityID string `json:"activity_id"`
	AppID      string `json:"app_id"`
	Created    string `json:"created"`

	Activity
}

// ActivityClient implements the activity client interface
type ActivityClient interface {
	CreateActivity(ctx context.Context, activity Activity) (*CreateActivityResponse, error)
}

type createActivityResponseWrapper struct {
	Data CreateActivityResponse `json:"data"`
}

type activityClient struct {
	basesdk.SDKClient
}

// NewActivityClient creates an activity client to allow calls to be made to marketplace.
func NewActivityClient(ctx context.Context, appID string, privateKey []byte, env config.Env, rootURLOverride string) ActivityClient {
	var rootURL string
	if rootURLOverride != "" {
		rootURL = rootURLOverride
	} else {
		rootURL = rootURLFromEnv(env)
	}
	oauthClient := NewOAuthClient(env, rootURL)
	return activityClient{
		basesdk.BaseClient{
			Authorization: NewMarketplaceAuthorization(ctx, appID, privateKey, oauthClient, nil),
			RootURL:       rootURL,
		},
	}
}

// Activity represents an activity-stream activity
type Activity struct {
	// AccountID is the ID of the account
	AccountID string `json:"account_id"`

	// Type is the Activity Type of the Activity. It must be configured by marketplace
	Type     string `json:"activity_type"`
	Title    string `json:"title"`
	Link     string `json:"link"`
	Content  string `json:"content"`
	MediaURL string `json:"media_url"`

	// SettingsTags allow users more granular control over which notifications they receive
	SettingsTags []string `json:"settings_tags"`
	// RequiresPlatformAuth specifies whether the Link is within a Vendasta App and requires authorization to access
	RequiresPlatformAuth bool `json:"requires_platform_auth"`
}

// Validate ensures the activity is valid for submission to the Activity Stream API
func (activity Activity) Validate() error {
	err := validation.NewValidator().Rule(
		validation.StringNotEmpty(activity.AccountID, util.FailedPrecondition, "Marketplace > CreateActivity: AccountID is required"),
		validation.StringNotEmpty(activity.Type, util.FailedPrecondition, "Marketplace > CreateActivity: Type is required"),
		validation.StringNotEmpty(activity.Title, util.FailedPrecondition, "Marketplace > CreateActivity: Title is required"),
	).ValidateAndJoinErrors()
	if err != nil {
		return err
	}

	if activity.MediaURL != "" {
		err = validation.ValidURL(activity.MediaURL, util.FailedPrecondition, "Marketplace > CreateActivity: MediaURL is invalid").Validate()
		if err != nil {
			return err
		}
	}
	if activity.Link != "" {
		err = validation.ValidURL(activity.Link, util.FailedPrecondition, "Marketplace > NewActivity: Link is invalid").Validate()
		if err != nil {
			return err
		}
	}
	return nil
}

// CreateActivity creates a new Activity in the Activity Stream
func (c activityClient) CreateActivity(ctx context.Context, activity Activity) (*CreateActivityResponse, error) {
	err := activity.Validate()
	if err != nil {
		return nil, err
	}

	params := map[string]interface{}{
		"account_id":             activity.AccountID,
		"activity_type":          activity.Type,
		"title":                  activity.Title,
		"link":                   activity.Link,
		"content":                activity.Content,
		"media_url":              activity.MediaURL,
		"settings_tags":          activity.SettingsTags,
		"requires_platform_auth": activity.RequiresPlatformAuth,
	}

	response, err := c.Post(ctx, newActivityPath, params, basesdk.Idempotent())
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	activityResponse := &createActivityResponseWrapper{}
	err = json.Unmarshal(body, activityResponse)
	if err != nil {
		return nil, err
	}
	return &activityResponse.Data, nil
}
