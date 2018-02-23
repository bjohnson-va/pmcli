# vStatic
- [Introduction](https://github.com/vendasta/gosdks/tree/master/vstatic#installation)
- [Usage](https://github.com/vendasta/gosdks/tree/master/vstatic#usage)
  - [Introduction](https://github.com/vendasta/gosdks/tree/master/vstatic#introduction)

## Installation

* YOLO development? `go get github.com/vendasta/gosdks`


### Introduction

vStatic serves static assets from vstore

## Usage

In main.go
```golang
c, err := vstatic.NewClient(ctx, "medical-compare-client", config.CurEnv(), grpc.WithUnaryInterceptor(logging.ClientInterceptor()))
    if err != nil {
        logging.Criticalf(ctx, "Error creating vstatic client: %s", err.Error())
        os.Exit(-1)
	}
    mux := http.NewServeMux()

    mux.HandleFunc("/", c.GetIndexHTMLHandler(ctx, func(r *http.Request) string { return "" }))
	mux.HandleFunc("/assets/", c.GetAssetHandler(ctx, false))
	mux.HandleFunc("/service-worker.js", c.GetAssetHandler(ctx, true))
```