package ci

import (
	"fmt"
	"github.com/vendasta/mscli/pkg/docker"
	"github.com/vendasta/mscli/pkg/io"
	"github.com/vendasta/mscli/pkg/spec"
	"github.com/vendasta/mscli/pkg/utils"
)

// RunVet runs go vet through mscli
func RunVet(specFile spec.MicroserviceFile, version string) error {
	fmt.Println("Running Vet...")

	var err error
	if utils.IsOnJenkins() {
		_, err = io.RunCommand("go", "vet", specFile.Microservice.GetPathToCodeImplementationDir()+"...").Execute()
	} else {
		dockerTag, err := docker.BuildDockerImage(specFile, version)
		if err != nil {
			return err
		}
		_, err = io.DockerCommand("run", dockerTag, "go", "tool", "vet", fmt.Sprintf("../%s/", specFile.Microservice.GetCodeImplementationPackageName())).Execute()
	}

	if err != nil {
		return fmt.Errorf("error(s) running vet %s", err.Error())
	}
	return nil
}
