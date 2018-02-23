package basesdk

import (
	"errors"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"context"

	"github.com/stretchr/testify/assert"
)

type MockHttpClientBadDo struct{}

func (c *MockHttpClientBadDo) Do(r *http.Request) (*http.Response, error) {
	return nil, errors.New("Something Failed In Do")
}

func TestPerformRequestReturnsErrorIfDoFails(t *testing.T) {
	mockHttpClient := &MockHttpClientBadDo{}
	auth := UserKey{APIUser: "RM", APIKey: "12345"}
	c := BaseClient{Authorization: auth, RootURL: "https://partner-central-test.appspot.com"}
	params := map[string]interface{}{"solutionId": "SO-123"}
	_, err := c.performRequest(context.Background(), mockHttpClient, "GET", "/marketplaceInternalApi/v1/solution/get", params)
	assert.EqualError(t, err, "Something Failed In Do")
}

type MockHTTPClientBadStatus struct{}

func (c *MockHTTPClientBadStatus) Do(r *http.Request) (*http.Response, error) {
	response := &http.Response{StatusCode: 400}
	return response, nil
}

func TestPerformRequestReturnsAnErrorIfResponseIsOver299(t *testing.T) {
	mockHTTPClient := &MockHTTPClientBadStatus{}
	auth := UserKey{APIUser: "RM", APIKey: "12345"}
	c := BaseClient{Authorization: auth, RootURL: "https://partner-central-test.appspot.com"}
	params := map[string]interface{}{"solutionId": "SO-123"}
	_, err := c.performRequest(context.Background(), mockHTTPClient, "GET", "/marketplaceInternalApi/v1/solution/get", params)
	assert.EqualError(t, err, "Bad Request: ")
}

func TestPerformRequestReturnsAnErrorIfMethodNotGetOrPostOrHead(t *testing.T) {
	auth := UserKey{APIUser: "RM", APIKey: "12345"}
	c := BaseClient{Authorization: auth}
	_, err := c.performRequest(context.Background(), nil, "INVALID_METHOD", "/test/path", nil)
	expected := "method INVALID_METHOD not valid, must be one of [GET POST HEAD]"
	assert.EqualError(t, err, expected)
}

// ******************
// buildRequest TESTS
// ******************
func TestBuildRequestReturnsRequestWithProperMethod(t *testing.T) {
	c := BaseClient{RootURL: "https://partner-central-test.appspot.com"}
	params := map[string]interface{}{}

	req, _ := c.buildRequest(context.Background(), "GET", "/test/path", params)
	assert.Equal(t, req.Method, "GET")

	req, _ = c.buildRequest(context.Background(), "POST", "/test/path", params)
	assert.Equal(t, req.Method, "POST")

	req, _ = c.buildRequest(context.Background(), "HEAD", "/test/path", nil)
	assert.Equal(t, req.Method, "HEAD")
}

func TestBuildRequestReturnsRequestWithProperURL(t *testing.T) {
	c := BaseClient{RootURL: "https://partner-central-test.appspot.com"}
	params := map[string]interface{}{}

	req, _ := c.buildRequest(context.Background(), "GET", "/test/path", params)
	expected := "https://partner-central-test.appspot.com/test/path"
	assert.Equal(t, req.URL.String(), expected)
}

func TestBuildRequestReturnsRequestWithParamsAsQueryStringIfMethodGET(t *testing.T) {
	c := BaseClient{RootURL: "https://partner-central-test.appspot.com"}
	params := map[string]interface{}{
		"param1": "value1",
		"param2": "value2",
	}

	req, _ := c.buildRequest(context.Background(), "GET", "/test/path", params)
	expected := "param1=value1&param2=value2"
	assert.Equal(t, req.URL.RawQuery, expected)
	assert.Nil(t, req.Body)
}

func TestBuildRequestReturnsRequestWithEmptyContentTypeHeaderIfMethodGET(t *testing.T) {
	c := BaseClient{RootURL: "https://partner-central-test.appspot.com"}
	params := map[string]interface{}{
		"param1": "value1",
		"param2": "value2",
	}

	req, _ := c.buildRequest(context.Background(), "GET", "/test/path", params)
	assert.Equal(t, "", req.Header.Get("Content-Type"))
}

func TestBuildRequestReturnsRequestWithEmptyContentTypeHeaderIfMethodHEAD(t *testing.T) {
	c := BaseClient{RootURL: "https://partner-central-test.appspot.com"}

	req, _ := c.buildRequest(context.Background(), "HEAD", "/test/path", nil)
	assert.Equal(t, "", req.Header.Get("Content-Type"))
}

func TestBuildRequestReturnsRequestWithParamsAsBodyIfMethodPOST(t *testing.T) {
	c := BaseClient{RootURL: "https://partner-central-test.appspot.com"}
	params := map[string]interface{}{
		"param1": "value1",
		"param2": "value2",
	}

	req, _ := c.buildRequest(context.Background(), "POST", "/test/path", params)
	expected := "param1=value1&param2=value2"
	assert.Empty(t, req.URL.RawQuery, expected)
	defer req.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(req.Body)
	body := string(bodyBytes)
	assert.Equal(t, body, "{\"param1\":\"value1\",\"param2\":\"value2\"}")
}

func TestBuildRequestReturnsRequestWithApplicationJsonContentHeaderIfMethodPOST(t *testing.T) {
	c := BaseClient{RootURL: "https://partner-central-test.appspot.com"}
	params := map[string]interface{}{
		"param1": "value1",
		"param2": "value2",
	}

	req, _ := c.buildRequest(context.Background(), "POST", "/test/path", params)
	assert.Equal(t, "application/json", req.Header.Get("Content-Type"))
}

// **********************
// buildQueryString TESTS
// **********************
func TestBuildQueryStringAddsApiUserApiKeyAndParamsWhenParamsPresent(t *testing.T) {
	c := BaseClient{RootURL: "https://partner-central-test.appspot.com"}
	params := map[string]interface{}{
		"param1": "value1",
	}

	req, _ := http.NewRequest("GET", "http://a.com/b/c", nil)
	req = c.buildQueryString(req, params)
	assert.Equal(t, req.URL.RawQuery, "param1=value1")
}

func TestBuildQueryStringAddsParamsWhenApiUserApiKeyNotPresent(t *testing.T) {
	c := BaseClient{RootURL: "https://partner-central-test.appspot.com"}
	params := map[string]interface{}{
		"param1": "value1",
	}

	req, _ := http.NewRequest("GET", "http://a.com/b/c", nil)
	req = c.buildQueryString(req, params)
	assert.Equal(t, req.URL.RawQuery, "param1=value1")
}

func TestBuildQueryStringAddsNothingWhenApiUserApiKeyAndParamsNotPresent(t *testing.T) {
	c := BaseClient{RootURL: "https://partner-central-test.appspot.com"}
	params := map[string]interface{}{}

	req, _ := http.NewRequest("GET", "http://a.com/b/c", nil)
	req = c.buildQueryString(req, params)
	assert.Empty(t, req.URL.RawQuery)
}

func Test_convertTimeToVAPITimestampReturnsEmptyStringWhenZeroTime(t *testing.T) {
	var none time.Time
	res := ConvertTimeToVAPITimestamp(none)
	if res != "" {
		t.Errorf("Expected zero time to be converted to empty string, got %s", res)
	}
}

func Test_convertTimeToVAPITimestampReturnsDateInCorrectFormatAsString(t *testing.T) {
	timestamp := time.Date(2017, 2, 2, 0, 0, 0, 0, time.UTC)
	res := ConvertTimeToVAPITimestamp(timestamp)
	if res != "2017-02-02T00:00:00Z" {
		t.Errorf("Expected [2017-02-02T00:00:00Z], got %s", res)
	}
}
