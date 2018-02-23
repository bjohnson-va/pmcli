package internal

import (
	"fmt"
	"strings"

	"github.com/vendasta/mscli/tools/sdkgen/util"
)

// Proto Services on a proto
type Proto struct {
	MicroserviceName string
	Services         []*Service
	NonServiceProtos []*Service
}

// GRPCServiceName grpc service name
func (p *Proto) GRPCServiceName() string {
	return fmt.Sprintf("_%sTransportOverGRPC", p.Services[0].Name)
}

// HTTPServiceName http service name
func (p *Proto) HTTPServiceName() string {
	return fmt.Sprintf("_%sTransportOverHttp", p.Services[0].Name)
}

// TsApiServiceFileName filename of the typescript api service file
func (p *Proto) TsFileName() string {
	return fmt.Sprintf("%s", util.ToKebabCase(p.MicroserviceName))
}

// TsApiServiceName service name of the typescript api service
func (p *Proto) TsName() string {
	return fmt.Sprintf("%s", util.ToPascalCase(p.MicroserviceName))
}

// UniqueRequestAndResponseClassNames used in proto file
func (p *Proto) UniqueRequestAndResponseClassNames() []string {
	var classNames []string
	for _, s := range p.Services {
		for _, n := range s.UniqueRequestAndResponseClassNames() {
			if !contains(classNames, n) {
				classNames = append(classNames, n)
			}
		}
	}
	return classNames
}

// IsGoogleProtoFileName returns if the provided proto is from google
func IsGoogleProtoFileName(fn string) bool {
	return strings.Contains(fn, "google/protobuf") || strings.Contains(fn, "google/api")
}
