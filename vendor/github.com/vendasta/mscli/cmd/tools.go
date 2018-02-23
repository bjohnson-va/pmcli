package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vendasta/mscli/pkg/tools"
)

var toolsCmd = &cobra.Command{
	Use:   "tools",
	Short: "Installs and verifies that all prerequisite tools are installed.",
	Long:  `Installs and verifies that all prerequisite tools are installed.`,
	RunE:  toolsInstall,
}

func toolsInstall(cmd *cobra.Command, args []string) error {
	return tools.InstallGoTools()
}

func init() {
	RootCmd.AddCommand(toolsCmd)
}
