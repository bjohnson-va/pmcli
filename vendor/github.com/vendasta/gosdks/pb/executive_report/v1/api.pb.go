// Code generated by protoc-gen-go. DO NOT EDIT.
// source: executive_report/v1/api.proto

/*
Package executivereport_v1 is a generated protocol buffer package.

It is generated from these files:
	executive_report/v1/api.proto
	executive_report/v1/executive_report.proto

It has these top-level messages:
	ReportData
	ListItem
	BarChart
	GraphMarker
	ReportDataRequest
	ListReportDataRequest
	ProductReportEntity
	ListReportDataPagedResponse
	ListReportDatesRequest
	ListReportDatesResponse
*/
package executivereport_v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf1 "github.com/golang/protobuf/ptypes/empty"
import _ "google.golang.org/genproto/googleapis/api/annotations"

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

// Client API for ExecutiveReport service

type ExecutiveReportClient interface {
	// Accept data for a report.
	CreateReportData(ctx context.Context, in *ReportDataRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
	// List data for a report category.
	ListReportData(ctx context.Context, in *ListReportDataRequest, opts ...grpc.CallOption) (*ListReportDataPagedResponse, error)
	// List all dates (weekly/monthly) for a time period
	ListReportDates(ctx context.Context, in *ListReportDatesRequest, opts ...grpc.CallOption) (*ListReportDatesResponse, error)
}

type executiveReportClient struct {
	cc *grpc.ClientConn
}

func NewExecutiveReportClient(cc *grpc.ClientConn) ExecutiveReportClient {
	return &executiveReportClient{cc}
}

func (c *executiveReportClient) CreateReportData(ctx context.Context, in *ReportDataRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/executivereport.v1.ExecutiveReport/CreateReportData", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *executiveReportClient) ListReportData(ctx context.Context, in *ListReportDataRequest, opts ...grpc.CallOption) (*ListReportDataPagedResponse, error) {
	out := new(ListReportDataPagedResponse)
	err := grpc.Invoke(ctx, "/executivereport.v1.ExecutiveReport/ListReportData", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *executiveReportClient) ListReportDates(ctx context.Context, in *ListReportDatesRequest, opts ...grpc.CallOption) (*ListReportDatesResponse, error) {
	out := new(ListReportDatesResponse)
	err := grpc.Invoke(ctx, "/executivereport.v1.ExecutiveReport/ListReportDates", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ExecutiveReport service

type ExecutiveReportServer interface {
	// Accept data for a report.
	CreateReportData(context.Context, *ReportDataRequest) (*google_protobuf1.Empty, error)
	// List data for a report category.
	ListReportData(context.Context, *ListReportDataRequest) (*ListReportDataPagedResponse, error)
	// List all dates (weekly/monthly) for a time period
	ListReportDates(context.Context, *ListReportDatesRequest) (*ListReportDatesResponse, error)
}

func RegisterExecutiveReportServer(s *grpc.Server, srv ExecutiveReportServer) {
	s.RegisterService(&_ExecutiveReport_serviceDesc, srv)
}

func _ExecutiveReport_CreateReportData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutiveReportServer).CreateReportData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/executivereport.v1.ExecutiveReport/CreateReportData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutiveReportServer).CreateReportData(ctx, req.(*ReportDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExecutiveReport_ListReportData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListReportDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutiveReportServer).ListReportData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/executivereport.v1.ExecutiveReport/ListReportData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutiveReportServer).ListReportData(ctx, req.(*ListReportDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExecutiveReport_ListReportDates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListReportDatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutiveReportServer).ListReportDates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/executivereport.v1.ExecutiveReport/ListReportDates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutiveReportServer).ListReportDates(ctx, req.(*ListReportDatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ExecutiveReport_serviceDesc = grpc.ServiceDesc{
	ServiceName: "executivereport.v1.ExecutiveReport",
	HandlerType: (*ExecutiveReportServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateReportData",
			Handler:    _ExecutiveReport_CreateReportData_Handler,
		},
		{
			MethodName: "ListReportData",
			Handler:    _ExecutiveReport_ListReportData_Handler,
		},
		{
			MethodName: "ListReportDates",
			Handler:    _ExecutiveReport_ListReportDates_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "executive_report/v1/api.proto",
}

func init() { proto.RegisterFile("executive_report/v1/api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 277 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x92, 0x4d, 0xad, 0x48, 0x4d,
	0x2e, 0x2d, 0xc9, 0x2c, 0x4b, 0x8d, 0x2f, 0x4a, 0x2d, 0xc8, 0x2f, 0x2a, 0xd1, 0x2f, 0x33, 0xd4,
	0x4f, 0x2c, 0xc8, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x82, 0x4b, 0x43, 0x64, 0xf5,
	0xca, 0x0c, 0xa5, 0xb4, 0xb0, 0x69, 0x41, 0x17, 0x83, 0xe8, 0x97, 0x92, 0x4e, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x07, 0xf3, 0x92, 0x4a, 0xd3, 0xf4, 0x53, 0x73, 0x0b, 0x4a, 0x2a, 0xa1, 0x92,
	0x32, 0x50, 0xc9, 0xc4, 0x82, 0x4c, 0xfd, 0xc4, 0xbc, 0xbc, 0xfc, 0x92, 0xc4, 0x92, 0xcc, 0xfc,
	0xbc, 0x62, 0x88, 0xac, 0xd1, 0x52, 0x66, 0x2e, 0x7e, 0x57, 0x98, 0xa9, 0x41, 0x60, 0x43, 0x85,
	0x4a, 0xb8, 0x04, 0x9c, 0x8b, 0x52, 0x13, 0x4b, 0xa0, 0x7c, 0x97, 0xc4, 0x92, 0x44, 0x21, 0x55,
	0x3d, 0x4c, 0x37, 0xea, 0x21, 0xe4, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0xa4, 0xc4, 0xf4,
	0x20, 0xb6, 0xe9, 0xc1, 0x9c, 0xa2, 0xe7, 0x0a, 0x72, 0x8a, 0x92, 0x62, 0xd3, 0xe5, 0x27, 0x93,
	0x99, 0xa4, 0x95, 0xc4, 0x40, 0xbe, 0x80, 0x18, 0xa0, 0x9b, 0x92, 0x58, 0x92, 0xa8, 0x9f, 0x0c,
	0xb6, 0xc8, 0x8a, 0x51, 0x4b, 0xa8, 0x97, 0x91, 0x8b, 0xcf, 0x27, 0xb3, 0xb8, 0x04, 0xc9, 0x52,
	0x4d, 0x6c, 0x96, 0xa2, 0xaa, 0x81, 0x59, 0xac, 0x4f, 0x58, 0x69, 0x40, 0x62, 0x7a, 0x6a, 0x4a,
	0x50, 0x6a, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0xaa, 0x92, 0x3c, 0xd8, 0x45, 0x92, 0x4a, 0x22, 0xe8,
	0x2e, 0xca, 0xc9, 0x2c, 0x2e, 0x01, 0xb9, 0xa7, 0x9b, 0x91, 0x8b, 0x1f, 0xc5, 0x80, 0xd4, 0x62,
	0x21, 0x2d, 0x82, 0xb6, 0xa4, 0x16, 0xc3, 0x5c, 0xa4, 0x4d, 0x94, 0x5a, 0xfc, 0xae, 0x49, 0x85,
	0xb9, 0x26, 0x89, 0x0d, 0x1c, 0xa0, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xcd, 0x81, 0xa5,
	0xe0, 0x4a, 0x02, 0x00, 0x00,
}
