package accountgroup

import "github.com/vendasta/gosdks/pb/account_group/v1"

// PageSize returns a LookupOption that includes the PageSize
func PageSize(size int64) LookupOption {
	return lookupOption{
		cb: func(req *accountgroup_v1.LookupRequest) {
			req.PageSize = size
		},
	}
}

// Cursor returns a LookupOption that includes the Cursor
func Cursor(cursor string) LookupOption {
	return lookupOption{
		cb: func(req *accountgroup_v1.LookupRequest) {
			req.Cursor = cursor
		},
	}
}

// CustomerID filters a lookup by a customer ID.
func CustomerID(customerID string) LookupOption {
	return lookupOption{
		cb: func(req *accountgroup_v1.LookupRequest) {
			req.Filters.CustomerId = customerID
		},
	}
}

// PartnerID filters a lookup by a partner ID.
func PartnerID(partnerID string) LookupOption {
	return lookupOption{
		cb: func(req *accountgroup_v1.LookupRequest) {
			req.Filters.PartnerId = partnerID
		},
	}
}

// MarketID filters a lookup by a market ID
func MarketID(marketIDs ...string) LookupOption {
	return lookupOption{
		cb: func(req *accountgroup_v1.LookupRequest) {
			req.Filters.MarketIds = marketIDs
		},
	}
}

// MarketplaceAppIDs filters a lookup by one or more marketplace app ids
func MarketplaceAppIDs(mpai []string) LookupOption {
	filters := []*accountgroup_v1.LookupRequest_AccountFilter{}
	for _, id := range mpai {
		filters = append(filters, &accountgroup_v1.LookupRequest_AccountFilter{
			MarketplaceAppId: id,
			AccountStatus:    accountgroup_v1.LookupRequest_Enabled,
		})
	}
	return lookupOption{
		cb: func(req *accountgroup_v1.LookupRequest) {
			req.Filters.AccountFilters = filters
		},
	}
}
