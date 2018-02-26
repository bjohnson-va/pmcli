package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const cliVersion = "1.2.1"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Outputs the current CLI version.",
	Long:  "Outputs the current CLI version.",
	Run:   version,
}

func version(cmd *cobra.Command, args []string) {
	fmt.Printf("%s\n", cliVersion)
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
