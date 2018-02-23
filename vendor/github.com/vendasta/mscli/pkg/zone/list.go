package zone

import (
	"github.com/vendasta/godatatypes/set"
	"github.com/vendasta/mscli/pkg/k8s"
	"github.com/vendasta/mscli/pkg/kubernetes"
	"github.com/vendasta/mscli/pkg/spec"
	"github.com/vendasta/mscli/pkg/utils"
)

type listConfig struct {
	enabledOnly bool
}

// ListOption provides options to
type ListOption func(config *listConfig)

// EnabledOnly will only return zones which are available to serve traffic.
func EnabledOnly() ListOption {
	return func(config *listConfig) {
		config.enabledOnly = true
	}
}

// List all zones for a given microservice.
func List(spec spec.MicroserviceFile, env utils.Environment, opts ...ListOption) ([]string, error) {
	config := &listConfig{enabledOnly: false}
	for _, o := range opts {
		o(config)
	}
	if env == utils.Local {
		return nil, nil
	}
	if !config.enabledOnly {
		return listSpecZones(spec, env)
	}
	return listEnabledOnlyZones(spec, env)
}

func listSpecZones(spec spec.MicroserviceFile, env utils.Environment) ([]string, error) {
	config, err := spec.Microservice.GetEnv(env)
	if err != nil {
		return nil, err
	}
	zones := config.GetZones()
	return zones, nil
}

func listEnabledOnlyZones(spec spec.MicroserviceFile, env utils.Environment) ([]string, error) {
	c, err := kubernetes.GetK8sClientSet(spec, env)
	if err != nil {
		return nil, err
	}
	envConfig, err := spec.Microservice.GetEnv(env)
	if err != nil {
		return nil, err
	}
	api := k8s.NewK8SApi(c)
	deployments, err := api.ListDeployments(envConfig.K8sNamespace, env.String())
	if err != nil {
		return nil, err
	}
	zones := set.NewStringSet()
	for _, d := range deployments {
		if *d.Spec.Replicas > 0 {
			zone := d.Labels["zone"]
			if zone != "" {
				zones.Add(zone)
			}
		}
	}
	return zones.ToSlice(), nil
}
