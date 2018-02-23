package reserveidsdk

import "github.com/vendasta/gosdks/config"

var envURLs = map[config.Env]string{
	config.Local: "http://localhost:8080/",
	config.Test:  "https://developers-vendasta-test.appspot.com/",
	config.Demo:  "https://developers-vendasta-demo.appspot.com/",
	config.Prod:  "https://developers-vendasta.appspot.com/",
}

func rootURLFromEnv(env config.Env) string {
	return envURLs[env]
}
