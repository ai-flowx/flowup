package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/cligpt/shup/config"
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check for updates to toolchains and shup",
	Long:  "Check for updates to toolchains and shup",
	Run: func(cmd *cobra.Command, args []string) {
		var _config config.Config
		err := viper.Unmarshal(&_config)
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err.Error())
			return
		}
	},
}

// nolint:gochecknoinits
func init() {
	rootCmd.AddCommand(checkCmd)
}
