package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vendasta/mscli/pkg/ci"
)

var vetCmd = &cobra.Command{
	Use:   "vet",
	Short: "Runs vet on your microservice",
	Long:  "Runs vet on your microservice",
	RunE:  vet,
}

func vet(cmd *cobra.Command, args []string) error {
	return ci.RunVet(microserviceSpec, tag)
}

func init() {
	appCmd.AddCommand(vetCmd)
}
