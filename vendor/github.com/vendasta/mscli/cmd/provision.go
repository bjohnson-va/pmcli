package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/vendasta/mscli/pkg/kubernetes"
)

var provisionCmd = &cobra.Command{
	Use:   "provision",
	Short: "Provision the k8s namespace, service account and secrets for your microservice.",
	Long:  `Provision the k8s namespace, service account and secrets for your microservice.`,
	RunE:  provision,
}

func provision(cmd *cobra.Command, args []string) error {
	serviceAccountName, err := kubernetes.ProvisionService(microserviceSpec, env)
	fmt.Printf("Your service account is: %s\n", serviceAccountName)
	return err
}

func init() {
	appCmd.AddCommand(provisionCmd)
}
