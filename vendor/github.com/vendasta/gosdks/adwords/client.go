package adwords

import (
	"fmt"

	"github.com/vendasta/gosdks/config"
	"github.com/vendasta/gosdks/pb/adwords_service/v1"
	"github.com/vendasta/gosdks/vax"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// NewClient returns a new adwords-service API client object
func NewClient(ctx context.Context, e config.Env, dialOptions ...grpc.DialOption) (Interface, error) {
	address := addresses[e]
	if address == "" {
		return nil, fmt.Errorf("Unable to create client with environment %d", e)
	}
	connection, err := vax.NewGRPCConnection(ctx, address, e != config.Local, scopes[e], true, dialOptions...)
	if err != nil {
		return nil, err
	}
	return &client{
		AccountsClient: adwords_v1.NewAccountsClient(connection),
		KeywordsClient: adwords_v1.NewKeywordsClient(connection),
		ReportsClient:  adwords_v1.NewReportsClient(connection),
	}, nil
}

type client struct {
	adwords_v1.AccountsClient
	adwords_v1.KeywordsClient
	adwords_v1.ReportsClient
}
