package parse

import (
	"fmt"
	"strings"
	"github.com/bjohnson-va/pmcli/protofiles"
	"github.com/emicklei/proto"
)

var (
	skippableProtos = []string{
		"google/protobuf/timestamp.proto",
		"google/protobuf/empty.proto",
	}
)

type FieldTypes struct {
	Messages []proto.Message
	Enums []proto.Enum
}

// AllFieldTypesFromProtos reads the given proto files (and all files it imports)
// and returns a collection of all Messages and Enums that were defined.
func AllFieldTypesFromProtos(rootPath string, definition *proto.Proto) (*FieldTypes, error) {
	messages := Messages(definition.Elements)
	enums := Enums(definition.Elements)
	otherFiles := Imports(definition.Elements)
	for _, o := range otherFiles {
		if isSkippableProto(o) {
			continue;
		}
		d, err := protofiles.Read(fmt.Sprintf("%s/%s", rootPath, o.Filename))
		if err != nil {
			// Don't die.  Attempt to process the rest of the files.
			continue
		}
		t, err := AllFieldTypesFromProtos(rootPath, d)
		if err != nil {
			err := fmt.Errorf("couldn't extract messages/enums: %s", err.Error())
			return nil, err
		}
		messages = append(messages, t.Messages...)
		enums = append(enums, t.Enums...)
	}
	// Nested messages
	for _, m := range messages {
		messages = append(messages, Messages(m.Elements)...)
		enums = append(enums, Enums(m.Elements)...)
	}
	return &FieldTypes{
		Messages: messages,
		Enums: enums,
	}, nil
}

func isSkippableProto(filename proto.Import) bool {
	for _, s := range skippableProtos {
		if strings.HasSuffix(filename.Filename, s) {
			return true
		}
	}
	return false
}
