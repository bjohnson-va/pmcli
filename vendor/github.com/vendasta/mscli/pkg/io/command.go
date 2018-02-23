package io

import (
	"bufio"
	"github.com/vendasta/mscli/pkg/utils"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

// Command interface to build and execute commands on the command line
type Command interface {
	Execute() (string, error)
}

type createFileCommand struct {
	file string
	data []byte
}

func (c *createFileCommand) Execute() error {
	log.Printf("Saving file %s", c.file)
	return ioutil.WriteFile(c.file, c.data, 0666)
}

type runCommand struct {
	cmd  string
	args []string
}

func (r runCommand) Execute() (string, error) {
	var err error
	var stdout io.ReadCloser
	cmd := exec.Command(r.cmd, r.args...)

	if stdout, err = cmd.StdoutPipe(); err != nil {
		log.Printf("Error creating stdoutpipe %s %s. Error: %s", r.cmd, r.args, err.Error())
		return "", err
	}

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	env := os.Environ()
	cmd.Env = env

	err = cmd.Start()
	if err != nil {
		log.Printf("Error starting command %s %s. Error: %s", r.cmd, r.args, err.Error())
		return "", err
	}

	reader := bufio.NewReader(stdout)
	output, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Fatalf("Error collecting minikube env %s", err.Error())
	}
	err = cmd.Wait()
	if err != nil {
		log.Printf("Error running command %s %s. Error: %s", r.cmd, r.args, err.Error())
		return "", err
	}
	return string(output), nil
}

// RunCommand Build a generic command to execute
func RunCommand(cmd string, args ...string) Command {
	return runCommand{cmd: cmd, args: args}
}

type dockerCommand struct {
	args []string
}

func (d dockerCommand) Execute() (string, error) {
	var err error
	cmd := exec.Command("docker", d.args...)
	cmd.Dir = "."
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if !utils.IsOnJenkins() {
		cmd.Env = os.Environ()
	}

	if err = cmd.Start(); err != nil {
		return "", err
	}
	if err = cmd.Wait(); err != nil {
		return "", err
	}
	return "", nil
}

// DockerCommand Builds a docker command to execute
func DockerCommand(args ...string) Command {
	return dockerCommand{args: args}
}
