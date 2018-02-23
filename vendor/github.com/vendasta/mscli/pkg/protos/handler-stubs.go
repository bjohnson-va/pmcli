package protos

import (
	"bytes"
	"fmt"
	"github.com/vendasta/mscli/pkg/grpc"
	"html/template"
	"log"
	"strings"
)

const (
	handlerStubServerDefinition = `type %s struct {
%s
}`
)

type HandlerStubs struct {
	services  map[string]serviceInfo
	types     map[string]typeInfo
	emptyStub EmptyStubs
}

func (hs HandlerStubs) GetHeaders() []string {
	return []string{`"github.com/vendasta/gosdks/util"`, `"context"`, `"github.com/vendasta/grpc-go-handler/pkg/grpchandler"`}
}

func (hs HandlerStubs) GetStubs() string {
	output := ""
	for _, s := range hs.services {
		fs := ""
		fh := []string{}
		for _, mod := range s.Methods {
			fs += hs.generateStub(s.ServiceName, mod.MethodName) + "\n"
			fh = append(fh, mod.MethodName+"TypedHandler")
		}
		output += hs.generateStubTypeDefinition(s.ServiceName, fh) + "\n"
		output += fs
	}
	return output
}

type FunctionSpec struct {
	GrpcServiceName     string
	FunctionName        string
	DomainNamespace     string
	GoRequestNamespace  string
	RequestType         string
	GoResponseNamespace string
	ResponseType        string
}

func (hs HandlerStubs) generateStub(serviceName, methodName string) string {
	s, ok := hs.services[serviceName]
	if !ok {
		return ""
	}
	method, ok := s.Methods[methodName]
	if !ok {
		return ""
	}

	if !(!method.ServerStreaming && !method.ClientStreaming) {
		return hs.emptyStub.GenerateStub(serviceName, methodName)
	}

	rq, ok := hs.types[method.RequestType]
	if !ok {
		log.Fatalf("Unknown request type while generating stubs %s", method.RequestType)
	}

	rs, ok := hs.types[method.ResponseType]
	if !ok {
		log.Fatalf("Unknown response type while generating stubs %s", method.ResponseType)
	}

	var tmpl *template.Template
	var err error
	fs := FunctionSpec{serviceName + "Server", methodName, strings.ToLower(serviceName), rq.GoNamespace, rq.GoTypeName, rs.GoNamespace, rs.GoTypeName}
	if tmpl, err = template.New("grpcFunc").Parse(grpc.GrpcFuncTemplate); err != nil {
		log.Fatalf("Error creating golang template: %s", err.Error())
	}

	buf := bytes.NewBufferString("")
	if err = tmpl.Execute(buf, fs); err != nil {
		log.Fatalf("Error executing golang template: %s", err.Error())
	}
	return buf.String()
}

func (hs HandlerStubs) generateStubTypeDefinition(serviceName string, functionHandlers []string) string {
	fs := ""
	for _, f := range functionHandlers {
		fs += fmt.Sprintf("\t%s %s\n", f, f)
	}
	output := fmt.Sprintf(handlerStubServerDefinition, serviceName+"Server", fs)
	return output
}
