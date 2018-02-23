// Code generated by protoc-gen-go. DO NOT EDIT.
// source: billing/v1/invoice.proto

package billing_v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf1 "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Currency int32

const (
	Currency_USD Currency = 0
	Currency_CAD Currency = 1
	Currency_EUR Currency = 2
	Currency_AUD Currency = 3
	Currency_GBP Currency = 4
	Currency_NZD Currency = 5
)

var Currency_name = map[int32]string{
	0: "USD",
	1: "CAD",
	2: "EUR",
	3: "AUD",
	4: "GBP",
	5: "NZD",
}
var Currency_value = map[string]int32{
	"USD": 0,
	"CAD": 1,
	"EUR": 2,
	"AUD": 3,
	"GBP": 4,
	"NZD": 5,
}

func (x Currency) String() string {
	return proto.EnumName(Currency_name, int32(x))
}
func (Currency) EnumDescriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

type Invoice_Status int32

const (
	Invoice_DRAFT   Invoice_Status = 0
	Invoice_DUE     Invoice_Status = 1
	Invoice_OVERDUE Invoice_Status = 2
	Invoice_PAID    Invoice_Status = 3
)

var Invoice_Status_name = map[int32]string{
	0: "DRAFT",
	1: "DUE",
	2: "OVERDUE",
	3: "PAID",
}
var Invoice_Status_value = map[string]int32{
	"DRAFT":   0,
	"DUE":     1,
	"OVERDUE": 2,
	"PAID":    3,
}

func (x Invoice_Status) String() string {
	return proto.EnumName(Invoice_Status_name, int32(x))
}
func (Invoice_Status) EnumDescriptor() ([]byte, []int) { return fileDescriptor5, []int{1, 0} }

type LineItem struct {
	// The identifier of what products/subscription this lineitem refers to. Does not conform to the SKU standard
	Sku string `protobuf:"bytes,1,opt,name=sku" json:"sku,omitempty"`
	// How much each item cost (in cents)
	UnitPrice int64 `protobuf:"varint,2,opt,name=unit_price,json=unitPrice" json:"unit_price,omitempty"`
	// The number of items purchased
	Quantity int64 `protobuf:"varint,3,opt,name=quantity" json:"quantity,omitempty"`
	// The name or human readable description of what is being purchased
	Description string `protobuf:"bytes,4,opt,name=description" json:"description,omitempty"`
	// UnitPrice * Quantity
	Total int64 `protobuf:"varint,5,opt,name=total" json:"total,omitempty"`
}

func (m *LineItem) Reset()                    { *m = LineItem{} }
func (m *LineItem) String() string            { return proto.CompactTextString(m) }
func (*LineItem) ProtoMessage()               {}
func (*LineItem) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

func (m *LineItem) GetSku() string {
	if m != nil {
		return m.Sku
	}
	return ""
}

func (m *LineItem) GetUnitPrice() int64 {
	if m != nil {
		return m.UnitPrice
	}
	return 0
}

func (m *LineItem) GetQuantity() int64 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

func (m *LineItem) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *LineItem) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

type Invoice struct {
	// Identifier for the invoice
	InvoiceId int64 `protobuf:"varint,1,opt,name=invoice_id,json=invoiceId" json:"invoice_id,omitempty"`
	// Identifier for the merchant
	MerchantId string `protobuf:"bytes,2,opt,name=merchant_id,json=merchantId" json:"merchant_id,omitempty"`
	// Status of the invoice
	Status Invoice_Status `protobuf:"varint,3,opt,name=status,enum=billing.v1.Invoice_Status" json:"status,omitempty"`
	// Total amount of the invoice (in cents)
	Total int64 `protobuf:"varint,4,opt,name=total" json:"total,omitempty"`
	// Amount paid on the invoice (in cents)
	Paid int64 `protobuf:"varint,5,opt,name=paid" json:"paid,omitempty"`
	// Start date of the invoice period
	PeriodStart *google_protobuf1.Timestamp `protobuf:"bytes,6,opt,name=period_start,json=periodStart" json:"period_start,omitempty"`
	// End date of the invoice period
	PeriodEnd *google_protobuf1.Timestamp `protobuf:"bytes,7,opt,name=period_end,json=periodEnd" json:"period_end,omitempty"`
	// Due date of the invoice
	Due *google_protobuf1.Timestamp `protobuf:"bytes,8,opt,name=due" json:"due,omitempty"`
	// The grace period (in days) from the period_end for this invoice to be paid
	Terms int64 `protobuf:"varint,9,opt,name=terms" json:"terms,omitempty"`
	// Created date of the invoice
	Created *google_protobuf1.Timestamp `protobuf:"bytes,10,opt,name=created" json:"created,omitempty"`
	// Last updated date of the invoice
	Updated *google_protobuf1.Timestamp `protobuf:"bytes,11,opt,name=updated" json:"updated,omitempty"`
	// The line items of the invoice
	LineItems []*LineItem `protobuf:"bytes,12,rep,name=line_items,json=lineItems" json:"line_items,omitempty"`
	// The posted date of the invoice
	Posted *google_protobuf1.Timestamp `protobuf:"bytes,13,opt,name=posted" json:"posted,omitempty"`
	// The currency that this invoice is in
	Currency Currency `protobuf:"varint,14,opt,name=currency,enum=billing.v1.Currency" json:"currency,omitempty"`
}

func (m *Invoice) Reset()                    { *m = Invoice{} }
func (m *Invoice) String() string            { return proto.CompactTextString(m) }
func (*Invoice) ProtoMessage()               {}
func (*Invoice) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{1} }

func (m *Invoice) GetInvoiceId() int64 {
	if m != nil {
		return m.InvoiceId
	}
	return 0
}

func (m *Invoice) GetMerchantId() string {
	if m != nil {
		return m.MerchantId
	}
	return ""
}

func (m *Invoice) GetStatus() Invoice_Status {
	if m != nil {
		return m.Status
	}
	return Invoice_DRAFT
}

func (m *Invoice) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *Invoice) GetPaid() int64 {
	if m != nil {
		return m.Paid
	}
	return 0
}

func (m *Invoice) GetPeriodStart() *google_protobuf1.Timestamp {
	if m != nil {
		return m.PeriodStart
	}
	return nil
}

func (m *Invoice) GetPeriodEnd() *google_protobuf1.Timestamp {
	if m != nil {
		return m.PeriodEnd
	}
	return nil
}

func (m *Invoice) GetDue() *google_protobuf1.Timestamp {
	if m != nil {
		return m.Due
	}
	return nil
}

func (m *Invoice) GetTerms() int64 {
	if m != nil {
		return m.Terms
	}
	return 0
}

func (m *Invoice) GetCreated() *google_protobuf1.Timestamp {
	if m != nil {
		return m.Created
	}
	return nil
}

func (m *Invoice) GetUpdated() *google_protobuf1.Timestamp {
	if m != nil {
		return m.Updated
	}
	return nil
}

func (m *Invoice) GetLineItems() []*LineItem {
	if m != nil {
		return m.LineItems
	}
	return nil
}

func (m *Invoice) GetPosted() *google_protobuf1.Timestamp {
	if m != nil {
		return m.Posted
	}
	return nil
}

func (m *Invoice) GetCurrency() Currency {
	if m != nil {
		return m.Currency
	}
	return Currency_USD
}

func init() {
	proto.RegisterType((*LineItem)(nil), "billing.v1.LineItem")
	proto.RegisterType((*Invoice)(nil), "billing.v1.Invoice")
	proto.RegisterEnum("billing.v1.Currency", Currency_name, Currency_value)
	proto.RegisterEnum("billing.v1.Invoice_Status", Invoice_Status_name, Invoice_Status_value)
}

func init() { proto.RegisterFile("billing/v1/invoice.proto", fileDescriptor5) }

var fileDescriptor5 = []byte{
	// 512 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x52, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xc5, 0x71, 0x3e, 0xec, 0x71, 0xa9, 0xac, 0x51, 0x0f, 0xab, 0x48, 0xa8, 0x51, 0x4f, 0x11,
	0x42, 0x0e, 0x4d, 0xb8, 0x70, 0x40, 0x22, 0xd4, 0x01, 0x45, 0x42, 0x10, 0x39, 0x0d, 0x07, 0x2e,
	0x91, 0xe3, 0x5d, 0xc2, 0x0a, 0x7b, 0x6d, 0xec, 0x75, 0xa4, 0xfe, 0x0a, 0x7e, 0x03, 0xff, 0x14,
	0xed, 0x7a, 0xdd, 0x86, 0x93, 0x6f, 0x6f, 0x66, 0xde, 0xdb, 0x37, 0x1f, 0x0b, 0xe4, 0xc0, 0xd3,
	0x94, 0x8b, 0xe3, 0xec, 0x74, 0x3b, 0xe3, 0xe2, 0x94, 0xf3, 0x84, 0x05, 0x45, 0x99, 0xcb, 0x1c,
	0xc1, 0x54, 0x82, 0xd3, 0xed, 0xf8, 0xfa, 0x98, 0xe7, 0xc7, 0x94, 0xcd, 0x74, 0xe5, 0x50, 0xff,
	0x98, 0x49, 0x9e, 0xb1, 0x4a, 0xc6, 0x59, 0xd1, 0x90, 0x6f, 0xfe, 0x58, 0xe0, 0x7c, 0xe6, 0x82,
	0xad, 0x25, 0xcb, 0xd0, 0x07, 0xbb, 0xfa, 0x55, 0x13, 0x6b, 0x62, 0x4d, 0xdd, 0x48, 0x41, 0x7c,
	0x01, 0x50, 0x0b, 0x2e, 0xf7, 0x45, 0xc9, 0x13, 0x46, 0x7a, 0x13, 0x6b, 0x6a, 0x47, 0xae, 0xca,
	0x6c, 0x54, 0x02, 0xc7, 0xe0, 0xfc, 0xae, 0x63, 0x21, 0xb9, 0x7c, 0x20, 0xb6, 0x2e, 0x3e, 0xc6,
	0x38, 0x01, 0x8f, 0xb2, 0x2a, 0x29, 0x79, 0x21, 0x79, 0x2e, 0x48, 0x5f, 0x3f, 0x7a, 0x9e, 0xc2,
	0x2b, 0x18, 0xc8, 0x5c, 0xc6, 0x29, 0x19, 0x68, 0x69, 0x13, 0xdc, 0xfc, 0x1d, 0xc0, 0x68, 0xdd,
	0x0c, 0xa4, 0xec, 0xcd, 0x6c, 0x7b, 0x4e, 0x75, 0x5f, 0x76, 0xe4, 0x9a, 0xcc, 0x9a, 0xe2, 0x35,
	0x78, 0x19, 0x2b, 0x93, 0x9f, 0xb1, 0x90, 0xaa, 0xde, 0xd3, 0x16, 0xd0, 0xa6, 0xd6, 0x14, 0xe7,
	0x30, 0xac, 0x64, 0x2c, 0xeb, 0x4a, 0x77, 0x77, 0x39, 0x1f, 0x07, 0x4f, 0xbb, 0x09, 0x8c, 0x49,
	0xb0, 0xd5, 0x8c, 0xc8, 0x30, 0x9f, 0xba, 0xea, 0x9f, 0x75, 0x85, 0x08, 0xfd, 0x22, 0xe6, 0xd4,
	0xb4, 0xaa, 0x31, 0xbe, 0x83, 0x8b, 0x82, 0x95, 0x3c, 0xa7, 0xfb, 0x4a, 0xc6, 0xa5, 0x24, 0xc3,
	0x89, 0x35, 0xf5, 0xe6, 0xe3, 0xa0, 0xd9, 0x79, 0xd0, 0xee, 0x3c, 0xb8, 0x6f, 0x77, 0x1e, 0x79,
	0x0d, 0x7f, 0xab, 0xe8, 0xf8, 0x16, 0xc0, 0xc8, 0x99, 0xa0, 0x64, 0xd4, 0x29, 0x76, 0x1b, 0xf6,
	0x4a, 0x50, 0x7c, 0x05, 0x36, 0xad, 0x19, 0x71, 0x3a, 0x35, 0x8a, 0xa6, 0x27, 0x62, 0x65, 0x56,
	0x11, 0xd7, 0x4c, 0xa4, 0x02, 0x7c, 0x03, 0xa3, 0xa4, 0x64, 0xb1, 0x64, 0x94, 0x40, 0xe7, 0x3b,
	0x2d, 0x55, 0xa9, 0xea, 0x82, 0x6a, 0x95, 0xd7, 0xad, 0x32, 0x54, 0x5c, 0x00, 0xa4, 0x5c, 0xb0,
	0x3d, 0x97, 0x2c, 0xab, 0xc8, 0xc5, 0xc4, 0x9e, 0x7a, 0xf3, 0xab, 0xf3, 0x5b, 0xb4, 0x5f, 0x30,
	0x72, 0x53, 0x83, 0x2a, 0x75, 0xbc, 0x22, 0xaf, 0x94, 0xd3, 0xf3, 0x4e, 0x27, 0xc3, 0xc4, 0xd7,
	0xe0, 0x24, 0x75, 0x59, 0x32, 0x91, 0x3c, 0x90, 0x4b, 0x7d, 0xf2, 0xff, 0x6c, 0xee, 0x4c, 0x2d,
	0x7a, 0x64, 0xdd, 0x2c, 0x60, 0xd8, 0x7c, 0x00, 0x74, 0x61, 0x10, 0x46, 0xcb, 0x8f, 0xf7, 0xfe,
	0x33, 0x1c, 0x81, 0x1d, 0xee, 0x56, 0xbe, 0x85, 0x1e, 0x8c, 0xbe, 0x7e, 0x5b, 0x45, 0x2a, 0xe8,
	0xa1, 0x03, 0xfd, 0xcd, 0x72, 0x1d, 0xfa, 0xf6, 0xcb, 0xf7, 0xe0, 0xb4, 0x4f, 0x29, 0xee, 0x6e,
	0x1b, 0x36, 0xa2, 0xbb, 0x65, 0xe8, 0x5b, 0x0a, 0xac, 0x76, 0x91, 0xdf, 0x53, 0x60, 0xb9, 0x0b,
	0x7d, 0x5b, 0x81, 0x4f, 0x1f, 0x36, 0x7e, 0x5f, 0x81, 0x2f, 0xdf, 0x43, 0x7f, 0x70, 0x18, 0xea,
	0x21, 0x16, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x82, 0xcf, 0x3e, 0x5d, 0xc7, 0x03, 0x00, 0x00,
}
