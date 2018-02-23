package vax

import (
	"errors"
	"testing"
	"time"

	"golang.org/x/net/context"
)

var canceledContext context.Context

func init() {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	canceledContext = ctx
}

// recordSleeper is a test implementation of sleeper.
type recordSleeper int

func (s *recordSleeper) sleep(ctx context.Context, _ time.Duration) error {
	*s++
	return ctx.Err()
}

type boolRetryer bool

func (r boolRetryer) Retry(ctx context.Context, err error) (time.Duration, bool) { return 0, bool(r) }

func TestInvokeSuccess(t *testing.T) {
	apiCall := func(context.Context, CallSettings) error { return nil }
	var sp recordSleeper
	err := invoke(context.Background(), apiCall, CallSettings{}, sp.sleep)

	if err != nil {
		t.Errorf("found error %s, want nil", err)
	}
	if sp != 0 {
		t.Errorf("slept %d times, should not have slept since the call succeeded", int(sp))
	}
}

func TestInvokeNoRetry(t *testing.T) {
	apiErr := errors.New("foo error")
	apiCall := func(context.Context, CallSettings) error { return apiErr }
	var sp recordSleeper
	err := invoke(context.Background(), apiCall, CallSettings{}, sp.sleep)

	if err != apiErr {
		t.Errorf("found error %s, want %s", err, apiErr)
	}
	if sp != 0 {
		t.Errorf("slept %d times, should not have slept since retry is not specified", int(sp))
	}
}

func TestInvokeNilRetry(t *testing.T) {
	apiErr := errors.New("foo error")
	apiCall := func(context.Context, CallSettings) error { return apiErr }
	var settings CallSettings
	WithRetry(func() Retryer { return nil }).Resolve(&settings)
	var sp recordSleeper
	err := invoke(context.Background(), apiCall, settings, sp.sleep)

	if err != apiErr {
		t.Errorf("found error %s, want %s", err, apiErr)
	}
	if sp != 0 {
		t.Errorf("slept %d times, should not have slept since retry is not specified", int(sp))
	}
}

func TestInvokeNeverRetry(t *testing.T) {
	apiErr := errors.New("foo error")
	apiCall := func(context.Context, CallSettings) error { return apiErr }
	var settings CallSettings
	WithRetry(func() Retryer { return boolRetryer(false) }).Resolve(&settings)
	var sp recordSleeper
	err := invoke(context.Background(), apiCall, settings, sp.sleep)

	if err != apiErr {
		t.Errorf("found error %s, want %s", err, apiErr)
	}
	if sp != 0 {
		t.Errorf("slept %d times, should not have slept since retry is not specified", int(sp))
	}
}

func TestInvokeRetry(t *testing.T) {
	const target = 3

	retryNum := 0
	apiErr := errors.New("foo error")
	apiCall := func(context.Context, CallSettings) error {
		retryNum++
		if retryNum < target {
			return apiErr
		}
		return nil
	}
	var settings CallSettings
	WithRetry(func() Retryer { return boolRetryer(true) }).Resolve(&settings)
	var sp recordSleeper
	err := invoke(context.Background(), apiCall, settings, sp.sleep)

	if err != nil {
		t.Errorf("found error %s, want nil, call should have succeeded after %d tries", err, target)
	}
	if sp != target-1 {
		t.Errorf("retried %d times, want %d", int(sp), int(target-1))
	}
}

func TestInvokeRetryTimeout(t *testing.T) {
	apiErr := errors.New("foo error")
	apiCall := func(context.Context, CallSettings) error { return apiErr }
	var settings CallSettings
	WithRetry(func() Retryer { return boolRetryer(true) }).Resolve(&settings)
	var sp recordSleeper

	err := invoke(canceledContext, apiCall, settings, sp.sleep)

	if err != context.Canceled {
		t.Errorf("found error %s, want %s", err, context.Canceled)
	}
}
