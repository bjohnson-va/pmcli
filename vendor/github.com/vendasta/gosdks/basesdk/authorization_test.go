package basesdk

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func Test_UserKey_SignRequest_http_request_has_key_and_user_added_to_query_string(t *testing.T) {
	testURL, _ := url.Parse("https://dundermifflin.com/api/paper?amount=25")
	request := &http.Request{
		Method: "POST",
		URL:    testURL,
	}
	authorization := UserKey{APIUser: "testuser", APIKey: "key-123"}
	authorization.SignRequest(request)
	assert.Equal(t, testURL.String(), "https://dundermifflin.com/api/paper?amount=25&apiKey=key-123&apiUser=testuser")
}

func Test_UserKey_SignRequest_will_ignore_key_if_not_provided(t *testing.T) {
	testURL, _ := url.Parse("https://dundermifflin.com/api/paper?amount=25")
	request := &http.Request{
		Method: "POST",
		URL:    testURL,
	}
	authorization := UserKey{APIUser: "testuser"}
	authorization.SignRequest(request)
	assert.Equal(t, testURL.String(), "https://dundermifflin.com/api/paper?amount=25&apiUser=testuser")
}

func Test_UserKey_SignRequest_will_ignore_user_if_not_provided(t *testing.T) {
	testURL, _ := url.Parse("https://dundermifflin.com/api/paper?amount=25")
	request := &http.Request{
		Method: "POST",
		URL:    testURL,
	}
	authorization := UserKey{APIKey: "key-123"}
	authorization.SignRequest(request)
	assert.Equal(t, testURL.String(), "https://dundermifflin.com/api/paper?amount=25&apiKey=key-123")
}
