// Code generated by protoc-gen-go.
// source: api.proto
// DO NOT EDIT!

package nap_v1

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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for NAPDataService service

type NAPDataServiceClient interface {
	ParsePhoneNumber(ctx context.Context, in *ParsePhoneNumberRequest, opts ...grpc.CallOption) (*ParsePhoneNumberResponse, error)
	ListCountries(ctx context.Context, in *ListCountriesRequest, opts ...grpc.CallOption) (*ListCountriesResponse, error)
	ListStates(ctx context.Context, in *ListStatesRequest, opts ...grpc.CallOption) (*ListStatesResponse, error)
}

type nAPDataServiceClient struct {
	cc *grpc.ClientConn
}

func NewNAPDataServiceClient(cc *grpc.ClientConn) NAPDataServiceClient {
	return &nAPDataServiceClient{cc}
}

func (c *nAPDataServiceClient) ParsePhoneNumber(ctx context.Context, in *ParsePhoneNumberRequest, opts ...grpc.CallOption) (*ParsePhoneNumberResponse, error) {
	out := new(ParsePhoneNumberResponse)
	err := grpc.Invoke(ctx, "/nap.v1.NAPDataService/ParsePhoneNumber", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nAPDataServiceClient) ListCountries(ctx context.Context, in *ListCountriesRequest, opts ...grpc.CallOption) (*ListCountriesResponse, error) {
	out := new(ListCountriesResponse)
	err := grpc.Invoke(ctx, "/nap.v1.NAPDataService/ListCountries", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nAPDataServiceClient) ListStates(ctx context.Context, in *ListStatesRequest, opts ...grpc.CallOption) (*ListStatesResponse, error) {
	out := new(ListStatesResponse)
	err := grpc.Invoke(ctx, "/nap.v1.NAPDataService/ListStates", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for NAPDataService service

type NAPDataServiceServer interface {
	ParsePhoneNumber(context.Context, *ParsePhoneNumberRequest) (*ParsePhoneNumberResponse, error)
	ListCountries(context.Context, *ListCountriesRequest) (*ListCountriesResponse, error)
	ListStates(context.Context, *ListStatesRequest) (*ListStatesResponse, error)
}

func RegisterNAPDataServiceServer(s *grpc.Server, srv NAPDataServiceServer) {
	s.RegisterService(&_NAPDataService_serviceDesc, srv)
}

func _NAPDataService_ParsePhoneNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParsePhoneNumberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NAPDataServiceServer).ParsePhoneNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nap.v1.NAPDataService/ParsePhoneNumber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NAPDataServiceServer).ParsePhoneNumber(ctx, req.(*ParsePhoneNumberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NAPDataService_ListCountries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCountriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NAPDataServiceServer).ListCountries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nap.v1.NAPDataService/ListCountries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NAPDataServiceServer).ListCountries(ctx, req.(*ListCountriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NAPDataService_ListStates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListStatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NAPDataServiceServer).ListStates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nap.v1.NAPDataService/ListStates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NAPDataServiceServer).ListStates(ctx, req.(*ListStatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NAPDataService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "nap.v1.NAPDataService",
	HandlerType: (*NAPDataServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ParsePhoneNumber",
			Handler:    _NAPDataService_ParsePhoneNumber_Handler,
		},
		{
			MethodName: "ListCountries",
			Handler:    _NAPDataService_ListCountries_Handler,
		},
		{
			MethodName: "ListStates",
			Handler:    _NAPDataService_ListStates_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

func init() { proto.RegisterFile("api.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 212 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x90, 0x3f, 0x4b, 0xc6, 0x30,
	0x10, 0xc6, 0xd5, 0xe1, 0x05, 0x03, 0x15, 0xcd, 0x66, 0x50, 0x94, 0x0e, 0x8e, 0x01, 0xf5, 0x13,
	0xf8, 0x6f, 0x93, 0x12, 0xec, 0xe0, 0x7c, 0x6d, 0x0e, 0xcd, 0xd0, 0x24, 0xe6, 0xae, 0x9d, 0xfd,
	0xe8, 0xd2, 0x86, 0x48, 0x95, 0xb7, 0xe3, 0xdd, 0xef, 0x79, 0x7e, 0xc3, 0x23, 0x8e, 0x21, 0x3a,
	0x1d, 0x53, 0xe0, 0x20, 0x77, 0x1e, 0xa2, 0x9e, 0x6e, 0xd5, 0x59, 0xfc, 0x0c, 0x1e, 0xfd, 0x38,
	0x74, 0x98, 0x32, 0x52, 0x15, 0x58, 0x9b, 0x90, 0x28, 0x9f, 0x77, 0xdf, 0x47, 0xe2, 0xa4, 0x79,
	0x30, 0xcf, 0xc0, 0xd0, 0x62, 0x9a, 0x5c, 0x8f, 0xf2, 0x5d, 0x9c, 0x1a, 0x48, 0x84, 0x66, 0xee,
	0x36, 0x4b, 0x57, 0x5e, 0xe9, 0x6c, 0xd4, 0xff, 0xc9, 0x1b, 0x7e, 0x8d, 0x48, 0xac, 0xae, 0xb7,
	0x03, 0x14, 0x83, 0x27, 0xac, 0x0f, 0x64, 0x23, 0xaa, 0x57, 0x47, 0xfc, 0x14, 0x46, 0xcf, 0xc9,
	0x21, 0xc9, 0x8b, 0x52, 0xfa, 0xf3, 0x2e, 0xca, 0xcb, 0x0d, 0xfa, 0xeb, 0x7b, 0x11, 0x62, 0x46,
	0x2d, 0x03, 0x23, 0xc9, 0xf3, 0x75, 0x3c, 0xff, 0x8a, 0x49, 0xed, 0x43, 0x45, 0xf3, 0x78, 0x23,
	0xea, 0x3e, 0x0c, 0x7a, 0x42, 0x6f, 0x81, 0x18, 0xf4, 0x7a, 0xb3, 0x0f, 0xf4, 0x98, 0x80, 0xd1,
	0x9a, 0xc3, 0x6e, 0xb7, 0x2c, 0x76, 0xff, 0x13, 0x00, 0x00, 0xff, 0xff, 0x1b, 0x05, 0x48, 0xfe,
	0x68, 0x01, 0x00, 0x00,
}
