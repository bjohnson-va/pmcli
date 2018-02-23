// This file has been borrowed from gax, which is a common gRPC util library used by Google's golang SDKs.
// Original found here: https://github.com/googleapis/gax-go

package vax

import (
	"time"

	"fmt"
	"net/http"

	"golang.org/x/net/context"
)

// A user defined call stub.
type APICall func(context.Context, CallSettings) error

// RetryableError indicates to vax.Invoke that the error returned is deemed retryable.
func RetryableError(err error) error {
	return retryableError{err}
}

// RetryOnResponse indicates to vax.Invoke that the response is an error and can be retried.
func RetryOnResponse(r *http.Response) error {
	if r.Request == nil {
		return retryableError{
			fmt.Errorf("Retrying on HTTP status %d", r.StatusCode),
		}
	}
	return retryableError{
		fmt.Errorf("Retrying on HTTP status %d to %s to %s", r.StatusCode, r.Request.Method, r.Request.URL),
	}
}

type retryableError struct {
	err error
}

func (re retryableError) Error() string {
	return re.err.Error()
}

// Invoke calls the given APICall,
// performing retries as specified by opts, if any.
func Invoke(ctx context.Context, call APICall, opts ...CallOption) error {
	var settings CallSettings
	for _, opt := range opts {
		opt.Resolve(&settings)
	}
	return invoke(ctx, call, settings, Sleep)
}

// Sleep is similar to time.Sleep, but it can be interrupted by ctx.Done() closing.
// If interrupted, Sleep returns ctx.Err().
func Sleep(ctx context.Context, d time.Duration) error {
	t := time.NewTimer(d)
	select {
	case <-ctx.Done():
		t.Stop()
		return ctx.Err()
	case <-t.C:
		return nil
	}
}

type sleeper func(ctx context.Context, d time.Duration) error

// invoke implements Invoke, taking an additional sleeper argument for testing.
func invoke(ctx context.Context, call APICall, settings CallSettings, sp sleeper) error {
	var retryer Retryer
	for {
		err := call(ctx, settings)
		if err == nil {
			return nil
		}
		if settings.Retry == nil {
			return err
		}
		if retryer == nil {
			if r := settings.Retry(); r != nil {
				retryer = r
			} else {
				return err
			}
		}
		if d, ok := retryer.Retry(ctx, err); !ok {
			return err
		} else if err = sp(ctx, d); err != nil {
			return err
		}
	}
}
