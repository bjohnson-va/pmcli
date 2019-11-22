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
	"strconv"
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
	MoreInfo      string                 `json:"moreInfo"`
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
	jsonData := MockServerJson{
		MoreInfo:      "https://github.com/bjohnson-va/pmcli",
	}
	if _, err := os.Stat("./" + config.FILENAME); !os.IsNotExist(err) {
		d, err := ioutil.ReadFile("./" + config.FILENAME)
		if err == nil {
			err = json.Unmarshal(d, &jsonData)
			if err != nil {
				warnf("Could not parse existing %s.\n"+
					"It will be overwritten [%s]\n", config.FILENAME, err.Error())
			}
		}
	}
	reader := bufio.NewReader(os.Stdin)

	jsonData.Port = promptForPort(reader, jsonData.Port)
	jsonData.AllowedOrigin = promptForAllowedOrigin(reader)
	jsonData.Https = promptForHttps(reader, jsonData.Https)
	jsonData.ProtoPaths = []string{promptForProtoPath(ctx, reader, jsonData.ProtoPaths[0])}
	jsonData.Overrides = map[string]interface{}{}
	jsonData.Instructions = map[string]interface{}{}
	jsonData.Exclusions = map[string]interface{}{}

	j, err := writeToFile(jsonData)
	if err != nil {
		logging.Errorf(ctx, "Failed to write server config file: %s", err.Error())
		return
	}
	fmt.Printf("Successfully wrote to %s:\n%s\n", config.FILENAME, j)
}

func writeToFile(o MockServerJson) ([]byte, error) {
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

func promptForPort(reader *bufio.Reader, current int64) int64 {
	defaultPort := 28000
	if current > 0 {
		defaultPort = int(current)
	}
	userValue := prompt(reader, "Enter the path for your API proto file", strconv.Itoa(defaultPort))
	atoi, err := strconv.Atoi(userValue)
	if err != nil {
		warnf("User input could not be understood [%s]\n", err.Error())
		return promptForPort(reader, current)
	}
	return int64(atoi)
}

func promptForHttps(reader *bufio.Reader, current bool) bool {
	defaultVal := "n"
	if current {
		defaultVal = "y"
	}
	userSelection := strings.ToLower(prompt(reader, "Use HTTPS? (y/n)", defaultVal))
	if userSelection != "y" && userSelection != "n" {
		warnf("User input could not be understood [%s]\n", userSelection)
		promptForHttps(reader, current)
	}
	return userSelection == "y"
}

func promptForAllowedOrigin(reader *bufio.Reader) string {
	return "null" // TODO: Prompt
}

func promptForProtoPath(ctx context.Context, reader *bufio.Reader, current string) string {
	// TODO: Stop assuming these will be in "$GOPATH/src/github.com/vendasta/vendastaapis"
	warnf("PMCLI will use %s as the root directory for protofiles\n", mockServerSource)
	defaultValue := current
	if current == "" {
		defaultValue = getDefaultProtoPath(ctx)
	}
	return prompt(reader, "Enter the path for your API proto file", defaultValue)
}

func getDefaultProtoPath(ctx context.Context) string {
	wd, err := os.Getwd()
	if err != nil {
		logging.Warningf(ctx, "Failed to guess protofile.  User will have to choose. (%s)", err.Error())
		return ""
	}
	parts := strings.Split(wd, string(os.PathSeparator))
	return fmt.Sprintf("%s/v1/api.proto", parts[len(parts)-1])
}

func prompt(reader *bufio.Reader, promptMsg string, defaultValue string) string {
	msg := promptMsg
	if defaultValue != "" {
		msg = fmt.Sprintf("%s [default %s]", msg, defaultValue)
	}
	fmt.Printf("%s: ", msg)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	if text == "" {
		return defaultValue
	}
	return text
}

func warnf(message string, args ...interface{}) {
	fmt.Printf("\033[1;36m%s\033[0m", fmt.Sprintf(message, args...))
}

func init() {
	RootCmd.AddCommand(generateCmd)
}
