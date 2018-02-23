package config

import "os"

//Env is an enum that represents the spectrum of possible environments can be configured to run on
type Env int64

//The possible values of Env
const (
	Prod Env = iota
	Demo
	Test
	Local
	Internal
)

// CurEnv returns the environment for this instance based on environment variables from deployment config
//
// This looks at an OS environment variable named "ENVIRONMENT" that should be set by MSCLI or whichever deployment process you use.
// Valid values for "ENVIRONMENT" include local, test, demo, prod
func CurEnv() Env {
	v := getEnv("ENVIRONMENT", "local")
	return GetEnv(v)
}

// GetEnv returns the environment enum for the environment string passed through
func GetEnv(env string) Env {
	if env == "local" {
		return Local
	} else if env == "internal" {
		return Internal
	} else if env == "test" {
		return Test
	} else if env == "demo" {
		return Demo
	} else if env == "prod" || env == "production" {
		return Prod
	}
	panic("Unable to determine environment.")
}

// Name returns a string format of the environment.
func (e Env) Name() string {
	switch e {
	case Prod:
		return "prod"
	case Demo:
		return "demo"
	case Test:
		return "test"
	default:
		return "local"
	}
}

// GetEnv returns the value of the environment variable specified by key, or fallback if no value is found
func getEnv(key string, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
