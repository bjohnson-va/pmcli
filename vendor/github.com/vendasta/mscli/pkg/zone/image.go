package zone

import (
	"fmt"
	"strings"

	"github.com/vendasta/mscli/pkg/k8s"
	"github.com/vendasta/mscli/pkg/kubernetes"
	"github.com/vendasta/mscli/pkg/spec"
	"github.com/vendasta/mscli/pkg/utils"
	"k8s.io/api/extensions/v1beta1"
)

type getImageConfig struct {
	tagOnly bool
}

// GetImageOption options for GetImage
type GetImageOption func(*getImageConfig)

// TagsOnly will return the image tags, and not the full path
func TagsOnly() GetImageOption {
	return func(config *getImageConfig) {
		config.tagOnly = true
	}
}

// GetImage returns the current image for the given zone
func GetImage(spec *spec.MicroserviceFile, env utils.Environment, zone string, opts ...GetImageOption) (string, error) {
	if env == utils.Local {
		return "", fmt.Errorf("local does not support zones")
	}
	if spec == nil {
		return "", fmt.Errorf("spec file is required")
	}
	config := &getImageConfig{
		tagOnly: false,
	}
	for _, o := range opts {
		o(config)
	}
	c, err := kubernetes.GetK8sClientSet(*spec, env)
	if err != nil {
		return "", err
	}
	envConfig, err := spec.Microservice.GetEnv(env)
	if err != nil {
		return "", err
	}
	api := k8s.NewK8SApi(c)
	deployments, err := api.ListDeployments(envConfig.K8sNamespace, env.String())
	if err != nil {
		return "", err
	}
	var deployment *v1beta1.Deployment
	for _, d := range deployments {
		if d.Labels["zone"] == zone {
			deployment = &d
			break
		}
	}
	if deployment == nil {
		return "", fmt.Errorf("%s has no deployments", zone)
	}
	for _, c := range deployment.Spec.Template.Spec.Containers {
		if c.Name == spec.Microservice.Name {
			if config.tagOnly {
				return strings.Split(c.Image, ":")[1], nil
			}
			return c.Image, nil
		}
	}
	return "", fmt.Errorf("%s is not a valid zone", zone)
}
