package cssdk

import (
	"encoding/json"
	"errors"
	"io"
)

// Source is a listing source
type Source struct {
	Name           string `json:"name"`
	SourceID       int64  `json:"sourceId"`
	IconURL16px    string `json:"iconUrl16px"`
	IconURL32px    string `json:"iconUrl32px"`
	IconURL50px    string `json:"iconUrl50px"`
	IconClass16px  string `json:"iconClass16px"`
	IconClass32px  string `json:"iconClass32px"`
	IconClass50px  string `json:"iconClass50px"`
	SourceTypeName string `json:"sourceTypeName"`
	SourceTypeID   string `json:"sourceTypeId"`
	EditListingURL string `json:"editListingUrl"`
}

// sourcesFromResponse converts an http response from core services to a list of sources
func sourcesFromResponse(r io.Reader) ([]*Source, error) {
	type Response struct {
		Sources []*Source `json:"data"`
	}
	res := Response{}
	if err := json.NewDecoder(r).Decode(&res); err != nil {
		reason := "Failed to convert response to Sources: " + err.Error()
		return nil, errors.New(reason)
	}
	return res.Sources, nil
}
