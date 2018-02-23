// Code generated by protoc-gen-go. DO NOT EDIT.
// source: listing_search/v1/api.proto

/*
Package listingsearch_v1 is a generated protocol buffer package.

It is generated from these files:
	listing_search/v1/api.proto

It has these top-level messages:
	Geo
	Location
	SearchRequest
	SearchResponse
	GetRequest
	GetResponse
*/
package listingsearch_v1

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

// Geographical location
type Geo struct {
	Latitude  float64 `protobuf:"fixed64,1,opt,name=latitude" json:"latitude,omitempty"`
	Longitude float64 `protobuf:"fixed64,2,opt,name=longitude" json:"longitude,omitempty"`
}

func (m *Geo) Reset()                    { *m = Geo{} }
func (m *Geo) String() string            { return proto.CompactTextString(m) }
func (*Geo) ProtoMessage()               {}
func (*Geo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Geo) GetLatitude() float64 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *Geo) GetLongitude() float64 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

// Location is a holder for search results from the corresponding api that a search is for
type Location struct {
	// The id that represents the location within the underlying service
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// The name of the location
	CompanyName string `protobuf:"bytes,2,opt,name=company_name,json=companyName" json:"company_name,omitempty"`
	// The postal address of the location
	Address string `protobuf:"bytes,3,opt,name=address" json:"address,omitempty"`
	// This typically refers to a city, but may refer to a suburb or a neighborhood in certain countries
	Locality string `protobuf:"bytes,4,opt,name=locality" json:"locality,omitempty"`
	// The state or territory abbreviation
	State string `protobuf:"bytes,5,opt,name=state" json:"state,omitempty"`
	// The country code
	Country string `protobuf:"bytes,6,opt,name=country" json:"country,omitempty"`
	// A string specifying the zip code or postal code
	ZipCode string `protobuf:"bytes,7,opt,name=zip_code,json=zipCode" json:"zip_code,omitempty"`
	// A location on the Earth specified by a latitude and longitude
	Point *Geo `protobuf:"bytes,8,opt,name=point" json:"point,omitempty"`
	// The website of the location
	Website string `protobuf:"bytes,9,opt,name=website" json:"website,omitempty"`
	// The primary phone number of the location
	Phone string `protobuf:"bytes,10,opt,name=phone" json:"phone,omitempty"`
	// The url location of where you can find the location on the internet
	Url string `protobuf:"bytes,11,opt,name=url" json:"url,omitempty"`
}

func (m *Location) Reset()                    { *m = Location{} }
func (m *Location) String() string            { return proto.CompactTextString(m) }
func (*Location) ProtoMessage()               {}
func (*Location) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Location) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Location) GetCompanyName() string {
	if m != nil {
		return m.CompanyName
	}
	return ""
}

func (m *Location) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Location) GetLocality() string {
	if m != nil {
		return m.Locality
	}
	return ""
}

func (m *Location) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *Location) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *Location) GetZipCode() string {
	if m != nil {
		return m.ZipCode
	}
	return ""
}

func (m *Location) GetPoint() *Geo {
	if m != nil {
		return m.Point
	}
	return nil
}

func (m *Location) GetWebsite() string {
	if m != nil {
		return m.Website
	}
	return ""
}

func (m *Location) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *Location) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

// A request to perform a search on behalf of a business
type SearchRequest struct {
	CompanyName string `protobuf:"bytes,1,opt,name=company_name,json=companyName" json:"company_name,omitempty"`
	Address     string `protobuf:"bytes,2,opt,name=address" json:"address,omitempty"`
	Locality    string `protobuf:"bytes,3,opt,name=locality" json:"locality,omitempty"`
	State       string `protobuf:"bytes,4,opt,name=state" json:"state,omitempty"`
	Country     string `protobuf:"bytes,5,opt,name=country" json:"country,omitempty"`
	ZipCode     string `protobuf:"bytes,6,opt,name=zip_code,json=zipCode" json:"zip_code,omitempty"`
	Phone       string `protobuf:"bytes,7,opt,name=phone" json:"phone,omitempty"`
	Location    *Geo   `protobuf:"bytes,8,opt,name=location" json:"location,omitempty"`
}

func (m *SearchRequest) Reset()                    { *m = SearchRequest{} }
func (m *SearchRequest) String() string            { return proto.CompactTextString(m) }
func (*SearchRequest) ProtoMessage()               {}
func (*SearchRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *SearchRequest) GetCompanyName() string {
	if m != nil {
		return m.CompanyName
	}
	return ""
}

func (m *SearchRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *SearchRequest) GetLocality() string {
	if m != nil {
		return m.Locality
	}
	return ""
}

func (m *SearchRequest) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *SearchRequest) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *SearchRequest) GetZipCode() string {
	if m != nil {
		return m.ZipCode
	}
	return ""
}

func (m *SearchRequest) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *SearchRequest) GetLocation() *Geo {
	if m != nil {
		return m.Location
	}
	return nil
}

// SearchResponse returns the location data that was found for the related search request
type SearchResponse struct {
	Locations []*Location `protobuf:"bytes,1,rep,name=locations" json:"locations,omitempty"`
}

func (m *SearchResponse) Reset()                    { *m = SearchResponse{} }
func (m *SearchResponse) String() string            { return proto.CompactTextString(m) }
func (*SearchResponse) ProtoMessage()               {}
func (*SearchResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *SearchResponse) GetLocations() []*Location {
	if m != nil {
		return m.Locations
	}
	return nil
}

// A request to lookup details of a business
type GetRequest struct {
	// The id that represents the location within the underlying service
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *GetRequest) Reset()                    { *m = GetRequest{} }
func (m *GetRequest) String() string            { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()               {}
func (*GetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *GetRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// GetResponse returns the location data for the related get request
type GetResponse struct {
	Location *Location `protobuf:"bytes,1,opt,name=location" json:"location,omitempty"`
}

func (m *GetResponse) Reset()                    { *m = GetResponse{} }
func (m *GetResponse) String() string            { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()               {}
func (*GetResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *GetResponse) GetLocation() *Location {
	if m != nil {
		return m.Location
	}
	return nil
}

func init() {
	proto.RegisterType((*Geo)(nil), "listingsearch.v1.Geo")
	proto.RegisterType((*Location)(nil), "listingsearch.v1.Location")
	proto.RegisterType((*SearchRequest)(nil), "listingsearch.v1.SearchRequest")
	proto.RegisterType((*SearchResponse)(nil), "listingsearch.v1.SearchResponse")
	proto.RegisterType((*GetRequest)(nil), "listingsearch.v1.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "listingsearch.v1.GetResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for SearchAdapter service

type SearchAdapterClient interface {
	Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
}

type searchAdapterClient struct {
	cc *grpc.ClientConn
}

func NewSearchAdapterClient(cc *grpc.ClientConn) SearchAdapterClient {
	return &searchAdapterClient{cc}
}

func (c *searchAdapterClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := grpc.Invoke(ctx, "/listingsearch.v1.SearchAdapter/Search", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchAdapterClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := grpc.Invoke(ctx, "/listingsearch.v1.SearchAdapter/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SearchAdapter service

type SearchAdapterServer interface {
	Search(context.Context, *SearchRequest) (*SearchResponse, error)
	Get(context.Context, *GetRequest) (*GetResponse, error)
}

func RegisterSearchAdapterServer(s *grpc.Server, srv SearchAdapterServer) {
	s.RegisterService(&_SearchAdapter_serviceDesc, srv)
}

func _SearchAdapter_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchAdapterServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/listingsearch.v1.SearchAdapter/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchAdapterServer).Search(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SearchAdapter_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchAdapterServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/listingsearch.v1.SearchAdapter/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchAdapterServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SearchAdapter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "listingsearch.v1.SearchAdapter",
	HandlerType: (*SearchAdapterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Search",
			Handler:    _SearchAdapter_Search_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _SearchAdapter_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "listing_search/v1/api.proto",
}

func init() { proto.RegisterFile("listing_search/v1/api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 451 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x53, 0x5d, 0x6f, 0xd3, 0x30,
	0x14, 0xc5, 0xc9, 0xfa, 0x75, 0x03, 0xd3, 0x64, 0x81, 0x64, 0x4a, 0x11, 0x25, 0x4f, 0x95, 0x90,
	0x3a, 0xb5, 0x48, 0x88, 0x37, 0x84, 0x00, 0x55, 0x42, 0xc0, 0x43, 0xf8, 0x01, 0x95, 0x17, 0x5f,
	0x6d, 0x96, 0x52, 0xdb, 0xc4, 0xee, 0x50, 0xf7, 0x7b, 0x78, 0xe5, 0x1f, 0xf2, 0x80, 0x62, 0x27,
	0xd9, 0xd6, 0x2e, 0x7d, 0xcb, 0xb9, 0xe7, 0xe6, 0xdc, 0x73, 0x8f, 0x6d, 0x78, 0x51, 0x48, 0xeb,
	0xa4, 0xba, 0x5c, 0x5b, 0xe4, 0x65, 0x7e, 0x75, 0x7e, 0xbd, 0x38, 0xe7, 0x46, 0xce, 0x4d, 0xa9,
	0x9d, 0xa6, 0x67, 0x35, 0x19, 0xb8, 0xf9, 0xf5, 0x22, 0xfd, 0x00, 0xf1, 0x0a, 0x35, 0x1d, 0xc3,
	0xb0, 0xe0, 0x4e, 0xba, 0xad, 0x40, 0x46, 0xa6, 0x64, 0x46, 0xb2, 0x16, 0xd3, 0x09, 0x8c, 0x0a,
	0xad, 0x2e, 0x03, 0x19, 0x79, 0xf2, 0xb6, 0x90, 0xfe, 0x8d, 0x60, 0xf8, 0x4d, 0xe7, 0xdc, 0x49,
	0xad, 0xe8, 0x29, 0x44, 0x52, 0x78, 0x81, 0x51, 0x16, 0x49, 0x41, 0x5f, 0xc3, 0xe3, 0x5c, 0x6f,
	0x0c, 0x57, 0xbb, 0xb5, 0xe2, 0x9b, 0xf0, 0xf7, 0x28, 0x4b, 0xea, 0xda, 0x0f, 0xbe, 0x41, 0xca,
	0x60, 0xc0, 0x85, 0x28, 0xd1, 0x5a, 0x16, 0x7b, 0xb6, 0x81, 0xde, 0x93, 0xce, 0x79, 0x21, 0xdd,
	0x8e, 0x9d, 0x78, 0xaa, 0xc5, 0xf4, 0x29, 0xf4, 0xac, 0xe3, 0x0e, 0x59, 0xcf, 0x13, 0x01, 0x54,
	0x5a, 0xb9, 0xde, 0x2a, 0x57, 0xee, 0x58, 0x3f, 0x68, 0xd5, 0x90, 0x3e, 0x87, 0xe1, 0x8d, 0x34,
	0xeb, 0x5c, 0x0b, 0x64, 0x83, 0x40, 0xdd, 0x48, 0xf3, 0x49, 0x0b, 0xa4, 0x6f, 0xa0, 0x67, 0xb4,
	0x54, 0x8e, 0x0d, 0xa7, 0x64, 0x96, 0x2c, 0x9f, 0xcd, 0xf7, 0x33, 0x9a, 0xaf, 0x50, 0x67, 0xa1,
	0xa7, 0x9a, 0xf0, 0x1b, 0x2f, 0xac, 0x74, 0xc8, 0x46, 0x41, 0xa6, 0x86, 0x95, 0x23, 0x73, 0xa5,
	0x15, 0x32, 0x08, 0x8e, 0x3c, 0xa0, 0x67, 0x10, 0x6f, 0xcb, 0x82, 0x25, 0xbe, 0x56, 0x7d, 0xa6,
	0xff, 0x08, 0x3c, 0xf9, 0xe9, 0xa5, 0x33, 0xfc, 0xb5, 0x45, 0xeb, 0x0e, 0x42, 0x22, 0x47, 0x43,
	0x8a, 0xba, 0x43, 0x8a, 0xbb, 0x42, 0x3a, 0xe9, 0x08, 0xa9, 0xd7, 0x1d, 0x52, 0xff, 0x7e, 0x48,
	0xed, 0x76, 0x83, 0xbb, 0xdb, 0x2d, 0xc2, 0xf0, 0xea, 0xe8, 0x8f, 0xa7, 0xd7, 0xb6, 0xa5, 0x5f,
	0xe1, 0xb4, 0xd9, 0xde, 0x1a, 0xad, 0x2c, 0xd2, 0xf7, 0xd5, 0xf5, 0x0a, 0xac, 0x65, 0x64, 0x1a,
	0xcf, 0x92, 0xe5, 0xf8, 0x50, 0xa5, 0xb9, 0x62, 0xd9, 0x6d, 0x73, 0x3a, 0x01, 0x58, 0xa1, 0x6b,
	0x62, 0xdc, 0xbb, 0x7b, 0xe9, 0x17, 0x48, 0x3c, 0x5b, 0x8f, 0x79, 0x77, 0xc7, 0x2b, 0xf1, 0x5e,
	0x8f, 0x4d, 0x69, 0x7b, 0x97, 0x7f, 0xda, 0xf3, 0xfa, 0x28, 0xb8, 0x71, 0x58, 0xd2, 0xef, 0xd0,
	0x0f, 0x05, 0xfa, 0xea, 0x50, 0xe1, 0xde, 0xd1, 0x8e, 0xa7, 0xdd, 0x0d, 0xc1, 0x56, 0xfa, 0x88,
	0x7e, 0xae, 0x5e, 0xa0, 0xa3, 0x93, 0x87, 0x92, 0x6b, 0x96, 0x1b, 0xbf, 0xec, 0x60, 0x1b, 0x95,
	0x8b, 0xbe, 0x7f, 0xe0, 0x6f, 0xff, 0x07, 0x00, 0x00, 0xff, 0xff, 0x0d, 0x6a, 0x52, 0x3a, 0xff,
	0x03, 0x00, 0x00,
}
