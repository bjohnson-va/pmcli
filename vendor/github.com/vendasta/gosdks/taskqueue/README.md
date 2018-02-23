# TaskQueue
`taskqueue` provides applications with the ability to put payloads into queues.

The payloads in these queues can be processed by applications asyncronously.

The semantics of this package are identical to [appengine pull queues](https://cloud.google.com/appengine/docs/standard/python/taskqueue/pull/). Push queues are not supported.

## Disclaimer
The current implementation of the API uses [Google Cloud Tasks](https://cloud.google.com/cloud-tasks/) which is still in alpha. They are still iterating fast and have broken the API recently. At this time, it is not recommended you use this package in production unless you really know what you are doing and are able to devote the time and care to keep up with future breaking changes.

Furthermore, `taskqueue` is not supported on local environments at this time. This is because each developer creating individual pull queues for their local environments would quickly exhaust our quotas. We may be able to supply a stub on local instead.

## TODOs:
- Pull many tasks at once, maintain a pool of work
- Non-integration Tests :agile:

## Overview
There are two main objects you will interact with from this package: The `taskqueue.Client` and the `taskqueue.Worker`.

The `taskqueue.Client` is responsible for scheduling tasks - putting work into the queue.

The `taskqueue.Worker` is responsible for processing tasks from the queue.

## Creating a client
This code creates a client configured for the test environment.
```golang
c, err := taskqueue.NewClient(ctx, config.Test)
if err != nil {
    log.Fatalf("error making client: %s", err.Error())
}
```
## Task Definitions
Under the hood, tasks are just serialized payloads. This SDK uses json serialization, so define your task as a JSON-tagged struct:
```golang
type MyTaskPayload struct {
	AccountGroupID string `json:"accountGroupId"`
	SnapshotID string `json:"snapshotId"`
	Section string `json:"section"`
}
```
## Scheduling Work
This code schedules a task into a queue called "data-collection". Note that you do not need to prefix or suffix the queue name with your environment, the client does this for you.
```golang
payload := &MyTaskPayload{
    AccountGroupID: "AG-123",
    SnapshotID: "S-456",
    Section: "listings",
}
err = c.ScheduleTask(ctx, "data-collection", payload)
if err != nil {
    log.Fatalf("error scheduling task: %s", err.Error())
}
```

### Delaying Task Processing
If you want to offset the processing of a task from the current time, you can pass the `DelayProcessingBy` option:
```golang
c.ScheduleTask(ctx, "data-collection", payload, taskqueue.DelayProcessingBy(time.Minute * 30) // this task can't be leased until 30 minutes from now
```

### Scheduling Task Processing at a specific time
If you know exactly when you want your task to be eligible for processing, you can pass the `ProcessAt` option:
```golang
c.ScheduleTask(ctx, "data-collection", payload, taskqueue.ProcessAt(time.Date(2018, time.January, 26, 3, 20, 13, 40, loc)) // this task can't be leased until exactly Jan 26, 2018 at 03:20:13.40 AM
```

### Scheduling a Task with a tag
Tags allow similar tasks to be processed in a batch. If you label tasks with a tag, your worker can lease tasks with the same tag using a filter.
```golang
c.ScheduleTask(ctx, "data-collection", payload, taskqueue.WithTag("review-data-collection"))
```

### Scheduling a Task with a name
Task names are unique within a queue for about 10 days. If you want to guarantee uniqueness of a unit of work you schedule it with a specific name.
The name can contain only letters ([A-Za-z]), numbers ([0-9]), hyphens (-), or underscores (_)
```golang
c.ScheduleTask(ctx, "data-collection", payload, taskqueue.WithName("AG-123-review-data-collection"))
```


## Processing Work
In order to consume tasks from the queue, you need to provide a handler function that works in terms of your task payload definition:
```golang
h := func(ctx context.Context, msg taskqueue.Task) error {
    log.Printf("raw task payload: %s", msg.Payload)
    p := &MyTaskPayload{}
    err := json.Unmarshal([]byte(msg.Payload), p)
    if err != nil {
        return err
    }
    log.Printf("deserialized account group id: %s", p.AccountGroupID)
    log.Printf("deserialized snapshot id: %s", p.SnapshotID)
    log.Printf("deserialized section: %s", p.Section)

    err = doWork(p.AccountGroupID, p.SnapshotID, p.Section)
    return err
}
```
If this function returns an error, the task will be nacked and placed back into the queue. If the function does not return an error, the task will be acknowledged as completed.

Now we create a worker and pass it our task handler and client as arguments:
```golang
w := taskqueue.NewWorker("data-collection", h, c)
```

To start working, we can call `w.Work(ctx)` with a context. We want to cancel this context when we want the worker to stop. It is common to use a goroutine to have the worker run asynchronously:
```golang
go w.Work(ctx)
```

### Lease Duration
By default we lease tasks for a minute. If you are concerned that processing your task will take longer than a minute or should take much shorter than a minute, you can set a custom lease duration on the worker:
```golang
w := taskqueue.NewWorker("data-collection", h, c, taskqueue.WithLeaseDuration(time.Minute * 5))
```

### Lease Filter
A `filter` can be used to specify a subset of tasks to lease. The tasks must have been previously scheduled with the specified tag.
```golang
w := taskqueue.NewWorker("data-collection", h, c, taskqueue.WithTagFilter("review-data-collection"))
```

### Cancel function
When the worker's context is cancelled, you may find yourself needing to do some kind of cleanup work before the worker stops pulling completely. You can provide a custom cancel func using the `WithCancelFunc` option:
```golang
cf := func(){
    logging.Infof(parentCtx, "Cleaning stuff up")
    //logic
}

w := taskqueue.NewWorker("data-collection", h, c, taskqueue.WithCancelFunc(cf))
```

### Backoff parameters

#### Cloud Tasks API Errors
You can provide a custom backoff configuration that controls how your worker will back off from errors from the Cloud Tasks API. We provide sensible default values that be overwritten using the `WithBackoff` option:
```golang
w := taskqueue.NewWorker("data-collection", h, c, taskqueue.WithBackoff(&backoff.Backoff{// see the backoff pkg}))
```

#### Handler Errors
You can also specify a custom backoff configuration to control how your worker will back off from errors returned by your task handler. No default backoff is instrumented, but one can be specified using the `WithHandlerBackoff` option:
```golang
w := taskqueue.NewWorker("send-emails", h, c, taskqueue.WithHandlerBackoff(&backoff.Backoff{// see the backoff pkg}))
```

## Contact

Dustin Walker