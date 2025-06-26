package views

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
)

type ProgressMessage struct {
	Language string
	Current  int
	Total    int
	Done     bool
}

type langProgess struct {
	Bar     progress.Model
	Current int
	Total   int
	Done    bool
}

type Model struct {
	langs map[string]*langProgess
	order []string
}

func NewModel(langTargets map[string]int) Model {
	m := Model{
		langs: make(map[string]*langProgess),
		order: make([]string, 0),
	}
	for lang, total := range langTargets {
		m.langs[lang] = &langProgess{
			Bar: progress.New(progress.WithWidth(40),
				progress.WithDefaultGradient()),
			Total:   total,
			Current: 0,
		}
		m.order = append(m.order, lang)
	}
	return m
}
func (m Model) Init() tea.Cmd {
	// Initialize with tick commands for all progress bars
	var cmds []tea.Cmd
	for _, lp := range m.langs {
		cmds = append(cmds, lp.Bar.Init())
	}
	return tea.Batch(cmds...)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case ProgressMessage:
		lp := m.langs[msg.Language]
		lp.Current = msg.Current
		lp.Done = msg.Done
		lp.Total = msg.Total

		percent := float64(lp.Current) / float64(lp.Total)
		cmd := lp.Bar.SetPercent(percent)

		// Check if all languages are done using percentage
		allDone := true
		for _, lang := range m.order {
			langProgress := m.langs[lang]
			if langProgress.Total > 0 {
				completion := float64(langProgress.Current) / float64(langProgress.Total)
				if completion < 1.0 { // Not 100% complete
					allDone = false
					break
				}
			} else {
				allDone = false // No target set yet
				break
			}
		}

		if allDone {
			return m, tea.Batch(cmd, tea.Quit)
		}

		return m, cmd

	case tea.KeyMsg:
		if msg.String() == "q" || msg.String() == "ctrl+c" {
			return m, tea.Quit
		}

	case progress.FrameMsg:
		// Update all progress bars with the frame message
		var cmds []tea.Cmd
		for _, lp := range m.langs {
			newModel, cmd := lp.Bar.Update(msg)
			lp.Bar = newModel.(progress.Model)
			cmds = append(cmds, cmd)
		}
		return m, tea.Batch(cmds...)
	}

	return m, nil
}

func (m Model) View() string {
	var b strings.Builder
	b.WriteString("📦 Training Progress:\n\n")
	for _, lang := range m.order {
		lp := m.langs[lang]
		status := "⏳"
		if lp.Done {
			status = "✅"
		}
		fmt.Fprintf(&b, "[%-10s] %s %4d / %4d %s\n", lang, lp.Bar.View(), lp.Current, lp.Total, status)
	}
	return b.String()

}
