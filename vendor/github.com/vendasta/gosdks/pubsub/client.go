package pubsub

import (
	"errors"
	"time"

	"golang.org/x/net/context"

	"cloud.google.com/go/pubsub"
	"github.com/vendasta/gosdks/logging"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
)

//Client defines what is required to be a pubsub client
type Client interface {
	CreateSubscription(ctx context.Context, id string, topic Topic, ackDeadline time.Duration) (Subscription, error)
	CreateTopic(ctx context.Context, id string) (Topic, error)
	Subscription(id string) Subscription
	Topic(id string) Topic
	TopicInProject(id, projectID string) Topic
}

//GooglePubsubClient implements PubsubSubscription for
type GooglePubsubClient struct {
	client *pubsub.Client
}

//Subscription creates a reference to the subscription
func (p *GooglePubsubClient) Subscription(id string) Subscription {
	return p.client.Subscription(id)
}

// Topic creates a reference to a topic.
func (p *GooglePubsubClient) Topic(id string) Topic {
	return &topic{p.client.Topic(id)}
}

// TopicInProject creates a reference to a topic.
func (p *GooglePubsubClient) TopicInProject(id, projectID string) Topic {
	return &topic{p.client.TopicInProject(id, projectID)}
}

// CreateSubscription creates a new subscription on a topic.
func (p *GooglePubsubClient) CreateSubscription(ctx context.Context, id string, top Topic, ackDeadline time.Duration) (Subscription, error) {
	t, ok := top.(*topic)
	if !ok {
		return nil, errors.New("Provided topic was not a valid Topic (see cloud.google.com/go/pubsub/topic)")
	}
	return p.client.CreateSubscription(ctx, id, pubsub.SubscriptionConfig{Topic: t.Topic, AckDeadline: ackDeadline})
}

// CreateTopic creates a new pubsub topic
func (p *GooglePubsubClient) CreateTopic(ctx context.Context, id string) (Topic, error) {
	top, err := p.client.CreateTopic(ctx, id)
	if err != nil {
		return nil, err
	}
	return &topic{top}, nil
}

//NewGooglePubsubClient creates a new google pubsub client
func NewGooglePubsubClient(ctx context.Context) (*GooglePubsubClient, error) {
	logging.Debugf(ctx, "Initializing pubsub client...")
	dts, err := google.DefaultTokenSource(ctx, pubsub.ScopePubSub)
	if err != nil {
		return nil, err
	}
	c, err := pubsub.NewClient(ctx, "repcore-prod", option.WithTokenSource(dts))
	if err != nil {
		return nil, err
	}
	return &GooglePubsubClient{c}, nil
}
