package mockserver

import (
	"bufio"
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/bjohnson-va/pmcli/config"
)

func showCustomizeEndpointsPrompts(
	ctx context.Context, reader *bufio.Reader, endpoints []string,
	updater ServerUpdater,
) {
	options := buildCustomizeMenuOptions(ctx, reader, endpoints, updater)
	for {
		fmt.Println(options)
		fmt.Printf("Choose an endpoint to customize: ")
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
	ctx context.Context, reader *bufio.Reader, endpoints []string,
	updater ServerUpdater,
) menuOptions {
	opts := make(menuOptions, len(endpoints)+1)
	for i, e := range endpoints {
		opts[i] = menuOption{
			Name: e,
			Fn: func() {
				showCustomizeSpecificEndpointPrompts(ctx, reader, e, updater)
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
	ctx context.Context, reader *bufio.Reader, endpoint string,
	updater ServerUpdater,
) {
	options := buildCustomizeSpecificEndpointMenuOptions(ctx, reader, endpoint, updater)
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
		if num > len(options) {
			fmt.Printf("Invalid selection: %d\n", num)
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
	ctx context.Context, reader *bufio.Reader, endpoint string,
	updater ServerUpdater,
) menuOptions {
	return []menuOption{
		{
			Name: "Set response delay",
			Fn: func() {
				err := setResponseDelay(ctx, reader, endpoint, updater)
				if err != nil {
					fmt.Printf("Could not set response delay: %s", err.Error())
				}
			},
		},
		{
			Name: "Set response status code",
			Fn: func() {
				err := setStatusCode(ctx, reader, endpoint, updater)
				if err != nil {
					fmt.Printf("Could not set response code: %s", err.Error())
				}
			},
		},
		{
			Name:      "Back to main",
			Fn:        func() {},
			ExitAfter: true,
		},
	}
}

func setResponseDelay(ctx context.Context, reader *bufio.Reader, endpoint string, updater ServerUpdater) error {
	secs, err := getNumberFromUser(reader, "How many seconds to delay?")
	if err != nil {
		return fmt.Errorf("failed to parse response: %s", err.Error())
	}
	err = updater.UpdateAndRestart(ctx, config.Mutation{
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

func setStatusCode(ctx context.Context, reader *bufio.Reader, endpoint string, updater ServerUpdater) error {
	secs, err := getNumberFromUser(reader, "Which status code?")
	if err != nil {
		return fmt.Errorf("failed to parse response: %s", err.Error())
	}
	err = updater.UpdateAndRestart(ctx, config.Mutation{
		ConfigMap: config.Map{
			Instructions: map[string]config.RPCInstructions{
				endpoint: {
					StatusCode: secs,
				},
			},
		},
	})
	if err != nil {
		return fmt.Errorf("failed to update config and restart server: %s", err.Error())
	}
	return nil
}

func getNumberFromUser(reader *bufio.Reader, msg string) (int, error) {
	fmt.Printf("%s ", msg)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	return strconv.Atoi(text)
}
