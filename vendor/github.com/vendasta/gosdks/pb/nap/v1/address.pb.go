// Code generated by protoc-gen-go.
// source: address.proto
// DO NOT EDIT!

/*
Package nap_v1 is a generated protocol buffer package.

It is generated from these files:
	address.proto
	api.proto
	phonenumber.proto

It has these top-level messages:
	Country
	State
	ListCountriesResponse
	ListStatesResponse
	ListCountriesRequest
	ListStatesRequest
	ParseResult
	ValidationResult
	FormatResult
	MetaDataResult
	ParsePhoneNumberRequest
	ParsePhoneNumberResponse
*/
package nap_v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Country struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// ISO 3166-1 alpha-2 codes (eg. "US" and "CA")
	Id string `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
}

func (m *Country) Reset()                    { *m = Country{} }
func (m *Country) String() string            { return proto.CompactTextString(m) }
func (*Country) ProtoMessage()               {}
func (*Country) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Country) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Country) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type State struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// ISO 3166-2 code without the country code prefix (eg. "AB", not "CA-AB")
	Id string `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
}

func (m *State) Reset()                    { *m = State{} }
func (m *State) String() string            { return proto.CompactTextString(m) }
func (*State) ProtoMessage()               {}
func (*State) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *State) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *State) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ListCountriesResponse struct {
	Countries []*Country `protobuf:"bytes,1,rep,name=countries" json:"countries,omitempty"`
}

func (m *ListCountriesResponse) Reset()                    { *m = ListCountriesResponse{} }
func (m *ListCountriesResponse) String() string            { return proto.CompactTextString(m) }
func (*ListCountriesResponse) ProtoMessage()               {}
func (*ListCountriesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ListCountriesResponse) GetCountries() []*Country {
	if m != nil {
		return m.Countries
	}
	return nil
}

type ListStatesResponse struct {
	States []*State `protobuf:"bytes,1,rep,name=states" json:"states,omitempty"`
}

func (m *ListStatesResponse) Reset()                    { *m = ListStatesResponse{} }
func (m *ListStatesResponse) String() string            { return proto.CompactTextString(m) }
func (*ListStatesResponse) ProtoMessage()               {}
func (*ListStatesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ListStatesResponse) GetStates() []*State {
	if m != nil {
		return m.States
	}
	return nil
}

type ListCountriesRequest struct {
}

func (m *ListCountriesRequest) Reset()                    { *m = ListCountriesRequest{} }
func (m *ListCountriesRequest) String() string            { return proto.CompactTextString(m) }
func (*ListCountriesRequest) ProtoMessage()               {}
func (*ListCountriesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type ListStatesRequest struct {
	// ISO 3166-1 alpha-2 code (eg. "US" and "CA")
	CountryId string `protobuf:"bytes,1,opt,name=country_id,json=countryId" json:"country_id,omitempty"`
}

func (m *ListStatesRequest) Reset()                    { *m = ListStatesRequest{} }
func (m *ListStatesRequest) String() string            { return proto.CompactTextString(m) }
func (*ListStatesRequest) ProtoMessage()               {}
func (*ListStatesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ListStatesRequest) GetCountryId() string {
	if m != nil {
		return m.CountryId
	}
	return ""
}

func init() {
	proto.RegisterType((*Country)(nil), "nap.v1.Country")
	proto.RegisterType((*State)(nil), "nap.v1.State")
	proto.RegisterType((*ListCountriesResponse)(nil), "nap.v1.ListCountriesResponse")
	proto.RegisterType((*ListStatesResponse)(nil), "nap.v1.ListStatesResponse")
	proto.RegisterType((*ListCountriesRequest)(nil), "nap.v1.ListCountriesRequest")
	proto.RegisterType((*ListStatesRequest)(nil), "nap.v1.ListStatesRequest")
}

func init() { proto.RegisterFile("address.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 212 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0x4c, 0x49, 0x29,
	0x4a, 0x2d, 0x2e, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xcb, 0x4b, 0x2c, 0xd0, 0x2b,
	0x33, 0x54, 0xd2, 0xe5, 0x62, 0x77, 0xce, 0x2f, 0xcd, 0x2b, 0x29, 0xaa, 0x14, 0x12, 0xe2, 0x62,
	0xc9, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x85, 0xf8, 0xb8,
	0x98, 0x32, 0x53, 0x24, 0x98, 0xc0, 0x22, 0x4c, 0x99, 0x29, 0x4a, 0xda, 0x5c, 0xac, 0xc1, 0x25,
	0x89, 0x25, 0xa9, 0x44, 0x29, 0x76, 0xe3, 0x12, 0xf5, 0xc9, 0x2c, 0x2e, 0x81, 0x98, 0x9f, 0x99,
	0x5a, 0x1c, 0x94, 0x5a, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x2a, 0xa4, 0xcb, 0xc5, 0x99, 0x0c, 0x13,
	0x94, 0x60, 0x54, 0x60, 0xd6, 0xe0, 0x36, 0xe2, 0xd7, 0x83, 0x38, 0x48, 0x0f, 0xea, 0x9a, 0x20,
	0x84, 0x0a, 0x25, 0x6b, 0x2e, 0x21, 0x90, 0x39, 0x60, 0x8b, 0x11, 0x86, 0xa8, 0x72, 0xb1, 0x15,
	0x83, 0x45, 0xa0, 0x26, 0xf0, 0xc2, 0x4c, 0x00, 0xab, 0x0b, 0x82, 0x4a, 0x2a, 0x89, 0x71, 0x89,
	0xa0, 0x39, 0xa2, 0xb0, 0x34, 0xb5, 0xb8, 0x44, 0xc9, 0x88, 0x4b, 0x10, 0xd9, 0x50, 0xb0, 0xa0,
	0x90, 0x2c, 0x17, 0x17, 0xc4, 0xda, 0xca, 0xf8, 0xcc, 0x14, 0xa8, 0xdf, 0xa0, 0x0e, 0xa9, 0xf4,
	0x4c, 0x49, 0x62, 0x03, 0x87, 0x9d, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x2c, 0xd3, 0x17, 0x43,
	0x4c, 0x01, 0x00, 0x00,
}