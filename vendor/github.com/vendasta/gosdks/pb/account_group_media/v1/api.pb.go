// Code generated by protoc-gen-go. DO NOT EDIT.
// source: account_group_media/v1/api.proto

/*
Package accountgroupmedia_v1 is a generated protocol buffer package.

It is generated from these files:
	account_group_media/v1/api.proto

It has these top-level messages:
	AssociateImageRequest
	DeleteImageRequest
	ListImagesRequest
	ImageAssociation
	ListImagesPagedResponse
*/
package accountgroupmedia_v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"
import google_protobuf1 "github.com/golang/protobuf/ptypes/empty"

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

// The possible image types, that can be associated with an image
type ImageType int32

const (
	ImageType_NOT_SPECIFIED ImageType = 0
	ImageType_PRIMARY       ImageType = 1
	ImageType_LOGO          ImageType = 2
)

var ImageType_name = map[int32]string{
	0: "NOT_SPECIFIED",
	1: "PRIMARY",
	2: "LOGO",
}
var ImageType_value = map[string]int32{
	"NOT_SPECIFIED": 0,
	"PRIMARY":       1,
	"LOGO":          2,
}

func (x ImageType) String() string {
	return proto.EnumName(ImageType_name, int32(x))
}
func (ImageType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// AssociateImageRequest associates an image id in the media service to an account group.
// Also specifies if an image is of a specific type.
type AssociateImageRequest struct {
	// Account group ID of the account with which to associate the image
	AccountGroupId string `protobuf:"bytes,1,opt,name=account_group_id,json=accountGroupId" json:"account_group_id,omitempty"`
	// The unique image id to associate with the account group id
	ImageId string `protobuf:"bytes,2,opt,name=image_id,json=imageId" json:"image_id,omitempty"`
	// The type to associate with the image.
	ImageType ImageType `protobuf:"varint,3,opt,name=image_type,json=imageType,enum=accountgroupmedia.v1.ImageType" json:"image_type,omitempty"`
}

func (m *AssociateImageRequest) Reset()                    { *m = AssociateImageRequest{} }
func (m *AssociateImageRequest) String() string            { return proto.CompactTextString(m) }
func (*AssociateImageRequest) ProtoMessage()               {}
func (*AssociateImageRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AssociateImageRequest) GetAccountGroupId() string {
	if m != nil {
		return m.AccountGroupId
	}
	return ""
}

func (m *AssociateImageRequest) GetImageId() string {
	if m != nil {
		return m.ImageId
	}
	return ""
}

func (m *AssociateImageRequest) GetImageType() ImageType {
	if m != nil {
		return m.ImageType
	}
	return ImageType_NOT_SPECIFIED
}

// DeleteImageRequest deletes an image association
type DeleteImageRequest struct {
	// Account group ID of the account with which the image to be deleted is associated
	AccountGroupId string `protobuf:"bytes,1,opt,name=account_group_id,json=accountGroupId" json:"account_group_id,omitempty"`
	// The unique image id to of the image to delete
	ImageId string `protobuf:"bytes,2,opt,name=image_id,json=imageId" json:"image_id,omitempty"`
}

func (m *DeleteImageRequest) Reset()                    { *m = DeleteImageRequest{} }
func (m *DeleteImageRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteImageRequest) ProtoMessage()               {}
func (*DeleteImageRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *DeleteImageRequest) GetAccountGroupId() string {
	if m != nil {
		return m.AccountGroupId
	}
	return ""
}

func (m *DeleteImageRequest) GetImageId() string {
	if m != nil {
		return m.ImageId
	}
	return ""
}

// ListImagesRequest is a request to get a list of image association models associated with an account group id
type ListImagesRequest struct {
	// Account group ID of the account with which the desired images are associated
	AccountGroupId string `protobuf:"bytes,1,opt,name=account_group_id,json=accountGroupId" json:"account_group_id,omitempty"`
	// The cursor from the previous response, or ""
	Cursor string `protobuf:"bytes,2,opt,name=cursor" json:"cursor,omitempty"`
	// The number of results to return
	PageSize int64 `protobuf:"varint,3,opt,name=page_size,json=pageSize" json:"page_size,omitempty"`
	// Filters for images that meet specific criteria, instead of just getting all non-deleted images back.
	Filters *ListImagesRequest_Filters `protobuf:"bytes,4,opt,name=filters" json:"filters,omitempty"`
}

func (m *ListImagesRequest) Reset()                    { *m = ListImagesRequest{} }
func (m *ListImagesRequest) String() string            { return proto.CompactTextString(m) }
func (*ListImagesRequest) ProtoMessage()               {}
func (*ListImagesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ListImagesRequest) GetAccountGroupId() string {
	if m != nil {
		return m.AccountGroupId
	}
	return ""
}

func (m *ListImagesRequest) GetCursor() string {
	if m != nil {
		return m.Cursor
	}
	return ""
}

func (m *ListImagesRequest) GetPageSize() int64 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListImagesRequest) GetFilters() *ListImagesRequest_Filters {
	if m != nil {
		return m.Filters
	}
	return nil
}

type ListImagesRequest_Filters struct {
	ImageType      []ImageType `protobuf:"varint,1,rep,packed,name=image_type,json=imageType,enum=accountgroupmedia.v1.ImageType" json:"image_type,omitempty"`
	IncludeDeleted bool        `protobuf:"varint,2,opt,name=include_deleted,json=includeDeleted" json:"include_deleted,omitempty"`
}

func (m *ListImagesRequest_Filters) Reset()                    { *m = ListImagesRequest_Filters{} }
func (m *ListImagesRequest_Filters) String() string            { return proto.CompactTextString(m) }
func (*ListImagesRequest_Filters) ProtoMessage()               {}
func (*ListImagesRequest_Filters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

func (m *ListImagesRequest_Filters) GetImageType() []ImageType {
	if m != nil {
		return m.ImageType
	}
	return nil
}

func (m *ListImagesRequest_Filters) GetIncludeDeleted() bool {
	if m != nil {
		return m.IncludeDeleted
	}
	return false
}

// ImageAssociation is the pairing of an image id to an account group id
type ImageAssociation struct {
	// Account group ID of the account with which the image is associated
	AccountGroupId string `protobuf:"bytes,1,opt,name=account_group_id,json=accountGroupId" json:"account_group_id,omitempty"`
	// The unique image id of the image associated with the account group
	ImageId string `protobuf:"bytes,2,opt,name=image_id,json=imageId" json:"image_id,omitempty"`
	// The timestap for when the image association was created
	Created *google_protobuf.Timestamp `protobuf:"bytes,3,opt,name=created" json:"created,omitempty"`
	// What type of image
	ImageType ImageType `protobuf:"varint,4,opt,name=image_type,json=imageType,enum=accountgroupmedia.v1.ImageType" json:"image_type,omitempty"`
	// The timestamp for when the image association was deleted
	Deleted *google_protobuf.Timestamp `protobuf:"bytes,5,opt,name=deleted" json:"deleted,omitempty"`
}

func (m *ImageAssociation) Reset()                    { *m = ImageAssociation{} }
func (m *ImageAssociation) String() string            { return proto.CompactTextString(m) }
func (*ImageAssociation) ProtoMessage()               {}
func (*ImageAssociation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ImageAssociation) GetAccountGroupId() string {
	if m != nil {
		return m.AccountGroupId
	}
	return ""
}

func (m *ImageAssociation) GetImageId() string {
	if m != nil {
		return m.ImageId
	}
	return ""
}

func (m *ImageAssociation) GetCreated() *google_protobuf.Timestamp {
	if m != nil {
		return m.Created
	}
	return nil
}

func (m *ImageAssociation) GetImageType() ImageType {
	if m != nil {
		return m.ImageType
	}
	return ImageType_NOT_SPECIFIED
}

func (m *ImageAssociation) GetDeleted() *google_protobuf.Timestamp {
	if m != nil {
		return m.Deleted
	}
	return nil
}

// ListImagesPagedResponse is the list of image association models that are associated to an account group
type ListImagesPagedResponse struct {
	Images []*ImageAssociation `protobuf:"bytes,1,rep,name=images" json:"images,omitempty"`
	// A cursor that can be provided to retrieve the next page of results
	NextCursor string `protobuf:"bytes,2,opt,name=next_cursor,json=nextCursor" json:"next_cursor,omitempty"`
	// Whether or not more results exist
	HasMore bool `protobuf:"varint,3,opt,name=has_more,json=hasMore" json:"has_more,omitempty"`
	// The total number of results for this query across all pages
	TotalResults int64 `protobuf:"varint,4,opt,name=total_results,json=totalResults" json:"total_results,omitempty"`
}

func (m *ListImagesPagedResponse) Reset()                    { *m = ListImagesPagedResponse{} }
func (m *ListImagesPagedResponse) String() string            { return proto.CompactTextString(m) }
func (*ListImagesPagedResponse) ProtoMessage()               {}
func (*ListImagesPagedResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ListImagesPagedResponse) GetImages() []*ImageAssociation {
	if m != nil {
		return m.Images
	}
	return nil
}

func (m *ListImagesPagedResponse) GetNextCursor() string {
	if m != nil {
		return m.NextCursor
	}
	return ""
}

func (m *ListImagesPagedResponse) GetHasMore() bool {
	if m != nil {
		return m.HasMore
	}
	return false
}

func (m *ListImagesPagedResponse) GetTotalResults() int64 {
	if m != nil {
		return m.TotalResults
	}
	return 0
}

func init() {
	proto.RegisterType((*AssociateImageRequest)(nil), "accountgroupmedia.v1.AssociateImageRequest")
	proto.RegisterType((*DeleteImageRequest)(nil), "accountgroupmedia.v1.DeleteImageRequest")
	proto.RegisterType((*ListImagesRequest)(nil), "accountgroupmedia.v1.ListImagesRequest")
	proto.RegisterType((*ListImagesRequest_Filters)(nil), "accountgroupmedia.v1.ListImagesRequest.Filters")
	proto.RegisterType((*ImageAssociation)(nil), "accountgroupmedia.v1.ImageAssociation")
	proto.RegisterType((*ListImagesPagedResponse)(nil), "accountgroupmedia.v1.ListImagesPagedResponse")
	proto.RegisterEnum("accountgroupmedia.v1.ImageType", ImageType_name, ImageType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for MediaService service

type MediaServiceClient interface {
	AssociateImage(ctx context.Context, in *AssociateImageRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
	ListImages(ctx context.Context, in *ListImagesRequest, opts ...grpc.CallOption) (*ListImagesPagedResponse, error)
	DeleteImage(ctx context.Context, in *DeleteImageRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
}

type mediaServiceClient struct {
	cc *grpc.ClientConn
}

func NewMediaServiceClient(cc *grpc.ClientConn) MediaServiceClient {
	return &mediaServiceClient{cc}
}

func (c *mediaServiceClient) AssociateImage(ctx context.Context, in *AssociateImageRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/accountgroupmedia.v1.MediaService/AssociateImage", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mediaServiceClient) ListImages(ctx context.Context, in *ListImagesRequest, opts ...grpc.CallOption) (*ListImagesPagedResponse, error) {
	out := new(ListImagesPagedResponse)
	err := grpc.Invoke(ctx, "/accountgroupmedia.v1.MediaService/ListImages", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mediaServiceClient) DeleteImage(ctx context.Context, in *DeleteImageRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/accountgroupmedia.v1.MediaService/DeleteImage", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MediaService service

type MediaServiceServer interface {
	AssociateImage(context.Context, *AssociateImageRequest) (*google_protobuf1.Empty, error)
	ListImages(context.Context, *ListImagesRequest) (*ListImagesPagedResponse, error)
	DeleteImage(context.Context, *DeleteImageRequest) (*google_protobuf1.Empty, error)
}

func RegisterMediaServiceServer(s *grpc.Server, srv MediaServiceServer) {
	s.RegisterService(&_MediaService_serviceDesc, srv)
}

func _MediaService_AssociateImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssociateImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediaServiceServer).AssociateImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accountgroupmedia.v1.MediaService/AssociateImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediaServiceServer).AssociateImage(ctx, req.(*AssociateImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MediaService_ListImages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListImagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediaServiceServer).ListImages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accountgroupmedia.v1.MediaService/ListImages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediaServiceServer).ListImages(ctx, req.(*ListImagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MediaService_DeleteImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediaServiceServer).DeleteImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accountgroupmedia.v1.MediaService/DeleteImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediaServiceServer).DeleteImage(ctx, req.(*DeleteImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MediaService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "accountgroupmedia.v1.MediaService",
	HandlerType: (*MediaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AssociateImage",
			Handler:    _MediaService_AssociateImage_Handler,
		},
		{
			MethodName: "ListImages",
			Handler:    _MediaService_ListImages_Handler,
		},
		{
			MethodName: "DeleteImage",
			Handler:    _MediaService_DeleteImage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "account_group_media/v1/api.proto",
}

func init() { proto.RegisterFile("account_group_media/v1/api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 589 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xb4, 0x54, 0xdf, 0x4f, 0xd3, 0x50,
	0x14, 0xa6, 0x1b, 0xd2, 0xee, 0x0c, 0xe6, 0xb8, 0x51, 0x84, 0xf2, 0xc0, 0x52, 0x13, 0x59, 0x34,
	0x76, 0x61, 0xea, 0x2b, 0x09, 0xe1, 0x57, 0x9a, 0x80, 0xe0, 0x85, 0xc4, 0xf0, 0xd4, 0x94, 0xf6,
	0x30, 0x6e, 0xd2, 0xee, 0xd6, 0xde, 0x5b, 0x22, 0xbc, 0xfa, 0x6f, 0xf8, 0xe4, 0x7f, 0xe0, 0xab,
	0x7f, 0x9d, 0xe9, 0xbd, 0xad, 0x30, 0x36, 0xc5, 0x25, 0xfa, 0xb6, 0x7e, 0xf7, 0x3b, 0xbf, 0xbe,
	0xf3, 0xed, 0x40, 0x27, 0x08, 0x43, 0x9e, 0x0f, 0xa5, 0x3f, 0xc8, 0x78, 0x9e, 0xfa, 0x09, 0x46,
	0x2c, 0xe8, 0x5d, 0x6d, 0xf4, 0x82, 0x94, 0xb9, 0x69, 0xc6, 0x25, 0x27, 0x4f, 0x4a, 0x86, 0x22,
	0xa8, 0x77, 0xf7, 0x6a, 0xc3, 0x5e, 0x1b, 0x70, 0x3e, 0x88, 0xb1, 0xa7, 0x38, 0xe7, 0xf9, 0x45,
	0x4f, 0xb2, 0x04, 0x85, 0x0c, 0x92, 0x54, 0x87, 0xd9, 0xab, 0xf7, 0x09, 0x98, 0xa4, 0xf2, 0x5a,
	0x3f, 0x3a, 0x5f, 0x0d, 0x78, 0xba, 0x25, 0x04, 0x0f, 0x59, 0x20, 0xd1, 0x4b, 0x82, 0x01, 0x52,
	0xfc, 0x94, 0xa3, 0x90, 0xa4, 0x0b, 0xed, 0xd1, 0x8e, 0x58, 0xb4, 0x6c, 0x74, 0x8c, 0x6e, 0x83,
	0xb6, 0x4a, 0x7c, 0xbf, 0x80, 0xbd, 0x88, 0xac, 0x80, 0xc5, 0x8a, 0xc8, 0x82, 0x51, 0x53, 0x0c,
	0x53, 0x7d, 0x7b, 0x11, 0xd9, 0x04, 0xd0, 0x4f, 0xf2, 0x3a, 0xc5, 0xe5, 0x7a, 0xc7, 0xe8, 0xb6,
	0xfa, 0x6b, 0xee, 0xa4, 0x39, 0x5c, 0x55, 0xfc, 0xf4, 0x3a, 0x45, 0xda, 0x60, 0xd5, 0x4f, 0xe7,
	0x0c, 0xc8, 0x0e, 0xc6, 0xf8, 0x1f, 0x5a, 0x73, 0xbe, 0xd7, 0x60, 0xf1, 0x80, 0x09, 0xa9, 0x32,
	0x8b, 0xe9, 0x53, 0x2f, 0xc1, 0x5c, 0x98, 0x67, 0x82, 0x67, 0x65, 0xe2, 0xf2, 0x8b, 0xac, 0x42,
	0x23, 0x2d, 0x2a, 0x0a, 0x76, 0xa3, 0x27, 0xae, 0x53, 0xab, 0x00, 0x4e, 0xd8, 0x0d, 0x12, 0x0f,
	0xcc, 0x0b, 0x16, 0x4b, 0xcc, 0xc4, 0xf2, 0x6c, 0xc7, 0xe8, 0x36, 0xfb, 0xbd, 0xc9, 0x62, 0x8c,
	0x35, 0xe6, 0xee, 0xe9, 0x30, 0x5a, 0xc5, 0xdb, 0x19, 0x98, 0x25, 0x76, 0x4f, 0x65, 0xa3, 0x53,
	0x9f, 0x4e, 0x65, 0xb2, 0x0e, 0x8f, 0xd9, 0x30, 0x8c, 0xf3, 0x08, 0xfd, 0x48, 0xa9, 0xad, 0xc5,
	0xb2, 0x68, 0xab, 0x84, 0xf5, 0x0e, 0x22, 0xe7, 0x4b, 0x0d, 0xda, 0x2a, 0x43, 0x65, 0x19, 0xc6,
	0x87, 0xff, 0xc6, 0x28, 0x6f, 0xc1, 0x0c, 0x33, 0x0c, 0x8a, 0xd2, 0x75, 0x25, 0x8c, 0xed, 0x6a,
	0xdb, 0xba, 0x95, 0x6d, 0xdd, 0xd3, 0xca, 0xd7, 0xb4, 0xa2, 0xde, 0x1b, 0x7c, 0x76, 0x5a, 0x7b,
	0x15, 0x55, 0xab, 0x81, 0x1f, 0x3d, 0x5c, 0xb5, 0xa4, 0x3a, 0x3f, 0x0c, 0x78, 0x76, 0xbb, 0xa0,
	0xe3, 0x60, 0x80, 0x11, 0x45, 0x91, 0xf2, 0xa1, 0x40, 0xb2, 0x09, 0x73, 0x2a, 0xbd, 0x50, 0x6b,
	0x68, 0xf6, 0x5f, 0xfc, 0xa1, 0x9b, 0x3b, 0x22, 0xd2, 0x32, 0x8a, 0xac, 0x41, 0x73, 0x88, 0x9f,
	0xa5, 0x3f, 0x62, 0x2d, 0x28, 0xa0, 0x6d, 0x6d, 0xaf, 0x15, 0xb0, 0x2e, 0x03, 0xe1, 0x27, 0x3c,
	0xd3, 0xee, 0xb2, 0xa8, 0x79, 0x19, 0x88, 0x43, 0x9e, 0x21, 0x79, 0x0e, 0x0b, 0x92, 0xcb, 0x20,
	0xf6, 0x33, 0x14, 0x79, 0x2c, 0xb5, 0xc5, 0xea, 0x74, 0x5e, 0x81, 0x54, 0x63, 0x2f, 0xdf, 0x41,
	0xe3, 0x97, 0x14, 0x64, 0x11, 0x16, 0xde, 0x1f, 0x9d, 0xfa, 0x27, 0xc7, 0xbb, 0xdb, 0xde, 0x9e,
	0xb7, 0xbb, 0xd3, 0x9e, 0x21, 0x4d, 0x30, 0x8f, 0xa9, 0x77, 0xb8, 0x45, 0xcf, 0xda, 0x06, 0xb1,
	0x60, 0xf6, 0xe0, 0x68, 0xff, 0xa8, 0x5d, 0xeb, 0x7f, 0xab, 0xc1, 0xfc, 0x61, 0xd1, 0xfd, 0x09,
	0x66, 0x57, 0x2c, 0x44, 0xf2, 0x11, 0x5a, 0xa3, 0x77, 0x83, 0xbc, 0x9a, 0x3c, 0xea, 0xc4, 0xeb,
	0x62, 0x2f, 0x8d, 0x09, 0xbd, 0x5b, 0x5c, 0x25, 0x67, 0x86, 0x5c, 0x00, 0xdc, 0x8a, 0x4b, 0xd6,
	0xff, 0xf2, 0xff, 0x61, 0xbf, 0x7e, 0x88, 0x38, 0xb2, 0x27, 0x67, 0x86, 0x7c, 0x80, 0xe6, 0x9d,
	0xd3, 0x42, 0xba, 0x93, 0xe3, 0xc7, 0xaf, 0xcf, 0xef, 0x5b, 0x3f, 0x9f, 0x53, 0xc8, 0x9b, 0x9f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x8e, 0x72, 0x72, 0xb3, 0xcb, 0x05, 0x00, 0x00,
}
