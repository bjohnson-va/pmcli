package basesdk

import (
	"net/http"
	"net/url"
)

// RequestAuthorization defines the functions to handle authorization between the sdk and server
type RequestAuthorization interface {
	SignRequest(in *http.Request)
}

// UserKey is a basic api user and api key combination for authorizing an http request
type UserKey struct {
	// APIUser is the name of the user performing the request
	APIUser string

	// APIKey is the key of the user performing the request
	APIKey string
}

// SignRequest takes in an http request pointer and adds the api user and key to the request
func (a UserKey) SignRequest(in *http.Request) {
	query := in.URL.Query()
	if a.APIKey != "" {
		query.Add("apiKey", a.APIKey)
	}
	if a.APIUser != "" {
		query.Add("apiUser", a.APIUser)
	}
	in.URL.RawQuery = url.Values.Encode(query)
}

// NoAuth satisfies the RequestAuthorization interface for http requests that do not require auth
type NoAuth struct{}

// SignRequest returns without signing the request
func (a NoAuth) SignRequest(in *http.Request) {
	return
}
