package logging

import (
	"net/http"
	"context"
	"google.golang.org/grpc/metadata"
)

// LogBundler bundles log messages together by surrounding them with a start and end function.
type logBundler interface {
	applyBundlingMetadata(ctx context.Context, bundleID string, rd *requestData,
		request *http.Request, response *loggedResponse) context.Context
}

// GoogleContainerEngineLogBundler ensures log messages are bundled under a single, collapsible request in cloud logging
type GoogleContainerEngineLogBundler struct {
}

func (g *GoogleContainerEngineLogBundler) applyBundlingMetadata(ctx context.Context, bundleID string, rd *requestData,
	request *http.Request, response *loggedResponse) context.Context {
	rd.HTTPRequest.Request.URL = request.URL
	rd.HTTPRequest.Request.Method = request.Method
	rd.HTTPRequest.Status = int(response.status)
	rd.HTTPRequest.ResponseSize = int64(response.length)
	rd.HTTPRequest.LocalIP = "127.0.0.1"
	rd.HTTPRequest.RemoteIP = request.RemoteAddr
	rd.Trace = bundleID
	md, _ := metadata.FromOutgoingContext(ctx)
	ctx = metadata.NewOutgoingContext(ctx, md)
	return ctx
}

// NonBundlingLogBundler satisfies the LogBundler interface but DOES NOT apply any bundling to logs.
type NonBundlingLogBundler struct {
}

func (n *NonBundlingLogBundler) applyBundlingMetadata(ctx context.Context, bundleID string, rd *requestData,
	request *http.Request, response *loggedResponse) context.Context {
	return ctx
}
