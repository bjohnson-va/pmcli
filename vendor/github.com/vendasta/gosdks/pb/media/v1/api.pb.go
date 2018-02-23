// Code generated by protoc-gen-go.
// source: api.proto
// DO NOT EDIT!

/*
Package media_v1 is a generated protocol buffer package.

It is generated from these files:
	api.proto

It has these top-level messages:
	CreateImageRequest
	CreateImageResponse
	Tag
*/
package media_v1

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

// The CreateImageRequest message comes from a client to create an image, after it has been uploaded
type CreateImageRequest struct {
	// The serving_url that is generated after an image is uploaded to google cloud storage
	ServingUrl string `protobuf:"bytes,1,opt,name=serving_url,json=servingUrl" json:"serving_url,omitempty"`
	// Tag identifiers for statistics
	Tags []*Tag `protobuf:"bytes,2,rep,name=tags" json:"tags,omitempty"`
}

func (m *CreateImageRequest) Reset()                    { *m = CreateImageRequest{} }
func (m *CreateImageRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateImageRequest) ProtoMessage()               {}
func (*CreateImageRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CreateImageRequest) GetServingUrl() string {
	if m != nil {
		return m.ServingUrl
	}
	return ""
}

func (m *CreateImageRequest) GetTags() []*Tag {
	if m != nil {
		return m.Tags
	}
	return nil
}

// The CreateImageResponse returns the image_id of the image that is indexed in vStore
type CreateImageResponse struct {
	// The image_id for the image indexed in vStore
	ImageId string `protobuf:"bytes,1,opt,name=image_id,json=imageId" json:"image_id,omitempty"`
}

func (m *CreateImageResponse) Reset()                    { *m = CreateImageResponse{} }
func (m *CreateImageResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateImageResponse) ProtoMessage()               {}
func (*CreateImageResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CreateImageResponse) GetImageId() string {
	if m != nil {
		return m.ImageId
	}
	return ""
}

// Key value pair
type Tag struct {
	Key   string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *Tag) Reset()                    { *m = Tag{} }
func (m *Tag) String() string            { return proto.CompactTextString(m) }
func (*Tag) ProtoMessage()               {}
func (*Tag) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Tag) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Tag) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateImageRequest)(nil), "media.v1.CreateImageRequest")
	proto.RegisterType((*CreateImageResponse)(nil), "media.v1.CreateImageResponse")
	proto.RegisterType((*Tag)(nil), "media.v1.Tag")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ImageService service

type ImageServiceClient interface {
	Create(ctx context.Context, in *CreateImageRequest, opts ...grpc.CallOption) (*CreateImageResponse, error)
}

type imageServiceClient struct {
	cc *grpc.ClientConn
}

func NewImageServiceClient(cc *grpc.ClientConn) ImageServiceClient {
	return &imageServiceClient{cc}
}

func (c *imageServiceClient) Create(ctx context.Context, in *CreateImageRequest, opts ...grpc.CallOption) (*CreateImageResponse, error) {
	out := new(CreateImageResponse)
	err := grpc.Invoke(ctx, "/media.v1.ImageService/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ImageService service

type ImageServiceServer interface {
	Create(context.Context, *CreateImageRequest) (*CreateImageResponse, error)
}

func RegisterImageServiceServer(s *grpc.Server, srv ImageServiceServer) {
	s.RegisterService(&_ImageService_serviceDesc, srv)
}

func _ImageService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/media.v1.ImageService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageServiceServer).Create(ctx, req.(*CreateImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ImageService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "media.v1.ImageService",
	HandlerType: (*ImageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ImageService_Create_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

func init() { proto.RegisterFile("api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 220 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0x2c, 0xc8, 0xd4,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xc8, 0x4d, 0x4d, 0xc9, 0x4c, 0xd4, 0x2b, 0x33, 0x54,
	0x8a, 0xe0, 0x12, 0x72, 0x2e, 0x4a, 0x4d, 0x2c, 0x49, 0xf5, 0xcc, 0x4d, 0x4c, 0x4f, 0x0d, 0x4a,
	0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x92, 0xe7, 0xe2, 0x2e, 0x4e, 0x2d, 0x2a, 0xcb, 0xcc, 0x4b,
	0x8f, 0x2f, 0x2d, 0xca, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0xe2, 0x82, 0x0a, 0x85, 0x16,
	0xe5, 0x08, 0x29, 0x72, 0xb1, 0x94, 0x24, 0xa6, 0x17, 0x4b, 0x30, 0x29, 0x30, 0x6b, 0x70, 0x1b,
	0xf1, 0xea, 0xc1, 0xcc, 0xd3, 0x0b, 0x49, 0x4c, 0x0f, 0x02, 0x4b, 0x29, 0x19, 0x70, 0x09, 0xa3,
	0x98, 0x5c, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x2a, 0x24, 0xc9, 0xc5, 0x91, 0x09, 0x12, 0x88, 0xcf,
	0x4c, 0x81, 0x9a, 0xcb, 0x0e, 0xe6, 0x7b, 0xa6, 0x28, 0xe9, 0x72, 0x31, 0x87, 0x24, 0xa6, 0x0b,
	0x09, 0x70, 0x31, 0x67, 0xa7, 0x56, 0x42, 0x25, 0x41, 0x4c, 0x21, 0x11, 0x2e, 0xd6, 0xb2, 0xc4,
	0x9c, 0xd2, 0x54, 0x09, 0x26, 0xb0, 0x18, 0x84, 0x63, 0x14, 0xce, 0xc5, 0x03, 0x36, 0x3a, 0x18,
	0xe4, 0xac, 0xe4, 0x54, 0x21, 0x77, 0x2e, 0x36, 0x88, 0x85, 0x42, 0x32, 0x08, 0xf7, 0x60, 0x7a,
	0x4e, 0x4a, 0x16, 0x87, 0x2c, 0xc4, 0x81, 0x4a, 0x0c, 0x49, 0x6c, 0xe0, 0x40, 0x32, 0x06, 0x04,
	0x00, 0x00, 0xff, 0xff, 0x6c, 0x0e, 0x34, 0x5e, 0x31, 0x01, 0x00, 0x00,
}