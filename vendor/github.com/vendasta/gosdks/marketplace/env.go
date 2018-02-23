package marketplace

import "github.com/vendasta/gosdks/config"

var envURLs = map[config.Env]string{
	config.Local: "http://localhost:8096",
	config.Test:  "https://marketplace-proxy-test.vendasta-internal.com",
	config.Demo:  "https://marketplace-proxy-demo.vendasta-internal.com",
	config.Prod:  "https://developers.vendasta.com",
}

func rootURLFromEnv(env config.Env) string {
	return envURLs[env]
}
