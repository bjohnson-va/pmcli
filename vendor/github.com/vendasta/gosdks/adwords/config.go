package adwords

import (
	"github.com/vendasta/gosdks/config"
	"github.com/vendasta/gosdks/vax"
	"google.golang.org/grpc/codes"
	"time"
)

var addresses = map[config.Env]string{
	config.Local: "adwords-service:11000",
	config.Test:  "adwords-service-api-test.vendasta-internal.com:443",
	config.Demo:  "adwords-service-api-demo.vendasta-internal.com:443",
	config.Prod:  "adwords-service-api-prod.vendasta-internal.com:443",
}

var scopes = map[config.Env]string{
	config.Local: "",
	config.Test:  "https://adwords-service-api-test.vendasta-internal.com",
	config.Demo:  "https://adwords-service-api-demo.vendasta-internal.com",
	config.Prod:  "https://adwords-service-api-prod.vendasta-internal.com",
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
