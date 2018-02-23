// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sales_orders/v1/api.proto

/*
Package salesorders_v1 is a generated protocol buffer package.

It is generated from these files:
	sales_orders/v1/api.proto
	sales_orders/v1/common.proto
	sales_orders/v1/sales_orders.proto

It has these top-level messages:
	RevenueComponent
	Revenue
	AddonKey
	Package
	Field
	CustomField
	CommonField
	ProductActivation
	AddonActivation
	Order
	CreateSalesOrderRequest
	CreateSalesOrderResponse
	GetSalesOrderRequest
	GetSalesOrderResponse
	ListSalesOrderRequest
	ListSalesOrderResponse
	ApproveSalesOrderRequest
	ApproveSalesOrderResponse
	DeclineSalesOrderRequest
	DeclineSalesOrderResponse
	ActivateProductsRequest
	ActivateProductsResponse
*/
package salesorders_v1

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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for SalesOrders service

type SalesOrdersClient interface {
	// Create data for a sales order
	CreateSalesOrder(ctx context.Context, in *CreateSalesOrderRequest, opts ...grpc.CallOption) (*CreateSalesOrderResponse, error)
	// Get data for a sales order
	GetSalesOrder(ctx context.Context, in *GetSalesOrderRequest, opts ...grpc.CallOption) (*GetSalesOrderResponse, error)
	// List sales orders
	ListSalesOrder(ctx context.Context, in *ListSalesOrderRequest, opts ...grpc.CallOption) (*ListSalesOrderResponse, error)
	// Approve a sales order
	ApproveSalesOrder(ctx context.Context, in *ApproveSalesOrderRequest, opts ...grpc.CallOption) (*ApproveSalesOrderResponse, error)
	// Decline a sales order
	DeclineSalesOrder(ctx context.Context, in *DeclineSalesOrderRequest, opts ...grpc.CallOption) (*DeclineSalesOrderResponse, error)
	// Activate products in a sales order
	ActivateProducts(ctx context.Context, in *ActivateProductsRequest, opts ...grpc.CallOption) (*ActivateProductsResponse, error)
}

type salesOrdersClient struct {
	cc *grpc.ClientConn
}

func NewSalesOrdersClient(cc *grpc.ClientConn) SalesOrdersClient {
	return &salesOrdersClient{cc}
}

func (c *salesOrdersClient) CreateSalesOrder(ctx context.Context, in *CreateSalesOrderRequest, opts ...grpc.CallOption) (*CreateSalesOrderResponse, error) {
	out := new(CreateSalesOrderResponse)
	err := grpc.Invoke(ctx, "/salesorders.v1.SalesOrders/CreateSalesOrder", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *salesOrdersClient) GetSalesOrder(ctx context.Context, in *GetSalesOrderRequest, opts ...grpc.CallOption) (*GetSalesOrderResponse, error) {
	out := new(GetSalesOrderResponse)
	err := grpc.Invoke(ctx, "/salesorders.v1.SalesOrders/GetSalesOrder", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *salesOrdersClient) ListSalesOrder(ctx context.Context, in *ListSalesOrderRequest, opts ...grpc.CallOption) (*ListSalesOrderResponse, error) {
	out := new(ListSalesOrderResponse)
	err := grpc.Invoke(ctx, "/salesorders.v1.SalesOrders/ListSalesOrder", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *salesOrdersClient) ApproveSalesOrder(ctx context.Context, in *ApproveSalesOrderRequest, opts ...grpc.CallOption) (*ApproveSalesOrderResponse, error) {
	out := new(ApproveSalesOrderResponse)
	err := grpc.Invoke(ctx, "/salesorders.v1.SalesOrders/ApproveSalesOrder", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *salesOrdersClient) DeclineSalesOrder(ctx context.Context, in *DeclineSalesOrderRequest, opts ...grpc.CallOption) (*DeclineSalesOrderResponse, error) {
	out := new(DeclineSalesOrderResponse)
	err := grpc.Invoke(ctx, "/salesorders.v1.SalesOrders/DeclineSalesOrder", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *salesOrdersClient) ActivateProducts(ctx context.Context, in *ActivateProductsRequest, opts ...grpc.CallOption) (*ActivateProductsResponse, error) {
	out := new(ActivateProductsResponse)
	err := grpc.Invoke(ctx, "/salesorders.v1.SalesOrders/ActivateProducts", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SalesOrders service

type SalesOrdersServer interface {
	// Create data for a sales order
	CreateSalesOrder(context.Context, *CreateSalesOrderRequest) (*CreateSalesOrderResponse, error)
	// Get data for a sales order
	GetSalesOrder(context.Context, *GetSalesOrderRequest) (*GetSalesOrderResponse, error)
	// List sales orders
	ListSalesOrder(context.Context, *ListSalesOrderRequest) (*ListSalesOrderResponse, error)
	// Approve a sales order
	ApproveSalesOrder(context.Context, *ApproveSalesOrderRequest) (*ApproveSalesOrderResponse, error)
	// Decline a sales order
	DeclineSalesOrder(context.Context, *DeclineSalesOrderRequest) (*DeclineSalesOrderResponse, error)
	// Activate products in a sales order
	ActivateProducts(context.Context, *ActivateProductsRequest) (*ActivateProductsResponse, error)
}

func RegisterSalesOrdersServer(s *grpc.Server, srv SalesOrdersServer) {
	s.RegisterService(&_SalesOrders_serviceDesc, srv)
}

func _SalesOrders_CreateSalesOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSalesOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SalesOrdersServer).CreateSalesOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/salesorders.v1.SalesOrders/CreateSalesOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SalesOrdersServer).CreateSalesOrder(ctx, req.(*CreateSalesOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SalesOrders_GetSalesOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSalesOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SalesOrdersServer).GetSalesOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/salesorders.v1.SalesOrders/GetSalesOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SalesOrdersServer).GetSalesOrder(ctx, req.(*GetSalesOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SalesOrders_ListSalesOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSalesOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SalesOrdersServer).ListSalesOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/salesorders.v1.SalesOrders/ListSalesOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SalesOrdersServer).ListSalesOrder(ctx, req.(*ListSalesOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SalesOrders_ApproveSalesOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApproveSalesOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SalesOrdersServer).ApproveSalesOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/salesorders.v1.SalesOrders/ApproveSalesOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SalesOrdersServer).ApproveSalesOrder(ctx, req.(*ApproveSalesOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SalesOrders_DeclineSalesOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeclineSalesOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SalesOrdersServer).DeclineSalesOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/salesorders.v1.SalesOrders/DeclineSalesOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SalesOrdersServer).DeclineSalesOrder(ctx, req.(*DeclineSalesOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SalesOrders_ActivateProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActivateProductsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SalesOrdersServer).ActivateProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/salesorders.v1.SalesOrders/ActivateProducts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SalesOrdersServer).ActivateProducts(ctx, req.(*ActivateProductsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SalesOrders_serviceDesc = grpc.ServiceDesc{
	ServiceName: "salesorders.v1.SalesOrders",
	HandlerType: (*SalesOrdersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSalesOrder",
			Handler:    _SalesOrders_CreateSalesOrder_Handler,
		},
		{
			MethodName: "GetSalesOrder",
			Handler:    _SalesOrders_GetSalesOrder_Handler,
		},
		{
			MethodName: "ListSalesOrder",
			Handler:    _SalesOrders_ListSalesOrder_Handler,
		},
		{
			MethodName: "ApproveSalesOrder",
			Handler:    _SalesOrders_ApproveSalesOrder_Handler,
		},
		{
			MethodName: "DeclineSalesOrder",
			Handler:    _SalesOrders_DeclineSalesOrder_Handler,
		},
		{
			MethodName: "ActivateProducts",
			Handler:    _SalesOrders_ActivateProducts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sales_orders/v1/api.proto",
}

func init() { proto.RegisterFile("sales_orders/v1/api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 236 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x92, 0x3f, 0x4b, 0xc7, 0x30,
	0x10, 0x86, 0x1d, 0xc4, 0x21, 0x62, 0xa9, 0xd9, 0xec, 0x58, 0xfc, 0xbb, 0xa4, 0x54, 0x3f, 0x41,
	0x51, 0x70, 0x11, 0x14, 0xdd, 0x95, 0xd8, 0x1e, 0x25, 0x52, 0x9a, 0x98, 0x4b, 0xf3, 0xe1, 0x9d,
	0xc4, 0x96, 0xd4, 0x34, 0xb1, 0xf6, 0x37, 0xe6, 0xde, 0xe7, 0xf2, 0xc0, 0xf1, 0x92, 0x13, 0xe4,
	0x1d, 0xe0, 0x9b, 0xd4, 0x0d, 0x68, 0x2c, 0x6c, 0x59, 0x70, 0x25, 0x98, 0xd2, 0xd2, 0x48, 0x9a,
	0x8c, 0xd1, 0x94, 0x30, 0x5b, 0x66, 0x79, 0x88, 0xfa, 0xef, 0x69, 0xe7, 0xfa, 0x6b, 0x9f, 0x1c,
	0xbe, 0xfc, 0x8c, 0x1f, 0xc7, 0x29, 0x6d, 0x49, 0x7a, 0xab, 0x81, 0x1b, 0xf8, 0x1d, 0xd2, 0x0b,
	0xb6, 0xfc, 0x98, 0x85, 0xc4, 0x33, 0x7c, 0x0e, 0x80, 0x26, 0xbb, 0xdc, 0x06, 0x51, 0xc9, 0x1e,
	0x21, 0xdf, 0xa3, 0xaf, 0xe4, 0xe8, 0x1e, 0x8c, 0x67, 0x39, 0x0d, 0x97, 0x17, 0xb1, 0x53, 0x9c,
	0x6d, 0x50, 0xf3, 0xff, 0x9c, 0x24, 0x0f, 0x02, 0x7d, 0x41, 0xb4, 0xba, 0xcc, 0x9d, 0xe1, 0x7c,
	0x0b, 0x9b, 0x15, 0x1f, 0xe4, 0xb8, 0x52, 0x4a, 0x4b, 0xeb, 0x1f, 0x2b, 0xba, 0x41, 0x84, 0x38,
	0xd1, 0xd5, 0x0e, 0xa4, 0xef, 0xba, 0x83, 0xba, 0x13, 0xfd, 0xbf, 0xae, 0x08, 0x59, 0x75, 0xfd,
	0x41, 0xce, 0xae, 0x96, 0xa4, 0x55, 0x6d, 0x84, 0xe5, 0x06, 0x9e, 0xb4, 0x6c, 0x86, 0xda, 0x60,
	0xdc, 0x81, 0x90, 0x58, 0xed, 0x40, 0x0c, 0x3a, 0xd1, 0xfb, 0xc1, 0xd8, 0xc1, 0x9b, 0xef, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x12, 0x48, 0x80, 0x09, 0xd4, 0x02, 0x00, 0x00,
}
