package cssdk

import (
	"context"
	"errors"
	"testing"

	"github.com/vendasta/gosdks/config"

	"github.com/vendasta/gosdks/basesdk"
)

var socialServicesResponse = `{"data": {
    "FB_USER": [
        {
            "agid": "AG-123654", 
            "clientTags": null, 
            "collectPostsFlag": true, 
            "collectStatsFlag": true, 
            "createdDateTime": "2017-05-31T20:19:30Z", 
            "disabledReason": null, 
            "expiryDateTime": "2017-06-05T00:19:30Z", 
            "facebookUserId": "iqurbewobqup1379nbva8yf", 
            "isAuthenticated": false, 
            "isDisabledFlag": null, 
            "modifiedDateTime": "2017-05-31T20:19:30Z", 
            "name": "USER-NAME", 
            "notes": null, 
            "permissions": [
                "manage_pages", 
                "read_insights"
            ], 
            "profileImageUrl": "http://www.facebook.com/!#/ajsngri24u9nabjanbiab", 
            "profileUrl": "http://www.facebook.com/USER-NAME", 
            "serviceId": "iqurbewobqup1379nbva8yf", 
            "serviceType": "FB_USER", 
            "socialSyncFlag": null, 
            "socialTokenBroken": null, 
            "spid": "SCP-5E5B0210F1384B44A14D067515A7950B", 
            "ssid": "FBU-iqurbewobqup1379nbva8yf"
        }
    ], 
    "TW_USER": [
        {
            "agid": "AG-123654", 
            "clientTags": null, 
            "collectPostsFlag": true, 
            "collectStatsFlag": true, 
            "createdDateTime": "2017-05-31T20:19:30Z", 
            "disabledReason": null, 
            "fullName": null, 
            "isAuthenticated": true, 
            "isDisabledFlag": null, 
            "modifiedDateTime": "2017-05-31T20:19:30Z", 
            "name": "USER-NAME", 
            "notes": null, 
            "profileImageUrl": "http://www.twitter.com/!#/ajsngri24u9nabjanbiab", 
            "profileUrl": "http://www.twitter.com/USER-NAME", 
            "serviceId": "678564567808645", 
            "serviceType": "TW_USER", 
            "socialSyncFlag": null, 
            "socialTokenBroken": null, 
            "spid": "SCP-5E5B0210F1384B44A14D067515A7950B", 
            "ssid": "TWU-678564567808645", 
            "twitterUserId": "678564567808645"
        }
    ]
}}`

func Test_ListSocialServicesReturnListOfSocialServicesOn200(t *testing.T) {

	client := BuildSocialServicesClient("user", "key", config.Local)
	client.SDKClient = &basesdk.BaseClientMock{JSONBody: socialServicesResponse}

	result, err := client.ListSocialServices(context.Background(), "accountGroupID")

	if err != nil {
		t.Errorf("Expected no error, got %s", err.Error())
	}
	if len(result) != 2 {
		t.Errorf("Expected 2 social services, got %d", len(result))
	}

	expected1 := &SocialService{
		SSID:            "TWU-678564567808645",
		Name:            "USER-NAME",
		Username:        "",
		AccountGroupID:  "AG-123654",
		IsAuthenticated: true,
		ProfileURL:      "http://www.twitter.com/USER-NAME",
		ProfileImageURL: "http://www.twitter.com/!#/ajsngri24u9nabjanbiab",
		ServiceType:     "TW_USER",
		SocialProfileID: "SCP-5E5B0210F1384B44A14D067515A7950B",
	}
	if *result[0] != *expected1 {
		t.Errorf("Expected: %#v as the first service, but got %#v", result[0], expected1)
	}

	expected2 := &SocialService{
		SSID:            "FBU-iqurbewobqup1379nbva8yf",
		Name:            "USER-NAME",
		Username:        "",
		AccountGroupID:  "AG-123654",
		IsAuthenticated: false,
		ProfileURL:      "http://www.facebook.com/USER-NAME",
		ProfileImageURL: "http://www.facebook.com/!#/ajsngri24u9nabjanbiab",
		ServiceType:     "FB_USER",
		SocialProfileID: "SCP-5E5B0210F1384B44A14D067515A7950B",
	}
	if *result[1] != *expected2 {
		t.Errorf("Expected: %v as the second service, but got %v", result[1], expected2)
	}
}

func Test_ListSocialServicesReturnsErrorWhenAccountGroupIDNotProvided(t *testing.T) {
	client := BuildSocialServicesClient("user", "key", config.Local)
	_, err := client.ListSocialServices(context.Background(), "")

	if err == nil {
		t.Errorf("Expected error, got no error")
	}
}

func Test_ListSocialServicesReturnsErrorWhenCoreReturnsError(t *testing.T) {
	client := BuildSocialServicesClient("user", "key", config.Local)
	expectedError := errors.New("New error")
	client.SDKClient = &basesdk.BaseClientMock{JSONBody: socialServicesResponse, Error: expectedError}

	_, err := client.ListSocialServices(context.Background(), "accountGroupID")

	if err != expectedError {
		t.Errorf("Expected error: %s, but got %s", expectedError.Error(), err.Error())
	}
}

func Test_ListSocialServicesReturnsErrorWhenInflatingResponseHasError(t *testing.T) {
	client := BuildSocialServicesClient("user", "key", config.Local)
	client.SDKClient = &basesdk.BaseClientMock{JSONBody: `{"data":"garbage"}`}

	_, err := client.ListSocialServices(context.Background(), "accountGroupID")

	if err == nil {
		t.Errorf("Expected error but got none")
	}
}
