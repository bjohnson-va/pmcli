package marketplace

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

// Price represents a solution's price and frequency
type Price struct {
	// Price of a solution
	Price int64 `json:"price"`

	// BillingFrequency that this price occurs in a solution (once, monthly, yearly, etc.)
	BillingFrequency BillingFrequency `json:"frequency"`
}

// Pricing represents a solution's currency and list of prices
type Pricing struct {
	// Currency for the list of prices
	Currency Currency `json:"currency"`
	// Prices is a list of multiple price options
	Prices []*Price `json:"prices"`
}

// Solution represents a marketplace solution/package
type Solution struct {
	// SolutionID is a unique identifier for a solution or package in marketplace
	SolutionID string `json:"solution_id"`

	// Name is the name of the solution
	Name string `json:"name"`

	// PartnerID is the identifier of the partner who created the solution
	PartnerID string `json:"partner_id"`

	// MarketID is the identifier of the market this solution is available for
	MarketID string `json:"market_id"`

	// IconURL is the url to the icon used for this solution. Solutions are not required to have an icon url
	IconURL string `json:"icon"`

	// Content is marketing content which describes the solution
	Content string `json:"content"`

	// Currency is the system of money used for this solution
	// Deprecated: Use the Currency on Pricing
	Currency Currency `json:"currency"`

	// BillingFrequency describes how often the purchaser would be billed for this solution
	// Deprecated: Use the BillingFrequency on the Price in Pricing
	BillingFrequency BillingFrequency `json:"billing_frequency"`

	// BillingFrequencyOther describes the billing frequency if it differs from the standard set
	BillingFrequencyOther string `json:"billing_frequency_other"`

	// Created is the date and time the solution was made
	Created time.Time `json:"created"`

	// Updated is the date and time the solution was updated
	Updated time.Time `json:"updated"`

	// Archived is the date and time the solution was archived
	Archived time.Time `json:"archived"`

	// Status describes if the solution is draft, published or archived
	Status string `json:"status"`

	// Products is a list of marketplace product ids which are included in this solution
	Products []string `json:"products"`

	// SellingPrice is the amount the solution will be sold to a customer
	// Deprecated: Use Pricing instead
	SellingPrice float64 `json:"selling_price"`

	// Pricing contains the currency and pricing information of this solution
	Pricing *Pricing `json:"pricing"`

	// HideProductDetails is a flag which indicates the product details should be hidden
	HideProductDetails bool `json:"hide_product_details"`

	// HideProductIconsAndNames is a flag which indicates the product icons and names should be hidden
	HideProductIconsAndNames bool `json:"hide_product_icons_and_names"`
}

// SolutionFromResponse converts an http response from marketplace to a Solution
func SolutionFromResponse(r *http.Response) (*Solution, error) {
	defer r.Body.Close()
	type Response struct {
		Solution *Solution `json:"data"`
	}
	res := Response{}
	if err := json.NewDecoder(r.Body).Decode(&res); err != nil {
		reason := "Failed to convert response to Solution: " + err.Error()
		return nil, errors.New(reason)
	}
	return res.Solution, nil
}

// SolutionListFromResponse converts an http response from marketplace to a list of solutions
func SolutionListFromResponse(r *http.Response) ([]*Solution, error) {
	defer r.Body.Close()
	type Response struct {
		Solutions []*Solution `json:"data"`
	}
	res := Response{}
	if err := json.NewDecoder(r.Body).Decode(&res); err != nil {
		reason := "Failed to convert response to Solution: " + err.Error()
		return nil, errors.New(reason)
	}
	return res.Solutions, nil
}
