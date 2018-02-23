package cssdk

import (
	"strconv"
	"time"

	"github.com/vendasta/gosdks/basesdk"
	"github.com/vendasta/gosdks/config"
	"github.com/vendasta/gosdks/util"
	"github.com/vendasta/gosdks/validation"
	"golang.org/x/net/context"
)

const (
	lookupReivewsPath  = "/internalApi/v2/review/lookup/"
	getReviewStatsPath = "/internalApi/v2/review/stats/"
)

// ReviewClientInterface defines the interface of a review client
type ReviewClientInterface interface {
	Lookup(context.Context, string, ...LookupOption) (*ReviewLookupResponse, error)
	GetStats(context.Context, string, string) (*ReviewStats, error)
}

type reviewClient struct {
	basesdk.SDKClient
}

// BuildReviewClient creates a review client.
func BuildReviewClient(apiUser string, apiKey string, env config.Env) ReviewClientInterface {
	auth := basesdk.UserKey{APIUser: apiUser, APIKey: apiKey}
	baseClient := basesdk.BaseClient{Authorization: auth, RootURL: rootURLFromEnv(env)}
	return reviewClient{baseClient}
}

type lookupOptions struct {
	cursor      string
	pageSize    int
	minDateTime time.Time
}

// LookupOption allows clients to optionally apply more specific filters to lookup.
type LookupOption func(o *lookupOptions)

//WithCursor uses the specified cursor as an input to a lookup API
func WithCursor(cursor string) LookupOption {
	return func(o *lookupOptions) {
		o.cursor = cursor
	}
}

//WithPageSize uses the specified pagesize as an input to a lookup API
func WithPageSize(pageSize int) LookupOption {
	return func(o *lookupOptions) {
		o.pageSize = pageSize
	}
}

//WithMinDateTime uses the specified minDateTime as an input to a lookup API
func WithMinDateTime(min time.Time) LookupOption {
	return func(o *lookupOptions) {
		o.minDateTime = min
	}
}

//Lookup fetches a list of reviews from CS
func (c reviewClient) Lookup(ctx context.Context, accountGroupID string, opts ...LookupOption) (*ReviewLookupResponse, error) {
	err := validation.NewValidator().
		Rule(validation.StringNotEmpty(accountGroupID, util.InvalidArgument, "account group id is required")).
		Validate()
	if err != nil {
		return nil, err
	}
	o := &lookupOptions{}
	for _, f := range opts {
		f(o)
	}
	params := map[string]interface{}{
		"agid": accountGroupID,
	}

	if o.cursor != "" {
		params["cursor"] = o.cursor
	}

	if o.pageSize > 0 {
		params["pageSize"] = strconv.FormatInt(int64(o.pageSize), 10)
	}

	if !o.minDateTime.IsZero() {
		params["minDateTime"] = basesdk.ConvertTimeToVAPITimestamp(o.minDateTime)
	}
	response, err := c.Get(ctx, lookupReivewsPath, params, basesdk.Idempotent())
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	return reviewLookupResponseFromResponse(response.Body)
}

func (c reviewClient) GetStats(ctx context.Context, accountGroupID string, sourceID string) (*ReviewStats, error) {
	err := validation.NewValidator().
		Rule(validation.StringNotEmpty(accountGroupID, util.InvalidArgument, "account group id is required")).
		Validate()
	if err != nil {
		return nil, err
	}

	params := map[string]interface{}{
		"agid":     accountGroupID,
		"sourceId": sourceID,
	}

	response, err := c.Get(ctx, getReviewStatsPath, params, basesdk.Idempotent())
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	return reviewReviewStatsFromResponse(response.Body)
}
