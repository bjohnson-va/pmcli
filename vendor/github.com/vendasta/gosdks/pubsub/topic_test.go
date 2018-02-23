package pubsub

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"golang.org/x/net/context"
)

func Test_GetOrCreateTopic_ReturnsAnErrorIfNameIsEmpty(t *testing.T) {
	_, err := GetOrCreateTopic(context.Background(), nil, "")
	assert.EqualError(t, err, "A topic is required to get or create")
}

func Test_GetOrCreateTopic_ReturnsAnErrorIfTopicExistsCheckFails(t *testing.T) {
	c := newMockClient(&mockErrorTopic{}, nil)
	_, err := GetOrCreateTopic(context.Background(), c, "topic-name")
	assert.EqualError(t, err, "Failed to check if topic exists")
}

func Test_GetOrCreateTopic_ReturnsTopicIfItAlreadyExists(t *testing.T) {
	expectedTopic := &mockHappyPathTopic{}
	c := newMockClient(&mockHappyPathTopic{}, nil)
	topic, err := GetOrCreateTopic(context.Background(), c, "topic-name")
	assert.Nil(t, err)
	assert.Equal(t, expectedTopic, topic)
}

func Test_GetOrCreateTopic_ReturnsErrorIfTopicCouldntBeCreated(t *testing.T) {
	c := newMockClient(&mockNonExistingTopic{}, errors.New("Failed to create topic"))
	_, err := GetOrCreateTopic(context.Background(), c, "topic-name")
	assert.EqualError(t, err, "Failed to create topic")
}

func Test_GetOrCreateTopic_ReturnsNewTopicIfTopicDoesNotExist(t *testing.T) {
	c := newMockClient(&mockNonExistingTopic{}, nil)
	topic, err := GetOrCreateTopic(context.Background(), c, "topic-name")
	assert.Nil(t, err)
	assert.NotNil(t, topic)
}
