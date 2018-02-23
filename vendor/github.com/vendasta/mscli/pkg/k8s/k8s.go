package k8s

import (
	"fmt"

	autoscaling "k8s.io/api/autoscaling/v1"
	"k8s.io/api/core/v1"
	"k8s.io/api/extensions/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// Interface interface for k8s api to apply
type Interface interface {
	K8SType() string
}

// Namespace k8s namespace
type Namespace interface {
	Interface
	Namespace() *v1.Namespace
}

// Deployment k8s deployment
type Deployment interface {
	Interface
	Deployment() *v1beta1.Deployment
}

// Service k8s service
type Service interface {
	Interface
	Service() *v1.Service
}

// ConfigMap k8s configmap
type ConfigMap interface {
	Interface
	ConfigMap() *v1.ConfigMap
}

// Secret k8s secret
type Secret interface {
	Interface
	Secret() *v1.Secret
}

// HPA k8s horizontal pod autoscaler
type HPA interface {
	Interface
	HPA() *autoscaling.HorizontalPodAutoscaler
}

// ImagePullSecrets k8s image pull secrets
type ImagePullSecrets interface {
	Interface
	ImagePullSecrets() v1.LocalObjectReference
	Namespace() string
}

// API kubernetes clientset
type API struct {
	clientset *kubernetes.Clientset
}

// Apply calls k8s api apply
func (k *API) Apply(k8sObjects ...Interface) error {
	for _, k8sObject := range k8sObjects {
		err := k.ApplyK8SObject(k8sObject)
		if err != nil {
			return fmt.Errorf("error applying k8s object. %s", err.Error())
		}
	}
	return nil
}

// ApplyK8SObject applies a k8s object to a cluster
func (k *API) ApplyK8SObject(k8sObject Interface) error {
	switch k8sObject.(type) {
	case Namespace:
		k.ApplyNamespace(k8sObject.(Namespace))
	case Deployment:
		k.ApplyDeployment(k8sObject.(Deployment))
	case Service:
		k.ApplyService(k8sObject.(Service))
	case ConfigMap:
		k.ApplyConfigMap(k8sObject.(ConfigMap))
	case Secret:
		k.ApplySecret(k8sObject.(Secret))
	case HPA:
		k.ApplyHPA(k8sObject.(HPA))
	case ImagePullSecrets:
		k.ApplyImagePullSecrets(k8sObject.(ImagePullSecrets))
	}
	return nil
}

// ApplyNamespace apply a namespace to a kubernetes cluster
func (k *API) ApplyNamespace(namespace Namespace) error {
	n := namespace.Namespace()
	fmt.Printf("Creating Namespace %s...\n", n.Name)
	exists := true
	if _, err := k.clientset.CoreV1().Namespaces().Get(n.Name, metav1.GetOptions{}); err != nil {
		if !errors.IsNotFound(err) {
			fmt.Printf("error getting namespace: %s", err.Error())
		}
		exists = false
	}
	if !exists {
		_, err := k.clientset.CoreV1().Namespaces().Create(n)
		if err != nil {
			fmt.Printf("error creating namespace: %s\n", err.Error())
		}
	}
	return nil
}

// ApplyDeployment apply a deployment to a kubernetes cluster
func (k *API) ApplyDeployment(deployment Deployment) error {
	fmt.Println("Creating Deployment")
	d := deployment.Deployment()

	exists := true
	if _, err := k.clientset.ExtensionsV1beta1().Deployments(d.Namespace).Get(d.Name, metav1.GetOptions{}); err != nil {
		if !errors.IsNotFound(err) {
			return fmt.Errorf("error getting deployment: %s", err.Error())
		}
		exists = false
	}

	if !exists {
		fmt.Println("Creating Deployment...")
		if _, err := k.clientset.ExtensionsV1beta1().Deployments(d.Namespace).Create(d); err != nil {
			return fmt.Errorf("error creating deployment: %s", err.Error())
		}
	} else {
		fmt.Printf("Deployment %s exists, updating...\n", d.Name)
		if _, err := k.clientset.ExtensionsV1beta1().Deployments(d.Namespace).Update(d); err != nil {
			return fmt.Errorf("error updating deployment: %s", err.Error())
		}
	}
	return nil
}

// ApplyService applies a service to a kubernetes cluster
func (k *API) ApplyService(service Service) error {
	s := service.Service()
	fmt.Printf("Creating Service %s...\n", s.Name)
	exists := true
	if _, err := k.clientset.CoreV1().Services(s.Namespace).Get(s.Name, metav1.GetOptions{}); err != nil {
		if !errors.IsNotFound(err) {
			return fmt.Errorf("error getting service: %s", err.Error())
		}
		exists = false
	}
	if !exists {
		_, err := k.clientset.CoreV1().Services(s.Namespace).Create(s)
		if err != nil {
			return fmt.Errorf("error creating service: %s\n", err.Error())
		}
	}
	return nil
}

// UpdateService updates a service
func (k *API) UpdateService(service Service) error {
	_, err := k.clientset.CoreV1().Services(service.Service().Namespace).Update(service.Service())
	return err
}

// ApplyConfigMap applies a configmap to a kubernetes cluster
func (k *API) ApplyConfigMap(configMap ConfigMap) error {
	c := configMap.ConfigMap()
	fmt.Printf("Creating ConfigMap %s...\n", c.Name)
	exists := true
	if _, err := k.clientset.CoreV1().ConfigMaps(c.Namespace).Get(c.Name, metav1.GetOptions{}); err != nil {
		if !errors.IsNotFound(err) {
			return fmt.Errorf("error getting config map: %s", err.Error())
		}
		exists = false
	}
	if !exists {
		_, err := k.clientset.CoreV1().ConfigMaps(c.Namespace).Create(c)
		if err != nil {
			return fmt.Errorf("error creating config map: %s\n", err.Error())
		}
	} else {
		_, err := k.clientset.CoreV1().ConfigMaps(c.Namespace).Update(c)
		if err != nil {
			return fmt.Errorf("error updating config map: %s\n", err.Error())
		}
	}
	return nil
}

// ApplySecret applies a kubernetes secret to a cluster
func (k *API) ApplySecret(secret Secret) error {
	s := secret.Secret()
	fmt.Printf("Creating Secret %s...\n", s.Name)
	exists := true
	if _, err := k.clientset.CoreV1().Secrets(s.Namespace).Get(s.Name, metav1.GetOptions{}); err != nil {
		if !errors.IsNotFound(err) {
			return fmt.Errorf("error getting secret: %s", err.Error())
		}
		exists = false
	}
	if !exists {
		_, err := k.clientset.CoreV1().Secrets(s.Namespace).Create(s)
		if err != nil {
			return fmt.Errorf("error creating secret: %s\n", err.Error())
		}
	}
	return nil
}

// ApplyHPA applies a horizontal pod autoscaler to a kubernetes cluster
func (k *API) ApplyHPA(hpa HPA) error {
	h := hpa.HPA()
	fmt.Printf("Creating HPA %s...\n", h.Name)
	exists := true
	hpaService := k.clientset.AutoscalingV1().HorizontalPodAutoscalers(h.Namespace)
	if _, err := hpaService.Get(h.Name, metav1.GetOptions{}); err != nil {
		if !errors.IsNotFound(err) {
			return fmt.Errorf("error getting HPA: %s", err.Error())
		}
		exists = false
	}
	if !exists {
		_, err := hpaService.Create(h)
		if err != nil {
			return fmt.Errorf("error creating HPA: %s\n", err.Error())
		}
	} else {
		fmt.Printf("HPA %s exists, updating...\n", h.Name)
		_, err := hpaService.Update(h)
		if err != nil {
			return fmt.Errorf("error creating HPA: %s\n", err.Error())
		}
	}
	return nil
}

// ApplyImagePullSecrets applies pull secrets to a kubernetes cluster
func (k *API) ApplyImagePullSecrets(imagePullSecrets ImagePullSecrets) error {
	secrets := imagePullSecrets.ImagePullSecrets()

	sai := k.clientset.CoreV1().ServiceAccounts(imagePullSecrets.Namespace())
	defaultSA, err := sai.Get("default", metav1.GetOptions{})
	if err != nil {
		return fmt.Errorf("error getting default service account %s", err.Error())
	}
	if len(defaultSA.ImagePullSecrets) >= 1 {
		return nil
	}
	defaultSA.ImagePullSecrets = append(defaultSA.ImagePullSecrets, secrets)
	_, err = sai.Update(defaultSA)
	if err != nil {
		return fmt.Errorf("error updating default service account. %s", err.Error())
	}
	return nil
}

// ListDeployments lists deployments for the namespace (optional, empty string means all), and environment
func (k *API) ListDeployments(namespace string, environment string) ([]v1beta1.Deployment, error) {
	deploys, err := k.clientset.ExtensionsV1beta1().Deployments(namespace).List(metav1.ListOptions{
		LabelSelector: fmt.Sprintf("environment=%s", environment),
	})
	if err != nil {
		return nil, fmt.Errorf("error listing deployments: %s", err.Error())
	}
	return deploys.Items, nil
}

// SetPodNumber set the number desired pods
func (k *API) SetPodNumber(namespace, deploymentName string, replicas int32) error {
	d, err := k.clientset.ExtensionsV1beta1().Deployments(namespace).Get(deploymentName, metav1.GetOptions{})
	if err != nil {
		return err
	}
	d.Spec.Replicas = &replicas
	_, err = k.clientset.ExtensionsV1beta1().Deployments(d.Namespace).Update(d)
	return err
}

// NewK8SApi creates a new k8s api client set
func NewK8SApi(clientset *kubernetes.Clientset) API {
	return API{clientset}
}
