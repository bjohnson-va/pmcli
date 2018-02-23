// Code generated by protoc-gen-go. DO NOT EDIT.
// source: snapshot/v1/api.proto

/*
Package snapshot_v1 is a generated protocol buffer package.

It is generated from these files:
	snapshot/v1/api.proto
	snapshot/v1/field_mask.proto
	snapshot/v1/reviews.proto
	snapshot/v1/snapshot.proto

It has these top-level messages:
	GetSectionRequest
	FieldMask
	ReviewSection
	ReviewConfig
	ReviewData
	SampleSourceCount
	ReviewDataItem
	UpdateReviewConfigRequest
	GetReviewSectionResponse
	GlobalConfig
	SectionContent
	Snapshot
	CreateSnapshotRequest
	CreateSnapshotResponse
	GetSnapshotRequest
	GetSnapshotResponse
	UpdateSnapshotConfigRequest
	GetSnapshotIDRequest
	GetSnapshotIDResponse
*/
package snapshot_v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/empty"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetSectionRequest struct {
	// The id of the snapshot to get a section for
	SnapshotId string `protobuf:"bytes,1,opt,name=snapshot_id,json=snapshotId" json:"snapshot_id,omitempty"`
}

func (m *GetSectionRequest) Reset()                    { *m = GetSectionRequest{} }
func (m *GetSectionRequest) String() string            { return proto.CompactTextString(m) }
func (*GetSectionRequest) ProtoMessage()               {}
func (*GetSectionRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetSectionRequest) GetSnapshotId() string {
	if m != nil {
		return m.SnapshotId
	}
	return ""
}

func init() {
	proto.RegisterType((*GetSectionRequest)(nil), "snapshot.v1.GetSectionRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for SnapshotService service

type SnapshotServiceClient interface {
	Create(ctx context.Context, in *CreateSnapshotRequest, opts ...grpc.CallOption) (*CreateSnapshotResponse, error)
	Get(ctx context.Context, in *GetSnapshotRequest, opts ...grpc.CallOption) (*GetSnapshotResponse, error)
	UpdateConfig(ctx context.Context, in *UpdateSnapshotConfigRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	// GetSnapshotID converts a partner/market combo to the snapshotID to be used to edit the default config
	GetSnapshotID(ctx context.Context, in *GetSnapshotIDRequest, opts ...grpc.CallOption) (*GetSnapshotIDResponse, error)
}

type snapshotServiceClient struct {
	cc *grpc.ClientConn
}

func NewSnapshotServiceClient(cc *grpc.ClientConn) SnapshotServiceClient {
	return &snapshotServiceClient{cc}
}

func (c *snapshotServiceClient) Create(ctx context.Context, in *CreateSnapshotRequest, opts ...grpc.CallOption) (*CreateSnapshotResponse, error) {
	out := new(CreateSnapshotResponse)
	err := grpc.Invoke(ctx, "/snapshot.v1.SnapshotService/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *snapshotServiceClient) Get(ctx context.Context, in *GetSnapshotRequest, opts ...grpc.CallOption) (*GetSnapshotResponse, error) {
	out := new(GetSnapshotResponse)
	err := grpc.Invoke(ctx, "/snapshot.v1.SnapshotService/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *snapshotServiceClient) UpdateConfig(ctx context.Context, in *UpdateSnapshotConfigRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/snapshot.v1.SnapshotService/UpdateConfig", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *snapshotServiceClient) GetSnapshotID(ctx context.Context, in *GetSnapshotIDRequest, opts ...grpc.CallOption) (*GetSnapshotIDResponse, error) {
	out := new(GetSnapshotIDResponse)
	err := grpc.Invoke(ctx, "/snapshot.v1.SnapshotService/GetSnapshotID", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SnapshotService service

type SnapshotServiceServer interface {
	Create(context.Context, *CreateSnapshotRequest) (*CreateSnapshotResponse, error)
	Get(context.Context, *GetSnapshotRequest) (*GetSnapshotResponse, error)
	UpdateConfig(context.Context, *UpdateSnapshotConfigRequest) (*google_protobuf.Empty, error)
	// GetSnapshotID converts a partner/market combo to the snapshotID to be used to edit the default config
	GetSnapshotID(context.Context, *GetSnapshotIDRequest) (*GetSnapshotIDResponse, error)
}

func RegisterSnapshotServiceServer(s *grpc.Server, srv SnapshotServiceServer) {
	s.RegisterService(&_SnapshotService_serviceDesc, srv)
}

func _SnapshotService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSnapshotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnapshotServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/snapshot.v1.SnapshotService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnapshotServiceServer).Create(ctx, req.(*CreateSnapshotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SnapshotService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSnapshotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnapshotServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/snapshot.v1.SnapshotService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnapshotServiceServer).Get(ctx, req.(*GetSnapshotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SnapshotService_UpdateConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSnapshotConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnapshotServiceServer).UpdateConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/snapshot.v1.SnapshotService/UpdateConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnapshotServiceServer).UpdateConfig(ctx, req.(*UpdateSnapshotConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SnapshotService_GetSnapshotID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSnapshotIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnapshotServiceServer).GetSnapshotID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/snapshot.v1.SnapshotService/GetSnapshotID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnapshotServiceServer).GetSnapshotID(ctx, req.(*GetSnapshotIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SnapshotService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "snapshot.v1.SnapshotService",
	HandlerType: (*SnapshotServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _SnapshotService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _SnapshotService_Get_Handler,
		},
		{
			MethodName: "UpdateConfig",
			Handler:    _SnapshotService_UpdateConfig_Handler,
		},
		{
			MethodName: "GetSnapshotID",
			Handler:    _SnapshotService_GetSnapshotID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "snapshot/v1/api.proto",
}

// Client API for ReviewSectionService service

type ReviewSectionServiceClient interface {
	Get(ctx context.Context, in *GetSectionRequest, opts ...grpc.CallOption) (*GetReviewSectionResponse, error)
	UpdateConfig(ctx context.Context, in *UpdateReviewConfigRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
}

type reviewSectionServiceClient struct {
	cc *grpc.ClientConn
}

func NewReviewSectionServiceClient(cc *grpc.ClientConn) ReviewSectionServiceClient {
	return &reviewSectionServiceClient{cc}
}

func (c *reviewSectionServiceClient) Get(ctx context.Context, in *GetSectionRequest, opts ...grpc.CallOption) (*GetReviewSectionResponse, error) {
	out := new(GetReviewSectionResponse)
	err := grpc.Invoke(ctx, "/snapshot.v1.ReviewSectionService/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewSectionServiceClient) UpdateConfig(ctx context.Context, in *UpdateReviewConfigRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/snapshot.v1.ReviewSectionService/UpdateConfig", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ReviewSectionService service

type ReviewSectionServiceServer interface {
	Get(context.Context, *GetSectionRequest) (*GetReviewSectionResponse, error)
	UpdateConfig(context.Context, *UpdateReviewConfigRequest) (*google_protobuf.Empty, error)
}

func RegisterReviewSectionServiceServer(s *grpc.Server, srv ReviewSectionServiceServer) {
	s.RegisterService(&_ReviewSectionService_serviceDesc, srv)
}

func _ReviewSectionService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewSectionServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/snapshot.v1.ReviewSectionService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewSectionServiceServer).Get(ctx, req.(*GetSectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReviewSectionService_UpdateConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateReviewConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewSectionServiceServer).UpdateConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/snapshot.v1.ReviewSectionService/UpdateConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewSectionServiceServer).UpdateConfig(ctx, req.(*UpdateReviewConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ReviewSectionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "snapshot.v1.ReviewSectionService",
	HandlerType: (*ReviewSectionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _ReviewSectionService_Get_Handler,
		},
		{
			MethodName: "UpdateConfig",
			Handler:    _ReviewSectionService_UpdateConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "snapshot/v1/api.proto",
}

func init() { proto.RegisterFile("snapshot/v1/api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 309 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x91, 0x51, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0xa9, 0x42, 0xc1, 0x55, 0x11, 0x17, 0x15, 0x5d, 0xc1, 0xd6, 0x88, 0xd2, 0xa7, 0x0d,
	0xad, 0xde, 0xa0, 0x95, 0x1a, 0x10, 0xd1, 0x04, 0x7d, 0x95, 0xb4, 0x99, 0xc6, 0x05, 0xcd, 0xae,
	0xd9, 0x6d, 0xc4, 0x7b, 0x79, 0x18, 0x8f, 0x23, 0xcd, 0x64, 0x43, 0x12, 0x6d, 0xf1, 0x71, 0xe6,
	0xff, 0xe7, 0x9f, 0xe1, 0x1b, 0xb2, 0xaf, 0x93, 0x50, 0xe9, 0x17, 0x69, 0xdc, 0xac, 0xef, 0x86,
	0x4a, 0x70, 0x95, 0x4a, 0x23, 0xe9, 0xa6, 0x6d, 0xf3, 0xac, 0xcf, 0x8e, 0x63, 0x29, 0xe3, 0x57,
	0x70, 0x73, 0x69, 0x32, 0x9f, 0xb9, 0xf0, 0xa6, 0xcc, 0x27, 0x3a, 0x19, 0xab, 0x06, 0x94, 0x53,
	0xa8, 0x1d, 0x55, 0xb5, 0x14, 0x32, 0x01, 0x1f, 0x1a, 0x25, 0xe7, 0x8a, 0xec, 0x8e, 0xc1, 0x04,
	0x30, 0x35, 0x42, 0x26, 0x3e, 0xbc, 0xcf, 0x41, 0x1b, 0xda, 0x21, 0xe5, 0xde, 0x67, 0x11, 0x1d,
	0xb6, 0xba, 0xad, 0xde, 0x86, 0x4f, 0x6c, 0xcb, 0x8b, 0x06, 0xdf, 0x6b, 0x64, 0x27, 0x28, 0xca,
	0x00, 0xd2, 0x4c, 0x4c, 0x81, 0x3e, 0x90, 0xf6, 0x30, 0x85, 0xd0, 0x00, 0x75, 0x78, 0xe5, 0x6a,
	0x8e, 0x4d, 0xeb, 0x2e, 0x56, 0xb0, 0xb3, 0x95, 0x1e, 0xad, 0x64, 0xa2, 0x81, 0xde, 0x90, 0xf5,
	0x31, 0x18, 0xda, 0xa9, 0x79, 0x17, 0xe7, 0x36, 0xc2, 0xba, 0xcb, 0x0d, 0x45, 0xd2, 0x3d, 0xd9,
	0x7a, 0x54, 0x51, 0x68, 0x60, 0x28, 0x93, 0x99, 0x88, 0x69, 0xaf, 0x36, 0x81, 0x92, 0x1d, 0x42,
	0x8b, 0xcd, 0x3e, 0xe0, 0x48, 0x9d, 0x5b, 0xea, 0xfc, 0x7a, 0x41, 0x9d, 0x3e, 0x91, 0xed, 0xca,
	0x22, 0x6f, 0x44, 0x4f, 0x97, 0x1d, 0xe1, 0x8d, 0x6c, 0x96, 0xb3, 0xca, 0x82, 0x97, 0x0e, 0xbe,
	0x5a, 0x64, 0xcf, 0xcf, 0x5f, 0x54, 0x3c, 0xc5, 0xf2, 0xbd, 0x45, 0x18, 0x27, 0xbf, 0x32, 0x6a,
	0xbf, 0x63, 0xe7, 0x4d, 0xbd, 0x16, 0x56, 0x02, 0xb9, 0x6b, 0x00, 0xb9, 0xf8, 0x03, 0x08, 0x4e,
	0xfe, 0x0b, 0xc7, 0xa4, 0x9d, 0xd7, 0x97, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xe3, 0xe4, 0x53,
	0x5f, 0xc8, 0x02, 0x00, 0x00,
}
