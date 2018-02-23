package group

import (
	"context"
	"fmt"

	"github.com/vendasta/gosdks/config"
	"github.com/vendasta/gosdks/pb/group/v1"
	"github.com/vendasta/gosdks/vax"
	"google.golang.org/grpc"
)

var addresses = map[config.Env]string{
	config.Test: "group-api-test.vendasta-internal.com:443",
	config.Demo: "group-api-demo.vendasta-internal.com:443",
	config.Prod: "group-api-prod.vendasta-internal.com:443",
}

var scopes = map[config.Env]string{
	config.Test: "https://group-api-test.vendasta-internal.com",
	config.Demo: "https://group-api-demo.vendasta-internal.com",
	config.Prod: "https://group-api-prod.vendasta-internal.com",
}

// NewClient returns a new Group API client object
func NewClient(ctx context.Context, e config.Env, dialOptions ...grpc.DialOption) (Interface, error) {
	address := addresses[e]
	if address == "" {
		return nil, fmt.Errorf("Unable to create client with environment %d", e)
	}
	connection, err := vax.NewGRPCConnection(ctx, address, true, scopes[e], true, dialOptions...)
	if err != nil {
		return nil, err
	}
	return &client{group_v1.NewGroupServiceClient(connection)}, nil
}

type client struct {
	grpcService group_v1.GroupServiceClient
}

// ListMembers all the values in a group, and it's children as a flat list
func (c *client) ListMembers(ctx context.Context, path Path, pagedRequestOptions ...PagedRequestOption) (members []string, totalMembers int64, nextCursor string, hasMore bool, err error) {
	if len(path.Nodes) == 0 {
		return nil, 0, "", false, ErrPathNodesRequired
	}
	pro := pagedRequestOptionsProto(pagedRequestOptions)
	res, err := c.grpcService.ListMembers(ctx, &group_v1.ListMembersRequest{
		Path:          &group_v1.Path{Nodes: path.Nodes},
		PagingOptions: pro,
	})
	return res.GetMembers(), res.GetTotalMembers(), res.GetPagingMetadata().GetNextCursor(), res.GetPagingMetadata().GetHasMore(), err
}

// List all the child groups for a given path
func (c *client) List(ctx context.Context, foreignKeys ForeignKeys, path Path, listFilters []ListFilter, pagedRequestOptions ...PagedRequestOption) (groups []*Group, nextCursor string, hasMore bool, err error) {
	if foreignKeys.PartnerID == "" {
		return nil, "", false, ErrPartnerIDRequired
	}
	pro := pagedRequestOptionsProto(pagedRequestOptions)
	lf := listFiltersProto(listFilters)
	res, err := c.grpcService.List(ctx, &group_v1.ListRequest{
		ForeignKeys:   &group_v1.ForeignKeys{PartnerId: foreignKeys.PartnerID, MarketId: foreignKeys.MarketID},
		Path:          &group_v1.Path{Nodes: path.Nodes},
		Filters:       lf,
		PagingOptions: pro,
	})
	if err != nil {
		return nil, "", false, err
	}
	groups = make([]*Group, len(res.Groups))
	for i, gp := range res.Groups {
		groups[i], err = fromProto(gp)
		if err != nil {
			return nil, "", false, err
		}
	}
	return groups, res.GetPagingMetadata().GetNextCursor(), res.GetPagingMetadata().GetHasMore(), err
}

// GetMulti groups
func (c *client) GetMulti(ctx context.Context, paths []Path) ([]*Group, error) {
	protoPaths := make([]*group_v1.Path, len(paths))
	for i, p := range paths {
		protoPaths[i] = &group_v1.Path{Nodes: p.Nodes}
	}
	res, err := c.grpcService.GetMulti(ctx, &group_v1.GetMultiRequest{
		Path: protoPaths,
	})
	if err != nil {
		return nil, err
	}
	groups := make([]*Group, len(res.Groups))
	for i, gp := range res.Groups {
		groups[i], err = fromProto(gp.GetGroup())
		if err != nil {
			return nil, err
		}
	}
	return groups, nil
}

// Get group
func (c *client) Get(ctx context.Context, path Path) (*Group, error) {
	groups, err := c.GetMulti(ctx, []Path{path})
	if err != nil {
		return nil, err
	}
	return groups[0], nil
}

func pagedRequestOptionsProto(pros []PagedRequestOption) *group_v1.PagedRequestOptions {
	proto := &group_v1.PagedRequestOptions{}
	for _, pro := range pros {
		pro(proto)
	}
	return proto
}

func listFiltersProto(lfs []ListFilter) *group_v1.ListRequest_Filters {
	proto := &group_v1.ListRequest_Filters{}
	for _, lf := range lfs {
		lf(proto)
	}
	return proto
}
