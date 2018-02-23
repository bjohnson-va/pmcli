package spec

import (
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
)

// LocalProxy sets up local proxy
func LocalProxy() v1.Container {
	cpuLimits := resource.MustParse("50m")
	cpuRequests := resource.MustParse("25m")
	memoryLimits := resource.MustParse("32Mi")
	memoryRequests := resource.MustParse("16Mi")

	env := []v1.EnvVar{}
	vm := []v1.VolumeMount{
		{Name: "vendasta-local-secret", MountPath: "/etc/local-proxy", ReadOnly: true},
	}
	return v1.Container{
		Name:            "local-proxy",
		Image:           "vendasta/local-proxy",
		ImagePullPolicy: v1.PullAlways,
		Ports: []v1.ContainerPort{
			{
				ContainerPort: 443,
			},
		},
		Resources: v1.ResourceRequirements{
			Limits: v1.ResourceList{
				v1.ResourceCPU:    cpuLimits,
				v1.ResourceMemory: memoryLimits,
			},
			Requests: v1.ResourceList{
				v1.ResourceCPU:    cpuRequests,
				v1.ResourceMemory: memoryRequests,
			},
		},
		VolumeMounts: vm,
		Env:          env,
	}
}
