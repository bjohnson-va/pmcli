package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vendasta/mscli/pkg/dns"
)

var dnsCmd = &cobra.Command{
	Use:   "dns",
	Short: "DNS related commands",
	Long:  "DNS related commands",
}

var dnsConfigureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Configures all DNS settings for you",
	Long:  "Configures all DNS settings for you",
	RunE:  dnsConfigure,
}

func dnsConfigure(cmd *cobra.Command, args []string) error {
	grpcIP, httpIP, err := dns.ConfigureDNS(microserviceSpec, env)
	if err != nil {
		return err
	}
	return dns.WriteLoadBalancerIP(microserviceSpec, env, grpcIP, httpIP)
}

var dnsTestCmd = &cobra.Command{
	Use:   "test",
	Short: "Tests the DNS settings",
	Long:  "Tests the DNS settings",
	RunE:  dnsTest,
}

func dnsTest(cmd *cobra.Command, args []string) error {
	return dns.TestDNS(microserviceSpec, env)
}

func init() {
	dnsCmd.AddCommand(dnsConfigureCmd, dnsTestCmd)
	appCmd.AddCommand(dnsCmd)
}
