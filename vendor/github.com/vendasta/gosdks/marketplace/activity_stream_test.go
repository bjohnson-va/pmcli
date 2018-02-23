package marketplace

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/vendasta/gosdks/basesdk"
)

type activityStreamTestSuite struct {
	suite.Suite
}

func TestActivityStreamTestSuite(t *testing.T) {
	suite.Run(t, &activityStreamTestSuite{})
}

func (s *activityStreamTestSuite) TestValidatesActivityParams() {
	type testCase struct {
		activity  Activity
		message   string
		expectErr bool
	}
	testCases := []testCase{
		{
			activity: Activity{
				AccountID:    "AG-1234",
				Content:      "content",
				Link:         "",
				MediaURL:     "",
				SettingsTags: []string{"tags"},
				Title:        "A Title",
				Type:         "A Type",
			},
			message:   "valid activity should not return an error",
			expectErr: false,
		},
		{
			activity: Activity{
				AccountID:    "",
				Content:      "content",
				Link:         "http://example.com",
				MediaURL:     "http://google.ca",
				SettingsTags: []string{"tags"},
				Title:        "A Title",
				Type:         "A Type",
			},
			message:   "activity with missing AccountID should return an error",
			expectErr: true,
		},
		{
			activity: Activity{
				AccountID:    "AG-1234",
				Content:      "content",
				Link:         "http://example.com",
				MediaURL:     "http://google.ca",
				SettingsTags: []string{"tags"},
				Title:        "",
				Type:         "A Type",
			},
			message:   "activity with missing Title should return an error",
			expectErr: true,
		},
		{
			activity: Activity{
				AccountID:    "AG-1234",
				Content:      "content",
				Link:         "http://example.com",
				MediaURL:     "http://google.ca",
				SettingsTags: []string{"tags"},
				Title:        "A Title",
				Type:         "",
			},
			message:   "activity with missing Type should return an error",
			expectErr: true,
		},
		{
			activity: Activity{
				AccountID:    "AG-1234",
				Content:      "content",
				Link:         "http://example.com",
				MediaURL:     "not a url",
				SettingsTags: []string{"tags"},
				Title:        "A Title",
				Type:         "A Type",
			},
			message:   "activity with invalid MediaURL should return an error",
			expectErr: true,
		},
		{
			activity: Activity{
				AccountID:    "AG-1234",
				Content:      "content",
				Link:         "not a url",
				MediaURL:     "http://google.ca",
				SettingsTags: []string{"tags"},
				Title:        "A Title",
				Type:         "A Type",
			},
			message:   "activity with invalid Link should return an error",
			expectErr: true,
		},
	}

	client := activityClient{&basesdk.BaseClientMock{JSONBody: "{}"}}
	for _, tc := range testCases {
		_, err := client.CreateActivity(context.Background(), tc.activity)
		if tc.expectErr {
			s.Error(err, tc.message)
		} else {
			s.Nil(err, tc.message)
		}
	}
}

func (s *activityStreamTestSuite) TestCreateActivityUnpacksJSONResponse() {
	response := `{
  "took": 500,
  "data": {
    "activity_id": "ACT-1234",
    "account_id": "AG-12345",
    "created": "2018-01-02T18:31:16Z",
    "settings_tags": [
      "tag"
    ],
    "title": "My Title",
    "app_id": "APP-1234",
    "content": "My Content",
    "link": "http://google.ca",
    "media_url": "http://example.com",
	"activity_type": "My Activity Type"
  }
}`

	activity := Activity{
		AccountID:    "AG-12345",
		Content:      "My Content",
		Link:         "http://google.ca",
		MediaURL:     "http://example.com",
		SettingsTags: []string{"tag"},
		Title:        "My Title",
		Type:         "My Activity Type",
	}
	expected := &CreateActivityResponse{
		ActivityID: "ACT-1234",
		AppID:      "APP-1234",
		Created:    "2018-01-02T18:31:16Z",

		Activity: activity,
	}
	client := activityClient{&basesdk.BaseClientMock{JSONBody: response}}
	resp, err := client.CreateActivity(context.Background(), activity)
	s.Nil(err)
	s.Equal(expected, resp)
}
