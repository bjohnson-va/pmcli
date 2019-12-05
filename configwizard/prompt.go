package configwizard

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
)

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

func Prompt(ctx context.Context, rootDir string) MockServerJson {
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

	safeGet := func(slice []string, index int, defaultVal string) string {
		if len(slice) > index + 1 {
			return slice[index]
		}
		return defaultVal
	}

	jsonData.Port = promptForPort(reader, jsonData.Port)
	jsonData.AllowedOrigin = promptForAllowedOrigin(reader)
	jsonData.Https = promptForHttps(reader, jsonData.Https)
	jsonData.ProtoPaths = []string{promptForProtoPath(ctx, reader, rootDir, safeGet(jsonData.ProtoPaths, 0, ""))}
	jsonData.Overrides = map[string]interface{}{}
	// TODO: Uncomment
	//jsonData.Overrides = PromptForOverrides(ctx, filepath.Join(os.Getwd(), config.FILENAME))
	jsonData.Instructions = map[string]interface{}{}
	jsonData.Exclusions = map[string]interface{}{}
	return jsonData
}

func promptForPort(reader *bufio.Reader, current int64) int64 {
	defaultPort := 28000
	if current > 0 {
		defaultPort = int(current)
	}
	userValue := prompt(reader, "Choose a port for the server", strconv.Itoa(defaultPort))
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

func promptForProtoPath(ctx context.Context, reader *bufio.Reader, rootDir string, current string) string {
	// TODO: Stop assuming these will be in "$GOPATH/src/github.com/vendasta/vendastaapis"
	warnf("PMCLI will use %s as the root directory for protofiles\n", rootDir)
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
	name := parts[len(parts)-1]
	name = strings.Replace(name, "-", "_", 1)
	return fmt.Sprintf("%s/v1/api.proto", name)
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
