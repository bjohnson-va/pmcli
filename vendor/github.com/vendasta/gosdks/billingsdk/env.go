package billingsdk

import "github.com/vendasta/gosdks/config"

var envURLs = map[config.Env]string{
	config.Local: "http://10.200.10.1:8088",
	config.Test:  "http://vbilling-test.appspot.com",
	config.Demo:  "http://vbilling-demo.appspot.com",
	config.Prod:  "http://vbilling-prod.appspot.com",
}

func rootURLFromEnv(env config.Env) string {
	return envURLs[env]
}
