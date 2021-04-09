package cmd

import (
	"os"
	"strings"

	"github.com/bjohnson-va/pmcli/mockserver"
	"github.com/spf13/cobra"
)

var (
	mockServerPort              int64
	mockServerAllowedOrigin     string
	mockServerSource            string
	mockServerConfigFile        string
	mockServerRandomValueSource string
	mockServerInteractive       bool

	mockServerCmd = &cobra.Command{
		Use:   "serve",
		Short: "Build a mock server from a proto specification",
		Long:  `Build a mock server from a proto specification in vendastaapis`,
		RunE:  runMockServer,
	}
)

func init() {
	mockServerCmd.Flags().Int64VarP(&mockServerPort, "port", "p",
		-1,
		"Sets the port of the mocked server")

	mockServerCmd.Flags().StringVarP(&mockServerAllowedOrigin, "allowedOrigin", "b",
		"null",
		"The origin from which requests will be made to this server")

	protoRootDir := getProtoRootDirectory()
	mockServerCmd.Flags().StringVarP(&mockServerSource, "source", "s",
		protoRootDir,
		"Directory containing source proto files")

	mockServerCmd.Flags().StringVarP(&mockServerConfigFile, "config", "c",
		"mockserver.json",
		"Config file")

	mockServerCmd.Flags().StringVarP(&mockServerRandomValueSource, "random", "r",
		"breadcrumb",
		"Randomization seed: Choose one of [breadcrumb, time]")

	mockServerCmd.Flags().BoolVarP(&mockServerInteractive, "interactive", "i",
		true,
		"Interactive prompts. Set True to configure endpoint at runtime. "+
			"False will reach from config file only with auto-reload.")

	RootCmd.AddCommand(mockServerCmd)
}

func getProtoRootDirectory() string {
	goPath := os.Getenv("GOPATH")
	rootOverride := os.Getenv("PMCLI_ROOT")
	protoRootDir := strings.TrimRight(rootOverride, "/")
	if protoRootDir == "" {
		protoRootDir = goPath + "/src/github.com/vendasta/vendastaapis"
	}
	return protoRootDir
}

func runMockServer(cmd *cobra.Command, args []string) error {
	return mockserver.BuildAndRun(
		mockServerPort, mockServerAllowedOrigin, mockServerSource,
		mockServerConfigFile, mockServerRandomValueSource, mockServerInteractive,
	)
}
