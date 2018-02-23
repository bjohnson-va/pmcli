package logging

import (
	"net/http"

	"time"

	gce_metadata "cloud.google.com/go/compute/metadata"
	"context"
)

func newLoggedResponse(w http.ResponseWriter) *loggedResponse {
	return &loggedResponse{w, 200, 0}
}

type loggedResponse struct {
	http.ResponseWriter
	status int
	length int
}

func (l *loggedResponse) WriteHeader(status int) {
	l.status = status
	l.ResponseWriter.WriteHeader(status)
}

func (l *loggedResponse) Write(b []byte) (n int, err error) {
	n, err = l.ResponseWriter.Write(b)
	l.length += n
	return
}

// HTTPMiddleware provides logging/tracing for incoming http requests.
func HTTPMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {

		ctx := request.Context()
		ctx, requestData := newRequest(ctx, GetLogger().RequestID())
		request = request.WithContext(ctx)

		ctx, span := startTracing(ctx, request.URL) // For StackDriver metrics
		defer span.Finish()

		response := newLoggedResponse(w)

		start := time.Now()
		h.ServeHTTP(response, request)
		end := time.Now()

		bundleID := GetLogger().RequestID()
		if span != nil {
			bundleID = span.TraceID()
		}
		logRequestWithBundling(ctx, bundleID, requestData, request, response, end.Sub(start))
	})
}

func logRequestWithBundling(ctx context.Context, bundleID string, requestData *requestData,
	request *http.Request, response *loggedResponse, latency time.Duration) {
	bundler := getBundler()
	ctx = bundler.applyBundlingMetadata(ctx, bundleID, requestData, request, response)
	requestData.HTTPRequest.Request = request
	requestData.HTTPRequest.Status = response.status
	requestData.HTTPRequest.Latency = latency
	logRequest(ctx, requestData)
}

func getBundler() logBundler {
	if gce_metadata.OnGCE() {
		return &GoogleContainerEngineLogBundler{}
	}
	return &NonBundlingLogBundler{}
}
