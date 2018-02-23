package kubernetes

import (
	"encoding/base64"
	"fmt"
	"github.com/vendasta/mscli/pkg/k8s"
	"github.com/vendasta/mscli/pkg/spec"
	"github.com/vendasta/mscli/pkg/utils"
	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/iam/v1"
)

// ProvisionService will create the k8s namespace, repcore-prod service account and secret
func ProvisionService(spec spec.MicroserviceFile, env utils.Environment) (string, error) {
	client, err := google.DefaultClient(context.Background(), "https://www.googleapis.com/auth/iam")
	if err != nil {
		return "", fmt.Errorf("could not create default client: %s", err.Error())
	}

	iamService, err := iam.New(client)
	if err != nil {
		return "", fmt.Errorf("could not create iam service: %s", err.Error())
	}

	serviceAccountName := fmt.Sprintf("%s-%s", spec.Microservice.Name, env.String())
	account, err := CreateServiceAccount(spec, iamService, serviceAccountName)
	if err != nil {
		return "", err
	}

	key, err := CreateServiceAccountKey(spec, iamService, account.UniqueId)
	if err != nil {
		return "", err
	}

	private, err := base64.StdEncoding.DecodeString(key.PrivateKeyData)
	if err != nil {
		return "", fmt.Errorf("could not decode private key for service account")
	}

	if spec.Microservice.Debug {
		fmt.Printf("Creating K8S Namespace & Secrets...\n")
	}

	k8sClient, err := GetK8sClientSet(spec, env)
	if err != nil {
		return "", err
	}

	k8sAPI := k8s.NewK8SApi(k8sClient)
	var k8s []k8s.Interface
	ns, err := spec.Microservice.CreateNamespace(env)
	if err != nil {
		return "", err
	}
	k8s = append(k8s, ns...)
	as, err := spec.Microservice.CreateServiceAccountSecret(env, private)
	if err != nil {
		return "", err
	}
	k8s = append(k8s, as...)
	gws, err := spec.Microservice.CreateApiGatewaySecret(env)
	if err != nil {
		return "", err
	}
	k8s = append(k8s, gws...)
	err = k8sAPI.Apply(k8s...)
	if err != nil {
		return "", err
	}

	return serviceAccountName, nil
}

// CreateServiceAccount will create a service account for the app and environment: https://cloud.google.com/iam/reference/rest/v1/projects.serviceAccounts/create
func CreateServiceAccount(spec spec.MicroserviceFile, service *iam.Service, serviceAccountName string) (*iam.ServiceAccount, error) {
	if spec.Microservice.Debug {
		fmt.Printf("Creating Service Account...\n")
	}

	createServiceAccountRequest := &iam.CreateServiceAccountRequest{
		AccountId: serviceAccountName,
		ServiceAccount: &iam.ServiceAccount{
			DisplayName: serviceAccountName,
		},
	}

	createCall := service.Projects.ServiceAccounts.Create("projects/repcore-prod", createServiceAccountRequest)
	s, err := createCall.Do()
	if err != nil {
		if spec.Microservice.Debug {
			fmt.Printf("Error creating Service: %s", err.Error())
			fmt.Printf("Trying to get existing Service... \n")
		}
		serviceAccountEmail := fmt.Sprintf("%s@repcore-prod.iam.gserviceaccount.com", serviceAccountName)
		getCall := service.Projects.ServiceAccounts.Get(fmt.Sprintf("projects/repcore-prod/serviceAccounts/%s", serviceAccountEmail))
		s, err = getCall.Do()
		if err != nil {
			return nil, fmt.Errorf("unable to get service account: %s", err.Error())
		}
	}
	return s, nil
}

// CreateServiceAccountKey creates a new service account key for the given service account: https://cloud.google.com/iam/reference/rest/v1/projects.serviceAccounts.keys/create
func CreateServiceAccountKey(spec spec.MicroserviceFile, service *iam.Service, uniqueID string) (*iam.ServiceAccountKey, error) {
	if spec.Microservice.Debug {
		fmt.Printf("Creating Service Account Key...\n")
	}
	createServiceAccountKeyRequest := &iam.CreateServiceAccountKeyRequest{}
	iamCall := service.Projects.ServiceAccounts.Keys.Create(fmt.Sprintf("projects/repcore-prod/serviceAccounts/%s", uniqueID), createServiceAccountKeyRequest)
	s, err := iamCall.Do()
	if err != nil {
		return nil, fmt.Errorf("unable to create service account: %s", err.Error())
	}
	return s, nil
}
