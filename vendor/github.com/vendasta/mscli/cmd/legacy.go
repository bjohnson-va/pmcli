package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vendasta/mscli/pkg/io"
)

// TODO: Delete these when project jenkinsfile are utilizing mscli's new interface correctly - dwalker

var lTestCmd = &cobra.Command{
	Use:    "test",
	Hidden: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		_, err := io.RunCommand("mscli", "app", "test").Execute()
		return err
	},
}

var lVetCmd = &cobra.Command{
	Use:    "vet",
	Hidden: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		_, err := io.RunCommand("mscli", "app", "vet").Execute()
		return err
	},
}

var lLintCmd = &cobra.Command{
	Use:    "lint",
	Hidden: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		_, err := io.RunCommand("mscli", "app", "lint").Execute()
		return err
	},
}

func init() {
	RootCmd.AddCommand(lTestCmd)
	RootCmd.AddCommand(lVetCmd)
	RootCmd.AddCommand(lLintCmd)
}
