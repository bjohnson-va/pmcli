package advertising

import (
	"time"

	"github.com/vendasta/gosdks/config"
	"github.com/vendasta/gosdks/vax"
	"google.golang.org/grpc/codes"
)

var addresses = map[config.Env]string{
	config.Local: "advertising:11000",
	config.Test:  "advertising-api-test.vendasta-internal.com:443",
	config.Demo:  "advertising-api-demo.vendasta-internal.com:443",
	config.Prod:  "advertising-api-prod.vendasta-internal.com:443",
}

var scopes = map[config.Env]string{
	config.Local: "",
	config.Test:  "https://advertising-api-test.vendasta-internal.com",
	config.Demo:  "https://advertising-api-demo.vendasta-internal.com",
	config.Prod:  "https://advertising-api-prod.vendasta-internal.com",
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