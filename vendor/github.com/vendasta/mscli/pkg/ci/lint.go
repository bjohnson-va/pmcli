package ci

import (
	"fmt"
	"github.com/vendasta/mscli/pkg/docker"
	"github.com/vendasta/mscli/pkg/io"
	"github.com/vendasta/mscli/pkg/spec"
	"github.com/vendasta/mscli/pkg/utils"
)

// RunLint runs golint through mscli
func RunLint(specFile spec.MicroserviceFile, version string) error {
	fmt.Println("Running Lint...")
	var err error
	if utils.IsOnJenkins() {
		_, err = io.RunCommand("golint", "-set_exit_status", specFile.Microservice.GetPathToCodeImplementationDir()+"...").Execute()
	} else {
		dockerTag, err := docker.BuildDockerImage(specFile, version)
		if err != nil {
			return err
		}
		_, err = io.DockerCommand("run", dockerTag, "golint", "-set_exit_status", fmt.Sprintf("../%s/...", specFile.Microservice.GetCodeImplementationPackageName())).Execute()
	}

	if err != nil {
		return fmt.Errorf("error(s) running lint %s", err.Error())
	}
	return nil
}
