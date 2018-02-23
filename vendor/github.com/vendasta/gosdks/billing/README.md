# Billing Microservice golang SDK
- [Introduction](https://github.com/vendasta/gosdks/tree/master/billing#installation)
- [Usage](https://github.com/vendasta/gosdks/tree/master/billing#usage)
  - [Introduction](https://github.com/vendasta/gosdks/tree/master/billing#introduction)

## Installation

* YOLO development? `go get github.com/vendasta/gosdks`
* Using [glide](https://github.com/Masterminds/glide) package manager? `glide get -s -v github.com/vendasta/gosdks`

## Usage

### Introduction

```golang
billing, err := billing.NewClient(ctx, config.CurEnv(), grpc.WithUnaryInterceptor(logging.ClientInterceptor()))
if err != nil {
    logging.Errorf(ctx, "Error initializing billing client %s", err.Error())
    os.Exit(-1)
}

ctx := context.Background()
merchantID := "ABC"
sku := "RM"
customerID := "AG-123"
orderID := "RM-123"
var expiry time.Time
billingStart := time.Now().UTC()
err := billing.CreateBillableItem(ctx, merchantID, sku, customerID, orderID, expiry, billingStart)
if err != nil {
    logging.Infof(ctx, "Error creating billable item for merchant %s: %s", merchantID, err.Error())
    return err
}
```
