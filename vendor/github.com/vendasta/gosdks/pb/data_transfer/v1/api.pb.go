// Code generated by protoc-gen-go.
// source: api.proto
// DO NOT EDIT!

/*
Package datatransfer_v1 is a generated protocol buffer package.

It is generated from these files:
	api.proto

It has these top-level messages:
	StartTransferRequest
	StartTransferResponse
*/
package datatransfer_v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/empty"

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

// A request to start a data transfer
type StartTransferRequest struct {
	// The URL of the file to be transferred (beginning with http or https)
	SourceUrl string `protobuf:"bytes,1,opt,name=source_url,json=sourceUrl" json:"source_url,omitempty"`
	// The name of the cloud storage bucket to place the transferred file in
	DestinationBucketName string `protobuf:"bytes,2,opt,name=destination_bucket_name,json=destinationBucketName" json:"destination_bucket_name,omitempty"`
	// The destination filename to be saved in cloud storage
	DestinationFilename string `protobuf:"bytes,3,opt,name=destination_filename,json=destinationFilename" json:"destination_filename,omitempty"`
	// An arbitrary label for the transfer
	Label string `protobuf:"bytes,4,opt,name=label" json:"label,omitempty"`
}

func (m *StartTransferRequest) Reset()                    { *m = StartTransferRequest{} }
func (m *StartTransferRequest) String() string            { return proto.CompactTextString(m) }
func (*StartTransferRequest) ProtoMessage()               {}
func (*StartTransferRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *StartTransferRequest) GetSourceUrl() string {
	if m != nil {
		return m.SourceUrl
	}
	return ""
}

func (m *StartTransferRequest) GetDestinationBucketName() string {
	if m != nil {
		return m.DestinationBucketName
	}
	return ""
}

func (m *StartTransferRequest) GetDestinationFilename() string {
	if m != nil {
		return m.DestinationFilename
	}
	return ""
}

func (m *StartTransferRequest) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

type StartTransferResponse struct {
	// Unique identifier for the data transfer
	DataTransferId string `protobuf:"bytes,1,opt,name=data_transfer_id,json=dataTransferId" json:"data_transfer_id,omitempty"`
}

func (m *StartTransferResponse) Reset()                    { *m = StartTransferResponse{} }
func (m *StartTransferResponse) String() string            { return proto.CompactTextString(m) }
func (*StartTransferResponse) ProtoMessage()               {}
func (*StartTransferResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *StartTransferResponse) GetDataTransferId() string {
	if m != nil {
		return m.DataTransferId
	}
	return ""
}

func init() {
	proto.RegisterType((*StartTransferRequest)(nil), "datatransfer.v1.StartTransferRequest")
	proto.RegisterType((*StartTransferResponse)(nil), "datatransfer.v1.StartTransferResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for DataTransferService service

type DataTransferServiceClient interface {
	// StartTransfer begins a data transfer from the provided source_url to the destination bucket
	StartTransfer(ctx context.Context, in *StartTransferRequest, opts ...grpc.CallOption) (*StartTransferResponse, error)
}

type dataTransferServiceClient struct {
	cc *grpc.ClientConn
}

func NewDataTransferServiceClient(cc *grpc.ClientConn) DataTransferServiceClient {
	return &dataTransferServiceClient{cc}
}

func (c *dataTransferServiceClient) StartTransfer(ctx context.Context, in *StartTransferRequest, opts ...grpc.CallOption) (*StartTransferResponse, error) {
	out := new(StartTransferResponse)
	err := grpc.Invoke(ctx, "/datatransfer.v1.DataTransferService/StartTransfer", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DataTransferService service

type DataTransferServiceServer interface {
	// StartTransfer begins a data transfer from the provided source_url to the destination bucket
	StartTransfer(context.Context, *StartTransferRequest) (*StartTransferResponse, error)
}

func RegisterDataTransferServiceServer(s *grpc.Server, srv DataTransferServiceServer) {
	s.RegisterService(&_DataTransferService_serviceDesc, srv)
}

func _DataTransferService_StartTransfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartTransferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataTransferServiceServer).StartTransfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/datatransfer.v1.DataTransferService/StartTransfer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataTransferServiceServer).StartTransfer(ctx, req.(*StartTransferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DataTransferService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "datatransfer.v1.DataTransferService",
	HandlerType: (*DataTransferServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StartTransfer",
			Handler:    _DataTransferService_StartTransfer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

func init() { proto.RegisterFile("api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 267 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x90, 0x5d, 0x4b, 0xc3, 0x30,
	0x14, 0x86, 0xad, 0x5f, 0xd0, 0x03, 0x7e, 0x90, 0x75, 0x58, 0x26, 0x82, 0x14, 0x94, 0x5d, 0x75,
	0x4c, 0xc1, 0x7b, 0x45, 0x04, 0x6f, 0xbc, 0xd8, 0xf4, 0xba, 0xa6, 0xed, 0xe9, 0x08, 0xa6, 0x49,
	0x4d, 0x4e, 0x26, 0xfe, 0x31, 0x7f, 0x9f, 0x2c, 0xed, 0xa0, 0x0e, 0x61, 0x97, 0x39, 0xef, 0xf3,
	0x86, 0x97, 0x07, 0x42, 0xde, 0x88, 0xb4, 0x31, 0x9a, 0x34, 0x3b, 0x29, 0x39, 0x71, 0x32, 0x5c,
	0xd9, 0x0a, 0x4d, 0xba, 0x9c, 0x8e, 0xce, 0x17, 0x5a, 0x2f, 0x24, 0x4e, 0x7c, 0x9c, 0xbb, 0x6a,
	0x82, 0x75, 0x43, 0xdf, 0x2d, 0x9d, 0xfc, 0x04, 0x10, 0xcd, 0x89, 0x1b, 0x7a, 0xed, 0x1a, 0x33,
	0xfc, 0x74, 0x68, 0x89, 0x5d, 0x00, 0x58, 0xed, 0x4c, 0x81, 0x99, 0x33, 0x32, 0x0e, 0x2e, 0x83,
	0x71, 0x38, 0x0b, 0xdb, 0xcb, 0x9b, 0x91, 0xec, 0x0e, 0xce, 0x4a, 0xb4, 0x24, 0x14, 0x27, 0xa1,
	0x55, 0x96, 0xbb, 0xe2, 0x03, 0x29, 0x53, 0xbc, 0xc6, 0x78, 0xd7, 0xb3, 0xc3, 0x5e, 0xfc, 0xe0,
	0xd3, 0x17, 0x5e, 0x23, 0x9b, 0x42, 0xd4, 0xef, 0x55, 0x42, 0xa2, 0x2f, 0xed, 0xf9, 0xd2, 0xa0,
	0x97, 0x3d, 0x75, 0x11, 0x8b, 0xe0, 0x40, 0xf2, 0x1c, 0x65, 0xbc, 0xef, 0x99, 0xf6, 0x91, 0xdc,
	0xc3, 0x70, 0x63, 0xb7, 0x6d, 0xb4, 0xb2, 0xc8, 0xc6, 0x70, 0xba, 0x32, 0x90, 0xad, 0x15, 0x64,
	0xa2, 0xec, 0xe6, 0x1f, 0xaf, 0xee, 0x6b, 0xfe, 0xb9, 0xbc, 0xf9, 0x82, 0xc1, 0x63, 0xef, 0x32,
	0x47, 0xb3, 0x14, 0x05, 0xb2, 0x77, 0x38, 0xfa, 0xf3, 0x33, 0xbb, 0x4a, 0x37, 0x94, 0xa6, 0xff,
	0x19, 0x1b, 0x5d, 0x6f, 0xc3, 0xda, 0x81, 0xc9, 0x4e, 0x7e, 0xe8, 0xdd, 0xdf, 0xfe, 0x06, 0x00,
	0x00, 0xff, 0xff, 0xcb, 0x58, 0x8f, 0xf9, 0xb6, 0x01, 0x00, 0x00,
}
