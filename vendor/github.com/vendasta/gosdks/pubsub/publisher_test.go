package pubsub

import (
	"testing"

	"cloud.google.com/go/pubsub"

	"github.com/stretchr/testify/assert"

	"golang.org/x/net/context"
)

func Test_Publish_ReturnsErrorIfTopicIsNil(t *testing.T) {
	err := Publish(context.Background(), nil, nil)
	assert.EqualError(t, err, "Topic is required to publish to")
}

func Test_Publish_ReturnsErrorIfTopicExistCheckFails(t *testing.T) {
	topic := &mockErrorTopic{}
	err := Publish(context.Background(), topic, nil)
	assert.EqualError(t, err, "Failed to check if topic exists")
}

func Test_Publish_ReturnsErrorIfTopicDoesNotExist(t *testing.T) {
	topic := &mockNonExistingTopic{}
	err := Publish(context.Background(), topic, nil)
	assert.EqualError(t, err, "Topic does not exist topic-name")
}

func Test_Publish_ReturnsNoErrorIfNoMessagesAreProvided(t *testing.T) {
	topic := &mockHappyPathTopic{}
	err := Publish(context.Background(), topic, nil)
	assert.Nil(t, err)
}

func Test_Publish_ReturnsNoErrorIfMessageIsPublished(t *testing.T) {
	topic := &mockHappyPathTopic{}
	m := []*pubsub.Message{&pubsub.Message{}}
	err := publish(context.Background(), happyPathPublisherMock, topic, m)
	assert.Nil(t, err)
}

func Test_Publish_ReturnsErrorIfPublishFails(t *testing.T) {
	topic := &mockHappyPathTopic{}
	m := []*pubsub.Message{&pubsub.Message{}}
	err := publish(context.Background(), errorPathPublisherMock, topic, m)
	assert.EqualError(t, err, "Failed to publish")
}
