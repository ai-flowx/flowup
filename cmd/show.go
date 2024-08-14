package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/cligpt/shup/config"
)

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show the installed toolchains",
	Long:  "Show the installed toolchains",
	Run: func(cmd *cobra.Command, args []string) {
		var cfg config.Config
		err := viper.Unmarshal(&cfg)
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err.Error())
			return
		}
		if err := runShow(&cfg); err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err.Error())
			return
		}
	},
}

// nolint:gochecknoinits
func init() {
	rootCmd.AddCommand(showCmd)
}

func runShow(_ *config.Config) error {
	info, _ := hostInfo()

	color.Set(color.Bold)
	fmt.Printf("host: ")
	color.Unset()
	fmt.Printf("%s\n", info)

	home, _ := os.UserHomeDir()
	path := filepath.Join(home, config.RootName)

	color.Set(color.Bold)
	fmt.Printf("home: ")
	color.Unset()
	fmt.Printf("%s\n\n", path)

	color.Set(color.Bold)
	fmt.Printf("installed toolchains\n")
	fmt.Printf("--------------------\n\n")
	color.Unset()

	fmt.Printf("shai 1.0.0\n")
	fmt.Printf("gitgpt 1.0.0\n")
	fmt.Printf("lintgpt 1.0.0\n")
	fmt.Printf("metalgpt 1.0.0\n")
	fmt.Printf("\n")

	return nil
}

func hostInfo() (string, error) {
	info, err := host.Info()
	if err != nil {
		return "unknown", err
	}

	return fmt.Sprintf("%s-%s-%s", info.Platform, info.PlatformVersion, info.KernelArch), nil
}
