package vax

import (
	"sync/atomic"

	"golang.org/x/net/context"
)

// A Hedge provides an API for invoking multiple calls that process the same thing but through a different route.
// This allows for racing go-routines, reducing your tail latency, improving the availability by increasing redundancy,
// and rolling out new functionality with increased safety.
type Hedge struct {
	parentCtx context.Context
	cancel    func()

	res   chan struct{}
	err   chan error
	total int64
}

// HedgeWithContext creates a new Hedge and an associated Context derived from ctx.
//
// The derived Context is canceled the first time a function passed to Go
// returns a nil error or the first time Wait returns, whichever occurs
// first.
func HedgeWithContext(ctx context.Context) (*Hedge, context.Context) {
	childCtx, cancel := context.WithCancel(ctx)
	return &Hedge{
		parentCtx: ctx,
		cancel:    cancel,
		res:       make(chan struct{}),
		err:       make(chan error),
		total:     0,
	}, childCtx
}

// Wait blocks until all function calls from the Go method have error'd, then
// returns the first non-nil error (if any) from them, OR a single function call results in success.
func (h *Hedge) Wait() error {
	defer h.cancel()
	if h.total == 0 {
		return nil
	}
	var errs []error
	var success bool
	for {
		if success {
			break
		}
		// all goroutines resulted in err
		if len(errs) == int(h.total) {
			return errs[0]
		}
		select {
		case <-h.res:
			success = true
		case <-h.parentCtx.Done():
			return h.parentCtx.Err()
		case err := <-h.err:
			errs = append(errs, err)
		}
	}
	return nil
}

// Go calls the given function in a new goroutine.
//
// The first call to return a nil error cancels the hedge; The first error will be returned if all fail.
func (h *Hedge) Go(f func() error) {
	atomic.AddInt64(&h.total, 1)
	go func() {
		err := f()
		if err != nil {
			h.err <- err
		} else {
			h.res <- struct{}{}
		}
	}()
}
