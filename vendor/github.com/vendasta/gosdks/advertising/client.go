package advertising

import (
	"fmt"

	"github.com/vendasta/gosdks/config"
	"github.com/vendasta/gosdks/pb/advertising/v1"
	"github.com/vendasta/gosdks/vax"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type client struct {
	campaignEventsClient advertising_v1.CampaignEventsClient
	adwordsClient        advertising_v1.AdwordsClient
}

// NewClient returns a new advertising API client object
func NewClient(ctx context.Context, e config.Env, dialOptions ...grpc.DialOption) (Client, error) {
	address, ok := addresses[e]
	if !ok {
		return nil, fmt.Errorf("unable to create client with environment %d", e)
	}
	connection, err := vax.NewGRPCConnection(ctx, address, e != config.Local, scopes[e], true, dialOptions...)
	if err != nil {
		return nil, err
	}
	return &client{
		campaignEventsClient: advertising_v1.NewCampaignEventsClient(connection),
		adwordsClient:        advertising_v1.NewAdwordsClient(connection),
	}, nil
}
