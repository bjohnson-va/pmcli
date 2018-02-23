package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vendasta/mscli/pkg/sdk"
	"github.com/vendasta/mscli/pkg/sdk/base"
)

var (
	sdkBootstrap bool
	sdkLanguage  string
	sdkSource    string

	sdkConvertedLanguage base.Language

	sdkCmd = &cobra.Command{
		Use:     "sdk",
		Short:   "Generate an SDK",
		Long:    `Generates an SDK for your microservice.`,
		PreRunE: sdkValidate,
		RunE:    generateSdk,
	}
)

func init() {
	sdkCmd.Flags().BoolVarP(&sdkBootstrap, "bootstrap", "b", false, "Generate the initial SDK files.")
	sdkCmd.Flags().StringVarP(&sdkLanguage, "language", "l", "", "Which language to generate an SDK for. Choices: python, typescript, java")
	sdkCmd.Flags().StringVarP(&sdkSource, "source", "s", "./vendor/github.com/vendasta/gosdks/pb/", "Directory containing source PB files")

	appCmd.AddCommand(sdkCmd)
}

func sdkValidate(cmd *cobra.Command, args []string) error {
	var err error
	sdkConvertedLanguage, err = base.LanguageFromString(sdkLanguage)
	return err
}

func generateSdk(cmd *cobra.Command, args []string) error {
	return sdk.GenerateSDK(microserviceSpec, sdkConvertedLanguage, sdkSource, sdkBootstrap)
}
