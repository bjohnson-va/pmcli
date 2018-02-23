package pubsub

import (
	"time"

	"golang.org/x/net/context"
)

func newMockClient(t Topic, createTopicError error) Client {
	return &mockClient{t: t, createTopicError: createTopicError}
}

type mockClient struct {
	t                Topic
	createTopicError error
}

func (c *mockClient) CreateSubscription(ctx context.Context, id string, topic Topic, ackDeadline time.Duration) (Subscription, error) {
	return nil, nil
}

func (c *mockClient) CreateTopic(ctx context.Context, id string) (Topic, error) {
	if c.createTopicError != nil {
		return nil, c.createTopicError
	}
	return c.t, nil
}

func (c *mockClient) Subscription(id string) Subscription {
	return nil
}

func (c *mockClient) Topic(id string) Topic {
	return c.t
}

func (c *mockClient) TopicInProject(id, projectID string) Topic {
	return c.t
}
