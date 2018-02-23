package whitelabel

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/vendasta/gosdks/basesdk"
	"github.com/vendasta/gosdks/config"
	"github.com/vendasta/gosdks/util"
	"github.com/vendasta/gosdks/validation"
	"golang.org/x/net/context"
)

// WhitelabelData contains whitelabel information for a partner/market
type WhitelabelData struct {
	SocialProfileGroupID string
	MaxRMCompetitors     int64
	MaxRMServices        int64

	UITheme      string
	PrimaryColor string

	FaviconSecureURL string
	FaviconURL       string

	MobileShortcutIconSecureURL string
	MobileShortcutIconURL       string

	LogoSecureURL string
	LogoURL       string

	BusinessCenterProductName       string
	ReputationManagementProductName string
	ListingBuilderProductName       string
	SocialMarketingProductName      string
	BrandAnalyticsProductName       string
}

const (
	getWhitelabelDataPath = "/internalApi/v3/configuration/get/"
)

// ErrorPartnerNotFound is returned when looking up the whitelabel data for a partner and it does not exist
var ErrorPartnerNotFound = util.Error(util.NotFound, "The partner id provided is not valid")

// ErrorMarketNotFound is returned when looking up whitelabel data for a market and it doesn't exist
var ErrorMarketNotFound = util.Error(util.NotFound, "The partner id and market id pair provided is not valid")

type getConfiguration struct {
	unmerged bool
}

// GetOption is a definition for an option provided to Get
type GetOption func(*getConfiguration)

// Unmerged is a get option which prevents market whitelabel data from being merged with the partners
func Unmerged() GetOption {
	return func(c *getConfiguration) {
		c.unmerged = true
	}
}

// WhitelabelClientInterface defines the interface for retrieving white label data for a partner/market
type WhitelabelClientInterface interface {
	// Get will return the whitelabel data for a partner/market.
	Get(ctx context.Context, partnerID string, marketID string, options ...GetOption) (*WhitelabelData, error)
}

type whiteLabelClient struct {
	basesdk.SDKClient
}

// BuildWhiteLabelClient creates a white label client
func BuildWhiteLabelClient(apiUser string, apiKey string, env config.Env) WhitelabelClientInterface {
	auth := basesdk.UserKey{APIUser: apiUser, APIKey: apiKey}
	baseClient := basesdk.BaseClient{Authorization: auth, RootURL: rootURLFromEnv(env)}
	return whiteLabelClient{baseClient}
}

var envURLs = map[config.Env]string{
	config.Local: "http://10.200.10.1:8081",
	config.Test:  "http://partner-central-test.appspot.com",
	config.Demo:  "http://partner-central-demo.appspot.com",
	config.Prod:  "http://partner-central.appspot.com",
}

func rootURLFromEnv(env config.Env) string {
	return envURLs[env]
}

type getWhitelabelDataResponseHandler func(body io.Reader) (*sparseWhiteLabelData, error)

// Get returns the white label data.
func (c whiteLabelClient) Get(ctx context.Context, partnerID string, marketID string, options ...GetOption) (*WhitelabelData, error) {
	configuration := &getConfiguration{
		unmerged: false,
	}
	for _, opt := range options {
		opt(configuration)
	}

	var wd *WhitelabelData
	var err error
	if !configuration.unmerged {
		wd, err = c.get(ctx, whiteLabelDataFromResponse, partnerID, "", &WhitelabelData{})
		if err != nil {
			if basesdk.StatusCode(err) == http.StatusNotFound {
				return nil, ErrorPartnerNotFound
			}
			return nil, err
		}
	}

	if marketID != "" || configuration.unmerged {
		wd, err = c.get(ctx, whiteLabelDataFromResponse, partnerID, marketID, wd)
	}
	if err != nil {
		if basesdk.StatusCode(err) == http.StatusNotFound {
			return nil, ErrorMarketNotFound
		}
		if basesdk.StatusCode(err) == http.StatusBadRequest && strings.Contains(err.Error(), "marketId") {
			return nil, util.Error(util.InvalidArgument, "the market Id: %s is not valid", marketID)
		}
		return nil, err
	}

	return wd, nil
}

func (c whiteLabelClient) get(ctx context.Context, responseHandler getWhitelabelDataResponseHandler, partnerID, marketID string, wd *WhitelabelData) (*WhitelabelData, error) {
	err := validation.NewValidator().Rule(
		validation.StringNotEmpty(partnerID, util.InvalidArgument, "partnerID is required"),
	).Validate()
	if err != nil {
		return nil, err
	}
	params := map[string]interface{}{"partnerId": partnerID, "marketId": marketID, "returnVObjectFlag": "true"}
	response, err := c.SDKClient.Get(ctx, getWhitelabelDataPath, params, basesdk.Idempotent())
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	sparseData, err := responseHandler(response.Body)
	if err != nil {
		return nil, err
	}

	if wd == nil {
		wd = &WhitelabelData{}
	}
	if sparseData.SocialProfileGroupID != nil {
		wd.SocialProfileGroupID = *sparseData.SocialProfileGroupID
	}
	if sparseData.MaxRMServices != nil {
		wd.MaxRMServices = *sparseData.MaxRMServices
	}
	if sparseData.MaxRMCompetitors != nil {
		wd.MaxRMCompetitors = *sparseData.MaxRMCompetitors
	}
	if sparseData.UITheme != nil {
		wd.UITheme = *sparseData.UITheme
	}
	if sparseData.PrimaryColor != nil {
		wd.PrimaryColor = *sparseData.PrimaryColor
	}
	if sparseData.FaviconURL != nil {
		wd.FaviconURL = *sparseData.FaviconURL
	}
	if sparseData.FaviconSecureURL != nil {
		wd.FaviconSecureURL = *sparseData.FaviconSecureURL
	}
	if sparseData.MobileShortcutIconSecureURL != nil {
		wd.MobileShortcutIconSecureURL = *sparseData.MobileShortcutIconSecureURL
	}
	if sparseData.MobileShortcutIconURL != nil {
		wd.MobileShortcutIconURL = *sparseData.MobileShortcutIconURL
	}
	if sparseData.LogoURL != nil {
		wd.LogoURL = *sparseData.LogoURL
	}
	if sparseData.LogoSecureURL != nil {
		wd.LogoSecureURL = *sparseData.LogoSecureURL
	}
	if sparseData.BusinessCenterProductName != nil {
		wd.BusinessCenterProductName = *sparseData.BusinessCenterProductName
	}
	if sparseData.ReputationManagementProductName != nil {
		wd.ReputationManagementProductName = *sparseData.ReputationManagementProductName
	}
	if sparseData.ListingBuilderProductName != nil {
		wd.ListingBuilderProductName = *sparseData.ListingBuilderProductName
	}
	if sparseData.SocialMarketingProductName != nil {
		wd.SocialMarketingProductName = *sparseData.SocialMarketingProductName
	}
	if sparseData.BrandAnalyticsProductName != nil {
		wd.BrandAnalyticsProductName = *sparseData.BrandAnalyticsProductName
	}

	return wd, nil
}

func whiteLabelDataFromResponse(body io.Reader) (*sparseWhiteLabelData, error) {
	type Response struct {
		Data string `json:"data"`
	}
	res := Response{}
	if err := json.NewDecoder(body).Decode(&res); err != nil {
		reason := "Failed to convert response to white label data: %s" + err.Error()
		return nil, errors.New(reason)
	}

	whitelabelData := &sparseWhiteLabelData{}
	err := json.Unmarshal([]byte(res.Data), whitelabelData)
	if err != nil {
		return nil, err
	}

	return whitelabelData, nil
}

type sparseWhiteLabelData struct {
	SocialProfileGroupID *string `json:"social_profile_group_id"`
	MaxRMCompetitors     *int64  `json:"rm_competition_max_competitors"`
	MaxRMServices        *int64  `json:"rm_competition_max_services"`

	UITheme      *string `json:"ui_theme"`
	PrimaryColor *string `json:"primary_color"`

	FaviconSecureURL *string `json:"favicon_secure_url"`
	FaviconURL       *string `json:"favicon_url"`

	MobileShortcutIconSecureURL *string `json:"mobile_shortcut_icon_secure_url"`
	MobileShortcutIconURL       *string `json:"mobile_shortcut_icon_url"`

	LogoSecureURL *string `json:"logo_secure_url"`
	LogoURL       *string `json:"logo_url"`

	BusinessCenterProductName       *string `json:"vbc_product_name"`
	ReputationManagementProductName *string `json:"rm_product_name"`
	ListingBuilderProductName       *string `json:"ms_product_name"`
	SocialMarketingProductName      *string `json:"sm_product_name"`
	BrandAnalyticsProductName       *string `json:"nb_product_name"`
}
