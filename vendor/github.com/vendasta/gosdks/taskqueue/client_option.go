package taskqueue

import "time"

// ClientOption can be provided to taskqueue.NewClient()
type ClientOption func(*clientConfig)

type clientConfig struct {
	projectID  string
	locationID string
}

// WithProjectID creates the client for a project other than repcore-prod
func WithProjectID(projectID string) ClientOption {
	return func(c *clientConfig) {
		c.projectID = projectID
	}
}

// WithLocationID creates the client for a location other than us-central1
func WithLocationID(locationID string) ClientOption {
	return func(c *clientConfig) {
		c.locationID = locationID
	}
}

type TaskSchedulingOption func(*taskScheduleConfig)

type taskScheduleConfig struct {
	scheduleTime *time.Time
	tag          string
	name         string
}

// DelayProcessingBy schedules a task to be processed some amount of time from now.
func DelayProcessingBy(delay time.Duration) TaskSchedulingOption {
	t := time.Now().UTC().Add(delay)
	return func(c *taskScheduleConfig) {
		c.scheduleTime = &t
	}
}

// ProcessAt schedules a task to be processed at a specific time
func ProcessAt(target time.Time) TaskSchedulingOption {
	return func(c *taskScheduleConfig) {
		c.scheduleTime = &target
	}
}

// WithTag schedules a task with the specified tag
func WithTag(tag string) TaskSchedulingOption {
	return func(c *taskScheduleConfig) {
		c.tag = tag
	}
}

// WithName schedules a task with the specified name
func WithName(name string) TaskSchedulingOption {
	return func(c *taskScheduleConfig) {
		c.name = name
	}
}
