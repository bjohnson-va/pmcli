package kubernetes

import (
	"fmt"
	"log"
	"os"
	"os/user"

	"github.com/vendasta/mscli/pkg/spec"
	"github.com/vendasta/mscli/pkg/utils"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp" // this makes the k8s client work locally, it's magic
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func getKubeConfigFile() string {
	var u *user.User
	var err error

	if u, err = user.Current(); err != nil {
		log.Fatalf("Error getting current user: %s", err.Error())
	}

	kubeFile := fmt.Sprintf("%s/.kube/config", u.HomeDir)
	if _, err = os.Stat(kubeFile); err != nil {
		log.Fatalf("Could not find kubernetes config at %s", kubeFile)
	}
	return kubeFile
}

// GetK8sClientSet get k8s client set for the microservice
func GetK8sClientSet(specFile spec.MicroserviceFile, env utils.Environment) (*kubernetes.Clientset, error) {
	clusterConfig, err := rest.InClusterConfig()
	if err != nil {
		cfg := clientcmd.GetConfigFromFileOrDie(getKubeConfigFile())
		se, err := specFile.Microservice.GetEnv(env)
		if err != nil {
			return nil, err
		}
		cfg2 := clientcmd.NewNonInteractiveClientConfig(*cfg, se.K8sContext, &clientcmd.ConfigOverrides{}, nil)
		cfg3, err := cfg2.ClientConfig()
		if err != nil {
			return nil, fmt.Errorf("error creating config: %s", err.Error())
		}
		// creates the clientset
		var clientset *kubernetes.Clientset
		//var err error
		if clientset, err = kubernetes.NewForConfig(cfg3); err != nil {
			return nil, fmt.Errorf("error generating kubernetes clientset: %s", err.Error())
		}

		return clientset, nil
	}
	clientset, err := kubernetes.NewForConfig(clusterConfig)
	if err != nil {
		panic(err)
	}
	return clientset, nil
}
