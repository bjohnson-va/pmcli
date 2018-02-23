package spec

import (
	"fmt"
	"strconv"

	"github.com/vendasta/mscli/pkg/utils"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
)

//QuantityToMB converts a resource.Quantity to a string expressing the quantity in MB
//EG 16Mi -> "17mb"
//This is the string format that redis expects
func quantityToMB(q resource.Quantity) string {
	n := q.ScaledValue(resource.Mega)
	s := strconv.FormatInt(n, 10)
	s += "mb"
	return s
}

//safeRedisLimit returns a safe memory limit based on a specified resource.Quantity request
//Right now this limit is basically q * 1.5
func safeRedisLimit(q resource.Quantity) resource.Quantity {
	n := q.Value() / 2
	nq := resource.NewQuantity(n, q.Format)
	q.Add(*nq)
	return resource.MustParse(q.String())
}

// Redis setup container for redis
func RedisContainer(environment utils.Environment, password string, maxMemory string) v1.Container {
	if maxMemory == "" {
		if environment == utils.Prod {
			maxMemory = "200Mi"
		} else {
			maxMemory = "16Mi"
		}
	}

	cpuLimits := resource.MustParse("50m")
	cpuRequests := resource.MustParse("25m")

	if environment == utils.Prod {
		cpuLimits = resource.MustParse("2000m")
		cpuRequests = resource.MustParse("1000m")
	}

	memoryLimits := safeRedisLimit(resource.MustParse(maxMemory))
	memoryRequests := resource.MustParse(maxMemory)

	env := []v1.EnvVar{}
	vm := []v1.VolumeMount{}
	return v1.Container{
		Name:  "redis",
		Image: "redis",
		Args: []string{
			"redis-server", fmt.Sprintf("--requirepass %s", password), fmt.Sprintf("--maxmemory %s", quantityToMB(memoryRequests)),
		},
		ImagePullPolicy: v1.PullAlways,
		Ports: []v1.ContainerPort{
			{
				ContainerPort: 6379,
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
