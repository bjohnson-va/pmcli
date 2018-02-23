package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
)

var debug bool

var RootCmd = &cobra.Command{
	Use:   "mscli",
	Short: "mscli helps manage Vendasta microservices.",
	Long: `mscli helps manage Vendasta microservices.
This includes both bootstrapping new projects as well as updating and maintaining existing projects.`,
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	RootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "show verbose information")
}
