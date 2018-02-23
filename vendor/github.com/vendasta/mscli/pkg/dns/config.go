package dns

import (
	"github.com/spf13/viper"
	"github.com/vendasta/mscli/pkg/spec"
	"github.com/vendasta/mscli/pkg/utils"
)

// WriteLoadBalancerIP set the load balancer IP in the microservice config
func WriteLoadBalancerIP(specFile spec.MicroserviceFile, env utils.Environment, grpcIP string, httpsIP string) error {
	e, err := specFile.Microservice.GetEnv(env)
	if err != nil {
		return err
	}
	e.Network.HTTPSLoadBalancerIP = httpsIP
	e.Network.GRPCLoadBalancerIP = grpcIP

	err = specFile.Microservice.ReplaceEnv(e)
	if err != nil {
		return err
	}

	viper.Set("microservice.environments", specFile.Microservice.Environments)
	return viper.WriteConfig()
}
