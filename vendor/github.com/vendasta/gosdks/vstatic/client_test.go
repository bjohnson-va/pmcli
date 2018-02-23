package vstatic

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"golang.org/x/net/context"

	"github.com/stretchr/testify/assert"
	"github.com/vendasta/gosdks/pb/vstatic/v1"
)

func Test_serveIndex(t *testing.T) {

	type testCase struct {
		name string

		tokenFunc          GetIAMToken
		writer             *httptest.ResponseRecorder
		request            *http.Request
		getterMock         *grpcGetterMock
		expectedHTML       string
		expectedStatusCode int
	}
	cases := []*testCase{
		{
			name:       "write the html with success when successful",
			writer:     httptest.NewRecorder(),
			request:    httptest.NewRequest("get", "/", nil),
			tokenFunc:  func(r *http.Request) (string, error) { return "", nil },
			getterMock: &grpcGetterMock{html: []byte("the html"), deploymentID: "123"},

			expectedHTML:       "the html",
			expectedStatusCode: 200,
		},
		{
			name:      "returns 500 when getting an iam token fails",
			writer:    httptest.NewRecorder(),
			request:   httptest.NewRequest("get", "/", nil),
			tokenFunc: func(r *http.Request) (string, error) { return "", errors.New("get token failed") },

			expectedHTML:       "get token failed\n",
			expectedStatusCode: 500,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ctx := context.Background()
			client := &vstaticClient{
				client: c.getterMock,
				env:    vstatic_v1.Environment_Test,
				appID:  "test-app",
			}
			client.serveIndex(ctx, c.writer, c.request, c.tokenFunc)

			assert.Equal(t, c.expectedStatusCode, c.writer.Code, "error validating code")
			body, _ := ioutil.ReadAll(c.writer.Body)
			assert.Equal(t, c.expectedHTML, string(body), "error validating body")
		})
	}
}

type grpcGetterMock struct {
	err          error
	html         HTML
	deploymentID string
	servingURL   string
}

func (m *grpcGetterMock) getIndex(ctx context.Context, appID string, env vstatic_v1.Environment, iamToken string, deploymentID string) (HTML, string, error) {
	return m.html, m.deploymentID, m.err
}
func (m *grpcGetterMock) getAssetURL(ctx context.Context, appID string, env vstatic_v1.Environment, assetPath string, deploymentID string) (string, error) {
	return m.servingURL, m.err
}

func Test_getCurrentDeployment(t *testing.T) {

	type testCase struct {
		name                 string
		request              *http.Request
		expectedDeploymnetID string
	}
	cases := []*testCase{
		{
			name:                 "returns the correct deployment id when found in cookie",
			request:              &http.Request{Header: http.Header{"Cookie": []string{"vstatic-deploymentId=123;"}}},
			expectedDeploymnetID: "123",
		},
		{
			name:                 "empty string when cookie not found",
			request:              &http.Request{},
			expectedDeploymnetID: "",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ctx := context.Background()
			id := getCurrentDeployment(ctx, c.request)

			assert.Equal(t, c.expectedDeploymnetID, id)
		})
	}
}
