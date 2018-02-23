package pubsub

import (
	"errors"
	"fmt"

	"cloud.google.com/go/pubsub"

	"github.com/vendasta/gosdks/logging"
	"golang.org/x/net/context"
)

// Publish a message to a topic
func Publish(ctx context.Context, topic Topic, messages []*pubsub.Message) error {
	return publish(ctx, publishMessage, topic, messages)
}

type publisherFunc func(ctx context.Context, topic Topic, messages []*pubsub.Message) error

func publish(ctx context.Context, publisher publisherFunc, topic Topic, messages []*pubsub.Message) error {
	if topic == nil {
		logging.Errorf(ctx, "Topic is required to publish to")
		return errors.New("Topic is required to publish to")
	}
	exists, err := topic.Exists(ctx)
	if err != nil {
		logging.Errorf(ctx, "Error checking the existance of topic %s: %s", topic.String(), err.Error())
		return err
	}
	if !exists {
		message := fmt.Sprintf("Topic does not exist %s", topic.String())
		logging.Errorf(ctx, message)
		return errors.New(message)
	}
	if len(messages) == 0 {
		logging.Debugf(ctx, "No messages were provided to publish to topic %s", topic.String())
		return nil
	}
	return publisher(ctx, topic, messages)
}

func publishMessage(ctx context.Context, topic Topic, messages []*pubsub.Message) error {
	defer topic.Stop()
	var res []PublishResult
	for _, msg := range messages {
		res = append(res, topic.Publish(ctx, msg))
	}
	for _, r := range res {
		_, err := r.Get(ctx)
		if err != nil {
			logging.Errorf(ctx, "Failed to publish message on topic %s: %s", topic.String(), err.Error())
			return err
		}
	}
	return nil
}
