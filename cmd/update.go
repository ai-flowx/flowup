package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/cligpt/shup/config"
	"github.com/cligpt/shup/view"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update toolchains and shup",
	Long:  "Update toolchains and shup",
	Run: func(cmd *cobra.Command, args []string) {
		var cfg config.Config
		err := viper.Unmarshal(&cfg)
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err.Error())
			return
		}
		if err := runUpdate(&cfg); err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err.Error())
			return
		}
	},
}

// nolint:gochecknoinits
func init() {
	rootCmd.AddCommand(updateCmd)
}

func runUpdate(cfg *config.Config) error {
	if _, err := tea.NewProgram(view.NewPackageModel()).Run(); err != nil {
		return err
	}

	return nil
}
