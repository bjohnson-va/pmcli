package vax

import (
	"errors"
	"testing"
	"time"

	"net/http"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

var _ Retryer = &boRetryer{}

func TestBackoffDefault(t *testing.T) {
	backoff := Backoff{}

	max := []time.Duration{1, 2, 4, 8, 16, 30, 30, 30, 30, 30}
	for i, m := range max {
		max[i] = m * time.Second
	}

	for i, w := range max {
		if d := backoff.Pause(); d > w {
			t.Errorf("Backoff duration should be at most %s, got %s", w, d)
		} else if i < len(max)-1 && backoff.cur != max[i+1] {
			t.Errorf("current envelope is %s, want %s", backoff.cur, max[i+1])
		}
	}
}

func TestBackoffExponential(t *testing.T) {
	backoff := Backoff{Initial: 1, Max: 20, Multiplier: 2}
	want := []time.Duration{1, 2, 4, 8, 16, 20, 20, 20, 20, 20}
	for _, w := range want {
		if d := backoff.Pause(); d > w {
			t.Errorf("Backoff duration should be at most %s, got %s", w, d)
		}
	}
}

func TestOnCodes(t *testing.T) {
	// Lint errors grpc.Errorf in 1.6. It mistakenly expects the first arg to Errorf to be a string.
	errf := grpc.Errorf
	apiErr := errf(codes.Unavailable, "")
	tests := []struct {
		c     []codes.Code
		retry bool
	}{
		{nil, false},
		{[]codes.Code{codes.DeadlineExceeded}, false},
		{[]codes.Code{codes.DeadlineExceeded, codes.Unavailable}, true},
		{[]codes.Code{codes.Unavailable}, true},
	}
	for _, tst := range tests {
		b := OnCodes(tst.c, Backoff{})
		if _, retry := b.Retry(context.Background(), apiErr); retry != tst.retry {
			t.Errorf("retriable codes: %v, error code: %s, retry: %t, want %t", tst.c, grpc.Code(apiErr), retry, tst.retry)
		}
	}
}

func TestOnRetryableError(t *testing.T) {

	tests := []struct {
		err   error
		retry bool
	}{
		{nil, false},
		{RetryableError(errors.New("Retry!")), true},
		{RetryOnResponse(&http.Response{StatusCode: 500}), true},
		{errors.New("Retry!"), false},
		{grpc.Errorf(codes.Internal, "No Retry."), false},
	}
	for _, tst := range tests {
		b := OnCodes(nil, Backoff{})
		if _, retry := b.Retry(context.Background(), tst.err); retry != tst.retry {
			t.Errorf("retry: %t, want %t", retry, tst.retry)
		}
	}
}

func TestErrorsInCode(t *testing.T) {
	tests := []struct {
		err           error
		retry         bool
		c             []codes.Code
		cancelContext bool
	}{
		{grpc.Errorf(codes.Canceled, "Retry."), false, []codes.Code{codes.Canceled}, true},
		{grpc.Errorf(codes.Canceled, "Retry."), true, []codes.Code{codes.Canceled}, false},
		{grpc.Errorf(codes.Canceled, "No Retry."), false, nil, false},
		{grpc.Errorf(codes.Internal, "No Retry."), false, nil, false},
		{grpc.Errorf(codes.Canceled, "Retry."), true, []codes.Code{codes.Canceled}, false},
		{RetryOnResponse(&http.Response{StatusCode: 500}), true, []codes.Code{codes.Canceled}, false},
	}

	for _, tst := range tests {
		ctx := context.Background()
		var cancel context.CancelFunc
		if tst.cancelContext {
			ctx, cancel = context.WithCancel(ctx)
			cancel()
		}

		b := OnCodes(tst.c, Backoff{})
		if _, retry := b.Retry(ctx, tst.err); retry != tst.retry {
			t.Errorf("retry: %t, want %t", retry, tst.retry)
		}
	}
}
