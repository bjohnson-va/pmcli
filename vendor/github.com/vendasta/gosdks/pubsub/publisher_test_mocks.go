package pubsub

import (
	"errors"

	"golang.org/x/net/context"

	"cloud.google.com/go/pubsub"
)

type mockErrorTopic struct{}

func (t *mockErrorTopic) Exists(ctx context.Context) (bool, error) {
	return false, errors.New("Failed to check if topic exists")
}

func (t *mockErrorTopic) String() string {
	return "topic-name"
}

func (t *mockErrorTopic) Stop() {}

func (t *mockErrorTopic) Publish(ctx context.Context, msg *pubsub.Message) PublishResult {
	return nil
}

type mockNonExistingTopic struct{}

func (t *mockNonExistingTopic) Exists(ctx context.Context) (bool, error) {
	return false, nil
}

func (t *mockNonExistingTopic) String() string {
	return "topic-name"
}

func (t *mockNonExistingTopic) Stop() {}

func (t *mockNonExistingTopic) Publish(ctx context.Context, msg *pubsub.Message) PublishResult {
	return nil
}

type mockBadPublishTopic struct{}

func (t *mockBadPublishTopic) Exists(ctx context.Context) (bool, error) {
	return true, nil
}

func (t *mockBadPublishTopic) String() string {
	return "topic-name"
}

func (t *mockBadPublishTopic) Stop() {}

func (t *mockBadPublishTopic) Publish(ctx context.Context, msg *pubsub.Message) *pubsub.PublishResult {
	return nil
}

type mockHappyPathTopic struct{}

func (t *mockHappyPathTopic) Exists(ctx context.Context) (bool, error) {
	return true, nil
}

func (t *mockHappyPathTopic) String() string {
	return "topic-name"
}

func (t *mockHappyPathTopic) Stop() {}

func (t *mockHappyPathTopic) Publish(ctx context.Context, msg *pubsub.Message) PublishResult {
	return nil
}

func happyPathPublisherMock(ctx context.Context, topic Topic, messages []*pubsub.Message) error {
	return nil
}

func errorPathPublisherMock(ctx context.Context, topic Topic, messages []*pubsub.Message) error {
	return errors.New("Failed to publish")
}
