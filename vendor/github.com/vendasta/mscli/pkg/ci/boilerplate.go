package ci

import (
	"fmt"
	mscliio "github.com/vendasta/mscli/pkg/io"
	"github.com/vendasta/mscli/pkg/spec"
)

// CreateBoilerplate handles the creation of all the files and directories needed to make sure that
//a microservice can be run on Jenkins (currently) or a suitable continuous integration service
func CreateBoilerplate(spec spec.MicroserviceFile) error {
	if spec.Microservice.Debug {
		fmt.Println("Creating Jenkinsfile")
	}
	return mscliio.WriteFile(spec, JenkinsfileTemplate, "./Jenkinsfile", "[[", "]]")
}
