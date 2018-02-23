package python

import (
	"fmt"
	"strings"

	google_protobuf "github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/vendasta/mscli/tools/sdkgen/internal"
)

// PythonMessageEnumDescriptors structure of message and enum descripors
type PythonMessageEnumDescriptors struct {
	*internal.MessageEnumDescriptors
}

// PyFileName filename of a python file
func (s *PythonMessageEnumDescriptors) PyFileName() string {
	fileName := s.GetFileNameFromFullPath(s.FileName)
	return fmt.Sprintf("%s", fileName)
}

// PyProtoFileName filename of a python file
func (s *PythonMessageEnumDescriptors) PyProtoFileName() string {
	fileName := s.GetFileNameFromFullPath(s.FileName)
	return fmt.Sprintf("%s_pb2", fileName)
}

// FieldToVobjectProperty converts a proto Field to a vobject property
func (s *PythonMessageEnumDescriptors) FieldToVobjectProperty(field google_protobuf.FieldDescriptorProto) string {
	property := ""
	propertyAttributes := []string{}
	switch *field.Type {
	case google_protobuf.FieldDescriptorProto_TYPE_DOUBLE:
		property = "FloatProperty"
	case google_protobuf.FieldDescriptorProto_TYPE_FLOAT:
		property = "FloatProperty"
	case google_protobuf.FieldDescriptorProto_TYPE_INT64:
		property = "IntegerProperty"
	case google_protobuf.FieldDescriptorProto_TYPE_UINT64:
		property = "IntegerProperty"
	case google_protobuf.FieldDescriptorProto_TYPE_INT32:
		property = "IntegerProperty"
	case google_protobuf.FieldDescriptorProto_TYPE_FIXED64:
		property = "IntegerProperty"
	case google_protobuf.FieldDescriptorProto_TYPE_FIXED32:
		property = "IntegerProperty"
	case google_protobuf.FieldDescriptorProto_TYPE_BOOL:
		property = "BooleanProperty"
	case google_protobuf.FieldDescriptorProto_TYPE_STRING:
		property = "StringProperty"
	case google_protobuf.FieldDescriptorProto_TYPE_MESSAGE:
		if internal.IsWrapperField(&field) {
			property = WrapperFieldToPropertyType(field)
		} else if internal.IsTimestampField(&field) {
			property = "DateTimeProperty"
		} else {
			property = "StructuredProperty"
			propertyAttributes = append(propertyAttributes, s.ExtractTypeName(*field.TypeName))
		}
	case google_protobuf.FieldDescriptorProto_TYPE_BYTES:
		property = "StringProperty"
	case google_protobuf.FieldDescriptorProto_TYPE_UINT32:
		property = "IntegerProperty"
	case google_protobuf.FieldDescriptorProto_TYPE_ENUM:
		property = "IntegerProperty"
		propertyAttributes = append(propertyAttributes, fmt.Sprintf("choices=%s", s.ExtractTypeName(*field.TypeName)))
	case google_protobuf.FieldDescriptorProto_TYPE_SFIXED32:
		property = "IntegerProperty"
	case google_protobuf.FieldDescriptorProto_TYPE_SFIXED64:
		property = "IntegerProperty"
	case google_protobuf.FieldDescriptorProto_TYPE_SINT32:
		property = "IntegerProperty"
	case google_protobuf.FieldDescriptorProto_TYPE_SINT64:
		property = "IntegerProperty"
	}

	switch *field.Label {
	case google_protobuf.FieldDescriptorProto_LABEL_REPEATED:
		propertyAttributes = append(propertyAttributes, "repeated=True")
	}

	property += fmt.Sprintf("(%s)", strings.Join(propertyAttributes, ", "))

	return property
}

// FieldToDomainValue converts a proto Field to it's domain value expression
func (s *PythonMessageEnumDescriptors) FieldToDomainValue(field google_protobuf.FieldDescriptorProto) string {
	messageValue := fmt.Sprintf("message.%s", *field.Name)
	expression := ""
	switch *field.Type {
	case google_protobuf.FieldDescriptorProto_TYPE_STRING:
		if *field.Label == google_protobuf.FieldDescriptorProto_LABEL_REPEATED {
			expression = fmt.Sprintf("extract_repeated(%s)", messageValue)
		} else {
			expression = fmt.Sprintf("extract_falsy_as_none(%s)", messageValue)
		}
	case google_protobuf.FieldDescriptorProto_TYPE_MESSAGE:
		if internal.IsWrapperField(&field) {
			if *field.Label == google_protobuf.FieldDescriptorProto_LABEL_REPEATED {
				expression = fmt.Sprintf("[x.value for x in %s]", messageValue)
			} else {
				expression = fmt.Sprintf("%s.value if message.HasField('%s') else None", messageValue, *field.Name)
			}
		} else if internal.IsTimestampField(&field) {
			if *field.Label == google_protobuf.FieldDescriptorProto_LABEL_REPEATED {
				expression = fmt.Sprintf("[extract_datetime(x) for x in %s]", messageValue)
			} else {
				expression = fmt.Sprintf("extract_datetime(%s)", messageValue)
			}
		} else {
			if *field.Label == google_protobuf.FieldDescriptorProto_LABEL_REPEATED {
				expression = fmt.Sprintf("[%s.from_proto(x) for x in %s]", s.ExtractTypeName(*field.TypeName), messageValue)
			} else {
				expression = fmt.Sprintf("%s.from_proto(%s) if message.HasField('%s') else None", s.ExtractTypeName(*field.TypeName), messageValue, *field.Name)
			}
		}
	case google_protobuf.FieldDescriptorProto_TYPE_ENUM:
		if *field.Label == google_protobuf.FieldDescriptorProto_LABEL_REPEATED {
			expression = fmt.Sprintf("[x for x in %s]", messageValue)
		} else {
			expression = messageValue
		}
	default:
		expression = messageValue
	}

	return expression
}

// FieldToProtoValue converts a domain Field to it's proto value expression
func (s *PythonMessageEnumDescriptors) FieldToProtoValue(field google_protobuf.FieldDescriptorProto) string {
	messageValue := fmt.Sprintf("self.%s", *field.Name)
	expression := ""
	switch *field.Type {
	case google_protobuf.FieldDescriptorProto_TYPE_ENUM:
		if *field.Label == google_protobuf.FieldDescriptorProto_LABEL_REPEATED {
			expression = fmt.Sprintf("[x.value for x in %s]", messageValue)
		} else {
			expression = fmt.Sprintf("%s.value", messageValue)
		}

	case google_protobuf.FieldDescriptorProto_TYPE_MESSAGE:
		if internal.IsWrapperField(&field) {
			expression = messageValue
		} else if *field.Label == google_protobuf.FieldDescriptorProto_LABEL_REPEATED {
			expression = fmt.Sprintf("[x.to_proto() for x in %s]", messageValue)
		} else {
			expression = fmt.Sprintf("%s.to_proto()", messageValue)
		}
	default:
		expression = messageValue
	}

	return expression
}

// SetInsideConstructorFields Builds a list of fields which are easily set in the constructor. Excluded fields are: timestamps and wrappers
func (s *PythonMessageEnumDescriptors) SetInsideConstructorFields(m google_protobuf.DescriptorProto) []*google_protobuf.FieldDescriptorProto {
	r := []*google_protobuf.FieldDescriptorProto{}

	for _, f := range m.Field {
		if !internal.IsTimestampField(f) && !internal.IsWrapperField(f) {
			r = append(r, f)
		}
	}

	return r
}

// DateTimeFields Builds a list of DateTime fields
func (s *PythonMessageEnumDescriptors) DateTimeFields(m google_protobuf.DescriptorProto) []*google_protobuf.FieldDescriptorProto {
	r := []*google_protobuf.FieldDescriptorProto{}

	for _, f := range m.Field {
		if internal.IsTimestampField(f) {
			r = append(r, f)
		}
	}

	return r
}

// WrapperFields Builds a list of wrapper fields
func (s *PythonMessageEnumDescriptors) WrapperFields(m google_protobuf.DescriptorProto) []*google_protobuf.FieldDescriptorProto {
	r := []*google_protobuf.FieldDescriptorProto{}

	for _, f := range m.Field {
		if internal.IsWrapperField(f) {
			r = append(r, f)
		}
	}

	return r
}

// WrapperFieldToPropertyType returns the vobject property type as a string for the provided wrapper field
func WrapperFieldToPropertyType(f google_protobuf.FieldDescriptorProto) string {
	switch f.GetTypeName() {
	case ".google.protobuf.DoubleValue":
		return "FloatProperty"
	case ".google.protobuf.FloatValue":
		return "FloatProperty"
	case ".google.protobuf.Int64Value":
		return "IntegerProperty"
	case ".google.protobuf.UInt64Value":
		return "IntegerProperty"
	case ".google.protobuf.Int32Value":
		return "IntegerProperty"
	case ".google.protobuf.UInt32Value":
		return "IntegerProperty"
	case ".google.protobuf.BoolValue":
		return "BooleanProperty"
	case ".google.protobuf.StringValue":
		return "StringProperty"
	case ".google.protobuf.BytesValue":
		return "StringProperty"
	default:
		panic(fmt.Sprintf("Unsupported Wrapper Field Encountered: %s", f.GetTypeName()))
	}
}