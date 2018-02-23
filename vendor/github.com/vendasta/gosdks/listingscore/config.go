package listingscore

import (
	"github.com/vendasta/gosdks/config"
	"github.com/vendasta/gosdks/vax"
	"google.golang.org/grpc/codes"
	"time"
)

var addresses = map[config.Env]string{
	config.Local: "listing-score:11000",
	config.Test:  "listing-score-api-test.vendasta-internal.com:443",
	config.Demo:  "listing-score-api-demo.vendasta-internal.com:443",
	config.Prod:  "listing-score-api-prod.vendasta-internal.com:443",
}

var scopes = map[config.Env]string{
	config.Local: "",
	config.Test:  "https://listing-score-api-test.vendasta-internal.com",
	config.Demo:  "https://listing-score-api-demo.vendasta-internal.com",
	config.Prod:  "https://listing-score-api-prod.vendasta-internal.com",
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
