// Code generated by protoc-gen-go.
// source: directory_service.proto
// DO NOT EDIT!

/*
Package vendasta_listingsproto is a generated protocol buffer package.

It is generated from these files:
	directory_service.proto
	listing.proto
	review.proto

It has these top-level messages:
	Geo
	Listing
	Stats
	GetListingRequest
	DeleteListingRequest
	SearchListingRequest
	SearchListingResponse
	Review
	GetReviewRequest
	DeleteReviewRequest
	ListReviewsRequest
	ListReviewsResponse
*/
package vendasta_listingsproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import google_protobuf2 "github.com/golang/protobuf/ptypes/empty"

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

// Client API for ListingService service

type ListingServiceClient interface {
	Put(ctx context.Context, in *Listing, opts ...grpc.CallOption) (*Listing, error)
	Get(ctx context.Context, in *GetListingRequest, opts ...grpc.CallOption) (*Listing, error)
	Delete(ctx context.Context, in *DeleteListingRequest, opts ...grpc.CallOption) (*google_protobuf2.Empty, error)
	GetStats(ctx context.Context, in *google_protobuf2.Empty, opts ...grpc.CallOption) (*Stats, error)
	Search(ctx context.Context, in *SearchListingRequest, opts ...grpc.CallOption) (*SearchListingResponse, error)
}

type listingServiceClient struct {
	cc *grpc.ClientConn
}

func NewListingServiceClient(cc *grpc.ClientConn) ListingServiceClient {
	return &listingServiceClient{cc}
}

func (c *listingServiceClient) Put(ctx context.Context, in *Listing, opts ...grpc.CallOption) (*Listing, error) {
	out := new(Listing)
	err := grpc.Invoke(ctx, "/vendasta.listingsproto.ListingService/Put", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *listingServiceClient) Get(ctx context.Context, in *GetListingRequest, opts ...grpc.CallOption) (*Listing, error) {
	out := new(Listing)
	err := grpc.Invoke(ctx, "/vendasta.listingsproto.ListingService/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *listingServiceClient) Delete(ctx context.Context, in *DeleteListingRequest, opts ...grpc.CallOption) (*google_protobuf2.Empty, error) {
	out := new(google_protobuf2.Empty)
	err := grpc.Invoke(ctx, "/vendasta.listingsproto.ListingService/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *listingServiceClient) GetStats(ctx context.Context, in *google_protobuf2.Empty, opts ...grpc.CallOption) (*Stats, error) {
	out := new(Stats)
	err := grpc.Invoke(ctx, "/vendasta.listingsproto.ListingService/GetStats", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *listingServiceClient) Search(ctx context.Context, in *SearchListingRequest, opts ...grpc.CallOption) (*SearchListingResponse, error) {
	out := new(SearchListingResponse)
	err := grpc.Invoke(ctx, "/vendasta.listingsproto.ListingService/Search", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ListingService service

type ListingServiceServer interface {
	Put(context.Context, *Listing) (*Listing, error)
	Get(context.Context, *GetListingRequest) (*Listing, error)
	Delete(context.Context, *DeleteListingRequest) (*google_protobuf2.Empty, error)
	GetStats(context.Context, *google_protobuf2.Empty) (*Stats, error)
	Search(context.Context, *SearchListingRequest) (*SearchListingResponse, error)
}

func RegisterListingServiceServer(s *grpc.Server, srv ListingServiceServer) {
	s.RegisterService(&_ListingService_serviceDesc, srv)
}

func _ListingService_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Listing)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListingServiceServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vendasta.listingsproto.ListingService/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListingServiceServer).Put(ctx, req.(*Listing))
	}
	return interceptor(ctx, in, info, handler)
}

func _ListingService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListingServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vendasta.listingsproto.ListingService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListingServiceServer).Get(ctx, req.(*GetListingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ListingService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteListingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListingServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vendasta.listingsproto.ListingService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListingServiceServer).Delete(ctx, req.(*DeleteListingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ListingService_GetStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf2.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListingServiceServer).GetStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vendasta.listingsproto.ListingService/GetStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListingServiceServer).GetStats(ctx, req.(*google_protobuf2.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ListingService_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchListingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListingServiceServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vendasta.listingsproto.ListingService/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListingServiceServer).Search(ctx, req.(*SearchListingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ListingService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "vendasta.listingsproto.ListingService",
	HandlerType: (*ListingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Put",
			Handler:    _ListingService_Put_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ListingService_Get_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ListingService_Delete_Handler,
		},
		{
			MethodName: "GetStats",
			Handler:    _ListingService_GetStats_Handler,
		},
		{
			MethodName: "Search",
			Handler:    _ListingService_Search_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "directory_service.proto",
}

// Client API for ReviewService service

type ReviewServiceClient interface {
	Put(ctx context.Context, in *Review, opts ...grpc.CallOption) (*Review, error)
	Get(ctx context.Context, in *GetReviewRequest, opts ...grpc.CallOption) (*Review, error)
	Delete(ctx context.Context, in *DeleteReviewRequest, opts ...grpc.CallOption) (*google_protobuf2.Empty, error)
	List(ctx context.Context, in *ListReviewsRequest, opts ...grpc.CallOption) (*ListReviewsResponse, error)
}

type reviewServiceClient struct {
	cc *grpc.ClientConn
}

func NewReviewServiceClient(cc *grpc.ClientConn) ReviewServiceClient {
	return &reviewServiceClient{cc}
}

func (c *reviewServiceClient) Put(ctx context.Context, in *Review, opts ...grpc.CallOption) (*Review, error) {
	out := new(Review)
	err := grpc.Invoke(ctx, "/vendasta.listingsproto.ReviewService/Put", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewServiceClient) Get(ctx context.Context, in *GetReviewRequest, opts ...grpc.CallOption) (*Review, error) {
	out := new(Review)
	err := grpc.Invoke(ctx, "/vendasta.listingsproto.ReviewService/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewServiceClient) Delete(ctx context.Context, in *DeleteReviewRequest, opts ...grpc.CallOption) (*google_protobuf2.Empty, error) {
	out := new(google_protobuf2.Empty)
	err := grpc.Invoke(ctx, "/vendasta.listingsproto.ReviewService/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewServiceClient) List(ctx context.Context, in *ListReviewsRequest, opts ...grpc.CallOption) (*ListReviewsResponse, error) {
	out := new(ListReviewsResponse)
	err := grpc.Invoke(ctx, "/vendasta.listingsproto.ReviewService/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ReviewService service

type ReviewServiceServer interface {
	Put(context.Context, *Review) (*Review, error)
	Get(context.Context, *GetReviewRequest) (*Review, error)
	Delete(context.Context, *DeleteReviewRequest) (*google_protobuf2.Empty, error)
	List(context.Context, *ListReviewsRequest) (*ListReviewsResponse, error)
}

func RegisterReviewServiceServer(s *grpc.Server, srv ReviewServiceServer) {
	s.RegisterService(&_ReviewService_serviceDesc, srv)
}

func _ReviewService_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Review)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServiceServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vendasta.listingsproto.ReviewService/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServiceServer).Put(ctx, req.(*Review))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReviewService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReviewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vendasta.listingsproto.ReviewService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServiceServer).Get(ctx, req.(*GetReviewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReviewService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteReviewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vendasta.listingsproto.ReviewService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServiceServer).Delete(ctx, req.(*DeleteReviewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReviewService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListReviewsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vendasta.listingsproto.ReviewService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServiceServer).List(ctx, req.(*ListReviewsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ReviewService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "vendasta.listingsproto.ReviewService",
	HandlerType: (*ReviewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Put",
			Handler:    _ReviewService_Put_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ReviewService_Get_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ReviewService_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _ReviewService_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "directory_service.proto",
}

func init() { proto.RegisterFile("directory_service.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 417 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x93, 0xd1, 0x4e, 0xe2, 0x40,
	0x14, 0x86, 0xd3, 0x85, 0x90, 0xcd, 0x64, 0x61, 0xb3, 0xb3, 0x1b, 0xc8, 0x76, 0x61, 0x4d, 0xe6,
	0x4a, 0x8b, 0xb6, 0x11, 0xef, 0xb8, 0xd6, 0x70, 0xe3, 0x85, 0x81, 0x07, 0xc0, 0x52, 0x8e, 0xb5,
	0x0a, 0x9d, 0xda, 0x39, 0xad, 0x21, 0xde, 0x19, 0xdf, 0xc0, 0xc7, 0xf1, 0x01, 0x7c, 0x00, 0x5f,
	0xc1, 0x07, 0x31, 0x9d, 0x99, 0x2a, 0x10, 0x0a, 0x84, 0xcb, 0xe1, 0xff, 0xe7, 0xff, 0xe6, 0x9c,
	0xfe, 0x90, 0xc6, 0x38, 0x88, 0xc1, 0x43, 0x1e, 0xcf, 0x86, 0x02, 0xe2, 0x34, 0xf0, 0xc0, 0x8e,
	0x62, 0x8e, 0x9c, 0xd6, 0x53, 0x08, 0xc7, 0xae, 0x40, 0xd7, 0x9e, 0x04, 0x02, 0x83, 0xd0, 0x17,
	0xf2, 0x77, 0xb3, 0xaa, 0x8f, 0xca, 0x66, 0xfe, 0x88, 0x21, 0x0d, 0xe0, 0x5e, 0x9f, 0x9a, 0x3e,
	0xe7, 0xfe, 0x04, 0x1c, 0x37, 0x0a, 0x1c, 0x37, 0x0c, 0x39, 0xba, 0x18, 0xf0, 0x50, 0x68, 0xf5,
	0x9f, 0x56, 0xe5, 0x69, 0x94, 0x5c, 0x39, 0x30, 0x8d, 0x70, 0xa6, 0xc4, 0xce, 0x4b, 0x99, 0xd4,
	0xce, 0x55, 0xf4, 0x40, 0x3d, 0x84, 0x7a, 0xa4, 0x74, 0x91, 0x20, 0xdd, 0xb3, 0x57, 0x3f, 0xc5,
	0xd6, 0x76, 0x73, 0x93, 0x81, 0x99, 0x8f, 0x6f, 0xef, 0xcf, 0xdf, 0xfe, 0xb0, 0x9f, 0x4e, 0x7a,
	0xec, 0x68, 0x8b, 0x13, 0x25, 0xd8, 0x35, 0x2c, 0x3a, 0x25, 0xa5, 0x1e, 0x20, 0x3d, 0x28, 0xca,
	0xe8, 0x01, 0xea, 0x98, 0x3e, 0xdc, 0x25, 0x20, 0x70, 0x47, 0x9c, 0x0f, 0x1a, 0x57, 0x39, 0x85,
	0x09, 0x20, 0xd0, 0xc3, 0xa2, 0x18, 0xa5, 0x2f, 0x41, 0xeb, 0xb6, 0x5a, 0x9e, 0x9d, 0x2f, 0xcf,
	0x3e, 0xcb, 0x96, 0xc7, 0x5a, 0x92, 0xd5, 0x60, 0x74, 0x9e, 0x35, 0x96, 0x09, 0x19, 0x6e, 0x48,
	0xbe, 0xf7, 0x00, 0x07, 0xe8, 0xa2, 0xa0, 0x05, 0x11, 0x66, 0xab, 0xe8, 0x21, 0xf2, 0x1a, 0x6b,
	0x4a, 0x42, 0x9d, 0xfd, 0x9a, 0x27, 0x88, 0x4c, 0xca, 0x00, 0x4f, 0x06, 0xa9, 0x0c, 0xc0, 0x8d,
	0xbd, 0xeb, 0xe2, 0x81, 0x94, 0xbe, 0x34, 0xd0, 0xd1, 0x96, 0x6e, 0x11, 0xf1, 0x50, 0xc0, 0xea,
	0x39, 0x85, 0xb4, 0x76, 0x0d, 0xab, 0xf3, 0x5a, 0x22, 0xd5, 0xbe, 0x6c, 0x62, 0x5e, 0x9e, 0x4b,
	0x55, 0x9e, 0xff, 0x45, 0x18, 0xe5, 0x36, 0x37, 0xe8, 0xec, 0xaf, 0xe4, 0xfe, 0x66, 0xb5, 0x8c,
	0xab, 0xba, 0x9e, 0x37, 0xe7, 0x46, 0x35, 0x67, 0x7f, 0x4d, 0x73, 0x54, 0x48, 0x3e, 0xf2, 0x2e,
	0x2c, 0x5d, 0x9b, 0xdb, 0xcf, 0xda, 0xb4, 0xd7, 0xd7, 0x66, 0x91, 0x58, 0xd4, 0x9a, 0x85, 0x6f,
	0xaa, 0x49, 0x5f, 0xa5, 0x79, 0x20, 0xe5, 0x6c, 0xfd, 0xd4, 0x5a, 0x57, 0x74, 0x05, 0x12, 0x39,
	0xa9, 0xbd, 0x95, 0x57, 0x7f, 0xcc, 0x85, 0x3f, 0x88, 0xc6, 0x67, 0xb7, 0xba, 0x86, 0x35, 0xaa,
	0xc8, 0x6b, 0x27, 0x1f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x36, 0xf7, 0x04, 0x6c, 0x99, 0x04, 0x00,
	0x00,
}