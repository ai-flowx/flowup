package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/ai-shflow/shup/config"
)

const (
	showVersionLen = 3
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
	hostInfo, _ := fetchHost()

	color.Set(color.Bold)
	fmt.Print("host: ")
	color.Unset()
	fmt.Printf("%s\n", hostInfo)

	home, _ := os.UserHomeDir()
	path := filepath.Join(home, config.RootName)

	color.Set(color.Bold)
	fmt.Print("home: ")
	color.Unset()
	fmt.Printf("%s\n\n", path)

	color.Set(color.Bold)
	fmt.Print("installed toolchains\n")
	fmt.Print("--------------------\n\n")
	color.Unset()

	path = filepath.Join(home, config.BinName)
	toolchainInfo, _ := fetchToolchain(path)

	for _, item := range toolchainInfo {
		fmt.Println(item)
	}

	fmt.Printf("\n")

	return nil
}

func fetchHost() (string, error) {
	info, err := host.Info()
	if err != nil {
		return "unknown", err
	}

	return fmt.Sprintf("%s-%s-%s", info.Platform, info.PlatformVersion, info.KernelArch), nil
}

func fetchToolchain(path string) ([]string, error) {
	helper := func(file string) (string, error) {
		cmd := exec.Command(file, "--version")
		out, err := cmd.CombinedOutput()
		if err != nil {
			return "", err
		}
		buf := strings.Split(string(out), " ")
		if len(buf) != showVersionLen {
			return "", err
		}
		// shup version v1.0.0-build-2024-08-16T11:33:19+0800
		name := buf[0]
		version := strings.TrimPrefix(strings.Split(buf[2], "-")[0], "v")
		return fmt.Sprintf("%s %s", name, version), nil
	}

	var files []string
	var buf []string

	_ = filepath.Walk(path, func(p string, _ os.FileInfo, _ error) error {
		files = append(files, p)
		return nil
	})

	for _, item := range files {
		if b, err := helper(item); err == nil {
			buf = append(buf, b)
		}
	}

	return buf, nil
}
