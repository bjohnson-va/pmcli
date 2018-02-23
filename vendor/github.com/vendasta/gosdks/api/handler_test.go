package api

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockHandler struct {
	AuthError         error
	Args              int
	ParseArgsError    error
	ValidateArgsError error
	ProcessString     string
	ProcessError      error
}

type MockHandlerArgs struct{}

func (h MockHandler) Authorize(r *http.Request) error { return h.AuthError }
func (h MockHandler) ParseArgs(r *http.Request) (map[string]interface{}, error) {
	return map[string]interface{}{}, h.ParseArgsError
}
func (h MockHandler) ValidateArgs(args map[string]interface{}) error { return h.ValidateArgsError }
func (h MockHandler) Process(ctx context.Context, args map[string]interface{}) (string, error) {
	return h.ProcessString, h.ProcessError
}

var handler = MockHandler{
	AuthError:         nil,
	ParseArgsError:    nil,
	ValidateArgsError: nil,
	ProcessString:     "hello",
	ProcessError:      nil,
}

func ResetMockHandler() {
	handler = MockHandler{
		AuthError:         nil,
		ParseArgsError:    nil,
		ValidateArgsError: nil,
		ProcessString:     "hello",
		ProcessError:      nil,
	}
	return
}

func TestReturns401ErrorIfAuthorizeFunctionReturnsError(t *testing.T) {
	defer ResetMockHandler()
	handler.AuthError = errors.New("Error")
	req := httptest.NewRequest("GET", "http://www.vendasta.com/api/v1", bytes.NewReader([]byte("{siteId: \"abc123\"}")))
	recorder := httptest.NewRecorder()
	handle_func := Handle(handler)
	handle_func(recorder, req)
	if recorder.Code != 401 {
		t.Errorf("Expected 401 code. Got %d", recorder.Code)
	}
}

func TestReturns400ErrorIfParseArgsReturnsError(t *testing.T) {
	defer ResetMockHandler()
	handler.ParseArgsError = errors.New("Error")
	req := httptest.NewRequest("GET", "http://www.vendasta.com/api/v1", bytes.NewReader([]byte("{siteId: \"abc123\"}")))
	recorder := httptest.NewRecorder()
	handle_func := Handle(handler)
	handle_func(recorder, req)
	if recorder.Code != 400 {
		t.Errorf("Expected 400 code. Got %d", recorder.Code)
	}
}

func TestReturns400ErrorIfValidateArgsReturnsError(t *testing.T) {
	defer ResetMockHandler()
	handler.ValidateArgsError = errors.New("Error")
	req := httptest.NewRequest("GET", "http://www.vendasta.com/api/v1", bytes.NewReader([]byte("{siteId: \"abc123\"}")))
	recorder := httptest.NewRecorder()
	handle_func := Handle(handler)
	handle_func(recorder, req)
	if recorder.Code != 400 {
		t.Errorf("Expected 400 code. Got %d", recorder.Code)
	}
}

func TestReturns500ErrorIfValidateArgsReturnsError(t *testing.T) {
	defer ResetMockHandler()
	handler.ProcessError = errors.New("Error")
	req := httptest.NewRequest("GET", "http://www.vendasta.com/api/v1", bytes.NewReader([]byte("{siteId: \"abc123\"}")))
	recorder := httptest.NewRecorder()
	handle_func := Handle(handler)
	handle_func(recorder, req)
	if recorder.Code != 500 {
		t.Errorf("Expected 500 code. Got %d", recorder.Code)
	}
}

func TestReturnsOutputFromProcessOnHttpResponseInDataJson(t *testing.T) {
	defer ResetMockHandler()
	processOutput := "This is my output"
	handler.ProcessString = processOutput
	req := httptest.NewRequest("GET", "http://www.vendasta.com/api/v1", bytes.NewReader([]byte("{siteId: \"abc123\"}")))
	recorder := httptest.NewRecorder()
	handle_func := Handle(handler)
	handle_func(recorder, req)
	result := recorder.Body.String()
	expected := fmt.Sprintf("{\"data\":\"%s\"}\n", processOutput)
	if result != expected {
		t.Errorf("Expected %s. Got %s", expected, result)
	}
}
