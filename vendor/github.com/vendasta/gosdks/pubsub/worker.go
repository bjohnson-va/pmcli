package pubsub

import (
	"time"

	gcloud_pubsub "cloud.google.com/go/pubsub"
	"github.com/jpillora/backoff"
	"github.com/vendasta/gosdks/logging"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

func newWorker(ctx context.Context, subscriptionName string, client Client, handler MessageHandler, cancelFunc context.CancelFunc, prefetch int64, ackExtension time.Duration) *pubsubWorker {
	if prefetch == 0 {
		prefetch = 10
	}
	b := &backoff.Backoff{
		Factor: 2,
		Jitter: true,
		Min:    time.Millisecond * 10,
		Max:    time.Minute * 30,
	}
	return &pubsubWorker{subscriptionName: subscriptionName, ctx: ctx, client: client, handler: handler, cancelFunc: cancelFunc, backoff: b, prefetch: prefetch, ackExtension: ackExtension}
}

type pubsubWorker struct {
	subscriptionName string
	prefetch         int64
	ackExtension     time.Duration
	backoff          *backoff.Backoff
	client           Client
	ctx              context.Context
	cancelFunc       context.CancelFunc
	handler          MessageHandler
}

//Work will begin polling a pubsub subscription for updates, processing messages using the handler function provided.
//Polling does not stop unless the worker's context is cancelled, but errors will cause exponential backoff on polling
//retries up to a maximum of 30 minutes.
func (w *pubsubWorker) Work() {
	defer w.cancelFunc()
	for {
		// TODO: Receive settings
		err := w.client.Subscription(w.subscriptionName).Receive(w.ctx, func(ctx context.Context, msg *gcloud_pubsub.Message) {
			err := w.handler(w.ctx, msg)
			if err != nil {
				logging.Infof(w.ctx, "Handler returned an error, not marking message as done: %s", err.Error())
				msg.Nack()
				return
			}
			msg.Ack()
		})
		if err != nil && grpc.Code(err) == codes.Canceled {
			//the server has cancelled the worker (likely kubernetes shutdown)
			return
		} else if err != nil {
			logging.Errorf(w.ctx, "Error polling for messages on subscription %s: %s", w.subscriptionName, err.Error())
		}
		time.Sleep(w.backoff.Duration())
	}

}

//WorkerOption is an optional argument that parameterizes the behaviour of the pubsub workers
type WorkerOption interface {
	setOptions(o *workerOptions)
}

type workerOptions struct {
	// maxExtension is the maximum period for which the worker should automatically extend the ack deadline for each message.
	maxExtension time.Duration

	// maxPrefetch is the maximum number of messages that each worker should attempt to pull and process
	maxPrefetch int64

	// numWorkers is the number of pubsub workers to spin up concurrently
	numWorkers int64
}

type maxPrefetch int64

func (max maxPrefetch) setOptions(o *workerOptions) {
	if o.maxPrefetch = int64(max); o.maxPrefetch < 1 {
		o.maxPrefetch = 10
	}
}

//MaxPrefetch is the maximum number of messages that each worker should attempt to pull and process
func MaxPrefetch(num int64) WorkerOption {
	return maxPrefetch(num)
}

type numWorkers int64

func (num numWorkers) setOptions(o *workerOptions) {
	if o.numWorkers = int64(num); o.numWorkers < 1 {
		o.numWorkers = 1
	}
}

//NumWorkers is the number of pubsub workers to spin up concurrently
func NumWorkers(num int64) WorkerOption {
	return numWorkers(num)
}

type maxExtension time.Duration

func (max maxExtension) setOptions(o *workerOptions) {
	if o.maxExtension = time.Duration(max); o.maxExtension < 0 {
		o.maxExtension = 0
	}
}

//MaxExtension is the maximum period for which the worker should automatically extend the ack deadline for each message.
func MaxExtension(duration time.Duration) WorkerOption {
	return maxExtension(duration)
}

// MessageHandler is the type of a pubsub message handler
type MessageHandler func(ctx context.Context, msg *gcloud_pubsub.Message) error

// DoWork spins up workers to process subscriptions using the give client and handler according to provided WorkerOptions.
func DoWork(ctx context.Context, subscriptionName string, client Client, handler MessageHandler, cancelFunc context.CancelFunc, opts ...WorkerOption) {
	po := &workerOptions{
		maxExtension: time.Minute * 1,
		maxPrefetch:  10,
		numWorkers:   1,
	}

	for _, o := range opts {
		o.setOptions(po)
	}

	for n := int64(0); n < po.numWorkers; n++ {
		w := newWorker(ctx, subscriptionName, client, handler, cancelFunc, po.maxPrefetch, po.maxExtension)
		go w.Work()
	}
}
