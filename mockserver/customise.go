package mockserver

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/bjohnson-va/pmcli/config"
)

func showCustomizeEndpointsPrompts(reader *bufio.Reader, endpoints []string, updater ServerUpdater) {
	options := buildCustomizeMenuOptions(reader, endpoints, updater)
	for {
		fmt.Println(options)
		fmt.Printf("Choose an option: ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		num, err := strconv.Atoi(text)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		c := options[num-1]
		c.Fn()
		if c.ExitAfter {
			break
		}
	}
}

func buildCustomizeMenuOptions(
	reader *bufio.Reader, endpoints []string, updater ServerUpdater,
) menuOptions {
	opts := make(menuOptions, len(endpoints)+1)
	for i, e := range endpoints {
		opts[i] = menuOption{
			Name: e,
			Fn: func() {
				showCustomizeSpecificEndpointPrompts(reader, e, updater)
			},
			ExitAfter: false,
		}
	}
	opts[len(endpoints)] = menuOption{
		Name: "Back To Main",
		Fn: func() {
			fmt.Println("Goodbye 👋")
		},
		ExitAfter: true,
	}
	return opts
}

func showCustomizeSpecificEndpointPrompts(
	reader *bufio.Reader, endpoint string, updater ServerUpdater,
) {
	options := buildCustomizeSpecificEndpointMenuOptions(reader, endpoint, updater)
	for {
		fmt.Println(options)
		fmt.Printf("Choose an option: ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		num, err := strconv.Atoi(text)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		c := options[num-1]
		c.Fn()
		if c.ExitAfter {
			break
		}
	}
}

func buildCustomizeSpecificEndpointMenuOptions(
	reader *bufio.Reader, endpoint string, updater ServerUpdater,
) menuOptions {
	return []menuOption{
		{
			Name: "Set response delay",
			Fn: func() {
				setResponseDelay(reader, endpoint, updater)
			},
		},
		{
			Name: "Quit",
			Fn: func() {
				fmt.Println("Goodbye 👋")
			},
			ExitAfter: true,
		},
	}
}

func setResponseDelay(reader *bufio.Reader, endpoint string, updater ServerUpdater) error {
	fmt.Printf("How many seconds to delay? ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	secs, err := strconv.Atoi(text)
	if err != nil {
		return fmt.Errorf("failed to parse response: %s", err.Error())
	}
	err = updater.UpdateAndRestart(config.File{
		ConfigMap: config.Map{
			Instructions: map[string]config.RPCInstructions{
				endpoint: {
					DelaySecs: secs,
				},
			},
		},
	})
	if err != nil {
		return fmt.Errorf("failed to update config and restart server: %s", err.Error())
	}
	return nil
}
