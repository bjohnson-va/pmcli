package cmd

import (
	"github.com/bjohnson-va/pmcli/config"
	"github.com/bjohnson-va/pmcli/mockserver"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var (
	mockServerPort              int64
	mockServerAllowedOrigin     string
	mockServerSource            string
	mockServerConfigFile        string
	mockServerRandomValueSource string

	mockServerCmd = &cobra.Command{
		Use:   "serve",
		Short: "Run a mock server from a proto specification",
		Long:  `Run a mock server from a proto specification in vendastaapis`,
		RunE:  runMockServer,
	}

	mockServerHelperCmd = &cobra.Command{
		Use:   "serve-ng",
		Short: "Run a mock server for an Angular App",
		Long:  `Run a mock server with extra prompts to help integrate with an Angular app`,
		RunE:  runMockServerForAngular,
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

	RootCmd.AddCommand(mockServerCmd)
	RootCmd.AddCommand(mockServerHelperCmd)
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
		mockServerConfigFile, mockServerRandomValueSource,
		config.AssistUnset,
	)
}

func runMockServerForAngular(cmd *cobra.Command, args []string) error {
	return mockserver.BuildAndRun(
		mockServerPort, mockServerAllowedOrigin, mockServerSource,
		mockServerConfigFile, mockServerRandomValueSource,
		config.AssistAngular,
	)
}
