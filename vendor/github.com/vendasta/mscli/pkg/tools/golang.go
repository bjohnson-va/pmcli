package tools

import (
	"fmt"
	"github.com/vendasta/mscli/pkg/utils"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

var tools map[string]string

func init() {
	tools = map[string]string{
		"gocode":     "github.com/nsf/gocode",
		"gopkgs":     "github.com/tpng/gopkgs",
		"go-outline": "github.com/lukehoban/go-outline",
		"go-symbols": "github.com/newhook/go-symbols",
		"guru":       "golang.org/x/tools/cmd/guru",
		"gorename":   "golang.org/x/tools/cmd/gorename",
		"goreturns":  "sourcegraph.com/sqs/goreturns",
		"golint":     "github.com/golang/lint/golint",
		"dlv":        "github.com/derekparker/delve/cmd/dlv",
		"godef":      "github.com/rogpeppe/godef",
		"godoc":      "golang.org/x/tools/cmd/godoc",
		"dep":        "github.com/golang/dep/cmd/dep",
	}
}

// InstallGoTools ensures that everything you need for local golang development is installed.
func InstallGoTools() error {
	err := verifyGoVersion("go1.9")
	if err != nil {
		return err
	}
	err = verifyGoEnvVariables()
	if err != nil {
		return err
	}
	return verifyAllToolsInstalled(tools)
}

func verifyGoVersion(expected string) error {
	if strings.HasPrefix(runtime.Version(), expected) {
		return nil
	}
	return fmt.Errorf("version must start with: %v", expected)
}

func verifyGoEnvVariables() error {
	home := os.Getenv("HOME")
	gopath := os.Getenv("GOPATH")
	gobin := os.Getenv("GOBIN")
	path := os.Getenv("PATH")
	if gopath == "" {
		return fmt.Errorf("your GOPATH env variable must be set. Recommend `$HOME/go`")
	}
	if !strings.HasPrefix(gopath, home) {
		fmt.Printf("WARNING: Your GOPATH (%s) is not under your HOME directory (%s). It is recommended that your GOPATH is set to `$HOME/go`", gopath, home)
	}
	if gobin == "" {
		return fmt.Errorf("your GOBIN env variable must be set. Recommend `$GOPATH/bin`")
	}
	if !strings.HasPrefix(gobin, gopath) {
		fmt.Printf("WARNING: Your GOBIN (%s) is not under your GOPATH (%s). It is recommended that your GOBIN is set to `$GOPATH/bin`", gobin, gopath)
	}
	if !strings.Contains(path, gobin) {
		fmt.Printf("WARNING: Your GOBIN (%s) is not in your PATH (%s). It is recommended that your GOBIN is in your path for easy tool execution", gobin, path)
	}
	return nil
}

func verifyToolInstalled(tool string) bool {
	out, _ := exec.Command("which", tool).Output()
	if len(out) == 0 {
		fmt.Printf("WARNING: %s is not installed or not found on your PATH", tool)
		return false
	}
	return true
}

func verifyAllToolsInstalled(tools map[string]string) error {
	all := false
	for k, v := range tools {
		if !verifyToolInstalled(k) {
			answer := ""
			if !all {
				answer = utils.GetUserInput(fmt.Sprintf("\nInstall %s now?: yes/no/all/none ", k), "")
			} else {
				answer = "y"
			}
			switch answer {
			case "a", "all":
				all = true
				fallthrough
			case "y", "yes":
				fmt.Printf("\nInstalling %s...\n", k)
				installTool(v)
				break
			case "n", "no":
				break
			case "none":
				return nil
			default:
				return fmt.Errorf("unrecognized answer")
			}
		}
	}
	return nil
}

func installTool(path string) error {
	_, err := exec.Command("go", "get", "-u", "-v", path).Output()
	return err
}
