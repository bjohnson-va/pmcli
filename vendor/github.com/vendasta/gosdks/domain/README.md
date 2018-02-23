# Domain
- [Introduction](https://github.com/vendasta/gosdks/tree/master/domain#installation)
- [Usage](https://github.com/vendasta/gosdks/tree/master/domain#usage)
  - [Introduction](https://github.com/vendasta/gosdks/tree/master/domain#introduction)

## Installation

* YOLO development? `go get github.com/vendasta/gosdks`


### Introduction

Domain serves domain related functionality such as looking up owners, dns records, etc.

## Usage

In main.go
```golang
c, err := domain.NewClient(ctx, "medical-compare-client", grpc.WithUnaryInterceptor(logging.ClientInterceptor()))
    if err != nil {
        logging.Criticalf(ctx, "Error creating domain client: %s", err.Error())
        os.Exit(-1)
	}
```