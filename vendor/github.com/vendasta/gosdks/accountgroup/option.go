package accountgroup

import "github.com/vendasta/gosdks/pb/account_group/v1"

// Option is "nearly empty" interface implmented by options
type Option interface {
	option()
}

// ProjectionFilterOption allows for projection filters to be set on outgoing requests
type ProjectionFilterOption interface {
	Option
	projectionFilter(filter *accountgroup_v1.ProjectionFilter)
}

type projectionFilterOption struct {
	cb func(projectionFilter *accountgroup_v1.ProjectionFilter)
}

// option implements the Option interface
func (pfo projectionFilterOption) option() {}

// projectionFilter implements the ProjectionFilterOption
func (pfo projectionFilterOption) projectionFilter(filter *accountgroup_v1.ProjectionFilter) {
	pfo.cb(filter)
}

// ReadFilterOption allows for read filters to be set on outgoing requests
type ReadFilterOption interface {
	Option
	readFilter(filter *accountgroup_v1.ReadFilter)
}

type readFilterOption struct {
	cb func(projectionFilter *accountgroup_v1.ReadFilter)
}

// option implements the Option interface
func (pfo readFilterOption) option() {}

// projectionFilter implements the ProjectionFilterOption
func (pfo readFilterOption) readFilter(filter *accountgroup_v1.ReadFilter) {
	pfo.cb(filter)
}

// projectionFilter implements the ProjectionFilterOption
func (pfo readFilterOption) projectionFilter(filter *accountgroup_v1.ReadFilter) {
	pfo.cb(filter)
}

// LookupOption allows for cursor and pageSize (lookup specific) options to be set on
// outgoing lookup requests
type LookupOption interface {
	Option
	lookupOption(filter *accountgroup_v1.LookupRequest)
}

type lookupOption struct {
	cb func(lookupRequest *accountgroup_v1.LookupRequest)
}

//option implements the Option interface
func (co lookupOption) option() {}

//lookupOption implements the LookupOption interface
func (co lookupOption) lookupOption(req *accountgroup_v1.LookupRequest) {
	co.cb(req)
}
