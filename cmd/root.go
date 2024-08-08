package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/cligpt/shup/config"
)

const (
	dirPerm  = 0755
	fileName = ".shai/shup.yml"
	filePerm = 0644
)

var (
	configFile string
)

var rootCmd = &cobra.Command{
	Use:     "shup",
	Version: config.Version + "-build-" + config.Build,
	Short:   "shai installer",
	Long:    fmt.Sprintf("shai installer %s (%s %s)", config.Version, config.Commit, config.Build),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			_ = cmd.Help()
			os.Exit(0)
		}
	},
}

// nolint:gochecknoinits
func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&configFile, "config", "", "config file (default \"$HOME/.shai/shup.yml\")")
	rootCmd.Root().CompletionOptions.DisableDefaultCmd = true
}

func initConfig() {
	helper := func(name string) error {
		if _, err := os.Stat(name); err != nil {
			_ = os.Mkdir(filepath.Dir(name), dirPerm)
			if err := os.WriteFile(name, []byte(config.ConfigData), filePerm); err != nil {
				return err
			}
		}
		return nil
	}

	if configFile == "" {
		home, _ := os.UserHomeDir()
		configFile = filepath.Join(home, fileName)
		if err := helper(configFile); err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err.Error())
			return
		}
	}

	viper.SetConfigFile(configFile)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err.Error())
	}
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
