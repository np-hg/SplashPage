package cmd

// Shamelessly stolen from https://github.com/bbriggs/bitbot/tree/dev/cmd

import (
	"os"

	log "gopkg.in/inconshreveable/log15.v2"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const version = "0.1"

var (
	AUTH2_CLIENT_ID     string
	AUTH2_CLIENT_SECRET string
	AUTH2_REDIRECT_URL  string
	AUTH2_PROVIDER_URL  string
	AUTH2_SCOPES        []string
	AUTH2_STATE         string
	cfgFile             string
	logger              log.Logger
)

var rootCmd = &cobra.Command{
	Version: version,
	Use: ""
}

// Execute does the thing
func Execute() {
	logger = createLogger()

	if err := rootCmd.Execute(); err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}

func createLogger() log.Logger {
	l := log.New()
	l.SetHandler(log.StreamHandler(os.Stderr, log.JsonFormat()))

	return l
}

func initConfig() {
	if cfgFile == "" {
		viper.AddConfigPath(".")
		viper.SetConfigName("config")
	} else {
		viper.SetConfigFile(cfgFile)
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		logger.Info("Reading config", "config file", viper.ConfigFileUsed())
	}
}
