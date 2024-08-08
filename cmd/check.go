package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/cligpt/shup/config"
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check for updates to toolchains and shup",
	Long:  "Check for updates to toolchains and shup",
	Run: func(cmd *cobra.Command, args []string) {
		var cfg config.Config
		err := viper.Unmarshal(&cfg)
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err.Error())
			return
		}
		if err := runCheck(&cfg); err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err.Error())
			return
		}
	},
}

// nolint:gochecknoinits
func init() {
	rootCmd.AddCommand(checkCmd)
}

func runCheck(cfg *config.Config) error {
	color.Set(color.Bold)
	fmt.Printf("shai - ")
	fmt.Printf(color.YellowString("update available "))
	color.Unset()
	fmt.Printf(": 1.0.0 -> 1.1.0\n")

	color.Set(color.Bold)
	fmt.Printf("gitgpt - ")
	fmt.Printf(color.YellowString("update available "))
	color.Unset()
	fmt.Printf(": 1.0.0 -> 1.1.0\n")

	color.Set(color.Bold)
	fmt.Printf("lintgpt - ")
	fmt.Printf(color.YellowString("update available "))
	color.Unset()
	fmt.Printf(": 1.0.0 -> 1.1.0\n")

	color.Set(color.Bold)
	fmt.Printf("metalgpt - ")
	fmt.Printf(color.GreenString("up to date "))
	color.Unset()
	fmt.Printf(": 1.1.0\n")

	return nil
}
