package accountgroup

import (
	"github.com/vendasta/gosdks/pb/account_group/v1"
)

// IncludeDeleted indicates to Get/GetMulti to include deleted account groups in the response.
func IncludeDeleted() Option {
	return readFilterOption{
		cb: func(readFilter *accountgroup_v1.ReadFilter) {
			readFilter.IncludeDeleted = true
		},
	}
}

// IncludeNAPData indicates to Get/GetMulti/Lookup to include the nap data.
func IncludeNAPData() Option {
	return projectionFilterOption{
		cb: func(projectionFilter *accountgroup_v1.ProjectionFilter) {
			projectionFilter.NapData = true
		},
	}
}

// IncludeAccounts returns a ReadOption that includes the Accounts
func IncludeAccounts() Option {
	return projectionFilterOption{
		cb: func(projectionFilter *accountgroup_v1.ProjectionFilter) {
			projectionFilter.Accounts = true
		},
	}
}

// IncludeListingDistribution returns a ReadOption that includes the ListingDistribution
func IncludeListingDistribution() Option {
	return projectionFilterOption{
		cb: func(projectionFilter *accountgroup_v1.ProjectionFilter) {
			projectionFilter.ListingDistribution = true
		},
	}
}

// IncludeListingSyncPro returns a ReadOption that includes the ListingSyncPro
func IncludeListingSyncPro() Option {
	return projectionFilterOption{
		cb: func(projectionFilter *accountgroup_v1.ProjectionFilter) {
			projectionFilter.ListingSyncPro = true
		},
	}
}

// IncludeAssociations returns a ReadOption that includes the Associations
func IncludeAssociations() Option {
	return projectionFilterOption{
		cb: func(projectionFilter *accountgroup_v1.ProjectionFilter) {
			projectionFilter.Associations = true
		},
	}
}

// IncludeAccountGroupExternalIdentifiers returns a ReadOption that includes the AccountGroupExternalIdentifiers
func IncludeAccountGroupExternalIdentifiers() Option {
	return projectionFilterOption{
		cb: func(projectionFilter *accountgroup_v1.ProjectionFilter) {
			projectionFilter.AccountGroupExternalIdentifiers = true
		},
	}
}

// IncludeSocialUrls returns a ReadOption that includes the SocialUrls
func IncludeSocialUrls() Option {
	return projectionFilterOption{
		cb: func(projectionFilter *accountgroup_v1.ProjectionFilter) {
			projectionFilter.SocialUrls = true
		},
	}
}

// IncludeHoursOfOperation returns a ReadOption that includes the HoursOfOperation
func IncludeHoursOfOperation() Option {
	return projectionFilterOption{
		cb: func(projectionFilter *accountgroup_v1.ProjectionFilter) {
			projectionFilter.HoursOfOperation = true
		},
	}
}

// IncludeContactDetails returns a ReadOption that includes the ContactDetails
func IncludeContactDetails() Option {
	return projectionFilterOption{
		cb: func(projectionFilter *accountgroup_v1.ProjectionFilter) {
			projectionFilter.ContactDetails = true
		},
	}
}

// IncludeSnapshotReports returns a ReadOption that includes the SnapshotReports
func IncludeSnapshotReports() Option {
	return projectionFilterOption{
		cb: func(projectionFilter *accountgroup_v1.ProjectionFilter) {
			projectionFilter.SnapshotReports = true
		},
	}
}

// IncludeLegacyProductDetails returns a ReadOption that includes the LegacyProductDetails
func IncludeLegacyProductDetails() Option {
	return projectionFilterOption{
		cb: func(projectionFilter *accountgroup_v1.ProjectionFilter) {
			projectionFilter.LegacyProductDetails = true
		},
	}
}

// IncludeRichData returns a ReadOption that includes the RichData
func IncludeRichData() Option {
	return projectionFilterOption{
		cb: func(projectionFilter *accountgroup_v1.ProjectionFilter) {
			projectionFilter.RichData = true
		},
	}
}

// IncludeStatus returns a ReadOption that includes the Status
func IncludeStatus() Option {
	return projectionFilterOption{
		cb: func(projectionFilter *accountgroup_v1.ProjectionFilter) {
			projectionFilter.Status = true
		},
	}
}
