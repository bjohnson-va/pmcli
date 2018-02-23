package internal

import (
	google_protobuf "github.com/golang/protobuf/protoc-gen-go/descriptor"

	"fmt"
)

type Message struct {
	Message *google_protobuf.DescriptorProto
	Parent  *Message
}

// FullyQualifiedName returns the qualified name containing the parent proto references
func (m *Message) FullyQualifiedName() string {
	str := *m.Message.Name
	parent := m.Parent
	for parent != nil {
		str = fmt.Sprintf("%s.%s", *parent.Message.Name, str)
		parent = parent.Parent
	}
	return str
}

// InterfaceName returns the name this message will have as an interface
func (m *Message) InterfaceName() string {
	return fmt.Sprintf("%sInterface", *m.Message.Name)
}

// IsMapEntry returns if the message was an automatically generated entry message for a proto map
func (m *Message) IsMapEntry() bool {
	return m.Message.GetOptions().GetMapEntry()
}
