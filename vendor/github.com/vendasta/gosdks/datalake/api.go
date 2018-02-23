package datalake

import (
	"context"
	"fmt"
	"strings"

	"github.com/vendasta/gosdks/pb/datalakeproto"
	"github.com/vendasta/gosdks/vax"
	"google.golang.org/grpc"
)

var (
	scopes = map[string]string{
		"LOCAL": "",
		"TEST":  "https://data-lake-api-test.vendasta-internal.com",
		"DEMO":  "https://data-lake-api-demo.vendasta-internal.com",
		"PROD":  "https://data-lake-api-prod.vendasta-internal.com",
	}

	urls = map[string]string{
		"LOCAL": "127.0.0.1:11000",
		"TEST":  "data-lake-api-test.vendasta-internal.com:443",
		"DEMO":  "data-lake-api-demo.vendasta-internal.com:443",
		"PROD":  "data-lake-api-prod.vendasta-internal.com:443",
	}
)

// Interface of the datalake api client
type Interface interface {
	GetReview(reviewId string) (*datalakeproto.Review, error)
	ReplaceReview(review *datalakeproto.Review) error
	/* Returns list of reviews, cursor, hasMore, total and error in that order */
	ListReviews(listingId string, cursor string, pageSize int64) ([]*datalakeproto.Review, string, bool, int64, error)
	DeleteReview(reviewId string) error

	GetListing(rawListingId string, externalId string, sourceId int64) (*datalakeproto.RawListing, error)
	GetListingByURL(url string, sourceId int64) (*datalakeproto.RawListing, error)
	ReplaceListing(rawListing *datalakeproto.RawListing) error
	/* Returns list of listings, cursor, hasMore, total and error in that order */
	SearchListings(cursor string, pageSize int64, companyName string, city string, state string, zipCode string, phone string, address string, sourceId int64) ([]*datalakeproto.RawListing, bool, string, error)
	DeleteListing(rawListingId string) error

	GetListingMetadata() ([]*datalakeproto.SourceStatistic, error)
	GetListingStats() ([]*datalakeproto.SourceStats, []*datalakeproto.SourceStats, error)
}

type datalakeGRPCClient struct {
	url        string
	context    context.Context
	gRPCClient datalakeproto.DataLakeClient
}

func (client *datalakeGRPCClient) GetReview(reviewId string) (*datalakeproto.Review, error) {
	request := &datalakeproto.GetReviewRequest{ReviewId: reviewId}
	response, err := client.gRPCClient.GetReview(client.context, request)
	if err != nil {
		return nil, err
	}
	return response.Review, nil
}

func (client *datalakeGRPCClient) ReplaceReview(review *datalakeproto.Review) error {
	request := &datalakeproto.ReplaceReviewRequest{Review: review}
	_, err := client.gRPCClient.ReplaceReview(client.context, request)
	return err
}

func (client *datalakeGRPCClient) ListReviews(rawListingId string, cursor string, pageSize int64) ([]*datalakeproto.Review, string, bool, int64, error) {
	request := &datalakeproto.ListReviewsRequest{ListingId: rawListingId, Cursor: cursor, PageSize: pageSize}
	response, err := client.gRPCClient.ListReviews(client.context, request)
	if err != nil {
		return nil, "", false, -1, err
	}
	return response.Reviews, response.Cursor, response.HasMore, response.TotalNumberOfReviews, nil

}

func (client *datalakeGRPCClient) DeleteReview(reviewId string) error {
	request := &datalakeproto.DeleteReviewRequest{ReviewId: reviewId}
	_, err := client.gRPCClient.DeleteReview(client.context, request)
	return err
}

func (client *datalakeGRPCClient) GetListing(rawListingId string, externalId string, sourceId int64) (*datalakeproto.RawListing, error) {
	request := &datalakeproto.GetListingRequest{RawListingId: rawListingId, ExternalId: externalId, SourceId: sourceId}
	response, err := client.gRPCClient.GetListing(client.context, request)
	if err != nil {
		return nil, err
	}
	return response.Listing, nil
}

func (client *datalakeGRPCClient) GetListingByURL(url string, sourceId int64) (*datalakeproto.RawListing, error) {
	request := &datalakeproto.GetListingByUrlRequest{Url: url, SourceId: sourceId}
	response, err := client.gRPCClient.GetListingByURL(client.context, request)
	if err != nil {
		return nil, err
	}
	return response.Listing, nil
}

func (client *datalakeGRPCClient) ReplaceListing(rawListing *datalakeproto.RawListing) error {
	request := &datalakeproto.ReplaceListingRequest{Listing: rawListing}
	_, err := client.gRPCClient.ReplaceListing(client.context, request)
	return err
}

func (client *datalakeGRPCClient) SearchListings(cursor string, pageSize int64, companyName string, city string, state string, zipCode string, phone string, address string, sourceId int64) ([]*datalakeproto.RawListing, bool, string, error) {
	request := &datalakeproto.SearchListingsRequest{Cursor: cursor, PageSize: pageSize, CompanyName: companyName, City: city,
		State: state, Zipcode: zipCode, Phone: phone, Address: address,
		SourceId: sourceId}
	response, err := client.gRPCClient.SearchListings(client.context, request)

	if err != nil {
		return nil, false, "", err
	}

	return response.Results, response.HasMore, response.Cursor, nil
}

func (client *datalakeGRPCClient) GetListingMetadata() ([]*datalakeproto.SourceStatistic, error) {
	request := &datalakeproto.GetListingMetadataRequest{}
	response, err := client.gRPCClient.GetListingMetadata(client.context, request)
	if err != nil {
		return nil, err
	}
	return response.Sources, nil
}

func (client *datalakeGRPCClient) GetListingStats() ([]*datalakeproto.SourceStats, []*datalakeproto.SourceStats, error) {
	request := &datalakeproto.GetListingStatsRequest{}
	response, err := client.gRPCClient.GetListingStats(client.context, request)
	if err != nil {
		return nil, nil, err
	}
	return response.ReplacedPerDay, response.TotalPerDay, err
}

func (client *datalakeGRPCClient) DeleteListing(rawListingId string) error {
	request := &datalakeproto.DeleteListingRequest{ListingId: rawListingId}
	_, err := client.gRPCClient.DeleteListing(client.context, request)
	return err
}

// NewGRPCClient creates a new datalake client
func NewGRPCClient(context context.Context, environment string, options ...grpc.DialOption) (Interface, error) {
	var url, scope, env string
	var err error
	var gRPCClient *grpc.ClientConn
	var gRPCDataLakeClient datalakeproto.DataLakeClient

	env = strings.ToUpper(environment)
	scope = scopes[env]
	url = urls[env]
	if url == "" {
		return nil, fmt.Errorf("Invalid environment selected: %s", environment)
	}

	if env == "LOCAL" {
		gRPCClient, err = vax.NewGRPCConnection(context, url, false, scope, false)
	} else {
		gRPCClient, err = vax.NewGRPCConnection(context, url, true, scope, true)
	}

	if err != nil {
		return nil, err
	}
	gRPCDataLakeClient = datalakeproto.NewDataLakeClient(gRPCClient)
	return &datalakeGRPCClient{url: url, context: context, gRPCClient: gRPCDataLakeClient}, nil
}
