package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vendasta/mscli/pkg/protos"
)

var (
	stubsStyle string

	stubsConvertedStyle protos.StubStyle

	stubsCmd = &cobra.Command{
		Use:     "stubs",
		Short:   "Compile and build protos with stubs",
		Long:    `Compile and build protos with stubs`,
		PreRunE: stubsValidate,
		RunE:    generateStubs,
	}
)

func init() {
	stubsCmd.Flags().StringVarP(&stubsStyle, "style", "s", "empty", "Which style stub should we generate? ['empty', 'handler']")

	appCmd.AddCommand(stubsCmd)
}

func stubsValidate(cmd *cobra.Command, args []string) error {
	var err error
	stubsConvertedStyle, err = protos.StubStyleFromString(stubsStyle)
	return err
}

func generateStubs(cmd *cobra.Command, args []string) error {
	return protos.CompileAndGenerateStubs(microserviceSpec, stubsConvertedStyle)
}
