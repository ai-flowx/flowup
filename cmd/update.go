package cmd

import (
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update toolchains and shup",
	Long:  "Update toolchains and shup",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

// nolint:gochecknoinits
func init() {
	rootCmd.AddCommand(updateCmd)
}
