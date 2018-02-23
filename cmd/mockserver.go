package cmd

import (
	"github.com/spf13/cobra"
	"os"
	"github.com/bjohnson-va/pmcli/mockserver"
)

var (
	mockServerPort int
	mockServerAllowedOrigin string
	mockServerSource string
	mockServerConfigFile string

	mockServerCmd = &cobra.Command{
		Use:   "serve",
		Short: "Build a mock server from a proto specification",
		Long:  `Build a mock server from a proto specification in vendastaapis`,
		RunE:  runMockServer,
	}

)

func init() {
	mockServerCmd.Flags().IntVarP(&mockServerPort, "port", "p",
		20001,
		"Sets the port of the mocked server")

	mockServerCmd.Flags().StringVarP(&mockServerAllowedOrigin, "allowedOrigin", "b",
		"http://localhost:4200",
		"The origin from which requests will be made to this server")

	goPath := os.Getenv("GOPATH")
	mockServerCmd.Flags().StringVarP(&mockServerSource, "source", "s",
		goPath+ "/src/github.com/vendasta/vendastaapis",
		"Directory containing source proto files")

	mockServerCmd.Flags().StringVarP(&mockServerConfigFile, "config", "c",
		"",
		"Config file")

	RootCmd.AddCommand(mockServerCmd)
}

func runMockServer(cmd *cobra.Command, args []string) error {
	return mockserver.BuildAndRun(mockServerPort, mockServerAllowedOrigin,
		mockServerSource, mockServerConfigFile)
}
