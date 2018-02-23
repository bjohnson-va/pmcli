package ci

import (
	"fmt"
	"github.com/vendasta/mscli/pkg/docker"
	"github.com/vendasta/mscli/pkg/io"
	"github.com/vendasta/mscli/pkg/spec"
	"github.com/vendasta/mscli/pkg/utils"
	"os"
	"path"
)

// RunCoverage runs gotestcover through mscli
func RunCoverage(specFile spec.MicroserviceFile, version string) error {
	fmt.Println("Running Coverage...")
	var err error
	err = io.EnsureDirExists(specFile, "./coverage")
	if err != nil {
		return err
	}
	if utils.IsOnJenkins() {
		os.Setenv("ENVIRONMENT", "test")
		fmt.Println("Running go test with coverage...")
		_, err = io.RunCommand("gotestcover", "-coverprofile=./coverage/c.out", specFile.Microservice.GetPathToCodeImplementationDir()+"...").Execute()
		if err != nil {
			return fmt.Errorf("error(s) running coverage %s", err.Error())
		}
		fmt.Println("Finished running go test with coverage...")
		fmt.Println("Building coverage artifact...")
		_, err = io.RunCommand("go", "tool", "cover", "-html=./coverage/c.out", "-o=./coverage/cover.html").Execute()
		if err != nil {
			return fmt.Errorf("error(s) running coverage %s", err.Error())
		}
		fmt.Println("Finished building coverage artifact...")
	} else {
		fmt.Println("Running go test with coverage...")
		currDir, _ := os.Getwd()
		coveragePathMount := path.Join(currDir, "coverage:/coverage")
		dockerTag, err := docker.BuildDockerImage(specFile, version)
		if err != nil {
			return err
		}
		_, err = io.DockerCommand("run", "-v", coveragePathMount, "--env", "ENVIRONMENT=test", dockerTag, "gotestcover", "-coverprofile=/coverage/c.out", fmt.Sprintf("../%s/%s/...", specFile.Microservice.GoPackageName, specFile.Microservice.GetPathToCodeImplementationDir())).Execute()
		if err != nil {
			return fmt.Errorf("error(s) running coverage %s", err.Error())
		}
		fmt.Println("Finished running go test with coverage...")
		fmt.Println("Building coverage artifact...")
		_, err = io.DockerCommand("run", "-v", coveragePathMount, dockerTag, "go", "tool", "cover", "-html=/coverage/c.out", "-o=/coverage/cover.html").Execute()
		if err != nil {
			return fmt.Errorf("Error(s) running coverage %s", err.Error())
		}
		fmt.Println("Finished building coverage artifact...")
		fmt.Println("Run `open coverage/cover.html` to view it.")
	}
	fmt.Println("Finished running Coverage...")
	return nil
}
