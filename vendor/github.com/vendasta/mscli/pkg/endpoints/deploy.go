package endpoints

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"path"

	"github.com/spf13/viper"
	"github.com/vendasta/mscli/pkg/protos"
	"github.com/vendasta/mscli/pkg/spec"
	"github.com/vendasta/mscli/pkg/utils"
)

// DeployAndGenerate preps and deploys endpoints
func DeployAndGenerate(spec spec.MicroserviceFile, skipGeneration bool, force bool, env utils.Environment) error {
	workingDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("error getting work directory: %s", err.Error())
	}

	err = protos.Compile(spec, workingDir, path.Join(workingDir, "vendor/github.com/vendasta/gosdks/pb"))
	if err != nil {
		return err
	}

	if !skipGeneration {
		err = GenerateEnvironmentSpecificEndpoints(spec, workingDir, env)
		if err != nil {
			return err
		}
	} else if spec.Microservice.Debug {
		fmt.Println("Skipping endpoints generation")
	}

	version, err := DeployEndpoints(spec, force, env)
	if err != nil {
		return err
	}
	fmt.Printf("Version: %s\n", version)
	err = WriteEndpointsVersion(spec, env.String(), version)
	return err
}

// DeployEndpoints deploy endpoints
func DeployEndpoints(spec spec.MicroserviceFile, force bool, env utils.Environment) (string, error) {

	//Make sure gcloud is up to date
	utils.AssertGCloudVersion(">= 148.0.1")
	utils.AssertGCloudBetaVersion(">= 2016.01.12")

	var err error
	command := []string{
		"beta",
		"service-management",
		"deploy",
		"pb/descriptor.pb",
		getAPIConfigFileName(spec, env),
		"--project=repcore-prod",
	}
	if force {
		command = append(command, "--force")
	}
	cmd := exec.CommandContext(context.Background(), "gcloud", command...)
	cmd.Stdout = os.Stdout
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return "", fmt.Errorf("error creating stderr pipe: %s", err.Error())
	}
	if err = cmd.Start(); err != nil {
		return "", fmt.Errorf("error starting gcloud: %s", err.Error())
	}

	//Get the version from the output
	var bytes []byte
	bytes, err = ioutil.ReadAll(stderr)
	if err != nil {
		return "", fmt.Errorf("error reading stderr: %s", err.Error())
	}

	//Find dat gcloud command with the id
	version := ""
	curEnv, err := spec.Microservice.GetEnv(env)
	if err != nil {
		return "", err
	}
	re := regexp.MustCompile(fmt.Sprintf(`Service Configuration \[(.*)\] uploaded for service \[%s\]`, regexp.QuoteMeta(curEnv.GRPCHost)))
	lines := strings.Split(string(bytes), "\n")
	for _, line := range lines {
		matches := re.FindStringSubmatch(line)
		if len(matches) > 1 {
			version = matches[1]
		}
	}

	if version == "" {
		fmt.Printf("Version was not found in the deploy output. There may have been an error.\nOutput: %s\n", string(bytes))
	}

	if err = cmd.Wait(); err != nil {
		fmt.Printf("%s", string(bytes))
		return "", fmt.Errorf("there were errors deploying endpoints: %s", err.Error())
	}

	return version, nil
}

// WriteEndpointsVersion set the endpoints version in the microservice config
func WriteEndpointsVersion(spec spec.MicroserviceFile, env string, version string) error {
	for i, e := range spec.Microservice.Environments {
		if e.Name == env {
			spec.Microservice.Environments[i].AppConfig.EndpointsVersion = version
			break
		}
	}

	viper.Set("microservice.environments", spec.Microservice.Environments)
	viper.WriteConfig()

	if env == "local" {
		err := writeEndpointsVersionToDockerComposeFile(version)
		return err
	}
	return nil
}

// writeEndpointsVersionToDockerComposeFile sets the new endpoints version in the docker-compose file
func writeEndpointsVersionToDockerComposeFile(version string) error {
	composeFile := "docker-compose.yaml"
	endpointsVersionArg := fmt.Sprintf(`"-v%s"`, version)
	input, err := ioutil.ReadFile(composeFile)
	if err != nil {
		return fmt.Errorf("Error reading %s: %s", composeFile, err.Error())
	}
	re := regexp.MustCompile(`"-v\d{4}-\d{2}-\d{2}r\d+"`) // e.g. "-v2018-01-01r0"
	lines := strings.Split(string(input), "\n")
	for i, line := range lines {
		lines[i] = re.ReplaceAllString(line, endpointsVersionArg)
	}
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(composeFile, []byte(output), 0644)
	if err != nil {
		return fmt.Errorf("Error writing %s: %s", composeFile, err.Error())
	}
	log.Printf("Updated %s with endpoints version: %s", composeFile, version)
	return nil
}
