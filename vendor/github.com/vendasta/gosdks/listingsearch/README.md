# Listing Search Client

## Installation
- Install gosdks `go get github.com/vendasta/gosdks`

### Introduction

Listing Search Client provides an adapter for api listing sources. This client allows
you to interact with the Search and Get API methods.

### Adding New Api service
1. Add the listing ID to the const in configs.go
2. Add the address and scope to the configs.go file
3. In the client.go and create the new client code in NewClient method


## Usage

In main.go
```
c, err := listingSearch.NewClient(ctx, config.CurEnv(), grpc.WithUnaryInterceptor(logging.ClientInterceptor()))
    if err != nil {
        logging.Criticalf(ctx, "Error creating listing search client: %s", err.Error())
        os.Exit(-1)
	}
```
