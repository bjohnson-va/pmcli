package orderformsdk

import "github.com/vendasta/gosdks/config"

var envURLs = map[config.Env]string{
	config.Local: "localhost:8082/",
	config.Test: "https://vbc-test.appspot.com/",
	config.Demo: "https://vbc-demo.appspot.com/",
	config.Prod: "https://vbc-prod.appspot.com/",
}

func rootURLFromEnv(env config.Env) string {
	return envURLs[env]
}