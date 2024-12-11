package cmd

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/hashicorp/go-version"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/ai-flowx/flowup/artifact"
	"github.com/ai-flowx/flowup/config"
	"github.com/ai-flowx/flowup/view"
)

const (
	envMessage = "To get started you may need to restart your current shell.\n" +
		"This would reload your PATH environment variable to include\n" +
		"flowx's bin directory ($HOME/.flowx/bin).\n" +
		"\n" +
		"To configure your current shell, you need to source\n" +
		"the corresponding env file under $HOME/.flowx.\n" +
		"\n" +
		"This is usually done by running the following:\n" +
		"source \"$HOME/.flowx/env\""
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update toolchains and flowup",
	Long:  "Update toolchains and flowup",
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

var (
	updateChannel string
)

// nolint:gochecknoinits
func init() {
	updateCmd.Flags().StringVarP(&updateChannel, "channel", "c", config.ChannelRelease, "update channel")

	rootCmd.AddCommand(updateCmd)
}

func runUpdate(cfg *config.Config) error {
	var buf []string
	var err error

	fmt.Printf("syncing %s channel updates\n", updateChannel)

	if updateChannel == config.ChannelRelease {
		buf, err = fetchRelease(cfg)
	} else if updateChannel == config.ChannelNightly {
		buf, err = fetchNightly(cfg)
	} else {
		return errors.New("invalid channel")
	}

	if err != nil {
		return errors.Wrap(err, "failed to fetch package")
	}

	if len(buf) == 0 {
		return nil
	}

	if _, err := tea.NewProgram(view.NewPackageModel(cfg, updateChannel, buf)).Run(); err != nil {
		return errors.Wrap(err, "failed to update package")
	}

	_, _ = fmt.Fprintln(os.Stdout, envMessage)

	return nil
}

func fetchRelease(cfg *config.Config) ([]string, error) {
	var ret []string

	home, _ := os.UserHomeDir()
	path := filepath.Join(home, config.BinName)

	localInfo, _ := localToolchain(path)
	remoteInfo, _ := remoteToolchain(cfg)

	buf, _ := matchToolchain(localInfo, remoteInfo)

	for _, item := range buf {
		b := strings.Split(item, " ")
		if len(b) != checkMatchLen {
			continue
		}
		v1, _ := version.NewVersion(b[1])
		v2, _ := version.NewVersion(b[2])
		if v1 != nil && v2 != nil {
			if v1.LessThan(v2) {
				ret = append(ret, fmt.Sprintf("%s %s", b[0], b[2]))
			} else {
				fmt.Printf("✓ %s unchanged\n", b[0])
			}
		} else {
			if v1 == nil && v2 != nil {
				ret = append(ret, fmt.Sprintf("%s %s", b[0], b[2]))
			} else {
				fmt.Printf("✓ %s unchanged\n", b[0])
			}
		}
	}

	return ret, nil
}

func fetchNightly(cfg *config.Config) ([]string, error) {
	ctx := context.Background()

	c := artifact.DefaultConfig()
	c.Config = *cfg
	a := artifact.New(ctx, c)

	defer func(a artifact.Artifact, ctx context.Context) {
		_ = a.Deinit(ctx)
	}(a, ctx)

	_ = a.Init(ctx)

	buf, err := a.Query(ctx, config.ChannelNightly, "")
	if err != nil {
		return []string{}, err
	}

	return buf, nil
}
