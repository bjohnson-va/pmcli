package typescript

import (
	"fmt"
	"os"

	mscliio "github.com/vendasta/mscli/pkg/io"
	"github.com/vendasta/mscli/pkg/sdk/base"
	"github.com/vendasta/mscli/pkg/sdk/typescript/templates"
	"github.com/vendasta/mscli/pkg/spec"
	"github.com/vendasta/mscli/pkg/utils"
)

// SDKGenerator generates a typescript SDK from proto files
type SDKGenerator struct {
	base.SDKGeneratorUtils
	name          string
	allEnvConfigs []spec.EnvironmentConfig
}

// NewSDKGenerator returns an instance of SDKGenerator
func NewSDKGenerator(spec spec.MicroserviceConfig, protoSourceDir string) *SDKGenerator {
	return &SDKGenerator{
		&base.SDKGeneratorUtilsImpl{
			Language:       base.LanguageTypescript,
			Debug:          spec.Debug,
			ProtoPaths:     spec.GetSDKProtoPaths(),
			ProtoSourceDir: protoSourceDir,
			Name:           spec.Name,
		},
		spec.Name,
		spec.Environments,
	}
}

// GenerateProtosForSDK Typescript SDK doesn't need this
func (tsdk *SDKGenerator) GenerateProtosForSDK() error {
	return nil
}

// GenerateLanguageExtras augments what was generated from compiling the protos so the sdk is easier to use
func (tsdk *SDKGenerator) GenerateLanguageExtras(bootstrap bool) error {
	os.MkdirAll(fmt.Sprintf("%s/_generated/", tsdk.SDKDir()), os.ModePerm)
	err := tsdk.GenerateHostService()
	if err != nil {
		return err
	}
	err = tsdk.GenerateInternalIndex()
	if err != nil {
		return err
	}
	if bootstrap {
		err = tsdk.GenerateNPMIgnore()
		if err != nil {
			return err
		}
		err = tsdk.GenerateChangelog()
		if err != nil {
			return err
		}
		err = tsdk.GenerateReadme()
		if err != nil {
			return err
		}
	}
	return nil
}

// GenerateNPMIgnore Generates the npmignore file
func (tsdk *SDKGenerator) GenerateNPMIgnore() error {
	return mscliio.CreateTemplatedFile(fmt.Sprintf("%s/.npmignore", tsdk.SDKDir()), nil, templates.NPMIgnoreTemplate)
}

// GenerateChangelog Generates the changelog
func (tsdk *SDKGenerator) GenerateChangelog() error {
	data := templates.ChangelogModel{
		Version: "0.1.0",
	}

	return mscliio.PrependOrCreateTemplateToFile(fmt.Sprintf("%s/CHANGELOG.md", tsdk.SDKDir()), data, templates.ChangelogTemplate)
}

// GenerateReadme Generates the README.md
func (tsdk *SDKGenerator) GenerateReadme() error {
	data := templates.ReadmeModel{
		Name: tsdk.name,
	}

	return mscliio.CreateTemplatedFile(fmt.Sprintf("%s/README.md", tsdk.SDKDir()), data, templates.ReadmeTemplate)
}

// GenerateHostService Generates the /_generated/host.service.ts
func (tsdk *SDKGenerator) GenerateHostService() error {
	data := templates.HostServiceModel{}

	for _, envConfig := range tsdk.allEnvConfigs {
		switch utils.EnvFromString(envConfig.Name) {
		case utils.Local:
			data.LocalHost = tsdk.PreferredHostFromEnvConfig(envConfig)
			data.LocalHttpsHost = tsdk.PreferredHttpsHostFromEnvConfig(envConfig)
		case utils.Test:
			data.TestHost = tsdk.PreferredHostFromEnvConfig(envConfig)
			data.TestHttpsHost = tsdk.PreferredHttpsHostFromEnvConfig(envConfig)
		case utils.Demo:
			data.DemoHost = tsdk.PreferredHostFromEnvConfig(envConfig)
			data.DemoHttpsHost = tsdk.PreferredHttpsHostFromEnvConfig(envConfig)
		case utils.Prod:
			data.ProdHost = tsdk.PreferredHostFromEnvConfig(envConfig)
			data.ProdHttpsHost = tsdk.PreferredHttpsHostFromEnvConfig(envConfig)
		}
	}

	return mscliio.CreateTemplatedFile(
		fmt.Sprintf("%s/_generated/host.service.ts", tsdk.SDKDir()),
		data, templates.HostServiceTemplate,
	)
}

// PreferredHostFromEnvConfig returns the preferred host from the microservice's environment configuration
func (tsdk *SDKGenerator) PreferredHostFromEnvConfig(ec spec.EnvironmentConfig) string {
	if ec.SecondarySSLConfig != nil {
		return ec.SecondarySSLConfig.Host
	}
	return ec.GRPCHost
}

// PreferredHttpsHostFromEnvConfig returns the preferred https host from the microservice's environment configuration
func (tsdk *SDKGenerator) PreferredHttpsHostFromEnvConfig(ec spec.EnvironmentConfig) string {
	if ec.SecondarySSLConfig != nil && ec.SecondarySSLConfig.HTTPSHost != "" {
		return ec.SecondarySSLConfig.HTTPSHost
	}
	return ec.HTTPSHost
}

// GenerateInternalIndex Generates the /_generated/index.ts
func (tsdk *SDKGenerator) GenerateInternalIndex() error {
	return mscliio.CreateTemplatedFile(
		fmt.Sprintf("%s/_generated/index.ts", tsdk.SDKDir()),
		nil, templates.GeneratedIndexTemplate,
	)
}
