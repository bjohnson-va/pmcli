package spec

import (
	"encoding/json"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/vendasta/godatatypes/set"
	"github.com/vendasta/mscli/pkg/k8s"
	"github.com/vendasta/mscli/pkg/utils"
	autoscaling "k8s.io/api/autoscaling/v1"
	"k8s.io/api/core/v1"
	"k8s.io/api/extensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

// K8S kubernetes objects configured for the microservice
func (s MicroserviceConfig) K8S(env utils.Environment, tag string, zones []string) ([]k8s.Interface, error) {
	var k []k8s.Interface

	deliveryZones, err := s.deliveryZones(env, zones)
	if err != nil {
		return nil, err
	}

	d, err := s.Deployments(env, tag, deliveryZones)
	if err != nil {
		return nil, err
	}
	k = append(k, d...)

	srv, err := s.Services(env)
	if err != nil {
		return nil, err
	}
	k = append(k, srv...)

	hpa, err := s.HPA(env, deliveryZones)
	if err != nil {
		return nil, err
	}
	k = append(k, hpa...)

	sec, err := s.Secrets(env)
	if err != nil {
		return nil, err
	}
	k = append(k, sec...)

	cm, err := s.ConfigMaps(env)
	if err != nil {
		return nil, err
	}
	k = append(k, cm...)
	if env == utils.Local {
		ips, err := s.ImagePullSecrets(env)
		if err != nil {
			return nil, err
		}
		k = append(k, ips)
	}
	return k, nil
}

func (s MicroserviceConfig) nameFromZone(name, zone string) string {
	if zone != DefaultGCPZone {
		return fmt.Sprintf("%s-%s", name, zone)
	}
	return name
}

func (s MicroserviceConfig) deliveryZones(env utils.Environment, zones []string) ([]string, error) {
	envConfig, err := s.GetEnv(env)
	if err != nil {
		return nil, err
	}

	if len(zones) == 0 {
		zones = envConfig.GetZones()
	} else {
		// Verify that these zone are actually specified in the yaml
		for _, z := range zones {
			found := false
			for _, availableZone := range envConfig.GetZones() {
				if availableZone == z {
					found = true
					break
				}
			}
			if !found {
				return nil, fmt.Errorf("%s is not a valid zone for %s (make sure %s is specified in the %s environment config of the yaml)", z, env.String(), z, env.String())
			}
		}
	}

	// Ensure we have a set of zones
	zoneSet := set.NewStringSet()
	for _, z := range zones {
		zoneSet.Add(z)
	}
	return zoneSet.ToSlice(), nil
}

// Deployments creates the k8s deployment configurations from the microservice spec.
//
// Zones may be specified to limit which deployments are updated to the new image. If no zones are specified a
// deployment will be created for each zone in the spec. If no zones are specified in the spec a deployment will be
// created for the default zone (DefaultGCPZone).
func (s MicroserviceConfig) Deployments(env utils.Environment, tag string, zones []string) ([]k8s.Interface, error) {
	replicas := int32(1)
	revisionHistoryLimit := int32(5)
	var maxUnavailable intstr.IntOrString
	var maxSurge intstr.IntOrString
	if env == utils.Prod {
		maxUnavailable = intstr.FromString("25%")
		maxSurge = intstr.FromString("50%")
	} else {
		maxUnavailable = intstr.FromString("0%")
		maxSurge = intstr.FromString("25%")
	}

	se, err := s.GetEnv(env)
	if err != nil {
		return nil, err
	}

	var podEnv []Env
	podEnv = append(podEnv, se.PodEnv...)
	// Add the public routes as an enviornment variable to send them through to IAM
	podEnv = append(podEnv, Env{Key: "PUBLIC_ROUTES", Value: strings.Join(s.PublicRoutes, ",")})
	if se.Redis != nil {
		podEnv = append(podEnv,
			Env{Key: "REDIS_HOST", Value: fmt.Sprintf("redis-%s.%s.svc.cluster.local:6379", s.Name, se.K8sNamespace)},
			Env{Key: "REDIS_PASSWORD", Value: se.Redis.Password},
		)
	}

	depContainers := []v1.Container{
		App(s.Name, tag, env, podEnv, se.Secrets, se.Resources, se.EndpointsVersion, se.Ports),
		GoogleAuth(env, se.HTTPSHost),
	}
	volumes := []v1.Volume{
		{
			Name: "vendasta-internal",
			VolumeSource: v1.VolumeSource{
				Secret: &v1.SecretVolumeSource{
					SecretName: "vendasta-internal-secret",
				},
			},
		},
	}

	if se.SecondarySSLConfig != nil && se.SecondarySSLConfig.Name != "" {
		volumes = append(volumes, v1.Volume{
			Name: se.SecondarySSLConfig.Name,
			VolumeSource: v1.VolumeSource{
				Secret: &v1.SecretVolumeSource{
					SecretName: se.SecondarySSLConfig.Name,
				},
			},
		})
	}

	if se.EndpointsVersion != "" {
		depContainers = append(depContainers, Endpoints(env, se.GRPCHost, se.EndpointsVersion, se.SecondarySSLConfig))
		volumes = append(volumes, v1.Volume{
			Name: "endpoints-nginx-conf",
			VolumeSource: v1.VolumeSource{
				ConfigMap: &v1.ConfigMapVolumeSource{
					LocalObjectReference: v1.LocalObjectReference{
						Name: "endpoints-nginx-conf",
					},
				},
			},
		})
	}

	if env == utils.Local {
		volumes = append(volumes,
			v1.Volume{
				Name: "local-auth",
				VolumeSource: v1.VolumeSource{
					HostPath: &v1.HostPathVolumeSource{
						Path: WellKnownFile(),
					},
				},
			},
			v1.Volume{
				Name: "local-app-creds",
				VolumeSource: v1.VolumeSource{
					Secret: &v1.SecretVolumeSource{
						SecretName: "vendasta-local-secret",
					},
				},
			},
		)
	}
	for _, secret := range se.PodConfig.Secrets {
		volumes = append(volumes, v1.Volume{
			Name: secret.Name,
			VolumeSource: v1.VolumeSource{
				Secret: &v1.SecretVolumeSource{
					SecretName: secret.Name,
				},
			},
		})
	}

	deployments := make([]k8s.Interface, len(zones))
	for _, zone := range zones {
		deployName := s.nameFromZone(s.Name, zone)
		d := v1beta1.Deployment{
			ObjectMeta: metav1.ObjectMeta{
				Name:      deployName,
				Namespace: se.K8sNamespace,
				Labels: map[string]string{
					"appId":       s.Name,
					"zone":        zone,
					"environment": env.String(),
				},
			},
			Spec: v1beta1.DeploymentSpec{
				Replicas:             &replicas,
				RevisionHistoryLimit: &revisionHistoryLimit,
				Strategy: v1beta1.DeploymentStrategy{
					Type: v1beta1.RollingUpdateDeploymentStrategyType,
					RollingUpdate: &v1beta1.RollingUpdateDeployment{
						MaxUnavailable: &maxUnavailable,
						MaxSurge:       &maxSurge,
					},
				},
				Template: v1.PodTemplateSpec{
					ObjectMeta: metav1.ObjectMeta{
						Labels: map[string]string{
							"app":             s.Name,
							"environment":     env.String(),
							"traffic-enabled": "true",
							zone:              "true",
						},
					},
					Spec: v1.PodSpec{
						Containers: depContainers,
						Volumes:    volumes,
						Affinity: &v1.Affinity{NodeAffinity: &v1.NodeAffinity{
							RequiredDuringSchedulingIgnoredDuringExecution: &v1.NodeSelector{
								NodeSelectorTerms: []v1.NodeSelectorTerm{
									{MatchExpressions: []v1.NodeSelectorRequirement{
										{
											Key:      "failure-domain.beta.kubernetes.io/zone",
											Operator: v1.NodeSelectorOpIn,
											Values:   []string{zone},
										},
									}},
								},
							},
						}},
					},
				},
			},
		}
		if env == utils.Local {
			secret := v1.LocalObjectReference{
				Name: "vendasta-local-gcr",
			}
			d.Spec.Template.Spec.ImagePullSecrets = append(d.Spec.Template.Spec.ImagePullSecrets, secret)
		}
		deployments = append(deployments, deployment{&d})
	}
	if se.Redis != nil {
		d, err := s.redisDeployment(env)
		if err != nil {
			return nil, err
		}
		deployments = append(deployments, d)
	}
	if env == utils.Local {
		deployments = append(deployments, s.localProxyDeployment())
	}
	return deployments, nil
}

func (s MicroserviceConfig) redisDeployment(env utils.Environment) (k8s.Interface, error) {
	se, err := s.GetEnv(env)
	if err != nil {
		return nil, err
	}

	if se.Redis.Password == "" {
		return nil, fmt.Errorf("redis password must be supplied")
	}
	depName := fmt.Sprintf("redis-%s", s.Name)

	replicas := int32(1)
	revisionHistoryLimit := int32(5)
	containers := []v1.Container{
		RedisContainer(env, se.Redis.Password, se.Redis.MaxMemory),
	}
	dep := v1beta1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      depName,
			Namespace: se.K8sNamespace,
		},
		Spec: v1beta1.DeploymentSpec{
			Replicas:             &replicas,
			RevisionHistoryLimit: &revisionHistoryLimit,
			Template: v1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": depName,
					},
				},
				Spec: v1.PodSpec{
					Containers: containers,
				},
			},
		},
	}
	return deployment{&dep}, nil
}

func (s MicroserviceConfig) localProxyDeployment() k8s.Interface {
	replicas := int32(1)
	revisionHistoryLimit := int32(5)
	maxUnavailable := intstr.FromString("25%")
	maxSurge := intstr.FromString("25%")
	containers := []v1.Container{
		LocalProxy(),
	}
	volumes := []v1.Volume{
		{
			Name: "vendasta-local-secret",
			VolumeSource: v1.VolumeSource{
				Secret: &v1.SecretVolumeSource{
					SecretName: "vendasta-local-secret",
				},
			},
		},
	}
	return deployment{
		&v1beta1.Deployment{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "local-proxy",
				Namespace: "default",
			},
			Spec: v1beta1.DeploymentSpec{
				Replicas:             &replicas,
				RevisionHistoryLimit: &revisionHistoryLimit,
				Strategy: v1beta1.DeploymentStrategy{
					Type: v1beta1.RollingUpdateDeploymentStrategyType,
					RollingUpdate: &v1beta1.RollingUpdateDeployment{
						MaxUnavailable: &maxUnavailable,
						MaxSurge:       &maxSurge,
					},
				},
				Template: v1.PodTemplateSpec{
					ObjectMeta: metav1.ObjectMeta{
						Labels: map[string]string{
							"app": "local-proxy",
						},
					},
					Spec: v1.PodSpec{
						Containers: containers,
						Volumes:    volumes,
					},
				},
			},
		},
	}
}

// ConfigMaps kubernetes services for the microservice spec
func (s MicroserviceConfig) ConfigMaps(env utils.Environment) ([]k8s.Interface, error) {
	se, err := s.GetEnv(env)
	if err != nil {
		return nil, err
	}
	if se.EndpointsVersion != "" {
		cm := EndpointsConfigMap(se.K8sNamespace, se.GRPCHost, se.SecondarySSLConfig)
		return []k8s.Interface{
			configMap{&cm},
		}, nil
	}
	return nil, nil
}

// Services kubernetes services for the microservice spec
func (s MicroserviceConfig) Services(env utils.Environment) ([]k8s.Interface, error) {
	grpcServiceName := fmt.Sprintf("%s-grpc-svc", s.Name)
	se, err := s.GetEnv(env)
	if err != nil {
		return nil, err
	}
	grpcService := &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      grpcServiceName,
			Namespace: se.K8sNamespace,
		},
		Spec: v1.ServiceSpec{
			Ports: []v1.ServicePort{
				{
					Name:       "grpc",
					Port:       int32(443),
					TargetPort: intstr.FromInt(11006),
					Protocol:   v1.ProtocolTCP,
				},
			},
			Type:           v1.ServiceTypeLoadBalancer,
			LoadBalancerIP: se.Network.GetGRPCLoadBalancerIP(),
			Selector: map[string]string{
				"app":         s.Name,
				"environment": env.String(),
			},
		},
	}
	if env == utils.Local {
		grpcService.Spec.LoadBalancerIP = ""
		grpcService.Spec.Ports[0].Port = int32(31957)
		grpcService.Spec.Ports[0].TargetPort = intstr.FromInt(11003)
		grpcService.ObjectMeta.Annotations = map[string]string{
			"vendasta-local.com/domain": se.GRPCHost,
			"vendasta-local.com/port":   "31957",
		}
	}

	httpsServiceName := fmt.Sprintf("%s-https-svc", s.Name)
	httpService := &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      httpsServiceName,
			Namespace: se.K8sNamespace,
			Annotations: map[string]string{
				"prometheus.io/scrape": "true",
			},
		},
		Spec: v1.ServiceSpec{
			Ports: []v1.ServicePort{
				{
					Name:       "https",
					Port:       int32(443),
					TargetPort: intstr.FromInt(11002),
					Protocol:   v1.ProtocolTCP,
				},
			},
			Type:           v1.ServiceTypeLoadBalancer,
			LoadBalancerIP: se.Network.HTTPSLoadBalancerIP,
			Selector: map[string]string{
				"app":         s.Name,
				"environment": env.String(),
			},
		},
	}
	if env == utils.Local {
		httpService.Spec.LoadBalancerIP = ""
		httpService.Spec.Ports[0].Port = int32(31958)
		httpService.Spec.Ports[0].TargetPort = intstr.FromInt(11003)
		httpService.ObjectMeta.Annotations = map[string]string{
			"vendasta-local.com/domain": se.HTTPSHost,
			"vendasta-local.com/port":   "31958",
		}
	}

	services := []k8s.Interface{
		service{grpcService},
		service{httpService},
	}

	if se.Redis != nil {
		redisSvc := &v1.Service{
			ObjectMeta: metav1.ObjectMeta{
				Name:      fmt.Sprintf("redis-%s", s.Name),
				Namespace: se.K8sNamespace,
			},
			Spec: v1.ServiceSpec{
				Ports: []v1.ServicePort{
					{
						Name:       "tcp",
						Port:       int32(6379),
						TargetPort: intstr.FromInt(6379),
						Protocol:   v1.ProtocolTCP,
					},
				},
				Type: v1.ServiceTypeLoadBalancer,
				Selector: map[string]string{
					"app": fmt.Sprintf("redis-%s", s.Name),
				},
			},
		}
		services = append(services, service{redisSvc})
	}

	if env == utils.Local {
		localProxySvc := &v1.Service{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "local-proxy",
				Namespace: "default",
			},
			Spec: v1.ServiceSpec{
				Ports: []v1.ServicePort{
					{
						Name:       "https",
						Port:       int32(443),
						TargetPort: intstr.FromInt(443),
						Protocol:   v1.ProtocolTCP,
						NodePort:   32000,
					}, {
						Name:       "http",
						Port:       int32(80),
						TargetPort: intstr.FromInt(80),
						Protocol:   v1.ProtocolTCP,
						NodePort:   32001,
					},
				},
				Type: v1.ServiceTypeLoadBalancer,
				Selector: map[string]string{
					"app": "local-proxy",
				},
			},
		}
		services = append(services, service{localProxySvc})
	}

	return services, nil
}

// HPA kubernetes horizontal pod autoscaler for the microservice spec
func (s MicroserviceConfig) HPA(env utils.Environment, zones []string) ([]k8s.Interface, error) {
	c, err := s.GetEnv(env)
	if err != nil {
		return nil, err
	}
	if c.Scaling.MinReplicas <= 0 {
		c.Scaling.MinReplicas = 1
	}
	if c.Scaling.MaxReplicas <= 0 {
		c.Scaling.MaxReplicas = 3
	}
	if c.Scaling.TargetCPU <= 0 {
		c.Scaling.TargetCPU = 50
	}

	minR := c.Scaling.MinReplicas
	maxR := c.Scaling.MaxReplicas
	cpu := c.Scaling.TargetCPU

	specs := make([]k8s.Interface, len(zones))
	for _, z := range zones {
		hpaName := s.nameFromZone(s.Name, z)
		hpaSpec := &autoscaling.HorizontalPodAutoscaler{
			ObjectMeta: metav1.ObjectMeta{
				Name:      hpaName,
				Namespace: c.K8sNamespace,
			},
			Spec: autoscaling.HorizontalPodAutoscalerSpec{
				ScaleTargetRef: autoscaling.CrossVersionObjectReference{
					Kind: "Deployment",
					Name: hpaName,
				},
				MinReplicas:                    &minR,
				MaxReplicas:                    maxR,
				TargetCPUUtilizationPercentage: &cpu,
			},
		}
		specs = append(specs, hpa{hpaSpec})
	}
	return specs, nil
}

// ImagePullSecrets pull secrets for the microservice spec
func (s MicroserviceConfig) ImagePullSecrets(env utils.Environment) (k8s.Interface, error) {
	se, err := s.GetEnv(env)
	if err != nil {
		return nil, err
	}
	return imagePullSecrets{
		imagePullSecrets: v1.LocalObjectReference{Name: "vendasta-local-gcr"},
		namespace:        se.K8sNamespace,
	}, nil
}

type dockerConfig struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// CreateNamespace will generate a namespace
func (s MicroserviceConfig) CreateNamespace(env utils.Environment) ([]k8s.Interface, error) {
	se, err := s.GetEnv(env)
	if err != nil {
		return nil, err
	}
	n := []k8s.Interface{
		namespace{
			&v1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Name: se.K8sNamespace,
				},
			},
		},
	}
	return n, nil
}

// CreateServiceAccountSecret will generate a secret for the provisioned service account
func (s MicroserviceConfig) CreateServiceAccountSecret(env utils.Environment, value []byte) ([]k8s.Interface, error) {
	se, err := s.GetEnv(env)
	if err != nil {
		return nil, err
	}
	secrets := []k8s.Interface{
		secret{
			&v1.Secret{
				ObjectMeta: metav1.ObjectMeta{
					Name:      fmt.Sprintf("%s-key", s.Name),
					Namespace: se.K8sNamespace,
				},
				Data: map[string][]byte{
					"key.json": value,
				},
			},
		},
	}
	return secrets, nil
}

// CreateApiGatewaySecret will generate a secret for the apigateway.co whitelabelled domain
func (s MicroserviceConfig) CreateApiGatewaySecret(env utils.Environment) ([]k8s.Interface, error) {
	se, err := s.GetEnv(env)
	if err != nil {
		return nil, err
	}
	secrets := []k8s.Interface{
		secret{
			&v1.Secret{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "wildcard-apigateway-co",
					Namespace: se.K8sNamespace,
				},
				Data: map[string][]byte{
					"nginx.crt": []byte(ApiGatewayCert),
					"nginx.key": []byte(ApiGatewayKey),
					"tls.crt":   []byte(ApiGatewayCert),
					"tls.key":   []byte(ApiGatewayKey),
				},
			},
		},
	}
	return secrets, nil
}

// Secrets secrets that the kubernetes
func (s MicroserviceConfig) Secrets(env utils.Environment) ([]k8s.Interface, error) {
	se, err := s.GetEnv(env)
	if err != nil {
		return nil, err
	}
	secrets := []k8s.Interface{
		secret{
			&v1.Secret{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "vendasta-internal-secret",
					Namespace: se.K8sNamespace,
				},
				Data: map[string][]byte{
					// Kubernetes Ingress format
					"tls.crt": []byte(VendastaInternalCert),
					"tls.key": []byte(VendastaInternalKey),
					// NGINX format
					"nginx.crt": []byte(VendastaInternalCert),
					"nginx.key": []byte(VendastaInternalKey),
				},
			},
		},
	}
	if env == utils.Local {
		dc := dockerConfig{
			Username: "_json_key",
			Password: VendastaLocalJSONKey,
			Email:    "123@3456.com",
			//Auth: auth,
		}
		dockerConfigBytes, err := json.Marshal(map[string]interface{}{"https://gcr.io": dc})
		if err != nil {
			return nil, fmt.Errorf("error creating json key for docker gcr.io auth %s", err.Error())
		}

		secrets = append(secrets,
			secret{
				&v1.Secret{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "vendasta-local-secret",
						Namespace: se.K8sNamespace,
					},
					Data: map[string][]byte{
						// vendasta local service account key
						"key.json": []byte(VendastaLocalJSONKey),

						// Kubernetes Ingress format
						"tls.crt": []byte(VendastaLocalCert),
						"tls.key": []byte(VendastaLocalKey),

						// NGINX format
						"nginx.crt": []byte(VendastaLocalCert),
						"nginx.key": []byte(VendastaLocalKey),
					},
				},
			},
			secret{
				&v1.Secret{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "vendasta-local-gcr",
						Namespace: se.K8sNamespace,
					},
					Data: map[string][]byte{
						// GCR.IO Pull Secrets
						".dockercfg": dockerConfigBytes,
					},
					Type: v1.SecretTypeDockercfg,
				},
			},
		)
	}
	return secrets, nil
}

//WellKnownFile returns the well-known gcloud cli config file
func WellKnownFile() string {
	return filepath.Join(guessUnixHomeDir(), ".config", "gcloud")
}

func guessUnixHomeDir() string {
	// Prefer $HOME over user.Current due to glibc bug: golang.org/issue/13470
	if v := os.Getenv("HOME"); v != "" {
		return v
	}
	// Else, fall back to user.Current:
	if u, err := user.Current(); err == nil {
		return u.HomeDir
	}
	return ""
}
