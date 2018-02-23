# Stub Server
The Stub Server can be used to test external services in cases where creating a stub
is difficult. This is a basic server which will return a response specified by the
caller to provide consistent test data.

## Initialization
To create a new server you must call `NewServer()`. This prepares everything needed
for the server to run and will return an struct which implements the StubServer interface.

To start the server you must call `Start()` on the StubServer. This will start the server
locally on a random port. To get the address to make requests to you may call *URL()*. HTTP
calls made to that address will return items from the server's response queue.

Once you are done with the server you should call `Stop()`. This will stop the server.
This can typically be deferred right after the server has been created.

## Responses
Responses from the server are handled by a queue of *Response*. If there
are no responses in the queue a response with the status code of 500 will be returned
along with the message: `internal error\n`

You can push your own responses to the queue using the *PushResponse* function on the server.
The method takes in a pointer to a *Response*. The *Response* has two properties
*StatusCode* and *Response*. *StatusCode* is the HTTP status code you want returned, and *Response* is
the response string you want returned.

Once your *Response* is passed to the server it will be added to the end of the queue. That
response will be removed from the queue once a caller has received that response.

## Example
```
// Create a new server
s := stubserver.NewServer()
// Make sure the server is stopped when we are done with it
defer s.Stop()
// Start the server
s.Start()

// Push responses into the server queue
s.PushResponse(&stubserver.Response{StatusCode: 200, Response: "Resp 1"})
s.PushResponse(&stubserver.Response{StatusCode: 401, Response: "Resp 2"})

// Get a response from the server
resp, _ := http.Get(s.URL())
defer resp.Body.Close()
body, _ := ioutil.ReadAll(resp.Body)

// This will print out Resp 1 since it is the first item in the queue
fmt.Println(string(body))
```
