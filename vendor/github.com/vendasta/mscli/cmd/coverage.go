package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vendasta/mscli/pkg/ci"
)

var coverageCmd = &cobra.Command{
	Use:   "coverage",
	Short: "Runs coverage on your microservice",
	Long:  "Runs coverage on your microservice",
	RunE:  coverage,
}

func coverage(cmd *cobra.Command, args []string) error {
	return ci.RunCoverage(microserviceSpec, tag)
}

func init() {
	appCmd.AddCommand(coverageCmd)
}
