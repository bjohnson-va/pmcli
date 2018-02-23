package utils

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"os/exec"

	version "github.com/hashicorp/go-version"
)

//GCloudVersions stores the relevant versions of gcloud and components within
type GCloudVersions struct {
	SDK  string `json:"Google Cloud SDK"`
	Beta string `json:"beta"`
}

//GetGCloudVersion queries the gcloud cli for versions
func GetGCloudVersion() GCloudVersions {
	var err error
	command := []string{
		"version",
		"--format=json",
	}
	cmd := exec.CommandContext(context.Background(), "gcloud", command...)
	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = os.Stderr
	if err = cmd.Start(); err != nil {
		log.Fatalf("Error starting gcloud: %s", err.Error())
	}
	versions := GCloudVersions{}
	if err = json.NewDecoder(stdout).Decode(&versions); err != nil {
		log.Fatalf("Error decoding versions: %s", err)
	}
	if err = cmd.Wait(); err != nil {
		log.Fatalf("There were errors deployting endpoints: %s", err.Error())
	}

	return versions
}

//AssertGCloudVersion does a semantic version check and exits if a version check fails
func AssertGCloudVersion(constraint string) {
	versions := GetGCloudVersion()
	v1, err := version.NewVersion(versions.SDK)
	if err != nil {
		log.Fatalf("Error parsing version of Gcloud SDK: %#v", versions)
	}
	// Constraints example.
	constraints, err := version.NewConstraint(constraint)
	if !constraints.Check(v1) {
		log.Fatalf("Error, gcloud version %s does not satisfy constraint %s", v1, constraint)
	}
}

//AssertGCloudBetaVersion does a semantic version check and exits if a version check fails
func AssertGCloudBetaVersion(constraint string) {
	versions := GetGCloudVersion()
	if versions.Beta == "" {
		log.Fatalf("Gcloud Beta component not installed.")
	}
	v1, err := version.NewVersion(versions.Beta)
	if err != nil {
		log.Fatalf("Error parsing version of Gcloud Beta: %#v", versions)
	}
	// Constraints example.
	constraints, err := version.NewConstraint(constraint)
	if !constraints.Check(v1) {
		log.Fatalf("Error, gcloud beta version %s does not satisfy constraint %s", v1, constraint)
	}
}
