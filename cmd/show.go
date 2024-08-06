package cmd

import (
	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show the active and installed toolchains",
	Long:  "Show the active and installed toolchains",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

// nolint:gochecknoinits
func init() {
	rootCmd.AddCommand(showCmd)
}
