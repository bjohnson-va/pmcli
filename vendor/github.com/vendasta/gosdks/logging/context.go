package logging

import (
	"sync"
	"context"
	"fmt"
	"github.com/pborman/uuid"
)

type tagsDataKey struct{}

type tagsData struct {
	mu sync.Mutex

	// Additional labels to add to the GKE request
	tags map[string]string
}

func (cd *tagsData) addTag(key, value string) {
	cd.mu.Lock()
	defer cd.mu.Unlock()

	key = fmt.Sprintf("vendasta.com/%s", key)

	cd.tags[key] = value
}

func (cd *tagsData) getLabels() map[string]string {
	return cd.tags
}

// TaggedContext returns a context.Context that will accept tags using Tag(ctx, k, v) and push them to Cloud Logging as labels
// This function will not clear old labels, so it is safe to call multiple times on the same ctx.
// Example:
// 	taggableCtx := logging.NewTaggedContext(ctx)
// 	logging.Tag(taggableCtx, "key", "value")
func NewTaggedContext(ctx context.Context) context.Context {
	cd, ok := taggedDataFromContext(ctx)
	if !ok {
		cd = &tagsData{
			tags: map[string]string{},
		}
	}
	return context.WithValue(ctx, tagsDataKey{}, cd)
}

func taggedDataFromContext(ctx context.Context) (md *tagsData, ok bool) {
	md, ok = ctx.Value(tagsDataKey{}).(*tagsData)
	return
}

// NewWorkerContext adds a worker id tag to the context with a unique id
func NewWorkerContext(ctx context.Context) context.Context {
	wrkCtx := NewTaggedContext(ctx)
	cd, _ := taggedDataFromContext(wrkCtx)
	cd.addTag("worker_id", uuid.New())
	return wrkCtx
}
