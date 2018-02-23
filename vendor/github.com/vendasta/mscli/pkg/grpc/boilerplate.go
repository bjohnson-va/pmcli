package grpc

import (
	"fmt"
	mscliio "github.com/vendasta/mscli/pkg/io"
	"github.com/vendasta/mscli/pkg/spec"
)

// CreateBoilerplate generates the initial server.go file required for mscli projects
func CreateBoilerplate(s spec.MicroserviceFile) error {
	if s.Microservice.Debug {
		fmt.Println("Creating GRPC boilerplate")
	}
	data := ServerTemplateData{
		Name:           s.Microservice.Name,
		VerifyIdentity: s.Microservice.IdentityType == spec.IdentityServiceAccountJwt,
	}
	return mscliio.CreateTemplatedFile("./server/main.go", data, ServerTemplate)
}
