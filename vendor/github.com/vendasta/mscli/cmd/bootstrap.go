package cmd

import (
	"github.com/spf13/cobra"
	bs "github.com/vendasta/mscli/pkg/bootstrap"
	"github.com/vendasta/mscli/pkg/protos"
)

var (
	bootstrapStyle string
	bootstrapName  string

	bootstrapConvertedStyle protos.StubStyle

	bootstrapCmd = &cobra.Command{
		Use:     "bootstrap",
		Short:   "Runs all the steps required for generating a microservice from proto file(s)",
		Long:    `Runs all the steps required for generating a microservice from proto file(s)`,
		PreRunE: bootstrapValidate,
		RunE:    bootstrap,
	}
)

func init() {
	bootstrapCmd.Flags().StringVarP(&bootstrapStyle, "style", "s", "empty", "Which style stub should we generate? ['empty', 'handler']")
	bootstrapCmd.Flags().StringVarP(&bootstrapName, "name", "n", "", "Name of the microservice.")

	RootCmd.AddCommand(bootstrapCmd)
}

func bootstrapValidate(cmd *cobra.Command, args []string) error {
	var err error
	bootstrapConvertedStyle, err = protos.StubStyleFromString(stubsStyle)
	return err
}

func bootstrap(cmd *cobra.Command, args []string) error {
	return bs.Bootstrap(bootstrapConvertedStyle, bootstrapName)
}
