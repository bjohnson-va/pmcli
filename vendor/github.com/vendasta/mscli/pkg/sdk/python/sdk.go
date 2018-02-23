package python

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/vendasta/gosdks/util"
	"github.com/vendasta/mscli/pkg/docker"
	mscliio "github.com/vendasta/mscli/pkg/io"
	"github.com/vendasta/mscli/pkg/sdk/base"
	"github.com/vendasta/mscli/pkg/sdk/python/templates"
	"github.com/vendasta/mscli/pkg/spec"
	"golang.org/x/net/context"
)

// DockerCmdSet is a Set of Docker Cmds
type DockerCmdSet []string

// SDKGenerator generates a python SDK from proto files
type SDKGenerator struct {
	base.SDKGeneratorUtils
	debug         bool
	name          string
	allEnvConfigs []spec.EnvironmentConfig
	repoURL       string
}

//NewSDKGenerator returns an instance of SDKGenerator
func NewSDKGenerator(spec spec.MicroserviceConfig, protoSourceDir string) *SDKGenerator {
	return &SDKGenerator{
		&base.SDKGeneratorUtilsImpl{
			Language:       base.LanguagePython,
			Debug:          spec.Debug,
			ProtoPaths:     spec.GetSDKProtoPaths(),
			ProtoSourceDir: protoSourceDir,
			Name:           spec.Name,
		},
		spec.Debug,
		spec.Name,
		spec.Environments,
		spec.RepoURL,
	}
}

// GenerateProtosForSDK runs docker commands to compile the proto files
func (psdk *SDKGenerator) GenerateProtosForSDK() error {
	var err error
	var wd string
	if wd, err = os.Getwd(); err != nil {
		return fmt.Errorf("error getting work directory: %s", err.Error())
	}

	files := psdk.ProtoFileLocations()
	if psdk.debug {
		filenames := ""
		for _, f := range files {
			filenames += f + " "
		}
		fmt.Printf("Found the proto file to be compiled: %q\n", strings.Trim(filenames, " "))
		fmt.Printf("Using working directory: %s\n", wd)
	}

	var cmds []DockerCmdSet
	for _, f := range files {
		cs, err := psdk.GetProtocSdkGenCommands(f)
		if err != nil {
			return err
		}
		cmds = append(cmds, cs...)
	}

	docker.Login()
	for _, c := range cmds {
		if psdk.debug {
			fmt.Printf("running command: docker %q\n", c)
		}

		cmd := exec.CommandContext(context.Background(), "docker", c...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err = cmd.Start(); err != nil {
			return fmt.Errorf("error starting protoc: %s", err.Error())
		}
		if err = cmd.Wait(); err != nil {
			return fmt.Errorf("there were errors compiling protos: %s", err.Error())
		}
	}
	return nil
}

// GetProtocSdkGenCommands formats the docker commands to run to compile the protos
func (psdk *SDKGenerator) GetProtocSdkGenCommands(file string) ([]DockerCmdSet, error) {
	sdkDirectory := psdk.SDKDir()
	wd, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("error getting work directory: %s", err.Error())
	}

	protoPath := "/src/vendor/github.com/vendasta/gosdks/pb"

	containerFilePathParts := strings.Split(file, "./vendor/github.com/vendasta/gosdks/pb/")
	containerFilePath := fmt.Sprintf("%s%s", protoPath, containerFilePathParts[len(containerFilePathParts)-1])

	protocContainer := "gcr.io/repcore-prod/protoc-python:latest"
	volumeMount := fmt.Sprintf("%s:/src", wd)

	grpcCmd := append([]string{"run", "--rm",
		"-v",
		volumeMount,
		protocContainer,
		"-I/usr/local/protoc/include/",
		fmt.Sprintf("--proto_path=%s", protoPath),
		fmt.Sprintf("--python_out=%s/_generated/grpc", sdkDirectory),
		fmt.Sprintf("--grpc_python_out=%s/_generated/grpc", sdkDirectory)},
		containerFilePath)
	os.MkdirAll(fmt.Sprintf("%s/_generated/grpc", sdkDirectory), 0777)

	protoCmd := append([]string{"run", "--rm",
		"-v",
		volumeMount,
		protocContainer,
		"-I/usr/local/protoc/include/",
		fmt.Sprintf("--proto_path=%s", protoPath),
		fmt.Sprintf("--python_out=%s/_generated/proto", sdkDirectory)},
		containerFilePath)
	os.MkdirAll(fmt.Sprintf("%s/_generated/proto", sdkDirectory), 0777)
	return []DockerCmdSet{grpcCmd, protoCmd}, nil
}

// GenerateLanguageExtras augments what was generated from compiling the protos so the sdk is easier to use
func (psdk *SDKGenerator) GenerateLanguageExtras(bootstrap bool) error {
	// Due to changes to pathing, we need to make sure we have the new setup.py.
	err := psdk.GenerateSetup()
	if err != nil {
		return err
	}
	if bootstrap {
		err = psdk.GenerateSDKInitFile()
		if err != nil {
			return err
		}
		err = psdk.GenerateVersionFile()
		if err != nil {
			return err
		}

		err = psdk.GenerateVirtualEnvReadme()
		if err != nil {
			return err
		}
		err = psdk.GenerateConfig()
		if err != nil {
			return err
		}
		err = psdk.GenerateReadmeFile()
		if err != nil {
			return err
		}
	}
	return nil
}

// GenerateVersionFile Generate the version file for the setup.py
func (psdk *SDKGenerator) GenerateVersionFile() error {
	data := templates.VersionTemplate{
		Version: "0.1.0",
	}

	return mscliio.CreateTemplatedFile(fmt.Sprintf("%s/VERSION", psdk.SDKDir()), data, templates.SDKVersionTemplate)
}

// GenerateReadmeFile Generate the README.md for the sdk
func (psdk *SDKGenerator) GenerateReadmeFile() error {
	data := templates.ReadmeTemplate{
		Version: "0.1.0",
	}

	return mscliio.CreateTemplatedFile(fmt.Sprintf("%s/README.md", psdk.SDKDir()), data, templates.SDKReadmeTemplate)
}

// GenerateSetup creates a setup.py file so the sdk can be packaged and published with pip
func (psdk *SDKGenerator) GenerateSetup() error {
	data := templates.SdkTemplate{
		Name:       psdk.name,
		SDKName:    psdk.SDKDirName(),
		URL:        psdk.repoURL,
		PythonName: strings.Replace(util.ToSnakeCase(psdk.name), "-", "_", -1),
		Version:    psdk.ProtoVersions()[0],
	}
	return mscliio.CreateTemplatedFile(fmt.Sprintf("%s/src/setup.py", psdk.SDKRootDir()), data, templates.SDKSetupTemplate)
}

// GenerateVirtualEnvReadme generates a readme with instructions on how to create a virtual env for your sdk
func (psdk *SDKGenerator) GenerateVirtualEnvReadme() error {
	data := templates.VirtualEnvTemplate{
		Name: psdk.name,
	}

	return mscliio.CreateTemplatedFile(fmt.Sprintf("%s/README.md", psdk.SDKRootDir()), data, templates.SDKVirtualEnvTemplate)
}

// GenerateConfig generates a config file for environment specific variables
func (psdk *SDKGenerator) GenerateConfig() error {
	envs := psdk.GenerateEnvironmentDataForPythonConfig()

	a := templates.Config{
		Environments: envs,
	}
	return mscliio.CreateTemplatedFile(fmt.Sprintf("%s/config.py", psdk.SDKDir()), a, templates.ConfigTemplate)
}

// GenerateSDKInitFile generates the init file for the sdk package
func (psdk *SDKGenerator) GenerateSDKInitFile() error {
	return mscliio.CreateTemplatedFile(fmt.Sprintf("%s/__init__.py", psdk.SDKDir()), nil, templates.SDKInitTemplate)
}

// GenerateEnvironmentDataForPythonConfig generates the data for the config file
func (psdk *SDKGenerator) GenerateEnvironmentDataForPythonConfig() []templates.Environment {
	var envs []templates.Environment
	for _, env := range psdk.allEnvConfigs {
		serviceAccount := ""
		for _, podEnv := range env.PodConfig.PodEnv {
			if podEnv.Key == "SERVICE_ACCOUNT" {
				serviceAccount = podEnv.Value
			}
		}
		if strings.ToUpper(env.Name) == "LOCAL" {
			envs = append(envs, templates.Environment{
				Host:           fmt.Sprintf("http://%s", env.Network.GRPCHost),
				Scope:          fmt.Sprintf("https://%s", env.Network.GRPCHost),
				ServiceAccount: serviceAccount,
				URL:            fmt.Sprintf("http://%s", env.Network.GRPCHost),
				Environment:    env.Name,
			})
		} else {
			envs = append(envs, templates.Environment{
				Host:           fmt.Sprintf("%s:443", env.Network.GRPCHost),
				Scope:          fmt.Sprintf("https://%s", env.Network.GRPCHost),
				ServiceAccount: serviceAccount,
				URL:            fmt.Sprintf("https://%s", env.Network.GRPCHost),
				Environment:    env.Name,
			})
		}
	}

	return envs
}
