package nap

import (
	"time"

	"github.com/vendasta/gosdks/config"
	"github.com/vendasta/gosdks/vax"
	"google.golang.org/grpc/codes"
)

var addresses = map[config.Env]string{
	config.Local: "nap-api-test.vendasta-internal.com:443",
	config.Test:  "nap-api-test.vendasta-internal.com:443",
	config.Demo:  "nap-api-demo.vendasta-internal.com:443",
	config.Prod:  "nap-api-prod.vendasta-internal.com:443",
}

var scopes = map[config.Env]string{
	config.Local: "https://nap-api-test.vendasta-internal.com",
	config.Test:  "https://nap-api-test.vendasta-internal.com",
	config.Demo:  "https://nap-api-demo.vendasta-internal.com",
	config.Prod:  "https://nap-api-prod.vendasta-internal.com",
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
