package logging

import (
	"sync"

	"time"

	"net/http"

	"cloud.google.com/go/logging"
	"golang.org/x/net/context"
	google_logging_type "google.golang.org/genproto/googleapis/logging/type"
)

type requestDataKey struct{}

type requestData struct {
	mu    sync.Mutex
	*logging.Entry
	lines []*logLine

	// Additional labels to add to the GKE request
	tags map[string]string

	// common labels will override tags and should only be filled with labels common to all requests
	commonLabels map[string]string
}

type logLine struct {
	time.Time
	Severity   google_logging_type.LogSeverity
	LogMessage string
}

func (rd *requestData) logLine(message string, severity google_logging_type.LogSeverity) {
	rd.mu.Lock()
	defer rd.mu.Unlock()

	rd.lines = append(rd.lines, &logLine{
		Time:       time.Now().UTC(),
		Severity:   severity,
		LogMessage: message,
	})
}

func (rd *requestData) addTag(key, value string) {
	rd.mu.Lock()
	defer rd.mu.Unlock()

	rd.tags[key] = value
}

func (rd *requestData) getLabels() map[string]string {
	r := map[string]string{}

	for k, v := range rd.tags {
		r[k] = v
	}

	for k, v := range rd.commonLabels {
		r[k] = v
	}

	return r
}

func newRequest(ctx context.Context, requestID string) (context.Context, *requestData) {
	rd := &requestData{
		Entry: &logging.Entry{
			HTTPRequest: &logging.HTTPRequest{
				Request: &http.Request{},
			},
		},
		tags: map[string]string{},
		commonLabels: map[string]string{
			"module_id":  "default",
			"project_id": "repcore-prod",
			"version_id": "default",
			"zone":       "us2",
		},
	}
	return context.WithValue(ctx, requestDataKey{}, rd), rd
}

func requestDataFromContext(ctx context.Context) (md *requestData, ok bool) {
	md, ok = ctx.Value(requestDataKey{}).(*requestData)
	return
}
