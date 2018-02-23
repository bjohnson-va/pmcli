package java

import (
	"github.com/vendasta/mscli/pkg/sdk/base"
	"github.com/vendasta/mscli/pkg/spec"
)

//SDKGenerator generates the sdk from proto files
type SDKGenerator struct {
	base.SDKGeneratorUtils
}

//NewSDKGenerator returns an instance of SDKGenerator
func NewSDKGenerator(spec spec.MicroserviceConfig, protoSourceDir string) *SDKGenerator {
	return &SDKGenerator{
		&base.SDKGeneratorUtilsImpl{
			Language:       base.LanguageJava,
			Debug:          spec.Debug,
			ProtoPaths:     spec.GetSDKProtoPaths(),
			ProtoSourceDir: protoSourceDir,
			Name:           spec.Name,
		},
	}
}

// GenerateProtosForSDK - Java SDK gen doesn't need this at the moment
func (sg *SDKGenerator) GenerateProtosForSDK() error {
	return nil
}

// GenerateLanguageExtras - No extras at this tiume
func (sg *SDKGenerator) GenerateLanguageExtras(bootstrap bool) error {
	return nil
}
