package cssdk

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_socialProfileFromResponse(t *testing.T) {
	type testCase struct {
		input          string
		expectedResult *SocialProfile
		expectedError  error
	}

	cases := []*testCase{
		{
			input:          "{\"version\": \"1.0\", \"data\": {\"spid\": \"SCP-123\", \"spgid\": \"SPG-444\"}}",
			expectedResult: &SocialProfile{SocialProfileID: "SCP-123"},
			expectedError:  nil,
		},
		{
			input:          "{\"version\": \"1.0\", \"data\": {\"spid\": \"SCP-123\"}}",
			expectedResult: &SocialProfile{SocialProfileID: "SCP-123"},
			expectedError:  nil,
		},
		{
			input:          "{\"version\": \"1.0\", \"data\": {\"spid\": \"SCP-123\", \"other\": \"stuff\"}}",
			expectedResult: &SocialProfile{SocialProfileID: "SCP-123"},
			expectedError:  nil,
		},
		{
			input:          "{badjson}",
			expectedResult: nil,
			expectedError:  errors.New("Failed to convert response to SocialProfile: invalid character 'b' looking for beginning of object key string"),
		},
	}

	for _, c := range cases {
		response := c.input

		handler := func(w http.ResponseWriter, r *http.Request) {
			io.WriteString(w, response)
		}
		req := httptest.NewRequest("GET", "http://example.com/foo", nil)
		w := httptest.NewRecorder()
		handler(w, req)
		resp := w.Result()
		result, err := socialProfileFromResponse(resp)

		assert.Equal(t, c.expectedResult, result)
		assert.Equal(t, c.expectedError, err)
	}
}

func Test_SocialRegistrationFromResponse(t *testing.T) {
	type testCase struct {
		input          string
		expectedResult *SocialProfileRegistration
		expectedError  error
	}

	cases := []*testCase{
		{
			input:          "{\"version\": \"1.0\", \"data\": {\"spid\": \"SCP-123\", \"uid\": \"vbc\",  \"accountId\": \"AG-123\"}}",
			expectedResult: &SocialProfileRegistration{SocialProfileID: "SCP-123", AppID: "vbc", AccountID: "AG-123"},
			expectedError:  nil,
		},
		{
			input:          "{badjson}",
			expectedResult: nil,
			expectedError:  errors.New("Failed to convert response to SocialProfileRegistration: invalid character 'b' looking for beginning of object key string"),
		},
	}

	for _, c := range cases {
		response := c.input

		handler := func(w http.ResponseWriter, r *http.Request) {
			io.WriteString(w, response)
		}
		req := httptest.NewRequest("GET", "http://example.com/foo", nil)
		w := httptest.NewRecorder()
		handler(w, req)
		resp := w.Result()
		result, err := socialProfileRegistrationFromResponse(resp)

		assert.Equal(t, c.expectedResult, result)
		assert.Equal(t, c.expectedError, err)
	}
}

func Test_SocialRegistrationsFromResponse(t *testing.T) {
	type testCase struct {
		input          string
		expectedResult []*SocialProfileRegistration
		expectedError  error
	}

	cases := []*testCase{
		{
			input: "{\"version\": \"1.0\", \"data\": [{\"spid\": \"SCP-123\", \"uid\": \"vbc\",  \"accountId\": \"AG-123\"}]}",
			expectedResult: []*SocialProfileRegistration{
				{SocialProfileID: "SCP-123", AppID: "vbc", AccountID: "AG-123"},
			},
			expectedError: nil,
		},
		{
			input:          "{\"version\": \"1.0\", \"data\": []}",
			expectedResult: []*SocialProfileRegistration{},
			expectedError:  nil,
		},
		{
			input:          "{badjson}",
			expectedResult: nil,
			expectedError:  errors.New("Failed to convert response to list of SocialProfileRegistrations: invalid character 'b' looking for beginning of object key string"),
		},
	}

	for _, c := range cases {
		response := c.input

		handler := func(w http.ResponseWriter, r *http.Request) {
			io.WriteString(w, response)
		}
		req := httptest.NewRequest("GET", "http://example.com/foo", nil)
		w := httptest.NewRecorder()
		handler(w, req)
		resp := w.Result()
		result, err := socialProfileRegistrationsFromResponse(resp)

		assert.Equal(t, c.expectedResult, result)
		assert.Equal(t, c.expectedError, err)
	}
}
