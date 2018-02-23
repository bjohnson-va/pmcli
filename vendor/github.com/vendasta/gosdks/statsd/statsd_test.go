package statsd

import (
	"testing"
	"time"

	"github.com/DataDog/datadog-go/statsd"
	"github.com/stretchr/testify/assert"
)

func Test_convertEvent(t *testing.T) {
	now := time.Now()
	ev := &Event{
		Title:          "Title",
		Text:           "Text",
		Timestamp:      now,
		Hostname:       "Hostname",
		AggregationKey: "aggregation_key",
		SourceTypeName: "Source Type Name",
		Tags:           []string{"Tag1", "Tag2"},
		AlertType:      Warning,
		Priority:       Low,
	}
	e := convertEvent(ev)
	assert.Equal(t, "Title", e.Title)
	assert.Equal(t, "Text", e.Text)
	assert.Equal(t, now, e.Timestamp)
	assert.Equal(t, "Hostname", e.Hostname)
	assert.Equal(t, "aggregation_key", e.AggregationKey)
	assert.Equal(t, "Source Type Name", e.SourceTypeName)
	assert.Equal(t, []string{"Tag1", "Tag2", "aggregation_key:aggregation_key"}, e.Tags)
	assert.Equal(t, statsd.Low, e.Priority)
	assert.Equal(t, statsd.Warning, e.AlertType)
}
