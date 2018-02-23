// Code generated by protoc-gen-go. DO NOT EDIT.
// source: billing/v1/product_summary.proto

package billing_v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf1 "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ProductSummary_Discount_DiscountType int32

const (
	ProductSummary_Discount_FIXED_AMOUNT          ProductSummary_Discount_DiscountType = 0
	ProductSummary_Discount_PERCENT_AMOUNT        ProductSummary_Discount_DiscountType = 1
	ProductSummary_Discount_FIXED_NUMBER_OF_UNITS ProductSummary_Discount_DiscountType = 2
	ProductSummary_Discount_FIXED_AMOUNT_PER_UNIT ProductSummary_Discount_DiscountType = 3
)

var ProductSummary_Discount_DiscountType_name = map[int32]string{
	0: "FIXED_AMOUNT",
	1: "PERCENT_AMOUNT",
	2: "FIXED_NUMBER_OF_UNITS",
	3: "FIXED_AMOUNT_PER_UNIT",
}
var ProductSummary_Discount_DiscountType_value = map[string]int32{
	"FIXED_AMOUNT":          0,
	"PERCENT_AMOUNT":        1,
	"FIXED_NUMBER_OF_UNITS": 2,
	"FIXED_AMOUNT_PER_UNIT": 3,
}

func (x ProductSummary_Discount_DiscountType) String() string {
	return proto.EnumName(ProductSummary_Discount_DiscountType_name, int32(x))
}
func (ProductSummary_Discount_DiscountType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor4, []int{2, 0, 0}
}

type Pricing_RuleType int32

const (
	Pricing_STANDARD  Pricing_RuleType = 0
	Pricing_STAIRSTEP Pricing_RuleType = 1
	Pricing_TIERED    Pricing_RuleType = 2
)

var Pricing_RuleType_name = map[int32]string{
	0: "STANDARD",
	1: "STAIRSTEP",
	2: "TIERED",
}
var Pricing_RuleType_value = map[string]int32{
	"STANDARD":  0,
	"STAIRSTEP": 1,
	"TIERED":    2,
}

func (x Pricing_RuleType) String() string {
	return proto.EnumName(Pricing_RuleType_name, int32(x))
}
func (Pricing_RuleType) EnumDescriptor() ([]byte, []int) { return fileDescriptor4, []int{3, 0} }

type ListProductSummariesRequest struct {
	// Unique ID of the merchant
	MerchantId string `protobuf:"bytes,1,opt,name=merchant_id,json=merchantId" json:"merchant_id,omitempty"`
	// The starting date of the billing period
	StartDate *google_protobuf1.Timestamp `protobuf:"bytes,2,opt,name=start_date,json=startDate" json:"start_date,omitempty"`
	// The ending date of the billing period
	EndDate *google_protobuf1.Timestamp `protobuf:"bytes,3,opt,name=end_date,json=endDate" json:"end_date,omitempty"`
	// Requested position in pagination
	Cursor string `protobuf:"bytes,4,opt,name=cursor" json:"cursor,omitempty"`
	// Limit of results
	PageSize int64 `protobuf:"varint,5,opt,name=page_size,json=pageSize" json:"page_size,omitempty"`
}

func (m *ListProductSummariesRequest) Reset()                    { *m = ListProductSummariesRequest{} }
func (m *ListProductSummariesRequest) String() string            { return proto.CompactTextString(m) }
func (*ListProductSummariesRequest) ProtoMessage()               {}
func (*ListProductSummariesRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *ListProductSummariesRequest) GetMerchantId() string {
	if m != nil {
		return m.MerchantId
	}
	return ""
}

func (m *ListProductSummariesRequest) GetStartDate() *google_protobuf1.Timestamp {
	if m != nil {
		return m.StartDate
	}
	return nil
}

func (m *ListProductSummariesRequest) GetEndDate() *google_protobuf1.Timestamp {
	if m != nil {
		return m.EndDate
	}
	return nil
}

func (m *ListProductSummariesRequest) GetCursor() string {
	if m != nil {
		return m.Cursor
	}
	return ""
}

func (m *ListProductSummariesRequest) GetPageSize() int64 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

type ListProductSummariesResponse struct {
	// The list of product summary results
	Results []*ProductSummary `protobuf:"bytes,1,rep,name=results" json:"results,omitempty"`
	// Next position in pagination
	NextCursor string `protobuf:"bytes,2,opt,name=next_cursor,json=nextCursor" json:"next_cursor,omitempty"`
	// Indicates if more data is available
	HasMore bool `protobuf:"varint,3,opt,name=has_more,json=hasMore" json:"has_more,omitempty"`
}

func (m *ListProductSummariesResponse) Reset()                    { *m = ListProductSummariesResponse{} }
func (m *ListProductSummariesResponse) String() string            { return proto.CompactTextString(m) }
func (*ListProductSummariesResponse) ProtoMessage()               {}
func (*ListProductSummariesResponse) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

func (m *ListProductSummariesResponse) GetResults() []*ProductSummary {
	if m != nil {
		return m.Results
	}
	return nil
}

func (m *ListProductSummariesResponse) GetNextCursor() string {
	if m != nil {
		return m.NextCursor
	}
	return ""
}

func (m *ListProductSummariesResponse) GetHasMore() bool {
	if m != nil {
		return m.HasMore
	}
	return false
}

type ProductSummary struct {
	// The unique SKU of the product
	Sku string `protobuf:"bytes,1,opt,name=sku" json:"sku,omitempty"`
	// Name of the product
	ProductName string `protobuf:"bytes,2,opt,name=product_name,json=productName" json:"product_name,omitempty"`
	// Quantity of all items sold
	TotalQuantity int64 `protobuf:"varint,3,opt,name=total_quantity,json=totalQuantity" json:"total_quantity,omitempty"`
	// Summed dollar amount of all items in cents
	TotalAmount int64 `protobuf:"varint,4,opt,name=total_amount,json=totalAmount" json:"total_amount,omitempty"`
	// Summed dollar amount of all discounts in cents
	TotalDiscount int64 `protobuf:"varint,5,opt,name=total_discount,json=totalDiscount" json:"total_discount,omitempty"`
	// Pricing of the product
	Pricing *Pricing `protobuf:"bytes,6,opt,name=pricing" json:"pricing,omitempty"`
	// Discounts of the product
	Discounts []*ProductSummary_Discount `protobuf:"bytes,7,rep,name=discounts" json:"discounts,omitempty"`
}

func (m *ProductSummary) Reset()                    { *m = ProductSummary{} }
func (m *ProductSummary) String() string            { return proto.CompactTextString(m) }
func (*ProductSummary) ProtoMessage()               {}
func (*ProductSummary) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2} }

func (m *ProductSummary) GetSku() string {
	if m != nil {
		return m.Sku
	}
	return ""
}

func (m *ProductSummary) GetProductName() string {
	if m != nil {
		return m.ProductName
	}
	return ""
}

func (m *ProductSummary) GetTotalQuantity() int64 {
	if m != nil {
		return m.TotalQuantity
	}
	return 0
}

func (m *ProductSummary) GetTotalAmount() int64 {
	if m != nil {
		return m.TotalAmount
	}
	return 0
}

func (m *ProductSummary) GetTotalDiscount() int64 {
	if m != nil {
		return m.TotalDiscount
	}
	return 0
}

func (m *ProductSummary) GetPricing() *Pricing {
	if m != nil {
		return m.Pricing
	}
	return nil
}

func (m *ProductSummary) GetDiscounts() []*ProductSummary_Discount {
	if m != nil {
		return m.Discounts
	}
	return nil
}

type ProductSummary_Discount struct {
	// Discount type eg. fixed_amount, percent_amount, etc.
	Type ProductSummary_Discount_DiscountType `protobuf:"varint,1,opt,name=type,enum=billing.v1.ProductSummary_Discount_DiscountType" json:"type,omitempty"`
	// Value of the discount
	Value int64 `protobuf:"varint,2,opt,name=value" json:"value,omitempty"`
	// Discounted dollar amount in cents
	TotalAmount int64 `protobuf:"varint,3,opt,name=total_amount,json=totalAmount" json:"total_amount,omitempty"`
	// Full english description
	Description string `protobuf:"bytes,4,opt,name=description" json:"description,omitempty"`
}

func (m *ProductSummary_Discount) Reset()                    { *m = ProductSummary_Discount{} }
func (m *ProductSummary_Discount) String() string            { return proto.CompactTextString(m) }
func (*ProductSummary_Discount) ProtoMessage()               {}
func (*ProductSummary_Discount) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2, 0} }

func (m *ProductSummary_Discount) GetType() ProductSummary_Discount_DiscountType {
	if m != nil {
		return m.Type
	}
	return ProductSummary_Discount_FIXED_AMOUNT
}

func (m *ProductSummary_Discount) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *ProductSummary_Discount) GetTotalAmount() int64 {
	if m != nil {
		return m.TotalAmount
	}
	return 0
}

func (m *ProductSummary_Discount) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type Pricing struct {
	// Pricing rule type eg. Standard, Stairstep, etc
	Type Pricing_RuleType `protobuf:"varint,1,opt,name=type,enum=billing.v1.Pricing_RuleType" json:"type,omitempty"`
	// Pricing rules
	Rules []*Pricing_Rule `protobuf:"bytes,2,rep,name=rules" json:"rules,omitempty"`
	// Prepaid sale information
	PrepaidSale *Pricing_PrepaidSale `protobuf:"bytes,3,opt,name=prepaid_sale,json=prepaidSale" json:"prepaid_sale,omitempty"`
}

func (m *Pricing) Reset()                    { *m = Pricing{} }
func (m *Pricing) String() string            { return proto.CompactTextString(m) }
func (*Pricing) ProtoMessage()               {}
func (*Pricing) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{3} }

func (m *Pricing) GetType() Pricing_RuleType {
	if m != nil {
		return m.Type
	}
	return Pricing_STANDARD
}

func (m *Pricing) GetRules() []*Pricing_Rule {
	if m != nil {
		return m.Rules
	}
	return nil
}

func (m *Pricing) GetPrepaidSale() *Pricing_PrepaidSale {
	if m != nil {
		return m.PrepaidSale
	}
	return nil
}

type Pricing_Rule struct {
	// Minimum unit range
	Min int64 `protobuf:"varint,1,opt,name=min" json:"min,omitempty"`
	// Maximum unit range
	Max int64 `protobuf:"varint,2,opt,name=max" json:"max,omitempty"`
	// Number of units sold in range
	Units int64 `protobuf:"varint,3,opt,name=units" json:"units,omitempty"`
	// Price per unit sold in range
	UnitPrice int64 `protobuf:"varint,4,opt,name=unit_price,json=unitPrice" json:"unit_price,omitempty"`
	// Total dollar amount sold in range in cents
	TotalAmount int64 `protobuf:"varint,5,opt,name=total_amount,json=totalAmount" json:"total_amount,omitempty"`
}

func (m *Pricing_Rule) Reset()                    { *m = Pricing_Rule{} }
func (m *Pricing_Rule) String() string            { return proto.CompactTextString(m) }
func (*Pricing_Rule) ProtoMessage()               {}
func (*Pricing_Rule) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{3, 0} }

func (m *Pricing_Rule) GetMin() int64 {
	if m != nil {
		return m.Min
	}
	return 0
}

func (m *Pricing_Rule) GetMax() int64 {
	if m != nil {
		return m.Max
	}
	return 0
}

func (m *Pricing_Rule) GetUnits() int64 {
	if m != nil {
		return m.Units
	}
	return 0
}

func (m *Pricing_Rule) GetUnitPrice() int64 {
	if m != nil {
		return m.UnitPrice
	}
	return 0
}

func (m *Pricing_Rule) GetTotalAmount() int64 {
	if m != nil {
		return m.TotalAmount
	}
	return 0
}

type Pricing_PrepaidSale struct {
	// Number of units prepaid
	UnitsPaid int64 `protobuf:"varint,1,opt,name=units_paid,json=unitsPaid" json:"units_paid,omitempty"`
	// Number of prepaid units sold, can not exceed units_paid
	UnitsSold int64 `protobuf:"varint,2,opt,name=units_sold,json=unitsSold" json:"units_sold,omitempty"`
	// Ending date for prepaid sale
	EndDate *google_protobuf1.Timestamp `protobuf:"bytes,3,opt,name=end_date,json=endDate" json:"end_date,omitempty"`
}

func (m *Pricing_PrepaidSale) Reset()                    { *m = Pricing_PrepaidSale{} }
func (m *Pricing_PrepaidSale) String() string            { return proto.CompactTextString(m) }
func (*Pricing_PrepaidSale) ProtoMessage()               {}
func (*Pricing_PrepaidSale) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{3, 1} }

func (m *Pricing_PrepaidSale) GetUnitsPaid() int64 {
	if m != nil {
		return m.UnitsPaid
	}
	return 0
}

func (m *Pricing_PrepaidSale) GetUnitsSold() int64 {
	if m != nil {
		return m.UnitsSold
	}
	return 0
}

func (m *Pricing_PrepaidSale) GetEndDate() *google_protobuf1.Timestamp {
	if m != nil {
		return m.EndDate
	}
	return nil
}

func init() {
	proto.RegisterType((*ListProductSummariesRequest)(nil), "billing.v1.ListProductSummariesRequest")
	proto.RegisterType((*ListProductSummariesResponse)(nil), "billing.v1.ListProductSummariesResponse")
	proto.RegisterType((*ProductSummary)(nil), "billing.v1.ProductSummary")
	proto.RegisterType((*ProductSummary_Discount)(nil), "billing.v1.ProductSummary.Discount")
	proto.RegisterType((*Pricing)(nil), "billing.v1.Pricing")
	proto.RegisterType((*Pricing_Rule)(nil), "billing.v1.Pricing.Rule")
	proto.RegisterType((*Pricing_PrepaidSale)(nil), "billing.v1.Pricing.PrepaidSale")
	proto.RegisterEnum("billing.v1.ProductSummary_Discount_DiscountType", ProductSummary_Discount_DiscountType_name, ProductSummary_Discount_DiscountType_value)
	proto.RegisterEnum("billing.v1.Pricing_RuleType", Pricing_RuleType_name, Pricing_RuleType_value)
}

func init() { proto.RegisterFile("billing/v1/product_summary.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 747 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x54, 0x4d, 0x6f, 0xf3, 0x44,
	0x10, 0x7e, 0x1d, 0xe7, 0x73, 0x92, 0x37, 0xb2, 0x96, 0x0f, 0xb9, 0x69, 0x51, 0x43, 0x10, 0x52,
	0x2f, 0x38, 0xfd, 0x80, 0x03, 0xc7, 0xb4, 0x71, 0xa5, 0x48, 0x34, 0x0d, 0x6b, 0x57, 0xe2, 0x66,
	0x6d, 0xe3, 0x25, 0x5d, 0xf0, 0x57, 0xbd, 0xeb, 0xaa, 0xe9, 0x91, 0x5f, 0xc0, 0x8d, 0xff, 0xc3,
	0x7f, 0xe1, 0xca, 0x6f, 0x40, 0xbb, 0x6b, 0x37, 0x89, 0xd2, 0x0a, 0xc4, 0x6d, 0xf6, 0x99, 0x67,
	0x9e, 0x19, 0x3f, 0x33, 0x09, 0x0c, 0xef, 0x59, 0x14, 0xb1, 0x64, 0x35, 0x7e, 0x3a, 0x1b, 0x67,
	0x79, 0x1a, 0x16, 0x4b, 0x11, 0xf0, 0x22, 0x8e, 0x49, 0xbe, 0x76, 0xb2, 0x3c, 0x15, 0x29, 0x82,
	0x92, 0xe1, 0x3c, 0x9d, 0x0d, 0x8e, 0x57, 0x69, 0xba, 0x8a, 0xe8, 0x58, 0x65, 0xee, 0x8b, 0x9f,
	0xc7, 0x82, 0xc5, 0x94, 0x0b, 0x12, 0x67, 0x9a, 0x3c, 0xfa, 0xcb, 0x80, 0xc3, 0x1f, 0x18, 0x17,
	0x0b, 0x2d, 0xe5, 0x29, 0x25, 0x46, 0x39, 0xa6, 0x8f, 0x05, 0xe5, 0x02, 0x1d, 0x43, 0x37, 0xa6,
	0xf9, 0xf2, 0x81, 0x24, 0x22, 0x60, 0xa1, 0x6d, 0x0c, 0x8d, 0x93, 0x0e, 0x86, 0x0a, 0x9a, 0x85,
	0xe8, 0x7b, 0x00, 0x2e, 0x48, 0x2e, 0x82, 0x90, 0x08, 0x6a, 0xd7, 0x86, 0xc6, 0x49, 0xf7, 0x7c,
	0xe0, 0xe8, 0xb6, 0x4e, 0xd5, 0xd6, 0xf1, 0xab, 0xb6, 0xb8, 0xa3, 0xd8, 0x53, 0x22, 0x28, 0xfa,
	0x0e, 0xda, 0x34, 0x09, 0x75, 0xa1, 0xf9, 0xaf, 0x85, 0x2d, 0x9a, 0x84, 0xaa, 0xec, 0x73, 0x68,
	0x2e, 0x8b, 0x9c, 0xa7, 0xb9, 0x5d, 0x57, 0xd3, 0x94, 0x2f, 0x74, 0x08, 0x9d, 0x8c, 0xac, 0x68,
	0xc0, 0xd9, 0x0b, 0xb5, 0x1b, 0x43, 0xe3, 0xc4, 0xc4, 0x6d, 0x09, 0x78, 0xec, 0x85, 0x8e, 0x7e,
	0x37, 0xe0, 0xe8, 0xed, 0xef, 0xe4, 0x59, 0x9a, 0x70, 0x8a, 0xbe, 0x85, 0x56, 0x4e, 0x79, 0x11,
	0x09, 0x6e, 0x1b, 0x43, 0x53, 0xcd, 0xb2, 0xf1, 0xd1, 0xd9, 0x29, 0x5b, 0xe3, 0x8a, 0x2a, 0xed,
	0x49, 0xe8, 0xb3, 0x08, 0xca, 0x81, 0x6a, 0xda, 0x1e, 0x09, 0x5d, 0xe9, 0xa1, 0x0e, 0xa0, 0xfd,
	0x40, 0x78, 0x10, 0xa7, 0xb9, 0xfe, 0xc6, 0x36, 0x6e, 0x3d, 0x10, 0x7e, 0x93, 0xe6, 0x74, 0xf4,
	0x67, 0x1d, 0xfa, 0xbb, 0xba, 0xc8, 0x02, 0x93, 0xff, 0x5a, 0x94, 0x2e, 0xcb, 0x10, 0x7d, 0x09,
	0xbd, 0x6a, 0xcb, 0x09, 0x89, 0x69, 0xd9, 0xa1, 0x5b, 0x62, 0x73, 0x12, 0x53, 0xf4, 0x35, 0xf4,
	0x45, 0x2a, 0x48, 0x14, 0x3c, 0x16, 0x24, 0x11, 0x4c, 0xac, 0x55, 0x23, 0x13, 0x7f, 0x54, 0xe8,
	0x8f, 0x25, 0x28, 0x95, 0x34, 0x8d, 0xc4, 0x69, 0x91, 0x08, 0x65, 0x9e, 0x89, 0xbb, 0x0a, 0x9b,
	0x28, 0x68, 0xa3, 0x14, 0x32, 0xbe, 0x54, 0xa4, 0xc6, 0x96, 0xd2, 0xb4, 0x04, 0xd1, 0x37, 0xd0,
	0xca, 0x72, 0xb6, 0x64, 0xc9, 0xca, 0x6e, 0xaa, 0xb5, 0x7d, 0xb2, 0x6b, 0x95, 0x4a, 0xe1, 0x8a,
	0x83, 0x26, 0xd0, 0xa9, 0xf4, 0xb8, 0xdd, 0x52, 0xde, 0x7e, 0xf5, 0xbe, 0xb7, 0x4e, 0xd5, 0x06,
	0x6f, 0xaa, 0x06, 0x7f, 0xd4, 0xa0, 0xfd, 0xda, 0x7e, 0x0a, 0x75, 0xb1, 0xce, 0xa8, 0x72, 0xa9,
	0x7f, 0x7e, 0xfa, 0x1f, 0xa4, 0x5e, 0x03, 0x7f, 0x9d, 0x51, 0xac, 0xaa, 0xd1, 0xa7, 0xd0, 0x78,
	0x22, 0x51, 0xa1, 0x1d, 0x35, 0xb1, 0x7e, 0xec, 0x99, 0x64, 0xee, 0x9b, 0x34, 0x84, 0x6e, 0x48,
	0xf9, 0x32, 0x67, 0x99, 0x60, 0x69, 0x52, 0xde, 0xe0, 0x36, 0x34, 0xfa, 0x05, 0x7a, 0xdb, 0x0d,
	0x91, 0x05, 0xbd, 0xeb, 0xd9, 0x4f, 0xee, 0x34, 0x98, 0xdc, 0xdc, 0xde, 0xcd, 0x7d, 0xeb, 0x03,
	0x42, 0xd0, 0x5f, 0xb8, 0xf8, 0xca, 0x9d, 0xfb, 0x15, 0x66, 0xa0, 0x03, 0xf8, 0x4c, 0xb3, 0xe6,
	0x77, 0x37, 0x97, 0x2e, 0x0e, 0x6e, 0xaf, 0x83, 0xbb, 0xf9, 0xcc, 0xf7, 0xac, 0xda, 0x26, 0xa5,
	0xc9, 0xc1, 0xc2, 0xc5, 0x2a, 0x67, 0x99, 0xa3, 0xbf, 0x4d, 0x68, 0x95, 0x8e, 0xa3, 0xd3, 0x1d,
	0x63, 0x8e, 0xde, 0x58, 0x8a, 0x83, 0x8b, 0x88, 0x6e, 0x99, 0xe0, 0x40, 0x23, 0x2f, 0x22, 0xca,
	0xed, 0x9a, 0x5a, 0x8b, 0xfd, 0x5e, 0x09, 0xd6, 0x34, 0x74, 0x29, 0xaf, 0x91, 0x66, 0x84, 0x85,
	0x01, 0x27, 0x51, 0xf5, 0xab, 0x3d, 0x7e, 0xab, 0x6c, 0xa1, 0x79, 0x1e, 0x89, 0xa8, 0x3c, 0xd7,
	0xd7, 0xc7, 0xe0, 0x37, 0x03, 0xea, 0x52, 0x53, 0x1e, 0x7b, 0xcc, 0x12, 0x35, 0xad, 0x89, 0x65,
	0xa8, 0x10, 0xf2, 0x5c, 0x6e, 0x44, 0x86, 0x72, 0x4b, 0x45, 0xc2, 0x04, 0x2f, 0x17, 0xa1, 0x1f,
	0xe8, 0x0b, 0x00, 0x19, 0x04, 0xf2, 0xc2, 0x68, 0x79, 0xc8, 0x1d, 0x89, 0xc8, 0xe6, 0xfb, 0x4b,
	0x6c, 0xec, 0x2d, 0x51, 0x0e, 0xd1, 0xdd, 0x9a, 0xb0, 0x52, 0xe4, 0x81, 0x44, 0xca, 0x91, 0x94,
	0x22, 0x5f, 0x10, 0x16, 0x6e, 0xd2, 0x3c, 0x8d, 0xc2, 0x72, 0x3e, 0x9d, 0xf6, 0xd2, 0x28, 0xfc,
	0x9f, 0x7f, 0x64, 0xa3, 0x0b, 0x68, 0x57, 0xfb, 0x40, 0x3d, 0x68, 0x7b, 0xfe, 0x64, 0x3e, 0x9d,
	0xe0, 0xa9, 0xf5, 0x01, 0x7d, 0x84, 0x8e, 0xe7, 0x4f, 0x66, 0xd8, 0xf3, 0xdd, 0x85, 0x65, 0x20,
	0x80, 0xa6, 0x3f, 0x73, 0xb1, 0x3b, 0xb5, 0x6a, 0xf7, 0x4d, 0xa5, 0x78, 0xf1, 0x4f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xca, 0x08, 0x0f, 0x19, 0x08, 0x06, 0x00, 0x00,
}
