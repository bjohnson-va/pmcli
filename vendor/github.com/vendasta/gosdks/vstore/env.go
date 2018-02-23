package vstore

import "os"

//env is an enum that represents the spectrum of possible environments VStore can be configured to run on
type env int64

//The possible values of env
const (
	Prod env = iota
	Demo
	Test
	Local
	Internal
)

// Env returns the vstore environment for this instance based on environment variables from deployment config
// This looks at an OS environment variable named "ENVIRONMENT" that should be set by MSCLI or whichever deployment process you use.
// Valid values for "ENVIRONMENT" include local, test, demo, prod
func Env() *env {
	v := getEnv("ENVIRONMENT", "local")
	var e env
	if v == "local" {
		e = env(Local)
	} else if v == "internal" {
		e = env(Internal)
	} else if v == "test" {
		e = env(Test)
	} else if v == "demo" {
		e = env(Demo)
	} else if v == "prod" {
		e = env(Prod)
	} else {
		panic("Unable to determine environment.")
	}
	return &e
}

// GetEnv returns the value of the environment variable specified by key, or fallback if no value is found
func getEnv(key string, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
