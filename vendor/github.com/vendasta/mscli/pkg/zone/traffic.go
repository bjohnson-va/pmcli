package zone

import (
	"fmt"

	"k8s.io/api/extensions/v1beta1"

	"github.com/vendasta/mscli/pkg/k8s"
	"github.com/vendasta/mscli/pkg/kubernetes"
	"github.com/vendasta/mscli/pkg/spec"
	"github.com/vendasta/mscli/pkg/utils"
)

// SetTraffic allows traffic to be served or restricted from the specified zone
func SetTraffic(spec spec.MicroserviceFile, env utils.Environment, zone string, enabled bool) error {
	zones, err := List(spec, env)
	if err != nil {
		return err
	}
	envConfig, err := spec.Microservice.GetEnv(env)
	if err != nil {
		return err
	}
	exists := false
	for _, z := range zones {
		if z == zone {
			exists = true
			break
		}
	}
	if !exists {
		return fmt.Errorf("%s is not a valid zone", zone)
	}
	c, err := kubernetes.GetK8sClientSet(spec, env)
	if err != nil {
		return err
	}
	podNumber := int32(0)
	if enabled {
		config, err := spec.Microservice.GetEnv(env)
		if err != nil {
			return err
		}
		podNumber = config.MinReplicas
	}
	api := k8s.NewK8SApi(c)
	deployments, err := api.ListDeployments(envConfig.K8sNamespace, env.String())
	if err != nil {
		return err
	}
	var deployment *v1beta1.Deployment
	for _, d := range deployments {
		if d.Labels["zone"] == zone {
			deployment = &d
			break
		}
	}
	if deployment == nil {
		return fmt.Errorf("could not find deployment for zone %s", zone)
	}
	return api.SetPodNumber(envConfig.K8sNamespace, deployment.Name, podNumber)
}
