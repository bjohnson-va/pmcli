package cssdk

import (
	"encoding/json"
	"errors"
	"io"
	"time"

	"github.com/vendasta/gosdks/basesdk"
)

// Review is a review
type Review struct {
	Rating            string    `json:"rating"`
	AccountGroupID    string    `json:"agid"`
	ReviewID          string    `json:"rid"`
	ListingID         string    `json:"lid"`
	SourceID          int64     `json:"sourceId"`
	PublishedDateTime time.Time `json:"publishedDateTime"`
}

// ReviewLookupResponse is the response of the review lookup api
type ReviewLookupResponse struct {
	Reviews    []*Review
	NextCursor string
}

// reviewLookupResponseFromResponse converts an http response from core services to a reviewLookupResponse
func reviewLookupResponseFromResponse(r io.Reader) (*ReviewLookupResponse, error) {
	type Response struct {
		Reviews         []*Review `json:"data"`
		NextQueryString string    `json:"nextQueryString"`
	}
	res := Response{}
	if err := json.NewDecoder(r).Decode(&res); err != nil {
		reason := "Failed to convert response to Review: " + err.Error()
		return nil, errors.New(reason)
	}
	nextCursor, err := basesdk.ParseCursorFromVAPINextQueryString(res.NextQueryString)
	if err != nil {
		reason := "Error parsing cursor: " + err.Error()
		return nil, errors.New(reason)
	}

	return &ReviewLookupResponse{
		Reviews:    res.Reviews,
		NextCursor: nextCursor,
	}, nil
}

// ReviewStats are the stats about reviews
type ReviewStats struct {
	IndustryAverageCount  float32          `json:"industryAverageCount"`
	IndustryAverageRating float32          `json:"industryAverageRating"`
	TotalCount            int64            `json:"totalCount"`
	SourceCounts          map[string]int64 `json:"sourceCounts"`
	RatingCounts          map[string]int64 `json:"ratingCounts"`
}

// reviewReviewStatsFromResponse converts an http response from core services to ReviewStats
func reviewReviewStatsFromResponse(r io.Reader) (*ReviewStats, error) {
	type Response struct {
		ReviewStats *ReviewStats `json:"data"`
	}
	res := Response{}
	if err := json.NewDecoder(r).Decode(&res); err != nil {
		reason := "Failed to convert response to Review: " + err.Error()
		return nil, errors.New(reason)
	}
	return res.ReviewStats, nil
}
