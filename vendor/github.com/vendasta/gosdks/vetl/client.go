package vetl

import (
	"github.com/vendasta/gosdks/config"
	"context"
	"fmt"
	"github.com/vendasta/gosdks/vax"
	"github.com/vendasta/gosdks/pb/vetl/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
	"github.com/vendasta/gosdks/util"
)

var addresses = map[config.Env]string{
	config.Test:  "vetl-api-test.vendasta-internal.com:443",
	config.Demo:  "vetl-api-demo.vendasta-internal.com:443",
	config.Prod:  "vetl-api-prod.vendasta-internal.com:443",
}

var scopes = map[config.Env]string{
	config.Test:  "https://vetl-api-test.vendasta-internal.com",
	config.Demo:  "https://vetl-api-demo.vendasta-internal.com",
	config.Prod:  "https://vetl-api-prod.vendasta-internal.com",
}

var serviceAccounts = map[config.Env]string{
	config.Test:  "vetl-test@repcore-prod.iam.gserviceaccount.com",
	config.Demo:  "vetl-demo@repcore-prod.iam.gserviceaccount.com",
	config.Prod:  "vetl-prod@repcore-prod.iam.gserviceaccount.com",
}

// NewClient returns a new vETL API client object
func NewClient(ctx context.Context, e config.Env, dialOptions ...grpc.DialOption) (Interface, error) {
	address := addresses[e]
	if address == "" {
		return nil, fmt.Errorf("Unable to create client with environment %d", e)
	}
	connection, err := vax.NewGRPCConnection(ctx, address, e != config.Local, scopes[e], true, dialOptions...)
	if err != nil {
		return nil, err
	}
	return &client{vetl_v1.NewVETLClient(connection)}, nil
}

type client struct {
	grpcService vetl_v1.VETLClient
}

// CreateDataSource creates a data source
func (c *client) CreateDataSource(ctx context.Context, sourceID string, dataSource Source) error {
	s := &sourceOption{}
	dataSource(s)

	// validate request
	if sourceID == "" {
		return util.Error(util.InvalidArgument, "sourceID is required")
	}
	if s.source == nil {
		return util.Error(util.InvalidArgument, "A Source is required, use DataSourceFromVStoreModel")
	}

	_, err := c.grpcService.CreateDataSource(ctx, &vetl_v1.CreateDataSourceRequest{
		SourceId: sourceID,
		Source: s.source,
	}, grpc.FailFast(false))

	//ignore errors on collision
	if err != nil {
		grpcErr, ok := status.FromError(err)
		if ok {
			if grpcErr.Code() == codes.AlreadyExists {
				return nil
			}
		}
	}
	return err
}

// UpsertTransform insert or update a transform
func (c *client) UpsertTransform(ctx context.Context, parentIDs []string, ID string, transform Transform, public bool) error {
	t := &transformOption{}
	transform(t)

	// validate request
	if ID == "" {
		return util.Error(util.InvalidArgument, "ID is required")
	}
	if len(parentIDs) < 1 {
		return util.Error(util.InvalidArgument, "A Transform must have a parent")
	}
	if t.transform == nil {
		return util.Error(util.InvalidArgument, "A Transform is required, use WithKeepKeys")
	}

	_, err := c.grpcService.UpsertTransform(ctx, &vetl_v1.UpsertTransformRequest{
		ParentIds: parentIDs,
		Id: ID,
		Transform: t.transform,
		Public: public,
	}, grpc.FailFast(false))
	return err
}

// CreateSubscription to a transform's output
func (c *client) CreateSubscription(ctx context.Context, ID, parentID string, sink Sink) error {
	s := &sinkOption{}
	sink(s)

	// validate request
	if ID == "" {
		return util.Error(util.InvalidArgument, "ID is required")
	}
	if parentID == "" {
		return util.Error(util.InvalidArgument, "parentID is required")
	}
	if s.sink == nil {
		return util.Error(util.InvalidArgument, "A Sink is required, use VStoreDataSink")
	}

	_, err := c.grpcService.CreateSubscription(ctx, &vetl_v1.CreateSubscriptionRequest{
		Id: ID,
		ParentId: parentID,
		Sink: s.sink,
	}, grpc.FailFast(false))

	//ignore errors on collision
	if err != nil {
		grpcErr, ok := status.FromError(err)
		if ok {
			if grpcErr.Code() == codes.AlreadyExists {
				return nil
			}
		}
	}
	return err
}

// BackfillSubscription backfill a subscription with all data from all its sources.
// This should be triggered manually and used to migrate/transform historical data.
func (c *client) BackfillSubscription(ctx context.Context, subscriptionID string) error {
	// validate request
	if subscriptionID == "" {
		return util.Error(util.InvalidArgument, "subscriptionID is required")
	}

	_, err := c.grpcService.BackfillSubscription(ctx, &vetl_v1.BackfillSubscriptionRequest{
		Id: subscriptionID,
	}, grpc.FailFast(false))
	return err
}

// GetServiceAccount returns the vetl service account based on the given environment.
// Use this function to give the right vetl service account access to your VStore namespaces/kinds
func GetServiceAccount(env config.Env) string {
	return serviceAccounts[env]
}
