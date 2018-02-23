package protos

import (
	"fmt"
	"github.com/vendasta/mscli/pkg/spec"
	"io/ioutil"
	"os"
	"path"
)

func CompileAndGenerateStubs(spec spec.MicroserviceFile, style StubStyle) error {
	workingDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("error getting work directory: %s", err.Error())
	}

	// we use vendastaapis ONLY for generating the initial stubs, since the service being generated will not have any vendored dependencies we can rely on
	err = Compile(spec, workingDir, path.Join(os.Getenv("GOPATH"), "src/github.com/vendasta/vendastaapis"))
	if err != nil {
		return err
	}

	if spec.Microservice.Debug {
		fmt.Printf("Generating stubs. Style: %s\n", style)
	}
	return GenerateStubs(spec, workingDir, style)
}

//GenerateStubs generates the stub implementations from the descriptor
func GenerateStubs(spec spec.MicroserviceFile, workingDir string, style StubStyle) error {
	bytes, err := ioutil.ReadFile(fmt.Sprintf("%s/pb/descriptor.pb", workingDir))
	if err != nil {
		return fmt.Errorf("error reading Proto descriptor file, protoc must have failed or something: %s", err.Error())
	}

	return GenerateStubImplementations(spec, bytes, style)
}
