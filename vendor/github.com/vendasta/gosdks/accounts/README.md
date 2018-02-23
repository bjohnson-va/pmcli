# Important
This is an alpha release and will contain future changes that are not backwards compatible.

# Accounts Microservice golang SDK

## Installation

* Waterfall development? `go get github.com/vendasta/gosdks`
* Using [glide](https://github.com/Masterminds/glide) package manager? `glide get -s -v github.com/vendasta/gosdks`

## Usage

### Introduction



```golang
accountsClient, err := accounts.NewClient(ctx, config.CurEnv(), grpc.WithUnaryInterceptor(logging.ClientInterceptor()))
if err != nil {
    logging.Errorf(ctx, "Error initializing accounts client %s", err.Error())
    os.Exit(-1)
}

businessID := "AG-1234"
accounts, err := accountsClient.List(ctx, businessID)
if err != nil {
    logging.Infof(ctx, "Error getting accounts from business %s: %s", businessID, err.Error())
    return nil, err
}
```
