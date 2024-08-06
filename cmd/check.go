package cmd

import (
	"github.com/spf13/cobra"
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check for updates to toolchains and shup",
	Long:  "Check for updates to toolchains and shup",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

// nolint:gochecknoinits
func init() {
	rootCmd.AddCommand(checkCmd)
}
