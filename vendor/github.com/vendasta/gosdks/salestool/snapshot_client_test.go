package salestool

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vendasta/gosdks/basesdk"
	"golang.org/x/net/context"
)

func TestCreateFromGooglePlaceReturnsAnErrorIfRequiredParamsMissing(t *testing.T) {
	baseClient := &basesdk.BaseClientMock{}
	client := SnapshotClient{SDKClient: baseClient}
	r := &CreateFromGooglePlaceRequest{
		WidgetID:         "SNAPSHOT-WIDGET-123",
		PlaceID:          "ChIJZXBHzOf2BFMRvDWDyvlI8h8",
		PartnerID:        "ABC",
		MarketID:         "market-1",
		SalespersonID:    "UID-1231231",
		ContactFirstName: "Joe",
		ContactLastName:  "Jackson",
		ContactEmail:     "joe@vendasta.com",
	}

	r.WidgetID = ""
	err := client.CreateFromGooglePlace(context.Background(), r)
	assert.EqualError(t, err, "widget id is required")
	r.WidgetID = "SNAPSHOT-WIDGET-123"

	r.PlaceID = ""
	err = client.CreateFromGooglePlace(context.Background(), r)
	assert.EqualError(t, err, "place id is required")
	r.PlaceID = "ChIJZXBHzOf2BFMRvDWDyvlI8h8"

	r.PartnerID = ""
	err = client.CreateFromGooglePlace(context.Background(), r)
	assert.EqualError(t, err, "partner id is required")
	r.PartnerID = "ABC"

	r.MarketID = ""
	err = client.CreateFromGooglePlace(context.Background(), r)
	assert.EqualError(t, err, "market id is required")
	r.MarketID = "market-1"

	r.ContactFirstName = ""
	err = client.CreateFromGooglePlace(context.Background(), r)
	assert.EqualError(t, err, "contact first name is required")
	r.ContactFirstName = "Joe"

	r.ContactEmail = ""
	err = client.CreateFromGooglePlace(context.Background(), r)
	assert.EqualError(t, err, "contact email is required")
	r.ContactEmail = "joe@vendasta.com"
}

func TestCreateFromGooglePlaceReturnsErrorIfPostFails(t *testing.T) {
	baseClient := &basesdk.BaseClientMock{Error: errors.New("Post Failed")}
	client := SnapshotClient{SDKClient: baseClient}
	r := &CreateFromGooglePlaceRequest{
		WidgetID:         "SNAPSHOT-WIDGET-123",
		PlaceID:          "ChIJZXBHzOf2BFMRvDWDyvlI8h8",
		PartnerID:        "ABC",
		MarketID:         "market-1",
		SalespersonID:    "UID-1231231",
		ContactFirstName: "Joe",
		ContactLastName:  "Jackson",
		ContactEmail:     "joe@vendasta.com",
	}
	err := client.CreateFromGooglePlace(context.Background(), r)
	assert.EqualError(t, err, "Post Failed")
}

func RequiredArgsCheckForCreate(t *testing.T, r *CreateRequest, m string) {
	baseClient := &basesdk.BaseClientMock{}
	client := SnapshotClient{SDKClient: baseClient}
	err := client.Create(context.Background(), r)
	assert.EqualError(t, err, m)
}

func TestCreateReturnsAnErrorIfRequiredParamsMissing(t *testing.T) {
	r := &CreateRequest{
		WidgetID:         "SNAPSHOT-WIDGET-123",
		PartnerID:        "ABC",
		MarketID:         "market-1",
		SalespersonID:    "UID-1231231",
		ContactFirstName: "Joe",
		ContactLastName:  "Jackson",
		ContactEmail:     "joe@vendasta.com",
		CompanyName:      "My Company",
		WorkNumber:       "3062341235",
		ZipCode:          "S0E1A0",
		TaxonomyID:       "other",
		Website:          "www.example.com",
		Address:          "220 3rd Ave",
		City:             "Saskatoon",
		State:            "SK",
		Country:          "CA",
		FacebookURL:      "www.facebook.com/url",
		TwitterURL:       "www.twitter.com/url",
	}

	r.WidgetID = ""
	RequiredArgsCheckForCreate(t, r, "widget id is required")
	r.WidgetID = "SNAPSHOT-WIDGET-123"

	r.PartnerID = ""
	RequiredArgsCheckForCreate(t, r, "partner id is required")
	r.PartnerID = "ABC"

	r.MarketID = ""
	RequiredArgsCheckForCreate(t, r, "market id is required")
	r.MarketID = "market-1"

	r.ContactFirstName = ""
	RequiredArgsCheckForCreate(t, r, "contact first name is required")
	r.ContactFirstName = "Joe"

	r.ContactEmail = ""
	RequiredArgsCheckForCreate(t, r, "contact email is required")
	r.ContactEmail = "joe@vendasta.com"

	r.CompanyName = ""
	RequiredArgsCheckForCreate(t, r, "company name is required")
	r.CompanyName = "My Company"

	r.Address = ""
	RequiredArgsCheckForCreate(t, r, "address is required")
	r.Address = "220 3rd Ave"

	r.City = ""
	RequiredArgsCheckForCreate(t, r, "city is required")
	r.City = "Saskatoon"

	r.Country = ""
	RequiredArgsCheckForCreate(t, r, "country is required")
	r.Country = "CA"

	r.WorkNumber = ""
	RequiredArgsCheckForCreate(t, r, "work number is required")
	r.WorkNumber = "3062341235"

	r.ZipCode = ""
	RequiredArgsCheckForCreate(t, r, "zip code is required")
	r.ZipCode = "S0E1A0"

	r.TaxonomyID = ""
	RequiredArgsCheckForCreate(t, r, "taxonomy id is required")
	r.TaxonomyID = "others"
}

func TestCreateReturnsErrorIfPostFails(t *testing.T) {
	baseClient := &basesdk.BaseClientMock{Error: errors.New("Post Failed")}
	client := SnapshotClient{SDKClient: baseClient}
	r := &CreateRequest{
		WidgetID:         "SNAPSHOT-WIDGET-123",
		PartnerID:        "ABC",
		MarketID:         "market-1",
		SalespersonID:    "UID-1231231",
		ContactFirstName: "Joe",
		ContactLastName:  "Jackson",
		ContactEmail:     "joe@vendasta.com",
		CompanyName:      "My Company",
		WorkNumber:       "3062341235",
		ZipCode:          "S0E1A0",
		TaxonomyID:       "other",
		Website:          "www.example.com",
		Address:          "220 3rd Ave",
		City:             "Saskatoon",
		State:            "SK",
		Country:          "CA",
		FacebookURL:      "www.facebook.com/url",
		TwitterURL:       "www.twitter.com/url",
	}
	err := client.Create(context.Background(), r)
	assert.EqualError(t, err, "Post Failed")
}
