package base

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

//SDKGeneratorUtils defines utility functions required for generating protos
type SDKGeneratorUtils interface {
	SDKDir() string
	SDKDirName() string
	SDKRootDir() string
	FileSystemSafeString(name string) string
	ProtoFileLocations() []string
	GetProtoFilesInDirectory(directory string) ([]string, error)
	ProtoVersions() []string
}

//SDKGenerator generates the sdk from proto files
type SDKGenerator interface {
	SDKGeneratorUtils
	GenerateProtosForSDK() error
	GenerateLanguageExtras(bootstrap bool) error
}

//SDKGeneratorUtilsImpl provides utility functions required for generating protos
type SDKGeneratorUtilsImpl struct {
	Language       Language
	Debug          bool
	Name           string
	ProtoPaths     []string
	ProtoSourceDir string
}

//SDKDirName returns the name of the sdk directory
func (sdk *SDKGeneratorUtilsImpl) SDKDirName() string {
	return fmt.Sprintf("%s_sdk", sdk.FileSystemSafeString(sdk.Name))
}

//SDKDir returns the path to the sdk directory
func (sdk *SDKGeneratorUtilsImpl) SDKDir() string {
	return fmt.Sprintf("%s/src/%s", sdk.SDKRootDir(), sdk.SDKDirName())
}

//SDKRootDir returns the root of the path to the language specific sdk
func (sdk *SDKGeneratorUtilsImpl) SDKRootDir() string {
	return fmt.Sprintf("sdks/%s", sdk.Language)
}

//FileSystemSafeString replaces - with _, not really sure why
func (sdk *SDKGeneratorUtilsImpl) FileSystemSafeString(name string) string {
	return strings.Replace(name, "-", "_", -1)
}

//ProtoFileLocations returns the location of where to find the proto files
//TODO: this path should be standardized throughout any tools that use protos.
//This is also a source of headache if gosdks isn't up to date
func (sdk *SDKGeneratorUtilsImpl) ProtoFileLocations() []string {
	locations := []string{}
	for _, p := range sdk.ProtoPaths {
		fmt.Println(p)
		locations = append(locations, fmt.Sprintf("%s/%s", sdk.ProtoSourceDir, p))
	}
	return locations
}

func (sdk *SDKGeneratorUtilsImpl) ProtoVersions() []string {
	protoVersions := map[string]bool{}
	for _, p := range sdk.ProtoPaths {
		protoVersions[convertFilePathToVersion(p)] = true
	}
	var keys []string
	for k := range protoVersions {
		keys = append(keys, k)
	}
	return keys
}

func convertFilePathToVersion(filePath string) string {
	directories := strings.Split(filePath, "/")
	if len(directories) >= 2 {
		return directories[len(directories)-2]
	}
	return "v1"
}

//GetProtoFilesInDirectory returns a list of files in the given directory that have a .proto extension
func (sdk *SDKGeneratorUtilsImpl) GetProtoFilesInDirectory(directory string) ([]string, error) {
	protoFiles := []string{}

	files, err := ioutil.ReadDir(directory)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".proto") {
			protoFiles = append(protoFiles, filepath.Join(directory, file.Name()))
		}
	}
	return protoFiles, nil
}
