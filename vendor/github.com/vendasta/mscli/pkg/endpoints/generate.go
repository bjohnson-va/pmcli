package endpoints

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	mscliio "github.com/vendasta/mscli/pkg/io"
	"github.com/vendasta/mscli/pkg/spec"
	"github.com/vendasta/mscli/pkg/utils"
)

// GenerateEnvironmentSpecificEndpoints Generate endpoints
func GenerateEnvironmentSpecificEndpoints(s spec.MicroserviceFile, workingDir string, env utils.Environment) error {
	fmt.Printf("Creating environment/microservice-name.yml\n")

	pbFile := fmt.Sprintf("%s/pb/descriptor.pb", workingDir)
	err := mscliio.EnsureFileExists(pbFile)
	if err != nil {
		return err
	}

	bytes, err := ioutil.ReadFile(pbFile)
	if err != nil {
		return fmt.Errorf("error, could not read %s", pbFile)
	}

	fds := descriptor.FileDescriptorSet{}
	proto.Unmarshal(bytes, &fds)

	err = mscliio.EnsureDirExists(s, fmt.Sprintf("./endpoints/%s", env.String()))
	if err != nil {
		return err
	}
	gRPCService, err := buildEnvironmentEndpointsData(s, &fds, s.Microservice.Name, env)
	if err != nil {
		return nil
	}

	if env == utils.Local {
		err := mscliio.CreateTemplatedFile("endpoints/local/local-service-account.json", nil, LocalServiceAccountJSON)
		if err != nil {
			return err
		}
	}

	return mscliio.CreateTemplatedFile(getAPIConfigFileName(s, env), gRPCService, EndpointsEnvironment)
}

func buildEnvironmentEndpointsData(s spec.MicroserviceFile, protos *descriptor.FileDescriptorSet, name string, env utils.Environment) (*GRPCService, error) {
	var apis []string
	for _, proto := range protos.File {
		for _, service := range proto.Service {
			apis = append(apis, fmt.Sprintf("%s.%s", *proto.Package, *service.Name))
		}
	}

	se, err := s.Microservice.GetEnv(env)
	if err != nil {
		return nil, err
	}

	publicRoutes := make([]string, len(s.Microservice.PublicRoutes))
	for idx, r := range s.Microservice.PublicRoutes {
		publicRoutes[idx] = convertEndpointsPathToEndpointsSelector(r)
	}

	return &GRPCService{
		Apis:            apis,
		Name:            se.GRPCHost,
		ProjectName:     name,
		EnvironmentName: env.String(),
		VerifyIdentity:  s.Microservice.IdentityType == spec.IdentityServiceAccountJwt,
		PublicRoutes:    publicRoutes,
	}, nil
}

func convertEndpointsPathToEndpointsSelector(p string) string {
	if strings.HasPrefix(p, "/") {
		p = strings.Replace(p, "/", "", 1)
	}
	return strings.Replace(p, "/", ".", -1)
}

func getAPIConfigFileName(spec spec.MicroserviceFile, env utils.Environment) string {
	return fmt.Sprintf("./endpoints/%s/%s.yml", env.String(), spec.Microservice.Name)
}
