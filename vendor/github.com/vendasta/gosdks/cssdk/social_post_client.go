package cssdk

import (
	"net/http"
	"time"

	"github.com/vendasta/gosdks/basesdk"
	"github.com/vendasta/gosdks/config"
	"github.com/vendasta/gosdks/util"
	"github.com/vendasta/gosdks/validation"
	"golang.org/x/net/context"
)

const (
	scheduleFacebookPostPath                       = "/internalApi/v2/socialPost/scheduleFacebookPost/"
	scheduleTwitterPostPath                        = "/internalApi/v2/socialPost/scheduleTwitterPost/"
	scheduleGooglePlusPostPath                     = "/internalApi/v2/socialPost/scheduleGooglePlusPost/"
	scheduleLinkedinPostPath                       = "/internalApi/v2/socialPost/scheduleLinkedInPost/"
	listSocialPostPath                             = "/internalApi/v2/socialPost/lookup/"
	listPartnerScheduledPostsPath                  = "/internalApi/v2/socialPost/scheduled/lookup/"
	listPartnerScheduledPostsByCreatedDateTimePath = "/internalApi/v2/socialPost/scheduled/lookup/byCreatedDate/"
)

// SocialPostClientInterface defines the interface of a social post client
type SocialPostClientInterface interface {
	ScheduleFacebookPost(ctx context.Context, accountGroupID string, postText string, socialPostID string, socialServiceID string, postDateTime time.Time, imageURL string) (*SocialPost, error)
	ScheduleTwitterPost(ctx context.Context, accountGroupID string, postText string, socialPostID string, socialServiceID string, postDateTime time.Time, imageURL string) (*SocialPost, error)
	ScheduleGooglePlusPost(ctx context.Context, accountGroupID string, postText string, socialPostID string, socialServiceID string, postDateTime time.Time, imageURL string) (*SocialPost, error)
	ScheduleLinkedinPost(ctx context.Context, accountGroupID string, postText string, socialPostID string, socialServiceID string, postDateTime time.Time, imageURL string) (*SocialPost, error)
	ListSocialPosts(ctx context.Context, accountGroupID string, startDateTime time.Time, endDateTime time.Time, socialServiceIDs []string, cursor string, pageSize int) ([]*SocialPost, string, error)
	ListPartnerScheduledPosts(ctx context.Context, partnerID string, startDateTime time.Time, endDateTime time.Time, cursor string, pageSize int) ([]*SocialPost, string, error)
	ListPartnerScheduledPostsByCreatedDateTime(ctx context.Context, partnerID string, startDateTime time.Time, endDateTime time.Time, cursor string, pageSize int) ([]*SocialPost, string, error)
}

// SocialPostClient is a client which handles calls to core services's social post apis, implements the interface
type SocialPostClient struct {
	basesdk.SDKClient
}

// BuildSocialPostClient creates a social post client.
func BuildSocialPostClient(apiUser string, apiKey string, env config.Env) SocialPostClient {
	auth := basesdk.UserKey{APIUser: apiUser, APIKey: apiKey}
	baseClient := basesdk.BaseClient{Authorization: auth, RootURL: rootURLFromEnv(env)}
	return SocialPostClient{baseClient}
}

//ScheduleFacebookPost schedule a post to facebook
func (c SocialPostClient) ScheduleFacebookPost(ctx context.Context, accountGroupID string, postText string, socialPostID string, socialServiceID string, postDateTime time.Time, imageURL string) (*SocialPost, error) {
	return c.schedulePost(ctx, accountGroupID, postText, socialPostID, socialServiceID, postDateTime, imageURL, scheduleFacebookPostPath)
}

//ScheduleTwitterPost schedule a post to twitter
func (c SocialPostClient) ScheduleTwitterPost(ctx context.Context, accountGroupID string, postText string, socialPostID string, socialServiceID string, postDateTime time.Time, imageURL string) (*SocialPost, error) {
	return c.schedulePost(ctx, accountGroupID, postText, socialPostID, socialServiceID, postDateTime, imageURL, scheduleTwitterPostPath)
}

//ScheduleGooglePlusPost schedule a post to google plus
func (c SocialPostClient) ScheduleGooglePlusPost(ctx context.Context, accountGroupID string, postText string, socialPostID string, socialServiceID string, postDateTime time.Time, imageURL string) (*SocialPost, error) {
	return c.schedulePost(ctx, accountGroupID, postText, socialPostID, socialServiceID, postDateTime, imageURL, scheduleGooglePlusPostPath)
}

//ScheduleLinkedinPost schedule a post to linkedin
func (c SocialPostClient) ScheduleLinkedinPost(ctx context.Context, accountGroupID string, postText string, socialPostID string, socialServiceID string, postDateTime time.Time, imageURL string) (*SocialPost, error) {
	return c.schedulePost(ctx, accountGroupID, postText, socialPostID, socialServiceID, postDateTime, imageURL, scheduleLinkedinPostPath)
}

func (c SocialPostClient) schedulePost(ctx context.Context, accountGroupID string, postText string, socialPostID string, socialServiceID string, postDateTime time.Time, imageURL string, path string) (*SocialPost, error) {
	err := validation.NewValidator().
		Rule(validation.StringNotEmpty(accountGroupID, util.InvalidArgument, "account group id is required")).
		Rule(validation.StringNotEmpty(socialServiceID, util.InvalidArgument, "social service id group id is required")).
		Rule(validation.AtLeastOneStringRequired([]string{postText, imageURL}, util.InvalidArgument, "Either postText or imageURL must be provided")).
		Validate()
	if err != nil {
		return nil, err
	}
	params := map[string]interface{}{
		"agid":         accountGroupID,
		"postText":     postText,
		"socialPostId": socialPostID,
		"ssid":         socialServiceID,
		"postDateTime": basesdk.ConvertTimeToVAPITimestamp(postDateTime),
		"imageUrl":     imageURL,
	}
	response, err := c.Post(ctx, path, params)
	if err != nil {
		return nil, err
	}
	socialPost, err := socialPostFromResponse(response)
	if err != nil {
		return nil, err
	}
	return socialPost, nil
}

//ListSocialPosts fetches a list of social posts from CS
func (c SocialPostClient) ListSocialPosts(ctx context.Context, accountGroupID string, startDateTime time.Time, endDateTime time.Time, socialServiceIDs []string, cursor string, pageSize int) ([]*SocialPost, string, error) {
	err := validation.NewValidator().
		Rule(validation.StringNotEmpty(accountGroupID, util.InvalidArgument, "account group id is required")).
		Validate()
	if err != nil {
		return nil, "", err
	}
	params := map[string]interface{}{
		"agid":          accountGroupID,
		"startDateTime": basesdk.ConvertTimeToVAPITimestamp(startDateTime),
		"endDateTime":   basesdk.ConvertTimeToVAPITimestamp(endDateTime),
		"ssid":          socialServiceIDs,
		"pageSize":      pageSize,
		"cursor":        cursor,
	}
	response, err := c.Get(ctx, listSocialPostPath, params, basesdk.Idempotent())
	return processListSocialPostsResponse(response, err)
}

func processListSocialPostsResponse(response *http.Response, err error) ([]*SocialPost, string, error) {
	if err != nil {
		return nil, "", err
	}
	socialPost, cursor, err := socialPostsFromResponse(response)
	if err != nil {
		return nil, "", err
	}
	return socialPost, cursor, nil
}

// ListPartnerScheduledPosts fetches a list of scheduled posts made across all accounts belonging to a partner filtered on scheduled datetime
func (c SocialPostClient) ListPartnerScheduledPosts(ctx context.Context, partnerID string, startDateTime time.Time, endDateTime time.Time, cursor string, pageSize int) ([]*SocialPost, string, error) {
	return c.listPartnerScheduledPosts(ctx, listPartnerScheduledPostsPath, partnerID, startDateTime, endDateTime, cursor, pageSize)
}

// ListPartnerScheduledPostsByCreatedDateTime fetches a list of scheduled posts made across all accounts belonging to a partner filtered on created datetime
func (c SocialPostClient) ListPartnerScheduledPostsByCreatedDateTime(ctx context.Context, partnerID string, startDateTime time.Time, endDateTime time.Time, cursor string, pageSize int) ([]*SocialPost, string, error) {
	return c.listPartnerScheduledPosts(ctx, listPartnerScheduledPostsByCreatedDateTimePath, partnerID, startDateTime, endDateTime, cursor, pageSize)
}

//Both PartnerScheduledPost Apis have the same interface, this is the helper method that does the work
func (c SocialPostClient) listPartnerScheduledPosts(ctx context.Context, path string, partnerID string, startDateTime time.Time, endDateTime time.Time, cursor string, pageSize int) ([]*SocialPost, string, error) {
	err := validation.NewValidator().
		Rule(validation.StringNotEmpty(partnerID, util.InvalidArgument, "partner id is required")).
		Validate()
	if err != nil {
		return nil, "", err
	}
	params := map[string]interface{}{
		"partnerId":     partnerID,
		"startDateTime": basesdk.ConvertTimeToVAPITimestamp(startDateTime),
		"endDateTime":   basesdk.ConvertTimeToVAPITimestamp(endDateTime),
		"pageSize":      pageSize,
		"cursor":        cursor,
	}
	response, err := c.Get(ctx, path, params, basesdk.Idempotent())
	return processListSocialPostsResponse(response, err)
}
