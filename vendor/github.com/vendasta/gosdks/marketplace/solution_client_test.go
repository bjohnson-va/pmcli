package marketplace

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vendasta/gosdks/basesdk"
	"golang.org/x/net/context"
)

func TestGetSolutionReturnsAnErrorIfNoSolutionIdIsProvided(t *testing.T) {
	baseClient := &basesdk.BaseClientMock{}
	solutionClient := SolutionClient{SDKClient: baseClient}
	_, err := solutionClient.getSolution(context.Background(), nil, "")
	assert.EqualError(t, err, "Solution id is required")
}

func handleSolutionResponseMock(_ *http.Response) (*Solution, error) {
	return &Solution{}, nil
}

func handleSolutionResponseErrorMock(_ *http.Response) (*Solution, error) {
	return nil, errors.New("Failed to handle response")
}

func TestGetSolutionReturnsAnErrorIfTheResponseFailsToConvertToASolution(t *testing.T) {
	baseClient := &basesdk.BaseClientMock{}
	solutionClient := SolutionClient{SDKClient: baseClient}
	_, err := solutionClient.getSolution(context.Background(), handleSolutionResponseErrorMock, "SOL-123")
	assert.EqualError(t, err, "Failed to handle response")
}

func TestGetSolutionReturnsASolution(t *testing.T) {
	baseClient := &basesdk.BaseClientMock{}
	solutionClient := SolutionClient{SDKClient: baseClient}
	s, err := solutionClient.getSolution(context.Background(), handleSolutionResponseMock, "SOL-123")
	assert.Nil(t, err)
	assert.NotNil(t, s)
}

func TestGetSolutionReturnsErrorIfGetFails(t *testing.T) {
	baseClient := &basesdk.BaseClientMock{Error: errors.New("Get Failed")}
	solutionClient := SolutionClient{SDKClient: baseClient}
	_, err := solutionClient.getSolution(context.Background(), nil, "SO-123")
	assert.EqualError(t, err, "Get Failed")
}

func TestListSolutionsByIdsReturnsAnErrorIfNoPartnerIdIsProvided(t *testing.T) {
	baseClient := &basesdk.BaseClientMock{}
	solutionClient := SolutionClient{SDKClient: baseClient}
	_, err := solutionClient.listSolutionsByIds(context.Background(), nil, "", "", []string{})
	assert.EqualError(t, err, "partnerID is required")
}

func TestListSolutionsByIdsReturnsAnErrorIfNoSolutionIdsAreProvided(t *testing.T) {
	baseClient := &basesdk.BaseClientMock{}
	solutionClient := SolutionClient{SDKClient: baseClient}
	_, err := solutionClient.listSolutionsByIds(context.Background(), nil, "RLW", "", []string{})
	assert.EqualError(t, err, "solutionIDs are required")
}

func TestListSolutionsByIdsReturnsAnErrorIfGetFails(t *testing.T) {
	baseClient := &basesdk.BaseClientMock{Error: errors.New("Get Failed")}
	solutionClient := SolutionClient{SDKClient: baseClient}
	_, err := solutionClient.listSolutionsByIds(context.Background(), nil, "RLW", "", []string{"SO-123"})
	assert.EqualError(t, err, "Get Failed")
}

func handleListSolutionResponseErrorMock(_ *http.Response) ([]*Solution, error) {
	return nil, errors.New("Failed to handle response")
}

func TestListSolutionsByIdsReturnsAnErrorIfHandleResponseFails(t *testing.T) {
	baseClient := &basesdk.BaseClientMock{}
	solutionClient := SolutionClient{SDKClient: baseClient}
	_, err := solutionClient.listSolutionsByIds(context.Background(), handleListSolutionResponseErrorMock, "RLW", "", []string{"SO-123"})
	assert.EqualError(t, err, "Failed to handle response")
}

func handleListSolutionResponseMock(_ *http.Response) ([]*Solution, error) {
	return []*Solution{}, nil
}

func TestListSolutionsByIdsReturnsAListOfSolutions(t *testing.T) {
	baseClient := &basesdk.BaseClientMock{}
	solutionClient := SolutionClient{SDKClient: baseClient}
	s, err := solutionClient.listSolutionsByIds(context.Background(), handleListSolutionResponseMock, "RLW", "", []string{"SO-123"})
	assert.Nil(t, err)
	assert.NotNil(t, s)
}
