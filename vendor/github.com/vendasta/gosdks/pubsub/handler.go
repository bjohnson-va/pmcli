package pubsub

import (
	"fmt"
	"time"

	"cloud.google.com/go/pubsub"
	gcloud_pubsub "cloud.google.com/go/pubsub"
	"github.com/vendasta/gosdks/config"
	"github.com/vendasta/gosdks/logging"
	"github.com/vendasta/gosdks/serverconfig"
	"github.com/vendasta/gosdks/statsd"
	"golang.org/x/net/context"
)

// Handler defines what's required to implement a pubsub handler
type Handler interface {
	GetTopicName() string
	GetSubscriptionName() string
	Handler(ctx context.Context, message *pubsub.Message) error
	Cancel()
}

// StartHandling starts work for the handler and returns an error if something went wrong settting that up
func StartHandling(ctx context.Context, c Client, h Handler, opts ...WorkerOption) error {
	ctx = serverconfig.SetUserOnContext(ctx, &serverconfig.UserInfo{Email: config.GetServiceAccount()})
	topicName, subscriptionName := h.GetTopicName(), h.GetSubscriptionName()
	s := fmt.Sprintf("%s-%s", topicName, subscriptionName)
	err := GetOrCreateSubscription(ctx, c, h.GetTopicName(), s)
	if err != nil {
		logging.Errorf(ctx, "Failed to get or create subscription to topic %s: %s", h.GetTopicName(), err.Error())
		return err
	}
	handler := AddPubsubHandlerMonitoring(topicName, subscriptionName, h.Handler)
	DoWork(ctx, s, c, handler, h.Cancel, opts...)
	return nil
}

// AddPubsubHandlerMonitoring adds a metric tracking pubsub handler latency
// at <namespace>.pubsub.handler.latency with the tags:
//
// topic:<topicID>
// subscription:<subscriptionID>
// either 'success' or 'error' accordingly
func AddPubsubHandlerMonitoring(topicName string, subscriptionName string, pubsubHandler MessageHandler) MessageHandler {
	return func(ctx context.Context, msg *gcloud_pubsub.Message) error {
		startTime := time.Now()
		err := pubsubHandler(ctx, msg)
		endTime := time.Now()
		latency := endTime.Sub(startTime) / time.Millisecond // Get latency in milliseconds
		result := "success"
		if err != nil {
			result = "error"
		}
		tags := []string{
			fmt.Sprintf("topic:%s", topicName),
			fmt.Sprintf("subscription:%s", subscriptionName),
			fmt.Sprintf("result:%s", result),
		}
		statsd.Histogram("pubsub.handler.latency", float64(latency), tags, 1)
		return err
	}
}
