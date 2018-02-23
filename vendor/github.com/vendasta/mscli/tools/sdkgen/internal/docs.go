package internal

import (
	"fmt"
	"os"
	"strings"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"
)

type SourceInfoTree struct {
	FileDocs      map[*descriptor.FileDescriptorProto]string
	MessageDocs   map[*descriptor.DescriptorProto]string
	EnumDocs      map[*descriptor.EnumDescriptorProto]string
	EnumValueDocs map[*descriptor.EnumValueDescriptorProto]string
	FieldDocs     map[*descriptor.FieldDescriptorProto]string
	ServiceDocs   map[*descriptor.ServiceDescriptorProto]string
	MethodDocs    map[*descriptor.MethodDescriptorProto]string
}

// ExtractSourceInfoTreeFromFiles extracts a SourceInfoTree from the files' source code info locations
func ExtractSourceInfoTreeFromFile(file *descriptor.FileDescriptorProto) *SourceInfoTree {
	sci := newSourceInfoTree()
	for _, loc := range file.GetSourceCodeInfo().GetLocation() {
		sci.PushLocation(file, loc)
	}
	return &sci
}

func newSourceInfoTree() SourceInfoTree {
	return SourceInfoTree{
		FileDocs:      map[*descriptor.FileDescriptorProto]string{},
		MessageDocs:   map[*descriptor.DescriptorProto]string{},
		EnumDocs:      map[*descriptor.EnumDescriptorProto]string{},
		EnumValueDocs: map[*descriptor.EnumValueDescriptorProto]string{},
		FieldDocs:     map[*descriptor.FieldDescriptorProto]string{},
		ServiceDocs:   map[*descriptor.ServiceDescriptorProto]string{},
		MethodDocs:    map[*descriptor.MethodDescriptorProto]string{},
	}
}

func (m *SourceInfoTree) PushEnum(enum *descriptor.EnumDescriptorProto, location *descriptor.SourceCodeInfo_Location, depth int) {
	subPath := location.Path[depth:]
	if len(subPath) == 0 {
		m.EnumDocs[enum] = strings.Trim(location.GetLeadingComments(), "\n \t")
		return
	}

	if subPath[0] == 1 { // Name?
		return
	} else if subPath[0] == 2 {
		valueType := enum.GetValue()[subPath[1]]
		if len(subPath) == 2 {
			m.EnumValueDocs[valueType] = strings.Trim(location.GetLeadingComments(), "\n \t")
			return
		}

		if subPath[2] == 1 { // Name?
			return
		} else if subPath[2] == 2 { //Value?
			return
		} else if subPath[2] == 3 { // Option?
			return
		}

		fmt.Fprintf(os.Stderr, "%s.%s Unknown Value Comment: %v\n", enum.GetName(), valueType.GetName(), subPath)
		// return
	} else if subPath[0] == 3 { // Options?
		return
	}

	fmt.Fprintf(os.Stderr, "%s Contained unknown Enum Type Docstring: %v\n", enum.GetName(), subPath)
	return
}

func (m *SourceInfoTree) PushType(message *descriptor.DescriptorProto, location *descriptor.SourceCodeInfo_Location, depth int) {
	subPath := location.Path[depth:]

	if len(subPath) == 0 { //Documentation of the message type itself
		m.MessageDocs[message] = strings.Trim(location.GetLeadingComments(), "\n \t")
		return
	}
	if subPath[0] == 1 { // Name?
		return
	} else if subPath[0] == 2 { // Fields
		fieldType := message.GetField()[subPath[1]]
		if len(subPath) == 2 {
			m.FieldDocs[fieldType] = strings.Trim(location.GetLeadingComments(), "\n \t")
			return
		}
		if subPath[2] == 1 { // Name?
			return
		} else if subPath[2] == 3 { // Number?
			return
		} else if subPath[2] == 4 { // Label?
			return
		} else if subPath[2] == 5 { // Type?
			return
		} else if subPath[2] == 6 { // TypeName?
			return
		}
		fmt.Fprintf(os.Stderr, "%s.%s Contained Unknown Field Docstring: %v\n", message.GetName(), fieldType.GetName(), subPath[2:])
		return
	} else if location.Path[2] == 3 { // Nested Type
		nestedType := message.GetNestedType()[subPath[1]]
		m.PushType(nestedType, location, depth+2)
		return
	} else if location.Path[2] == 4 { // Enum Type
		enumType := message.GetEnumType()[subPath[1]]
		m.PushEnum(enumType, location, depth+2)
		return
	}
	fmt.Fprintf(os.Stderr, "%s Contained Unknown Message Docstring: %v\n", message.GetName(), subPath)
}

func (m *SourceInfoTree) PushLocation(file *descriptor.FileDescriptorProto, location *descriptor.SourceCodeInfo_Location) {
	//Parse the file location
	if len(location.Path) == 0 { //File documentation
		m.FileDocs[file] = strings.Trim(location.GetLeadingComments(), "\n \t")
		return
	}
	if location.Path[0] == 2 { //Package
		return
	} else if location.Path[0] == 3 { //Dependency
		return
	} else if location.Path[0] == 4 { //MessageType
		messageType := file.GetMessageType()[location.Path[1]]
		m.PushType(messageType, location, 2)
		return
	} else if location.Path[0] == 5 {
		enumType := file.GetEnumType()[location.Path[1]]
		m.PushEnum(enumType, location, 2)
		return
	} else if location.Path[0] == 6 {
		serviceType := file.GetService()[location.Path[1]]
		if len(location.Path) == 2 {
			m.ServiceDocs[serviceType] = strings.Trim(location.GetLeadingComments(), "\n \t")
			return
		}
		if location.Path[2] == 1 { // Name?
			return
		} else if location.Path[2] == 2 {
			methodType := serviceType.GetMethod()[location.Path[3]]
			m.MethodDocs[methodType] = strings.Trim(location.GetLeadingComments(), "\n \t")
			return
		}
		fmt.Fprintf(os.Stderr, "%s Contained Unknown Service Docstring: %v\n", serviceType.GetName(), location.Path[2:])
		return
	} else if location.Path[0] == 8 { //Options
		return
	} else if location.Path[0] == 12 { //Syntax (proto3 vs proto2)
		return
	}
	fmt.Fprintf(os.Stderr, "Unknown Location: %v\n", location.Path)
}
