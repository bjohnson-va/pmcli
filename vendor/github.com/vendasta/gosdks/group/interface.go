package group

import (
	"context"

	"github.com/vendasta/gosdks/pb/group/v1"
)

// PagedRequestOption is an option related to paging
type PagedRequestOption func(pro *group_v1.PagedRequestOptions)

// Cursor option sets the cursor
func Cursor(cursor string) PagedRequestOption {
	return func(pro *group_v1.PagedRequestOptions) {
		pro.Cursor = cursor
	}
}

// PageSize option sets the page size
func PageSize(pageSize int64) PagedRequestOption {
	return func(pro *group_v1.PagedRequestOptions) {
		pro.PageSize = pageSize
	}
}

// ListFilter is an option related to filtering on the List call
type ListFilter func(lf *group_v1.ListRequest_Filters)

// MemberTypeFilter filters results to groups which match the member type
func MemberTypeFilter(memberType string) ListFilter {
	return func(lf *group_v1.ListRequest_Filters) {
		lf.MemberType = memberType
	}
}

// NamespaceFilter filters results to groups in the namespace
func NamespaceFilter(namespace string) ListFilter {
	return func(lf *group_v1.ListRequest_Filters) {
		lf.Namespace = namespace
	}
}

// Interface to group microservice API
type Interface interface {
	// Create a new group

	// Add members to a specific group

	// Remove members to a specific group

	// ListMembers all the values in a group, and it's children as a flat list
	ListMembers(ctx context.Context, path Path, pagedRequestOptions ...PagedRequestOption) (members []string, totalMembers int64, nextCursor string, hasMore bool, err error)
	// List all the child groups for a given path
	List(ctx context.Context, foreignKeys ForeignKeys, path Path, listFilters []ListFilter, pagedRequestOptions ...PagedRequestOption) (groups []*Group, nextCursor string, hasMore bool, err error)
	// Update a group

	// Delete a group

	// GetMulti groups
	GetMulti(ctx context.Context, paths []Path) ([]*Group, error)
	// Get single group
	Get(ctx context.Context, path Path) (*Group, error)
}
