package pubsub

import (
	"sync"
	"testing"
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

//mockPubsubClient implements pubsubClient
type mockPubsubClient struct {
	mock.Mock
}

func (p *mockPubsubClient) Subscription(id string) Subscription {
	args := p.Called(id)
	res := args.Get(0)
	return res.(*mockPubsubSubscription)
}

func (p *mockPubsubClient) Topic(id string) Topic {
	args := p.Called(id)
	res := args.Get(0)
	return res.(*mockPubsubTopic)
}

func (p *mockPubsubClient) TopicInProject(id, projectID string) Topic {
	args := p.Called(id)
	res := args.Get(0)
	return res.(*mockPubsubTopic)
}

func (p *mockPubsubClient) CreateSubscription(ctx context.Context, id string, topic Topic, ackDeadline time.Duration) (Subscription, error) {
	args := p.Called(ctx, id, topic, ackDeadline)
	sub := args.Get(0)
	if sub == nil {
		return nil, args.Error(1)
	}
	return sub.(*mockPubsubSubscription), args.Error(1)
}

func (p *mockPubsubClient) CreateTopic(ctx context.Context, id string) (Topic, error) {
	return nil, nil
}

//mockPubsubSubscription implements pubsubSubscription
type mockPubsubSubscription struct {
	mock.Mock
}

func (p *mockPubsubSubscription) Exists(ctx context.Context) (bool, error) {
	args := p.Called(ctx)
	return args.Bool(0), args.Error(1)
}

type mockPubsubTopic struct {
	mock.Mock
}

//mockPubsubTopic implements pubsubTopic
func (p *mockPubsubSubscription) Receive(ctx context.Context, f func(context.Context, *pubsub.Message)) error {
	args := p.Called(ctx, f)
	return args.Error(0)
}

func (p *mockPubsubTopic) Exists(ctx context.Context) (bool, error) {
	args := p.Called(ctx)
	return args.Bool(0), args.Error(1)
}

func (p *mockPubsubTopic) Stop() {}

func (p *mockPubsubTopic) Publish(ctx context.Context, msg *pubsub.Message) PublishResult {
	return nil
}

func (p *mockPubsubTopic) String() string {
	return "name"
}

func TestNewWorkerUsesExpectedDefaultValues(t *testing.T) {
	p := &mockPubsubClient{}
	w := newWorker(context.Background(), "subscription", p, func(ctx context.Context, message *pubsub.Message) error {
		return nil
	}, func() {}, 0, time.Second)
	assert.Equal(t, int64(10), w.prefetch)
	assert.Equal(t, time.Second, w.ackExtension)
	assert.Equal(t, float64(2), w.backoff.Factor)
	assert.Equal(t, true, w.backoff.Jitter)
	assert.Equal(t, time.Millisecond*10, w.backoff.Min)
	assert.Equal(t, time.Minute*30, w.backoff.Max)
}

func TestCancelFuncCalledOnExit(t *testing.T) {
	p := &mockPubsubClient{}

	called := false
	cancelFunc := func() {
		called = true
	}

	var wg sync.WaitGroup
	wg.Add(1)

	w := newWorker(context.Background(), "subscription", p, func(ctx context.Context, message *pubsub.Message) error {
		return nil
	}, cancelFunc, 0, time.Second)

	s := &mockPubsubSubscription{}
	s.On("Receive", mock.Anything, mock.Anything, mock.Anything).Return(grpc.Errorf(codes.Canceled, "Cancelled"))
	p.On("Subscription", w.subscriptionName).Return(s)

	go func() {
		defer wg.Done()
		w.Work()
	}()

	wg.Wait()

	assert.True(t, called)
}
