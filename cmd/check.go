package cmd

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	"github.com/hashicorp/go-version"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/cligpt/shup/artifact"
	"github.com/cligpt/shup/config"
)

const (
	checkMatchLen   = 3
	checkVersionLen = 2
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
		n := b[0]
		v1, _ := version.NewVersion(b[1])
		v2, _ := version.NewVersion(b[2])
		color.Set(color.Bold)
		fmt.Printf("%s - ", n)
		if v1 != nil && v2 != nil {
			if v1.LessThan(v2) {
				fmt.Print(color.YellowString("update available "))
				color.Unset()
				fmt.Printf(": %s -> %s\n", b[1], b[2])
			}
		} else {
			if v1 == nil && v2 != nil {
				fmt.Print(color.YellowString("update available "))
				color.Unset()
				fmt.Printf(": %s -> %s\n", b[1], b[2])
			} else {
				fmt.Print(color.GreenString("up to date "))
				color.Unset()
				fmt.Printf(": %s\n", b[1])
			}
		}
	}

	return nil
}

func localToolchain(path string) ([]string, error) {
	return fetchToolchain(path)
}

func remoteToolchain(cfg *config.Config) ([]string, error) {
	latest := func(ver []string) (string, error) {
		ret, _ := version.NewVersion("0.0.0")
		for _, item := range ver {
			b := strings.TrimPrefix(item, "v")
			if v, e := version.NewVersion(b); e == nil {
				if ret.LessThan(v) {
					ret = v
				}
			}
		}
		return ret.String(), nil
	}

	rebuild := func(files []string, ver string) []string {
		var buf []string
		for _, item := range files {
			buf = append(buf, fmt.Sprintf("%s %s", item, ver))
		}
		return buf
	}

	var buf []string

	ctx := context.Background()

	c := artifact.DefaultConfig()
	c.Config = *cfg
	a := artifact.New(ctx, c)

	defer func(a artifact.Artifact, ctx context.Context) {
		_ = a.Deinit(ctx)
	}(a, ctx)

	_ = a.Init(ctx)

	buf, err := a.Query(ctx, config.ChannelRelease, "")
	if err != nil {
		return []string{}, err
	}

	ver, err := latest(buf)
	if err != nil {
		return []string{}, err
	}

	buf, err = a.Query(ctx, config.ChannelRelease, "v"+ver)
	if err != nil {
		return []string{}, err
	}

	return rebuild(buf, ver), nil
}

func matchToolchain(local, remote []string) ([]string, error) {
	helper := func(local []string, key string) (string, error) {
		var found string
		k := strings.Split(key, " ")
		if len(k) != checkVersionLen {
			return "", errors.New("invalid key")
		}
		for _, item := range local {
			n := strings.Split(item, " ")
			if len(n) != checkVersionLen {
				continue
			}
			if k[0] == n[0] {
				found = fmt.Sprintf("%s %s %s", k[0], n[1], k[1])
				break
			}
		}
		if found == "" {
			found = fmt.Sprintf("%s %s %s", k[0], "", k[1])
		}
		return found, nil
	}

	var buf []string

	for _, item := range remote {
		if found, err := helper(local, item); err == nil {
			buf = append(buf, found)
		}
	}

	return buf, nil
}
