#Package `logging`

Provide a set of functions to flow your applications' logs to Google Cloud Console

## Getting Started

### Prerequisites

You have to define environment variable `HOSTNAME` in your deployment yaml file in order to user this package 
([mscli](https://github.com/vendasta/mscli) may already did this for you):

```yaml
- name: HOSTNAME
  valueFrom:
    fieldRef:
      fieldPath: metadata.name
```

### Example

```go
package main
import (
    "log"
    "context"
    
    "github.com/vendasta/gosdks/logging"
)

func main() {
    namespace := "Your app's namespace"
    podName := "Your app's pod name"
    appName := "Your app's name"
    ctx := context.Background()
    
    err := logging.Initialize(namespace, podName, appName)
	if err != nil {
		log.Fatalln("error initializing logging module, err: %s", err.Error())
	}
    
    logging.Infof(ctx, "info blahblah...")
    logging.Warningf(ctx, "warning blahblah...")
    logging.Errorf(ctx, "error blahblah...")
}
```


## Using logging to create "span"s in traces in Google Cloud Platform

If your function takes in a Context you can use it to set more specific information about your request. To do so is simple:

```go
func myFunc(ctx Context) err {
  trace := logging.FromContext(ctx)
  
  span := trace.NewChild("span-name")
  // Do some stuff! Call other functions, whatever.
  span.Finish()
  
  span := trace.NewChild("another-span-name")
  defer span.Finish()
  
  // Do a bunch of other stuff
  return nil
}
```

In the example above, `span-name` and `another-span-name` will be sibling spans in the stack trace graph, since they're created
at the same level in code. You can (and should!) nest spans as well. Any spans you create while you have a `NewChild` open
(i.e. `Finish()` has not been called on it) will nest inside of the current span. Here's a code example:

```go
func mySecondFunc(ctx Context) err {
  trace := logging.FromContext(ctx)
  
  span := trace.NewChild("parent-span")
  someFunc(ctx)
  span.Finish()
  
  return nil
}

func someFunc(ctx Context) {
  span := logging.FromContext(ctx).NewChild("child-span")
  logging.Printf("Hello world from someFunc!")
  span.Finish()
}
```

It may be important to note that you could have deferred the span.Finish() and it will still nest, as defer'ing the call 
will have it run after the function is complete.

## Tagging Your Logs

It might be useful to add tags to your logs so that you can filter them easier.

```go
logging.Tag(ctx, "key", "value")
```

Will attach a label `"key": "value"` to entries logged with that particular `ctx`.

Tagging contexts that originate from a GRPC or HTTP entrypoint will just work. In order to tag a context that is not associated with a request (such as a background job), you will need to use the context returned by `NewTaggedContext`.

```go
// this can be any context, not just context.Background()
ctx := logging.NewTaggedContext(context.Background())
logging.Tag(ctx, "key", "value")
```

If you need to differentiate between different contexts operating in the same function, you may find it useful to have a UUID in order to tie together a single stream of execution. You can use the helper:
```go
ctx := logging.NewWorkerContext(context.Background())
```