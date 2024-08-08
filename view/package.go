package view

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	doneMessage = "shai is installed now. Great!\n"
)

var packages = []string{
	"shai",
	"gitgpt",
	"lintgpt",
	"metalgpt",
}

// nolint:mnd
var (
	currentPkgNameStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("211"))
	doneStyle           = lipgloss.NewStyle().Margin(1, 0)
	checkMark           = lipgloss.NewStyle().Foreground(lipgloss.Color("42")).SetString("✓")
)

type PackageModel struct {
	packages []string
	index    int
	width    int
	height   int
	spinner  spinner.Model
	done     bool
}

type installedPkgMsg string

// nolint:mnd
func NewPackageModel() PackageModel {
	s := spinner.New()
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("63"))
	return PackageModel{
		packages: getPackages(),
		spinner:  s,
	}
}

// nolint:gocritic
func (m PackageModel) Init() tea.Cmd {
	return tea.Batch(downloadAndInstall(m.packages[m.index]), m.spinner.Tick)
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
			tea.Printf("%s %s", checkMark, pkg),     // print success message above our program
			downloadAndInstall(m.packages[m.index]), // download the next package
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

// nolint:gosec,mnd
func downloadAndInstall(pkg string) tea.Cmd {
	// This is where you'd do i/o stuff to download and install packages. In
	// our case we're just pausing for a moment to simulate the process.
	d := time.Millisecond * time.Duration(rand.Intn(500))
	return tea.Tick(d, func(t time.Time) tea.Msg {
		return installedPkgMsg(pkg)
	})
}

// nolint:gosec,mnd
func getPackages() []string {
	pkgs := packages
	copy(pkgs, packages)

	rand.Shuffle(len(pkgs), func(i, j int) {
		pkgs[i], pkgs[j] = pkgs[j], pkgs[i]
	})

	for k := range pkgs {
		pkgs[k] += fmt.Sprintf(" %d.%d.%d", 1, 0, 0)
	}

	return pkgs
}
