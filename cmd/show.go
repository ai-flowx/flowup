package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
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
	color.Set(color.Bold)
	fmt.Printf("default host: ")
	color.Unset()
	fmt.Printf("%s\n", "ubuntu-22.04-x86_64")

	color.Set(color.Bold)
	fmt.Printf("shai home: ")
	color.Unset()
	fmt.Printf("%s\n\n", "/home/zte/.shai")

	color.Set(color.Bold)
	fmt.Printf("installed toolchains\n")
	fmt.Printf("--------------------\n\n")
	color.Unset()

	fmt.Printf("TBD\n")
	fmt.Printf("\n")

	return nil
}
