package cmd

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/cligpt/shup/config"
	"github.com/cligpt/shup/view"
)

const (
	envMessage = "To get started you may need to restart your current shell.\n" +
		"This would reload your PATH environment variable to include\n" +
		"shai's bin directory ($HOME/.shai/bin).\n" +
		"\n" +
		"To configure your current shell, you need to source\n" +
		"the corresponding env file under $HOME/.shai.\n" +
		"\n" +
		"This is usually done by running one of the following (note the leading DOT):\n" +
		". \"$HOME/.shai/env\""
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
		_, _ = fmt.Fprintln(os.Stdout, envMessage)
	},
}

var (
	updateChannel string
)

// nolint:gochecknoinits
func init() {
	updateCmd.Flags().StringVarP(&updateChannel, "channel", "c", config.ChannelRelease, "update channel")

	rootCmd.AddCommand(updateCmd)
}

func runUpdate(cfg *config.Config) error {
	if _, err := tea.NewProgram(view.NewPackageModel()).Run(); err != nil {
		return err
	}

	return nil
}
