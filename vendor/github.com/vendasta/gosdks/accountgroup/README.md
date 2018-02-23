# Important
This is an alpha release and will contain future changes that are not backwards compatible.

# Account Group Microservice golang SDK
- [Introduction](https://github.com/vendasta/gosdks/tree/master/accountgroup#installation)
- [Usage](https://github.com/vendasta/gosdks/tree/master/accountgroup#usage)
  - [Introduction](https://github.com/vendasta/gosdks/tree/master/accountgroup#introduction)

## Installation

* YOLO development? `go get github.com/vendasta/gosdks`
* Using [glide](https://github.com/Masterminds/glide) package manager? `glide get -s -v github.com/vendasta/gosdks`

## Usage

### Introduction

```golang
accountGroupClient, err := accountgroup.NewClient(ctx, config.CurEnv(), grpc.WithUnaryInterceptor(logging.ClientInterceptor()))
if err != nil {
    logging.Errorf(ctx, "Error initializing account group client %s", err.Error())
    os.Exit(-1)
}

accountGroupID := []string{"AG-1234"}
accountGroups, err := accountGroupClient.GetMulti(ctx, accountGroupIDs)
if err != nil {
    logging.Infof(ctx, "Error getting account group %s: %s", accountGroupID, err.Error())
    return nil, err
}
```
