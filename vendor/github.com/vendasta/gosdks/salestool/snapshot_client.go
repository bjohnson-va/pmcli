package salestool

import (
	"errors"

	"github.com/vendasta/gosdks/basesdk"
	"golang.org/x/net/context"
)

const (
	createAccountPath                = "/internalApi/v1/snapshot/create/"
	createAccountFromGooglePlacePath = "/internalApi/v1/snapshot/create/from-google-place/"
)

// SnapshotClient is a client which handles calls to salestool's snapshot apis
type SnapshotClient struct {
	basesdk.SDKClient
}

// CreateFromGooglePlaceRequest defines structure of request for create account from google place id
type CreateFromGooglePlaceRequest struct {
	WidgetID         string
	PartnerID        string
	MarketID         string
	CampaignID       string
	SalespersonID    string
	ContactFirstName string
	ContactLastName  string
	ContactEmail     string
	PlaceID          string
}

// CreateRequest defines structure of request for create account from NAP data
type CreateRequest struct {
	WidgetID         string
	PartnerID        string
	MarketID         string
	CampaignID       string
	SalespersonID    string
	ContactFirstName string
	ContactLastName  string
	ContactEmail     string
	CompanyName      string
	WorkNumber       string
	ZipCode          string
	TaxonomyID       string
	Website          string
	Address          string
	City             string
	State            string
	Country          string
	FacebookURL      string
	TwitterURL       string
}

// BuildSnapshotClient creates a snapshot client
func BuildSnapshotClient(apiUser string, apiKey string, rootURL string) SnapshotClient {
	auth := basesdk.UserKey{APIUser: apiUser, APIKey: apiKey}
	baseClient := basesdk.BaseClient{Authorization: auth, RootURL: rootURL}
	return SnapshotClient{baseClient}
}

// CreateFromGooglePlace creates snapshot account from place id
func (c SnapshotClient) CreateFromGooglePlace(ctx context.Context, r *CreateFromGooglePlaceRequest) error {
	if r.WidgetID == "" {
		return errors.New("widget id is required")
	}
	if r.PlaceID == "" {
		return errors.New("place id is required")
	}
	if r.PartnerID == "" {
		return errors.New("partner id is required")
	}
	if r.MarketID == "" {
		return errors.New("market id is required")
	}
	if r.ContactFirstName == "" {
		return errors.New("contact first name is required")
	}
	if r.ContactEmail == "" {
		return errors.New("contact email is required")
	}
	params := map[string]interface{}{
		"widgetId":         r.WidgetID,
		"placeId":          r.PlaceID,
		"partnerId":        r.PartnerID,
		"marketId":         r.MarketID,
		"contactFirstName": r.ContactFirstName,
		"contactEmail":     r.ContactEmail,
	}

	if r.SalespersonID != "" {
		params["salespersonId"] = r.SalespersonID
	}
	if r.ContactLastName != "" {
		params["contactLastName"] = r.ContactLastName
	}
	if r.CampaignID != "" {
		params["campaignId"] = r.CampaignID
	}
	_, err := c.Post(ctx, createAccountFromGooglePlacePath, params)
	if err != nil {
		return err
	}
	return nil
}

// Create creates snapshot account from NAP data
func (c SnapshotClient) Create(ctx context.Context, r *CreateRequest) error {
	if r.WidgetID == "" {
		return errors.New("widget id is required")
	}
	if r.PartnerID == "" {
		return errors.New("partner id is required")
	}
	if r.MarketID == "" {
		return errors.New("market id is required")
	}
	if r.ContactFirstName == "" {
		return errors.New("contact first name is required")
	}
	if r.ContactEmail == "" {
		return errors.New("contact email is required")
	}
	if r.CompanyName == "" {
		return errors.New("company name is required")
	}
	if r.Address == "" {
		return errors.New("address is required")
	}
	if r.City == "" {
		return errors.New("city is required")
	}
	if r.Country == "" {
		return errors.New("country is required")
	}
	if r.WorkNumber == "" {
		return errors.New("work number is required")
	}
	if r.ZipCode == "" {
		return errors.New("zip code is required")
	}
	if r.TaxonomyID == "" {
		return errors.New("taxonomy id is required")
	}

	params := map[string]interface{}{
		"widgetId":         r.WidgetID,
		"partnerId":        r.PartnerID,
		"marketId":         r.MarketID,
		"contactFirstName": r.ContactFirstName,
		"contactEmail":     r.ContactEmail,
		"companyName":      r.CompanyName,
		"workNumber":       r.WorkNumber,
		"zipCode":          r.ZipCode,
		"taxonomyId":       r.TaxonomyID,
	}

	if r.SalespersonID != "" {
		params["salespersonId"] = r.SalespersonID
	}
	if r.ContactLastName != "" {
		params["contactLastName"] = r.ContactLastName
	}
	if r.CampaignID != "" {
		params["campaignId"] = r.CampaignID
	}
	if r.Website != "" {
		params["website"] = r.Website
	}
	if r.Address != "" {
		params["address"] = r.Address
	}
	if r.City != "" {
		params["city"] = r.City
	}
	if r.State != "" {
		params["state"] = r.State
	}
	if r.Country != "" {
		params["country"] = r.Country
	}
	if r.FacebookURL != "" {
		params["facebookUrl"] = r.FacebookURL
	}
	if r.TwitterURL != "" {
		params["twitterUrl"] = r.TwitterURL
	}

	_, err := c.Post(ctx, createAccountPath, params)
	if err != nil {
		return err
	}
	return nil
}
