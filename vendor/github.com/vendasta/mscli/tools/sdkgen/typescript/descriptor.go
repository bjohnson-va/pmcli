package typescript

import (
	"fmt"

	google_protobuf "github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/vendasta/mscli/tools/sdkgen/internal"
	"github.com/vendasta/mscli/tools/sdkgen/util"
)

// TypescriptMessageEnumDescriptors structure of message and enum descripors
type TypescriptMessageEnumDescriptors struct {
	*internal.MessageEnumDescriptors
}

// TsFileName filename of the descriptor
func (s *TypescriptMessageEnumDescriptors) TsFileName() string {
	fileName := s.GetFileNameFromFullPath(s.FileName)
	return fmt.Sprintf("%s", util.ToKebabCase(fileName))
}

// TsInterfaceFileName interface filename for the descriptor
func (s *TypescriptMessageEnumDescriptors) TsInterfaceFileName() string {
	fileName := s.GetFileNameFromFullPath(s.FileName)
	return fmt.Sprintf("%s.interface", util.ToKebabCase(fileName))
}

// TsEnumFileName enum file name for the descriptor
func (s *TypescriptMessageEnumDescriptors) TsEnumFileName() string {
	fileName := s.GetFileNameFromFullPath(s.FileName)
	return fmt.Sprintf("%s.enum", util.ToKebabCase(fileName))
}

// FieldToProperty converts a proto Field to a typescript property
func (s *TypescriptMessageEnumDescriptors) FieldToProperty(field google_protobuf.FieldDescriptorProto, isInterface bool) string {
	property := ""
	switch *field.Type {
	case google_protobuf.FieldDescriptorProto_TYPE_DOUBLE:
		property = "number"
	case google_protobuf.FieldDescriptorProto_TYPE_FLOAT:
		property = "number"
	case google_protobuf.FieldDescriptorProto_TYPE_INT64:
		property = "number"
	case google_protobuf.FieldDescriptorProto_TYPE_UINT64:
		property = "number"
	case google_protobuf.FieldDescriptorProto_TYPE_INT32:
		property = "number"
	case google_protobuf.FieldDescriptorProto_TYPE_FIXED64:
		property = "number"
	case google_protobuf.FieldDescriptorProto_TYPE_FIXED32:
		property = "number"
	case google_protobuf.FieldDescriptorProto_TYPE_BOOL:
		property = "boolean"
	case google_protobuf.FieldDescriptorProto_TYPE_STRING:
		property = "string"
	case google_protobuf.FieldDescriptorProto_TYPE_MESSAGE:
		m := s.GetMessageByName(s.ExtractTypeName(field.GetTypeName()))
		if m != nil && m.IsMapEntry() {
			mapKeyField := m.Message.GetField()[0]
			if mapKeyField.GetType() != google_protobuf.FieldDescriptorProto_TYPE_STRING {
				panic(fmt.Sprintf("Maps keyed by types other than strings is not currently supported in typescript sdkgen: %s", m.Message.GetName()))
			}
			mapValueField := m.Message.GetField()[1]
			mapValueProperty := s.FieldToProperty(*mapValueField, isInterface)
			return fmt.Sprintf("{[key: string]: %s}", mapValueProperty)
		}
		if internal.IsWrapperField(&field) {
			property = WrapperFieldToPropertyType(field)
		} else if internal.IsTimestampField(&field) {
			property = "Date"
		} else {
			if isInterface {
				property = "i." + s.ExtractInterfaceTypeName(*field.TypeName)
			} else {
				property = "o." + s.ExtractTypeName(*field.TypeName)
			}
		}
	case google_protobuf.FieldDescriptorProto_TYPE_BYTES:
		property = "string"
	case google_protobuf.FieldDescriptorProto_TYPE_UINT32:
		property = "number"
	case google_protobuf.FieldDescriptorProto_TYPE_ENUM:
		property = "e." + s.ExtractTypeName(*field.TypeName)
	case google_protobuf.FieldDescriptorProto_TYPE_SFIXED32:
		property = "number"
	case google_protobuf.FieldDescriptorProto_TYPE_SFIXED64:
		property = "number"
	case google_protobuf.FieldDescriptorProto_TYPE_SINT32:
		property = "number"
	case google_protobuf.FieldDescriptorProto_TYPE_SINT64:
		property = "number"
	}

	switch *field.Label {
	case google_protobuf.FieldDescriptorProto_LABEL_REPEATED:
		property += "[]"
	}

	return property
}

// FieldFromProto contains special rules when converting a proto to an object
func (s *TypescriptMessageEnumDescriptors) FieldFromProto(field google_protobuf.FieldDescriptorProto) string {
	if internal.IsWrapperField(&field) && !isInt64Type(&field) {
		// Nothing to add, wrapper fields behave like scalars over json http
		return ""
	}

	if field.GetType() != google_protobuf.FieldDescriptorProto_TYPE_MESSAGE &&
		field.GetType() != google_protobuf.FieldDescriptorProto_TYPE_ENUM &&
		!isInt64Type(&field) {
		// No special handling, they will just be set directly
		return ""
	}

	extra := fmt.Sprintf("if (proto.%s) {m.%s = ", field.GetJsonName(), field.GetJsonName())
	extra += s.FieldValueFromProtoValue(field, fmt.Sprintf("proto.%s", field.GetJsonName()))
	extra += ";}"
	return extra
}

// FieldValueFromProtoValue returns the field value expression from the proto value expression
func (s *TypescriptMessageEnumDescriptors) FieldValueFromProtoValue(field google_protobuf.FieldDescriptorProto, protoValueExpression string) string {
	if internal.IsWrapperField(&field) && !isInt64Type(&field) {
		// Nothing to add, wrapper fields behave like scalars over json http
		return protoValueExpression
	}

	exp := protoValueExpression

	if isInt64Type(&field) {
		if *field.Label == google_protobuf.FieldDescriptorProto_LABEL_REPEATED {
			exp = fmt.Sprintf("%s.map((i: any) => parseInt(i, 10))", protoValueExpression)
		} else {
			exp = fmt.Sprintf("parseInt(%s, 10)", protoValueExpression)
		}
		return exp
	}

	switch *field.Type {
	case google_protobuf.FieldDescriptorProto_TYPE_MESSAGE:
		m := s.GetMessageByName(s.ExtractTypeName(field.GetTypeName()))
		if m != nil && m.IsMapEntry() {
			mapValueField := m.Message.GetField()[1]
			exp = fmt.Sprintf(`Object.keys(%s).reduce((obj, k) => { obj[k] = %s; return obj; }, {})`, protoValueExpression, s.FieldValueFromProtoValue(*mapValueField, protoValueExpression+"[k]"))
			break
		}
		if internal.IsTimestampField(&field) {
			if *field.Label == google_protobuf.FieldDescriptorProto_LABEL_REPEATED {
				exp = fmt.Sprintf("%s.map((d: any) => new Date(d))", protoValueExpression)
			} else {
				exp = fmt.Sprintf("new Date(%s)", protoValueExpression)
			}
		} else {
			if *field.Label == google_protobuf.FieldDescriptorProto_LABEL_REPEATED {
				exp = fmt.Sprintf("%s.map(o.%s.fromProto)", protoValueExpression, s.ExtractTypeName(*field.TypeName))
			} else {
				exp = fmt.Sprintf("o.%s.fromProto(%s)", s.ExtractTypeName(*field.TypeName), protoValueExpression)
			}
		}
	case google_protobuf.FieldDescriptorProto_TYPE_ENUM:
		typeName := s.ExtractTypeName(*field.TypeName)
		if *field.Label == google_protobuf.FieldDescriptorProto_LABEL_REPEATED {
			exp = fmt.Sprintf("%s.map((v: string) => enumStringToValue<e.%s>(e.%s, v))", protoValueExpression, typeName, typeName)
		} else {
			exp = fmt.Sprintf("enumStringToValue<e.%s>(e.%s, %s)", typeName, typeName, protoValueExpression)
		}
	}
	return exp
}

// FieldToProtoValue converts a domain Field to it's proto value expression
func (s *TypescriptMessageEnumDescriptors) FieldToProtoValue(field google_protobuf.FieldDescriptorProto) string {
	messageValue := fmt.Sprintf("this.%s", *field.JsonName)
	exp := fmt.Sprintf("'%s': (typeof %s !== 'undefined'", *field.JsonName, messageValue)
	switch *field.Type {
	case google_protobuf.FieldDescriptorProto_TYPE_MESSAGE:
		exp += fmt.Sprintf(" && %s !== null) ? ", messageValue)
	default:
		exp += fmt.Sprintf(") ? ")
	}
	exp += s.FieldValueToProtoValue(field, messageValue)

	return exp + " : null"
}

// FieldValueToProtoValue converts a domain value expression to it's proto value expression
func (s *TypescriptMessageEnumDescriptors) FieldValueToProtoValue(field google_protobuf.FieldDescriptorProto, fieldValueExpression string) string {
	if internal.IsWrapperField(&field) {
		return fieldValueExpression
	}

	exp := fieldValueExpression
	switch *field.Type {
	case google_protobuf.FieldDescriptorProto_TYPE_MESSAGE:
		m := s.GetMessageByName(s.ExtractTypeName(field.GetTypeName()))
		if m != nil && m.IsMapEntry() {
			mapValueField := m.Message.GetField()[1]
			exp = fmt.Sprintf(`Object.keys(%s).reduce((obj, k) => { obj[k] = %s; return obj; }, {})`, fieldValueExpression, s.FieldValueToProtoValue(*mapValueField, fieldValueExpression+"[k]"))
			break
		}
		if *field.Label == google_protobuf.FieldDescriptorProto_LABEL_REPEATED {
			if internal.IsTimestampField(&field) {
				exp = fmt.Sprintf("%s.map(t => t.toISOString())", fieldValueExpression)
			} else {
				exp = fmt.Sprintf("%s.map(obj => obj.toApiJson())", fieldValueExpression)
			}
		} else {
			if internal.IsTimestampField(&field) {
				exp = fmt.Sprintf("%s.toISOString()", fieldValueExpression)
			} else {
				exp = fmt.Sprintf("%s.toApiJson()", fieldValueExpression)
			}
		}
	}

	return exp
}

// ExtractInterfaceTypeName returns the interface type name from the type reference
func (s *TypescriptMessageEnumDescriptors) ExtractInterfaceTypeName(typeRef string) string {
	return s.ExtractTypeName(typeRef) + "Interface"
}

// WrapperFieldToPropertyType returns the vobject property type as a string for the provided wrapper field
func WrapperFieldToPropertyType(f google_protobuf.FieldDescriptorProto) string {
	switch f.GetTypeName() {
	case ".google.protobuf.DoubleValue":
		return "number"
	case ".google.protobuf.FloatValue":
		return "number"
	case ".google.protobuf.Int64Value":
		return "number"
	case ".google.protobuf.UInt64Value":
		return "number"
	case ".google.protobuf.Int32Value":
		return "number"
	case ".google.protobuf.UInt32Value":
		return "number"
	case ".google.protobuf.BoolValue":
		return "boolean"
	case ".google.protobuf.StringValue":
		return "string"
	case ".google.protobuf.BytesValue":
		return "string"
	default:
		panic(fmt.Sprintf("Unsupported Wrapper Field Encountered: %s", f.GetTypeName()))
	}
}

func isInt64Type(f *google_protobuf.FieldDescriptorProto) bool {
	if internal.IsWrapperField(f) {
		return f.GetTypeName() == ".google.protobuf.Int64Value" || f.GetTypeName() == ".google.protobuf.UInt64Value"
	} else {
		return *f.Type == google_protobuf.FieldDescriptorProto_TYPE_INT64 ||
			*f.Type == google_protobuf.FieldDescriptorProto_TYPE_UINT64 ||
			*f.Type == google_protobuf.FieldDescriptorProto_TYPE_FIXED64
	}
}