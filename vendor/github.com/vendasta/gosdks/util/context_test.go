package util

import (
	"testing"

	"time"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
)

func Test_NewContextWithEmptyMetadata(t *testing.T) {
	ctx := NewContext(context.TODO())
	assert.NotNil(t, ctx)
}

func Test_NewContextWithTracingMetadata(t *testing.T) {
	ctx := metadata.NewIncomingContext(context.TODO(), metadata.New(map[string]string{"x-cloud-trace-context": "trace-id"}))
	ctx = NewContext(ctx)
	assert.NotNil(t, ctx)
	md, _ := metadata.FromOutgoingContext(ctx)
	assert.Equal(t, []string{"trace-id"}, md["x-cloud-trace-context"])
}

func Test_NewContextWithOutTimeout(t *testing.T) {
	ctx := NewContext(context.Background())
	deadline, deadlineSet := ctx.Deadline()
	assert.False(t, deadline.IsZero())
	assert.True(t, deadlineSet)
}

func Test_NewContextWithTimeout(t *testing.T) {
	ctx := NewContext(context.TODO(), WithTimeout(time.Second))
	deadline, deadlineSet := ctx.Deadline()
	assert.False(t, deadline.IsZero())
	assert.True(t, deadlineSet)
}
