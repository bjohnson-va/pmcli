package marketplace

import (
	"github.com/stretchr/testify/suite"
	"github.com/vendasta/gosdks/marketplace/testing_utils"
	"golang.org/x/net/context"
	"net/http"
	"testing"
)

func TestAccountClient_GetSessionTransferURL(t *testing.T) {
	suite.Run(t, new(GetSessionTransferURLTests))
}

func TestAccountClient_buildSessionTransferAPIURL(t *testing.T) {
	suite.Run(t, new(BuildSessionTransferAPIURLTests))
}

type GetSessionTransferURLTests struct {
	BaseAccountClientTest
}

func (suite GetSessionTransferURLTests) TestReturnsErrorIfEmptyAccountIDProvided() {
	_, err := suite.client.GetSessionTransferURL(context.Background(), "")
	suite.EqualError(err, "accountID is required")
}

type BuildSessionTransferAPIURLTests struct {
	BaseAccountClientTest
}

func (suite BuildSessionTransferAPIURLTests) TestReturnsCorrectlyFormattedURL() {
	path := suite.client.buildSessionTransferAPIURL("MyAccount")
	expectedPath := "/api/v1/account/MyAccount/session-transfer"
	suite.Equal(expectedPath, path)
}

func TestSessionTransferURLFromHttpResponse_ProcessResponse(t *testing.T) {
	suite.Run(t, new(ProcessTests))

}

type ProcessTests struct {
	suite.Suite
	handler  SessionTransferURLFromHTTPResponse
	response *http.Response
}

func (suite *ProcessTests) SetupTest() {
	suite.handler = SessionTransferURLFromHTTPResponse{}
}

func (suite *ProcessTests) TestReturnsTheCorrectURLFromValidResponse() {
	response := testingutils.ResponseFromString(`{
		"took": 12345,
		"data": {
			"session_transfer_url": "http://example.com"
		}
	}`)
	url, _ := suite.handler.ProcessResponse(response)
	expectedURL := "http://example.com"
	suite.Equal(expectedURL, url)
}

func (suite *ProcessTests) TestReturnsErrorFromBadResponse() {
	response := testingutils.ResponseFromString(`{bad-data}`)
	_, err := suite.handler.ProcessResponse(response)
	suite.EqualError(err, "failed to convert response to URL: invalid character 'b' looking for beginning of object key string")
}
