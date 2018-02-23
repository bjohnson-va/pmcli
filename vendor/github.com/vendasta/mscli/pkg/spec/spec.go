package spec

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/spf13/viper"
	"github.com/vendasta/mscli/pkg/utils"
)

const (
	// IdentityServiceAccountJwt identity service account
	IdentityServiceAccountJwt = "service-account-jwt"
	// IdentityStaticToken identity static token
	IdentityStaticToken = "static-token"
	// DefaultGCPZone is the default zone we work with if no specific zones are specified in the microservice configuration
	DefaultGCPZone = "us-central1-c"
)

// ProtoPath indicates which protos should be compiled or used for SDK generation.
type ProtoPath struct {
	Path           string `json:"path" yaml:"path"`
	ExcludeFromSDK bool   `json:"excludeFromSdk" yaml:"excludeFromSdk"`
}

// Network allows for custom network settings for the microservice.
type Network struct {
	GRPCHost            string `json:"grpcHost" yaml:"grpcHost"`
	GRPCLoadBalancerIP  string `json:"grpcLoadBalancerIp" yaml:"grpcLoadBalancerIp"`
	HTTPSHost           string `json:"httpsHost" yaml:"httpsHost"`
	HTTPSLoadBalancerIP string `json:"httpsLoadBalancerIp" yaml:"httpsLoadBalancerIp"`

	DeprecatedGRPCLoadBalancerIP string `json:"loadBalancerIp" yaml:"loadBalancerIp"`
}

func (n Network) GetGRPCLoadBalancerIP() string {
	if n.GRPCLoadBalancerIP != "" {
		return n.GRPCLoadBalancerIP
	}
	return n.DeprecatedGRPCLoadBalancerIP
}

// Scaling allows for custom scaling settings for the microservice.
type Scaling struct {
	MaxReplicas int32 `json:"maxReplicas" yaml:"maxReplicas"`
	MinReplicas int32 `json:"minReplicas" yaml:"minReplicas"`
	TargetCPU   int32 `json:"targetCPU" yaml:"targetCPU"`
}

// Resources allows for setting the memory/cpu requirements/limits for the microservice.
type Resources struct {
	MemoryRequest string `json:"memoryRequest" yaml:"memoryRequest"`
	MemoryLimit   string `json:"memoryLimit" yaml:"memoryLimit"`
	CPURequest    string `json:"cpuRequest" yaml:"cpuRequest"`
	CPULimit      string `json:"cpuLimit" yaml:"cpuLimit"`
}

// AppConfig configures sidecar applications.
type AppConfig struct {
	EndpointsVersion string `json:"endpointsVersion" yaml:"endpointsVersion"`
}

type Port struct {
	Name          string `json:"name" yaml:"name"`
	ContainerPort int32  `json:"containerPort" yaml:"containerPort"`
}

// PodConfig allows for secrets and environment variables to be injected into the microservice.
type PodConfig struct {
	Secrets []Secret `json:"secrets" yaml:"secrets"`
	PodEnv  []Env    `json:"podEnv" yaml:"podEnv"`
	Ports   []Port   `json:"ports" yaml:"ports"`
}

// JwtConfig allows for jwt generation
type JwtConfig struct {
	Type                    string `json:"type" yaml:"type"`
	ProjectID               string `json:"project_id" yaml:"project_id"`
	PrivateKeyID            string `json:"private_key_id" yaml:"private_key_id"`
	PrivateKey              string `json:"private_key" yaml:"private_key"`
	ClientEmail             string `json:"client_email" yaml:"client_email"`
	ClientID                string `json:"client_id" yaml:"client_id"`
	AuthURI                 string `json:"auth_uri" yaml:"auth_uri"`
	TokenURI                string `json:"token_uri" yaml:"token_uri"`
	AuthProviderX509CertURL string `json:"auth_provider_x509_cert_url" yaml:"auth_provider_x509_cert_url"`
	ClientX509CertURL       string `json:"client_x509_cert_url" yaml:"client_x509_cert_url"`
}

// SecondarySSLConfig allows for configuration of an additional SSL certificate
type SecondarySSLConfig struct {
	// Host in the form sub.domain.com (no scheme, optional subdomain)
	Host string `json:"host" yaml:"host"`
	// HTTPSHost in the form sub.domain.com (no scheme, optional subdomain)
	HTTPSHost string `json:"httpsHost" yaml:"httpsHost"`
	// Name is the name of the secret that contains the certificate/key (provisioned by SRE)
	Name string `json:"name" yaml:"name"`
}

// EnvironmentConfig controls the settings for each environment of the microservice.
type EnvironmentConfig struct {
	Name               string              `json:"name" yaml:"name"`
	K8sContext         string              `json:"k8sContext" yaml:"k8sContext"`
	K8sNamespace       string              `json:"k8sNamespace" yaml:"k8sNamespace"`
	JwtConfig          *JwtConfig          `json:"jwtConfig" yaml:"jwtConfig"`
	SecondarySSLConfig *SecondarySSLConfig `json:"secondarySSLConfig" yaml:"secondarySSLConfig"`

	Apps      `json:"apps" yaml:"apps"`
	Network   `json:"network" yaml:"network"`
	Scaling   `json:"scaling" yaml:"scaling"`
	Resources `json:"resources" yaml:"resources"`
	AppConfig `json:"appConfig" yaml:"appConfig"`
	PodConfig `json:"podConfig" yaml:"podConfig"`
	Zones     []string `json:"zones" yaml:"zones"`
}

// GetZones returns the list of zones this microservice is in.
func (c EnvironmentConfig) GetZones() []string {
	if len(c.Zones) <= 0 {
		return []string{DefaultGCPZone}
	}
	return c.Zones
}

// Env is a key-value string pair for a pods environment.
type Env struct {
	Key   string `json:"key" yaml:"key"`
	Value string `json:"value" yaml:"value"`
}

// Secret is a custom secret created outside of mscli
type Secret struct {
	Name      string `json:"name" yaml:"name"`
	MountPath string `json:"mountPath" yaml:"mountPath"`
}

// Redis is a config to be supplied if the microservice required it.
type Redis struct {
	Password  string `json:"password" yaml:"password"`
	MaxMemory string `json:"maxMemory" yaml:"maxMemory"`
}

// Apps are 3rd party apps that are commonly used by microservices.
type Apps struct {
	Redis *Redis `json:"redis" yaml:"redis"`
}

// MicroserviceConfig controls the microservice and all of its environments.
type MicroserviceConfig struct {
	Name               string              `json:"name" yaml:"name"`
	GoPackageName      string              `json:"goPackageName" yaml:"goPackageName"`
	ProtoPath          string              `json:"protoPath" yaml:"protoPath"` // Deprecated: Use ProtoPaths instead
	ProtoPaths         []*ProtoPath        `json:"protoPaths" yaml:"protoPaths"`
	Environments       []EnvironmentConfig `json:"environments" yaml:"environments"`
	RepoURL            string              `json:"repoUrl" yaml:"repoUrl"`
	IdentityType       string              `json:"identityType" yaml:"identityType"`
	Dockerfile         string              `json:"dockerfile" yaml:"dockerfile"`
	UseInternalPackage bool                `json:"useInternalPackage" yaml:"useInternalPackage"`
	DatadogDashboardID int                 `json:"datadogDashboardId" yaml:"datadogDashboardId"`
	PublicRoutes       []string            `json:"publicRoutes" yaml:"publicRoutes"`

	Debug bool
}

func (s MicroserviceConfig) ReplaceEnv(cfg EnvironmentConfig) error {
	for i, e := range s.Environments {
		if e.Name == cfg.Name {
			s.Environments[i] = e
			return nil
		}
	}
	return fmt.Errorf("couldn't find env to replace for %s", cfg.Name)
}

func (s MicroserviceConfig) GetEnv(env utils.Environment) (EnvironmentConfig, error) {
	st := env.String()
	for _, e := range s.Environments {
		if e.Name == st {
			return e, nil
		}
	}
	return EnvironmentConfig{}, fmt.Errorf("microservice.yaml did not an EnvironmentConfig for the environment: %s", st)
}

// MicroserviceFile The top level config that controls the schema for `microservice.yaml`
type MicroserviceFile struct {
	Syntax       string             `json:"syntax" yaml:"syntax"`
	Microservice MicroserviceConfig `json:"microservice" yaml:"microservice"`
}

// GenConfig creates the microservice.yaml
func GenConfig(name string) (*MicroserviceFile, error) {
	if _, err := os.Stat("./microservice.yaml"); !os.IsNotExist(err) {
		return nil, fmt.Errorf("error, ./microservice.yaml exists. Please remove and try again")
	}

	if name == "" {
		name = utils.GetUserInput("Service name?\n", "my-service")
	}

	ms, err := newConfig(name)
	if err != nil {
		return nil, err
	}
	config := MicroserviceFile{
		Syntax:       "v1",
		Microservice: *ms,
	}

	viper.SetConfigFile("./microservice.yaml")
	viper.Set("microservice", config.Microservice)
	viper.Set("syntax", config.Syntax)
	viper.WriteConfig()

	return &config, nil
}

func newConfig(name string) (*MicroserviceConfig, error) {
	match, _ := regexp.MatchString("^[a-z\\-]+$", name)
	formattedName := strings.Replace(name, "-", "_", -1)
	if !match {
		return nil, fmt.Errorf("the service name can only contain lowercase letters and dashes")
	}

	protoPaths := []*ProtoPath{
		{
			Path: fmt.Sprintf("%s/v1/%s.proto", formattedName, formattedName),
		},
	}

	config := &MicroserviceConfig{
		Name:               name,
		GoPackageName:      inferPackageNameFromDirectoryStructure(name),
		ProtoPaths:         protoPaths,
		RepoURL:            fmt.Sprintf("https://github.com/vendasta/%s", name),
		IdentityType:       IdentityServiceAccountJwt,
		UseInternalPackage: true,
		Environments: []EnvironmentConfig{
			NewEnvironment(name, "prod"),
			NewEnvironment(name, "demo"),
			NewEnvironment(name, "test"),
			NewEnvironment(name, "local"),
		},
	}
	return config, nil
}

func inferPackageNameFromDirectoryStructure(name string) string {
	defaultPackageName := fmt.Sprintf("github.com/vendasta/%s", name)
	r := regexp.MustCompile("github.com/vendasta.+")
	if dir, err := filepath.Abs(filepath.Dir(os.Args[0])); err == nil {
		if match := r.FindString(dir); match != "" {
			defaultPackageName = match
		}

	}
	return defaultPackageName
}

// NewEnvironment Sets up a new environment (ex: test, demo, prod, local)
func NewEnvironment(service string, env string) EnvironmentConfig {
	var k8sContext string
	var k8sNamespace string
	if env == "local" {
		k8sContext = "minikube"
		k8sNamespace = "default"
	} else {
		k8sContext = "gke_repcore-prod_us-central1-c_vendasta-central"
		k8sNamespace = fmt.Sprintf("%s-%s", service, env)
	}

	var podEnvs []Env
	var secrets []Secret
	var sslConfig SecondarySSLConfig
	if env != "local" {
		podEnvs = append(podEnvs, Env{Key: "GOOGLE_APPLICATION_CREDENTIALS", Value: fmt.Sprintf("/etc/%s/keys/key.json", service)})
		podEnvs = append(podEnvs, Env{Key: "SERVICE_ACCOUNT", Value: fmt.Sprintf("%s-%s@repcore-prod.iam.gserviceaccount.com", service, env)})

		secrets = append(secrets, Secret{Name: fmt.Sprintf("%s-key", service), MountPath: fmt.Sprintf("/etc/%s/keys", service)})

		sslConfig = SecondarySSLConfig{
			Host:      fmt.Sprintf("%s-api-%s.apigateway.co", service, env),
			HTTPSHost: fmt.Sprintf("%s-%s.apigateway.co", service, env),
			Name:      "wildcard-apigateway-co",
		}
	}

	cfg := EnvironmentConfig{
		Name:         env,
		K8sNamespace: k8sNamespace,
		K8sContext:   k8sContext,

		Resources: Resources{
			CPULimit:      "50m",
			CPURequest:    "25m",
			MemoryLimit:   "32Mi",
			MemoryRequest: "16Mi",
		},
		Network: Network{
			GRPCHost:  fmt.Sprintf("%s-api-%s.vendasta-internal.com", service, env),
			HTTPSHost: fmt.Sprintf("%s-%s.vendasta-internal.com", service, env),
		},
		PodConfig: PodConfig{
			PodEnv:  podEnvs,
			Secrets: secrets,
		},
		Scaling: Scaling{
			MaxReplicas: 3,
			MinReplicas: 1,
			TargetCPU:   50,
		},
	}
	if env == "local" {
		cfg.Network.GRPCHost = fmt.Sprintf("%s-api.vendasta-local.com", service)
		cfg.Network.HTTPSHost = fmt.Sprintf("%s.vendasta-local.com", service)
	} else {
		cfg.SecondarySSLConfig = &sslConfig
	}
	if env == "prod" {
		cfg.Scaling.MinReplicas = 2
	}

	return cfg
}
