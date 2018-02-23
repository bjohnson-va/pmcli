package vax

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Hedge(t *testing.T) {
	tsts := []struct {
		calls []func() error
		err   error
		dsc   string
	}{
		{
			calls: []func() error{},
			dsc:   "No deferred calls will return no error",
		},
		{
			calls: []func() error{
				func() error {
					return nil
				},
			},
			dsc: "Single success fn will return in no error",
		},
		{
			calls: []func() error{
				func() error {
					return nil
				},
				func() error {
					return errors.New("Broked!")
				},
			},
			dsc: "Single success fn will return in no error",
		},
		{
			calls: []func() error{
				func() error {
					return errors.New("Broked!")
				},
				func() error {
					return errors.New("Broked!")
				},
			},
			err: errors.New("Broked!"),
			dsc: "All failures will result in an error",
		},
		{
			calls: []func() error{
				func() error {
					return errors.New("Broked!")
				},
			},
			err: errors.New("Broked!"),
			dsc: "Single failure will result in an error",
		},
	}
	for _, tst := range tsts {
		ctx, cancel := context.WithCancel(context.Background())
		h, childCtx := HedgeWithContext(ctx)
		for _, c := range tst.calls {
			h.Go(c)
		}
		err := h.Wait()
		assert.Equal(t, tst.err, err)
		assert.NotNil(t, childCtx.Err())
		cancel()
	}
}
