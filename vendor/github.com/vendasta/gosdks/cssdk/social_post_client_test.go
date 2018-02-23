package cssdk

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/vendasta/gosdks/config"

	"github.com/vendasta/gosdks/basesdk"
	"golang.org/x/net/context"
)

var scheduleSocialPostResponse = `{
	"data": {
		"username": "Star Hotel Restaurant",
		"postText": "#June is here and it's going to be a GREAT month!! Start it off with the very best right here at the Star! http://bit.ly/1DPJTWq",
		"attachments": [{
			"url": "http://www.thestarhotelwv.com/Menu/0/Menus.aspx",
			"imageUrl": "https://digitalmarketing.blob.core.windows.net/6099/skins/StarHotelD/images/StarHotelDesignSliced_01.png",
			"objectId": null,
			"objectType": "article"
		}],
		"userId": "112978898143021632739",
		"agid": "AG-MHM2HKZ4",
		"postedDateTime": null,
		"scheduledDateTime": "2017-06-01T15:13:17Z",
		"ssid": "GPP-DF704C2BDE8D47FA80E57731EDFEFEA2",
		"socialPostId": "GooglePlusPost-8bed45260a464ee9825c3d7822b3555b-GPP-DF704C2BDE8D47FA80E57731EDFEFEA2",
		"isError": null,
		"deletionStatus": "NONE",
		"status": "IN_PROGRESS",
		"multiLocationPostId": null,
		"tags": ["SCHEDULED"],
		"clientTags": [],
		"profileImageUrl": "https://lh3.googleusercontent.com/-q1Smh9d8d0g/AAAAAAAAAAM/AAAAAAAAAAA/3YaY0XeTIPc/photo.jpg?sz=50",
		"verb": null,
		"postedByOwner": true,
		"annotation": "Welcome to Star Hotel & Restaurant! We are located at 76 North Main Street in Franklin, WV 26807. Our phone number is 304-358-3580, and our fax is 304-358-3581. We look forward to seeing you!",
		"permalink": null,
		"name": "Star Hotel Restaurant",
		"imageUrl": "http://lh3.googleusercontent.com/4G1_uSeCq5TEx5LukNmT0jxxVWuqOA8mSvJOd8-EQkt9pSLuOx-IAiqvSTA-a6SqhZVmk4ZkOiqJDyHrCFkLulsXWw=s1600",
		"profileUrl": "https://plus.google.com/112978898143021632739",
		"postCreatedDateTime": "2017-06-01T14:46:24Z"
	}
}`

func Test_schedulePostReturnsSocialPostOnSuccess(t *testing.T) {

	client := BuildSocialPostClient("user", "key", config.Local)
	client.SDKClient = &basesdk.BaseClientMock{JSONBody: scheduleSocialPostResponse}

	result, err := client.schedulePost(context.Background(), "accountGroupID", "post text", "socialPostID", "SSID", time.Now(), "http://imageUrl.com", "/schedulePostToFacebook")

	if err != nil {
		t.Errorf("Expected no error, got %s", err.Error())
	}

	expected1 := &SocialPost{
		AccountGroupID:      "AG-MHM2HKZ4",
		IsError:             false,
		DeletionStatus:      "NONE",
		Permalink:           "",
		PostCreatedDateTime: "2017-06-01T14:46:24Z",
		PostText:            "#June is here and it's going to be a GREAT month!! Start it off with the very best right here at the Star! http://bit.ly/1DPJTWq",
		PostedDateTime:      "",
		ProfileURL:          "https://plus.google.com/112978898143021632739",
		ProfileImageURL:     "https://lh3.googleusercontent.com/-q1Smh9d8d0g/AAAAAAAAAAM/AAAAAAAAAAA/3YaY0XeTIPc/photo.jpg?sz=50",
		ScheduledDateTime:   "2017-06-01T15:13:17Z",
		SocialPostID:        "GooglePlusPost-8bed45260a464ee9825c3d7822b3555b-GPP-DF704C2BDE8D47FA80E57731EDFEFEA2",
		Ssid:                "GPP-DF704C2BDE8D47FA80E57731EDFEFEA2",
		Status:              "IN_PROGRESS",
		ImageURL:            "http://lh3.googleusercontent.com/4G1_uSeCq5TEx5LukNmT0jxxVWuqOA8mSvJOd8-EQkt9pSLuOx-IAiqvSTA-a6SqhZVmk4ZkOiqJDyHrCFkLulsXWw=s1600",
		Name:                "Star Hotel Restaurant",
		Username:            "Star Hotel Restaurant",
	}
	if *result != *expected1 {
		t.Errorf("Expected: %#v, but got %#v", result, expected1)
	}
}

func Test_schedulePostReturnsErrorWhenNoAccountGroupIdProvided(t *testing.T) {
	client := BuildSocialPostClient("user", "key", config.Local)
	client.SDKClient = &basesdk.BaseClientMock{JSONBody: scheduleSocialPostResponse}
	_, err := client.schedulePost(context.Background(), "", "post text", "socialPostID", "SSID", time.Now(), "http://imageUrl.com", "/schedulePostToFacebook")

	if err == nil {
		t.Errorf("Expected error, got no error")
	}
}

func Test_schedulePostReturnsErrorWhenNoSocialServiceIdProvided(t *testing.T) {
	client := BuildSocialPostClient("user", "key", config.Local)
	client.SDKClient = &basesdk.BaseClientMock{JSONBody: scheduleSocialPostResponse}
	_, err := client.schedulePost(context.Background(), "accountGroupId", "post text", "socialPostID", "", time.Now(), "http://imageUrl.com", "/schedulePostToFacebook")

	if err == nil {
		t.Errorf("Expected error, got no error")
	}
}

func Test_schedulePostReturnsErrorWhenNoPostTextOrImageUrlProvided(t *testing.T) {
	client := BuildSocialPostClient("user", "key", config.Local)
	client.SDKClient = &basesdk.BaseClientMock{JSONBody: scheduleSocialPostResponse}
	_, err := client.schedulePost(context.Background(), "accountGroupId", "", "socialPostID", "SSID", time.Now(), "", "/schedulePostToFacebook")

	if err == nil {
		t.Errorf("Expected error, got no error")
	}
}

func Test_schedulePostReturnsErrorWhenErrorConvertingResponse(t *testing.T) {
	client := BuildSocialPostClient("user", "key", config.Local)
	client.SDKClient = &basesdk.BaseClientMock{JSONBody: `{"data":"garbage"}`}
	_, err := client.schedulePost(context.Background(), "accountGroupId", "posttext", "socialPostID", "SSID", time.Now(), "", "/schedulePostToFacebook")

	if err == nil {
		t.Errorf("Expected error, got no error")
	}
}

func Test_schedulePostReturnsErrorWhenCoreReturnsError(t *testing.T) {
	client := BuildSocialPostClient("user", "key", config.Local)
	expectedError := errors.New("New error")
	client.SDKClient = &basesdk.BaseClientMock{JSONBody: scheduleSocialPostResponse, Error: expectedError}

	_, err := client.schedulePost(context.Background(), "accountGroupId", "posttext", "socialPostID", "SSID", time.Now(), "", "/schedulePostToFacebook")

	if err != expectedError {
		t.Errorf("Expected error: %s, but got %s", expectedError.Error(), err.Error())
	}
}

func Test_ScheduleFacebookPostMakesCorrectCallToCS(t *testing.T) {
	client := BuildSocialPostClient("user", "key", config.Local)
	mock := &basesdk.BaseClientMock{JSONBody: scheduleSocialPostResponse, Error: nil}
	client.SDKClient = mock
	now := time.Now()
	client.ScheduleFacebookPost(context.Background(), "accountGroupId", "posttext", "socialPostID", "SSID", now, "img.com")

	expectedParams := map[string]interface{}{
		"agid":         "accountGroupId",
		"postText":     "posttext",
		"socialPostId": "socialPostID",
		"ssid":         "SSID",
		"postDateTime": basesdk.ConvertTimeToVAPITimestamp(now),
		"imageUrl":     "img.com",
	}
	if !reflect.DeepEqual(mock.ParamsSent, expectedParams) {
		t.Errorf("Expected params: %v, but got %v", expectedParams, mock.ParamsSent)
	}
	if mock.PathRequested != scheduleFacebookPostPath {
		t.Errorf("Expected path: %s, but got %s", scheduleFacebookPostPath, mock.PathRequested)
	}
}

func Test_ScheduleTwitterPostMakesCorrectCallToCS(t *testing.T) {
	client := BuildSocialPostClient("user", "key", config.Local)
	mock := &basesdk.BaseClientMock{JSONBody: scheduleSocialPostResponse, Error: nil}
	now := time.Now()
	client.SDKClient = mock
	client.ScheduleTwitterPost(context.Background(), "accountGroupId", "posttext", "socialPostID", "SSID", now, "img.com")

	expectedParams := map[string]interface{}{
		"agid":         "accountGroupId",
		"postText":     "posttext",
		"socialPostId": "socialPostID",
		"ssid":         "SSID",
		"postDateTime": basesdk.ConvertTimeToVAPITimestamp(now),
		"imageUrl":     "img.com",
	}
	if !reflect.DeepEqual(mock.ParamsSent, expectedParams) {
		t.Errorf("Expected params: %v, but got %v", expectedParams, mock.ParamsSent)
	}
	if mock.PathRequested != scheduleTwitterPostPath {
		t.Errorf("Expected path: %s, but got %s", scheduleTwitterPostPath, mock.PathRequested)
	}
}

func Test_ScheduleGooglePlusPostMakesCorrectCallToCS(t *testing.T) {
	client := BuildSocialPostClient("user", "key", config.Local)
	mock := &basesdk.BaseClientMock{JSONBody: scheduleSocialPostResponse, Error: nil}
	now := time.Now()
	client.SDKClient = mock
	client.ScheduleGooglePlusPost(context.Background(), "accountGroupId", "posttext", "socialPostID", "SSID", now, "img.com")

	expectedParams := map[string]interface{}{
		"agid":         "accountGroupId",
		"postText":     "posttext",
		"socialPostId": "socialPostID",
		"ssid":         "SSID",
		"postDateTime": basesdk.ConvertTimeToVAPITimestamp(now),
		"imageUrl":     "img.com",
	}
	if !reflect.DeepEqual(mock.ParamsSent, expectedParams) {
		t.Errorf("Expected params: %v, but got %v", expectedParams, mock.ParamsSent)
	}
	if mock.PathRequested != scheduleGooglePlusPostPath {
		t.Errorf("Expected path: %s, but got %s", scheduleGooglePlusPostPath, mock.PathRequested)
	}
}

func Test_ScheduleLinkedInPostMakesCorrectCallToCS(t *testing.T) {
	client := BuildSocialPostClient("user", "key", config.Local)
	mock := &basesdk.BaseClientMock{JSONBody: scheduleSocialPostResponse, Error: nil}
	now := time.Now()
	client.SDKClient = mock
	client.ScheduleLinkedinPost(context.Background(), "accountGroupId", "posttext", "socialPostID", "SSID", now, "img.com")

	expectedParams := map[string]interface{}{
		"agid":         "accountGroupId",
		"postText":     "posttext",
		"socialPostId": "socialPostID",
		"ssid":         "SSID",
		"postDateTime": basesdk.ConvertTimeToVAPITimestamp(now),
		"imageUrl":     "img.com",
	}
	if !reflect.DeepEqual(mock.ParamsSent, expectedParams) {
		t.Errorf("Expected params: %v, but got %v", expectedParams, mock.ParamsSent)
	}
	if mock.PathRequested != scheduleLinkedinPostPath {
		t.Errorf("Expected path: %s, but got %s", scheduleLinkedinPostPath, mock.PathRequested)
	}
}

var listSocialPostCSResposne = `{
	"version": "2.0",
	"data": [{
		"services": [{
			"username": "Warden's Grand Casino",
			"postText": "Have you heard? Like our page between now October 8th for a chance to win $100 in game play!",
			"userId": "FBU-1214610698612373",
			"agid": "AG-G5XZWXD3",
			"postedDateTime": "2016-10-01T19:25:01Z",
			"likes": 0,
			"scheduledDateTime": "2016-10-01T19:25:00Z",
			"ssid": "FBP-D960CCA681F9402B984AC9F01B5AA51F",
			"socialPostId": "FacebookPost-ed01320972ef4ef4b04a79366a184633-FBP-D960CCA681F9402B984AC9F01B5AA51F",
			"deletionStatus": "NONE",
			"facebookTags": [],
			"attachment": null,
			"status": "COMPLETED",
			"multiLocationPostId": null,
			"tags": ["SCHEDULED", "ACTION_REQUIRED", "POSTED"],
			"clientTags": [],
			"profileImageUrl": "http://graph.facebook.com/1214610698612373/picture",
			"postedByOwner": true,
			"isError": null,
			"permalink": "http://www.facebook.com/permalink.php?story_fbid=1316168381789937&id=1214610698612373",
			"name": "The Grand Warden's Casino",
			"reshares": 0,
			"imageUrl": "http://lh3.googleusercontent.com/QmMUj5vxO3mrOEEwOFUuifiIhtWRNyeVUFUXORsFRq-5vuSmQRGO9ld6UU3iU5YHkFQpMjnarXoe8BG76fRoBJjy=s900",
			"profileUrl": "http://www.facebook.com/profile.php?id=1214610698612373",
			"place": null,
			"postCreatedDateTime": "2016-09-07T00:21:30Z"
		}],
		"socialPostId": "SocialPosts-ed01320972ef4ef4b04a79366a184633"
	}, {
		"services": [{
			"username": "Carmileta Lehman",
			"postText": "",
			"userId": "FBU-10207040226706770",
			"agid": "AG-G5XZWXD3",
			"postedDateTime": "2016-09-12T18:12:43Z",
			"likes": 0,
			"scheduledDateTime": null,
			"ssid": "FBP-D960CCA681F9402B984AC9F01B5AA51F",
			"socialPostId": "FacebookPost-5962cfdad77144a8ad59af6f404eb85a-FBP-D960CCA681F9402B984AC9F01B5AA51F",
			"deletionStatus": "NONE",
			"facebookTags": [],
			"attachment": {
				"caption": null,
				"fbObjectId": "1214610698612373_10210238868630819",
				"description": null,
				"name": null,
				"media": [{
					"src": "https://scontent.xx.fbcdn.net/v/t1.0-9/14316778_1282276271845815_5489864147119555648_n.jpg?oh=eecd93282da7b866f661811a2c512332&oe=587B8E68",
					"alt": null,
					"href": "https://www.facebook.com/1214610698612373/photos/a.1218066388266804.1073741828.1214610698612373/1282276271845815/?type=3",
					"type": "photo"
				}]
			},
			"status": "COMPLETED",
			"multiLocationPostId": null,
			"tags": ["ACTION_REQUIRED", "POSTED"],
			"clientTags": ["SM"],
			"profileImageUrl": "http://graph.facebook.com/10207040226706770/picture",
			"postedByOwner": false,
			"isError": null,
			"permalink": "http://www.facebook.com/permalink.php?story_fbid=10210238868630819&id=1214610698612373",
			"name": "Carmileta Lehman",
			"reshares": 0,
			"imageUrl": "https://scontent.xx.fbcdn.net/v/t1.0-9/14316778_1282276271845815_5489864147119555648_n.jpg?oh=eecd93282da7b866f661811a2c512332&oe=587B8E68",
			"profileUrl": "http://www.facebook.com/profile.php?id=10207040226706770",
			"place": null,
			"postCreatedDateTime": "2016-09-14T03:54:57Z"
		}],
		"socialPostId": "SocialPosts-5962cfdad77144a8ad59af6f404eb85a"
	}],
	"requestId": "5930982800ff0545fabf33f5b80001737e726570636f72652d70726f6400016170693a3534342d656e2d7475726e2d6f66662d64656275670001012b",
	"responseTime": 137,
	"statusCode": 200,
	"nextQueryString": "cursor=CrsBCiIKFW9yZGVyQnlQb3N0ZWREYXRlVGltZRIJCMDn_cTkqtQCEpABag5zfnJlcGNvcmUtdGVzdHJ-CxIXU29jaWFsUHJvZmlsZVBvc3RQYXJlbnQiJFNDUC01Nzk3QkRDQTMxOUY0MzY2QjY0QTQ2RkZCRkU1MTE3NQwLEgtTb2NpYWxQb3N0cyIsU29jaWFsUG9zdHMtN2ViOTA0ZjdkNDcwNDFmYWEyY2Y0MmI3Njg5YTNmYzkMGAAgAQ%3D%3D&agid=AG-TPV5TMG5"
}`

func Test_ListSocialPostsReturnsSocialPostsOnSuccess(t *testing.T) {
	client := BuildSocialPostClient("user", "key", config.Local)
	mock := &basesdk.BaseClientMock{JSONBody: listSocialPostCSResposne, Error: nil}
	endDateTime := time.Now()
	startDateTime := endDateTime.AddDate(0, 0, -1)
	client.SDKClient = mock
	res, cursor, err := client.ListSocialPosts(context.Background(), "accountGroupID", startDateTime, endDateTime, []string{"ssid1", "ssid2"}, "cursor", 10)

	if err != nil {
		t.Errorf("Expected no error, got %s", err.Error())
	}

	if cursor != "CrsBCiIKFW9yZGVyQnlQb3N0ZWREYXRlVGltZRIJCMDn_cTkqtQCEpABag5zfnJlcGNvcmUtdGVzdHJ-CxIXU29jaWFsUHJvZmlsZVBvc3RQYXJlbnQiJFNDUC01Nzk3QkRDQTMxOUY0MzY2QjY0QTQ2RkZCRkU1MTE3NQwLEgtTb2NpYWxQb3N0cyIsU29jaWFsUG9zdHMtN2ViOTA0ZjdkNDcwNDFmYWEyY2Y0MmI3Njg5YTNmYzkMGAAgAQ==" {
		t.Errorf("Expected CrsBCiIKFW9yZGVyQnlQb3N0ZWREYXRlVGltZRIJCMDn_cTkqtQCEpABag5zfnJlcGNvcmUtdGVzdHJ-CxIXU29jaWFsUHJvZmlsZVBvc3RQYXJlbnQiJFNDUC01Nzk3QkRDQTMxOUY0MzY2QjY0QTQ2RkZCRkU1MTE3NQwLEgtTb2NpYWxQb3N0cyIsU29jaWFsUG9zdHMtN2ViOTA0ZjdkNDcwNDFmYWEyY2Y0MmI3Njg5YTNmYzkMGAAgAQ== got %s", cursor)
	}

	if len(res) != 2 {
		t.Errorf("Expected 2 posts, got %d", len(res))
	}

	expected1 := &SocialPost{
		AccountGroupID:      "AG-G5XZWXD3",
		IsError:             false,
		DeletionStatus:      "NONE",
		Permalink:           "http://www.facebook.com/permalink.php?story_fbid=1316168381789937&id=1214610698612373",
		PostCreatedDateTime: "2016-09-07T00:21:30Z",
		PostText:            "Have you heard? Like our page between now October 8th for a chance to win $100 in game play!",
		PostedDateTime:      "2016-10-01T19:25:01Z",
		ProfileURL:          "http://www.facebook.com/profile.php?id=1214610698612373",
		ProfileImageURL:     "http://graph.facebook.com/1214610698612373/picture",
		ScheduledDateTime:   "2016-10-01T19:25:00Z",
		SocialPostID:        "FacebookPost-ed01320972ef4ef4b04a79366a184633-FBP-D960CCA681F9402B984AC9F01B5AA51F",
		Ssid:                "FBP-D960CCA681F9402B984AC9F01B5AA51F",
		Status:              "COMPLETED",
		ImageURL:            "http://lh3.googleusercontent.com/QmMUj5vxO3mrOEEwOFUuifiIhtWRNyeVUFUXORsFRq-5vuSmQRGO9ld6UU3iU5YHkFQpMjnarXoe8BG76fRoBJjy=s900",
		Name:                "The Grand Warden's Casino",
		Username:            "Warden's Grand Casino",
		ParentSocialPostID:  "SocialPosts-ed01320972ef4ef4b04a79366a184633",
	}
	if *res[0] != *expected1 {
		t.Errorf("Expected: %#v, but got %#v", res[0], expected1)
	}

	expected2 := &SocialPost{
		AccountGroupID:      "AG-G5XZWXD3",
		IsError:             false,
		DeletionStatus:      "NONE",
		Permalink:           "http://www.facebook.com/permalink.php?story_fbid=10210238868630819&id=1214610698612373",
		PostCreatedDateTime: "2016-09-14T03:54:57Z",
		PostText:            "",
		PostedDateTime:      "2016-09-12T18:12:43Z",
		ProfileURL:          "http://www.facebook.com/profile.php?id=10207040226706770",
		ProfileImageURL:     "http://graph.facebook.com/10207040226706770/picture",
		ScheduledDateTime:   "",
		SocialPostID:        "FacebookPost-5962cfdad77144a8ad59af6f404eb85a-FBP-D960CCA681F9402B984AC9F01B5AA51F",
		Ssid:                "FBP-D960CCA681F9402B984AC9F01B5AA51F",
		Status:              "COMPLETED",
		ImageURL:            "https://scontent.xx.fbcdn.net/v/t1.0-9/14316778_1282276271845815_5489864147119555648_n.jpg?oh=eecd93282da7b866f661811a2c512332&oe=587B8E68",
		Name:                "Carmileta Lehman",
		Username:            "Carmileta Lehman",
		ParentSocialPostID:  "SocialPosts-5962cfdad77144a8ad59af6f404eb85a",
	}
	if *res[1] != *expected2 {
		t.Errorf("Expected: %#v, but got %#v", res[1], expected2)
	}
}

var listSocialPostEmptyCursorResponse = `{
	"version": "2.0",
	"data": [],
	"requestId": "5930982800ff0545fabf33f5b80001737e726570636f72652d70726f6400016170693a3534342d656e2d7475726e2d6f66662d64656275670001012b",
	"responseTime": 137,
	"statusCode": 200,
	"nextQueryString": ""
}`

func Test_ListSocialPostsReturnsSocialPostsOnErrorParsingCursor(t *testing.T) {
	client := BuildSocialPostClient("user", "key", config.Local)
	mock := &basesdk.BaseClientMock{JSONBody: listSocialPostEmptyCursorResponse, Error: nil}
	endDateTime := time.Now()
	startDateTime := endDateTime.AddDate(0, 0, -1)
	client.SDKClient = mock
	_, cursor, err := client.ListSocialPosts(context.Background(), "accountGroupID", startDateTime, endDateTime, []string{"ssid1", "ssid2"}, "cursor", 10)

	if err != nil {
		t.Errorf("Expected no error, got %s", err.Error())
	}
	if cursor != "" {
		t.Errorf("Expected empty cursor, got %s", cursor)
	}
}

func Test_ListSocialPostMakesCorrectCallToCS(t *testing.T) {
	client := BuildSocialPostClient("user", "key", config.Local)
	mock := &basesdk.BaseClientMock{JSONBody: listSocialPostCSResposne, Error: nil}
	endDateTime := time.Now()
	startDateTime := endDateTime.AddDate(0, 0, -1)
	client.SDKClient = mock
	client.ListSocialPosts(context.Background(), "accountGroupID", startDateTime, endDateTime, []string{"ssid1", "ssid2"}, "cursor", 10)

	expectedParams := map[string]interface{}{
		"agid":          "accountGroupID",
		"startDateTime": basesdk.ConvertTimeToVAPITimestamp(startDateTime),
		"endDateTime":   basesdk.ConvertTimeToVAPITimestamp(endDateTime),
		"ssid":          []string{"ssid1", "ssid2"},
		"pageSize":      10,
		"cursor":        "cursor",
	}
	if !reflect.DeepEqual(mock.ParamsSent, expectedParams) {
		t.Errorf("Expected params: %v, but got %v", expectedParams, mock.ParamsSent)
	}
	if mock.PathRequested != listSocialPostPath {
		t.Errorf("Expected path: %s, but got %s", listSocialPostPath, mock.PathRequested)
	}
}

func Test_ListSocialPostsReturnsErrorWhenErrorConvertingResponse(t *testing.T) {
	client := BuildSocialPostClient("user", "key", config.Local)
	client.SDKClient = &basesdk.BaseClientMock{JSONBody: `{"data":"garbage"}`}
	endDateTime := time.Now()
	startDateTime := endDateTime.AddDate(0, 0, -1)
	_, _, err := client.ListSocialPosts(context.Background(), "accountGroupID", startDateTime, endDateTime, []string{"ssid1", "ssid2"}, "cursor", 10)

	if err == nil {
		t.Errorf("Expected error, got no error")
	}
}

func Test_ListSocialPostsReturnsErrorWhenCoreReturnsError(t *testing.T) {
	client := BuildSocialPostClient("user", "key", config.Local)
	expectedError := errors.New("New error")
	mock := &basesdk.BaseClientMock{JSONBody: listSocialPostCSResposne, Error: expectedError}
	endDateTime := time.Now()
	startDateTime := endDateTime.AddDate(0, 0, -1)
	client.SDKClient = mock
	_, _, err := client.ListSocialPosts(context.Background(), "accountGroupID", startDateTime, endDateTime, []string{"ssid1", "ssid2"}, "cursor", 10)

	if err != expectedError {
		t.Errorf("Expected error: %s, but got %s", expectedError.Error(), err.Error())
	}
}

func Test_ListSocialPostsReturnsErrorWhenNoAccountGroupIDProvided(t *testing.T) {
	client := BuildSocialPostClient("user", "key", config.Local)
	mock := &basesdk.BaseClientMock{JSONBody: listSocialPostCSResposne, Error: nil}
	endDateTime := time.Now()
	startDateTime := endDateTime.AddDate(0, 0, -1)
	client.SDKClient = mock
	_, _, err := client.ListSocialPosts(context.Background(), "", startDateTime, endDateTime, []string{"ssid1", "ssid2"}, "cursor", 10)

	if err == nil {
		t.Errorf("Expected error but got none")
	}
}

var listPartnerScheduledSocialPostCSResponse = `{
	"nextUrl": "http://repcore-test.appspot.com/internalApi/v2/socialPost/scheduled/lookup/?pageSize=2&endDateTime=2017-12-13T14%3A15%3A16Z&cursor=CrsBCiIKFW9yZGVyQnlQb3N0ZWREYXRlVGltZRIJCP7S6oTM4dUCEpABag5zfnJlcGNvcmUtdGVzdHJ-CxIXU29jaWFsUHJvZmlsZVBvc3RQYXJlbnQiJFNDUC01Nzk3QkRDQTMxOUY0MzY2QjY0QTQ2RkZCRkU1MTE3NQwLEgtTb2NpYWxQb3N0cyIsU29jaWFsUG9zdHMtNWQwYTUzOWNiMjgwNDhkOWIzYWEzOTZlY2U5ZTU1ZDMMGAAgAA%3D%3D&partnerId=ABC&startDateTime=2015-12-13T14%3A15%3A16Z",
	"log": [],
	"version": "2.0",
	"requestId": "599c9f9800ff039e5c4fe0157c0001737e726570636f72652d7465737400016170693a636f6e74696e756f75730001011f",
	"responseTime": 73,
	"nextQueryString": "pageSize=2&endDateTime=2017-12-13T14%3A15%3A16Z&cursor=CrsBCiIKFW9yZGVyQnlQb3N0ZWREYXRlVGltZRIJCP7S6oTM4dUCEpABag5zfnJlcGNvcmUtdGVzdHJ-CxIXU29jaWFsUHJvZmlsZVBvc3RQYXJlbnQiJFNDUC01Nzk3QkRDQTMxOUY0MzY2QjY0QTQ2RkZCRkU1MTE3NQwLEgtTb2NpYWxQb3N0cyIsU29jaWFsUG9zdHMtNWQwYTUzOWNiMjgwNDhkOWIzYWEzOTZlY2U5ZTU1ZDMMGAAgAA%3D%3D&partnerId=ABC&startDateTime=2015-12-13T14%3A15%3A16Z",
	"data": [
	  {
		"services": [
		  {
			"username": "All the pages",
			"postText": "Facebook only",
			"userId": "FBU-808721995867241",
			"agid": "AG-TPV5TMG5",
			"postedDateTime": "2017-08-18T20:12:13Z",
			"likes": 0,
			"scheduledDateTime": "2017-08-18T20:12:11Z",
			"ssid": "FBP-0887CD42C8B041F795395922D63654DC",
			"socialPostId": "FacebookPost-ba1244865a48429eb2bcaac42e17a24b-FBP-0887CD42C8B041F795395922D63654DC",
			"deletionStatus": "NONE",
			"facebookTags": [],
			"attachment": null,
			"status": "COMPLETED",
			"multiLocationPostId": null,
			"tags": [
			  "SCHEDULED",
			  "POSTED"
			],
			"clientTags": [],
			"profileImageUrl": "http://graph.facebook.com/808721995867241/picture",
			"postedByOwner": true,
			"isError": null,
			"permalink": "http://www.facebook.com/permalink.php?story_fbid=1677498532322912&id=808721995867241",
			"name": "All the pages",
			"reshares": 0,
			"imageUrl": null,
			"profileUrl": "http://www.facebook.com/profile.php?id=808721995867241",
			"place": null,
			"postCreatedDateTime": "2017-08-18T20:12:11Z"
		  }
		],
		"socialPostId": "SocialPosts-ba1244865a48429eb2bcaac42e17a24b"
	  },
	  {
		"services": [
		  {
			"username": "SrepGrys",
			"postText": "facebook and twitter",
			"userId": "2445757524",
			"agid": "AG-TPV5TMG5",
			"retweetCount": 0,
			"postedDateTime": "2017-08-18T20:12:51Z",
			"scheduledDateTime": "2017-08-18T20:12:47Z",
			"ssid": "TWU-2445757524",
			"favouriteCount": 0,
			"socialPostId": "TwitterPost-5d0a539cb28048d9b3aa396ece9e55d3-TWU-2445757524",
			"deletionStatus": "NONE",
			"displayCoordinatesFlag": true,
			"latitude": null,
			"replyToPostId": null,
			"inReplyToScreenName": null,
			"status": "COMPLETED",
			"multiLocationPostId": null,
			"retweetedFlag": false,
			"tags": [
			  "SCHEDULED",
			  "POSTED"
			],
			"clientTags": [],
			"profileImageUrl": "http://pbs.twimg.com/profile_images/540599330386616321/aQJWV-Pg_normal.png",
			"possiblySensitiveFlag": false,
			"postedByOwner": true,
			"isError": null,
			"permalink": "http://twitter.com/SrepGrys/status/898638845578682368",
			"name": "Shawn",
			"imageUrl": null,
			"placeId": null,
			"longitude": null,
			"profileUrl": "http://twitter.com/SrepGrys",
			"postCreatedDateTime": "2017-08-18T20:12:47Z"
		  }
		],
		"socialPostId": "SocialPosts-5d0a539cb28048d9b3aa396ece9e55d3"
	  }
	],
	"statusCode": 200
  }`

func Test_ListPartnerScheduledPostsReturnsSocialPostsOnSuccess(t *testing.T) {
	client := BuildSocialPostClient("user", "key", config.Local)
	mock := &basesdk.BaseClientMock{JSONBody: listPartnerScheduledSocialPostCSResponse, Error: nil}
	endDateTime := time.Now()
	startDateTime := endDateTime.AddDate(0, 0, -1)
	client.SDKClient = mock
	res, cursor, err := client.ListPartnerScheduledPosts(context.Background(), "ABC", startDateTime, endDateTime, "cursor", 10)

	if err != nil {
		t.Errorf("Expected no error, got %s", err.Error())
	}

	if cursor != "CrsBCiIKFW9yZGVyQnlQb3N0ZWREYXRlVGltZRIJCP7S6oTM4dUCEpABag5zfnJlcGNvcmUtdGVzdHJ-CxIXU29jaWFsUHJvZmlsZVBvc3RQYXJlbnQiJFNDUC01Nzk3QkRDQTMxOUY0MzY2QjY0QTQ2RkZCRkU1MTE3NQwLEgtTb2NpYWxQb3N0cyIsU29jaWFsUG9zdHMtNWQwYTUzOWNiMjgwNDhkOWIzYWEzOTZlY2U5ZTU1ZDMMGAAgAA==" {
		t.Errorf("Expected CrsBCiIKFW9yZGVyQnlQb3N0ZWREYXRlVGltZRIJCP7S6oTM4dUCEpABag5zfnJlcGNvcmUtdGVzdHJ-CxIXU29jaWFsUHJvZmlsZVBvc3RQYXJlbnQiJFNDUC01Nzk3QkRDQTMxOUY0MzY2QjY0QTQ2RkZCRkU1MTE3NQwLEgtTb2NpYWxQb3N0cyIsU29jaWFsUG9zdHMtNWQwYTUzOWNiMjgwNDhkOWIzYWEzOTZlY2U5ZTU1ZDMMGAAgAA== got %s", cursor)
	}

	if len(res) != 2 {
		t.Errorf("Expected 2 posts, got %d", len(res))
	}

	expected1 := &SocialPost{
		AccountGroupID:      "AG-TPV5TMG5",
		IsError:             false,
		DeletionStatus:      "NONE",
		Permalink:           "http://www.facebook.com/permalink.php?story_fbid=1677498532322912&id=808721995867241",
		PostCreatedDateTime: "2017-08-18T20:12:11Z",
		PostText:            "Facebook only",
		PostedDateTime:      "2017-08-18T20:12:13Z",
		ProfileURL:          "http://www.facebook.com/profile.php?id=808721995867241",
		ProfileImageURL:     "http://graph.facebook.com/808721995867241/picture",
		ScheduledDateTime:   "2017-08-18T20:12:11Z",
		SocialPostID:        "FacebookPost-ba1244865a48429eb2bcaac42e17a24b-FBP-0887CD42C8B041F795395922D63654DC",
		Ssid:                "FBP-0887CD42C8B041F795395922D63654DC",
		Status:              "COMPLETED",
		ImageURL:            "",
		Name:                "All the pages",
		Username:            "All the pages",
		ParentSocialPostID:  "SocialPosts-ba1244865a48429eb2bcaac42e17a24b",
	}
	assert.Equal(t, expected1, res[0])

	expected2 := &SocialPost{
		AccountGroupID:      "AG-TPV5TMG5",
		IsError:             false,
		DeletionStatus:      "NONE",
		Permalink:           "http://twitter.com/SrepGrys/status/898638845578682368",
		PostCreatedDateTime: "2017-08-18T20:12:47Z",
		PostText:            "facebook and twitter",
		PostedDateTime:      "2017-08-18T20:12:51Z",
		ProfileURL:          "http://twitter.com/SrepGrys",
		ProfileImageURL:     "http://pbs.twimg.com/profile_images/540599330386616321/aQJWV-Pg_normal.png",
		ScheduledDateTime:   "2017-08-18T20:12:47Z",
		SocialPostID:        "TwitterPost-5d0a539cb28048d9b3aa396ece9e55d3-TWU-2445757524",
		Ssid:                "TWU-2445757524",
		Status:              "COMPLETED",
		ImageURL:            "",
		Name:                "Shawn",
		Username:            "SrepGrys",
		ParentSocialPostID:  "SocialPosts-5d0a539cb28048d9b3aa396ece9e55d3",
	}
	assert.Equal(t, expected2, res[1])
}

var listPartnerScheduledSocialPostsEmptyCursorResponse = `{
	"log": [],
	"version": "2.0",
	"requestId": "599ca06b00ff03741d711652b30001737e726570636f72652d7465737400016170693a636f6e74696e756f757300010122",
	"responseTime": 72,
	"data": [],
	"statusCode": 200
  }`

func Test_ListPartnerScheduledSocialPostsReturnsErrorOnErrorParsingCursor(t *testing.T) {
	client := BuildSocialPostClient("user", "key", config.Local)
	mock := &basesdk.BaseClientMock{JSONBody: listSocialPostEmptyCursorResponse, Error: nil}
	endDateTime := time.Now()
	startDateTime := endDateTime.AddDate(0, 0, -1)
	client.SDKClient = mock
	_, cursor, err := client.ListPartnerScheduledPosts(context.Background(), "ABC", startDateTime, endDateTime, "cursor", 10)

	if err != nil {
		t.Errorf("Expected no error, got %s", err.Error())
	}
	if cursor != "" {
		t.Errorf("Expected empty cursor, got %s", cursor)
	}
}

func Test_ListPartnerScheduledSocialPostsMakesAGetCallWithTheCorrectPath(t *testing.T) {
	client := BuildSocialPostClient("user", "key", config.Local)
	mock := &basesdk.BaseClientMock{JSONBody: listSocialPostEmptyCursorResponse, Error: nil}
	endDateTime := time.Now()
	startDateTime := endDateTime.AddDate(0, 0, -1)
	client.SDKClient = mock
	client.ListPartnerScheduledPosts(context.Background(), "ABC", startDateTime, endDateTime, "cursor", 10)
	assert.Equal(t, "/internalApi/v2/socialPost/scheduled/lookup/", mock.PathRequested)
}

func Test_ListPartnerScheduledPostsByCreatedDateTimeMakesAGetCallWithTheCorrectPath(t *testing.T) {
	client := BuildSocialPostClient("user", "key", config.Local)
	mock := &basesdk.BaseClientMock{JSONBody: listSocialPostEmptyCursorResponse, Error: nil}
	endDateTime := time.Now()
	startDateTime := endDateTime.AddDate(0, 0, -1)
	client.SDKClient = mock
	client.ListPartnerScheduledPostsByCreatedDateTime(context.Background(), "ABC", startDateTime, endDateTime, "cursor", 10)
	assert.Equal(t, "/internalApi/v2/socialPost/scheduled/lookup/byCreatedDate/", mock.PathRequested)
}
