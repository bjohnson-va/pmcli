// Code generated by protoc-gen-go. DO NOT EDIT.
// source: marketplace_packages/v1/restrictions.proto

package marketplace_packages_v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// A white and black list used for defining restrictions
type PermissionLists struct {
	Whitelist []string `protobuf:"bytes,1,rep,name=whitelist" json:"whitelist,omitempty"`
	Blacklist []string `protobuf:"bytes,2,rep,name=blacklist" json:"blacklist,omitempty"`
}

func (m *PermissionLists) Reset()                    { *m = PermissionLists{} }
func (m *PermissionLists) String() string            { return proto.CompactTextString(m) }
func (*PermissionLists) ProtoMessage()               {}
func (*PermissionLists) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *PermissionLists) GetWhitelist() []string {
	if m != nil {
		return m.Whitelist
	}
	return nil
}

func (m *PermissionLists) GetBlacklist() []string {
	if m != nil {
		return m.Blacklist
	}
	return nil
}

// Restictions is used to specify any restrictions on an app/addon for a given field
type Restrictions struct {
	// Allowed or not allowed countries for this product/addon
	Country *PermissionLists `protobuf:"bytes,1,opt,name=country" json:"country,omitempty"`
}

func (m *Restrictions) Reset()                    { *m = Restrictions{} }
func (m *Restrictions) String() string            { return proto.CompactTextString(m) }
func (*Restrictions) ProtoMessage()               {}
func (*Restrictions) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

func (m *Restrictions) GetCountry() *PermissionLists {
	if m != nil {
		return m.Country
	}
	return nil
}

func init() {
	proto.RegisterType((*PermissionLists)(nil), "marketplace_packages.v1.PermissionLists")
	proto.RegisterType((*Restrictions)(nil), "marketplace_packages.v1.Restrictions")
}

func init() { proto.RegisterFile("marketplace_packages/v1/restrictions.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 171 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xd2, 0xca, 0x4d, 0x2c, 0xca,
	0x4e, 0x2d, 0x29, 0xc8, 0x49, 0x4c, 0x4e, 0x8d, 0x2f, 0x48, 0x4c, 0xce, 0x4e, 0x4c, 0x4f, 0x2d,
	0xd6, 0x2f, 0x33, 0xd4, 0x2f, 0x4a, 0x2d, 0x2e, 0x29, 0xca, 0x4c, 0x2e, 0xc9, 0xcc, 0xcf, 0x2b,
	0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0xc7, 0xa6, 0x56, 0xaf, 0xcc, 0x50, 0xc9, 0x97,
	0x8b, 0x3f, 0x20, 0xb5, 0x28, 0x37, 0xb3, 0xb8, 0x38, 0x33, 0x3f, 0xcf, 0x27, 0xb3, 0xb8, 0xa4,
	0x58, 0x48, 0x86, 0x8b, 0xb3, 0x3c, 0x23, 0xb3, 0x24, 0x35, 0x27, 0xb3, 0xb8, 0x44, 0x82, 0x51,
	0x81, 0x59, 0x83, 0x33, 0x08, 0x21, 0x00, 0x92, 0x4d, 0xca, 0x49, 0x4c, 0xce, 0x06, 0xcb, 0x32,
	0x41, 0x64, 0xe1, 0x02, 0x4a, 0x41, 0x5c, 0x3c, 0x41, 0x48, 0xb6, 0x0b, 0x39, 0x71, 0xb1, 0x27,
	0xe7, 0x97, 0xe6, 0x95, 0x14, 0x55, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x1b, 0x69, 0xe8, 0xe1,
	0x70, 0x89, 0x1e, 0x9a, 0x33, 0x82, 0x60, 0x1a, 0x93, 0xd8, 0xc0, 0x5e, 0x30, 0x06, 0x04, 0x00,
	0x00, 0xff, 0xff, 0xaa, 0xc6, 0x95, 0xbe, 0xf0, 0x00, 0x00, 0x00,
}
