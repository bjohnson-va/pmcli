package docker

import (
	"fmt"
	mscliio "github.com/vendasta/mscli/pkg/io"
	"github.com/vendasta/mscli/pkg/spec"
	"github.com/vendasta/mscli/pkg/utils"
)

// CreateDockerComposeYaml generates the initial docker-compose.yaml and docker-compose.debug.yaml required by mscli projects
func CreateDockerComposeYaml(spec spec.MicroserviceFile) error {
	if spec.Microservice.Debug {
		fmt.Println("Creating docker-compose.yaml")
	}

	grpcPort := utils.GetUserInput("The Grpc Port for local. Must be unique with the other µs you have running locally", "21000")
	httpPort := utils.GetUserInput("The Http Port for local. Must be unique with the other µs you have running locally", "21001")
	delvePort := utils.GetUserInput("The Delve port for local. Must be unique with the other µs you have running locally. This is what you'll use to hit your µs", "21002")
	endpointsPort := utils.GetUserInput("The Endpoints Port for local. Must be unique with the other µs you have running locally. This is what you'll use to hit your µs", "21003")

	data := DockerComposeYamlTemplateData{
		Name:          spec.Microservice.Name,
		GRPCPort:      grpcPort,
		HTTPPort:      httpPort,
		EndpointsPort: endpointsPort,
	}
	err := mscliio.CreateTemplatedFile("./docker-compose.yaml", data, DockerComposeYamlTemplate)
	if err != nil {
		return err
	}

	dbgData := DockerDebugComposeYamlTemplateData{
		Name:          spec.Microservice.Name,
		DelvePort:     delvePort,
		GoPackageName: spec.Microservice.GoPackageName,
	}
	return mscliio.CreateTemplatedFile("./docker-compose.debug.yaml", dbgData, DockerDebugComposeYamlTemplate)
}
