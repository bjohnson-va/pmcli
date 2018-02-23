package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/vendasta/mscli/pkg/spec"
	"github.com/vendasta/mscli/pkg/utils"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

var cliEnv string
var env utils.Environment
var tag string
var cfgFile string
var microserviceSpec spec.MicroserviceFile

var appCmd = &cobra.Command{
	Use:    "app",
	Short: "Commands that require a microservice.yaml (test, build, deploy, etc.)",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if cliEnv != "" {
			env = utils.EnvFromString(cliEnv)
		}
		return parseConfig(cfgFile)
	},
}

func parseConfig(cfgFile string) error {
	// get the filepath
	abs, err := filepath.Abs(cfgFile)
	if err != nil {
		return fmt.Errorf("error reading filepath: %s", err.Error())
	}

	// get the config name
	base := filepath.Base(abs)

	// get the path
	path := filepath.Dir(abs)

	// set up viper
	viper.SetConfigName(strings.Split(base, ".")[0])
	viper.SetConfigType(strings.Split(base, ".")[1])
	viper.AddConfigPath(path)

	// Find and read the config file; Handle errors reading the config file
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("failed to read config file: %s", err.Error())
	}

	m := spec.MicroserviceFile{}
	err = viper.Unmarshal(&m)
	if err != nil {
		return err
	}

	microserviceSpec = m
	microserviceSpec.Microservice.Debug = debug
	if microserviceSpec.Microservice.IdentityType == "" {
		microserviceSpec.Microservice.IdentityType = spec.IdentityServiceAccountJwt
	}

	return nil
}

func init() {
	RootCmd.AddCommand(appCmd)
	appCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "./microservice.yaml", "config file")
	appCmd.PersistentFlags().StringVarP(&cliEnv, "env", "e", "local", "environment to target")
	appCmd.PersistentFlags().StringVarP(&tag, "tag", "t", strconv.FormatInt(time.Now().UTC().Unix(), 10), "image target")
}
