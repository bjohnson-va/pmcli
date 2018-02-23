package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vendasta/mscli/pkg/datadog"
)

var (
	datadogForce          bool

	datadogCommand = &cobra.Command{
		Use:   "datadog",
		Short: "Build the default dashboard.",
		Long:  `Build the default dashboard.`,
		RunE:  provisionDashboards,
	}
)

func init() {
	datadogCommand.Flags().BoolVarP(&datadogForce, "force", "f", false, "Passes the --force flag to to overwrite existing dashboards.")

	appCmd.AddCommand(datadogCommand)
}

func provisionDashboards(cmd *cobra.Command, args []string) error {
	return datadog.EnsureDashboard(microserviceSpec, datadogForce)
}
