package vstore

import (
	"strconv"
	"sync"

	"strings"

	"sort"

	"github.com/vendasta/gosdks/pb/vstorepb"
	"github.com/vendasta/gosdks/pubsub"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

// NewHappyPathStub returns an in memory validation that has little to no validation but supports creating
// namespaces/kinds as well as getting/creating/updating entities in vstore. It does not do any secondary indexing.
func NewHappyPathStub() *HappyPathStub {
	return &HappyPathStub{
		entities: map[string][]*entity{},
		kinds:    map[string][]string{},
	}
}

// HappyPathStub implements vstore.Interface
type HappyPathStub struct {
	entities   map[string][]*entity
	namespaces []string
	kinds      map[string][]string
	sync.Mutex
}

type entity struct {
	m  Model
	ks *KeySet
}

type entities []*entity

func (e entities) Len() int {
	return len(e)
}

func (e entities) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func (e entities) Less(i, j int) bool {
	lenJ := len(e[j].ks.keys)
	for n, si := range e[i].ks.keys {
		if n == lenJ {
			return false
		}
		sj := e[j].ks.keys[n]
		if si < sj {
			return true
		}
	}
	if len(e[i].ks.keys) < lenJ {
		return true
	}
	return false
}

// GetMulti stub
func (vs *HappyPathStub) GetMulti(ctx context.Context, keysets []*KeySet) ([]Model, error) {
	vs.Lock()
	defer vs.Unlock()

	toReturn := []Model{}
	for _, ks := range keysets {
		kspb := ks.ToKeySetPB()
		if !StringInSlice(kspb.Namespace, vs.namespaces) {
			return nil, grpc.Errorf(codes.InvalidArgument, "Namespace doesnt exist")
		}
		if !StringInSlice(kspb.Kind, vs.kinds[kspb.Namespace]) {
			return nil, grpc.Errorf(codes.InvalidArgument, "Kind doesnt exist")
		}
		found := false
		for _, existing := range vs.entities[kspb.Namespace+kspb.Kind] {
			if StringsEqual(existing.ks.keys, ks.keys) {
				found = true
				toReturn = append(toReturn, clone(existing.ks.namespace, existing.ks.kind, existing.m))
				break
			}
		}
		if !found {
			toReturn = append(toReturn, nil)
		}
	}
	return toReturn, nil
}

// Get stub
func (vs *HappyPathStub) Get(ctx context.Context, ks *KeySet) (Model, error) {
	r, err := vs.GetMulti(ctx, []*KeySet{ks})
	if err != nil {
		return nil, err
	}
	return r[0], err
}

func (vs *HappyPathStub) isFilterMatch(lo *lookupOption, keys []string) bool {
	if len(lo.filters) > 0 {
		for i, prefix := range lo.filters {
			if len(keys) < i+1 {
				return false
			} else if lo.partialFilter && i == len(lo.filters)-1 {
				if !strings.HasPrefix(keys[i], prefix) {
					return false
				}
			} else if keys[i] != prefix {
				return false
			}
		}
	}
	if len(lo.beginFilters) > 0 {
		for i2, prefix2 := range lo.beginFilters {
			if len(keys) < i2+1 {
				return false
			} else if lo.partialFilter && i2 == len(lo.beginFilters)-1 {
				if !strings.HasPrefix(keys[i2], prefix2) {
					return false
				}
			} else if keys[i2] < prefix2 {
				return false
			}
		}
	}
	if len(lo.endFilters) > 0 {
		for i3, prefix3 := range lo.endFilters {
			if len(keys) < i3+1 {
				return false
			} else if lo.partialFilter && i3 == len(lo.endFilters)-1 {
				if !strings.HasPrefix(keys[i3], prefix3) {
					return false
				}
			} else if keys[i3] > prefix3 {
				return false
			} else if i3+1 == len(lo.endFilters) && keys[i3] == prefix3 {
				return false
			}
		}
	}
	return true
}

// Lookup stub
func (vs *HappyPathStub) Lookup(ctx context.Context, namespace, kind string, opts ...LookupOption) (*LookupResult, error) {
	vs.Lock()
	defer vs.Unlock()

	if !StringInSlice(namespace, vs.namespaces) {
		return nil, grpc.Errorf(codes.InvalidArgument, "Namespace doesnt exist")
	}
	if !StringInSlice(kind, vs.kinds[namespace]) {
		return nil, grpc.Errorf(codes.InvalidArgument, "Kind doesnt exist")
	}
	var err error

	lo := &lookupOption{pageSize: 10, cursor: ""}
	for _, opt := range opts {
		opt(lo)
	}

	// set index to start scanning, based on cursor
	startIdx := 0
	if lo.cursor != "" {
		startIdx, err = strconv.Atoi(lo.cursor)
		if err != nil {
			return nil, grpc.Errorf(codes.InvalidArgument, "Invalid cursor")
		}
	}

	hasMore := false
	nextCursor := ""
	toReturn := []Model{}

	// The real vstore returns items sorted by "least key value"
	sorted := entities{}
	for _, e := range vs.entities[namespace+kind] {
		sorted = append(sorted, e)
	}
	sort.Sort(sorted)

	for i, e := range sorted {
		// skip entries until we reach the correct startIdx defined by the cursor
		if i < startIdx {
			continue
		}

		// if there are filters, we need to actually see if the keys of the entities match our filters, else just add the next entity
		if vs.isFilterMatch(lo, e.ks.keys) {
			toReturn = append(toReturn, e.m)
		}

		// if this happens, there is more than a single page of results remaining, so set the information the client
		// will need to iterate over the next 1<=n<=lo.pageSize results
		if int64(len(toReturn)) > lo.pageSize {
			// remove the last item, we need to check to see if there are actually more results
			toReturn = toReturn[:len(toReturn)-1]
			hasMore = true
			nextCursor = strconv.FormatInt(int64(i), 10)
			break
		}
	}
	return &LookupResult{Results: toReturn, NextCursor: nextCursor, HasMore: hasMore}, nil
}

func (vs *HappyPathStub) Scan(ctx context.Context, namespace, kind string, cb func(m Model) bool, opts ...ScanOption) error {
	// Leverage lookup to implement scan

	scanOpts := scanOption{}
	for _, opt := range opts {
		opt(&scanOpts)
	}
	cursor := ""
	for {
		lookupOption := []LookupOption{}
		if len(scanOpts.filters) > 0 {
			if scanOpts.partialFilter {
				lookupOption = append(lookupOption, PartialFilter(scanOpts.filters))
			} else {
				lookupOption = append(lookupOption, Filter(scanOpts.filters))
			}
		} else if len(scanOpts.beginFilters) > 0 && len(scanOpts.endFilters) > 0 {
			lookupOption = append(lookupOption, RangeFilter(scanOpts.beginFilters, scanOpts.endFilters))
		}
		lookupOption = append(lookupOption, Cursor(cursor))

		lr, err := vs.Lookup(ctx, namespace, kind, lookupOption...)
		if err != nil {
			return err
		}
		for _, e := range lr.Results {
			ok := cb(e)
			if !ok {
				return nil
			}
		}
		if !lr.HasMore || len(lr.Results) == 0 {
			return nil
		}
		cursor = lr.NextCursor
	}
}

// Transaction stub
func (vs *HappyPathStub) Transaction(ctx context.Context, ks *KeySet, tx func(Transaction, Model) error, opts ...TransactionOption) error {
	o := &txOpts{}
	for _, apply := range opts {
		apply(o)
	}
	// TODO: Support proto transaction logic in stub - dwalker
	if o.pbTransactionCallback != nil {
		return nil
	}

	t := &transaction{}
	m, err := vs.Get(ctx, ks)
	if err != nil {
		return err
	}

	vs.Lock()
	defer vs.Unlock()

	err = tx(t, m)
	if err != nil {
		return err
	}
	if t.toSave == nil {
		return nil
	}
	index := -1
	for i, existing := range vs.entities[ks.namespace+ks.kind] {
		if StringsEqual(existing.ks.keys, ks.keys) {
			index = i
			break
		}
	}
	e := &entity{m: clone(ks.namespace, ks.kind, t.toSave), ks: ks}
	if index == -1 {
		vs.entities[ks.namespace+ks.kind] = append(vs.entities[ks.namespace+ks.kind], e)
	} else {
		vs.entities[ks.namespace+ks.kind][index] = e
	}

	return nil
}

// CreateNamespace stub
func (vs *HappyPathStub) CreateNamespace(ctx context.Context, namespace string, authorizedServiceAccounts []string) error {
	vs.Lock()
	defer vs.Unlock()

	if StringInSlice(namespace, vs.namespaces) {
		return grpc.Errorf(codes.AlreadyExists, "Namespace already exists")
	}
	vs.namespaces = append(vs.namespaces, namespace)
	return nil
}

// UpdateNamespace stub
func (vs *HappyPathStub) UpdateNamespace(ctx context.Context, namespace string, authorizedServiceAccounts []string) error {
	vs.Lock()
	defer vs.Unlock()

	if !StringInSlice(namespace, vs.namespaces) {
		return grpc.Errorf(codes.NotFound, "Namespace doesnt exist.")
	}
	return nil
}

// DeleteNamespace stub
func (vs *HappyPathStub) DeleteNamespace(ctx context.Context, namespace string) error {
	vs.Lock()
	defer vs.Unlock()
	if !StringInSlice(namespace, vs.namespaces) {
		return grpc.Errorf(codes.NotFound, "Namespace doesnt exist.")
	}
	namespaces := []string{}
	for _, ns := range vs.namespaces {
		if ns == namespace {
			continue
		}
		namespaces = append(namespaces, ns)
	}
	vs.namespaces = namespaces
	delete(vs.kinds, namespace)
	return nil
}

// CreateKind stub
func (vs *HappyPathStub) CreateKind(ctx context.Context, schema *Schema) error {
	vs.Lock()
	defer vs.Unlock()
	if !StringInSlice(schema.Namespace, vs.namespaces) {
		return grpc.Errorf(codes.NotFound, "Namespace doesnt exist.")
	}

	if StringInSlice(schema.Kind, vs.kinds[schema.Namespace]) {
		return grpc.Errorf(codes.AlreadyExists, "Kind already exists")
	}
	vs.kinds[schema.Namespace] = append(vs.kinds[schema.Namespace], schema.Kind)
	return nil
}

// UpdateKind stub
func (vs *HappyPathStub) UpdateKind(ctx context.Context, schema *Schema) error {
	vs.Lock()
	defer vs.Unlock()
	if !StringInSlice(schema.Namespace, vs.namespaces) {
		return grpc.Errorf(codes.NotFound, "Namespace doesnt exist.")
	}

	if !StringInSlice(schema.Kind, vs.kinds[schema.Namespace]) {
		return grpc.Errorf(codes.NotFound, "Kind doenst exist.")
	}
	return nil
}

// GetKind stub
func (vs *HappyPathStub) GetKind(ctx context.Context, namespace string, kind string) (*vstorepb.GetKindResponse, error) {
	vs.Lock()
	defer vs.Unlock()
	return nil, nil
}

// DeleteKind stub
func (vs *HappyPathStub) DeleteKind(ctx context.Context, namespace, kind string) error {
	vs.Lock()
	defer vs.Unlock()
	if !StringInSlice(namespace, vs.kinds[namespace]) {
		return grpc.Errorf(codes.NotFound, "Kind doesnt exist.")
	}
	delete(vs.kinds, namespace)
	return nil
}

// RegisterKind stub
func (vs *HappyPathStub) RegisterKind(ctx context.Context, namespace, kind string, serviceAccounts []string, model Model) (*vstorepb.GetKindResponse, error) {
	vs.Lock()
	RegisterModel(namespace, kind, model)
	vs.Unlock()

	err := vs.UpdateKind(ctx, model.Schema())
	if err == nil {
		sch, err := vs.GetKind(ctx, namespace, kind)
		if err != nil {
			return nil, err
		}
		return sch, nil
	}
	if grpc.Code(err) != codes.NotFound {
		return nil, err
	}

	// Kind doesn't exist yet; create it
	err = vs.CreateNamespace(ctx, namespace, serviceAccounts)
	if err != nil && grpc.Code(err) != codes.AlreadyExists {
		return nil, err
	}

	err = vs.CreateKind(ctx, model.Schema())
	if err != nil {
		return nil, err
	}
	return &vstorepb.GetKindResponse{}, nil
}

// RegisterSubscriptionCallback stub
func (vs *HappyPathStub) RegisterSubscriptionCallback(ctx context.Context, namespace, kind, indexID, subscriptionName string, handler MessageHandler, cancelFunc context.CancelFunc, opts ...pubsub.WorkerOption) error {
	vs.Lock()
	defer vs.Unlock()

	return nil
}

// GetSecondaryIndexName stub
func (vs *HappyPathStub) GetSecondaryIndexName(ctx context.Context, namespace, kind string, indexID string) (string, error) {
	vs.Lock()
	defer vs.Unlock()

	return "", nil
}

// StringInSlice determines if the given string is in the provided slice.
func StringInSlice(target string, list []string) bool {
	for _, candidate := range list {
		if candidate == target {
			return true
		}
	}
	return false
}

// StringsEqual determines if the list of strings are equal.
func StringsEqual(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func clone(namespace, kind string, m Model) Model {
	s, err := ModelToStructPB(m)
	if err != nil {
		panic(err.Error())
	}
	mClone, err := StructPBToModel(namespace, kind, s)
	if err != nil {
		panic(err.Error())
	}
	return mClone
}

// SetTestMessageModel sets the model on a vstore message so we can construct a message stub
func (m *message) SetTestMessageModel(model Model) {
	m.model = model
}

// SetTestMessageVersion sets the version on a vstore message so we can construct a message stub
func (m *message) SetTestMessageVersion(version int64) {
	m.version = version
}
