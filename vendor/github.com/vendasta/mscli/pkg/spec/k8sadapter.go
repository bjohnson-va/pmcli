package spec

import (
	autoscaling "k8s.io/api/autoscaling/v1"
	"k8s.io/api/core/v1"
	"k8s.io/api/extensions/v1beta1"
)

// namespace implements k8s.Namespace
type namespace struct {
	namespace *v1.Namespace
}

func (n namespace) K8SType() string {
	return "namespace"
}

func (n namespace) Namespace() *v1.Namespace {
	return n.namespace
}

// deployment implements k8s.Deployment
type deployment struct {
	deployment *v1beta1.Deployment
}

func (d deployment) K8SType() string {
	return "deployment"
}

func (d deployment) Deployment() *v1beta1.Deployment {
	return d.deployment
}

// service implements k8s.Service
type service struct {
	service *v1.Service
}

func (s service) K8SType() string {
	return "service"
}

func (s service) Service() *v1.Service {
	return s.service
}

// configMap implements k8s.ConfigMap
type configMap struct {
	configMap *v1.ConfigMap
}

func (c configMap) K8SType() string {
	return "ConfigMap"
}

func (c configMap) ConfigMap() *v1.ConfigMap {
	return c.configMap
}

// secret implements k8s.Secret
type secret struct {
	secret *v1.Secret
}

func (s secret) K8SType() string {
	return "Secret"
}

func (s secret) Secret() *v1.Secret {
	return s.secret
}

// hpa implements k8s.HPA
type hpa struct {
	hpa *autoscaling.HorizontalPodAutoscaler
}

func (h hpa) K8SType() string {
	return "HPA"
}

func (h hpa) HPA() *autoscaling.HorizontalPodAutoscaler {
	return h.hpa
}

// imagePullSecrets implements k8s.ImagePullSecrets
type imagePullSecrets struct {
	namespace        string
	imagePullSecrets v1.LocalObjectReference
}

func (i imagePullSecrets) K8SType() string {
	return "ImagePullSecrets"
}

func (i imagePullSecrets) ImagePullSecrets() v1.LocalObjectReference {
	return i.imagePullSecrets
}

func (i imagePullSecrets) Namespace() string {
	return i.namespace
}
