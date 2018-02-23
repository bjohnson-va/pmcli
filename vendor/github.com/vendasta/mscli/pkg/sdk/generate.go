package sdk

import (
	"fmt"
	"os"

	"github.com/vendasta/mscli/pkg/docker"
	mscliio "github.com/vendasta/mscli/pkg/io"
	"github.com/vendasta/mscli/pkg/sdk/base"
	"github.com/vendasta/mscli/pkg/sdk/java"
	"github.com/vendasta/mscli/pkg/sdk/python"
	"github.com/vendasta/mscli/pkg/sdk/typescript"
	"github.com/vendasta/mscli/pkg/spec"
)

// dockerImageVersion Which docker image version to use for SDK GEN
const dockerImageVersion = "1.1.0"

// GenerateSDK generates protos for the specified language and places them into the correct SDK directory
func GenerateSDK(spec spec.MicroserviceFile, language base.Language, protoSourceDir string, bootstrap bool) error {
	//generate protos into the specified language
	if spec.Microservice.Debug {
		fmt.Printf("Compiling protos for source dir %s and language %s\n", protoSourceDir, language)
	}
	sdkGenerator, err := getSDKGenerator(language, spec.Microservice, protoSourceDir)
	if err != nil {
		return err
	}
	createSDKDirectory(sdkGenerator.SDKDir())
	err = sdkGenerator.GenerateProtosForSDK()
	if err != nil {
		return err
	}
	err = sdkGenerator.GenerateLanguageExtras(bootstrap)
	if err != nil {
		return err
	}

	return runGenSDK(spec.Microservice, language, sdkGenerator.SDKDir())
}

func getSDKGenerator(language base.Language, spec spec.MicroserviceConfig, protoSourceDir string) (base.SDKGenerator, error) {
	switch language {
	case base.LanguagePython:
		return python.NewSDKGenerator(spec, protoSourceDir), nil
	case base.LanguageTypescript:
		return typescript.NewSDKGenerator(spec, protoSourceDir), nil
	case base.LanguageJava:
		return java.NewSDKGenerator(spec, protoSourceDir), nil
	default:
		return nil, fmt.Errorf("language not supported %s", language)
	}
}

func createSDKDirectory(dirPath string) {
	os.MkdirAll(dirPath, os.ModePerm)
}

func runGenSDK(spec spec.MicroserviceConfig, language base.Language, sdkDir string) error {
	wd, _ := os.Getwd()

	protoPaths := []string{}
	for _, p := range spec.GetSDKProtoPaths() {
		protoPaths = append(protoPaths, fmt.Sprintf("/protos/%s", p))
	}
	c := []string{
		"run",
		"-t",
		"-v",
		fmt.Sprintf("%s/vendor/github.com/vendasta/gosdks/pb/:/protos/", wd),
		"-v",
		fmt.Sprintf("%s/%s:/%s", wd, sdkDir, sdkDir),
		fmt.Sprintf("gcr.io/repcore-prod/sdkgen:%s", dockerImageVersion),
		"--proto_path=/protos/",
		"--plugin=/usr/bin/protoc-gen-sdkgen",
		fmt.Sprintf("--sdkgen_out=%s&%s:/%s", language, spec.Name, sdkDir),
	}
	c = append(c, protoPaths...)
	if spec.Debug {
		fmt.Printf("running command: docker %q\n", c)
	}

	docker.Login()
	dc := mscliio.DockerCommand(c...)
	_, err := dc.Execute()
	if err != nil {
		return fmt.Errorf("error generating the SDK: %s", err.Error())
	}
	return nil
}
