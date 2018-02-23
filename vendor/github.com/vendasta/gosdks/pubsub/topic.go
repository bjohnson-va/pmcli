package pubsub

import (
	"errors"

	"cloud.google.com/go/pubsub"
	"github.com/vendasta/gosdks/logging"
	"golang.org/x/net/context"
)

// Topic describes what you can do with a pubsub topic
type Topic interface {
	Exists(ctx context.Context) (bool, error)
	String() string
	Stop()
	Publish(ctx context.Context, msg *pubsub.Message) PublishResult
}

type topic struct {
	*pubsub.Topic
}

// the Publish method of the underlying topic does not match the interface correctly so we need to override it
// just to change the signature
func (t *topic) Publish(ctx context.Context, msg *pubsub.Message) PublishResult {
	return t.Topic.Publish(ctx, msg)
}

// GetOrCreateTopic gets a pubsub topic if it exists. Otherwise the topic will be created and returned to the caller
func GetOrCreateTopic(ctx context.Context, client Client, name string) (Topic, error) {
	if name == "" {
		logging.Errorf(ctx, "A topic is required to get or create")
		return nil, errors.New("A topic is required to get or create")
	}
	t := client.Topic(name)
	exists, err := t.Exists(ctx)
	if err != nil {
		logging.Errorf(ctx, "Failed to check for the existence of topic %s: %s", t.String(), err.Error())
		return nil, err
	}

	if exists {
		return t, nil
	}

	logging.Debugf(ctx, "Creating pubsub topic: %s", name)
	topic, err := client.CreateTopic(ctx, name)
	if err != nil {
		logging.Errorf(ctx, "Failed to create topic %s: %s", t, err.Error())
		return nil, err
	}
	return topic, nil
}
