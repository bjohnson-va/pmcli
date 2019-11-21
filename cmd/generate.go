package cmd

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"github.com/bjohnson-va/pmcli/config"
	"github.com/vendasta/gosdks/logging"
	"io/ioutil"
	"os"
	"strings"

	"github.com/spf13/cobra"
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

type MockServerJson struct {
	Port          int64                  `json:"port"`
	AllowedOrigin string                 `json:"allowedOrigin"`
	Https         bool                   `json:"useHttps"`
	ProtoPaths    []string               `json:"protofiles"`
	Overrides     map[string]interface{} `json:"overrides"`
	Instructions  map[string]interface{} `json:"instructions"`
	Exclusions    map[string]interface{} `json:"exclusions"`
}

func generate(cmd *cobra.Command, args []string) {
	ctx := context.Background()
	if _, err := os.Stat("./" + config.FILENAME); !os.IsNotExist(err) {
		fmt.Println(config.FILENAME + " already exists. Please remove it before running the generator.")
		return
	}
	reader := bufio.NewReader(os.Stdin)

	port := promptForPort(reader)
	https := promptForHttps(reader)
	allowedOrigin := promptForAllowedOrigin(reader)
	protopath := promptForProtoPath(reader)

	j, err := json.MarshalIndent(&MockServerJson{
		Port:          port,
		AllowedOrigin: allowedOrigin,
		Https:         https,
		ProtoPaths:    []string{protopath},
		Overrides:     map[string]interface{}{},
		Instructions:  map[string]interface{}{},
		Exclusions:    map[string]interface{}{},
	}, "", "  ")
	if err != nil {
		logging.Errorf(ctx, "Error marshalling %s: %s", config.FILENAME, err.Error())
		return
	}
	err = ioutil.WriteFile("./"+config.FILENAME, append(j, ([]byte)("\n")...), 0644)
	if err != nil {
		logging.Errorf(ctx, "Error writing to %s: %s", config.FILENAME, err.Error())
		return
	}
	fmt.Printf("Successfully wrote to %s:\n%s\n", config.FILENAME, j)
}

func promptForPort(reader *bufio.Reader) int64 {
	return 28000 // TODO: Prompt
}

func promptForHttps(reader *bufio.Reader) bool {
	return true // TODO: Prompt
}

func promptForAllowedOrigin(reader *bufio.Reader) string {
	return "localhost:4000" // TODO: Prompt
}

func promptForProtoPath(reader *bufio.Reader) string {
	fmt.Printf("Enter the path for your API proto file (relative to %s): ", mockServerSource)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	// TODO: Stop assuming these will be in "$GOPATH/src/github.com/vendasta/vendastaapis"
	return text
}

func init() {
	RootCmd.AddCommand(generateCmd)
}
