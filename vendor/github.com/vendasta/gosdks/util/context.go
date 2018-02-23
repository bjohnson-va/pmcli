package util

import (
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
)

type options struct {
	timeout time.Duration
}

// Option is used as to configure NewContext.
type Option func(*options)

// WithTimeout overrides the default timeout of NewContext.
// Example: NewContext(ctx, WithTimeout(time.Second))
func WithTimeout(d time.Duration) Option {
	return func(o *options) {
		o.timeout = d
	}
}

// NewContext returns a new context with tracing metadata included and a specific timeout set. The timeout can also
// optionally be overridden.
//
// This method makes it easy to include tracing data when reusing a context for outbound gRPC requests.
func NewContext(ctx context.Context, opts ...Option) context.Context {
	o := &options{
		timeout: time.Second * 10,
	}
	for _, opt := range opts {
		opt(o)
	}

	ctx, _ = context.WithTimeout(ctx, o.timeout)
	newMd := metadata.MD{}
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		traceID, ok := md["x-cloud-trace-context"]
		if ok {
			newMd["x-cloud-trace-context"] = traceID
		}
	}
	return metadata.NewOutgoingContext(ctx, newMd)
}
