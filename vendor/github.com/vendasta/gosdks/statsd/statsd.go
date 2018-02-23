package statsd

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/DataDog/datadog-go/statsd"
	"github.com/vendasta/gosdks/config"
)

var client statsdInterface
var clientNotInitialized = errors.New("StatsD client has not initialized")

// Initialize must be called before any tracing is done.
func Initialize(statsNamespace string, tags []string) error {
	env := config.CurEnv().Name()
	tags = append(tags,
		fmt.Sprintf("env:%s", env),
		fmt.Sprintf("service:%s", statsNamespace),
		fmt.Sprintf("namespace:%s", config.GetGkeNamespace()),
	)

	ddAgentAddr := os.Getenv("DD_AGENT_ADDR")
	if ddAgentAddr == "" {
		ddAgentAddr = "dd-agent.default.svc.cluster.local:8125"
	}

	//use a fake client on local
	if config.IsLocal() {
		client = &fakeStatsD{}
		return nil
	}
	//use the datadog client on real environments
	c, err := statsd.New(ddAgentAddr)
	if err != nil {
		fmt.Printf("Error initializing statsd client. %s", err.Error())
		return err
	}

	client = &dataDogStatsD{
		Client: c,
	}
	client.SetNamespace(statsNamespace)
	client.SetGlobalTags(tags)

	containerInitializedEvent(statsNamespace, tags)
	return nil
}

// EventPriority is the priority for the event
type EventPriority string

const (
	// Normal is the "normal" Priority for events
	Normal EventPriority = "normal"
	// Low is the "low" Priority for events
	Low EventPriority = "low"
)

// EventAlertType is the alert type of the event
type EventAlertType string

const (
	// Info is the "info" AlertType for events
	Info EventAlertType = "info"
	// Error is the "error" AlertType for events
	Error EventAlertType = "error"
	// Warning is the "warning" AlertType for events
	Warning EventAlertType = "warning"
	// Success is the "success" AlertType for events
	Success EventAlertType = "success"
)

// An Event is an object that can be posted to your event stream.
// Mirrored from https://github.com/DataDog/datadog-go/blob/master/statsd/statsd.go#L384
type Event struct {
	// Title of the event.  Required.
	Title string
	// Text is the description of the event.  Required.
	Text string
	// Timestamp is a timestamp for the event.  If not provided, the dogstatsd
	// server will set this to the current time.
	Timestamp time.Time
	// Hostname for the event.
	Hostname string
	// AggregationKey groups this event with others of the same key.
	AggregationKey string
	// Priority of the event.  Can be statsd.Low or statsd.Normal.
	Priority EventPriority
	// SourceTypeName is a source type for the event.
	SourceTypeName string
	// AlertType can be statsd.Info, statsd.Error, statsd.Warning, or statsd.Success.
	// If absent, the default value applied by the dogstatsd server is Info.
	AlertType EventAlertType
	// Tags for the event.
	Tags []string
}

// containerOnlineEvent sends an event signifying that a container initialized
func containerInitializedEvent(namespace string, tags []string) {
	event := &Event{
		Title:          fmt.Sprintf("%s Container Online", namespace),
		Text:           "A container initialized",
		AggregationKey: "container_initialized",
		Priority:       Low,
		AlertType:      Info,
		Tags:           tags,
	}
	LogEvent(event)
}

// Gauge measures the value of a metric at a particular time.
func Gauge(name string, value float64, tags []string, rate float64) error {
	if client == nil {
		return clientNotInitialized
	}
	return client.Gauge(name, value, tags, rate)
}

// Count tracks how many times something happened per second.
func Count(name string, value int64, tags []string, rate float64) error {
	if client == nil {
		return clientNotInitialized
	}
	return client.Count(name, value, tags, rate)
}

// Histogram tracks the statistical distribution of a set of values.
func Histogram(name string, value float64, tags []string, rate float64) error {
	if client == nil {
		return clientNotInitialized
	}
	return client.Histogram(name, value, tags, rate)
}

// Decr is just Count of 1
func Decr(name string, tags []string, rate float64) error {
	if client == nil {
		return clientNotInitialized
	}
	return client.Decr(name, tags, rate)
}

// Incr is just Count of 1
func Incr(name string, tags []string, rate float64) error {
	if client == nil {
		return clientNotInitialized
	}
	return client.Incr(name, tags, rate)
}

// Set counts the number of unique elements in a group.
func Set(name string, value string, tags []string, rate float64) error {
	if client == nil {
		return clientNotInitialized
	}
	return client.Set(name, value, tags, rate)
}

// Timing sends timing information, it is an alias for TimeInMilliseconds
func Timing(name string, value time.Duration, tags []string, rate float64) error {
	if client == nil {
		return clientNotInitialized
	}
	return client.Timing(name, value, tags, rate)
}

// TimeInMilliseconds sends timing information in milliseconds.
// It is flushed by statsd with percentiles, mean and other info (https://github.com/etsy/statsd/blob/master/docs/metric_types.md#timing)
func TimeInMilliseconds(name string, value float64, tags []string, rate float64) error {
	if client == nil {
		return clientNotInitialized
	}
	return client.TimeInMilliseconds(name, value, tags, rate)
}

// LogEvent logs the event
func LogEvent(event *Event) error {
	if client == nil {
		return clientNotInitialized
	}
	return client.Event(convertEvent(event))
}

// converts our event to the underlying datadog statsd event
func convertEvent(e *Event) *statsd.Event {
	eventType := fmt.Sprintf("aggregation_key:%s", e.AggregationKey)
	ev := &statsd.Event{
		Title:          e.Title,
		Text:           e.Text,
		Timestamp:      e.Timestamp,
		Hostname:       e.Hostname,
		AggregationKey: e.AggregationKey,
		SourceTypeName: e.SourceTypeName,
		Tags:           append(e.Tags, eventType),
	}
	setEventPriority(ev, e.Priority)
	setEventAlertType(ev, e.AlertType)
	return ev
}

// sets the statsd event's priority
func setEventPriority(ev *statsd.Event, ep EventPriority) {
	// Set type with enum values directly since dogstats type is private
	switch ep {
	case Low:
		ev.Priority = statsd.Low
	case Normal:
		ev.Priority = statsd.Normal
	}
}

// sets the statsd event's alert type
func setEventAlertType(ev *statsd.Event, eat EventAlertType) {
	// Set type with enum values directly since dogstats type is private
	switch eat {
	case Info:
		ev.AlertType = statsd.Info
	case Error:
		ev.AlertType = statsd.Error
	case Warning:
		ev.AlertType = statsd.Warning
	case Success:
		ev.AlertType = statsd.Success
	}
}

type statsdInterface interface {
	// Gauge measures the value of a metric at a particular time.
	Gauge(name string, value float64, tags []string, rate float64) error

	// Count tracks how many times something happened per second.
	Count(name string, value int64, tags []string, rate float64) error

	// Histogram tracks the statistical distribution of a set of values.
	Histogram(name string, value float64, tags []string, rate float64) error

	// Decr is just Count of 1
	Decr(name string, tags []string, rate float64) error

	// Incr is just Count of 1
	Incr(name string, tags []string, rate float64) error

	// Set counts the number of unique elements in a group.
	Set(name string, value string, tags []string, rate float64) error

	// Timing sends timing information, it is an alias for TimeInMilliseconds
	Timing(name string, value time.Duration, tags []string, rate float64) error

	// TimeInMilliseconds sends timing information in milliseconds.
	// It is flushed by statsd with percentiles, mean and other info (https://github.com/etsy/statsd/blob/master/docs/metric_types.md#timing)
	TimeInMilliseconds(name string, value float64, tags []string, rate float64) error

	// Event sends an event
	Event(event *statsd.Event) error

	// SetNamespace configures the prefix for all stats being pushed by this application.
	SetNamespace(string)
	// SetGlobalTags configures the set of global tags that are attached to each increment
	SetGlobalTags([]string)
}

type dataDogStatsD struct {
	*statsd.Client
}

func (d *dataDogStatsD) SetNamespace(namespace string) {
	d.Client.Namespace = fmt.Sprintf("%s.", namespace)
}

func (d *dataDogStatsD) SetGlobalTags(tags []string) {
	d.Client.Tags = tags
}
