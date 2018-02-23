package bootstrap

import (
	"fmt"
	"github.com/vendasta/mscli/pkg/ci"
	"github.com/vendasta/mscli/pkg/docker"
	"github.com/vendasta/mscli/pkg/grpc"
	mscliio "github.com/vendasta/mscli/pkg/io"
	"github.com/vendasta/mscli/pkg/protos"
	"github.com/vendasta/mscli/pkg/spec"
	"github.com/vendasta/mscli/pkg/utils"
)

func Bootstrap(stubStyle protos.StubStyle, name string) error {
	s, err := spec.GenConfig(name)
	if err != nil {
		return err
	}
	sp := *s
	if sp.Microservice.Debug {
		fmt.Println("Bootstrapping a new microservice")
	}
	err = utils.AssertDepIsInstalled()
	if err != nil {
		return err
	}
	err = mscliio.EnsureDirExists(sp, sp.Microservice.GetPathToServerDir())
	if err != nil {
		return err
	}

	//Create a configuration
	err = grpc.CreateBoilerplate(sp)
	if err != nil {
		return err
	}
	err = ci.CreateBoilerplate(sp)
	if err != nil {
		return err
	}
	err = docker.CreateDockerComposeYaml(sp)
	if err != nil {
		return err
	}

	RunDepInit()

	if sp.Microservice.Debug {
		fmt.Println("Completed Bootstrapping.")
	}
	return protos.CompileAndGenerateStubs(sp, stubStyle)
}

// RunDepInit Runs dep through mscli
func RunDepInit() error {
	fmt.Println("Running `dep init`, this will take a while")
	_, err := mscliio.RunCommand("dep", "init").Execute()
	if err != nil {
		return fmt.Errorf("error running `dep init` %s", err.Error())
	}
	return nil
}
