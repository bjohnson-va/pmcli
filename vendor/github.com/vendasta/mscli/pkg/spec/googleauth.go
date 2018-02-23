package spec

import (
	"fmt"

	"github.com/vendasta/mscli/pkg/utils"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
)

// GoogleAuth sets up google auth secrets and env vars
func GoogleAuth(environment utils.Environment, httpsHost string) v1.Container {
	cpuLimits := resource.MustParse("50m")
	cpuRequests := resource.MustParse("25m")
	memoryLimits := resource.MustParse("32Mi")
	memoryRequests := resource.MustParse("16Mi")

	if environment == utils.Prod {
		cpuLimits = resource.MustParse("100m")
		cpuRequests = resource.MustParse("50m")
		memoryLimits = resource.MustParse("128Mi")
		memoryRequests = resource.MustParse("64Mi")
	}

	container := v1.Container{
		Name:            "auth-proxy",
		Image:           "gcr.io/repcore-prod/google_auth_proxy:v11",
		ImagePullPolicy: v1.PullIfNotPresent,
		Ports: []v1.ContainerPort{
			{
				Name:          "https",
				ContainerPort: int32(11002),
			},
		},
		Env: []v1.EnvVar{
			{
				Name:  "REDIRECT_URL",
				Value: fmt.Sprintf("https://%s/oauth2/callback", httpsHost),
			},
			{
				Name:  "EMAIL_DOMAIN",
				Value: "vendasta.com",
			},
			{
				Name:  "UPSTREAM_URL",
				Value: "http://127.0.0.1:11001",
			},
			{
				Name:  "HTTPS_ADDRESS",
				Value: "0.0.0.0:11002",
			},
			{
				Name:  "SECURE_COOKIE",
				Value: "true",
			},
			{
				Name:  "CLIENT_ID",
				Value: "999898651218.apps.googleusercontent.com",
			},
			{
				Name:  "CLIENT_SECRET",
				Value: "HxfjhCvesf9cmymiUIptoT88",
			},
			{
				Name:  "COOKIE_SECRET",
				Value: "ZDRhMTlkNTczNDYyNDY2ZGJhMDFiYWQ2M2YyM2IxMTYK",
			},
			{
				Name:  "SERVER_CERT",
				Value: "/etc/vendasta-internal/tls.crt",
			},
			{
				Name:  "SERVER_KEY",
				Value: "/etc/vendasta-internal/tls.key",
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
		VolumeMounts: []v1.VolumeMount{
			{
				Name:      "vendasta-internal",
				MountPath: "/etc/vendasta-internal",
			},
		},
	}

	if environment == utils.Local {
		container.VolumeMounts = append(container.VolumeMounts,
			v1.VolumeMount{
				Name:      "local-app-creds",
				MountPath: "/etc/local-app-creds/",
			})
	}
	return container
}
