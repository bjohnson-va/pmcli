package pubsub

import (
	"cloud.google.com/go/pubsub"
	"github.com/stretchr/testify/mock"
	"golang.org/x/net/context"
)

// MockSubscription mocks the Subscription interface
type MockSubscription struct{ mock.Mock }

// Exists mock
func (m *MockSubscription) Exists(ctx context.Context) (bool, error) {
	call := m.Called(ctx)
	return call.Bool(0), call.Error(1)
}

// Receive mock
func (m *MockSubscription) Receive(ctx context.Context, f func(context.Context, *pubsub.Message)) error {
	call := m.Called(ctx, f)
	return call.Error(0)
}

// MockTopic mocks the  interface
type MockTopic struct{ mock.Mock }

// Exists mock
func (m *MockTopic) Exists(ctx context.Context) (bool, error) {
	call := m.Called(ctx)
	return call.Bool(0), call.Error(1)
}

// String mock
func (m *MockTopic) String() string {
	call := m.Called()
	return call.String()
}

// Stop mock
func (m *MockTopic) Stop() {
	m.Called()
}

// Publish mock
func (m *MockTopic) Publish(ctx context.Context, msg *pubsub.Message) PublishResult {
	call := m.Called(ctx, msg)
	return call.Get(0).(PublishResult)
}

// MockPublishResult mocks the interface
type MockPublishResult struct{ mock.Mock }

// Ready mock
func (m *MockPublishResult) Ready() <-chan struct{} {
	call := m.Called()
	return call.Get(0).(<-chan struct{})
}

// Get mock
func (m *MockPublishResult) Get(ctx context.Context) (serverID string, err error) {
	call := m.Called(ctx)
	return call.String(0), call.Error(1)
}
