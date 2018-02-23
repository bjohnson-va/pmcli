package internal

import (
	"strings"

	google_protobuf "github.com/golang/protobuf/protoc-gen-go/descriptor"
)

// MessageEnumDescriptors structure of message and enum descripors
type MessageEnumDescriptors struct {
	FileName     string
	Dependencies []string
	Messages     []*Message
	Enums        []*google_protobuf.EnumDescriptorProto
}

// NonGoogleDependenciesFileNames Builds a list of non google dependency file names
func (s *MessageEnumDescriptors) NonGoogleDependenciesFileNames(dependencies []string) []string {
	deps := []string{}

	for _, d := range dependencies {
		if !IsGoogleProtoFileName(d) {
			deps = append(deps, s.GetFileNameFromFullPath(d))
		}
	}

	return deps
}

// GetFileNameFromFullPath filename of a python file
func (s *MessageEnumDescriptors) GetFileNameFromFullPath(str string) string {
	fileWithNoExtension := strings.Split(str, ".")[0]
	pathParts := strings.Split(fileWithNoExtension, "/")
	fileWithNoPath := pathParts[len(pathParts)-1]
	return fileWithNoPath
}

// ExtractTypeName returns the type name from the type reference
func (s *MessageEnumDescriptors) ExtractTypeName(typeRef string) string {
	parts := strings.Split(typeRef, ".")
	return parts[len(parts)-1]
}

// GetMessageByName returns the message with the name
func (s *MessageEnumDescriptors) GetMessageByName(name string) *Message {
	for _, m := range s.Messages {
		if m.Message.GetName() == name {
			return m
		}
	}
	return nil
}

// ExtractMessagesAndEnums extracts all messages and enums from the proto
func ExtractMessagesAndEnums(p *google_protobuf.FileDescriptorProto) ([]*Message, []*google_protobuf.EnumDescriptorProto) {
	mr := []*Message{}
	er := []*google_protobuf.EnumDescriptorProto{}
	for _, m := range p.MessageType {
		mr, er = recExtractMessagesAndEnums(m, nil, mr, er)
	}
	for _, e := range p.EnumType {
		er = append(er, e)
	}
	return mr, er
}

func recExtractMessagesAndEnums(m *google_protobuf.DescriptorProto, p *Message, mr []*Message, er []*google_protobuf.EnumDescriptorProto) ([]*Message, []*google_protobuf.EnumDescriptorProto) {
	curMessage := &Message{Message: m, Parent: p}
	for _, nt := range m.NestedType {
		mr, er = recExtractMessagesAndEnums(nt, curMessage, mr, er)
	}
	mr = append(mr, curMessage)

	for _, et := range m.EnumType {
		er = append(er, et)
	}

	return mr, er
}


// IsWrapperField tests if the provided field is one of the wrapper types from well-known types
func IsWrapperField(f *google_protobuf.FieldDescriptorProto) bool {
	if *f.Type == google_protobuf.FieldDescriptorProto_TYPE_MESSAGE {
		wrapperTypes := []string{
			".google.protobuf.DoubleValue",
			".google.protobuf.FloatValue",
			".google.protobuf.Int64Value",
			".google.protobuf.UInt64Value",
			".google.protobuf.Int32Value",
			".google.protobuf.UInt32Value",
			".google.protobuf.BoolValue",
			".google.protobuf.StringValue",
			".google.protobuf.BytesValue",
		}
		for _, t := range wrapperTypes {
			if t == f.GetTypeName() {
				return true
			}
		}
	}
	return false
}

// IsTimestampField tests if the provided field is a timestamp
func IsTimestampField(f *google_protobuf.FieldDescriptorProto) bool {
	return *f.Type == google_protobuf.FieldDescriptorProto_TYPE_MESSAGE && f.GetTypeName() == ".google.protobuf.Timestamp"
}
