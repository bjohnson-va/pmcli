package billingsdk

import (
	"encoding/json"
	"errors"
	"io"
)

// LineItem is an item purchased
type LineItem struct {
	Count        int    `json:"count"`
	Total        int    `json:"total"`
	ProductID    string `json:"product_id"`
	PricePerUnit int    `json:"price_per_unit"`
	ProductName  string `json:"product_name"`
}

// LineItemResponse contains the lineitems and the offset
type LineItemResponse struct {
	LineItems []*LineItem `json:"line_items"`
	Offset    int         `json:"offset"`
}

// lineItemsFromResponse converts an http response from billing system to a list of lineitems
func lineItemsFromResponse(r io.Reader) (*LineItemResponse, error) {
	type Response struct {
		Resp *LineItemResponse `json:"data"`
	}
	res := Response{}
	if err := json.NewDecoder(r).Decode(&res); err != nil {
		reason := "Failed to convert response to Lineitem: " + err.Error()
		return nil, errors.New(reason)
	}
	return res.Resp, nil
}
