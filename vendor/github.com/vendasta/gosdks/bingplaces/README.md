# Bing Places
- [Introduction](https://github.com/vendasta/gosdks/tree/master/bingplacessdk#installation)

## Installation

* YOLO development? `go get github.com/vendasta/gosdks`


### Introduction

Bing Places is an adapter for the bing places javascript api. This client allows you to 
interact with the Search and Get API methods.

## Usage

In main.go
```golang
c, err := bingplaces.NewClient(ctx, config.CurEnv(), grpc.WithUnaryInterceptor(logging.ClientInterceptor()))
    if err != nil {
        logging.Criticalf(ctx, "Error creating bing places client: %s", err.Error())
        os.Exit(-1)
	}
```