package mssdk

import "github.com/vendasta/gosdks/config"

var envURLs = map[config.Env]string{
	config.Local: "http://10.200.10.1:8084",
	config.Test:  "https://microsite-test.appspot.com",
	config.Demo:  "https://microsite-demo.appspot.com",
	config.Prod:  "https://microsite-prod.appspot.com",
}

func rootURLFromEnv(env config.Env) string {
	return envURLs[env]
}
