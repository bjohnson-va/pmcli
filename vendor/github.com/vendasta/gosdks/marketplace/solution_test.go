package marketplace

import (
	"github.com/stretchr/testify/suite"
	"github.com/vendasta/gosdks/marketplace/testing_utils"
	"testing"
	"time"
)

type solutionFromResponseBytesTestSuite struct {
	suite.Suite
	solution *Solution
	err      error
}

func (suite *solutionFromResponseBytesTestSuite) SetupTest() {
	resp := testingutils.ResponseFromString(`
	{
		"version": "1.0",
		"data": {
			"status": "published",
			"updated": "2017-05-07T22:14:38Z",
			"solution_id": "SOL-2906285AAB0C4589ABEE49F7F0A862CB",
			"name": "Awesome Possum",
			"billing_frequency": "daily",
			"market_id": "market-1",
			"_type": "SolutionModel",
			"created": "2017-02-07T22:14:38Z",
			"archived": "2017-05-15T22:14:38Z",
			"partner_id": "RLW",
			"content": null,
			"currency": "USD",
			"version": 2,
			"products": [
				"MP-5fa7a24ec7ac4f088d7e5c2a6dff4d6f"
			],
			"billing_frequency_other": null,
			"selling_price": 123,
			"pricing": {
				"_object_version": "1.0.2",
				"currency": "USD",
				"prices": [
					{
						"_object_version": "1.0.2",
						"frequency": "daily",
						"price": 12300
					}
				]
			},
			"hide_product_details": false,
			"icon": "//storage.googleapis.com/mediarepo-test/marketplace-products/89cf4a74-eb04-4664-934b-b47239ec617c/possum_icon.jpg",
			"hide_product_icons_and_names": true,
			"_object_version": "3.0.0",
			"updated_by": null
		},
		"requestId": "58e3aa5c00ff0958e70b259e230001737e706172746e65722d63656e7472616c2d7465737400016d61726b6574706c6163653a636f6e74696e756f757300010105",
		"responseTime": 49,
		"statusCode": 200
	}`)
	s, err := SolutionFromResponse(resp)
	suite.solution = s
	suite.err = err
}

func TestSolutionFromResponseBytesTestSuite(t *testing.T) {
	suite.Run(t, new(solutionFromResponseBytesTestSuite))
}

func (suite *solutionFromResponseBytesTestSuite) TestSolutionFromResponseBytesDidNotReturnError() {
	suite.Nil(suite.err)
}

func (suite *solutionFromResponseBytesTestSuite) TestSolutionFromResponseBytesReturnsSolutionWithCorrectSolutionId() {
	suite.Equal("SOL-2906285AAB0C4589ABEE49F7F0A862CB", suite.solution.SolutionID)
}

func (suite *solutionFromResponseBytesTestSuite) TestSolutionFromResponseBytesReturnsSolutionWithCorrectName() {
	suite.Equal("Awesome Possum", suite.solution.Name)
}

func (suite *solutionFromResponseBytesTestSuite) TestSolutionFromResponseBytesReturnsSolutionWithCorrectStatus() {
	suite.Equal("published", suite.solution.Status)
}

func (suite *solutionFromResponseBytesTestSuite) TestSolutionFromResponseBytesReturnsSolutionWithCorrectBillingFrequency() {
	suite.Equal(Daily, suite.solution.BillingFrequency)
}

func (suite *solutionFromResponseBytesTestSuite) TestSolutionFromResponseBytesReturnsSolutionWithCorrectBillingPartnerId() {
	suite.Equal("RLW", suite.solution.PartnerID)
}

func (suite *solutionFromResponseBytesTestSuite) TestSolutionFromResponseBytesReturnsSolutionWithCorrectBillingMarketId() {
	suite.Equal("market-1", suite.solution.MarketID)
}

func (suite *solutionFromResponseBytesTestSuite) TestSolutionFromResponseBytesReturnsSolutionWithCorrectBillingCurrency() {
	suite.Equal(USD, suite.solution.Currency)
}

func (suite *solutionFromResponseBytesTestSuite) TestSolutionFromResponseBytesReturnsSolutionWithCorrectProducts() {
	suite.Equal(1, len(suite.solution.Products))
	suite.Equal("MP-5fa7a24ec7ac4f088d7e5c2a6dff4d6f", suite.solution.Products[0])
}

func (suite *solutionFromResponseBytesTestSuite) TestSolutionFromResponseBytesReturnsSolutionWitCorrectContent() {
	suite.Empty(suite.solution.Content)
}

func (suite *solutionFromResponseBytesTestSuite) TestSolutionFromResponseBytesReturnsSolutionWithCorrectBillingFrequencyOther() {
	suite.Empty(suite.solution.BillingFrequencyOther)
}

func (suite *solutionFromResponseBytesTestSuite) TestSolutionFromResponseBytesReturnsSolutionWitCorrectSellingPrice() {
	suite.Equal(123.0, suite.solution.SellingPrice)
}

func (suite *solutionFromResponseBytesTestSuite) TestSolutionFromResponseBytesReturnsSolutionWithCorrectPrice() {
	suite.Equal(int64(12300), suite.solution.Pricing.Prices[0].Price)
}

func (suite *solutionFromResponseBytesTestSuite) TestSolutionFromResponseBytesReturnsSolutionWithCorrectFrequency() {
	suite.Equal(Daily, suite.solution.Pricing.Prices[0].BillingFrequency)
}

func (suite *solutionFromResponseBytesTestSuite) TestSolutionFromResponseBytesReturnsSolutionWithCorrectCurrency() {
	suite.Equal(USD, suite.solution.Pricing.Currency)
}

func (suite *solutionFromResponseBytesTestSuite) TestSolutionFromResponseBytesReturnsSolutionWithCorrectHideProductDetails() {
	suite.False(suite.solution.HideProductDetails)
}

func (suite *solutionFromResponseBytesTestSuite) TestSolutionFromResponseBytesReturnsSolutionWithCorrectHideProductIconsAndNames() {
	suite.True(suite.solution.HideProductIconsAndNames)
}

func (suite *solutionFromResponseBytesTestSuite) TestSolutionFromResponseBytesReturnsSolutionWithCorrectIconUrl() {
	suite.Equal("//storage.googleapis.com/mediarepo-test/marketplace-products/89cf4a74-eb04-4664-934b-b47239ec617c/possum_icon.jpg", suite.solution.IconURL)
}

func (suite *solutionFromResponseBytesTestSuite) TestSolutionFromResponseBytesReturnsSolutionWithCorrectCreatedDate() {
	created := time.Date(2017, 2, 7, 22, 14, 38, 0, time.UTC)
	suite.Equal(created, suite.solution.Created)
}

func (suite *solutionFromResponseBytesTestSuite) TestSolutionFromResponseBytesReturnsSolutionWithCorrectUpdatedDate() {
	updated := time.Date(2017, 5, 7, 22, 14, 38, 0, time.UTC)
	suite.Equal(updated, suite.solution.Updated)
}

func (suite *solutionFromResponseBytesTestSuite) TestSolutionFromResponseBytesReturnsSolutionWithCorrectArchivedDate() {
	archived := time.Date(2017, 5, 15, 22, 14, 38, 0, time.UTC)
	suite.Equal(archived, suite.solution.Archived)
}

type solutionFromResponseBytesMissingDataTestSuite struct {
	suite.Suite
	solution *Solution
	err      error
}

func (suite *solutionFromResponseBytesMissingDataTestSuite) SetupTest() {
	resp := testingutils.ResponseFromString(`{
		"version": "1.0",
		"data": {
			"status": "published",
			"updated": "2017-05-07T22:14:38Z",
			"solution_id": "SOL-2906285AAB0C4589ABEE49F7F0A862CB",
			"name": "Awesome Possum",
			"billing_frequency": "daily",
			"market_id": "market-1",
			"_type": "SolutionModel",
			"created": "2017-02-07T22:14:38Z",
			"archived": null,
			"partner_id": "RLW",
			"content": null,
			"currency": "USD",
			"version": 2,
			"products": [
				"MP-5fa7a24ec7ac4f088d7e5c2a6dff4d6f"
			],
			"billing_frequency_other": null,
			"selling_price": 123.0,
			"hide_product_details": false,
			"icon": "//storage.googleapis.com/mediarepo-test/marketplace-products/89cf4a74-eb04-4664-934b-b47239ec617c/possum_icon.jpg",
			"hide_product_icons_and_names": true,
			"_object_version": "3.0.0",
			"updated_by": null
		},
		"requestId": "58e3aa5c00ff0958e70b259e230001737e706172746e65722d63656e7472616c2d7465737400016d61726b6574706c6163653a636f6e74696e756f757300010105",
		"responseTime": 49,
		"statusCode": 200
	}`)
	s, err := SolutionFromResponse(resp)
	suite.solution = s
	suite.err = err
}

func TestSolutionFromResponseBytesMissingDataTestSuite(t *testing.T) {
	suite.Run(t, new(solutionFromResponseBytesMissingDataTestSuite))
}

func (suite *solutionFromResponseBytesMissingDataTestSuite) TestSolutionFromResponseBytesSetsArchivedToZeroTime() {
	suite.Nil(suite.err)
}

type solutionFromResponseBytesBadJSONTestSuite struct {
	suite.Suite
	solution *Solution
	err      error
}

func (suite *solutionFromResponseBytesBadJSONTestSuite) SetupTest() {
	resp := testingutils.ResponseFromString(`{badjson}`)
	s, err := SolutionFromResponse(resp)
	suite.solution = s
	suite.err = err
}

func TestSolutionFromResponseBytesBadJSONTestSuite(t *testing.T) {
	suite.Run(t, new(solutionFromResponseBytesBadJSONTestSuite))
}

func (suite *solutionFromResponseBytesBadJSONTestSuite) TestSolutionFromResponseBytesReturnsErrorIfJsonIsInvalid() {
	suite.EqualError(suite.err, "Failed to convert response to Solution: invalid character 'b' looking for beginning of object key string")
}
