package protos

import (
	"fmt"
	mscliio "github.com/vendasta/mscli/pkg/io"
	"github.com/vendasta/mscli/pkg/spec"
	"path"
)

// Compile builds the initial protos for a mscli project
func Compile(spec spec.MicroserviceFile, workingDir, baseProtoDir string) error {
	if spec.Microservice.Debug {
		fmt.Printf("%s\n", "Compiling protos")
	}
	err := mscliio.EnsureDirExists(spec, spec.Microservice.GetPathToServerDir())
	if err != nil {
		return err
	}
	err = mscliio.EnsureDirExists(spec, spec.Microservice.GetPathToCodeImplementationDir())
	if err != nil {
		return err
	}
	err = mscliio.EnsureDirExists(spec, spec.Microservice.GetPathToAPIDir())
	if err != nil {
		return err
	}

	pathToProtos := []string{}
	for _, p := range spec.Microservice.GetProtoPaths() {
		pathToProtos = append(pathToProtos, path.Dir(p.Path))
	}

	return Build(spec, baseProtoDir, pathToProtos, workingDir)
}
