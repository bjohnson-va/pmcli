package vstore

import (
	"errors"
	"fmt"
	"io"
	"time"

	gcloud_pubsub "cloud.google.com/go/pubsub"
	"github.com/vendasta/gosdks/pubsub"

	"math/rand"

	"github.com/vendasta/gosdks/logging"
	"github.com/vendasta/gosdks/pb/vstorepb"
	"github.com/vendasta/gosdks/vax"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ProtoStruct is the wire format for a vstore struct
type ProtoStruct = vstorepb.Struct

// ProtoStruct is the wire format for a vstore entity
type ProtoEntity = vstorepb.Entity

//ClientOption modifies the settings and behaviour of the vstore client.
type ClientOption func(*clientOption)

type clientOption struct {
	client               vstorepb.VStoreClient
	adminClient          vstorepb.VStoreAdminClient
	pubsubClient         pubsub.Client
	environment          *env
	dialOptions          []grpc.DialOption
	useInternalTransport bool
}

// Environment sets the environment vStore should connect to. Either `Environment` or `Client` must be supplied.
func Environment(e env) ClientOption {
	return func(c *clientOption) {
		c.environment = &e
	}
}

// Client manually sets the client for vStore to use. Either `Environment` or `Client` must be supplied.
func Client(client vstorepb.VStoreClient) ClientOption {
	return func(c *clientOption) {
		c.client = client
	}
}

// AdminClient manually sets the admin client for vStore to use.
func AdminClient(client vstorepb.VStoreAdminClient) ClientOption {
	return func(c *clientOption) {
		c.adminClient = client
	}
}

// PubsubClient manually sets the pubsub client for vStore to use.
func PubsubClient(client pubsub.Client) ClientOption {
	return func(c *clientOption) {
		c.pubsubClient = client
	}
}

// GRPCDialOptions manually sets the dial options for vStore to use when establishing connections.
func GRPCDialOptions(dialOptions ...grpc.DialOption) ClientOption {
	return func(c *clientOption) {
		c.dialOptions = dialOptions
	}
}

// WithInternalTransport tells vStore to use an experimental internal transport method that skips the load balancer.
func WithInternalTransport() ClientOption {
	return func(c *clientOption) {
		c.useInternalTransport = true
	}
}

// New returns a new vStore client
func New(cs ...ClientOption) (Interface, error) {
	co := clientOption{}
	for _, c := range cs {
		c(&co)
	}
	if co.environment == nil {
		co.environment = Env()
	}
	if co.client == nil {
		var err error
		co.client, co.adminClient, err = newClient(*co.environment, co.useInternalTransport, co.dialOptions...)
		if err != nil {
			return nil, err
		}
	}
	if co.pubsubClient == nil {
		var err error
		co.pubsubClient, err = pubsub.NewGooglePubsubClient(context.Background())
		if err != nil {
			return nil, err
		}
	}

	return &vStore{client: co.client, adminClient: co.adminClient, pubsubClient: co.pubsubClient}, nil
}

type lookupOption struct {
	pageSize      int64
	cursor        string
	filters       []string
	beginFilters  []string
	endFilters    []string
	partialFilter bool
}

func (l *lookupOption) FilterProto() *vstorepb.LookupFilter {
	// range filter
	if len(l.beginFilters) > 0 && len(l.endFilters) > 0 {
		return &vstorepb.LookupFilter{
			Filters: &vstorepb.LookupFilter_RangeFilter{
				RangeFilter: &vstorepb.RangeFilter{
					Begin: l.beginFilters,
					End:   l.endFilters,
				},
			},
		}
	}

	if len(l.filters) == 0 {
		return nil
	}

	// key filter
	return &vstorepb.LookupFilter{
		Filters: &vstorepb.LookupFilter_KeyFilter{
			KeyFilter: &vstorepb.KeyFilter{
				Keys:   l.filters,
				Prefix: l.partialFilter,
			},
		},
	}
}

//LookupOption augments the behavior of the lookup API
type LookupOption func(*lookupOption)

//PageSize sets a maximum page size on the lookup request
func PageSize(pageSize int64) LookupOption {
	return func(l *lookupOption) {
		l.pageSize = pageSize
	}
}

//Cursor sets a cursor on the lookup request, enabling easy paging
func Cursor(cursor string) LookupOption {
	return func(l *lookupOption) {
		l.cursor = cursor
	}
}

// Filter sets a prefix filter on the lookup
// Entities in VStore:
// ['LIS-101', 'RVW-223']
// ['LIS-101', 'RVW-676']
// ['LIS-555', 'RVW-444']
// Applying a filter like Filter([]string{'LIS-101'}) would mean that the lookup api would return ['LIS-101', 'RVW-223'] and ['LIS-101', 'RVW-676'] but not ['LIS-555', 'RVW-444']
func Filter(filters []string) LookupOption {
	return func(l *lookupOption) {
		l.filters = filters
	}
}

// PartialFilter sets a partial prefix filter on the lookup
// Entities in VStore:
// ['LIS-101', 'RVW-223']
// ['LIS-101', 'RVW-676']
// ['LIS-101', 'RVW-244']
// ['LIS-201', 'RVW-245']
// Applying a filter like PartialFilter([]string{'LIS-101', "RVW-2"}) would mean that the lookup api would return ['LIS-101', 'RVW-223'] and ['LIS-101', 'RVW-244'] but not ['LIS-101', 'RVW-676'] or ['LIS-201', 'RVW-245']
func PartialFilter(filters []string) LookupOption {
	return func(l *lookupOption) {
		l.filters = filters
		l.partialFilter = true
	}
}

// RangeFilter sets a range filter on the lookup
// Entities in VStore:
// ['LIS-101', 'RVW-223']
// ['LIS-101', 'RVW-244']
// ['LIS-101', 'RVW-676']
// ['LIS-201', 'RVW-245']
// Applying a filter like RangeFilter([]string{'LIS-101', "RVW-244"}, []string{'LIS-101', "RVW-900"}) would mean that the lookup api would return ['LIS-101', 'RVW-244'] and ['LIS-101', 'RVW-676'] but not ['LIS-101', 'RVW-223'] or ['LIS-201', 'RVW-245']
func RangeFilter(begin, end []string) LookupOption {
	return func(l *lookupOption) {
		l.beginFilters = begin
		l.endFilters = end
	}
}

// ScanOption augments the behaviour of the scan API
type ScanOption func(*scanOption)

type scanOption struct {
	filters       []string
	beginFilters  []string
	endFilters    []string
	partialFilter bool
	protoCallback func(*ProtoEntity) bool
}

func (s *scanOption) FilterProto() *vstorepb.LookupFilter {
	// range filter
	if len(s.beginFilters) > 0 && len(s.endFilters) > 0 {
		return &vstorepb.LookupFilter{
			Filters: &vstorepb.LookupFilter_RangeFilter{
				RangeFilter: &vstorepb.RangeFilter{
					Begin: s.beginFilters,
					End:   s.endFilters,
				},
			},
		}
	}

	if len(s.filters) == 0 {
		return nil
	}

	// key filter
	return &vstorepb.LookupFilter{
		Filters: &vstorepb.LookupFilter_KeyFilter{
			KeyFilter: &vstorepb.KeyFilter{
				Keys:   s.filters,
				Prefix: s.partialFilter,
			},
		},
	}
}

// WithProto executes the scan given the provided callback operating over a ProtoStruct.
// This allows the scan to be agnostic of the entity schema and work in terms of the raw proto format.
// You probably don't need to use this.
func WithProto(f func(*ProtoEntity) bool) ScanOption {
	return func(o *scanOption) {
		o.protoCallback = f
	}
}

// Filter sets a prefix filter on the scan
// Entities in VStore:
// ['LIS-101', 'RVW-223']
// ['LIS-101', 'RVW-676']
// ['LIS-555', 'RVW-444']
// Applying a filter like Filter([]string{'LIS-101'}) would mean that the scan api would return ['LIS-101', 'RVW-223'] and ['LIS-101', 'RVW-676'] but not ['LIS-555', 'RVW-444']
func ScanFilter(filters []string) ScanOption {
	return func(s *scanOption) {
		s.filters = filters
	}
}

// ScanPartialFilter sets a partial prefix filter on the scan
// Entities in VStore:
// ['LIS-101', 'RVW-223']
// ['LIS-101', 'RVW-676']
// ['LIS-101', 'RVW-244']
// ['LIS-201', 'RVW-245']
// Applying a filter like ScanPartialFilter([]string{'LIS-101', "RVW-2"}) would mean that the scan would start at ['LIS-101', 'RVW-223'] then go to ['LIS-101', 'RVW-244'] but not ['LIS-101', 'RVW-676'] or ['LIS-201', 'RVW-245']
func ScanPartialFilter(filters []string) ScanOption {
	return func(s *scanOption) {
		s.filters = filters
		s.partialFilter = true
	}
}

// ScanRangeFilter sets a range filter on the scan
// Entities in VStore:
// ['LIS-101', 'RVW-223']
// ['LIS-101', 'RVW-244']
// ['LIS-101', 'RVW-676']
// ['LIS-201', 'RVW-245']
// Applying a filter like ScanRangeFilter([]string{'LIS-101', "RVW-244"}, []string{'LIS-101', "RVW-900"}) would mean that the lookup api scan through ['LIS-101', 'RVW-244'] and ['LIS-101', 'RVW-676'] but not ['LIS-101', 'RVW-223'] or ['LIS-201', 'RVW-245']
func ScanRangeFilter(begin, end []string) ScanOption {
	return func(s *scanOption) {
		s.beginFilters = begin
		s.endFilters = end
	}
}

// Interface is the complete VStore client interface
type Interface interface {
	GetMulti(context.Context, []*KeySet) ([]Model, error)
	Get(context.Context, *KeySet) (Model, error)
	Lookup(ctx context.Context, namespace, kind string, opts ...LookupOption) (*LookupResult, error)
	Transaction(context.Context, *KeySet, func(Transaction, Model) error, ...TransactionOption) error
	Scan(ctx context.Context, namespace, kind string, cb func(m Model) bool, opts ...ScanOption) error

	CreateNamespace(ctx context.Context, namespace string, authorizedServiceAccounts []string) error
	UpdateNamespace(ctx context.Context, namespace string, authorizedServiceAccounts []string) error
	DeleteNamespace(ctx context.Context, namespace string) error

	CreateKind(ctx context.Context, schema *Schema) error
	UpdateKind(ctx context.Context, schema *Schema) error
	GetKind(ctx context.Context, namespace string, kind string) (*vstorepb.GetKindResponse, error)
	DeleteKind(ctx context.Context, namespace, kind string) error

	RegisterKind(ctx context.Context, namespace, kind string, serviceAccounts []string, model Model) (*vstorepb.GetKindResponse, error)
	GetSecondaryIndexName(ctx context.Context, namespace, kind string, indexID string) (string, error)

	RegisterSubscriptionCallback(ctx context.Context, namespace, kind, indexID, subscriptionName string, handler MessageHandler, cancelFunc context.CancelFunc, opts ...pubsub.WorkerOption) error
}

// defaultRetryCallOptions controls the errors that we will automatically retry on. This is due to the case where the
// server has given us an error that is deemed retry-able.
var defaultRetryCallOptions = vax.WithRetry(func() vax.Retryer {
	return vax.OnCodes([]codes.Code{
		codes.DeadlineExceeded,
		codes.Unavailable,
		codes.Unknown,
		codes.Canceled,
	}, vax.Backoff{
		Initial:    10 * time.Millisecond,
		Max:        300 * time.Millisecond,
		Multiplier: 3,
	})
})

// Implements vstore.Interface
type vStore struct {
	client       vstorepb.VStoreClient
	adminClient  vstorepb.VStoreAdminClient
	pubsubClient pubsub.Client
	user         *UserInfo
}

// GetMulti returns a set of rows for the given set of keysets.
func (v *vStore) GetMulti(ctx context.Context, keysets []*KeySet) ([]Model, error) {
	kspbs := make([]*vstorepb.KeySet, len(keysets))
	for i, ks := range keysets {
		kspbs[i] = ks.ToKeySetPB()
	}

	var res *vstorepb.GetResponse
	err := vax.Invoke(ctx, func(ctx context.Context, settings vax.CallSettings) error {
		var err error
		res, err = v.client.Get(ctx, &vstorepb.GetRequest{KeySets: kspbs}, grpc.FailFast(false))
		return err
	}, defaultRetryCallOptions)
	if err != nil {
		return nil, err
	}

	var models []Model
	for _, e := range res.Entities {
		var m Model
		if e.Entity != nil {
			m, err = StructPBToModel(e.Entity.Namespace, e.Entity.Kind, e.Entity.Values)
			if err != nil {
				return nil, err
			}
		}
		models = append(models, m)
	}
	return models, nil
}

// Get returns a single row from vStore by its KeySet.
func (v *vStore) Get(ctx context.Context, keyset *KeySet) (Model, error) {
	entities, err := v.GetMulti(ctx, []*KeySet{keyset})
	if err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Lookup supports fetching a page of rows from vStore for a single namespace/kind.
func (v *vStore) Lookup(ctx context.Context, namespace, kind string, opts ...LookupOption) (*LookupResult, error) {
	options := lookupOption{pageSize: 10}
	for _, opt := range opts {
		opt(&options)
	}
	req := &vstorepb.LookupRequest{
		Namespace: namespace,
		Kind:      kind,
		PageSize:  options.pageSize,
		Cursor:    options.cursor,
		Filter:    options.FilterProto(),
	}
	var r *vstorepb.LookupResponse
	err := vax.Invoke(ctx, func(ctx context.Context, settings vax.CallSettings) error {
		var err error
		r, err = v.client.Lookup(ctx, req, grpc.FailFast(false))
		return err
	}, defaultRetryCallOptions)
	if err != nil {
		return nil, err
	}

	var m = make([]Model, len(r.Entities))
	for i, e := range r.Entities {
		mod, err := StructPBToModel(e.Entity.Namespace, e.Entity.Kind, e.Entity.Values)
		if err != nil {
			return nil, err
		}
		m[i] = mod
	}
	return &LookupResult{Results: m, NextCursor: r.NextCursor, HasMore: r.HasMore}, nil
}

// Scan supports streaming a set of results. This is considered experimental. Use at your own risk!
func (v *vStore) Scan(ctx context.Context, namespace, kind string, cb func(m Model) bool, opts ...ScanOption) error {
	options := scanOption{}
	for _, opt := range opts {
		opt(&options)
	}
	req := &vstorepb.ScanRequest{
		Namespace: namespace,
		Kind:      kind,
		Filter:    options.FilterProto(),
	}
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	var previousRowKey *vstorepb.KeySet

	return vax.Invoke(ctx, func(ctx context.Context, cs vax.CallSettings) error {
		if previousRowKey != nil {
			req.StartingKeySet = previousRowKey
		}

		sc, err := v.client.Scan(ctx, req, grpc.FailFast(false))
		if err != nil {
			return err
		}

		for {
			er, err := sc.Recv()
			if err != nil {
				if err == io.EOF {
					break
				}
				return err
			}

			previousRowKey = er.KeySet

			var ok bool
			if options.protoCallback != nil {
				ok = options.protoCallback(er.Entity)
			} else {
				m, err := StructPBToModel(er.Entity.Namespace, er.Entity.Kind, er.Entity.Values)
				if err != nil {
					return err
				}
				ok = cb(m)
			}
			if !ok {
				break
			}
		}
		return nil
	}, defaultRetryCallOptions)
}

// Transaction allows updating a single row in vStore.
func (v *vStore) Transaction(ctx context.Context, ks *KeySet, f func(Transaction, Model) error, opts ...TransactionOption) error {
	o := &txOpts{}
	for _, apply := range opts {
		apply(o)
	}

	// Only retry the transaction as a whole when there has been a transaction collision.
	txRetryOption := vax.WithRetry(func() vax.Retryer {
		return vax.OnCodes([]codes.Code{
			codes.FailedPrecondition,
			codes.AlreadyExists, // Race on Creates
		}, vax.Backoff{
			Initial:    time.Duration(rand.Int63n(50)) * time.Millisecond,
			Max:        300 * time.Millisecond,
			Multiplier: 3,
		})
	})
	return vax.Invoke(ctx, func(ctx context.Context, settings vax.CallSettings) error {
		return v.transactionHelper(
			ctx, ks, f, o,
		)
	}, txRetryOption)
}

func (v *vStore) transactionHelper(ctx context.Context, ks *KeySet, f func(Transaction, Model) error, opts *txOpts) error {
	tx := &transaction{}

	var res *vstorepb.GetResponse
	err := vax.Invoke(ctx, func(ctx context.Context, settings vax.CallSettings) error {
		var err error
		res, err = v.client.Get(ctx, &vstorepb.GetRequest{KeySets: []*vstorepb.KeySet{ks.ToKeySetPB()}}, grpc.FailFast(false))
		return err
	}, defaultRetryCallOptions)
	if err != nil {
		return err
	}

	// invoke client's callback
	entity := res.Entities[0]
	var m Model
	if opts.pbTransactionCallback != nil {
		err = opts.pbTransactionCallback(tx, entity.GetEntity().GetValues())
	} else {
		if entity.Entity != nil {
			m, err = StructPBToModel(entity.Entity.Namespace, entity.Entity.Kind, entity.Entity.Values)
			if err != nil {
				return err
			}
		}
		err = f(tx, m)
	}
	if err != nil {
		return err
	}
	if tx.toSave == nil && tx.pbToSave == nil {
		return nil
	}

	e := vstorepb.Entity{
		Namespace: ks.namespace,
		Kind:      ks.kind,
	}

	// set proto values
	if opts.pbTransactionCallback != nil {
		e.Values = tx.pbToSave
	} else {
		s, err := ModelToStructPB(tx.toSave)
		if err != nil {
			return err
		}
		e.Values = s
	}

	// figure out if the version should be 1, if there's already a version or model set, then it doesn't need to change
	if opts.pbTransactionCallback != nil {
		ver := entity.GetEntity().GetVersion()
		if ver == 0 {
			e.Version = 1
		}
	} else {
		if m == nil {
			e.Version = 1
		}
	}

	// if the version is 1, we want to attempt to create the entity
	if e.Version == 1 {
		return vax.Invoke(ctx, func(ctx context.Context, settings vax.CallSettings) error {
			_, err := v.client.Create(ctx, &vstorepb.CreateRequest{Entity: &e})
			return err
		}, defaultRetryCallOptions)
	}

	// otherwise, we want to attempt to update it
	e.Version = entity.Entity.Version
	return vax.Invoke(ctx, func(ctx context.Context, settings vax.CallSettings) error {
		_, err := v.client.Update(ctx, &vstorepb.UpdateRequest{Entity: &e})
		return err
	}, defaultRetryCallOptions)
}

// CreateNamespace
func (v *vStore) CreateNamespace(ctx context.Context, namespace string, authorizedServiceAccounts []string) error {
	if v.adminClient == nil {
		return errors.New("Admin client must be initialized.")
	}
	_, err := v.adminClient.CreateNamespace(ctx,
		&vstorepb.CreateNamespaceRequest{Namespace: namespace, AuthorizedServiceAccounts: authorizedServiceAccounts},
		grpc.FailFast(false),
	)
	return err
}

// UpdateNamespace allows the updating of authorized service accounts.
func (v *vStore) UpdateNamespace(ctx context.Context, namespace string, authorizedServiceAccounts []string) error {
	if v.adminClient == nil {
		return errors.New("Admin client must be initialized.")
	}
	_, err := v.adminClient.UpdateNamespace(ctx,
		&vstorepb.UpdateNamespaceRequest{Namespace: namespace, AuthorizedServiceAccounts: authorizedServiceAccounts},
	)
	return err
}

// DeleteNamespace removes the given namespace from vStore and all of its kinds.  This is a permanent process and
// can not be reversed.
func (v *vStore) DeleteNamespace(ctx context.Context, namespace string) error {
	if v.adminClient == nil {
		return errors.New("Admin client must be initialized.")
	}
	_, err := v.adminClient.DeleteNamespace(ctx,
		&vstorepb.DeleteNamespaceRequest{Namespace: namespace},
	)
	return err
}

// CreateKind makes a new kind in a specific namespace.
func (v *vStore) CreateKind(ctx context.Context, schema *Schema) error {
	if v.adminClient == nil {
		return errors.New("Admin client must be initialized.")
	}
	properties, err := PropertiesToPb(schema.Properties...)
	if err != nil {
		return err
	}
	_, err = v.adminClient.CreateKind(ctx,
		&vstorepb.CreateKindRequest{
			Namespace:         schema.Namespace,
			Kind:              schema.Kind,
			PrimaryKey:        schema.PrimaryKey,
			Properties:        properties,
			SecondaryIndexes:  SecondaryIndexesToPb(schema.SecondaryIndexes...),
			BackupConfig:      BackupConfigToPb(schema.BackupConfig),
			ExtendedKeyLength: schema.ExtendedKeyLength,
		}, grpc.FailFast(false))
	if err != nil {
		s, ok := status.FromError(err)
		if !ok {
			return err
		}
		if s.Code() == codes.ResourceExhausted {
			logging.Debugf(ctx, "Received Rate Limit error from VStore on creating kind %s-%s, continuing", schema.Namespace, schema.Kind)
			return nil
		}
		return err
	}
	return nil
}

// GetKind returns the kind by its namespace/name pair.
func (v *vStore) GetKind(ctx context.Context, namespace string, kind string) (*vstorepb.GetKindResponse, error) {
	if v.adminClient == nil {
		return nil, errors.New("Admin client must be initialized.")
	}
	r, err := v.adminClient.GetKind(ctx,
		&vstorepb.GetKindRequest{
			Namespace: namespace,
			Kind:      kind,
		}, grpc.FailFast(false))
	if err != nil {
		return nil, err
	}
	return r, nil
}

// UpdateKind supports the addition of new fields and updates any supported settings.
func (v *vStore) UpdateKind(ctx context.Context, schema *Schema) error {
	if v.adminClient == nil {
		return errors.New("Admin client must be initialized.")
	}
	properties, err := PropertiesToPb(schema.Properties...)
	if err != nil {
		return err
	}
	_, err = v.adminClient.UpdateKind(ctx,
		&vstorepb.UpdateKindRequest{
			Namespace:         schema.Namespace,
			Kind:              schema.Kind,
			Properties:        properties,
			SecondaryIndexes:  SecondaryIndexesToPb(schema.SecondaryIndexes...),
			ExtendedKeyLength: schema.ExtendedKeyLength,
		}, grpc.FailFast(false))
	return err
}

// DeleteKind removes a kind from vStore and deletes all of its associated data and secondary indexes. This is a
// permanent process and can not be reversed.
func (v *vStore) DeleteKind(ctx context.Context, namespace, kind string) error {
	if v.adminClient == nil {
		return errors.New("Admin client must be initialized.")
	}
	_, err := v.adminClient.DeleteKind(ctx,
		&vstorepb.DeleteKindRequest{
			Namespace: namespace,
			Kind:      kind,
		}, grpc.FailFast(false))
	if err != nil {
		s, ok := status.FromError(err)
		if !ok {
			return err
		}
		if s.Code() == codes.ResourceExhausted {
			logging.Debugf(ctx, "Received Rate Limit error from VStore on deleting kind %s-%s, continuing", namespace, kind)
			return nil
		}
		return err
	}
	return nil
}

//LookupResult holds all of the information returned by the Lookup API that is relevant for a client
type LookupResult struct {
	Results    []Model
	NextCursor string
	HasMore    bool
}

// Register handles the registration of a certain kind with VStore, returning the kind as it exists in VStore
func (v *vStore) RegisterKind(ctx context.Context, namespace, kind string, serviceAccounts []string, model Model) (*vstorepb.GetKindResponse, error) {
	RegisterModel(namespace, kind, model)
	err := v.UpdateKind(ctx, model.Schema())
	if err == nil {
		return v.GetKind(ctx, namespace, kind)
	}
	if grpc.Code(err) != codes.NotFound {
		return nil, err
	}

	// Kind doesn't exist yet; create it
	err = v.CreateNamespace(ctx, namespace, serviceAccounts)
	if err != nil && grpc.Code(err) != codes.AlreadyExists {
		return nil, err
	}

	err = v.CreateKind(ctx, model.Schema())
	if err != nil {
		return nil, err
	}

	sch, err := v.GetKind(ctx, namespace, kind)
	if err != nil {
		return nil, err
	}

	return sch, nil
}

// GetSecondaryIndexName will tell you the name of table on the secondary index that VStore has created for the
// secondary index specified by indexID. The possible valid values for indexID are the same as the identifiers for
// the secondary indexes in your model's Schema.
func (v *vStore) GetSecondaryIndexName(ctx context.Context, namespace, kind string, indexID string) (string, error) {
	sch, err := v.GetKind(ctx, namespace, kind)
	if err != nil {
		return "", err
	}

	si := sch.GetSecondaryIndexes()
	if si == nil {
		return "", errors.New("Could not find any secondary indexes on the schema.")
	}
	if len(si) < 1 {
		return "", errors.New("Could not find any secondary indexes on the schema.")
	}

	var r *vstorepb.SecondaryIndex
	for _, i := range si {
		if i.Name == indexID {
			r = i
			break
		}
	}
	if r == nil {
		return "", errors.New("Could not find the specified secondary index on the schema.")
	}

	sql := r.GetCloudSqlConfig()
	if sql != nil {
		return sql.GetIndexName(), nil
	}
	es := r.GetEsConfig()
	if es != nil {
		return es.GetIndexName(), nil
	}
	esRaw := r.GetEsRawConfig()
	if esRaw != nil {
		return esRaw.GetIndexName(), nil
	}
	ps := r.GetPubsubConfig()
	if ps != nil {
		return ps.GetIndexName(), nil
	}
	return "", errors.New("Could not determine the type of secondary index")
}

//RegisterSubscriptionCallback will associate a handler with a pubsub subscription so that each time that subscription receives a message, it will be
//processed by the provided MessageHandler. You may also provide a cancelFunc that executes when the provided context is cancelled to allow for
//graceful shutdown logic. Lastly, you may provide zero or more worker options to parameterize the behaviour of the pubsub message processing mechanism,
//such as adjusting how many workers are started and how many messages they attempt to fetch at once. SubscriptionName can be an identifier of your own
//choosing, but it should not change, as this function creates new subscriptions if they do not yet exist.
func (v *vStore) RegisterSubscriptionCallback(ctx context.Context, namespace, kind, indexID, subscriptionName string, handler MessageHandler, cancelFunc context.CancelFunc, opts ...pubsub.WorkerOption) error {
	topicName, err := v.GetSecondaryIndexName(ctx, namespace, kind, indexID)
	if err != nil {
		return err
	}

	logging.Debugf(ctx, "Topic name: %s", topicName)

	s := fmt.Sprintf("%s-%s", topicName, subscriptionName)
	err = pubsub.GetOrCreateSubscription(ctx, v.pubsubClient, topicName, s)
	if err != nil {
		logging.Errorf(ctx, "Failed to get or create subscription to topic %s: %s", topicName, err.Error())
		return err
	}
	pubsubHandler := newPubsubMessageHandler(handler)
	pubsubHandler = pubsub.AddPubsubHandlerMonitoring(topicName, subscriptionName, pubsubHandler)
	pubsub.DoWork(ctx, s, v.pubsubClient, pubsubHandler, cancelFunc, opts...)

	return nil
}

func newPubsubMessageHandler(handler MessageHandler) pubsub.MessageHandler {
	return func(ctx context.Context, msg *gcloud_pubsub.Message) error {
		m, err := FromPubsubMessage(msg)
		if err != nil {
			logging.Errorf(ctx, "Failed to deserialize pubsub message: %s", err.Error())
			return err
		}
		err = handler(ctx, m)
		return err
	}
}
