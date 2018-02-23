package smsdk

import "github.com/vendasta/gosdks/config"

var envURLs = map[config.Env]string{
	config.Local: "http://10.200.10.1:8084",
	config.Test:  "https://socmktg-test.appspot.com",
	config.Demo:  "https://socmktg-demo.appspot.com",
	config.Prod:  "https://socmktg-prod.appspot.com",
}

func rootURLFromEnv(env config.Env) string {
	return envURLs[env]
}
