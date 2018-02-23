package protos

import "fmt"

const (
	emptyStubServerDefinition = `type %s struct {}
`

	emptyStubFunctionDefinition = `func (s *%s) %s(ctx context.Context, req *%s) (*%s, error) {
	//TODO: CHANGE THIS TO DO STUFF
	logging.Debugf(ctx, "You should change this")
	res := &%s{}
	return res, nil
}`
	emptyStubStreamingFunctionDefinition = `func (s *%s) %s(ctx context.Context, stream *%s_%sServer) error {
	//TODO: CHANGE THIS TO DO STUFF
	logging.Debugf(ctx, "You should change this")
	res := &%s{}
	stream.Send(res)
	return nil
}`
)

type EmptyStubs struct {
	services map[string]serviceInfo
	types    map[string]typeInfo
}

func (es EmptyStubs) GetHeaders() []string {
	return []string{`"github.com/vendasta/gosdks/logging"`, `"context"`}
}

func (es EmptyStubs) GetStubs() string {
	output := ""
	for _, s := range es.services {
		output += es.generateStubTypeDefinition(s.ServiceName) + "\n"
		for _, mod := range s.Methods {
			output += es.GenerateStub(s.ServiceName, mod.MethodName) + "\n"
		}
	}
	return output
}

func (es EmptyStubs) generateStubTypeDefinition(serviceName string) string {
	return fmt.Sprintf(emptyStubServerDefinition, serviceName+"Server")
}

func (es EmptyStubs) GenerateStub(serviceName, methodName string) string {
	s, ok := es.services[serviceName]
	if !ok {
		return ""
	}
	method, ok := s.Methods[methodName]
	if !ok {
		return ""
	}

	requestType := method.RequestType
	rq, ok := es.types[method.RequestType]
	if ok {
		requestType = fmt.Sprintf("%s.%s", rq.GoNamespace, rq.GoTypeName)
	}

	responseType := method.ResponseType
	rs, ok := es.types[method.ResponseType]
	if ok {
		responseType = fmt.Sprintf("%s.%s", rs.GoNamespace, rs.GoTypeName)
	}

	gc := ""
	if !method.ServerStreaming {
		if !method.ClientStreaming {
			//No Streaming
			gc = fmt.Sprintf(emptyStubFunctionDefinition,
				s.GoStructName, method.MethodName,
				requestType, responseType, responseType)
		} else {
			//Client Streaming
			gc = "Client Streaming not supported"
		}
	} else {
		if !method.ClientStreaming {
			//Server Streaming
			gc = fmt.Sprintf(emptyStubStreamingFunctionDefinition,
				s.GoStructName, methodName,
				fmt.Sprintf("pb.%s", serviceName), rs.GoTypeName,
				responseType)
		} else {
			//Bidirectional Streaming
			gc = "Bi-directional streaming not supported"
		}
	}
	return gc
}
