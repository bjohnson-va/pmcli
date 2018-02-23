package utils

import (
	"fmt"
	"os/exec"
	"strings"
)

// AssertDepIsInstalled asserts that dep is installed and on your path.
func AssertDepIsInstalled() error {
	out, _ := exec.Command("which", "dep").Output()
	if len(out) == 0 || strings.Contains(string(out), "not found") {
		return fmt.Errorf("dep was not found on your path. Install dep by running `go get -u github.com/golang/dep/cmd/dep`")
	}
	return nil
}
