package pubsub

import (
	"fmt"
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/vendasta/gosdks/logging"
	"golang.org/x/net/context"
)

type subscriptionOptions struct {
	projectID string
}

// SubscriptionOptions are different options which can be specified for a subscription
type SubscriptionOptions func(o *subscriptionOptions)

// ProjectID allows for a topic in another project to be subscribed to
func ProjectID(projectID string) SubscriptionOptions {
	return func(o *subscriptionOptions) {
		o.projectID = projectID
	}
}

//GetOrCreateSubscription checks if the subscription exists and returns it if it does or creates and returns it if it doesn't exist
func GetOrCreateSubscription(ctx context.Context, client Client, topicName, subscriptionName string, options ...SubscriptionOptions) error {
	opts := &subscriptionOptions{
		projectID: "",
	}

	for _, o := range options {
		o(opts)
	}

	logging.Debugf(ctx, "Trying to get subscription %s", subscriptionName)
	exists, err := client.Subscription(subscriptionName).Exists(ctx)
	if err != nil {
		logging.Errorf(ctx, "Error checking if subscription exists: %s", subscriptionName)
		return err
	}

	if !exists {
		logging.Debugf(ctx, "Subscription %s does not exist. Creating it for topic %s and project id: %s", subscriptionName, topicName, opts.projectID)
		err := createSubscription(ctx, client, topicName, subscriptionName, opts)
		if err != nil {
			return err
		}
	} else {
		logging.Debugf(ctx, "Subscription %s already exists for topic %s.", subscriptionName, topicName)
	}
	return nil
}

func createSubscription(ctx context.Context, client Client, topicName, subscriptionName string, options *subscriptionOptions) error {
	var topic Topic
	if options.projectID == "" {
		topic = client.Topic(topicName)
	} else {
		topic = client.TopicInProject(topicName, options.projectID)
	}
	exists, err := topic.Exists(ctx)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("Topic %s does not exist", topicName)
	}

	_, err = client.CreateSubscription(ctx, subscriptionName, topic, time.Second*60)
	if err != nil {
		return err
	}
	logging.Debugf(ctx, "Created subscription %s to topic %s.", subscriptionName, topicName)
	return nil
}

//Subscription describes what you can do with a pubsub subscription
type Subscription interface {
	Exists(ctx context.Context) (bool, error)
	Receive(ctx context.Context, f func(context.Context, *pubsub.Message)) error
}
