package reserveidsdk

import (
	"encoding/json"
	"github.com/vendasta/gosdks/util"
	"io/ioutil"
	"net/http"
)

// ReserveIDResponse represents the result of reserving vendor internal ID
type ReserveIDResponse struct {
	Status           int    `json:"status"`
	VendorInternalID string `json:"internal_id"`
	Reason           string `json:"reason"`
}

// reserveIDResponseFromResponse converts an http response from VDC to a ReserveIDResponse
// The response body of ReserveId could be:
//    {"data": {"status": 200, "internal_id": "ID-123", "reason": ""}}
// or {"data": {"status": 200, "internal_id": "", "reason": "No reserve_id_url"}}
// or {"data": {"status": 404, "internal_id": "", "reason": "some error"}}
func reserveIDResponseFromResponse(r *http.Response) (*ReserveIDResponse, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, util.Error(util.Internal, "Failed to read the body of reserve id response: %s", err.Error())
	}
	defer r.Body.Close()
	type Response struct {
		ReserveIDResponse *ReserveIDResponse `json:"data"`
	}
	res := Response{}

	if err := json.Unmarshal(body, &res); err != nil {
		return nil, util.Error(util.Internal, "Failed to convert response to ReserveIDResult: %s,", err.Error())
	}

	return res.ReserveIDResponse, nil
}
