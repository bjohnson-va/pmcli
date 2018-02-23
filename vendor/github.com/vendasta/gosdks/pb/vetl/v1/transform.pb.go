// Code generated by protoc-gen-go. DO NOT EDIT.
// source: vetl/v1/transform.proto

/*
Package vetl_v1 is a generated protocol buffer package.

It is generated from these files:
	vetl/v1/transform.proto
	vetl/v1/api.proto
	vetl/v1/source.proto
	vetl/v1/entity.proto
	vetl/v1/sink.proto

It has these top-level messages:
	Transform
	KeepProperties
	RenameProperties
	CreateDataSourceRequest
	UpsertTransformRequest
	CreateSubscriptionRequest
	BackfillSubscriptionRequest
	DataSource
	VStoreSource
	DatastoreSource
	Property
	Schema
	Entity
	Struct
	ListValue
	Value
	GeoPoint
	DataSink
	TesseractSink
	VStoreSink
*/
package vetl_v1

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

// Transform is an operation performed on an input row that changes its shape, structure, cardinality, etc.
type Transform struct {
	// Types that are valid to be assigned to Transform:
	//	*Transform_KeepProperties
	//	*Transform_RenameProperties
	Transform isTransform_Transform `protobuf_oneof:"transform"`
}

func (m *Transform) Reset()                    { *m = Transform{} }
func (m *Transform) String() string            { return proto.CompactTextString(m) }
func (*Transform) ProtoMessage()               {}
func (*Transform) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type isTransform_Transform interface {
	isTransform_Transform()
}

type Transform_KeepProperties struct {
	KeepProperties *KeepProperties `protobuf:"bytes,4,opt,name=keep_properties,json=keepProperties,oneof"`
}
type Transform_RenameProperties struct {
	RenameProperties *RenameProperties `protobuf:"bytes,5,opt,name=rename_properties,json=renameProperties,oneof"`
}

func (*Transform_KeepProperties) isTransform_Transform()   {}
func (*Transform_RenameProperties) isTransform_Transform() {}

func (m *Transform) GetTransform() isTransform_Transform {
	if m != nil {
		return m.Transform
	}
	return nil
}

func (m *Transform) GetKeepProperties() *KeepProperties {
	if x, ok := m.GetTransform().(*Transform_KeepProperties); ok {
		return x.KeepProperties
	}
	return nil
}

func (m *Transform) GetRenameProperties() *RenameProperties {
	if x, ok := m.GetTransform().(*Transform_RenameProperties); ok {
		return x.RenameProperties
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Transform) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Transform_OneofMarshaler, _Transform_OneofUnmarshaler, _Transform_OneofSizer, []interface{}{
		(*Transform_KeepProperties)(nil),
		(*Transform_RenameProperties)(nil),
	}
}

func _Transform_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Transform)
	// transform
	switch x := m.Transform.(type) {
	case *Transform_KeepProperties:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.KeepProperties); err != nil {
			return err
		}
	case *Transform_RenameProperties:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.RenameProperties); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Transform.Transform has unexpected type %T", x)
	}
	return nil
}

func _Transform_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Transform)
	switch tag {
	case 4: // transform.keep_properties
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(KeepProperties)
		err := b.DecodeMessage(msg)
		m.Transform = &Transform_KeepProperties{msg}
		return true, err
	case 5: // transform.rename_properties
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(RenameProperties)
		err := b.DecodeMessage(msg)
		m.Transform = &Transform_RenameProperties{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Transform_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Transform)
	// transform
	switch x := m.Transform.(type) {
	case *Transform_KeepProperties:
		s := proto.Size(x.KeepProperties)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Transform_RenameProperties:
		s := proto.Size(x.RenameProperties)
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// KeepProperties removes all properties from a row except for a subset of specified properties
type KeepProperties struct {
	// The properties names to keep, all other names will be dropped
	Names []string `protobuf:"bytes,1,rep,name=names" json:"names,omitempty"`
}

func (m *KeepProperties) Reset()                    { *m = KeepProperties{} }
func (m *KeepProperties) String() string            { return proto.CompactTextString(m) }
func (*KeepProperties) ProtoMessage()               {}
func (*KeepProperties) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *KeepProperties) GetNames() []string {
	if m != nil {
		return m.Names
	}
	return nil
}

// RenameProperties renames the properties with the provided mapping
type RenameProperties struct {
	// map of old the property names to their new one
	// a missing property from this mapping, will keep its name
	Mappings map[string]string `protobuf:"bytes,1,rep,name=mappings" json:"mappings,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *RenameProperties) Reset()                    { *m = RenameProperties{} }
func (m *RenameProperties) String() string            { return proto.CompactTextString(m) }
func (*RenameProperties) ProtoMessage()               {}
func (*RenameProperties) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *RenameProperties) GetMappings() map[string]string {
	if m != nil {
		return m.Mappings
	}
	return nil
}

func init() {
	proto.RegisterType((*Transform)(nil), "vetl.v1.Transform")
	proto.RegisterType((*KeepProperties)(nil), "vetl.v1.KeepProperties")
	proto.RegisterType((*RenameProperties)(nil), "vetl.v1.RenameProperties")
}

func init() { proto.RegisterFile("vetl/v1/transform.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 247 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x2f, 0x4b, 0x2d, 0xc9,
	0xd1, 0x2f, 0x33, 0xd4, 0x2f, 0x29, 0x4a, 0xcc, 0x2b, 0x4e, 0xcb, 0x2f, 0xca, 0xd5, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0x62, 0x07, 0x49, 0xe8, 0x95, 0x19, 0x2a, 0x2d, 0x63, 0xe4, 0xe2, 0x0c,
	0x81, 0x49, 0x0a, 0x39, 0x71, 0xf1, 0x67, 0xa7, 0xa6, 0x16, 0xc4, 0x17, 0x14, 0xe5, 0x17, 0xa4,
	0x16, 0x95, 0x64, 0xa6, 0x16, 0x4b, 0xb0, 0x28, 0x30, 0x6a, 0x70, 0x1b, 0x89, 0xeb, 0x41, 0x35,
	0xe8, 0x79, 0xa7, 0xa6, 0x16, 0x04, 0xc0, 0xa5, 0x3d, 0x18, 0x82, 0xf8, 0xb2, 0x51, 0x44, 0x84,
	0x3c, 0xb8, 0x04, 0x8b, 0x52, 0xf3, 0x12, 0x73, 0x53, 0x91, 0x4d, 0x61, 0x05, 0x9b, 0x22, 0x09,
	0x37, 0x25, 0x08, 0xac, 0x02, 0xc5, 0x1c, 0x81, 0x22, 0x34, 0x31, 0x27, 0x6e, 0x2e, 0x4e, 0xb8,
	0xbb, 0x95, 0xd4, 0xb8, 0xf8, 0x50, 0xad, 0x16, 0x12, 0xe1, 0x62, 0x05, 0x69, 0x28, 0x96, 0x60,
	0x54, 0x60, 0xd6, 0xe0, 0x0c, 0x82, 0x70, 0x94, 0xa6, 0x30, 0x72, 0x09, 0xa0, 0x9b, 0x2e, 0xe4,
	0xcc, 0xc5, 0x91, 0x9b, 0x58, 0x50, 0x90, 0x99, 0x97, 0x0e, 0x51, 0xcd, 0x6d, 0xa4, 0x8e, 0xd3,
	0x29, 0x7a, 0xbe, 0x50, 0x95, 0xae, 0x79, 0x25, 0x45, 0x95, 0x41, 0x70, 0x8d, 0x52, 0xd6, 0x5c,
	0xbc, 0x28, 0x52, 0x42, 0x02, 0x5c, 0xcc, 0xd9, 0xa9, 0x95, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c,
	0x41, 0x20, 0x26, 0xc8, 0x49, 0x65, 0x89, 0x39, 0xa5, 0xa9, 0x12, 0x4c, 0x60, 0x31, 0x08, 0xc7,
	0x8a, 0xc9, 0x82, 0x31, 0x89, 0x0d, 0x1c, 0xee, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd0,
	0xba, 0x19, 0xc7, 0x92, 0x01, 0x00, 0x00,
}