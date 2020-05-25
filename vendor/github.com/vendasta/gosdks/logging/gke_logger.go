package logging

import (
	"fmt"
	"os"
	"strconv"
	"time"

	gce_metadata "cloud.google.com/go/compute/metadata"
	"cloud.google.com/go/logging"
	"github.com/mattheath/kala/bigflake"
	"github.com/vendasta/gosdks/statsd"
	"golang.org/x/net/context"
	mrpb "google.golang.org/genproto/googleapis/api/monitoredres"
	google_logging_type "google.golang.org/genproto/googleapis/logging/type"
	logpb "google.golang.org/genproto/googleapis/logging/v2"
)

type gkeLogger struct {
	appLogger     *logging.Logger
	requestLogger *logging.Logger
	config        *config
	flake         *bigflake.Bigflake
}

func (l *gkeLogger) request(ctx context.Context, rl *requestData) {
	l.logRequest(ctx, rl)
}

func (l *gkeLogger) Debugf(ctx context.Context, f string, a ...interface{}) {
	l.log(ctx, logging.Debug, f, a...)
}

func (l *gkeLogger) Infof(ctx context.Context, f string, a ...interface{}) {
	l.log(ctx, logging.Info, f, a...)
}

func (l *gkeLogger) Noticef(ctx context.Context, f string, a ...interface{}) {
	l.log(ctx, logging.Notice, f, a...)
}

func (l *gkeLogger) Warningf(ctx context.Context, f string, a ...interface{}) {
	l.log(ctx, logging.Warning, f, a...)
}

func (l *gkeLogger) Errorf(ctx context.Context, f string, a ...interface{}) {
	l.log(ctx, logging.Error, f, a...)
}

func (l *gkeLogger) Criticalf(ctx context.Context, f string, a ...interface{}) {
	l.log(ctx, logging.Critical, f, a...)
}

func (l *gkeLogger) Alertf(ctx context.Context, f string, a ...interface{}) {
	l.log(ctx, logging.Alert, f, a...)
}

func (l *gkeLogger) Emergencyf(ctx context.Context, f string, a ...interface{}) {
	l.log(ctx, logging.Emergency, f, a...)
}

func (l *gkeLogger) Tag(ctx context.Context, key, value string) Logger {
	l.tag(ctx, key, value)
	return l
}

func (l *gkeLogger) log(ctx context.Context, severity logging.Severity, f string, a ...interface{}) {
	requestData, ok := requestDataFromContext(ctx)

	//Synchronously log Emergency and Critical logs as the program is exiting soon
	if severity == logging.Emergency || severity == logging.Critical {
		l.appLogger.LogSync(ctx, logging.Entry{
			Timestamp: time.Now().UTC(),
			Severity:  severity,
			Payload:   fmt.Sprintf(f, a...),
			Operation: &logpb.LogEntryOperation{
				Producer: "gke-logger",
			},
		})
		return
	}

	if !ok {
		cd, ok := taggedDataFromContext(ctx)
		if !ok {
			cd = &tagsData{
				tags: map[string]string{},
			}
		}
		// Log message to cloud logging without bundling into an existing request log.
		l.appLogger.Log(logging.Entry{
			Timestamp: time.Now().UTC(),
			Severity:  severity,
			Labels:    cd.getLabels(),
			Payload:   fmt.Sprintf(f, a...),
			Operation: &logpb.LogEntryOperation{
				Producer: "gke-logger",
			},
		})
	} else {
		// We have an existing request associated to the provided context, add log message to existing
		// request data so we have our application logs bundled with our request log.
		requestData.logLine(fmt.Sprintf(f, a...), loggingSeverityToPB(severity))
	}
}

func (l *gkeLogger) tag(ctx context.Context, key, value string) {
	requestData, ok := requestDataFromContext(ctx)

	if ok {
		// We have an existing request associated to the provided context, add tag to existing
		// request data
		requestData.addTag(key, value)
	} else {
		// otherwise try to use context data
		// this requires the context to have been created using NewTaggedContext
		contextData, ok := taggedDataFromContext(ctx)
		if ok {
			contextData.addTag(key, value)
		}
	}
}

func loggingSeverityToPB(s logging.Severity) google_logging_type.LogSeverity {
	return google_logging_type.LogSeverity(s)
}

func getLabelsFromRequest(rd *requestData) map[string]string {
	labels := rd.getLabels()
	labels["appengine.googleapis.com/trace_id"] = rd.Entry.Trace
	return labels
}

func (l *gkeLogger) logRequest(ctx context.Context, rd *requestData) {
	sev := logging.Info
	if rd.HTTPRequest.Request.URL.Path == "/healthz" {
		sev = logging.Debug
	}
	if rd.HTTPRequest.Status >= 500 {
		sev = logging.Error
	}
	rd.Entry.Labels = getLabelsFromRequest(rd)
	rd.Entry.Severity = sev
	l.requestLogger.Log(*rd.Entry)
	for _, line := range rd.lines {
		l.appLogger.Log(logging.Entry{
			Timestamp: line.Time,
			Severity:  logging.Severity(line.Severity),
			Payload:   line.LogMessage,
			Trace:     rd.Trace,
			Resource:  rd.Entry.Resource,
		})
	}
	tags := []string{
		fmt.Sprintf("status:%d", rd.HTTPRequest.Status),
		fmt.Sprintf("namespace:%s", l.config.Namespace),
		fmt.Sprintf("path:%s", rd.HTTPRequest.Request.URL.Path),
	}
	go func() {
		statsd.Histogram("gRPC.Latency", float64(rd.HTTPRequest.Latency.Nanoseconds()/1e6), tags, 1)
	}()
}

func (l *gkeLogger) RequestID() string {
	for {
		f, err := l.flake.Mint()
		if err == bigflake.ErrSequenceOverflow {
			time.Sleep(time.Millisecond)
			continue
		} else if err != nil {
			Errorf(context.Background(), "Unable to use flake to generate a unique request id, returning empty string. Error: %s", err.Error())
			return ""
		}
		return f.String()
	}
}

func newGkeLogger(config *config, client *logging.Client) (*gkeLogger, error) {
	projectID, err := gce_metadata.ProjectID()
	if err != nil {
		return nil, err
	}
	zone, err := gce_metadata.Zone()
	if err != nil {
		return nil, err
	}
	clusterName, err := gce_metadata.InstanceAttributeValue("cluster-name")
	if err != nil {
		return nil, err
	}
	instanceID, err := gce_metadata.InstanceID()
	if err != nil {
		return nil, err
	}
	instanceIDInt, err := strconv.Atoi(instanceID)
	if err != nil {
		return nil, err
	}
	labels := map[string]string{
		"project_id":     projectID,
		"cluster_name":   clusterName,
		"namespace_id":   config.Namespace,
		"instance_id":    instanceID,
		"pod_id":         os.Getenv("HOSTNAME"),
		"container_name": config.AppName,
		"zone":           zone,
	}
	mr := &mrpb.MonitoredResource{Type: "container", Labels: labels}
	workerID := uint64(instanceIDInt) & uint64((1<<31)-1) //take 32 bits from the 64-bit integer.
	flake, _ := bigflake.New(workerID)
	client.OnError = func(err error) {
		fmt.Printf("Error flushing logs: %s", err.Error())
	}
	logger := gkeLogger{
		config:        config,
		flake:         flake,
		appLogger:     client.Logger(config.AppName, logging.CommonResource(mr)),
		requestLogger: client.Logger("request", logging.CommonResource(mr)),
	}
	return &logger, nil
}
