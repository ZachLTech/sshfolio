package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	pages       []string
	selectedIdx int
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "left", "h":
			if m.selectedIdx > 0 {
				m.selectedIdx--
			}
		case "right", "l":
			if m.selectedIdx < len(m.pages)-1 {
				m.selectedIdx++
			}
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	var view string
	dot := lipgloss.NewStyle().Foreground(lipgloss.Color("5")).Render("•")

	for i, page := range m.pages {
		style := lipgloss.NewStyle().Foreground(lipgloss.Color("2")).Width(10).Align(lipgloss.Center)
		if i == m.selectedIdx {
			view += fmt.Sprintf(" %s %s ", dot, style.Render(page))
		} else {
			view += fmt.Sprintf("   %s ", style.Render(page))
		}
	}

	return view
}

func main() {
	pages := []string{"Home", "About", "Projects", "Contact"}
	initialModel := model{pages: pages, selectedIdx: 0}

	p := tea.NewProgram(initialModel)
	if err := p.Start(); err != nil {
		fmt.Printf("Error running program: %v", err)
		os.Exit(1)
	}
}
