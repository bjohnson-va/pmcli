package cssdk

import (
	"encoding/json"
	"errors"
	"net/http"
)

// SocialProfile represents a social profile which could be connected to a social profile group
type SocialProfile struct {
	SocialProfileID string `json:"spid"`
}

// SocialProfileRegistration represents a social profile registration which connects an app (vbc, sr, etc..) and account (AG-123, etc...) with a social profile
type SocialProfileRegistration struct {
	SocialProfileID string `json:"spid"`
	AppID           string `json:"uid"`
	AccountID       string `json:"accountId"`
}

// socialProfileFromResponse converts an http response from core services to a social profile
func socialProfileFromResponse(r *http.Response) (*SocialProfile, error) {
	defer r.Body.Close()
	type Response struct {
		SocialProfile *SocialProfile `json:"data"`
	}
	res := Response{}
	if err := json.NewDecoder(r.Body).Decode(&res); err != nil {
		reason := "Failed to convert response to SocialProfile: " + err.Error()
		return nil, errors.New(reason)
	}
	return res.SocialProfile, nil
}

// socialProfileRegistrationFromResponse converts an http response from core services to a social profile registration
func socialProfileRegistrationFromResponse(r *http.Response) (*SocialProfileRegistration, error) {
	defer r.Body.Close()
	type Response struct {
		SocialProfileRegistration *SocialProfileRegistration `json:"data"`
	}
	res := Response{}
	if err := json.NewDecoder(r.Body).Decode(&res); err != nil {
		reason := "Failed to convert response to SocialProfileRegistration: " + err.Error()
		return nil, errors.New(reason)
	}
	return res.SocialProfileRegistration, nil
}

// socialProfileRegistrationsFromResponse converts an http response from core services to a list of social profile registrations
func socialProfileRegistrationsFromResponse(r *http.Response) ([]*SocialProfileRegistration, error) {
	defer r.Body.Close()
	type Response struct {
		SocialProfileRegistrations []*SocialProfileRegistration `json:"data"`
	}
	res := Response{}
	if err := json.NewDecoder(r.Body).Decode(&res); err != nil {
		reason := "Failed to convert response to list of SocialProfileRegistrations: " + err.Error()
		return nil, errors.New(reason)
	}
	return res.SocialProfileRegistrations, nil
}
