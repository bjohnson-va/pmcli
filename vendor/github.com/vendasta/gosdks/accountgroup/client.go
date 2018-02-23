package accountgroup

import (
	"fmt"
	"time"

	"github.com/vendasta/gosdks/config"
	"github.com/vendasta/gosdks/pb/account_group/v1"
	"github.com/vendasta/gosdks/util"
	"github.com/vendasta/gosdks/vax"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

var addresses = map[config.Env]string{
	config.Local: "account-group:11000",
	config.Test:  "account-group-api-test.vendasta-internal.com:443",
	config.Demo:  "account-group-api-demo.vendasta-internal.com:443",
	config.Prod:  "account-group-api-prod.vendasta-internal.com:443",
}

var scopes = map[config.Env]string{
	config.Local: "",
	config.Test:  "https://account-group-api-test.vendasta-internal.com",
	config.Demo:  "https://account-group-api-demo.vendasta-internal.com",
	config.Prod:  "https://account-group-api-prod.vendasta-internal.com",
}

// Interface to account group microservice API
type Interface interface {
	GetMulti(ctx context.Context, accountGroupIDs []string, ro ...Option) ([]*AccountGroup, error)
	Get(ctx context.Context, accountGroupID string, ro ...Option) (*AccountGroup, error)
	Lookup(ctx context.Context, lo ...Option) (*PagedResult, error)

	Create(ctx context.Context, socialProfileGroupID string, napData *NAPData, externalIds *ExternalIdentifiers, opts ...MutateOption) (string, error)
	Update(ctx context.Context, accountGroupID string, opts ...MutateOption) error
	Delete(ctx context.Context, accountGroupID string) error
}

// NewClient returns a new account group API client object
func NewClient(ctx context.Context, e config.Env, dialOptions ...grpc.DialOption) (Interface, error) {
	address := addresses[e]
	if address == "" {
		return nil, fmt.Errorf("Unable to create client with environment %d", e)
	}
	connection, err := vax.NewGRPCConnection(ctx, address, e != config.Local, scopes[e], true, dialOptions...)
	if err != nil {
		return nil, err
	}
	return &client{accountgroup_v1.NewAccountGroupServiceClient(connection)}, nil
}

type client struct {
	accountgroup_v1.AccountGroupServiceClient
}

// GetMulti returns a set of account groups by their IDs. The list of account groups returned will match the order of
// account group ids passed in.
//
// ReadOption allows you to control which parts of the account group that you want to be returned.
//
// Example that only pulls back the account group's NAP data:
// c.GetMulti(ctx, accountGroupIds, accountgroup.IncludeNAPData(), accountgroup.IncludeDeleted())
//
func (c *client) GetMulti(ctx context.Context, accountGroupIDs []string, ro ...Option) ([]*AccountGroup, error) {
	var resp *accountgroup_v1.GetMultiResponse
	req := &accountgroup_v1.GetMultiRequest{
		AccountGroupIds:  accountGroupIDs,
		ProjectionFilter: &accountgroup_v1.ProjectionFilter{},
		ReadFilter:       &accountgroup_v1.ReadFilter{},
	}
	for _, opt := range ro {
		switch c := opt.(type) {
		case ProjectionFilterOption:
			c.projectionFilter(req.ProjectionFilter)
		case ReadFilterOption:
			c.readFilter(req.ReadFilter)
		default:
			return nil, fmt.Errorf("not able to use option %#v", c)
		}
	}
	err := vax.Invoke(ctx, func(ctx context.Context, cs vax.CallSettings) error {
		var err error
		resp, err = c.AccountGroupServiceClient.GetMulti(ctx,
			req,
			grpc.FailFast(false),
		)
		return err
	}, defaultRetryCallOptions)

	if err != nil {
		return nil, err
	}
	accountGroups := make([]*AccountGroup, len(accountGroupIDs))
	for n, accountGroup := range resp.AccountGroups {
		accountGroups[n], err = FromPB(accountGroup.GetAccountGroup(), req.ProjectionFilter)
		if err != nil {
			return nil, err
		}
	}

	return accountGroups, nil
}

// Get is a wrapper for retrieving a single account group
// Will return the account group or nil
func (c *client) Get(ctx context.Context, accountGroupID string, ro ...Option) (*AccountGroup, error) {
	ags, err := c.GetMulti(ctx, []string{accountGroupID}, ro...)
	if err != nil {
		return nil, err
	}
	if ags[0] == nil {
		return nil, util.Error(util.NotFound, "No Business found for id %s", accountGroupID)
	}
	return ags[0], nil
}

// Lookup queries account groups across a partner id.
func (c *client) Lookup(ctx context.Context, opts ...Option) (*PagedResult, error) {
	req := &accountgroup_v1.LookupRequest{
		ProjectionFilter: &accountgroup_v1.ProjectionFilter{},
		Filters:          &accountgroup_v1.LookupRequest_Filters{},
	}
	for _, opt := range opts {
		switch c := opt.(type) {
		case ProjectionFilterOption:
			c.projectionFilter(req.ProjectionFilter)
		case LookupOption:
			c.lookupOption(req)
		default:
			return nil, fmt.Errorf("not able to use option %#v", c)
		}
	}

	var resp *accountgroup_v1.PagedResponse
	err := vax.Invoke(ctx, func(ctx context.Context, cs vax.CallSettings) error {
		var err error
		resp, err = c.AccountGroupServiceClient.Lookup(ctx, req)
		return err
	}, defaultRetryCallOptions)

	if err != nil {
		return nil, err
	}
	accountGroups := make([]*AccountGroup, len(resp.AccountGroups))
	for n, accountGroup := range resp.AccountGroups {
		accountGroups[n], err = FromPB(accountGroup, req.ProjectionFilter)
		if err != nil {
			return nil, err
		}
	}
	return &PagedResult{
		AccountGroups: accountGroups,
		Cursor:        resp.NextCursor,
		HasMore:       resp.HasMore,
		TotalResults:  resp.TotalResults,
	}, nil
}

// MutateOption is used for create/update account group operations.
type MutateOption interface {
	toUpdateOperation() (*accountgroup_v1.UpdateOperation, error)
}

// Create will create a new account group and return its account group id
func (c *client) Create(ctx context.Context, socialProfileGroupID string, napData *NAPData, externalIds *ExternalIdentifiers, opts ...MutateOption) (string, error) {
	if napData == nil || externalIds == nil {
		return "", ErrInvalidCreate
	}
	var err error
	updateOps := make([]*accountgroup_v1.UpdateOperation, len(opts)+1)
	updateOps[0], err = externalIds.toUpdateOperation()
	if err != nil {
		return "", err
	}
	for n, opt := range opts {
		updateOps[n+1], err = opt.toUpdateOperation()
		if err != nil {
			return "", err
		}
	}

	resp, err := c.AccountGroupServiceClient.Create(ctx, &accountgroup_v1.CreateAccountGroupRequest{
		AccountGroupNap:      napData.toProto(),
		SocialProfileGroupId: socialProfileGroupID,
		UpdateOperations:     updateOps,
	})
	if err != nil {
		return "", err
	}
	return resp.AccountGroupId, nil
}

// Update will update an existing account group.
func (c *client) Update(ctx context.Context, accountGroupID string, opts ...MutateOption) error {
	var err error
	updateOps := make([]*accountgroup_v1.UpdateOperation, len(opts))
	for n, opt := range opts {
		updateOps[n], err = opt.toUpdateOperation()
		if err != nil {
			return err
		}
	}

	_, err = c.AccountGroupServiceClient.BulkUpdate(ctx, &accountgroup_v1.BulkUpdateRequest{
		AccountGroupId:   accountGroupID,
		UpdateOperations: updateOps,
	})
	if err != nil {
		return err
	}
	return nil
}

// Delete will mark an existing account group as deleted
func (c *client) Delete(ctx context.Context, accountGroupID string) error {
	deleteRequest := &accountgroup_v1.DeleteAccountGroupRequest{
		AccountGroupId: accountGroupID,
	}
	_, err := c.AccountGroupServiceClient.Delete(ctx, deleteRequest)
	if err != nil {
		return err
	}
	return nil
}

var defaultRetryCallOptions = vax.WithRetry(func() vax.Retryer {
	return vax.OnCodes([]codes.Code{
		codes.DeadlineExceeded,
		codes.Unavailable,
		codes.Unknown,
	}, vax.Backoff{
		Initial:    10 * time.Millisecond,
		Max:        300 * time.Millisecond,
		Multiplier: 3,
	})
})

// PagedResult is a container for paged set of account group results
type PagedResult struct {
	AccountGroups []*AccountGroup
	Cursor        string
	HasMore       bool
	TotalResults  int64
}
