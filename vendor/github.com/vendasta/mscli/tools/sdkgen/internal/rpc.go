package internal

import (
	"strings"
	"fmt"

	google_protobuf "github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/vendasta/mscli/tools/sdkgen/util"
)

// RPC remote procedure call structure
type RPC struct {
	Name string
	Path string

	Response string
	Request  string

	ResponseDescriptor *google_protobuf.DescriptorProto
	RequestDescriptor  *google_protobuf.DescriptorProto
}

func (r *RPC) TsName() string {
	return util.ToCamelCase(util.ToSnakeCase(r.Name))
}

// ResponseIsEmpty returns whether or not the response type is empty
func (r *RPC) ResponseIsEmpty() bool {
	pieces := strings.Split(r.Response, ".")
	name := pieces[len(pieces)-1]

	return strings.ToLower(name) == "empty"
}

// ResponseClassName The response class name
func (r *RPC) ResponseClassName() string {
	if r.ResponseIsEmpty() {
		return "google_dot_protobuf_dot_empty__pb2.Empty"
	}
	pieces := strings.Split(r.Response, ".")
	return pieces[len(pieces)-1]
}

// RequestClassName the request class name
func (r *RPC) RequestClassName() string {
	pieces := strings.Split(r.Request, ".")
	return pieces[len(pieces)-1]
}

// Args rpc arguments
func (r *RPC) Args() string {
	return strings.Join(r.ArgsList(), ", ")
}

// ArgsList rpc arguments as a list
func (r *RPC) ArgsList() []string {
	fields := []string{}
	if r.RequestDescriptor == nil {
		return fields
	}
	for _, f := range r.RequestDescriptor.Field {
		fields = append(fields, *f.Name)
	}
	return fields
}

// RequestConstructorArgs rpc request constructor args
func (r *RPC) RequestConstructorArgs() string {
	fields := []string{}
	if r.RequestDescriptor == nil {
		return ""
	}
	for _, f := range r.RequestDescriptor.Field {
		fields = append(fields, fmt.Sprintf("%s=%s", *f.Name, *f.Name))
	}
	return strings.Join(fields, ", ")
}
