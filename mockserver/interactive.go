package mockserver

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/bjohnson-va/pmcli/config"
)

type ServerUpdater interface {
	UpdateAndRestart(ctx context.Context, d config.Mutation) error
}

func showInteractivePrompts(ctx context.Context, endpoints []string, srv ServerUpdater) error {
	reader := bufio.NewReader(os.Stdin)
	options := buildMainMenuOptions(ctx, reader, endpoints, srv)
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
			os.Exit(0)
		}
	}
}

type menuOption struct {
	Name      string
	Fn        func()
	ExitAfter bool
}

type menuOptions []menuOption

func (m menuOptions) String() string {
	names := make([]string, len(m))
	for i, o := range m {
		names[i] = fmt.Sprintf("[%d] %s", i+1, o.Name)
	}
	return strings.Join(names, "\n")
}

func buildMainMenuOptions(
	ctx context.Context, reader *bufio.Reader, endpoints []string,
	srv ServerUpdater,
) menuOptions {
	return []menuOption{
		{
			Name: "Customize an endpoint's behaviour",
			Fn: func() {
				showCustomizeEndpointsPrompts(ctx, reader, endpoints, srv)
			},
		},
		{
			Name: "Open insecure chromium browser",
			Fn: func() {
				go func() {
					err := exec.Command("chromium", "--disable-web-security", "--user-data-dir=~/chromeTemp").Run()
					if err != nil {
						fmt.Printf("Failed to open chromium: %s", err.Error())
					}
				}()
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
