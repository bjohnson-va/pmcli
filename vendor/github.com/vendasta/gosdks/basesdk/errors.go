package basesdk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/vendasta/gosdks/util"
)

func parseError(r *http.Response) error {
	body := ""
	if r.Body != nil {
		defer r.Body.Close()
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return err
		}
		body = string(bodyBytes)
	}

	vapiError := &VAPIError{}
	err := json.Unmarshal([]byte(body), vapiError)
	if err != nil || !vapiError.isPopulated() {
		errorBody := fmt.Sprintf("%s: %s", http.StatusText(r.StatusCode), body)
		return &HTTPError{StatusCode: r.StatusCode, Body: errorBody}
	}
	return vapiError
}

// HTTPError represents an error when calling a service over http
type HTTPError struct {
	Body       string
	StatusCode int
}

func (e *HTTPError) Error() string {
	if e == nil {
		return ""
	}
	return e.Body
}

//StatusCode returns the http status code of a response or http.StatusInternalServerError if it could not be decoded
func StatusCode(e error) int {
	switch err := e.(type) {
	case *HTTPError:
		return err.StatusCode
	case *VAPIError:
		return err.StatusCode
	default:
		return http.StatusInternalServerError
	}
}

//VAPIError is an api error returned from our vapi endpoints
type VAPIError struct {
	Message      string `json:"message"`
	Version      string `json:"version"`
	RequestID    string `json:"requestId"`
	ResponseTime int64  `json:"responseTime"`
	StatusCode   int    `json:"statusCode"`
}

func (e *VAPIError) Error() string {
	if e == nil {
		return ""
	}
	return e.Message
}

func (e *VAPIError) isPopulated() bool {
	if e.Message == "" || e.StatusCode == 0 {
		return false
	}
	return true
}

func ConvertHttpErrorToGRPC(err error) error {
	switch e := err.(type) {
	case *HTTPError:
		return util.Error(util.StatusCodeToGRPCError(e.StatusCode), e.Error())
	case *VAPIError:
		return util.Error(util.StatusCodeToGRPCError(e.StatusCode), e.Error())
	default:
		return util.Error(util.Internal, "%s", err.Error())
	}
}
