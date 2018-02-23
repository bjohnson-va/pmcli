# Listing Score Microservice golang SDK
- [Introduction](https://github.com/vendasta/gosdks/tree/master/listingscore#installation)
- [Usage](https://github.com/vendasta/gosdks/tree/master/listingscore#usage)
  - [Introduction](https://github.com/vendasta/gosdks/tree/master/listingscore#introduction)

## Installation

* YOLO development? `go get github.com/vendasta/gosdks`
* Using [glide](https://github.com/Masterminds/glide) package manager? `glide get -s -v github.com/vendasta/gosdks`

## Usage

### Introduction

```golang
listingscore, err := listingscore.NewClient(ctx, config.CurEnv(), grpc.WithUnaryInterceptor(logging.ClientInterceptor()))
if err != nil {
    logging.Errorf(ctx, "Error initializing listing score client %s", err.Error())
    os.Exit(-1)
}

accountGroupID := "AG-1234"
score, err := listingscore.GetScore(ctx, accountGroupID)
if err != nil {
    logging.Infof(ctx, "Error getting score for account group %s: %s", accountGroupID, err.Error())
    return nil, err
}
```
