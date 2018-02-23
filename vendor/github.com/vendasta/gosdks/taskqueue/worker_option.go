package taskqueue

import (
	"context"
	"fmt"
	"github.com/jpillora/backoff"
	"time"
)

const defaultLeaseDuration = time.Minute

// WorkerOption can be provided to taskqueue.NewWorker()
type WorkerOption func(*workerConfig)

type workerConfig struct {
	leaseDuration  time.Duration
	filter         string
	cancelFunc     context.CancelFunc
	leaseBackoff   *backoff.Backoff
	handlerBackoff *backoff.Backoff
	maxToLease     int64
}

// WithLeaseDuration controls how long the worker leases tasks from the task queue for
func WithLeaseDuration(duration time.Duration) WorkerOption {
	return func(c *workerConfig) {
		c.leaseDuration = duration
	}
}

// WithCancelFunc defines the work done when the worker is cleaned up
func WithCancelFunc(cancelFunc context.CancelFunc) WorkerOption {
	return func(c *workerConfig) {
		c.cancelFunc = cancelFunc
	}
}

// WithBackoff defines a custom backoff configuration for the worker.
// This backoff controls how often the worker is attempting to lease tasks when there is an error with the cloud tasks API.
func WithBackoff(config *backoff.Backoff) WorkerOption {
	return func(c *workerConfig) {
		c.leaseBackoff = config
	}
}

// WithMaximumTasksPerLease controls how many tasks the worker leases at once.
func WithMaximumTasksPerLease(max int64) WorkerOption {
	return func(c *workerConfig) {
		c.maxToLease = max
	}
}

// WithTagFilter can be used to specify a subset of tasks to lease
func WithTagFilter(tag string) WorkerOption {
	return func(c *workerConfig) {
		c.filter = fmt.Sprintf("tag=%s", tag)
	}
}

// WithHandlerBackoff controls how the worker will back off when errors are returned by the task handler
func WithHandlerBackoff(bo *backoff.Backoff) WorkerOption {
	return func(c *workerConfig) {
		c.handlerBackoff = bo
	}
}
