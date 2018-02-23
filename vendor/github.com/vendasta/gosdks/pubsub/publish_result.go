package pubsub

import (
	"golang.org/x/net/context"
)

// PublishResult describes what you can do with the result of a pubsub Publish
type PublishResult interface {
	Ready() <-chan struct{}
	Get(ctx context.Context) (serverID string, err error)
}
