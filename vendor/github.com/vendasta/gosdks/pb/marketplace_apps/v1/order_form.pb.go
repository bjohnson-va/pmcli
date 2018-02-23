// Code generated by protoc-gen-go. DO NOT EDIT.
// source: marketplace_apps/v1/order_form.proto

package marketplaceapps_v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Field asking user for information when they try to activate the app
type OrderFormField struct {
	Label       string   `protobuf:"bytes,1,opt,name=label" json:"label,omitempty"`
	Id          string   `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
	Type        string   `protobuf:"bytes,3,opt,name=type" json:"type,omitempty"`
	Options     []string `protobuf:"bytes,4,rep,name=options" json:"options,omitempty"`
	Description string   `protobuf:"bytes,6,opt,name=description" json:"description,omitempty"`
	Required    bool     `protobuf:"varint,7,opt,name=required" json:"required,omitempty"`
	// For file upload fields this is where we upload and store the file
	UploadUrl string `protobuf:"bytes,8,opt,name=upload_url,json=uploadUrl" json:"upload_url,omitempty"`
}

func (m *OrderFormField) Reset()                    { *m = OrderFormField{} }
func (m *OrderFormField) String() string            { return proto.CompactTextString(m) }
func (*OrderFormField) ProtoMessage()               {}
func (*OrderFormField) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *OrderFormField) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *OrderFormField) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *OrderFormField) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *OrderFormField) GetOptions() []string {
	if m != nil {
		return m.Options
	}
	return nil
}

func (m *OrderFormField) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *OrderFormField) GetRequired() bool {
	if m != nil {
		return m.Required
	}
	return false
}

func (m *OrderFormField) GetUploadUrl() string {
	if m != nil {
		return m.UploadUrl
	}
	return ""
}

// Common form fields among apps, can be autofilled
type IncludedCommonFormFields struct {
	BusinessName           bool `protobuf:"varint,1,opt,name=business_name,json=businessName" json:"business_name,omitempty"`
	BusinessAddress        bool `protobuf:"varint,2,opt,name=business_address,json=businessAddress" json:"business_address,omitempty"`
	BusinessPhoneNumber    bool `protobuf:"varint,3,opt,name=business_phone_number,json=businessPhoneNumber" json:"business_phone_number,omitempty"`
	BusinessAccountGroupId bool `protobuf:"varint,4,opt,name=business_account_group_id,json=businessAccountGroupId" json:"business_account_group_id,omitempty"`
	SalespersonName        bool `protobuf:"varint,5,opt,name=salesperson_name,json=salespersonName" json:"salesperson_name,omitempty"`
	SalespersonPhoneNumber bool `protobuf:"varint,6,opt,name=salesperson_phone_number,json=salespersonPhoneNumber" json:"salesperson_phone_number,omitempty"`
	SalespersonEmail       bool `protobuf:"varint,7,opt,name=salesperson_email,json=salespersonEmail" json:"salesperson_email,omitempty"`
	ContactName            bool `protobuf:"varint,8,opt,name=contact_name,json=contactName" json:"contact_name,omitempty"`
	ContactPhoneNumber     bool `protobuf:"varint,9,opt,name=contact_phone_number,json=contactPhoneNumber" json:"contact_phone_number,omitempty"`
	ContactEmail           bool `protobuf:"varint,10,opt,name=contact_email,json=contactEmail" json:"contact_email,omitempty"`
}

func (m *IncludedCommonFormFields) Reset()                    { *m = IncludedCommonFormFields{} }
func (m *IncludedCommonFormFields) String() string            { return proto.CompactTextString(m) }
func (*IncludedCommonFormFields) ProtoMessage()               {}
func (*IncludedCommonFormFields) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *IncludedCommonFormFields) GetBusinessName() bool {
	if m != nil {
		return m.BusinessName
	}
	return false
}

func (m *IncludedCommonFormFields) GetBusinessAddress() bool {
	if m != nil {
		return m.BusinessAddress
	}
	return false
}

func (m *IncludedCommonFormFields) GetBusinessPhoneNumber() bool {
	if m != nil {
		return m.BusinessPhoneNumber
	}
	return false
}

func (m *IncludedCommonFormFields) GetBusinessAccountGroupId() bool {
	if m != nil {
		return m.BusinessAccountGroupId
	}
	return false
}

func (m *IncludedCommonFormFields) GetSalespersonName() bool {
	if m != nil {
		return m.SalespersonName
	}
	return false
}

func (m *IncludedCommonFormFields) GetSalespersonPhoneNumber() bool {
	if m != nil {
		return m.SalespersonPhoneNumber
	}
	return false
}

func (m *IncludedCommonFormFields) GetSalespersonEmail() bool {
	if m != nil {
		return m.SalespersonEmail
	}
	return false
}

func (m *IncludedCommonFormFields) GetContactName() bool {
	if m != nil {
		return m.ContactName
	}
	return false
}

func (m *IncludedCommonFormFields) GetContactPhoneNumber() bool {
	if m != nil {
		return m.ContactPhoneNumber
	}
	return false
}

func (m *IncludedCommonFormFields) GetContactEmail() bool {
	if m != nil {
		return m.ContactEmail
	}
	return false
}

// Order form fields
type OrderForm struct {
	// Order form: fields in the order form
	OrderForm []*OrderFormField `protobuf:"bytes,1,rep,name=order_form,json=orderForm" json:"order_form,omitempty"`
	// Order form: common form fields
	CommonForm *IncludedCommonFormFields `protobuf:"bytes,2,opt,name=common_form,json=commonForm" json:"common_form,omitempty"`
	// Order form: message shown to users upon activating the app
	ActivationMessage string `protobuf:"bytes,3,opt,name=activation_message,json=activationMessage" json:"activation_message,omitempty"`
}

func (m *OrderForm) Reset()                    { *m = OrderForm{} }
func (m *OrderForm) String() string            { return proto.CompactTextString(m) }
func (*OrderForm) ProtoMessage()               {}
func (*OrderForm) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

func (m *OrderForm) GetOrderForm() []*OrderFormField {
	if m != nil {
		return m.OrderForm
	}
	return nil
}

func (m *OrderForm) GetCommonForm() *IncludedCommonFormFields {
	if m != nil {
		return m.CommonForm
	}
	return nil
}

func (m *OrderForm) GetActivationMessage() string {
	if m != nil {
		return m.ActivationMessage
	}
	return ""
}

func init() {
	proto.RegisterType((*OrderFormField)(nil), "marketplaceapps.v1.OrderFormField")
	proto.RegisterType((*IncludedCommonFormFields)(nil), "marketplaceapps.v1.IncludedCommonFormFields")
	proto.RegisterType((*OrderForm)(nil), "marketplaceapps.v1.OrderForm")
}

func init() { proto.RegisterFile("marketplace_apps/v1/order_form.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 501 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x53, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x95, 0xf3, 0x55, 0x7b, 0xd2, 0x96, 0x76, 0x29, 0x68, 0x41, 0x42, 0x32, 0x29, 0x87, 0x20,
	0x20, 0xa5, 0xe1, 0x02, 0xc7, 0x0a, 0x51, 0x54, 0xa4, 0x16, 0x64, 0x89, 0xb3, 0xb5, 0xf1, 0x0e,
	0xc5, 0x62, 0xbf, 0xd8, 0xb5, 0x23, 0xf1, 0x0b, 0xb9, 0xf0, 0x37, 0xf8, 0x1f, 0xc8, 0xbb, 0xb1,
	0xe3, 0x08, 0xb8, 0x65, 0xde, 0x7b, 0x33, 0xf3, 0x26, 0xfb, 0x0c, 0x4f, 0x24, 0xb3, 0xdf, 0xb0,
	0x32, 0x82, 0x15, 0x98, 0x33, 0x63, 0xdc, 0xd9, 0xfa, 0xfc, 0x4c, 0x5b, 0x8e, 0x36, 0xff, 0xa2,
	0xad, 0x5c, 0x18, 0xab, 0x2b, 0x4d, 0x48, 0x4f, 0xd5, 0x88, 0x16, 0xeb, 0xf3, 0xd9, 0xcf, 0x08,
	0x0e, 0x3f, 0x36, 0xc2, 0x4b, 0x6d, 0xe5, 0x65, 0x89, 0x82, 0x93, 0x13, 0x18, 0x0b, 0xb6, 0x42,
	0x41, 0xa3, 0x34, 0x9a, 0x27, 0x59, 0x28, 0xc8, 0x21, 0x0c, 0x4a, 0x4e, 0x07, 0x1e, 0x1a, 0x94,
	0x9c, 0x10, 0x18, 0x55, 0x3f, 0x0c, 0xd2, 0xa1, 0x47, 0xfc, 0x6f, 0x42, 0x61, 0x4f, 0x9b, 0xaa,
	0xd4, 0xca, 0xd1, 0x51, 0x3a, 0x9c, 0x27, 0x59, 0x5b, 0x92, 0x14, 0xa6, 0x1c, 0x5d, 0x61, 0x4b,
	0x5f, 0xd3, 0x89, 0x6f, 0xea, 0x43, 0xe4, 0x21, 0xc4, 0x16, 0xbf, 0xd7, 0xa5, 0x45, 0x4e, 0xf7,
	0xd2, 0x68, 0x1e, 0x67, 0x5d, 0x4d, 0x1e, 0x01, 0xd4, 0x46, 0x68, 0xc6, 0xf3, 0xda, 0x0a, 0x1a,
	0xfb, 0xe6, 0x24, 0x20, 0x9f, 0xad, 0xf8, 0x30, 0x8a, 0xc7, 0x47, 0x93, 0xd9, 0xef, 0x21, 0xd0,
	0x2b, 0x55, 0x88, 0x9a, 0x23, 0x7f, 0xab, 0xa5, 0xd4, 0xaa, 0x3b, 0xc9, 0x91, 0x53, 0x38, 0x58,
	0xd5, 0xae, 0x54, 0xe8, 0x5c, 0xae, 0x98, 0x44, 0x7f, 0x5b, 0x9c, 0xed, 0xb7, 0xe0, 0x0d, 0x93,
	0x48, 0x9e, 0xc2, 0x51, 0x27, 0x62, 0x9c, 0x5b, 0x74, 0xce, 0x1f, 0x1c, 0x67, 0x77, 0x5a, 0xfc,
	0x22, 0xc0, 0x64, 0x09, 0xf7, 0x3a, 0xa9, 0xf9, 0xaa, 0x15, 0xe6, 0xaa, 0x96, 0x2b, 0xb4, 0xfe,
	0xef, 0x88, 0xb3, 0xbb, 0x2d, 0xf9, 0xa9, 0xe1, 0x6e, 0x3c, 0x45, 0xde, 0xc0, 0x83, 0xed, 0xf8,
	0xa2, 0xd0, 0xb5, 0xaa, 0xf2, 0x5b, 0xab, 0x6b, 0x93, 0x97, 0x9c, 0x8e, 0x7c, 0xdf, 0xfd, 0x6e,
	0x4f, 0xe0, 0xdf, 0x37, 0xf4, 0x15, 0x6f, 0x9c, 0x39, 0x26, 0xd0, 0x19, 0xb4, 0x4e, 0xab, 0x70,
	0xc1, 0x38, 0x38, 0xeb, 0xe1, 0xfe, 0x88, 0xd7, 0x40, 0xfb, 0xd2, 0x1d, 0x73, 0x93, 0xb0, 0xa4,
	0xc7, 0xf7, 0xfd, 0x3d, 0x83, 0xe3, 0x7e, 0x27, 0x4a, 0x56, 0x8a, 0xcd, 0x53, 0xf4, 0xb7, 0xbf,
	0x6b, 0x70, 0xf2, 0x18, 0xf6, 0x0b, 0xad, 0x2a, 0x56, 0x54, 0xc1, 0x4d, 0xec, 0x75, 0xd3, 0x0d,
	0xe6, 0x9d, 0xbc, 0x84, 0x93, 0x56, 0xb2, 0xe3, 0x22, 0xf1, 0x52, 0xb2, 0xe1, 0xfa, 0x0e, 0x4e,
	0xe1, 0xa0, 0xed, 0x08, 0xdb, 0x21, 0xbc, 0xd2, 0x06, 0xf4, 0x9b, 0x67, 0xbf, 0x22, 0x48, 0xba,
	0xc4, 0x92, 0x0b, 0x80, 0x6d, 0xce, 0x69, 0x94, 0x0e, 0xe7, 0xd3, 0xe5, 0x6c, 0xf1, 0x77, 0xd0,
	0x17, 0xbb, 0x21, 0xcf, 0x12, 0xdd, 0x8d, 0xb8, 0x86, 0x69, 0xe1, 0xf3, 0x12, 0x66, 0x34, 0x2f,
	0x3e, 0x5d, 0x3e, 0xff, 0xd7, 0x8c, 0xff, 0xc5, 0x2b, 0x83, 0xa2, 0x43, 0xc8, 0x0b, 0x20, 0xac,
	0xa8, 0xca, 0x35, 0x6b, 0x62, 0x9d, 0x4b, 0x74, 0x8e, 0xdd, 0xb6, 0x9f, 0xc9, 0xf1, 0x96, 0xb9,
	0x0e, 0xc4, 0x6a, 0xe2, 0xbf, 0xcd, 0x57, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x34, 0xc4, 0x6d,
	0xda, 0xc3, 0x03, 0x00, 0x00,
}
