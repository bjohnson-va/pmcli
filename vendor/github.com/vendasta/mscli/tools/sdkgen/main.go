package main

import (
	"flag"
	"io"
	"io/ioutil"
	"os"

	"fmt"
	"strings"

	"github.com/golang/glog"
	"github.com/golang/protobuf/proto"
	google_protobuf "github.com/golang/protobuf/protoc-gen-go/descriptor"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/vendasta/mscli/tools/sdkgen/internal"
	"github.com/vendasta/mscli/tools/sdkgen/java"
	"github.com/vendasta/mscli/tools/sdkgen/python"
	"github.com/vendasta/mscli/tools/sdkgen/typescript"
)

// Attempt to parse the incoming CodeGeneratorRequest being written by `protoc` to our stdin
func parseReq(r io.Reader) (*plugin.CodeGeneratorRequest, error) {
	glog.V(1).Info("Parsing code generator request")
	input, err := ioutil.ReadAll(r)

	if err != nil {
		glog.Errorf("Failed to read code generator request from stdin: %v", err)
		return nil, err
	}

	req := new(plugin.CodeGeneratorRequest)
	if err = proto.Unmarshal(input, req); err != nil {
		glog.Errorf("Failed to unmarshal code generator request: %v", err)
		return nil, err
	}
	glog.V(1).Info("Successfully parsed code generator request")
	return req, nil
}

// DescriptorProtoByName get a descriptor proto by its name
func DescriptorProtoByName(dp []*google_protobuf.DescriptorProto, replyClass string) *google_protobuf.DescriptorProto {
	pieces := strings.Split(replyClass, ".")
	for _, d := range dp {
		if *d.Name == pieces[len(pieces)-1] {
			return d
		}
	}
	return nil
}

func main() {
	flag.Parse()
	defer glog.Flush()

	glog.V(1).Info("Processing the CodeGeneratorRequest")
	request, err := parseReq(os.Stdin)
	if err != nil {
		glog.Fatal(err)
	}

	allDescriptorProtos := []*google_protobuf.DescriptorProto{}
	for _, p := range request.ProtoFile {
		allDescriptorProtos = append(allDescriptorProtos, p.MessageType...)
	}

	services := []*internal.Service{}
	nonServiceProtos := []*internal.Service{}
	messageEnumDescriptors := []*internal.MessageEnumDescriptors{}

	for _, p := range request.ProtoFile {
		if len(p.Service) == 0 {
			if !internal.IsGoogleProtoFileName(*p.Name) {
				nonServiceProtos = append(nonServiceProtos, &internal.Service{FileName: *p.Name})
			}
		} else {
			for _, svcProto := range p.Service {
				svc := &internal.Service{}
				svc.Name = *svcProto.Name
				svc.FileName = *p.Name
				services = append(services, svc)
				for _, method := range svcProto.Method {
					svc.RPCS = append(svc.RPCS, &internal.RPC{
						Name:               *method.Name,
						Response:           *method.OutputType,
						Request:            *method.InputType,
						Path:               fmt.Sprintf("/%s.%s/%s", p.GetPackage(), svc.Name, method.GetName()),
						RequestDescriptor:  DescriptorProtoByName(allDescriptorProtos, *method.InputType),
						ResponseDescriptor: DescriptorProtoByName(allDescriptorProtos, *method.OutputType),
					})

				}
			}
		}

		if !internal.IsGoogleProtoFileName(*p.Name) {
			if len(p.MessageType)+len(p.EnumType) > 0 {

				messages, enums := internal.ExtractMessagesAndEnums(p)

				messageEnumDescriptors = append(messageEnumDescriptors, &internal.MessageEnumDescriptors{
					FileName:     *p.Name,
					Dependencies: p.Dependency,
					Messages:     messages,
					Enums:        enums,
				})
			}
		}
	}

	params := strings.Split(request.GetParameter(), "&")
	p := &internal.Proto{
		MicroserviceName: params[1],
		Services:         services,
		NonServiceProtos: nonServiceProtos,
	}

	if len(services) == 0 {
		glog.Fatal("No proto services found, is your protopath pointing towards the correct proto file?")
	}

	var outputStruct *plugin.CodeGeneratorResponse
	switch params[0] {
	case "python":
		outputStruct = python.Generate(p, messageEnumDescriptors)
	case "typescript":
		outputStruct = typescript.Generate(p, messageEnumDescriptors)
	case "java":
		outputStruct = java.GenerateFiles(request)
	default:
		outputStruct = python.Generate(p, messageEnumDescriptors)
	}

	buf, err := proto.Marshal(outputStruct)

	if _, err := os.Stdout.Write(buf); err != nil {
		glog.Fatal(err)
	}
}
