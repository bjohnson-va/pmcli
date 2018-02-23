// Code generated by protoc-gen-go. DO NOT EDIT.
// source: datalakeproto/listing.proto

package datalakeproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf1 "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Geo struct {
	Latitude  float64 `protobuf:"fixed64,1,opt,name=latitude" json:"latitude,omitempty"`
	Longitude float64 `protobuf:"fixed64,2,opt,name=longitude" json:"longitude,omitempty"`
}

func (m *Geo) Reset()                    { *m = Geo{} }
func (m *Geo) String() string            { return proto.CompactTextString(m) }
func (*Geo) ProtoMessage()               {}
func (*Geo) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

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

type RawListing struct {
	RawListingId string `protobuf:"bytes,1,opt,name=raw_listing_id,json=rawListingId" json:"raw_listing_id,omitempty"`
	ExternalId   string `protobuf:"bytes,2,opt,name=external_id,json=externalId" json:"external_id,omitempty"`
	Url          string `protobuf:"bytes,3,opt,name=url" json:"url,omitempty"`
	// Basic NAP data
	CompanyName            string   `protobuf:"bytes,4,opt,name=company_name,json=companyName" json:"company_name,omitempty"`
	Address                string   `protobuf:"bytes,5,opt,name=address" json:"address,omitempty"`
	City                   string   `protobuf:"bytes,6,opt,name=city" json:"city,omitempty"`
	State                  string   `protobuf:"bytes,7,opt,name=state" json:"state,omitempty"`
	Country                string   `protobuf:"bytes,8,opt,name=country" json:"country,omitempty"`
	ZipCode                string   `protobuf:"bytes,9,opt,name=zip_code,json=zipCode" json:"zip_code,omitempty"`
	GeoLocation            *Geo     `protobuf:"bytes,10,opt,name=geo_location,json=geoLocation" json:"geo_location,omitempty"`
	Phone                  string   `protobuf:"bytes,11,opt,name=phone" json:"phone,omitempty"`
	AdditionalPhoneNumbers []string `protobuf:"bytes,12,rep,name=additional_phone_numbers,json=additionalPhoneNumbers" json:"additional_phone_numbers,omitempty"`
	Website                string   `protobuf:"bytes,13,opt,name=website" json:"website,omitempty"`
	// Extended NAP data
	NumberOfReviews     int64                       `protobuf:"varint,14,opt,name=number_of_reviews,json=numberOfReviews" json:"number_of_reviews,omitempty"`
	AverageReviewRating float64                     `protobuf:"fixed64,15,opt,name=average_review_rating,json=averageReviewRating" json:"average_review_rating,omitempty"`
	BusinessCategories  []string                    `protobuf:"bytes,16,rep,name=business_categories,json=businessCategories" json:"business_categories,omitempty"`
	SourceId            int64                       `protobuf:"varint,17,opt,name=source_id,json=sourceId" json:"source_id,omitempty"`
	Modified            *google_protobuf1.Timestamp `protobuf:"bytes,18,opt,name=modified" json:"modified,omitempty"`
	DeletedOn           *google_protobuf1.Timestamp `protobuf:"bytes,19,opt,name=deleted_on,json=deletedOn" json:"deleted_on,omitempty"`
	ExpiresOn           *google_protobuf1.Timestamp `protobuf:"bytes,20,opt,name=expires_on,json=expiresOn" json:"expires_on,omitempty"`
	NotifyAfter         *google_protobuf1.Timestamp `protobuf:"bytes,21,opt,name=notify_after,json=notifyAfter" json:"notify_after,omitempty"`
	ClickTrackedWebsite string                      `protobuf:"bytes,22,opt,name=click_tracked_website,json=clickTrackedWebsite" json:"click_tracked_website,omitempty"`
}

func (m *RawListing) Reset()                    { *m = RawListing{} }
func (m *RawListing) String() string            { return proto.CompactTextString(m) }
func (*RawListing) ProtoMessage()               {}
func (*RawListing) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *RawListing) GetRawListingId() string {
	if m != nil {
		return m.RawListingId
	}
	return ""
}

func (m *RawListing) GetExternalId() string {
	if m != nil {
		return m.ExternalId
	}
	return ""
}

func (m *RawListing) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *RawListing) GetCompanyName() string {
	if m != nil {
		return m.CompanyName
	}
	return ""
}

func (m *RawListing) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *RawListing) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *RawListing) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *RawListing) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *RawListing) GetZipCode() string {
	if m != nil {
		return m.ZipCode
	}
	return ""
}

func (m *RawListing) GetGeoLocation() *Geo {
	if m != nil {
		return m.GeoLocation
	}
	return nil
}

func (m *RawListing) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *RawListing) GetAdditionalPhoneNumbers() []string {
	if m != nil {
		return m.AdditionalPhoneNumbers
	}
	return nil
}

func (m *RawListing) GetWebsite() string {
	if m != nil {
		return m.Website
	}
	return ""
}

func (m *RawListing) GetNumberOfReviews() int64 {
	if m != nil {
		return m.NumberOfReviews
	}
	return 0
}

func (m *RawListing) GetAverageReviewRating() float64 {
	if m != nil {
		return m.AverageReviewRating
	}
	return 0
}

func (m *RawListing) GetBusinessCategories() []string {
	if m != nil {
		return m.BusinessCategories
	}
	return nil
}

func (m *RawListing) GetSourceId() int64 {
	if m != nil {
		return m.SourceId
	}
	return 0
}

func (m *RawListing) GetModified() *google_protobuf1.Timestamp {
	if m != nil {
		return m.Modified
	}
	return nil
}

func (m *RawListing) GetDeletedOn() *google_protobuf1.Timestamp {
	if m != nil {
		return m.DeletedOn
	}
	return nil
}

func (m *RawListing) GetExpiresOn() *google_protobuf1.Timestamp {
	if m != nil {
		return m.ExpiresOn
	}
	return nil
}

func (m *RawListing) GetNotifyAfter() *google_protobuf1.Timestamp {
	if m != nil {
		return m.NotifyAfter
	}
	return nil
}

func (m *RawListing) GetClickTrackedWebsite() string {
	if m != nil {
		return m.ClickTrackedWebsite
	}
	return ""
}

func init() {
	proto.RegisterType((*Geo)(nil), "datalakeproto.Geo")
	proto.RegisterType((*RawListing)(nil), "datalakeproto.RawListing")
}

func init() { proto.RegisterFile("datalakeproto/listing.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 563 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x92, 0x5f, 0x6b, 0xdb, 0x30,
	0x14, 0xc5, 0x49, 0xd3, 0x3f, 0xc9, 0x4d, 0xfa, 0x4f, 0x69, 0x8b, 0xd6, 0x0e, 0x9a, 0x95, 0x3d,
	0x84, 0x3d, 0x24, 0xd0, 0xb1, 0xb1, 0x3d, 0x8c, 0x31, 0xfa, 0x50, 0x0a, 0xa5, 0x1d, 0xa6, 0xb0,
	0x47, 0xa1, 0x58, 0xd7, 0x9e, 0xa8, 0x2d, 0x19, 0x49, 0x6e, 0x9a, 0x7e, 0xeb, 0x7d, 0x83, 0x21,
	0xc9, 0x4e, 0xe9, 0x53, 0xde, 0x7c, 0xce, 0xef, 0xdc, 0xcb, 0xc1, 0x57, 0x70, 0x26, 0xb8, 0xe3,
	0x05, 0x7f, 0xc4, 0xca, 0x68, 0xa7, 0x67, 0x85, 0xb4, 0x4e, 0xaa, 0x7c, 0x1a, 0x14, 0xd9, 0x7d,
	0x03, 0x4f, 0xcf, 0x73, 0xad, 0xf3, 0x02, 0x67, 0x41, 0xcd, 0xeb, 0x6c, 0xe6, 0x64, 0x89, 0xd6,
	0xf1, 0xb2, 0x8a, 0xf9, 0x8b, 0x9f, 0xd0, 0xbd, 0x46, 0x4d, 0x4e, 0xa1, 0x57, 0x70, 0x27, 0x5d,
	0x2d, 0x90, 0x76, 0xc6, 0x9d, 0x49, 0x27, 0x59, 0x69, 0xf2, 0x1e, 0xfa, 0x85, 0x56, 0x79, 0x84,
	0x1b, 0x01, 0xbe, 0x1a, 0x17, 0xff, 0xb6, 0x01, 0x12, 0xbe, 0xb8, 0x8d, 0x2d, 0xc8, 0x47, 0xd8,
	0x33, 0x7c, 0xc1, 0x9a, 0x52, 0x4c, 0x8a, 0xb0, 0xae, 0x9f, 0x0c, 0xcd, 0x2a, 0x73, 0x23, 0xc8,
	0x39, 0x0c, 0xf0, 0xd9, 0xa1, 0x51, 0xbc, 0xf0, 0x91, 0x8d, 0x10, 0x81, 0xd6, 0xba, 0x11, 0xe4,
	0x00, 0xba, 0xb5, 0x29, 0x68, 0x37, 0x00, 0xff, 0x49, 0x3e, 0xc0, 0x30, 0xd5, 0x65, 0xc5, 0xd5,
	0x92, 0x29, 0x5e, 0x22, 0xdd, 0x0c, 0x68, 0xd0, 0x78, 0x77, 0xbc, 0x44, 0x42, 0x61, 0x87, 0x0b,
	0x61, 0xd0, 0x5a, 0xba, 0x15, 0x68, 0x2b, 0x09, 0x81, 0xcd, 0x54, 0xba, 0x25, 0xdd, 0x0e, 0x76,
	0xf8, 0x26, 0x47, 0xb0, 0x65, 0x1d, 0x77, 0x48, 0x77, 0x82, 0x19, 0x85, 0xdf, 0x91, 0xea, 0x5a,
	0x39, 0xb3, 0xa4, 0xbd, 0xb8, 0xa3, 0x91, 0xe4, 0x1d, 0xf4, 0x5e, 0x64, 0xc5, 0x52, 0x2d, 0x90,
	0xf6, 0x23, 0x7a, 0x91, 0xd5, 0x95, 0x16, 0x48, 0xbe, 0xc0, 0x30, 0x47, 0xcd, 0x0a, 0x9d, 0x72,
	0x27, 0xb5, 0xa2, 0x30, 0xee, 0x4c, 0x06, 0x97, 0x64, 0xfa, 0xe6, 0x16, 0xd3, 0x6b, 0xd4, 0xc9,
	0x20, 0x47, 0x7d, 0xdb, 0xc4, 0x7c, 0x83, 0xea, 0xaf, 0x56, 0x48, 0x07, 0xb1, 0x41, 0x10, 0xe4,
	0x1b, 0x50, 0x2e, 0x84, 0xf4, 0x09, 0x5e, 0xb0, 0xe0, 0x31, 0x55, 0x97, 0x73, 0x34, 0x96, 0x0e,
	0xc7, 0xdd, 0x49, 0x3f, 0x39, 0x79, 0xe5, 0xbf, 0x3d, 0xbe, 0x8b, 0xd4, 0x77, 0x5f, 0xe0, 0xdc,
	0x4a, 0x87, 0x74, 0x37, 0x16, 0x6c, 0x24, 0xf9, 0x04, 0x87, 0x71, 0x05, 0xd3, 0x19, 0x33, 0xf8,
	0x24, 0x71, 0x61, 0xe9, 0xde, 0xb8, 0x33, 0xe9, 0x26, 0xfb, 0x11, 0xdc, 0x67, 0x49, 0xb4, 0xc9,
	0x25, 0x1c, 0xf3, 0x27, 0x34, 0x3c, 0xc7, 0x26, 0xc9, 0x0c, 0xf7, 0x67, 0xa3, 0xfb, 0xe1, 0xf4,
	0xa3, 0x06, 0xc6, 0x78, 0x12, 0x10, 0x99, 0xc1, 0x68, 0x5e, 0x5b, 0xa9, 0xd0, 0x5a, 0x96, 0x72,
	0x87, 0xb9, 0x36, 0x12, 0x2d, 0x3d, 0x08, 0x75, 0x49, 0x8b, 0xae, 0x56, 0x84, 0x9c, 0x41, 0xdf,
	0xea, 0xda, 0xa4, 0xe8, 0xcf, 0x7f, 0x18, 0x8a, 0xf4, 0xa2, 0x71, 0x23, 0xc8, 0x57, 0xe8, 0x95,
	0x5a, 0xc8, 0x4c, 0xa2, 0xa0, 0x24, 0xfc, 0xca, 0xd3, 0x69, 0x7c, 0xc7, 0xd3, 0xf6, 0x1d, 0x4f,
	0x1f, 0xda, 0x77, 0x9c, 0xac, 0xb2, 0xe4, 0x3b, 0x80, 0xc0, 0x02, 0x1d, 0x0a, 0xa6, 0x15, 0x1d,
	0xad, 0x9d, 0xec, 0x37, 0xe9, 0x7b, 0xe5, 0x47, 0xf1, 0xb9, 0x92, 0x06, 0xad, 0x1f, 0x3d, 0x5a,
	0x3f, 0xda, 0xa4, 0xef, 0x15, 0xf9, 0x01, 0x43, 0xa5, 0x9d, 0xcc, 0x96, 0x8c, 0x67, 0x0e, 0x0d,
	0x3d, 0x5e, 0x3b, 0x3c, 0x88, 0xf9, 0x5f, 0x3e, 0xee, 0x7f, 0x77, 0x5a, 0xc8, 0xf4, 0x91, 0x39,
	0xc3, 0xd3, 0x47, 0x14, 0xac, 0x3d, 0xe1, 0x49, 0x38, 0xe1, 0x28, 0xc0, 0x87, 0xc8, 0xfe, 0x44,
	0x34, 0xdf, 0x0e, 0x4b, 0x3f, 0xff, 0x0f, 0x00, 0x00, 0xff, 0xff, 0x39, 0x71, 0xed, 0x64, 0x0a,
	0x04, 0x00, 0x00,
}