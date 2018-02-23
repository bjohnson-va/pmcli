package ci

import (
	"fmt"

	"github.com/vendasta/gosdks/datadogapi"
	"github.com/vendasta/mscli/pkg/docker"
	"github.com/vendasta/mscli/pkg/k8s"
	"github.com/vendasta/mscli/pkg/kubernetes"
	"github.com/vendasta/mscli/pkg/spec"
	"github.com/vendasta/mscli/pkg/utils"
)

func Deliver(specFile spec.MicroserviceFile, version string, env utils.Environment, zones []string) error {
	fmt.Printf("Deploying to Kubernetes. Version: %s\n", version)

	var tag string
	var err error
	if env != utils.Local {
		//Verify some settings
		err = spec.VerifyPodEnv(specFile, "GOOGLE_APPLICATION_CREDENTIALS", env)
		if err != nil {
			return err
		}
		err = spec.VerifyPodEnv(specFile, "SERVICE_ACCOUNT", env)
		if err != nil {
			return err
		}
		fmt.Println("Verified")
		tag = docker.DockerImageTag(specFile, version)
	} else {
		fmt.Println("building docker image")
		tag, err = docker.BuildDockerImage(specFile, version)
		if err != nil {
			return err
		}
	}

	if env == utils.Prod {
		datadogapi.PushEventToDatadog(version, specFile.Microservice.Name, env.String(), datadogapi.Info)
	}

	client, err := kubernetes.GetK8sClientSet(specFile, env)
	if err != nil {
		return err
	}

	k8sAPI := k8s.NewK8SApi(client)

	d, err := specFile.Microservice.K8S(env, tag, zones)
	if err != nil {
		return err
	}
	return k8sAPI.Apply(d...)
}
