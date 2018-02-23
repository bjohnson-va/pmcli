package taskqueue

import (
	"context"
	"time"

	"github.com/jpillora/backoff"
	"github.com/vendasta/gosdks/logging"
)

// TaskHandler is a function that receives a cloud task message and processes it
// If no error is returned, the task is marked as acknowledged and completed.
// If an error is returned, the task is not acknowledged will remain in the task queue.
type TaskHandler func(ctx context.Context, msg Task) error

// PayloadDefinition is a JSON-tagged struct
// eg:
// type MyPayload struct {
//      field1 string `json:"field1,omitempty"`
//      field2 []string `json:"field2,omitempty"`
// }
// The JSON tags define how the payload is serialized into the task queue, as well as how it is deserialized from the queue.
type PayloadDefinition interface{}

// Worker processes tasks for a specific Task Queue
type Worker struct {
	queueID         string
	maxTasksToLease int64
	filter          string
	leaseDuration   time.Duration
	cancelFunc      context.CancelFunc
	leaseBackoff    *backoff.Backoff
	handlerBackoff  *backoff.Backoff
	handler         TaskHandler
	client          Interface
}

// Work processes tasks
func (w *Worker) Work(inCtx context.Context) {
	ctx := logging.NewWorkerContext(inCtx)
	logging.Tag(ctx, "cloudTasks/queueID", w.queueID)

	if w.cancelFunc != nil {
		defer w.cancelFunc()
	}

	for {
		select {
		case <-ctx.Done():
			return
		default:
			t, err := w.client.LeaseTask(ctx, w.queueID, w.leaseDuration, w.maxTasksToLease, w.filter)
			if err != nil {
				if err == NoTaskToPull {
					time.Sleep(time.Second)
					continue
				}
				logging.Errorf(ctx, "Error leasing tasks for queue %s: %s", w.queueID, err.Error())
				time.Sleep(w.leaseBackoff.Duration())
				continue
			}

			// if we successfully pull a task, reset the backoff mechanism
			w.leaseBackoff.Reset()

			// invoke the client's handler on the pulled task
			err = w.handler(ctx, t)
			if err == nil {
				err = w.client.AckTask(ctx, t.taskID, t.ackID)
				if err != nil {
					logging.Errorf(ctx, "Error acking task %s for queue %s: %s", t.taskID, w.queueID, err.Error())
				}
			} else {
				logging.Warningf(ctx, "Error processing task %s for queue %s: %s", t.taskID, w.queueID, err.Error())
				if w.handlerBackoff != nil {
					duration := w.handlerBackoff.ForAttempt(float64(t.attempt))
					logging.Debugf(ctx, "Backing off %ds for attempt %d", duration.Seconds(), t.attempt)
					err = w.client.RenewLease(ctx, t, duration)
					if err != nil {
						logging.Errorf(ctx, "Error renewing lease for task %s for queue %s: %s",
							t.taskID, w.queueID, err.Error())
					}
				} else {
					err = w.client.NackTask(ctx, t.taskID, t.ackID)
					if err != nil {
						logging.Errorf(ctx, "Error nacking task %s for queue %s: %s", t.taskID, w.queueID, err.Error())
					}
				}
			}
		}
	}
}

// NewWorker creates a new Worker
func NewWorker(queueID string, handler TaskHandler, client Interface, opts ...WorkerOption) *Worker {
	c := &workerConfig{
		leaseDuration: defaultLeaseDuration,
		filter:        "",
		cancelFunc:    func() {},
		maxToLease:    1,
		leaseBackoff: &backoff.Backoff{
			Factor: 2,
			Jitter: true,
			Min:    time.Millisecond * 10,
			Max:    time.Minute * 30,
		},
		handlerBackoff: nil,
	}

	for _, o := range opts {
		o(c)
	}

	worker := &Worker{
		queueID:        queueID,
		handler:        handler,
		client:         client,
		leaseBackoff:   c.leaseBackoff,
		handlerBackoff: c.handlerBackoff,
		filter:         c.filter,
		leaseDuration:  c.leaseDuration,
		cancelFunc:     c.cancelFunc,
	}
	return worker
}
