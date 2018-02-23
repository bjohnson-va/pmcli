Core Services SDK
=================
An sdk to communicate with Core Services's apis via Go.

## Versions
### 4.1.0 [2017-08-22]
Added ScheduledPartnerPosts endpoint
Removed the private function convertTimeToVAPITimestamp and added it to basesdk as a public method
### 4.0.0 [2017-07-06]
BREAKING CHANGE: Contexts must be provided to client calls -- this enables tracing
### 3.3.0 [2017-06-27]
Added Taxonomy validation rule
Added URL and Email Validator
### 3.2.0 [2017-06-21]
Added Infer Geolocation API
### 3.1.0 [2017-06-20]
Added taxonomy APIs.
Fix signatures of Social Post client and Social Service client to use env as an input rather than rootUrl, like the previous 3.0.0 breaking change claimed it did.
### 3.0.0 [2017-06-01]
Require the environment be passed in to the clients to determine the proper host.
### 2.0.0 [2017-06-01]
Require the root url to be passed in to the clients. The environment switch for which url to pass in should be in the environment variables for the µs
### 1.1.0 [2017-06-01]
Added social service and social posting apis
### 1.0.0
Initial Release

## Usage
Core Services has a lot of APIs. Each category of API has its own client object so that interfaces are kept tight and simple.

### Instantiating a Client
You need to provide a valid apiUser and apiKey that is registered with Core Services. You must also specify the environment you want to connect to using one of the constants from the gosdks config package.

```
client := BuildTaxonomyClient("myApiUser", "myApiKey", config.Test)
```

### Validation Rules & Utilities

```golang
v := validation.NewValidator()

v.Rule(cssdk.TaxonomyIDValidationRule(cssdkClient, ctx, taxIDs, util.InvalidArgument, "Invalid Taxonomy ID"))
```
