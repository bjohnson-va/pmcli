package java

import (
	"fmt"
	"os"

	"strings"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/vendasta/mscli/tools/sdkgen/internal"
)

type Tree struct {
	Name     string
	Index    int
	Parent   *Tree
	Children []*Tree
}

type ResolverInterface interface {
	GetMessageJavaName(message *descriptor.DescriptorProto) string
	GetMessageProtoName(message *descriptor.DescriptorProto) string
	GetMessageDoc(message *descriptor.DescriptorProto) string
	GetMessageByProtoName(protoName string) *descriptor.DescriptorProto
	IsMessageNested(field *descriptor.DescriptorProto) bool
	GetMessageJavaOuterClassName(message *descriptor.DescriptorProto) string

	GetFieldJavaName(field *descriptor.FieldDescriptorProto) string
	GetFieldProtoName(field *descriptor.FieldDescriptorProto) string
	GetFieldDoc(field *descriptor.FieldDescriptorProto) string
	GetFieldJavaBaseType(field *descriptor.FieldDescriptorProto) string
	GetFieldJavaType(field *descriptor.FieldDescriptorProto) string
	GetFieldJavaDefault(field *descriptor.FieldDescriptorProto) string

	GetEnumJavaName(enum *descriptor.EnumDescriptorProto) string
	GetEnumProtoName(enum *descriptor.EnumDescriptorProto) string
	GetEnumDoc(enum *descriptor.EnumDescriptorProto) string
	GetEnumByProtoName(protoName string) *descriptor.EnumDescriptorProto
	GetEnumJavaOuterClassName(enum *descriptor.EnumDescriptorProto) string

	GetEnumValueJavaName(value *descriptor.EnumValueDescriptorProto) string
	GetEnumValueProtoName(value *descriptor.EnumValueDescriptorProto) string
	GetEnumValueDoc(value *descriptor.EnumValueDescriptorProto) string
}

type EnumData struct {
	Doc                string
	Name               string
	ProtoPackageName   string
	JavaOuterClassName string
	Parent             *descriptor.DescriptorProto
}

type EnumValueData struct {
	Doc    string
	Name   string
	Parent *descriptor.EnumDescriptorProto
}

type MessageData struct {
	Doc                string
	Name               string
	ProtoPackageName   string
	JavaOuterClassName string
	Parent             *descriptor.DescriptorProto
}

type FieldData struct {
	Doc    string
	Name   string
	Parent *descriptor.DescriptorProto
}

type Resolver struct {
	javaPackageName string
	enumLookup      map[*descriptor.EnumDescriptorProto]*EnumData
	enumValueLookup map[*descriptor.EnumValueDescriptorProto]*EnumValueData
	messageLookup   map[*descriptor.DescriptorProto]*MessageData
	fieldLookup     map[*descriptor.FieldDescriptorProto]*FieldData

	enumProtoNameLookup    map[string]*descriptor.EnumDescriptorProto
	messageProtoNameLookup map[string]*descriptor.DescriptorProto

	typeMap    map[descriptor.FieldDescriptorProto_Type]string
	defaultMap map[descriptor.FieldDescriptorProto_Type]string
}

func NewResolver(req *plugin.CodeGeneratorRequest) ResolverInterface {
	rv := &Resolver{
		enumLookup:             map[*descriptor.EnumDescriptorProto]*EnumData{},
		enumValueLookup:        map[*descriptor.EnumValueDescriptorProto]*EnumValueData{},
		messageLookup:          map[*descriptor.DescriptorProto]*MessageData{},
		fieldLookup:            map[*descriptor.FieldDescriptorProto]*FieldData{},
		enumProtoNameLookup:    map[string]*descriptor.EnumDescriptorProto{},
		messageProtoNameLookup: map[string]*descriptor.DescriptorProto{},
	}

	//Initialize type lookup
	rv.typeMap = map[descriptor.FieldDescriptorProto_Type]string{
		descriptor.FieldDescriptorProto_TYPE_DOUBLE: "double",
		descriptor.FieldDescriptorProto_TYPE_FLOAT:  "float",
		descriptor.FieldDescriptorProto_TYPE_STRING: "String",
		descriptor.FieldDescriptorProto_TYPE_BOOL:   "boolean",
		descriptor.FieldDescriptorProto_TYPE_INT32:  "int",
		descriptor.FieldDescriptorProto_TYPE_SINT32: "int",
		descriptor.FieldDescriptorProto_TYPE_UINT32: "int",
		descriptor.FieldDescriptorProto_TYPE_INT64:  "long",
		descriptor.FieldDescriptorProto_TYPE_SINT64: "long",
		descriptor.FieldDescriptorProto_TYPE_UINT64: "long",
	}

	//Initialize default lookup
	rv.defaultMap = map[descriptor.FieldDescriptorProto_Type]string{
		descriptor.FieldDescriptorProto_TYPE_DOUBLE: "0",
		descriptor.FieldDescriptorProto_TYPE_FLOAT:  "0",
		descriptor.FieldDescriptorProto_TYPE_STRING: `""`,
		descriptor.FieldDescriptorProto_TYPE_BOOL:   "false",
		descriptor.FieldDescriptorProto_TYPE_INT32:  "0",
		descriptor.FieldDescriptorProto_TYPE_SINT32: "0",
		descriptor.FieldDescriptorProto_TYPE_UINT32: "0",
		descriptor.FieldDescriptorProto_TYPE_INT64:  "0",
		descriptor.FieldDescriptorProto_TYPE_SINT64: "0",
		descriptor.FieldDescriptorProto_TYPE_UINT64: "0",
	}

	for _, file := range req.ProtoFile {
		//Parse the package name
		fileJavaPackage := file.GetOptions().GetJavaPackage()
		if strings.Contains(fileJavaPackage, "com.google.") {
			//Skip google built-in protos
			fmt.Fprintf(os.Stderr, "Skipping Google proto file %s\n", file.GetName())
			continue
		} else if rv.javaPackageName == "" {
			rv.javaPackageName = fileJavaPackage
		} else if rv.javaPackageName != fileJavaPackage {
			fmt.Fprintf(os.Stderr, "Multi-pacakge proto not supported: %s & %s found\n", rv.javaPackageName, fileJavaPackage)
			os.Exit(-1)
		}

		fmt.Fprintf(os.Stderr, "Processing file %s\n", file.GetName())
		//Figure out the proto package name
		protoPackage := file.GetPackage()

		//Figure out the Java outer classname
		javaOuterClassName := file.GetOptions().GetJavaOuterClassname()
		fmt.Fprintf(os.Stderr, "Producing javaOuterClassName: %s\n", javaOuterClassName)

		//Parse root Enums
		for _, e := range file.GetEnumType() {
			rv.addEnum(e, nil, protoPackage, javaOuterClassName)
		}

		//Recurse throuth Messages
		rv.addMessages(file.GetMessageType(), nil, protoPackage, javaOuterClassName)

		//Extract comments
		sci := internal.ExtractSourceInfoTreeFromFile(file)

		//Tie Comments to the tree
		for k, v := range sci.MessageDocs {
			if _, ok := rv.messageLookup[k]; ok {
				rv.messageLookup[k].Doc = v
			} else {
				fmt.Fprintf(os.Stderr, "Missing message %s\n", k.GetName())
			}
		}
		for k, v := range sci.EnumDocs {
			if _, ok := rv.enumLookup[k]; ok {
				rv.enumLookup[k].Doc = v
			} else {
				fmt.Fprintf(os.Stderr, "Missing enum %s\n", k.GetName())
			}
		}
		for k, v := range sci.EnumValueDocs {
			if _, ok := rv.enumValueLookup[k]; ok {
				rv.enumValueLookup[k].Doc = v
			} else {
				fmt.Fprintf(os.Stderr, "Missing enum value %s\n", k.GetName())
			}
		}
		for k, v := range sci.FieldDocs {
			if _, ok := rv.fieldLookup[k]; ok {
				rv.fieldLookup[k].Doc = v
			} else {
				fmt.Fprintf(os.Stderr, "Missing field %s\n", k.GetName())
			}
		}
		// ServiceDocs   map[*descriptor.ServiceDescriptorProto]string
		// MethodDocs    map[*descriptor.MethodDescriptorProto]string
	}

	return rv
}

func (r *Resolver) addEnum(enum *descriptor.EnumDescriptorProto, parent *descriptor.DescriptorProto, protoPackageName, javaouterClassName string) {
	ed := &EnumData{
		Doc:                "",
		Name:               enum.GetName(),
		Parent:             parent,
		JavaOuterClassName: javaouterClassName,
		ProtoPackageName:   protoPackageName,
	}
	r.enumLookup[enum] = ed

	protoName := ""
	if parent == nil {
		protoName = fmt.Sprintf(".%s.%s", protoPackageName, ed.Name)
	} else {
		protoName = fmt.Sprintf(".%s.%s.%s", protoPackageName, r.GetMessageJavaName(parent), ed.Name)
	}
	r.enumProtoNameLookup[protoName] = enum

	for _, v := range enum.GetValue() {
		r.enumValueLookup[v] = &EnumValueData{
			Doc:    "",
			Name:   v.GetName(),
			Parent: enum,
		}
	}
}

func (r *Resolver) addMessages(messages []*descriptor.DescriptorProto, parent *descriptor.DescriptorProto, protoPackageName, javaouterClassName string) {
	for _, m := range messages {
		//Add the message
		r.messageLookup[m] = &MessageData{
			Doc:                "",
			Name:               m.GetName(),
			Parent:             parent,
			JavaOuterClassName: javaouterClassName,
			ProtoPackageName:   protoPackageName,
		}

		//Add the protoName
		protoName := ""
		if parent == nil {
			protoName = fmt.Sprintf(".%s.%s", protoPackageName, m.GetName())
		} else {
			protoName = fmt.Sprintf(".%s.%s.%s", protoPackageName, r.GetMessageJavaName(parent), m.GetName())
		}
		r.messageProtoNameLookup[protoName] = m

		//Add the enums
		for _, e := range m.GetEnumType() {
			r.addEnum(e, m, protoPackageName, javaouterClassName)
		}

		//Add the nested messages
		r.addMessages(m.GetNestedType(), m, protoPackageName, javaouterClassName)

		//Add the fields
		for _, f := range m.GetField() {
			r.fieldLookup[f] = &FieldData{
				Doc:    "",
				Name:   f.GetName(),
				Parent: m,
			}
		}
	}
}

func (r *Resolver) GetMessageJavaName(message *descriptor.DescriptorProto) string {
	if val, ok := r.messageLookup[message]; ok {
		if val.Parent != nil {
			return r.GetMessageJavaName(val.Parent) + "." + val.Name
		}
		return val.Name
	}
	return "(Unknown Message Type)"
}
func (r *Resolver) GetMessageProtoName(message *descriptor.DescriptorProto) string {
	if val, ok := r.messageLookup[message]; ok {
		if val.Parent != nil {
			return r.GetMessageProtoName(val.Parent) + "." + val.Name
		}
		return fmt.Sprintf("%s.%s", val.JavaOuterClassName, val.Name)
	}
	return "(Unknown Message Type)"
}
func (r *Resolver) GetMessageDoc(message *descriptor.DescriptorProto) string {
	if val, ok := r.messageLookup[message]; ok {
		return val.Doc
	}
	return "(Unknown Message Type)"
}

func (r *Resolver) GetMessageByProtoName(protoName string) *descriptor.DescriptorProto {
	if val, ok := r.messageProtoNameLookup[protoName]; ok {
		return val
	}
	fmt.Fprintf(os.Stderr, "ProtoName: %s\n", protoName)
	for k, _ := range r.enumProtoNameLookup {
		fmt.Fprintf(os.Stderr, " - %s\n", k)
	}
	os.Exit(0)
	return nil
}
func (r *Resolver) IsMessageNested(message *descriptor.DescriptorProto) bool {
	if val, ok := r.messageLookup[message]; ok {
		return val.Parent != nil
	}
	return false
}
func (r *Resolver) GetMessageJavaOuterClassName(message *descriptor.DescriptorProto) string {
	if val, ok := r.messageLookup[message]; ok {
		return val.JavaOuterClassName
	}
	return "ProtosClassName"
}

func (r *Resolver) GetFieldJavaName(field *descriptor.FieldDescriptorProto) string {
	if _, ok := r.fieldLookup[field]; ok {
		return field.GetJsonName()
	}
	return "(Unknown Field Type)"
}

func (r *Resolver) GetFieldProtoName(field *descriptor.FieldDescriptorProto) string {
	if _, ok := r.fieldLookup[field]; ok {
		return field.GetJsonName()
	}
	return "(Unknown Field Type)"
}

func (r *Resolver) GetFieldDoc(field *descriptor.FieldDescriptorProto) string {
	if val, ok := r.fieldLookup[field]; ok {
		return val.Doc
	}
	return "(Unknown Field Type)"
}

func (r *Resolver) GetFieldJavaBaseType(field *descriptor.FieldDescriptorProto) string {
	if val, ok := r.typeMap[field.GetType()]; ok {
		return val
	} else if field.GetType() == descriptor.FieldDescriptorProto_TYPE_ENUM {
		e := r.GetEnumByProtoName(field.GetTypeName())
		return r.GetEnumJavaName(e)
	} else if field.GetType() == descriptor.FieldDescriptorProto_TYPE_MESSAGE {
		if field.GetTypeName() == ".google.protobuf.Timestamp" {
			return "Date"
		} else if field.GetTypeName() == ".google.protobuf.BoolValue" {
			return "com.google.protobuf.BoolValue"
		} else if m := r.GetMessageByProtoName(field.GetTypeName()); m != nil {
			return r.GetMessageJavaName(m)
		}
	}
	return "(Unknown Field Type)"

}

func (r *Resolver) GetFieldJavaType(field *descriptor.FieldDescriptorProto) string {
	rv := r.GetFieldJavaBaseType(field)
	if field.GetLabel() == descriptor.FieldDescriptorProto_LABEL_REPEATED {
		return fmt.Sprintf("List<%s>", rv)
	}
	return rv
}
func (r *Resolver) GetFieldJavaDefault(field *descriptor.FieldDescriptorProto) string {
	if val, ok := r.defaultMap[field.GetType()]; ok {
		if field.GetLabel() == descriptor.FieldDescriptorProto_LABEL_REPEATED {
			return fmt.Sprintf("new ArrayList<%s>()", r.GetFieldJavaBaseType(field))
		}
		return val
	} else if field.GetType() == descriptor.FieldDescriptorProto_TYPE_ENUM {
		return "null"
	} else if field.GetType() == descriptor.FieldDescriptorProto_TYPE_MESSAGE {
		return "null"
	}
	return "(Unknown Field Type)"
}

func (r *Resolver) GetEnumJavaName(enum *descriptor.EnumDescriptorProto) string {
	if val, ok := r.enumLookup[enum]; ok {
		if val.Parent != nil {
			return r.GetMessageJavaName(val.Parent) + "." + val.Name
		}
		return val.Name
	}
	return "(Unknown Enum Type)"
}
func (r *Resolver) GetEnumProtoName(enum *descriptor.EnumDescriptorProto) string {
	if val, ok := r.enumLookup[enum]; ok {
		if val.Parent != nil {
			return r.GetMessageProtoName(val.Parent) + "." + val.Name
		}
		return fmt.Sprintf("%s.%s", val.JavaOuterClassName, val.Name)
	}
	return "(Unknown Enum Type)"
}
func (r *Resolver) GetEnumDoc(enum *descriptor.EnumDescriptorProto) string {
	if val, ok := r.enumLookup[enum]; ok {
		return val.Doc
	}
	return "(Unknown Enum Type)"
}

func (r *Resolver) GetEnumByProtoName(protoName string) *descriptor.EnumDescriptorProto {
	if val, ok := r.enumProtoNameLookup[protoName]; ok {
		return val
	}
	fmt.Fprintf(os.Stderr, "ProtoName: %s\n", protoName)
	for k, _ := range r.enumProtoNameLookup {
		fmt.Fprintf(os.Stderr, " - %s\n", k)
	}
	os.Exit(0)
	return nil
}
func (r *Resolver) GetEnumJavaOuterClassName(enum *descriptor.EnumDescriptorProto) string {
	if val, ok := r.enumLookup[enum]; ok {
		return val.JavaOuterClassName
	}
	return "ProtoClassName"
}

func (r *Resolver) GetEnumValueJavaName(value *descriptor.EnumValueDescriptorProto) string {
	if val, ok := r.enumValueLookup[value]; ok {
		return val.Name
	}
	return "(Unknown EnumValue Type)"
}
func (r *Resolver) GetEnumValueProtoName(value *descriptor.EnumValueDescriptorProto) string {
	if val, ok := r.enumValueLookup[value]; ok {
		return val.Name
	}
	return "(Unknown EnumValue Type)"
}
func (r *Resolver) GetEnumValueDoc(value *descriptor.EnumValueDescriptorProto) string {
	if val, ok := r.enumValueLookup[value]; ok {
		return val.Doc
	}
	return "(Unknown EnumValue Type)"
}
