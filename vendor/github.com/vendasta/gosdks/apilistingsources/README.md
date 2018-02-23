# Api listing sources
- [Introduction](https://github.com/vendasta/gosdks/tree/master/bingplacessdk#installation)

## Installation

 `go get github.com/vendasta/gosdks`


### Introduction

api listing sources is an adapter calling external listing api's

## Usage

In main.go
```golang
c, err := apilistingsources.NewClient(ctx, config.CurEnv(), grpc.WithUnaryInterceptor(logging.ClientInterceptor()))
    if err != nil {
        logging.Criticalf(ctx, "Error creating api listing sources client: %s", err.Error())
        os.Exit(-1)
	}
```