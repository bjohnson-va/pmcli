package basesdk

import (
	"bytes"
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vendasta/gosdks/util"
)

func TestParseErrorReturnsErrorWithStatusCodeAndBodyMessage(t *testing.T) {
	resp := &http.Response{StatusCode: 400, Body: nopCloser{bytes.NewBufferString("Error doing stuff")}}
	actual := parseError(resp)
	assert.EqualError(t, actual, "Bad Request: Error doing stuff")
	assert.Equal(t, 400, StatusCode(actual))
}

func TestParseErrorReturnsErrorWithStatusCodeAndNoBodyMessageWhenThereIsNoBody(t *testing.T) {
	resp := &http.Response{StatusCode: 401, Body: nil}
	actual := parseError(resp)
	assert.EqualError(t, actual, "Unauthorized: ")
	assert.Equal(t, 401, StatusCode(actual))
}

var vapiErrorBody = `{
	"message": "Arg \"marketId\": Market Id GARBAGE is not valid. ",
	"version": "3.0",
	"requestId": "thisisalongrequestid",
	"responseTime": 14,
	"statusCode": 400
  }`

func TestParseErrorReturnsVAPIErrorWhenTheErrorLooksLikeAVAPIError(t *testing.T) {
	resp := &http.Response{StatusCode: 400, Body: nopCloser{bytes.NewBufferString(vapiErrorBody)}}
	actual := parseError(resp)
	assert.EqualError(t, actual, "Arg \"marketId\": Market Id GARBAGE is not valid. ")
	assert.Equal(t, 400, StatusCode(actual))

	vapiError, ok := actual.(*VAPIError)
	assert.Equal(t, true, ok)
	assert.EqualError(t, vapiError, "Arg \"marketId\": Market Id GARBAGE is not valid. ")
	assert.Equal(t, 400, StatusCode(vapiError))
	assert.Equal(t, "3.0", vapiError.Version)
	assert.Equal(t, "thisisalongrequestid", vapiError.RequestID)
	assert.Equal(t, int64(14), vapiError.ResponseTime)
}

var almostVapiErrorBody = `{
	"message": "Arg \"marketId\": Market Id GARBAGE is not valid. ",
	"version": "3.0",
	"requestId": "thisisalongrequestid",
	"responseTime": 14,
  }`

func TestParseErrorReturnsReturnsFullErrorIfErrorResponseLooksSimilarToAVAPIErrorButIsNot(t *testing.T) {
	resp := &http.Response{StatusCode: 400, Body: nopCloser{bytes.NewBufferString(almostVapiErrorBody)}}
	actual := parseError(resp)
	assert.EqualError(t, actual, "Bad Request: "+almostVapiErrorBody)
	assert.Equal(t, 400, StatusCode(actual))
}

func TestStatusCode(t *testing.T) {
	var tests = []struct {
		in  error
		out int
	}{
		{errors.New("non http erorr"), http.StatusInternalServerError},
		{&HTTPError{StatusCode: 500}, http.StatusInternalServerError},
		{&HTTPError{StatusCode: 400}, http.StatusBadRequest},
		{&HTTPError{StatusCode: 200}, http.StatusOK},
		{&VAPIError{StatusCode: 500}, http.StatusInternalServerError},
		{&VAPIError{StatusCode: 400}, http.StatusBadRequest},
		{&VAPIError{StatusCode: 200}, http.StatusOK},
	}
	for _, tt := range tests {
		s := StatusCode(tt.in)
		if s != tt.out {
			t.Errorf("Expected converting from %v to status code: %d but got %d", tt.in, tt.out, s)
		}
	}
}

func TestConvertHttpErrorToGRPC(t *testing.T) {
	var tests = []struct {
		in error
		out error
	}{
		{errors.New("not an http error"), util.Error(util.Internal, "not an http error")},
		{&HTTPError{"500 Error", 500}, util.Error(util.Internal, "500 Error")},
		{&HTTPError{"400 Error", 400}, util.Error(util.InvalidArgument, "400 Error")},
		{&HTTPError{"404 Error", 404}, util.Error(util.NotFound, "404 Error")},
		{&HTTPError{"409 Error", 409}, util.Error(util.AlreadyExists, "409 Error")},
		{&HTTPError{"403 Error", 403}, util.Error(util.PermissionDenied, "403 Error")},
		{&VAPIError{Message:"500 Error", StatusCode:500}, util.Error(util.Internal, "500 Error")},
		{&VAPIError{Message:"400 Error", StatusCode:400}, util.Error(util.InvalidArgument, "400 Error")},
		{&VAPIError{Message:"404 Error", StatusCode:404}, util.Error(util.NotFound, "404 Error")},
		{&VAPIError{Message:"409 Error", StatusCode:409}, util.Error(util.AlreadyExists, "409 Error")},
		{&VAPIError{Message:"403 Error", StatusCode:403}, util.Error(util.PermissionDenied, "403 Error")},
	}

	for _, tt := range tests {
		err := ConvertHttpErrorToGRPC(tt.in)
		if err != tt.out {
			t.Errorf("Expected to get %v, but got %v", tt.out, err)
		}

	}
}
