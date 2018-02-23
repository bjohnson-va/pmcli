package smsdk

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vendasta/gosdks/basesdk"
	"github.com/vendasta/gosdks/util"
)

func Test_CreateReturnsErrorIfNoAccountGroupIDProvided(t *testing.T) {
	client := &AccountClient{SDKClient: &basesdk.BaseClientMock{JSONBody: "{}"}}

	_, err := client.Create(context.Background(), "", "partnerID")

	assert.Equal(t, err, util.Error(util.InvalidArgument, "accountGroupID is required"))
}

func Test_CreateReturnsErrorIfNoPartnerIDProvided(t *testing.T) {
	client := &AccountClient{SDKClient: &basesdk.BaseClientMock{JSONBody: "{}"}}

	_, err := client.Create(context.Background(), "accountGroupID", "")

	assert.Equal(t, err, util.Error(util.InvalidArgument, "partnerID is required"))
}

func Test_CreateShouldCallSocialMarketingApiWithCorrectParams(t *testing.T) {
	baseMock := &basesdk.BaseClientMock{JSONBody: "{}"}
	client := &AccountClient{SDKClient: baseMock}
	client.Create(context.Background(), "AG-123", "ABC")

	expectedParams := map[string]interface{}{
		"country":        "ca",
		"accountGroupId": "AG-123",
		"address":        "dummy",
		"city":           "dummy",
		"state":          "sk",
		"zip":            "dummy",
		"companyName":    "dummy",
		"pid":            "ABC",
	}
	if !reflect.DeepEqual(baseMock.ParamsSent, expectedParams) {
		t.Errorf("Expected params: %v, but got %v", expectedParams, baseMock.ParamsSent)
	}
}

var createAccountResponse = `{
"data": {
    "accountGroupId": "AG-KLNC662J", 
    "accountId": "AC-JWGGGWH7", 
    "address": "555 white rd", 
    "billingCode": null, 
    "callTrackingNumber": [
        "7777777777", 
        "8888888888"
    ], 
    "city": "Toronto", 
    "companyName": "company", 
    "country": "Canada", 
    "createdDateTime": "2014-01-15T22:35:43Z", 
    "customerIdentifier": null, 
    "latitude": "52.125948", 
    "longitude": "-106.663286", 
    "marketId": "Eastern Canada", 
    "pid": "ABC", 
    "ssoToken": "AC-JWGGGWH7", 
    "state": "Ontario", 
    "taxonomyId": [
        "Other"
    ], 
    "updatedDateTime": "2014-01-15T22:35:43Z", 
    "workNumber": [
        "5555555555", 
        "5555555666"
    ], 
    "zip": "H0H 0H0"
}}
`

func Test_CreateShouldReturnTheAccountOnSuccessFromSocialMarketing(t *testing.T) {
	baseMock := &basesdk.BaseClientMock{JSONBody: createAccountResponse}
	client := &AccountClient{SDKClient: baseMock}
	result, err := client.Create(context.Background(), "AG-123", "ABC")

	assert.Nil(t, err, "There should be no error returned on a successful create")

	expectedAccount := &Account{
		AccountGroupID: "AG-KLNC662J",
		AccountID:      "AC-JWGGGWH7",
		PartnerID:      "ABC",
		MarketID:       "Eastern Canada",
		SsoToken:       "AC-JWGGGWH7",
	}
	assert.Equal(t, expectedAccount, result)
}

func Test_CreateReturnsErrorIfSocialMarketingReturnsError(t *testing.T) {
	smErr := errors.New("SM is down")
	baseMock := &basesdk.BaseClientMock{Error: smErr}
	client := &AccountClient{SDKClient: baseMock}
	_, err := client.Create(context.Background(), "AG-123", "ABC")

	assert.Equal(t, smErr, err)
}

func Test_CreateReturnsErrorIfUnmarshalingDataReturnsError(t *testing.T) {
	baseMock := &basesdk.BaseClientMock{JSONBody: "garbage json"}
	client := &AccountClient{SDKClient: baseMock}
	_, err := client.Create(context.Background(), "AG-123", "ABC")

	assert.NotNil(t, err)
}

func Test_ActivateReturnsErrorIfNoAccountGroupIDProvided(t *testing.T) {
	client := &AccountClient{SDKClient: &basesdk.BaseClientMock{JSONBody: "{}"}}

	_, err := client.Activate(context.Background(), "", "partnerID", "")

	assert.Equal(t, err, util.Error(util.InvalidArgument, "accountGroupID is required"))
}

func Test_ActivateReturnsErrorIfNoPartnerIDProvided(t *testing.T) {
	client := &AccountClient{SDKClient: &basesdk.BaseClientMock{JSONBody: "{}"}}

	_, err := client.Activate(context.Background(), "accountGroupID", "", "")

	assert.Equal(t, err, util.Error(util.InvalidArgument, "partnerID is required"))
}

func Test_ActivateShouldCallSocialMarketingApiWithCorrectParams(t *testing.T) {
	baseMock := &basesdk.BaseClientMock{JSONBody: "{}"}
	client := &AccountClient{SDKClient: baseMock}
	client.Activate(context.Background(), "AG-123", "ABC", "sso-token")

	expectedParams := map[string]interface{}{
		"accountGroupId": "AG-123",
		"ssoToken":       "sso-token",
		"partnerId":      "ABC",
	}
	if !reflect.DeepEqual(baseMock.ParamsSent, expectedParams) {
		t.Errorf("Expected params: %v, but got %v", expectedParams, baseMock.ParamsSent)
	}
}

var activateAccountResponse = `{
	"data": {
		"accountId": "AC-JWGGGWH7"
	}
}`

func Test_ActivateShouldReturnTheAccountIDOnSuccessFromSocialMarketing(t *testing.T) {
	baseMock := &basesdk.BaseClientMock{JSONBody: activateAccountResponse}
	client := &AccountClient{SDKClient: baseMock}
	result, err := client.Activate(context.Background(), "AG-123", "ABC", "")

	assert.Nil(t, err, "There should be no error returned on a successful activation")

	assert.Equal(t, "AC-JWGGGWH7", result)
}

func Test_ActivateReturnsErrorIfSocialMarketingReturnsError(t *testing.T) {
	smErr := errors.New("SM is down")
	baseMock := &basesdk.BaseClientMock{Error: smErr}
	client := &AccountClient{SDKClient: baseMock}
	_, err := client.Activate(context.Background(), "AG-123", "ABC", "")

	assert.Equal(t, smErr, err)
}

func Test_ActivateReturnsErrorIfUnmarshalingDataReturnsError(t *testing.T) {
	baseMock := &basesdk.BaseClientMock{JSONBody: "garbage json"}
	client := &AccountClient{SDKClient: baseMock}
	_, err := client.Activate(context.Background(), "AG-123", "ABC", "")

	assert.NotNil(t, err)
}
