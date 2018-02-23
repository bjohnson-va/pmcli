package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vendasta/mscli/pkg/endpoints"
)

var (
	endpointsForce          bool
	endpointsSkipGeneration bool

	endpointsCmd = &cobra.Command{
		Use:   "endpoints",
		Short: "Build and deploy endpoints.",
		Long:  `Build and deploy endpoints. WARNING: set your kubectl context to vendasta-central before running this.`,
		RunE:  deployEndpoints,
	}
)

func init() {
	endpointsCmd.Flags().BoolVarP(&endpointsForce, "force", "f", false, "Passes the --force flag to the endpoints deploy.")
	endpointsCmd.Flags().BoolVarP(&endpointsSkipGeneration, "skip-generation", "s", false, "Skips generation of the endpoints yaml files and just uses the existing ones for deployment.")

	appCmd.AddCommand(endpointsCmd)
}

func deployEndpoints(cmd *cobra.Command, args []string) error {
	return endpoints.DeployAndGenerate(microserviceSpec, endpointsSkipGeneration, endpointsForce, env)
}
