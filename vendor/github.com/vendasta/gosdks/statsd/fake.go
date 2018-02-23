package statsd

import (
	"time"
	"github.com/DataDog/datadog-go/statsd"
)

//fakeStatsD just swallows all calls and doesn't actually ship any stats. Meant to be used on local and in tests.
type fakeStatsD struct{}

func (l *fakeStatsD) Gauge(name string, value float64, tags []string, rate float64) error {
	return nil
}

func (l *fakeStatsD) Count(name string, value int64, tags []string, rate float64) error {
	return nil
}

func (l *fakeStatsD) Histogram(name string, value float64, tags []string, rate float64) error {
	return nil
}

func (l *fakeStatsD) Decr(name string, tags []string, rate float64) error {
	return nil
}

func (l *fakeStatsD) Incr(name string, tags []string, rate float64) error {
	return nil
}

func (l *fakeStatsD) Set(name string, value string, tags []string, rate float64) error {
	return nil
}

func (l *fakeStatsD) Timing(name string, value time.Duration, tags []string, rate float64) error {
	return nil
}

func (l *fakeStatsD) TimeInMilliseconds(name string, value float64, tags []string, rate float64) error {
	return nil
}

func (l *fakeStatsD) Event(event *statsd.Event) error {
	return nil
}


func (l *fakeStatsD) SetNamespace(string) {
	return
}

func (l *fakeStatsD) SetGlobalTags([]string) {
	return
}
