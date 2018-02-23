// Code generated by protoc-gen-go. DO NOT EDIT.
// source: marketplace_analytics/v1/api.proto

/*
Package marketplaceanalytics_v1 is a generated protocol buffer package.

It is generated from these files:
	marketplace_analytics/v1/api.proto
	marketplace_analytics/v1/engagement.proto

It has these top-level messages:
	SortOptions
	GetPartnerEngagementRequest
	GetPartnerEngagementResponse
	ListPartnerEngagementRequest
	ListPartnerEngagementResponse
	PartnerEngagement
*/
package marketplaceanalytics_v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

// The directions that a result set can be sorted in.
type SortDirection int32

const (
	SortDirection_DESCENDING SortDirection = 0
	SortDirection_ASCENDING  SortDirection = 1
)

var SortDirection_name = map[int32]string{
	0: "DESCENDING",
	1: "ASCENDING",
}
var SortDirection_value = map[string]int32{
	"DESCENDING": 0,
	"ASCENDING":  1,
}

func (x SortDirection) String() string {
	return proto.EnumName(SortDirection_name, int32(x))
}
func (SortDirection) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// The field to sort a result set on.
type SortField int32

const (
	SortField_TOTAL_ACTIVATIONS                 SortField = 0
	SortField_PARTNER_NAME                      SortField = 1
	SortField_ENABLED_ON                        SortField = 2
	SortField_SELLING                           SortField = 3
	SortField_TOTAL_STORE_VIEWS                 SortField = 4
	SortField_TOTAL_STORE_CALL_TO_ACTION_EVENTS SortField = 5
	SortField_TOTAL_DEACTIVATIONS               SortField = 6
	SortField_TOTAL_PARTNER_VIEWS               SortField = 7
)

var SortField_name = map[int32]string{
	0: "TOTAL_ACTIVATIONS",
	1: "PARTNER_NAME",
	2: "ENABLED_ON",
	3: "SELLING",
	4: "TOTAL_STORE_VIEWS",
	5: "TOTAL_STORE_CALL_TO_ACTION_EVENTS",
	6: "TOTAL_DEACTIVATIONS",
	7: "TOTAL_PARTNER_VIEWS",
}
var SortField_value = map[string]int32{
	"TOTAL_ACTIVATIONS":                 0,
	"PARTNER_NAME":                      1,
	"ENABLED_ON":                        2,
	"SELLING":                           3,
	"TOTAL_STORE_VIEWS":                 4,
	"TOTAL_STORE_CALL_TO_ACTION_EVENTS": 5,
	"TOTAL_DEACTIVATIONS":               6,
	"TOTAL_PARTNER_VIEWS":               7,
}

func (x SortField) String() string {
	return proto.EnumName(SortField_name, int32(x))
}
func (SortField) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// Options for controlling the order of query results
type SortOptions struct {
	// A direction to sort results in
	Direction SortDirection `protobuf:"varint,1,opt,name=direction,enum=marketplaceanalytics.v1.SortDirection" json:"direction,omitempty"`
	// Field to sort on
	Field SortField `protobuf:"varint,2,opt,name=field,enum=marketplaceanalytics.v1.SortField" json:"field,omitempty"`
}

func (m *SortOptions) Reset()                    { *m = SortOptions{} }
func (m *SortOptions) String() string            { return proto.CompactTextString(m) }
func (*SortOptions) ProtoMessage()               {}
func (*SortOptions) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SortOptions) GetDirection() SortDirection {
	if m != nil {
		return m.Direction
	}
	return SortDirection_DESCENDING
}

func (m *SortOptions) GetField() SortField {
	if m != nil {
		return m.Field
	}
	return SortField_TOTAL_ACTIVATIONS
}

type GetPartnerEngagementRequest struct {
	// the app id to get engagement for
	AppId string `protobuf:"bytes,1,opt,name=app_id,json=appId" json:"app_id,omitempty"`
	// the partner id to get engagement for
	PartnerId string `protobuf:"bytes,2,opt,name=partner_id,json=partnerId" json:"partner_id,omitempty"`
}

func (m *GetPartnerEngagementRequest) Reset()                    { *m = GetPartnerEngagementRequest{} }
func (m *GetPartnerEngagementRequest) String() string            { return proto.CompactTextString(m) }
func (*GetPartnerEngagementRequest) ProtoMessage()               {}
func (*GetPartnerEngagementRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GetPartnerEngagementRequest) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *GetPartnerEngagementRequest) GetPartnerId() string {
	if m != nil {
		return m.PartnerId
	}
	return ""
}

type GetPartnerEngagementResponse struct {
	PartnerEngagement *PartnerEngagement `protobuf:"bytes,1,opt,name=partner_engagement,json=partnerEngagement" json:"partner_engagement,omitempty"`
}

func (m *GetPartnerEngagementResponse) Reset()                    { *m = GetPartnerEngagementResponse{} }
func (m *GetPartnerEngagementResponse) String() string            { return proto.CompactTextString(m) }
func (*GetPartnerEngagementResponse) ProtoMessage()               {}
func (*GetPartnerEngagementResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GetPartnerEngagementResponse) GetPartnerEngagement() *PartnerEngagement {
	if m != nil {
		return m.PartnerEngagement
	}
	return nil
}

type ListPartnerEngagementRequest struct {
	// the app id to get engagement for
	AppId string `protobuf:"bytes,1,opt,name=app_id,json=appId" json:"app_id,omitempty"`
	// A cursor to be provided when fetching additional pages of results beyond the first
	Cursor string `protobuf:"bytes,2,opt,name=cursor" json:"cursor,omitempty"`
	// Size of the page to return
	PageSize int64 `protobuf:"varint,3,opt,name=page_size,json=pageSize" json:"page_size,omitempty"`
	// The options you can sort a table on
	SortOptions *SortOptions `protobuf:"bytes,4,opt,name=sort_options,json=sortOptions" json:"sort_options,omitempty"`
}

func (m *ListPartnerEngagementRequest) Reset()                    { *m = ListPartnerEngagementRequest{} }
func (m *ListPartnerEngagementRequest) String() string            { return proto.CompactTextString(m) }
func (*ListPartnerEngagementRequest) ProtoMessage()               {}
func (*ListPartnerEngagementRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ListPartnerEngagementRequest) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *ListPartnerEngagementRequest) GetCursor() string {
	if m != nil {
		return m.Cursor
	}
	return ""
}

func (m *ListPartnerEngagementRequest) GetPageSize() int64 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListPartnerEngagementRequest) GetSortOptions() *SortOptions {
	if m != nil {
		return m.SortOptions
	}
	return nil
}

type ListPartnerEngagementResponse struct {
	PartnerEngagement []*PartnerEngagement `protobuf:"bytes,1,rep,name=partner_engagement,json=partnerEngagement" json:"partner_engagement,omitempty"`
}

func (m *ListPartnerEngagementResponse) Reset()                    { *m = ListPartnerEngagementResponse{} }
func (m *ListPartnerEngagementResponse) String() string            { return proto.CompactTextString(m) }
func (*ListPartnerEngagementResponse) ProtoMessage()               {}
func (*ListPartnerEngagementResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ListPartnerEngagementResponse) GetPartnerEngagement() []*PartnerEngagement {
	if m != nil {
		return m.PartnerEngagement
	}
	return nil
}

func init() {
	proto.RegisterType((*SortOptions)(nil), "marketplaceanalytics.v1.SortOptions")
	proto.RegisterType((*GetPartnerEngagementRequest)(nil), "marketplaceanalytics.v1.GetPartnerEngagementRequest")
	proto.RegisterType((*GetPartnerEngagementResponse)(nil), "marketplaceanalytics.v1.GetPartnerEngagementResponse")
	proto.RegisterType((*ListPartnerEngagementRequest)(nil), "marketplaceanalytics.v1.ListPartnerEngagementRequest")
	proto.RegisterType((*ListPartnerEngagementResponse)(nil), "marketplaceanalytics.v1.ListPartnerEngagementResponse")
	proto.RegisterEnum("marketplaceanalytics.v1.SortDirection", SortDirection_name, SortDirection_value)
	proto.RegisterEnum("marketplaceanalytics.v1.SortField", SortField_name, SortField_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for PartnerEngagementService service

type PartnerEngagementServiceClient interface {
	// Get all of the partner engagement for a partner and app
	GetPartnerEngagement(ctx context.Context, in *GetPartnerEngagementRequest, opts ...grpc.CallOption) (*GetPartnerEngagementResponse, error)
	// List all of the partner engagements for an app
	ListPartnerEngagement(ctx context.Context, in *ListPartnerEngagementRequest, opts ...grpc.CallOption) (*ListPartnerEngagementResponse, error)
}

type partnerEngagementServiceClient struct {
	cc *grpc.ClientConn
}

func NewPartnerEngagementServiceClient(cc *grpc.ClientConn) PartnerEngagementServiceClient {
	return &partnerEngagementServiceClient{cc}
}

func (c *partnerEngagementServiceClient) GetPartnerEngagement(ctx context.Context, in *GetPartnerEngagementRequest, opts ...grpc.CallOption) (*GetPartnerEngagementResponse, error) {
	out := new(GetPartnerEngagementResponse)
	err := grpc.Invoke(ctx, "/marketplaceanalytics.v1.PartnerEngagementService/GetPartnerEngagement", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *partnerEngagementServiceClient) ListPartnerEngagement(ctx context.Context, in *ListPartnerEngagementRequest, opts ...grpc.CallOption) (*ListPartnerEngagementResponse, error) {
	out := new(ListPartnerEngagementResponse)
	err := grpc.Invoke(ctx, "/marketplaceanalytics.v1.PartnerEngagementService/ListPartnerEngagement", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PartnerEngagementService service

type PartnerEngagementServiceServer interface {
	// Get all of the partner engagement for a partner and app
	GetPartnerEngagement(context.Context, *GetPartnerEngagementRequest) (*GetPartnerEngagementResponse, error)
	// List all of the partner engagements for an app
	ListPartnerEngagement(context.Context, *ListPartnerEngagementRequest) (*ListPartnerEngagementResponse, error)
}

func RegisterPartnerEngagementServiceServer(s *grpc.Server, srv PartnerEngagementServiceServer) {
	s.RegisterService(&_PartnerEngagementService_serviceDesc, srv)
}

func _PartnerEngagementService_GetPartnerEngagement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPartnerEngagementRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartnerEngagementServiceServer).GetPartnerEngagement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/marketplaceanalytics.v1.PartnerEngagementService/GetPartnerEngagement",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartnerEngagementServiceServer).GetPartnerEngagement(ctx, req.(*GetPartnerEngagementRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PartnerEngagementService_ListPartnerEngagement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPartnerEngagementRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartnerEngagementServiceServer).ListPartnerEngagement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/marketplaceanalytics.v1.PartnerEngagementService/ListPartnerEngagement",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartnerEngagementServiceServer).ListPartnerEngagement(ctx, req.(*ListPartnerEngagementRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PartnerEngagementService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "marketplaceanalytics.v1.PartnerEngagementService",
	HandlerType: (*PartnerEngagementServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPartnerEngagement",
			Handler:    _PartnerEngagementService_GetPartnerEngagement_Handler,
		},
		{
			MethodName: "ListPartnerEngagement",
			Handler:    _PartnerEngagementService_ListPartnerEngagement_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "marketplace_analytics/v1/api.proto",
}

func init() { proto.RegisterFile("marketplace_analytics/v1/api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 539 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x8d, 0x93, 0x26, 0xc5, 0x93, 0xb6, 0x72, 0x07, 0x42, 0xa3, 0xb4, 0x95, 0x8a, 0x05, 0xa8,
	0xe4, 0x90, 0xaa, 0xe1, 0x43, 0x5c, 0x4d, 0xbc, 0x44, 0x96, 0x8c, 0x5d, 0xd9, 0x56, 0x10, 0x27,
	0xcb, 0x24, 0x4b, 0x64, 0x91, 0xda, 0x8b, 0xbd, 0x8d, 0xd4, 0xdc, 0x91, 0xb8, 0xf0, 0x53, 0x38,
	0xf1, 0x0b, 0xf8, 0x67, 0xc8, 0x76, 0x12, 0x07, 0x1a, 0x07, 0x45, 0x1c, 0x67, 0x76, 0xde, 0x9b,
	0x37, 0x6f, 0xc6, 0x06, 0xf9, 0xda, 0x8b, 0x3e, 0x53, 0xce, 0x26, 0xde, 0x90, 0xba, 0x5e, 0xe0,
	0x4d, 0x6e, 0xb9, 0x3f, 0x8c, 0x2f, 0xa6, 0x97, 0x17, 0x1e, 0xf3, 0x3b, 0x2c, 0x0a, 0x79, 0x88,
	0x47, 0x2b, 0x35, 0xcb, 0x92, 0xce, 0xf4, 0xb2, 0xf5, 0xac, 0x10, 0x4c, 0x83, 0xb1, 0x37, 0xa6,
	0xd7, 0x34, 0xe0, 0x19, 0x87, 0xfc, 0x5d, 0x80, 0xba, 0x1d, 0x46, 0xdc, 0x64, 0xdc, 0x0f, 0x83,
	0x18, 0x55, 0x10, 0x47, 0x7e, 0x44, 0x87, 0x49, 0xd4, 0x14, 0xce, 0x84, 0xf3, 0x83, 0xee, 0xd3,
	0x4e, 0x41, 0x9f, 0x4e, 0x02, 0x54, 0x17, 0xd5, 0x56, 0x0e, 0xc4, 0xd7, 0x50, 0xfd, 0xe4, 0xd3,
	0xc9, 0xa8, 0x59, 0x4e, 0x19, 0xe4, 0x8d, 0x0c, 0x6f, 0x93, 0x4a, 0x2b, 0x03, 0xc8, 0x36, 0x1c,
	0xf7, 0x29, 0xbf, 0xf2, 0x22, 0x1e, 0xd0, 0x88, 0x2c, 0xd5, 0x5a, 0xf4, 0xcb, 0x0d, 0x8d, 0x39,
	0x36, 0xa0, 0xe6, 0x31, 0xe6, 0xfa, 0xa3, 0x54, 0x9b, 0x68, 0x55, 0x3d, 0xc6, 0xb4, 0x11, 0x9e,
	0x02, 0xb0, 0x0c, 0x92, 0x3c, 0x95, 0xd3, 0x27, 0x71, 0x9e, 0xd1, 0x46, 0xf2, 0x2d, 0x9c, 0xac,
	0x27, 0x8d, 0x59, 0x18, 0xc4, 0x14, 0x3f, 0x00, 0x2e, 0xe0, 0xb9, 0x41, 0x69, 0x87, 0x7a, 0xb7,
	0x5d, 0xa8, 0xfd, 0x2e, 0xdf, 0x21, 0xfb, 0x3b, 0x25, 0xff, 0x14, 0xe0, 0x44, 0xf7, 0xe3, 0xad,
	0x27, 0x7a, 0x08, 0xb5, 0xe1, 0x4d, 0x14, 0x87, 0xd1, 0x7c, 0x9a, 0x79, 0x84, 0xc7, 0x20, 0x32,
	0x6f, 0x4c, 0xdd, 0xd8, 0x9f, 0xd1, 0x66, 0xe5, 0x4c, 0x38, 0xaf, 0x58, 0xf7, 0x92, 0x84, 0xed,
	0xcf, 0x28, 0xf6, 0x61, 0x2f, 0x0e, 0x23, 0xee, 0x86, 0xd9, 0x32, 0x9b, 0x3b, 0xe9, 0x04, 0x8f,
	0x37, 0xba, 0x3f, 0x5f, 0xbc, 0x55, 0x8f, 0xf3, 0x40, 0x9e, 0xc1, 0x69, 0x81, 0xe8, 0x7f, 0x38,
	0x56, 0xf9, 0x6f, 0xc7, 0xda, 0x1d, 0xd8, 0xff, 0xe3, 0xae, 0xf0, 0x00, 0x40, 0x25, 0x76, 0x8f,
	0x18, 0xaa, 0x66, 0xf4, 0xa5, 0x12, 0xee, 0x83, 0xa8, 0x2c, 0x43, 0xa1, 0xfd, 0x4b, 0x00, 0x71,
	0x79, 0x46, 0xd8, 0x80, 0x43, 0xc7, 0x74, 0x14, 0xdd, 0x55, 0x7a, 0x8e, 0x36, 0x50, 0x1c, 0xcd,
	0x34, 0x6c, 0xa9, 0x84, 0x12, 0xec, 0x5d, 0x29, 0x96, 0x63, 0x10, 0xcb, 0x35, 0x94, 0x77, 0x44,
	0x12, 0x12, 0x56, 0x62, 0x28, 0x6f, 0x74, 0xa2, 0xba, 0xa6, 0x21, 0x95, 0xb1, 0x0e, 0xbb, 0x36,
	0xd1, 0xf5, 0x84, 0xb3, 0x92, 0xb3, 0xd8, 0x8e, 0x69, 0x11, 0x77, 0xa0, 0x91, 0xf7, 0xb6, 0xb4,
	0x83, 0x4f, 0xe0, 0xd1, 0x6a, 0xba, 0xa7, 0xe8, 0xba, 0xeb, 0x98, 0x69, 0x2b, 0xd3, 0x70, 0xc9,
	0x80, 0x18, 0x8e, 0x2d, 0x55, 0xf1, 0x08, 0xee, 0x67, 0x65, 0x2a, 0x59, 0x55, 0x51, 0xcb, 0x1f,
	0x16, 0x5a, 0x32, 0xe2, 0xdd, 0xee, 0x8f, 0x32, 0x34, 0xef, 0x98, 0x63, 0xd3, 0x68, 0xea, 0x0f,
	0x29, 0x7e, 0x15, 0xe0, 0xc1, 0xba, 0xf3, 0xc5, 0x17, 0x85, 0x46, 0x6f, 0xf8, 0x84, 0x5a, 0x2f,
	0xb7, 0x44, 0x65, 0x1b, 0x97, 0x4b, 0xf8, 0x4d, 0x80, 0xc6, 0xda, 0xab, 0xc0, 0x62, 0xca, 0x4d,
	0xa7, 0xdf, 0x7a, 0xb5, 0x2d, 0x6c, 0x21, 0xe5, 0x63, 0x2d, 0xfd, 0x79, 0x3d, 0xff, 0x1d, 0x00,
	0x00, 0xff, 0xff, 0xa4, 0x4f, 0xba, 0x19, 0x26, 0x05, 0x00, 0x00,
}
