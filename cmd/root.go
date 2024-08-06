package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/cligpt/shup/config"
)

var (
	configFile string
)

var rootCmd = &cobra.Command{
	Use:     "shup",
	Version: config.Version + "-build-" + config.Build,
	Short:   "shai installer",
	Long:    "shai installer",
	Run:     func(cmd *cobra.Command, args []string) {},
}

// nolint:gochecknoinits
func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&configFile, "config", "$HOME/.shai/shup.yml", "config file")
	rootCmd.Root().CompletionOptions.DisableDefaultCmd = true
}

func initConfig() {
	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.AddConfigPath(".shai")
		viper.SetConfigName("shup")
		viper.SetConfigType("yml")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Printf("Using config file: %v\n", viper.ConfigFileUsed())
	} else {
		fmt.Println(err.Error())
	}
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
