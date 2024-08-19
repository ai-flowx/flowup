package view

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/cligpt/shup/artifact"
	"github.com/cligpt/shup/config"
)

const (
	doneMessage = "shai is installed now. Great!\n"
)

// nolint:mnd
var (
	currentPkgNameStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("211"))
	doneStyle           = lipgloss.NewStyle().Margin(1, 0)
	checkMark           = lipgloss.NewStyle().Foreground(lipgloss.Color("42")).SetString("✓")
)

type PackageModel struct {
	cfg      *config.Config
	channel  string
	packages []string
	index    int
	width    int
	height   int
	spinner  spinner.Model
	done     bool
}

type installedPkgMsg string

// nolint:mnd
func NewPackageModel(cfg *config.Config, channel string, packages []string) PackageModel {
	s := spinner.New()
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("63"))
	return PackageModel{
		cfg:      cfg,
		channel:  channel,
		packages: packages,
		spinner:  s,
	}
}

// nolint:gocritic
func (m PackageModel) Init() tea.Cmd {
	return tea.Batch(m.downloadAndInstall(m.packages[m.index]), m.spinner.Tick)
}

// nolint:gocritic
func (m PackageModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc", "q":
			return m, tea.Quit
		}
	case installedPkgMsg:
		pkg := m.packages[m.index]
		if m.index >= len(m.packages)-1 {
			// Everything's been installed. We're done!
			m.done = true
			return m, tea.Sequence(
				tea.Printf("%s %s", checkMark, pkg), // print the last success message
				tea.Quit,                            // exit the program
			)
		}
		m.index++
		return m, tea.Batch(
			tea.Printf("%s %s", checkMark, pkg),       // print success message above our program
			m.downloadAndInstall(m.packages[m.index]), // download the next package
		)
	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}
	return m, nil
}

// nolint:gocritic
func (m PackageModel) View() string {
	n := len(m.packages)
	w := lipgloss.Width(fmt.Sprintf("%d", n))

	if m.done {
		return doneStyle.Render(fmt.Sprintf(doneMessage))
	}

	pkgCount := fmt.Sprintf(" %*d/%*d", w, m.index, w, n)

	spin := m.spinner.View() + " "
	cellsAvail := max(0, m.width-lipgloss.Width(spin+pkgCount))

	pkgName := currentPkgNameStyle.Render(m.packages[m.index])
	info := lipgloss.NewStyle().MaxWidth(cellsAvail).Render("Installing " + pkgName)

	cellsRemaining := max(0, m.width-lipgloss.Width(spin+info+pkgCount))
	gap := strings.Repeat(" ", cellsRemaining)

	return spin + info + gap + pkgCount
}

// nolint:gocritic
func (m PackageModel) downloadAndInstall(pkg string) tea.Cmd {
	return func() tea.Msg {
		var version string

		ctx := context.Background()

		c := artifact.DefaultConfig()
		c.Config = *m.cfg
		a := artifact.New(ctx, c)

		defer func(a artifact.Artifact, ctx context.Context) {
			_ = a.Deinit(ctx)
		}(a, ctx)

		_ = a.Init(ctx)

		if m.channel == config.ChannelRelease {
			version = strings.Split(pkg, " ")[1]
		} else if m.channel == config.ChannelNightly {
			version = ""
		} else {
			return installedPkgMsg("")
		}

		name := strings.Split(pkg, " ")[0]

		home, _ := os.UserHomeDir()
		install := filepath.Join(home, config.BinName, name)

		if err := a.Fetch(ctx, m.channel, "v"+version, name, install); err != nil {
			return installedPkgMsg("")
		}

		return installedPkgMsg(pkg)
	}
}
