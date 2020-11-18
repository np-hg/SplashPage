package cmd

// Shamelessly stolen from https://github.com/bbriggs/bitbot/tree/dev/cmd

import (
	"os"

	log "gopkg.in/inconshreveable/log15.v2"

	"splashserver/components"
	"splashserver/server"

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
	Use:     "",
	Short:   "",
	Run: func(cmd *cobra.Command, args []string) {
		s := viper.GetViper().ConfigFileUsed()
		splash := components.NewSplash(s)
		logger.Info("Starting splash page...")
		server.Serve(splash)
	},
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

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./config.yaml)")
	rootCmd.PersistentFlags().StringVarP(&AUTH2_CLIENT_ID, "id", "i", AUTH2_CLIENT_ID, "client ID")
	rootCmd.PersistentFlags().StringVarP(&AUTH2_CLIENT_SECRET, "secret", "s", AUTH2_CLIENT_SECRET, "client secret")
	rootCmd.PersistentFlags().StringVarP(&AUTH2_REDIRECT_URL, "redirect-url", "r", AUTH2_REDIRECT_URL, "Redirect URL")
	rootCmd.PersistentFlags().StringVarP(&AUTH2_PROVIDER_URL, "provider-url", "p", AUTH2_PROVIDER_URL, "Provider URL")
	rootCmd.PersistentFlags().StringSliceVarP(&AUTH2_SCOPES, "scopes", "o", AUTH2_SCOPES, "Comma separated list of scopes")
	rootCmd.PersistentFlags().StringVarP(&AUTH2_STATE, "state", "t", AUTH2_STATE, "Auth2 state")

	viper.BindPFlag("server", rootCmd.PersistentFlags().Lookup("server"))
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
