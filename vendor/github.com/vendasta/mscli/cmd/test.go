package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vendasta/mscli/pkg/ci"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Runs tests on your microservice",
	Long:  "Runs tests on your microservice",
	RunE:  test,
}

func test(cmd *cobra.Command, args []string) error {
	return ci.RunTests(microserviceSpec, tag)
}

func init() {
	appCmd.AddCommand(testCmd)
}
