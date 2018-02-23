package marketplace

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/vendasta/gosdks/basesdk"
	"github.com/vendasta/gosdks/config"
	"golang.org/x/net/context"
)

// SolutionClient is a client which handles calls to marketplace's solution apis
type SolutionClient struct {
	basesdk.SDKClient
}

// BuildSolutionClient creates a solution client to allow calls to be made to marketplace.
func BuildSolutionClient(apiUser string, apiKey string, env config.Env, rootURLOverride string) SolutionClient {
	var rootURL string
	if rootURLOverride != "" {
		rootURL = rootURLOverride
	} else {
		rootURL = rootURLFromEnv(env)
	}
	auth := basesdk.UserKey{APIUser: apiUser, APIKey: apiKey}
	baseClient := basesdk.BaseClient{Authorization: auth, RootURL: rootURL}
	return SolutionClient{baseClient}
}

// GetSolutionResponseHandler is a function which handle the http response from the solutions api
type GetSolutionResponseHandler func(r *http.Response) (*Solution, error)

// GetSolution returns a solution based on the solution id provided
func (c SolutionClient) GetSolution(ctx context.Context, solutionID string) (*Solution, error) {
	return c.getSolution(ctx, SolutionFromResponse, solutionID)
}

func (c SolutionClient) getSolution(ctx context.Context, responseHandler GetSolutionResponseHandler, solutionID string) (*Solution, error) {
	if solutionID == "" {
		return nil, errors.New("Solution id is required")
	}
	path := "/marketplaceInternalApi/v1/solution/get"
	params := map[string]interface{}{"solutionId": solutionID}
	response, err := c.Get(ctx, path, params, basesdk.Idempotent())
	if err != nil {
		return nil, err
	}
	solution, err := responseHandler(response)
	if err != nil {
		return nil, err
	}
	return solution, nil
}

// ListSolutionResponseHandler is an http response handler which returns a list of solutions
type ListSolutionResponseHandler func(r *http.Response) ([]*Solution, error)

// ListSolutionsByIds returns a list of solutions based on the partner id, market id and solution ids provided
func (c SolutionClient) ListSolutionsByIds(ctx context.Context, partnerID string, marketID string, solutionIDs []string) ([]*Solution, error) {
	return c.listSolutionsByIds(ctx, SolutionListFromResponse, partnerID, marketID, solutionIDs)
}

func (c SolutionClient) listSolutionsByIds(ctx context.Context, responseHandler ListSolutionResponseHandler, partnerID string, marketID string, solutionIDs []string) ([]*Solution, error) {
	if partnerID == "" {
		return nil, errors.New("partnerID is required")
	}
	numberOfSolutionIds := len(solutionIDs)
	if numberOfSolutionIds == 0 {
		return nil, errors.New("solutionIDs are required")
	}
	pageSize := strconv.Itoa(numberOfSolutionIds)
	path := "/marketplaceInternalApi/v1/solution/list"
	params := map[string]interface{}{
		"partnerId":   partnerID,
		"marketId":    marketID,
		"solutionIds": solutionIDs,
		"pageSize":    pageSize,
		"cursor":      "0",
	}
	response, err := c.Get(ctx, path, params, basesdk.Idempotent())
	if err != nil {
		return nil, err
	}
	solutions, err := responseHandler(response)
	if err != nil {
		return nil, err
	}
	return solutions, nil
}
