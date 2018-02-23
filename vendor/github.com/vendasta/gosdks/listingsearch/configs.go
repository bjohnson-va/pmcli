package listingsearch

import "github.com/vendasta/gosdks/config"

type Source int64

const (
	BingPlaces Source = 10020
	ForRent    Source = 10860
)


var addresses = map[Source]map[config.Env]string {
	BingPlaces:  {
	config.Local: "domain:11000",
	config.Test:  "bing-places-api-test.vendasta-internal.com:443",
	config.Demo:  "bing-places-api-demo.vendasta-internal.com:443",
	config.Prod:  "bing-places-api-prod.vendasta-internal.com:443",
},
ForRent: {
	config.Local: "domain:11000",
	config.Test:  "api-listing-sources-api-test.vendasta-internal.com:443",
	config.Demo:  "api-listing-sources-api-demo.vendasta-internal.com:443",
	config.Prod:  "api-listing-sources-api-prod.vendasta-internal.com:443",
},
}

var scopes = map[Source]map[config.Env]string {
	BingPlaces: {
		config.Local: "",
		config.Test:  "https://bing-places-api-test.vendasta-internal.com",
		config.Demo:  "https://bing-places-api-demo.vendasta-internal.com",
		config.Prod:  "https://bing-places-api-prod.vendasta-internal.com",
	},
	ForRent: {
		config.Local: "",
		config.Test:  "https://api-listing-sources-api-test.vendasta-internal.com",
		config.Demo:  "https://api-listing-sources-api-demo.vendasta-internal.com",
		config.Prod:  "https://api-listing-sources-api-prod.vendasta-internal.com",
	},
}