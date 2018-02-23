package concierge

import (
	"github.com/vendasta/gosdks/basesdk"
)

var (
	noHTTPConciergeClient = conciergeClient{&basesdk.BaseClientMock{}}
)
