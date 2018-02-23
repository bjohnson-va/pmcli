package cssdk

import "github.com/vendasta/gosdks/config"

var envURLs = map[config.Env]string{
	config.Local: "http://10.200.10.1:8080",
	config.Test:  "http://repcore-test.appspot.com",
	config.Demo:  "http://repcore-demo.appspot.com",
	config.Prod:  "http://repcore-prod.appspot.com",
}

func rootURLFromEnv(env config.Env) string {
	return envURLs[env]
}
