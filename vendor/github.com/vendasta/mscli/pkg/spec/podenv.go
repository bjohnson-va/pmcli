package spec

import (
	"fmt"
	"github.com/vendasta/mscli/pkg/utils"
	"strings"
)

func VerifyPodEnv(spec MicroserviceFile, key string, env utils.Environment) error {
	ec, err := spec.Microservice.GetEnv(env)
	if err != nil {
		return err
	}
	for _, e := range ec.PodConfig.PodEnv {
		if e.Key != key {
			continue
		}
		if strings.Contains(e.Value, "TODO") || strings.Trim(e.Value, " \r\n\t") == "" {
			return fmt.Errorf("deploying to non-local environment requires %s podEnv to be set, talk to SRE", key)
		} else {
			//It's cool
			return nil
		}
	}
	return fmt.Errorf("deploying to environment %s requires %s podEnv (missing) to be set, talk to SRE", env.String(), key)
}
