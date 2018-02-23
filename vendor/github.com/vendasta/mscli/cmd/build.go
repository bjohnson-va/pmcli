package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vendasta/mscli/pkg/docker"
	"github.com/vendasta/mscli/pkg/utils"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Builds a docker image",
	Long:  `Builds a new Docker Image for your microservice with CloudBuilder.`,
	RunE:  build,
}

func build(cmd *cobra.Command, args []string) error {
	if env == utils.Local {
		dockerTag, err := docker.BuildDockerImage(microserviceSpec, tag)
		if err != nil {
			return err
		}
		err = docker.CopyLocalDockerTag(dockerTag, microserviceSpec.Microservice.Name)
		return err
	}
	return docker.BuildCloudContainer(microserviceSpec, tag)
}

func init() {
	appCmd.AddCommand(buildCmd)
}
