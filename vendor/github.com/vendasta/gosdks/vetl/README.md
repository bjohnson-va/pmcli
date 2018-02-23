# Important
This is an alpha release and will contain future changes that are not backwards compatible.

# VETL golang SDK
- [Installation](#installation)
- [Usage](#usage)
  - [Initialization](#initialization)
  - [Registering Data Sources](#registering-data-sources)
    - [VStore Source](#vstore-source)
  - [Adding Transforms](#adding-transforms)
    - [Keep Properties](#keep-properties)
  - [Registering Subscriptions](#registering-subscriptions)
    - [VStore Sink](#vstore-sink)
  - [Backfilling Subscriptions](#backfilling-subscriptions)
- [Contact](#contact)

## Installation

* Using dep? `dep ensure -update github.com/vendasta/gosdks`

## Usage

### Initialization

To initialize a client:

```golang
client, err := vetl.NewClient(ctx, config.Test)
if err != nil {
    logging.Criticalf(ctx, "Failed to initialize vetl client because %s", err.Error())
    os.Exit(-1)
} 
```

### Registering Data Sources

In order to be able to transform your microservice's data using `vETL`, you need to register your data with `vETL` as a `DataSource`.

Doing this allows `vETL` to read your data directly from it's storage location.

#### VStore Source

Currently, `VStore` is the only supported `DataSource`. You can register a `vstore.Model` against `vETL` like so:

```golang
vetlSchema, err := vetl.DataSourceFromVStoreModel((*MyVStoreModel)(nil), "myPubsubSecondaryIndexID")
if err != nil {
    logging.Criticalf(ctx, "Failed to derive vetl schema from vstore model because %s", err.Error())
    os.Exit(-1)
}
err = client.CreateDataSource(ctx, "my-unique-source-id", vetlSchema)
if err != nil {
    logging.Criticalf(ctx, "Failed to register vstore model with vetl client because %s", err.Error())
    os.Exit(-1)
}
```

`"myPubsubSecondaryIndexID"` is the name of the pubsub secondary index already registered on *MyVStoreModel that we want vETL to read from.

### Adding Transforms

Data is transformed in `vETL` by adding transform definitions to existing data sources or existing transforms.

This allows us to programmatically define and build up a pipeline of operations that can be represented as a tree.
```
      datasource
          |
      transform1
       |      |
transform2  transform3
```

#### Keep Properties

The **Keep Properties** transform has the client specify a subset of properties they want kept as an output of this transformation step.

Use this transform to control and strongly define the properties available to downstream transforms and subscribers. 

```golang
// regardless of which properties are available from the datasource, the only properties available after this step will be "partner_id", "account_group_id", and "market_id"
err := client.UpsertTransform(ctx, []string{"my-unique-source-id"}, "keep-id-properties", KeepPropertiesTransform([]string{"partner_id", "account_group_id", "market_id"}))
if err != nil {
    logging.Criticalf(ctx, "Failed to attach keep properties transform because %s", err.Error())
    os.Exit(-1)
}
```

### Registering Subscriptions

In order for `vETL` to write the output of a transform definition anywhere, a client must subscribe to a transform, providing enough information for `vETL` to be able to write this output somewhere.

This output definition is called a `sink` and it should be unique for each subscription that is registered. A sink might be a VStore or CloudSQL table, but it might also be a Cloud Pubsub topic.

Do note that clients can only subscribe to transforms they either own, or have been marked public. Client A can not subscribe to Client B's private transforms.

Also note that data sources can not be subscribed to directly, regardless of who owns them. This is to prevent laziness - write transforms that establish a proper API - you have to support this forever!

#### VStore Sink

Specify a VStore Sink by using the `vetl.VStoreDataSink` function. This allows you specify the destination namespace, kind and primary key.

`vETL` will figure out the schema based on the transform you're subscribing to, as well as check that your primary key is a valid option based on that schema.

TODO: Provide a way to get a schema for any transform so that it is easy for the client to know how to read vETL's written model from VStore - dwalker

```golang
err := client.CreateSubscription(ctx, "my-subscription-id", "keep-id-properties", vetl.VStoreDataSink("mynamespace", "destinationkind", []string{"account_group_id"}))
if err != nil {
    logging.Criticalf(ctx, "Failed to subscribe to `keep-id-properties` because %s", err.Error())
    os.Exit(-1)
}
```

#### CloudSQL via VStore

You can push the output of a transform definition into CloudSQL by registering a VStore Sink with a CloudSQL secondary index like so:

```golang
err := client.CreateSubscription(ctx, "my-subscription-id", "keep-id-properties",
    vetl.VStoreDataSink("mynamespace", "destinationkind", []string{"account_group_id"},
        vetl.CloudSQLIndex("cloud-sql", "104.154.165.235", "root", "password", sqlClientKey, sqlClientCert, sqlCA, "repcore-prod", "my-instance-name")
        )
    )
if err != nil {
    logging.Criticalf(ctx, "Failed to subscribe to `keep-id-properties` because %s", err.Error())
    os.Exit(-1)
}
```

#### Tesseract Sink (Alpha)

Specify a Tesseract Sink by using the `vetl.TesseractDataSink` function. This allows you specify the destination namespace, kind and primary key.

`vETL` will figure out the schema based on the transform you're subscribing to, as well as check that your primary key is a valid option based on that schema.

Notes:
 - Tesseract does not support Structured properties and will ignore any that are provided.
 - Tesseract requires that you provide how you would like it to manage the concurrency control for the underlink sink. The two mechanisms
 supported are version and last modified. The version concurrency control requires that you provide a `version` field that is a monotonically increasing
 integer (vStore manages this for you automatically). The last modified concurrency control requires that you provide a `last_modified` field that is a
 timestamp field that is updated on each write. This is typically how most datastore models will leverage concurrency control.


```golang
err := client.CreateSubscription(ctx, "my-subscription-id", "keep-id-properties", vetl.TesseractDataSink("mynamespace", "destinationkind", []string{"account_group_id"}, vetl.TesseractLastModifiedConcurrencyControl("last_modified")))
if err != nil {
    logging.Criticalf(ctx, "Failed to subscribe to `keep-id-properties` because %s", err.Error())
    os.Exit(-1)
}
```

### Backfilling Subscriptions

After registering a subscription, `vETL` will only be sending new messages. In order to backfill historical data, you will need to call

`BackfillSubscription` and simply pass the subscription ID you defined during the subscription registration.

Since this request will essentially provision a migration pipeline, you should ensure you only call it once. Instead of using the golang SDK in your microservice to call it, which can be a problem if your service is replicated across more than a single pod or if it restarts due to some other reason, you can use the postman collection defined in the `vetl` repository: https://github.com/vendasta/vETL/tree/master/integration . Follow the instruction there to safely invoke this command.

```golang
err := client.BackfillSubscription(ctx, "my-subscription-id")
if err != nil {
    logging.Criticalf(ctx, "Failed to start the subscription backfill: %s", err.Error())
    os.Exit(-1)
}
```

## Contact

Dustin Walker

