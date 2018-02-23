package utils

import (
	"fmt"
	"os"
	"strings"
)

//Environment is an enum to capture the target environment of the tool
type Environment int64

const (
	//Local is running on the developer's laptop
	Local Environment = iota
	//Test is for internal testing only
	Test
	//Demo is for customer demos
	Demo
	//Prod is what all of our customers use
	Prod
	//Integration is for running integration tests against a live environment (may have running stubs)
	Integration
)

//String just converts the enum to a user-friendly representation
func (e Environment) String() string {
	switch e {
	case Local:
		return "local"
	case Integration:
		return "integration"
	case Test:
		return "test"
	case Demo:
		return "demo"
	case Prod:
		return "prod"
	}
	panic("Unable to determine environment string.")
}

//EnvFromString converts a string to an enum value (should only be used when reading in configs from file)
func EnvFromString(envName string) Environment {
	envName = strings.ToLower(envName)
	switch envName {
	case "local":
		return Local
	case "test":
		return Test
	case "demo":
		return Demo
	case "prod":
		return Prod
	case "integration":
		return Integration
	}
	panic(fmt.Sprintf("Unable to determine environment from string: %s", envName))
}

func IsOnJenkins() bool {
	return os.Getenv("JENKINS_NAME") != ""
}
