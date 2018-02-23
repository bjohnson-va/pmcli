package lb

import (
	"google.golang.org/grpc/naming"
	"context"
	"fmt"
	"github.com/vendasta/gosdks/logging"
)

// PoolResolver provides a fixed list of addresses to load balance between
// and does not provide further updates.
type PoolResolver struct {
	dialOpt *DialSettings
	ch      chan []*naming.Update
	watcher *K8SPeerWatcher
}

// NewPoolResolver returns a PoolResolver
// This is an EXPERIMENTAL API and may be changed or removed in the future.
func NewPoolResolver(o *DialSettings) (*PoolResolver, error) {
	w, err := NewK8SWatcher(o.Namespace, o.Labels)
	if err != nil {
		return nil, err
	}
	go w.WatchPods()
	return &PoolResolver{dialOpt: o, watcher: w}, nil
}

// Resolve returns a Watcher for the endpoint defined by the DialSettings
// provided to NewPoolResolver.
func (r *PoolResolver) Resolve(target string) (naming.Watcher, error) {
	return r, nil
}

var podUIDToIP = map[string]string{}

// Next returns a static list of updates on the first call,
// and blocks indefinitely until Close is called on subsequent calls.
func (r *PoolResolver) Next() ([]*naming.Update, error) {

	for {
		select {
		case podAdded := <-r.watcher.podsAdded:
			addr := fmt.Sprintf("%s:10002", podAdded.podIP)
			podUIDToIP[podAdded.UID] = podAdded.podIP
			logging.Debugf(context.Background(), "Added new vstore pod %s", addr)
			return []*naming.Update{{Op: naming.Add, Addr: addr, Metadata: podAdded.UID}}, nil
		case podDeleted := <-r.watcher.podsDeleted:
			podIP := podUIDToIP[podDeleted.UID]
			delete(podUIDToIP, podDeleted.UID)
			addr := fmt.Sprintf("%s:10002", podIP)
			logging.Debugf(context.Background(), "Removed vstore pod %s", addr)
			return []*naming.Update{{Op: naming.Delete, Addr: addr, Metadata: podDeleted.UID}}, nil
		}
	}
}

// Closes the pool
func (r *PoolResolver) Close() {
	close(r.ch)
}

// DialSettings holds information needed to establish a connection with vStore.
type DialSettings struct {
	Namespace string
	Labels    map[string]string
}
