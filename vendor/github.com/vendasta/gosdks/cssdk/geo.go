package cssdk

import (
	"io"
	"encoding/json"
	"errors"
)

type GeoPoint struct {
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// taxonomiesFromResponse converts an http response from core services to a list of taxonomies
func geopointFromResponse(r io.Reader) (*GeoPoint, error) {
	type Response struct {
		GeoPoint *GeoPoint `json:"data"`
	}
	res := Response{}
	if err := json.NewDecoder(r).Decode(&res); err != nil {
		reason := "Failed to convert response to Geopoint: " + err.Error()
		return nil, errors.New(reason)
	}

	return res.GeoPoint, nil
}
