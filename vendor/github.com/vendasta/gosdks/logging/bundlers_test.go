package logging

import (
	"context"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestGoogleContainerEngineLogBundlerInterfaceSanityCheck(t *testing.T) {
	var bundler logBundler
	bundler = &GoogleContainerEngineLogBundler{}
	_ = bundler // satisfy compiler
}

func TestNonBundlingLogBundlerInterfaceSanityCheck(t *testing.T) {
	var bundler logBundler
	bundler = &NonBundlingLogBundler{}
	_ = bundler // satisfy compiler
}

func Test_GCELogBundler_ApplyBundlingMetadataTests(t *testing.T) {
	suite.Run(t, new(ApplyBundlingMetadataTests))
}

type ApplyBundlingMetadataTests struct {
	suite.Suite
	bundler GoogleContainerEngineLogBundler
}

func (b *ApplyBundlingMetadataTests) SetupTest() {
	b.bundler = GoogleContainerEngineLogBundler{}
}

func (b *ApplyBundlingMetadataTests) Test_SetsCorrectHTTPRequestValues() {
	request := http.Request{
		URL:    &url.URL{Path: "http://example.com"},
		Method: "POST",
	}
	response := loggedResponse{length: 100}
	_, requestData := newRequest(context.Background(), "")

	bundleID := "abc123"
	b.bundler.applyBundlingMetadata(context.Background(), bundleID, requestData, &request, &response)

	b.Equal("http://example.com", requestData.HTTPRequest.Request.URL.Path)
	b.Equal("POST", requestData.HTTPRequest.Request.Method)
	b.EqualValues(100, requestData.HTTPRequest.ResponseSize)
	b.NotNil(requestData.HTTPRequest.LocalIP)  // Could probably be more specific
	b.NotNil(requestData.HTTPRequest.RemoteIP) // Could probably be more specific
	b.Equal(bundleID, requestData.Trace)
}
