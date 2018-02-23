package protos

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"

	"github.com/vendasta/mscli/pkg/spec"
)

// protoFilesFromDir builds a list of proto files in the specified directory
func protoFilesFromDir(baseProtoDir string, protoSourceDir string) []string {
	fullPath := path.Join(baseProtoDir, protoSourceDir)
	validateProtoDir(fullPath)

	var err error

	var files, paths []string
	if paths, err = filepath.Glob(fmt.Sprintf("%s/*.proto", fullPath)); err != nil {
		log.Fatalf("Error expanding glob: %s", err.Error())
	}
	for i := range paths {
		r, _ := filepath.Rel(fullPath, paths[i])
		files = append(files, fmt.Sprintf("%s/%s", protoSourceDir, r))
	}
	return files
}

// protoFilesFromDirs builds a list of the proto files for each proto source directory provided
func protoFilesFromDirs(baseProtoDir string, protoSourceDirs []string) []string {
	files := []string{}
	for _, dir := range protoSourceDirs {
		f := protoFilesFromDir(baseProtoDir, dir)
		files = append(files, f...)
	}
	return files
}

// Build build from protos
func Build(spec spec.MicroserviceFile, baseProtoDir string, protoSourceDirs []string, workingDir string) error {
	files := protoFilesFromDirs(baseProtoDir, protoSourceDirs)
	if spec.Microservice.Debug {
		fmt.Printf("Found the following proto files to be compiled: %q\n", files)
	}

	command := append([]string{
		"run",
		"--rm",
		"-v", fmt.Sprintf("%s:/src", baseProtoDir),
		"-v", fmt.Sprintf("%s/pb:/dest", workingDir),
		"gcr.io/repcore-prod/protoc-go",
		"--include_imports",
		"--go_out=plugins=grpc,import_path=pb:/dest",
		"--include_source_info",
		"--descriptor_set_out=/dest/descriptor.pb",
		"-I=.",
	})
	command = append(command, fmt.Sprintf("--proto_path=."))
	command = append(command, files...)

	if spec.Microservice.Debug {
		fmt.Printf("Running Command: docker %v\n", command)
	}
	var err error
	cmd := exec.CommandContext(context.Background(), "docker", command...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err = cmd.Start(); err != nil {
		return fmt.Errorf("error starting protoc: %s", err.Error())
	}
	if err = cmd.Wait(); err != nil {
		return fmt.Errorf("there were errors compiling protos: %s", err.Error())
	}
	return nil
}

func validateProtoDir(protoSourceDir string) {
	info, err := os.Stat(protoSourceDir)
	if err != nil {
		log.Fatalf("Error validating the proto path: %v", err)
	}
	if !info.IsDir() {
		log.Fatal("The path to the protos must be a directory, not a file")
	}
}
