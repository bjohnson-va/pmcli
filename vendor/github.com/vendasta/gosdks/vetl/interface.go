package vetl

import (
	"context"
)

// Interface to vETL microservice API
type Interface interface {
	CreateDataSource(ctx context.Context, sourceID string, dataSource Source) error
	UpsertTransform(ctx context.Context, parentIDs []string, ID string, transform Transform, public bool) error
	CreateSubscription(ctx context.Context, ID string, parentID string, sink Sink) error
	BackfillSubscription(ctx context.Context, subscriptionID string) error
}
