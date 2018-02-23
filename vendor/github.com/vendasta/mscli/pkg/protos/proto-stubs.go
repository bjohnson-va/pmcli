package protos

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/vendasta/mscli/pkg/grpc"
	"github.com/vendasta/mscli/pkg/spec"
)

const (
	//TODO generate the stub client???????
	stubServerHeaderTemplate = `package api

import (
%s
)
`
)

//GenerateStubImplementations produces a stub implmentation for the user to be able to
//get going quickly.  It also updates the server/main.go to include the stub
//implementation and serve it.
func GenerateStubImplementations(spec spec.MicroserviceFile, protoData []byte, style StubStyle) error {
	protoSet, err := NewProtoSetFromBytes(protoData, style)
	if err != nil {
		return fmt.Errorf("error could not extract the Proto data: %s", err.Error())
	}

	headers := protoSet.GetHeaders()
	headerText := ""
	for _, h := range headers {
		headerText += fmt.Sprintf("    %s\n", h)
	}
	fileContent := fmt.Sprintf(stubServerHeaderTemplate, headerText)
	fileContent += "\n"
	fileContent += protoSet.getStubs()

	destination := fmt.Sprintf("%s/servers.go", spec.Microservice.GetPathToAPIDir())
	if _, err := os.Stat(destination); os.IsNotExist(err) {
		if err := ioutil.WriteFile(destination, []byte(fileContent), 0644); err != nil {
			return fmt.Errorf("error writing %s: %s", destination, err.Error())
		}
	} else if spec.Microservice.Debug {
		fmt.Printf("%s already exists, not generating the stub", destination)
	}

	//Update the server to have the right servers and imports
	return augmentMainGoToRegisterServices(protoSet.services, spec)
}

func insertRegisteringOfServers(mainGoText string, services map[string]serviceInfo, spec spec.MicroserviceFile) string {
	if spec.Microservice.Debug {
		fmt.Println("Registering Servers")
	}
	registerString := ""
	for _, s := range services {
		regString := fmt.Sprintf("%s.Register%sServer(grpcServer, &api.%sServer{})", s.GoNamespace, s.ServiceName, s.ServiceName)
		if !strings.Contains(mainGoText, regString) {
			registerString = registerString + regString + "\n    "
		} else if spec.Microservice.Debug {
			fmt.Printf("Registration for %s already exists, skipping", s.ServiceName)
		}
	}
	return strings.Replace(mainGoText, grpc.GRPCServerRegisterTag, registerString+grpc.GRPCServerRegisterTag, 1)
}

func insertProtobufImport(mainGoText string, spec spec.MicroserviceFile) string {
	//All of the proto imports (not the internal/api import) should really be built from the descriptor.pb (probably the services in them)
	//Instead of building them from the listed proto paths
	for _, p := range spec.Microservice.GetProtoImportPaths() {
		if !strings.Contains(mainGoText, p) {
			packageImport := fmt.Sprintf("import ( \n    \"%s\"    ", p)
			mainGoText = strings.Replace(mainGoText, "import (", packageImport, 1)
		} else if spec.Microservice.Debug {
			fmt.Printf("main.go already imports %s, skipping", spec.Microservice.GoPackageName)
		}
	}

	if !strings.Contains(mainGoText, spec.Microservice.GoPackageName) {
		packageImport := fmt.Sprintf("import ( \n    \"%s/internal/api\"", spec.Microservice.GoPackageName)
		mainGoText = strings.Replace(mainGoText, "import (", packageImport, 1)
	}

	return mainGoText
}

func augmentMainGoToRegisterServices(services map[string]serviceInfo, spec spec.MicroserviceFile) error {
	pathToMainGo := spec.Microservice.GetPathToMainGo()
	b, err := ioutil.ReadFile(pathToMainGo)
	if err != nil {
		return fmt.Errorf("error reading %s", pathToMainGo)
	}

	augmentedMainGo := insertRegisteringOfServers(string(b[:]), services, spec)
	augmentedMainGo = insertProtobufImport(augmentedMainGo, spec)

	if err := ioutil.WriteFile(pathToMainGo, []byte(augmentedMainGo), 0644); err != nil {
		return fmt.Errorf("error writing %s", pathToMainGo)
	}
	return nil
}

//methodInfo stores handy information about a method for code gen
type methodInfo struct {
	MethodOffset    int
	MethodName      string
	RequestType     string
	ResponseType    string
	ServerStreaming bool
	ClientStreaming bool
}

//serviceInfo stores handy information about a server for code gen
type serviceInfo struct {
	FileOffset    int
	FileName      string
	ServiceOffset int
	ServiceName   string
	GoStructName  string
	Methods       map[string]methodInfo
	GoNamespace   string
}

//typeInfo stores handy information about a protobuf type for code gen
type typeInfo struct {
	FileOffset    int
	FileName      string
	ProtoTypeName string
	GoTypeName    string
	GoNamespace   string
	PackageName   string
	Referenced    bool
	GoImport      string
}

type StubWriter interface {
	GetHeaders() []string
	GetStubs() string
}

//ProtoSet provides methods for generating stubs from a Protobuf descriptor
type ProtoSet struct {
	descriptorSet *descriptor.FileDescriptorSet
	services      map[string]serviceInfo
	types         map[string]typeInfo
	writer        StubWriter
}

//NewProtoSetFromBytes creates a ProtoSet from the raw bytes exported from
//  protoc --descriptor_set_out=<file>
func NewProtoSetFromBytes(data []byte, style StubStyle) (*ProtoSet, error) {
	//Extract the proto
	fileSet := descriptor.FileDescriptorSet{}
	err := proto.Unmarshal(data, &fileSet)
	if err != nil {
		return nil, fmt.Errorf("error could not unmarshal the FileDescriptorSet")
	}
	return NewProtoSetFromProto(fileSet, style)
}

//NewProtoSetFromProto creates a ProtoSet from a populated FileDescriptorSet
func NewProtoSetFromProto(fds descriptor.FileDescriptorSet, style StubStyle) (*ProtoSet, error) {
	services := extractServiceInformation(fds)
	types := extractTypeInformation(fds)
	var writer StubWriter
	switch style {
	case StubEmpty:
		writer = EmptyStubs{services, types}
	case StubHandler:
		writer = HandlerStubs{services, types, EmptyStubs{services, types}}
	}
	ps := ProtoSet{
		descriptorSet: &fds,
		services:      services,
		types:         types,
		writer:        writer,
	}

	return &ps, nil
}

func extractTypeInformation(fds descriptor.FileDescriptorSet) map[string]typeInfo {
	types := map[string]typeInfo{}
	//Put the built-in google protos we know about (gross)
	emptyProto := typeInfo{
		GoImport:      "github.com/golang/protobuf/ptypes/empty",
		ProtoTypeName: ".google.protobuf.Empty",
		GoNamespace:   "pb1",
		GoTypeName:    "Empty",
	}
	timestampProto := typeInfo{
		GoImport:      "github.com/golang/protobuf/ptypes/timestamp",
		ProtoTypeName: ".google.protobuf.Timestamp",
		GoNamespace:   "pb2",
		GoTypeName:    "Timestamp",
	}
	anyProto := typeInfo{
		GoImport:      "github.com/golang/protobuf/ptypes/any",
		ProtoTypeName: ".google.protobuf.Any",
		GoNamespace:   "pb3",
		GoTypeName:    "Any",
	}

	exclude := map[string]bool{"google/protobuf/descriptor.proto": true, "google/api/http.proto": true}
	for fi, f := range fds.File {
		if exclude[f.GetName()] {
			continue
		}

		//Extract the custom messages defined in our files
		for _, msg := range f.MessageType {
			name := fmt.Sprintf(".%s.%s", f.GetPackage(), msg.GetName())
			types[name] = typeInfo{
				FileName:      f.GetName(),
				FileOffset:    fi,
				ProtoTypeName: name,
				PackageName:   f.GetPackage(),
				GoNamespace:   strings.Replace(f.GetPackage(), ".", "_", -1), //golang package names use _ instead of .
				GoTypeName:    msg.GetName(),
				Referenced:    true,
				GoImport:      "github.com/vendasta/gosdks/pb/" + strings.Replace(f.GetPackage(), ".", "/", -1),
			}
		}
		for _, d := range f.Dependency {
			switch d {
			case "google/protobuf/empty.proto":
				emptyProto.Referenced = true
			case "google/protobuf/timestamp.proto":
				timestampProto.Referenced = true
			case "google/protobuf/any.proto":
				anyProto.Referenced = true
			}
		}
	}

	types[".google.protobuf.Empty"] = emptyProto
	types[".google.protobuf.Timestamp"] = timestampProto
	types[".google.protobuf.Any"] = anyProto

	return types
}

func extractServiceInformation(fds descriptor.FileDescriptorSet) map[string]serviceInfo {
	services := map[string]serviceInfo{}
	for fi, f := range fds.File {
		for si, s := range f.Service {
			methods := map[string]methodInfo{}
			for mi, m := range s.Method {
				methods[m.GetName()] = methodInfo{
					MethodName:      m.GetName(),
					MethodOffset:    mi,
					RequestType:     m.GetInputType(),
					ResponseType:    m.GetOutputType(),
					ServerStreaming: m.GetServerStreaming(),
					ClientStreaming: m.GetClientStreaming(),
				}
			}
			services[s.GetName()] = serviceInfo{
				FileOffset:    fi,
				ServiceOffset: si,
				FileName:      f.GetName(),
				ServiceName:   s.GetName(),
				GoStructName:  fmt.Sprintf("%sServer", s.GetName()),
				Methods:       methods,
				GoNamespace:   strings.Replace(f.GetPackage(), ".", "_", -1), //golang package names use _ instead of .
			}
		}
	}
	return services
}

// GetHeaders get headers
func (m ProtoSet) GetHeaders() []string {
	headers := map[string]bool{}
	for _, h := range m.writer.GetHeaders() {
		headers[h] = true
	}
	for _, t := range m.types {
		if !t.Referenced {
			continue
		}
		line := fmt.Sprintf("%s \"%s\"", t.GoNamespace, t.GoImport)
		headers[line] = true
	}
	rv := []string{}
	for k := range headers {
		rv = append(rv, k)
	}
	return rv
}

func (m ProtoSet) getStubs() string {
	return m.writer.GetStubs()
}
