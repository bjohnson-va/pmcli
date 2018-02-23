package mssdk

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vendasta/gosdks/basesdk"
	"github.com/vendasta/gosdks/util"
)

func Test_ActivateReturnsErrorIfNoPartnerIDProvided(t *testing.T) {
	client := &AccountClient{SDKClient: &basesdk.BaseClientMock{JSONBody: "{}"}}

	_, err := client.Activate(context.Background(), "accountGroupID", "", "")

	assert.Equal(t, err, util.Error(util.InvalidArgument, "partnerID is required"))
}

func Test_ActivateShouldCallPresenceBuilderApiWithCorrectParams(t *testing.T) {
	baseMock := &basesdk.BaseClientMock{JSONBody: "{}"}
	client := &AccountClient{SDKClient: baseMock}
	client.Activate(context.Background(), "AG-123", "ABC", "sso-token")

	expectedParams := map[string]interface{}{
		"agid":     "AG-123",
		"ssoToken": "sso-token",
		"pid":      "ABC",
	}
	if !reflect.DeepEqual(baseMock.ParamsSent, expectedParams) {
		t.Errorf("Expected params: %v, but got %v", expectedParams, baseMock.ParamsSent)
	}
}

var activateAccountResponse = `{
	"data": {
		"msid": "AC-JWGGGWH7"
	}
}`

func Test_ActivateShouldReturnTheAccountIDOnSuccessFromPresenceBuilder(t *testing.T) {
	baseMock := &basesdk.BaseClientMock{JSONBody: activateAccountResponse}
	client := &AccountClient{SDKClient: baseMock}
	result, err := client.Activate(context.Background(), "AG-123", "ABC", "")

	assert.Nil(t, err, "There should be no error returned on a successful activation")

	assert.Equal(t, "AC-JWGGGWH7", result)
}

func Test_ActivateReturnsErrorIfPresenceBuilderReturnsError(t *testing.T) {
	msErr := errors.New("LB is down")
	baseMock := &basesdk.BaseClientMock{Error: msErr}
	client := &AccountClient{SDKClient: baseMock}
	_, err := client.Activate(context.Background(), "AG-123", "ABC", "")

	assert.Equal(t, msErr, err)
}

func Test_ActivateReturnsErrorIfUnmarshalingDataReturnsError(t *testing.T) {
	baseMock := &basesdk.BaseClientMock{JSONBody: "garbage json"}
	client := &AccountClient{SDKClient: baseMock}
	_, err := client.Activate(context.Background(), "AG-123", "ABC", "")

	assert.NotNil(t, err)
}
