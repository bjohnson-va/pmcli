package vstore

import (
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	"github.com/stretchr/testify/mock"
	"github.com/vendasta/gosdks/pb/vstorepb"
	"github.com/vendasta/gosdks/pubsub"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// TransactionMock implements vstore.Transaction
type TransactionMock struct {
	mock.Mock
}

func (tm *TransactionMock) Save(m Model) error {
	args := tm.Called(m)
	return args.Error(0)
}

// VStoreMock implements vstore.Interface
type VStoreMock struct {
	mock.Mock
}

func (v *VStoreMock) GetMulti(ctx context.Context, ks []*KeySet) ([]Model, error) {
	args := v.Called(ctx, ks)
	res := args.Get(0)
	if res == nil {
		return nil, args.Error(1)
	}
	return res.([]Model), args.Error(1)
}

func (v *VStoreMock) Get(ctx context.Context, ks *KeySet) (Model, error) {
	args := v.Called(ctx, ks)
	res := args.Get(0)
	if res == nil {
		return nil, args.Error(1)
	}
	return res.(Model), args.Error(1)
}

func (v *VStoreMock) Lookup(ctx context.Context, namespace, kind string, opts ...LookupOption) (*LookupResult, error) {
	args := v.Called(ctx, namespace, kind, opts)
	res := args.Get(0)
	if res == nil {
		return nil, args.Error(1)
	}
	return res.(*LookupResult), args.Error(1)
}

func (v *VStoreMock) Scan(ctx context.Context, namespace, kind string, cb func(m Model) bool, opts ...ScanOption) error {
	args := v.Called(ctx, namespace, kind, cb, opts)
	return args.Error(0)
}

func (v *VStoreMock) Transaction(ctx context.Context, ks *KeySet, f func(Transaction, Model) error, opts ...TransactionOption) error {
	args := v.Called(ctx, ks, f, opts)
	return args.Error(0)
}

func (v *VStoreMock) CreateNamespace(ctx context.Context, namespace string, authorizedServiceAccounts []string) error {
	args := v.Called(ctx, namespace, authorizedServiceAccounts)
	return args.Error(0)
}

func (v *VStoreMock) UpdateNamespace(ctx context.Context, namespace string, authorizedServiceAccounts []string) error {
	args := v.Called(ctx, namespace, authorizedServiceAccounts)
	return args.Error(0)
}

func (v *VStoreMock) DeleteNamespace(ctx context.Context, namespace string) error {
	args := v.Called(ctx, namespace)
	return args.Error(0)
}

func (v *VStoreMock) CreateKind(ctx context.Context, schema *Schema) error {
	args := v.Called(ctx, schema)
	return args.Error(0)
}

func (v *VStoreMock) UpdateKind(ctx context.Context, schema *Schema) error {
	args := v.Called(ctx, schema)
	return args.Error(0)
}

func (v *VStoreMock) GetKind(ctx context.Context, namespace string, kind string) (*vstorepb.GetKindResponse, error) {
	args := v.Called(ctx, namespace, kind)
	res := args.Get(0)
	if res == nil {
		return nil, args.Error(1)
	}
	return res.(*vstorepb.GetKindResponse), args.Error(1)
}

func (v *VStoreMock) RegisterKind(ctx context.Context, namespace, kind string, serviceAccounts []string, model Model) (*vstorepb.GetKindResponse, error) {
	args := v.Called(ctx, namespace, kind, serviceAccounts, model)
	res := args.Get(0)
	if res == nil {
		return nil, args.Error(1)
	}
	return res.(*vstorepb.GetKindResponse), args.Error(1)
}

func (v *VStoreMock) DeleteKind(ctx context.Context, namespace string, kind string) error {
	args := v.Called(ctx, namespace, kind)
	return args.Error(0)
}

func (v *VStoreMock) GetSecondaryIndexName(ctx context.Context, namespace, kind string, indexID string) (string, error) {
	args := v.Called(ctx, namespace, kind, indexID)
	return args.String(0), args.Error(0)
}

func (v *VStoreMock) RegisterSubscriptionCallback(ctx context.Context, namespace, kind, indexID, subscriptionName string, handler MessageHandler, cancelFunc context.CancelFunc, opts ...pubsub.WorkerOption) error {
	args := v.Called(ctx, namespace, kind, indexID, subscriptionName, handler, cancelFunc, opts)
	return args.Error(0)
}

// ClientMock implements vstorepb.VStoreClient
type ClientMock struct {
	mock.Mock
}

// Create mocks VStoreClient.Create
func (v *ClientMock) Create(ctx context.Context, in *vstorepb.CreateRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	args := v.Called(ctx, in, opts)
	g := args.Get(0)
	if g == nil {
		return nil, args.Error(1)
	}
	return g.(*google_protobuf.Empty), args.Error(1)
}

// Get mocks VStoreClient.Get
func (v *ClientMock) Get(ctx context.Context, in *vstorepb.GetRequest, opts ...grpc.CallOption) (*vstorepb.GetResponse, error) {
	args := v.Called(ctx, in, opts)
	g := args.Get(0)
	if g == nil {
		return nil, args.Error(1)
	}
	return g.(*vstorepb.GetResponse), args.Error(1)
}

// Lookup mocks VStoreClient.Lookup
func (v *ClientMock) Lookup(ctx context.Context, in *vstorepb.LookupRequest, opts ...grpc.CallOption) (*vstorepb.LookupResponse, error) {
	args := v.Called(ctx, in, opts)
	g := args.Get(0)
	if g == nil {
		return nil, args.Error(1)
	}
	return g.(*vstorepb.LookupResponse), args.Error(1)
}

// Update mocks VStoreClient.Update
func (v *ClientMock) Update(ctx context.Context, in *vstorepb.UpdateRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	args := v.Called(ctx, in, opts)
	g := args.Get(0)
	if g == nil {
		return nil, args.Error(1)
	}
	return g.(*google_protobuf.Empty), args.Error(1)
}

// Scan mocks VStoreClient.Scan
func (v *ClientMock) Scan(ctx context.Context, in *vstorepb.ScanRequest, opts ...grpc.CallOption) (vstorepb.VStore_ScanClient, error) {
	args := v.Called(ctx, in, opts)
	g := args.Get(0)
	if g == nil {
		return nil, args.Error(1)
	}
	return g.(vstorepb.VStore_ScanClient), args.Error(1)
}

// AdminClientMock implements vstorepb.VStoreAdminClient
type AdminClientMock struct {
	mock.Mock
}

// CreateNamespace mocks VStoreAdminClient.CreateNamespace
func (v *AdminClientMock) CreateNamespace(ctx context.Context, in *vstorepb.CreateNamespaceRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	args := v.Called(ctx, in, opts)
	g := args.Get(0)
	if g == nil {
		return nil, args.Error(1)
	}
	return g.(*google_protobuf.Empty), args.Error(1)
}

// UpdateNamespace mocks VStoreAdminClient.UpdateNamespace
func (v *AdminClientMock) UpdateNamespace(ctx context.Context, in *vstorepb.UpdateNamespaceRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	args := v.Called(ctx, in, opts)
	g := args.Get(0)
	if g == nil {
		return nil, args.Error(1)
	}
	return g.(*google_protobuf.Empty), args.Error(1)
}

// DeleteNamespace mocks VStoreAdminClient.DeleteNamespace
func (v *AdminClientMock) DeleteNamespace(ctx context.Context, in *vstorepb.DeleteNamespaceRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	args := v.Called(ctx, in, opts)
	g := args.Get(0)
	if g == nil {
		return nil, args.Error(1)
	}
	return g.(*google_protobuf.Empty), args.Error(1)
}

// CreateKind mocks VStoreAdminClient.CreateKind
func (v *AdminClientMock) CreateKind(ctx context.Context, in *vstorepb.CreateKindRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	args := v.Called(ctx, in, opts)
	g := args.Get(0)
	if g == nil {
		return nil, args.Error(1)
	}
	return g.(*google_protobuf.Empty), args.Error(1)
}

// UpdateKind mocks VStoreAdminClient.UpdateKind
func (v *AdminClientMock) UpdateKind(ctx context.Context, in *vstorepb.UpdateKindRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	args := v.Called(ctx, in, opts)
	g := args.Get(0)
	if g == nil {
		return nil, args.Error(1)
	}
	return g.(*google_protobuf.Empty), args.Error(1)
}

// GetKind mocks VStoreAdminClient.GetKind
func (v *AdminClientMock) GetKind(ctx context.Context, in *vstorepb.GetKindRequest, opts ...grpc.CallOption) (*vstorepb.GetKindResponse, error) {
	args := v.Called(ctx, in, opts)
	g := args.Get(0)
	if g == nil {
		return nil, args.Error(1)
	}
	return g.(*vstorepb.GetKindResponse), args.Error(1)
}

// DeleteKind mocks VStoreAdminClient.DeleteKind
func (v *AdminClientMock) DeleteKind(ctx context.Context, in *vstorepb.DeleteKindRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	args := v.Called(ctx, in, opts)
	g := args.Get(0)
	if g == nil {
		return nil, args.Error(1)
	}
	return g.(*google_protobuf.Empty), args.Error(1)
}
