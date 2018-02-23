// Code generated by protoc-gen-go. DO NOT EDIT.
// source: marketplace_packages/v1/package.proto

package marketplace_packages_v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Statuses_Status int32

const (
	Statuses_DRAFT     Statuses_Status = 0
	Statuses_PUBLISHED Statuses_Status = 1
	Statuses_ARCHIVED  Statuses_Status = 2
)

var Statuses_Status_name = map[int32]string{
	0: "DRAFT",
	1: "PUBLISHED",
	2: "ARCHIVED",
}
var Statuses_Status_value = map[string]int32{
	"DRAFT":     0,
	"PUBLISHED": 1,
	"ARCHIVED":  2,
}

func (x Statuses_Status) String() string {
	return proto.EnumName(Statuses_Status_name, int32(x))
}
func (Statuses_Status) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{0, 0} }

type Currencies_Currency int32

const (
	Currencies_USD Currencies_Currency = 0
	Currencies_AUD Currencies_Currency = 1
	Currencies_BRL Currencies_Currency = 2
	Currencies_CAD Currencies_Currency = 3
	Currencies_CHF Currencies_Currency = 4
	Currencies_CNY Currencies_Currency = 5
	Currencies_CZK Currencies_Currency = 6
	Currencies_EUR Currencies_Currency = 7
	Currencies_GBP Currencies_Currency = 8
	Currencies_HKD Currencies_Currency = 9
	Currencies_INR Currencies_Currency = 10
	Currencies_JPY Currencies_Currency = 11
	Currencies_KHR Currencies_Currency = 12
	Currencies_KRW Currencies_Currency = 13
	Currencies_MXN Currencies_Currency = 14
	Currencies_NOK Currencies_Currency = 15
	Currencies_NZD Currencies_Currency = 16
	Currencies_RUB Currencies_Currency = 17
	Currencies_SEK Currencies_Currency = 18
	Currencies_SGD Currencies_Currency = 19
	Currencies_TRY Currencies_Currency = 20
	Currencies_ZAR Currencies_Currency = 21
)

var Currencies_Currency_name = map[int32]string{
	0:  "USD",
	1:  "AUD",
	2:  "BRL",
	3:  "CAD",
	4:  "CHF",
	5:  "CNY",
	6:  "CZK",
	7:  "EUR",
	8:  "GBP",
	9:  "HKD",
	10: "INR",
	11: "JPY",
	12: "KHR",
	13: "KRW",
	14: "MXN",
	15: "NOK",
	16: "NZD",
	17: "RUB",
	18: "SEK",
	19: "SGD",
	20: "TRY",
	21: "ZAR",
}
var Currencies_Currency_value = map[string]int32{
	"USD": 0,
	"AUD": 1,
	"BRL": 2,
	"CAD": 3,
	"CHF": 4,
	"CNY": 5,
	"CZK": 6,
	"EUR": 7,
	"GBP": 8,
	"HKD": 9,
	"INR": 10,
	"JPY": 11,
	"KHR": 12,
	"KRW": 13,
	"MXN": 14,
	"NOK": 15,
	"NZD": 16,
	"RUB": 17,
	"SEK": 18,
	"SGD": 19,
	"TRY": 20,
	"ZAR": 21,
}

func (x Currencies_Currency) String() string {
	return proto.EnumName(Currencies_Currency_name, int32(x))
}
func (Currencies_Currency) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{1, 0} }

type Frequencies_Frequency int32

const (
	Frequencies_MONTHLY Frequencies_Frequency = 0
	Frequencies_DAILY   Frequencies_Frequency = 1
	Frequencies_ONCE    Frequencies_Frequency = 2
	Frequencies_YEARLY  Frequencies_Frequency = 3
	Frequencies_WEEKLY  Frequencies_Frequency = 4
	Frequencies_OTHER   Frequencies_Frequency = 5
)

var Frequencies_Frequency_name = map[int32]string{
	0: "MONTHLY",
	1: "DAILY",
	2: "ONCE",
	3: "YEARLY",
	4: "WEEKLY",
	5: "OTHER",
}
var Frequencies_Frequency_value = map[string]int32{
	"MONTHLY": 0,
	"DAILY":   1,
	"ONCE":    2,
	"YEARLY":  3,
	"WEEKLY":  4,
	"OTHER":   5,
}

func (x Frequencies_Frequency) String() string {
	return proto.EnumName(Frequencies_Frequency_name, int32(x))
}
func (Frequencies_Frequency) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{2, 0} }

type Statuses struct {
}

func (m *Statuses) Reset()                    { *m = Statuses{} }
func (m *Statuses) String() string            { return proto.CompactTextString(m) }
func (*Statuses) ProtoMessage()               {}
func (*Statuses) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

type Currencies struct {
}

func (m *Currencies) Reset()                    { *m = Currencies{} }
func (m *Currencies) String() string            { return proto.CompactTextString(m) }
func (*Currencies) ProtoMessage()               {}
func (*Currencies) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

type Frequencies struct {
}

func (m *Frequencies) Reset()                    { *m = Frequencies{} }
func (m *Frequencies) String() string            { return proto.CompactTextString(m) }
func (*Frequencies) ProtoMessage()               {}
func (*Frequencies) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

type AddonKey struct {
	// The parent app for the addon
	AppId string `protobuf:"bytes,1,opt,name=app_id,json=appId" json:"app_id,omitempty"`
	// The addon id
	AddonId string `protobuf:"bytes,2,opt,name=addon_id,json=addonId" json:"addon_id,omitempty"`
}

func (m *AddonKey) Reset()                    { *m = AddonKey{} }
func (m *AddonKey) String() string            { return proto.CompactTextString(m) }
func (*AddonKey) ProtoMessage()               {}
func (*AddonKey) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

func (m *AddonKey) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *AddonKey) GetAddonId() string {
	if m != nil {
		return m.AddonId
	}
	return ""
}

type Price struct {
	// Price of the product
	Price int32 `protobuf:"varint,1,opt,name=price" json:"price,omitempty"`
	// Billing frequency of the product
	Frequency Frequencies_Frequency `protobuf:"varint,2,opt,name=frequency,enum=marketplace_packages.v1.Frequencies_Frequency" json:"frequency,omitempty"`
}

func (m *Price) Reset()                    { *m = Price{} }
func (m *Price) String() string            { return proto.CompactTextString(m) }
func (*Price) ProtoMessage()               {}
func (*Price) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

func (m *Price) GetPrice() int32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *Price) GetFrequency() Frequencies_Frequency {
	if m != nil {
		return m.Frequency
	}
	return Frequencies_MONTHLY
}

type Pricing struct {
	// Currency of the pricing
	Currency Currencies_Currency `protobuf:"varint,1,opt,name=currency,enum=marketplace_packages.v1.Currencies_Currency" json:"currency,omitempty"`
	// Prices of the package
	Prices []*Price `protobuf:"bytes,2,rep,name=prices" json:"prices,omitempty"`
}

func (m *Pricing) Reset()                    { *m = Pricing{} }
func (m *Pricing) String() string            { return proto.CompactTextString(m) }
func (*Pricing) ProtoMessage()               {}
func (*Pricing) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{5} }

func (m *Pricing) GetCurrency() Currencies_Currency {
	if m != nil {
		return m.Currency
	}
	return Currencies_USD
}

func (m *Pricing) GetPrices() []*Price {
	if m != nil {
		return m.Prices
	}
	return nil
}

type Package struct {
	// Created time
	Created *google_protobuf.Timestamp `protobuf:"bytes,1,opt,name=created" json:"created,omitempty"`
	// Updated time
	Updated *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=updated" json:"updated,omitempty"`
	// Archived time
	Archived *google_protobuf.Timestamp `protobuf:"bytes,3,opt,name=archived" json:"archived,omitempty"`
	// User that last updated the package
	UpdatedBy string `protobuf:"bytes,4,opt,name=updated_by,json=updatedBy" json:"updated_by,omitempty"`
	// Unique id for the package
	PackageId string `protobuf:"bytes,5,opt,name=package_id,json=packageId" json:"package_id,omitempty"`
	// Partner Id for package
	PartnerId string `protobuf:"bytes,6,opt,name=partner_id,json=partnerId" json:"partner_id,omitempty"`
	// Market Id for the package
	MarketId string `protobuf:"bytes,7,opt,name=market_id,json=marketId" json:"market_id,omitempty"`
	// Package name
	Name string `protobuf:"bytes,8,opt,name=name" json:"name,omitempty"`
	// Icon for the package
	Icon string `protobuf:"bytes,9,opt,name=icon" json:"icon,omitempty"`
	// Status for the package
	Status Statuses_Status `protobuf:"varint,10,opt,name=status,enum=marketplace_packages.v1.Statuses_Status" json:"status,omitempty"`
	// URL for the header image
	HeaderImageUrl string `protobuf:"bytes,11,opt,name=header_image_url,json=headerImageUrl" json:"header_image_url,omitempty"`
	// The slogan for the package
	Tagline string `protobuf:"bytes,12,opt,name=tagline" json:"tagline,omitempty"`
	// What is in the package
	Content string `protobuf:"bytes,13,opt,name=content" json:"content,omitempty"`
	// Product ids for a package
	Products []string `protobuf:"bytes,14,rep,name=products" json:"products,omitempty"`
	// Hide details of the products in the package
	HideProductDetails bool `protobuf:"varint,15,opt,name=hide_product_details,json=hideProductDetails" json:"hide_product_details,omitempty"`
	// Hide product icons and names
	HideProductIconsAndNames bool `protobuf:"varint,16,opt,name=hide_product_icons_and_names,json=hideProductIconsAndNames" json:"hide_product_icons_and_names,omitempty"`
	// Pricing of the package
	Pricing *Pricing `protobuf:"bytes,17,opt,name=pricing" json:"pricing,omitempty"`
	// Price for the year
	NormalizedAnnualPrice int64 `protobuf:"varint,18,opt,name=normalized_annual_price,json=normalizedAnnualPrice" json:"normalized_annual_price,omitempty"`
	// Addon ids for a package
	AddonKeys []*AddonKey `protobuf:"bytes,19,rep,name=addon_keys,json=addonKeys" json:"addon_keys,omitempty"`
}

func (m *Package) Reset()                    { *m = Package{} }
func (m *Package) String() string            { return proto.CompactTextString(m) }
func (*Package) ProtoMessage()               {}
func (*Package) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{6} }

func (m *Package) GetCreated() *google_protobuf.Timestamp {
	if m != nil {
		return m.Created
	}
	return nil
}

func (m *Package) GetUpdated() *google_protobuf.Timestamp {
	if m != nil {
		return m.Updated
	}
	return nil
}

func (m *Package) GetArchived() *google_protobuf.Timestamp {
	if m != nil {
		return m.Archived
	}
	return nil
}

func (m *Package) GetUpdatedBy() string {
	if m != nil {
		return m.UpdatedBy
	}
	return ""
}

func (m *Package) GetPackageId() string {
	if m != nil {
		return m.PackageId
	}
	return ""
}

func (m *Package) GetPartnerId() string {
	if m != nil {
		return m.PartnerId
	}
	return ""
}

func (m *Package) GetMarketId() string {
	if m != nil {
		return m.MarketId
	}
	return ""
}

func (m *Package) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Package) GetIcon() string {
	if m != nil {
		return m.Icon
	}
	return ""
}

func (m *Package) GetStatus() Statuses_Status {
	if m != nil {
		return m.Status
	}
	return Statuses_DRAFT
}

func (m *Package) GetHeaderImageUrl() string {
	if m != nil {
		return m.HeaderImageUrl
	}
	return ""
}

func (m *Package) GetTagline() string {
	if m != nil {
		return m.Tagline
	}
	return ""
}

func (m *Package) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *Package) GetProducts() []string {
	if m != nil {
		return m.Products
	}
	return nil
}

func (m *Package) GetHideProductDetails() bool {
	if m != nil {
		return m.HideProductDetails
	}
	return false
}

func (m *Package) GetHideProductIconsAndNames() bool {
	if m != nil {
		return m.HideProductIconsAndNames
	}
	return false
}

func (m *Package) GetPricing() *Pricing {
	if m != nil {
		return m.Pricing
	}
	return nil
}

func (m *Package) GetNormalizedAnnualPrice() int64 {
	if m != nil {
		return m.NormalizedAnnualPrice
	}
	return 0
}

func (m *Package) GetAddonKeys() []*AddonKey {
	if m != nil {
		return m.AddonKeys
	}
	return nil
}

func init() {
	proto.RegisterType((*Statuses)(nil), "marketplace_packages.v1.Statuses")
	proto.RegisterType((*Currencies)(nil), "marketplace_packages.v1.Currencies")
	proto.RegisterType((*Frequencies)(nil), "marketplace_packages.v1.Frequencies")
	proto.RegisterType((*AddonKey)(nil), "marketplace_packages.v1.AddonKey")
	proto.RegisterType((*Price)(nil), "marketplace_packages.v1.Price")
	proto.RegisterType((*Pricing)(nil), "marketplace_packages.v1.Pricing")
	proto.RegisterType((*Package)(nil), "marketplace_packages.v1.Package")
	proto.RegisterEnum("marketplace_packages.v1.Statuses_Status", Statuses_Status_name, Statuses_Status_value)
	proto.RegisterEnum("marketplace_packages.v1.Currencies_Currency", Currencies_Currency_name, Currencies_Currency_value)
	proto.RegisterEnum("marketplace_packages.v1.Frequencies_Frequency", Frequencies_Frequency_name, Frequencies_Frequency_value)
}

func init() { proto.RegisterFile("marketplace_packages/v1/package.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 857 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x94, 0xdd, 0x6e, 0xdb, 0x36,
	0x14, 0xc7, 0xe3, 0x4f, 0x49, 0xc7, 0x89, 0xcb, 0xb2, 0x09, 0xaa, 0x65, 0x5f, 0x9e, 0x81, 0x01,
	0xbe, 0x18, 0x9c, 0x36, 0x1b, 0x72, 0x31, 0x0c, 0x43, 0xe5, 0x48, 0xa9, 0x35, 0xbb, 0x8e, 0xc1,
	0xd8, 0xeb, 0x9c, 0x1b, 0x81, 0x91, 0x58, 0x47, 0x88, 0x2d, 0x69, 0x92, 0x1c, 0xc0, 0x7b, 0x85,
	0x3d, 0xce, 0x5e, 0x64, 0x6f, 0xb0, 0x57, 0x19, 0x0e, 0x29, 0x25, 0x1d, 0x30, 0x37, 0x77, 0xbf,
	0x73, 0xfe, 0xe7, 0x4f, 0x8a, 0x3c, 0x87, 0x82, 0x6f, 0xd7, 0x3c, 0xbd, 0x13, 0x79, 0xb2, 0xe2,
	0xbe, 0xf0, 0x12, 0xee, 0xdf, 0xf1, 0xa5, 0xc8, 0x4e, 0xee, 0x5f, 0x9f, 0x14, 0xdc, 0x4f, 0xd2,
	0x38, 0x8f, 0xe9, 0xcb, 0xff, 0x2b, 0xeb, 0xdf, 0xbf, 0x3e, 0xfe, 0x7a, 0x19, 0xc7, 0xcb, 0x95,
	0x38, 0x91, 0x65, 0x37, 0x9b, 0x0f, 0x27, 0x79, 0xb8, 0x16, 0x59, 0xce, 0xd7, 0x89, 0x72, 0x76,
	0x7f, 0x02, 0xfd, 0x2a, 0xe7, 0xf9, 0x26, 0x13, 0x59, 0xf7, 0x15, 0x34, 0x15, 0x53, 0x03, 0x1a,
	0x36, 0xb3, 0x2e, 0x66, 0x64, 0x8f, 0x1e, 0x80, 0x31, 0x9d, 0x0f, 0xc6, 0xee, 0xd5, 0xd0, 0xb1,
	0x49, 0x85, 0xee, 0x83, 0x6e, 0xb1, 0xf3, 0xa1, 0xfb, 0xab, 0x63, 0x93, 0x6a, 0xf7, 0x9f, 0x0a,
	0xc0, 0xf9, 0x26, 0x4d, 0x45, 0xe4, 0x87, 0x22, 0xeb, 0xfe, 0x5d, 0x01, 0xbd, 0x08, 0xb7, 0x54,
	0x83, 0xda, 0xfc, 0xca, 0x26, 0x7b, 0x08, 0xd6, 0x1c, 0xbd, 0x1a, 0xd4, 0x06, 0x6c, 0x4c, 0xaa,
	0x08, 0xe7, 0x96, 0x4d, 0x6a, 0x12, 0x86, 0x17, 0xa4, 0x2e, 0x61, 0xb2, 0x20, 0x0d, 0x09, 0xd7,
	0x23, 0xd2, 0x44, 0x70, 0xe6, 0x8c, 0x68, 0x08, 0x6f, 0x07, 0x53, 0xa2, 0x23, 0x0c, 0x47, 0x36,
	0x31, 0x10, 0xdc, 0x09, 0x23, 0x80, 0xf0, 0xcb, 0x74, 0x41, 0x5a, 0x08, 0xa3, 0x21, 0x23, 0xfb,
	0x12, 0xd8, 0x7b, 0x72, 0x80, 0xf0, 0xee, 0xb7, 0x09, 0x69, 0x23, 0x4c, 0x2e, 0x47, 0xe4, 0x99,
	0x84, 0x6b, 0x9b, 0x10, 0x04, 0x36, 0x1f, 0x90, 0xe7, 0x08, 0x57, 0xce, 0x88, 0x50, 0x09, 0x6f,
	0x6d, 0xf2, 0x02, 0x61, 0xc6, 0x16, 0xe4, 0x10, 0xe1, 0xda, 0x62, 0xe4, 0xa8, 0xeb, 0x41, 0xeb,
	0x22, 0x15, 0xbf, 0x6f, 0x8a, 0x13, 0x4e, 0xc1, 0x28, 0xc3, 0x2d, 0x6d, 0x81, 0xf6, 0xee, 0x72,
	0x32, 0x1b, 0x8e, 0x17, 0x64, 0x4f, 0x5e, 0x99, 0xe5, 0x8e, 0x17, 0xa4, 0x42, 0x75, 0xa8, 0x5f,
	0x4e, 0xce, 0x1d, 0x52, 0xa5, 0x00, 0xcd, 0x85, 0x63, 0xb1, 0xf1, 0x82, 0xd4, 0x90, 0xdf, 0x3b,
	0xce, 0x68, 0xbc, 0x20, 0x75, 0x2c, 0xbe, 0x9c, 0x0d, 0x1d, 0x46, 0x1a, 0xd8, 0x00, 0x2b, 0x08,
	0xe2, 0x68, 0x24, 0xb6, 0xf4, 0x08, 0x9a, 0x3c, 0x49, 0xbc, 0x30, 0x30, 0x2b, 0x9d, 0x4a, 0xcf,
	0x60, 0x0d, 0x9e, 0x24, 0x6e, 0x40, 0x3f, 0x03, 0x9d, 0x63, 0x09, 0x0a, 0x55, 0x29, 0x68, 0x32,
	0x76, 0x83, 0xee, 0x1d, 0x34, 0xa6, 0x69, 0xe8, 0x0b, 0x7a, 0x08, 0x8d, 0x04, 0x41, 0x3a, 0x1b,
	0x4c, 0x05, 0x74, 0x0c, 0xc6, 0x87, 0xf2, 0x73, 0xa5, 0xb5, 0x7d, 0xda, 0xef, 0xef, 0x98, 0x95,
	0xfe, 0x47, 0xe7, 0x7c, 0xe0, 0x2d, 0x7b, 0x5c, 0xa0, 0xfb, 0x67, 0x05, 0x34, 0xdc, 0x2d, 0x8c,
	0x96, 0x74, 0x08, 0xba, 0x5f, 0x74, 0x5a, 0x6e, 0xd9, 0x3e, 0xfd, 0x6e, 0xe7, 0xc2, 0x8f, 0x13,
	0x52, 0xe2, 0x96, 0x3d, 0xb8, 0xe9, 0x19, 0x34, 0xe5, 0xc7, 0x66, 0x66, 0xb5, 0x53, 0xeb, 0xb5,
	0x4e, 0xbf, 0xda, 0xb9, 0x8e, 0x3c, 0x29, 0x2b, 0xaa, 0xbb, 0x7f, 0x35, 0x41, 0x9b, 0x2a, 0x95,
	0xfe, 0x00, 0x9a, 0x9f, 0x0a, 0x9e, 0x0b, 0x75, 0x73, 0xad, 0xd3, 0xe3, 0xbe, 0x1a, 0xfc, 0x7e,
	0x39, 0xf8, 0xfd, 0x59, 0x39, 0xf8, 0xac, 0x2c, 0x45, 0xd7, 0x26, 0x09, 0xa4, 0xab, 0xfa, 0xb4,
	0xab, 0x28, 0xa5, 0x67, 0xa0, 0xf3, 0xd4, 0xbf, 0x0d, 0xef, 0x45, 0x60, 0xd6, 0x9e, 0xb4, 0x3d,
	0xd4, 0xd2, 0x2f, 0x01, 0x8a, 0x25, 0xbc, 0x9b, 0xad, 0x59, 0x97, 0x7d, 0x34, 0x8a, 0xcc, 0x60,
	0x8b, 0x72, 0x71, 0x56, 0x6c, 0x73, 0x43, 0xc9, 0x45, 0xc6, 0x0d, 0x94, 0x9c, 0xe6, 0x91, 0x48,
	0x51, 0x6e, 0x96, 0xb2, 0xcc, 0xb8, 0x01, 0xfd, 0x1c, 0x0c, 0x75, 0x6b, 0xa8, 0x6a, 0x52, 0xd5,
	0x55, 0xc2, 0x0d, 0x28, 0x85, 0x7a, 0xc4, 0xd7, 0xc2, 0xd4, 0x65, 0x5e, 0x32, 0xe6, 0x42, 0x3f,
	0x8e, 0x4c, 0x43, 0xe5, 0x90, 0xe9, 0x1b, 0x68, 0x66, 0xf2, 0xfd, 0x9b, 0x20, 0x3b, 0xda, 0xdb,
	0xd9, 0x89, 0xf2, 0x97, 0x51, 0x00, 0x2b, 0x7c, 0xb4, 0x07, 0xe4, 0x56, 0xf0, 0x00, 0x3f, 0x72,
	0x8d, 0x27, 0xd9, 0xa4, 0x2b, 0xb3, 0x25, 0x77, 0x68, 0xab, 0xbc, 0x8b, 0xe9, 0x79, 0xba, 0xa2,
	0x26, 0x68, 0x39, 0x5f, 0xae, 0xc2, 0x48, 0x98, 0xfb, 0x6a, 0xa4, 0x8b, 0x10, 0x15, 0x3f, 0x8e,
	0x72, 0x11, 0xe5, 0xe6, 0x81, 0x52, 0x8a, 0x90, 0x1e, 0x83, 0x9e, 0xa4, 0x71, 0xb0, 0xf1, 0xf3,
	0xcc, 0x6c, 0x77, 0x6a, 0x78, 0xc6, 0x32, 0xa6, 0xaf, 0xe0, 0xf0, 0x36, 0x0c, 0x84, 0x57, 0x24,
	0xbc, 0x40, 0xe4, 0x3c, 0x5c, 0x65, 0xe6, 0xb3, 0x4e, 0xa5, 0xa7, 0x33, 0x8a, 0xda, 0x54, 0x49,
	0xb6, 0x52, 0xe8, 0xcf, 0xf0, 0xc5, 0x7f, 0x1c, 0x78, 0x05, 0x99, 0xc7, 0xa3, 0xc0, 0xc3, 0x0b,
	0xca, 0x4c, 0x22, 0x9d, 0xe6, 0x47, 0x4e, 0x17, 0x2b, 0xac, 0x28, 0x98, 0xa0, 0x4e, 0x7f, 0x04,
	0x2d, 0x51, 0x8f, 0xc1, 0x7c, 0x2e, 0xc7, 0xa0, 0xf3, 0xc9, 0xc1, 0x0d, 0xa3, 0x25, 0x2b, 0x0d,
	0xf4, 0x0c, 0x5e, 0x46, 0x71, 0xba, 0xe6, 0xab, 0xf0, 0x0f, 0x11, 0x78, 0x3c, 0x8a, 0x36, 0x7c,
	0xe5, 0xa9, 0xf7, 0x4b, 0x3b, 0x95, 0x5e, 0x8d, 0x1d, 0x3d, 0xca, 0x96, 0x54, 0xd5, 0x2b, 0x7f,
	0x03, 0xa0, 0xfe, 0x04, 0x77, 0x62, 0x9b, 0x99, 0x2f, 0xe4, 0x7b, 0xf9, 0x66, 0xe7, 0xb6, 0xe5,
	0x7f, 0x85, 0x19, 0xbc, 0xa0, 0xec, 0xa6, 0x29, 0x67, 0xf4, 0xfb, 0x7f, 0x03, 0x00, 0x00, 0xff,
	0xff, 0xb7, 0xe7, 0x04, 0x71, 0x59, 0x06, 0x00, 0x00,
}