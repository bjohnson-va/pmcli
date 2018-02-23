package spec

import (
	"github.com/vendasta/mscli/pkg/utils"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/apimachinery/pkg/util/intstr"
)

// App sets up environment variables, volumes, and some container stuff
func App(name string, tag string, environment utils.Environment, env []Env, secrets []Secret, resources Resources, endpointsVersion string, ports []Port) v1.Container {
	envvars := []v1.EnvVar{
		{
			Name: "GKE_PODNAME",
			ValueFrom: &v1.EnvVarSource{
				FieldRef: &v1.ObjectFieldSelector{
					FieldPath: "metadata.name",
				},
			},
		},
		{
			Name: "GKE_NAMESPACE",
			ValueFrom: &v1.EnvVarSource{
				FieldRef: &v1.ObjectFieldSelector{
					FieldPath: "metadata.namespace",
				},
			},
		},
		{
			Name: "GKE_PODIP",
			ValueFrom: &v1.EnvVarSource{
				FieldRef: &v1.ObjectFieldSelector{
					FieldPath: "status.podIP",
				},
			},
		},
		{
			Name:  "ENVIRONMENT",
			Value: environment.String(),
		},
	}

	for _, customEnv := range env {
		envvars = append(envvars, v1.EnvVar{Name: customEnv.Key, Value: customEnv.Value})
	}

	// Volume Mounts
	vm := []v1.VolumeMount{}
	for _, secret := range secrets {
		vm = append(vm, v1.VolumeMount{Name: secret.Name, MountPath: secret.MountPath})
	}

	if environment == utils.Local {
		vm = append(vm,
			v1.VolumeMount{Name: "local-auth", MountPath: "/etc/local-auth/"},
			v1.VolumeMount{Name: "local-app-creds", MountPath: "/etc/local-app-creds/"},
		)
		envvars = append(envvars, v1.EnvVar{
			Name:  "GOOGLE_APPLICATION_CREDENTIALS",
			Value: "/etc/local-auth/application_default_credentials.json",
		})
	}

	var probe *v1.Probe
	if endpointsVersion == "" {
		probe = &v1.Probe{
			Handler: v1.Handler{
				HTTPGet: &v1.HTTPGetAction{
					Path: "/healthz",
					Port: intstr.FromInt(11001),
				},
			},
		}
	}

	appPorts := []v1.ContainerPort{
		{
			Name:          "http",
			ContainerPort: 11001,
		},
		{
			Name:          "grpc",
			ContainerPort: 11000,
		},
	}

	for _, port := range ports {
		appPorts = append(appPorts, v1.ContainerPort{
			Name:          port.Name,
			ContainerPort: port.ContainerPort,
		})
	}

	rvalue := v1.Container{
		Name:            name,
		Image:           tag,
		ImagePullPolicy: v1.PullIfNotPresent,
		Ports:           appPorts,
		Env:             envvars,
		Resources: v1.ResourceRequirements{
			Limits: v1.ResourceList{
				v1.ResourceCPU:    resource.MustParse(resources.CPULimit),
				v1.ResourceMemory: resource.MustParse(resources.MemoryLimit),
			},
			Requests: v1.ResourceList{
				v1.ResourceCPU:    resource.MustParse(resources.CPURequest),
				v1.ResourceMemory: resource.MustParse(resources.MemoryRequest),
			},
		},
		VolumeMounts:   vm,
		LivenessProbe:  probe,
		ReadinessProbe: probe,
	}

	if environment == utils.Local {
		rvalue.ImagePullPolicy = v1.PullNever
	}
	return rvalue
}
