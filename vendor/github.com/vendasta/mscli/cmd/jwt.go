package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vendasta/mscli/pkg/jwt"
)

var jwtScope string

var jwtCmd = &cobra.Command{
	Use:   "jwt",
	Short: "Generates a JWT with vendasta local creds (not IAM creds) to call a microservice",
	Long:  "Generates a JWT with vendasta local creds (not IAM creds) to call a microservice",
	RunE:  generateJwt,
}

func generateJwt(cmd *cobra.Command, args []string) error {
	return jwt.GenerateJWT(microserviceSpec, jwtScope, env)
}

func init() {
	jwtCmd.Flags().StringVar(&jwtScope, "scope", "", "Scope is specified in the auth section for the endpoints audience. If omitted, uses grpcHost from the microservice.yaml for the given environment")
	appCmd.AddCommand(jwtCmd)
}
