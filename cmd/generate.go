package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/bjohnson-va/pmcli/config"
	"github.com/bjohnson-va/pmcli/configwizard"
	"github.com/spf13/cobra"
	"github.com/vendasta/gosdks/logging"
	"io/ioutil"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate " + config.FILENAME,
	Long: fmt.Sprintf(
		"An interactive helper for automatically generating the %s file required by PMCLI",
		config.FILENAME,
	),
	Run: generate,
}

func generate(cmd *cobra.Command, args []string) {
	ctx := context.Background()
	jsonData := configwizard.Prompt(ctx, mockServerSource)

	j, err := writeToFile(jsonData)
	if err != nil {
		logging.Errorf(ctx, "Failed to write server config file: %s", err.Error())
		return
	}
	fmt.Printf("Successfully wrote to %s:\n%s\n", config.FILENAME, j)
}

func writeToFile(o configwizard.MockServerJson) ([]byte, error) {
	j, err := json.MarshalIndent(&o, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("Error marshalling %s: %s", config.FILENAME, err.Error())
	}
	err = ioutil.WriteFile("./"+config.FILENAME, append(j, ([]byte)("\n")...), 0644)
	if err != nil {
		return nil, fmt.Errorf("Error writing to %s: %s", config.FILENAME, err.Error())
	}
	return j, nil
}

func init() {
	RootCmd.AddCommand(generateCmd)
}
