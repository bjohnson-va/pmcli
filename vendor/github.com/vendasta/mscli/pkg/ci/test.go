package ci

import (
	"fmt"
	"github.com/vendasta/mscli/pkg/docker"
	"github.com/vendasta/mscli/pkg/io"
	"github.com/vendasta/mscli/pkg/spec"
	"github.com/vendasta/mscli/pkg/utils"
	"os"
)

// RunTests runs the tests through mscli
func RunTests(specFile spec.MicroserviceFile, version string) error {
	fmt.Println("Running Tests")
	var err error
	if utils.IsOnJenkins() {
		os.Setenv("ENVIRONMENT", "test")
		_, err = io.RunCommand("go", "test", specFile.Microservice.GetPathToCodeImplementationDir()+"...").Execute()
	} else {
		dockerTag, err := docker.BuildDockerImage(specFile, version)
		if err != nil {
			return err
		}
		_, err = io.DockerCommand("run", "--env", "ENVIRONMENT=test", dockerTag, "go", "test", fmt.Sprintf("../%s/...", specFile.Microservice.GetCodeImplementationPackageName())).Execute()
	}

	if err != nil {
		return fmt.Errorf("error running tests: %s", err.Error())
	}
	return nil
}
