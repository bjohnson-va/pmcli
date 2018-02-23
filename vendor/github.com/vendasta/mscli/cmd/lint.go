package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vendasta/mscli/pkg/ci"
)

var lintCmd = &cobra.Command{
	Use:   "lint",
	Short: "Runs lint on your microservice",
	Long:  "Runs lint on your microservice",
	RunE:  lint,
}

func lint(cmd *cobra.Command, args []string) error {
	return ci.RunLint(microserviceSpec, tag)
}

func init() {
	appCmd.AddCommand(lintCmd)
}
