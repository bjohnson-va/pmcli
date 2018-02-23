# Marketplace Webhook Handling
See https://www.vendasta.com/developers/vendors/web-hooks for webhook documentation before getting started.

## Usage

This package takes care of some of the complexity around decoding marketplace webhook JWTs. Because the body of the
webhook request consists solely of a JWT and not valid JSON, we need to handle them using our HTTP server instead of our
GRPC one.

Products that need to listen to marketplace webhooks should attach their handlers using this package. Everything to do
with decoding and authenticating the JWT is taken care of leaving your code to simply deal with processing the payload.

### Attaching Handlers
This is an example of what you might add to your `main.go` for attaching webhook handlers:
```
mux := http.NewServeMux()
webhookServer, err := marketplacewebhooks.NewServer(ctx, config.CurEnv())
if err != nil {
    logging.Criticalf(ctx, "Error initializing marketplace webhook server: %s", err.Error())
    os.Exit(-1)
}
mux.HandleFunc("/webhook/purchase", webhookServer.GetPurchaseWebhookHandler(webhooks.HandlePurchaseWebhook))
mux.HandleFunc("/webhook/account", webhookServer.GetAccountWebhookHandler(webhooks.HandleAccountWebhook))
mux.HandleFunc("/webhook/session", webhookServer.GetSessionWebhookHandler(webhooks.HandleSessionWebhook))
```

### Handling Payloads
#### Purchase Webhook Example
```
func HandlePurchaseWebhook(payload *marketplacewebhooks.PurchaseWebhookPayload, url *url.URL) error {
  myOrderForm := MyAppOrderForm{}
  if err := json.Unmarshal(payload.OrderForm, &myAppOrderForm); err != nil {
    // probably want to log this and return an error
  }

  // provision the account

  return nil
}
```

If your handler function returns an error, the webhook will retry as per marketplace's retry semantics.

The `url` of the request is provided for the case where your handling code needs to know about query parameters on
the request, but can otherwise be ignored.

For an actual purchase webhook implementation example, check out: https://github.com/vendasta/medical-compare/blob/master/internal/webhook/purchase.go

#### Session Webhook Example
```
func HandleSessionWebhook(payload *marketplacewebhooks.SessionPayload) (code string, err error) {
    code := uuid.New()
    // store session somewhere keyed with the code you just generated (could be redis, could be vstore, up to you)
    return code
}
```
For an actual session webhook implementation example, check out: https://github.com/vendasta/medical-compare/blob/master/internal/webhook/session.go

#### Order Form
The order form on the purchase webhook (`PurchaseWebhookPayload.OrderForm`) exists as a JSON byte string (`[]byte`)
rather than an instance of a struct. This is because the order form can be different for every app. You will
likely want to `json.Unmarshal` this into your own struct representation of your app's order form.