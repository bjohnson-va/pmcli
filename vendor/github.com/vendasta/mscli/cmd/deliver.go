package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/vendasta/mscli/pkg/ci"
	"github.com/vendasta/mscli/pkg/spec"
)

var (
	zones           []string
	longDescription = fmt.Sprintf(`Delivers a service

Zone Flag:
The zone flag limits which zones in the cluster the image (tag) is deployed to. Multiple zone flags can be speficied.
In order for an image to be deployed to a zone, the zone must appear in the microservice yaml for the specified
environment (env). If no zones flags are set, every zone listed in the environment configuration of the microservice
yaml will be updated. If no zones are configured for an environment, the image will be updated for the deployment in
the default zone (%s).`, spec.DefaultGCPZone)

	deliverCmd = &cobra.Command{
		Use:     "deliver",
		Aliases: []string{"deploy"},
		Short:   "Delivers a service",
		Long:    longDescription,
		RunE:    deliver,
	}
)

func deliver(cmd *cobra.Command, args []string) error {
	return ci.Deliver(microserviceSpec, tag, env, zones)
}

func init() {
	deliverCmd.Flags().StringArrayVarP(&zones, "zone", "z", []string{}, "limits which cluster zones the image is deployed to")
	appCmd.AddCommand(deliverCmd)
}
