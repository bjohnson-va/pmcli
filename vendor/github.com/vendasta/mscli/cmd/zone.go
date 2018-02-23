package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/vendasta/mscli/pkg/zone"
)

var (
	zoneLongDesc = `Manage microservice zones

List:
  mscli app zone list

  List will display the available zones a microservice can be in. If the --enabled flag is provided, only zones that the
  microservice is currently serving traffic to will be shown.

Traffic:
  Traffic manages which zones are serving traffic for the microservice. The environment (env) flag is required.

  mscli app zone traffic enable [zone name] (--env=)

  Enable will allow traffic to be served to the given zone. The zone must be specified in the microservice yaml before
  traffic can be served to it.

  mscli app zone traffic disable [zone name] (--env=)

  Disable will stop routing traffic to the specified zone.`
	enabledOnly = false
	zoneCmd     = &cobra.Command{
		Use:   "zone",
		Short: "manage microservice zones",
		Long:  zoneLongDesc,
	}

	zoneListCmd = &cobra.Command{
		Use:   "list",
		Short: "lists all the zones available for the microservice",
		Long:  "lists all the zones available for the microservice",
		RunE:  listZones,
	}

	zoneTrafficCmd = &cobra.Command{
		Use:   "traffic",
		Short: "enable or disable traffic for a zone",
		Long:  "enable or disable traffic for a zone",
	}
	zoneEnableTrafficCmd = &cobra.Command{
		Use:     "enable",
		Short:   "enable traffic for a zone",
		Long:    "enable traffic for a zone",
		RunE:    enableTraffic,
		PreRunE: setTrafficPreRun,
	}
	zoneDisableTrafficCmd = &cobra.Command{
		Use:     "disable",
		Short:   "disable traffic for a zone",
		Long:    "disable traffic for a zone",
		RunE:    disableTraffic,
		PreRunE: setTrafficPreRun,
	}

	tagOnly      = false
	zoneImageCmd = &cobra.Command{
		Use:     "image",
		Short:   "get the current image running in a zone",
		Long:    "get the current image running in a zone",
		RunE:    image,
		PreRunE: imagePreRun,
	}
)

func listZones(cmd *cobra.Command, args []string) error {
	var opts []zone.ListOption
	if enabledOnly {
		opts = append(opts, zone.EnabledOnly())
	}
	zones, err := zone.List(microserviceSpec, env, opts...)
	if err != nil {
		return err
	}
	if len(zones) == 0 {
		fmt.Printf("%s does not support zones\n", env.String())
	}
	for _, z := range zones {
		fmt.Println(z)
	}
	return nil
}

func setTrafficPreRun(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("zone name is required")
	}
	return nil
}

func enableTraffic(cmd *cobra.Command, args []string) error {
	fmt.Printf("enabling %s traffic\n", args[0])
	return zone.SetTraffic(microserviceSpec, env, args[0], true)
}

func disableTraffic(cmd *cobra.Command, args []string) error {
	fmt.Printf("disabling %s traffic\n", args[0])
	return zone.SetTraffic(microserviceSpec, env, args[0], false)
}

func imagePreRun(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("zone name is required")
	}
	return nil
}

func image(cmd *cobra.Command, args []string) error {
	var opts []zone.GetImageOption
	if tagOnly {
		opts = append(opts, zone.TagsOnly())
	}
	img, err := zone.GetImage(&microserviceSpec, env, args[0], opts...)
	if err != nil {
		return err
	}
	fmt.Println(img)
	return nil
}

func init() {
	zoneListCmd.Flags().BoolVar(&enabledOnly, "enabled", false, "show zones which are currently enabled for an environment")
	zoneCmd.AddCommand(zoneListCmd)

	zoneTrafficCmd.AddCommand(zoneEnableTrafficCmd)
	zoneTrafficCmd.AddCommand(zoneDisableTrafficCmd)
	zoneCmd.AddCommand(zoneTrafficCmd)

	zoneImageCmd.Flags().BoolVar(&tagOnly, "tag-only", false, "returns the image tag instead of the full path")
	zoneCmd.AddCommand(zoneImageCmd)

	appCmd.AddCommand(zoneCmd)
}
