package accounts

import "github.com/vendasta/gosdks/config"

var addresses = map[config.Env]string{
	config.Local: "accounts:11000",
	config.Test:  "accounts-api-test.vendasta-internal.com:443",
	config.Demo:  "accounts-api-demo.vendasta-internal.com:443",
	config.Prod:  "accounts-api-prod.vendasta-internal.com:443",
}

var scopes = map[config.Env]string{
	config.Local: "",
	config.Test:  "https://accounts-api-test.vendasta-internal.com",
	config.Demo:  "https://accounts-api-demo.vendasta-internal.com",
	config.Prod:  "https://accounts-api-prod.vendasta-internal.com",
}
