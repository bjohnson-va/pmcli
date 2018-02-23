# PUBSUB

## This is not for subscribing to vStore callbacks
How to write a handler and register it for vstore callbacks can be found in the vstore documentation.  The pubsub code vstore uses for that uses a subset of this lib so they will be similar though.

## So, you want to listen to a arbitrary pubsub message and so something about with it.
1. In main.go initialize a pubsub client (this can be used multiple times and you only need one client per project)
```
pubsubClient, err := pubsub.NewGooglePubsubClient(ctx)
if err != nil {
	logging.Criticalf(ctx, "Error creating pubsub client: %s", err.Error())
	os.Exit(-1)
}
```
2. Setup the code that will handle the function.
The interface:
```
type Handler interface {
	GetTopicName() string
	GetSubscriptionName() string
	Handler(ctx context.Context, message *pubsub.Message) error
	Cancel()
}
```
GetTopicName: should return the name of the topic you want to listen to (not including the /project/repcore-prod/topic/ part. EG: `medicare-physician-listing-transfer`

GetSubscriptionName: should return the name of your specific subscription (EG: `data-transfer-medicare-physician-listing-transfer`)

Handler: does the actual work for your message, if it returns an error the message will not be acknowledged and you will recieve it again. If it returns nil the message will be acknowledged

Cancel: Do cleanup work if necessary because the workers for your handler are being shutdown

3. Now that you have the handler set up and the pubsub client created you can start the work from main.go with StartHandling
```
err = pubsub.StartHandling(ctx, pubsubClient, handler)
if err != nil {
	logging.Criticalf(ctx, "Error setting up pubsub workers: %s", err.Error())
	os.Exit(-1)
}
```

