package basesdk

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"golang.org/x/net/context"
)

type nopCloser struct {
	io.Reader
}

func (nopCloser) Close() error { return nil }

type BaseClientMock struct {
	Error         error
	JSONBody      string
	PathRequested string
	ParamsSent    map[string]interface{}
}

func (client *BaseClientMock) Get(ctx context.Context, path string, params map[string]interface{}, ro ...RequestOpts) (*http.Response, error) {
	resp := nopCloser{bytes.NewBufferString(client.JSONBody)}
	client.PathRequested = path
	client.ParamsSent = params
	return &http.Response{Body: resp}, client.Error
}

func (client *BaseClientMock) Post(ctx context.Context, path string, params map[string]interface{}, ro ...RequestOpts) (*http.Response, error) {
	resp := nopCloser{bytes.NewBufferString(client.JSONBody)}
	client.PathRequested = path
	client.ParamsSent = params
	return &http.Response{Body: resp}, client.Error
}

func (client *BaseClientMock) Head(ctx context.Context, path string, ro ...RequestOpts) error {
	client.PathRequested = path
	return client.Error
}

type BaseClientMultiCallMock struct {
	Mocks     []*BaseClientMock
	CallCount int
}

func (c *BaseClientMultiCallMock) Get(ctx context.Context, path string, params map[string]interface{}, ro ...RequestOpts) (*http.Response, error) {
	resp, err := c.Mocks[c.CallCount].Get(ctx, path, params, ro...)
	c.CallCount = c.CallCount + 1
	return resp, err
}

func (c *BaseClientMultiCallMock) Post(ctx context.Context, path string, params map[string]interface{}, ro ...RequestOpts) (*http.Response, error) {
	fmt.Print("POST")
	resp, err := c.Mocks[c.CallCount].Post(ctx, path, params, ro...)
	c.CallCount = c.CallCount + 1
	return resp, err
}

func (c *BaseClientMultiCallMock) Head(ctx context.Context, path string, ro ...RequestOpts) error {
	err := c.Mocks[c.CallCount].Head(ctx, path, ro...)
	return err
}
