package app

import (
	tea "github.com/charmbracelet/bubbletea"
)

// Empty init for now since there's not much hard logic
func (m Model) Init() tea.Cmd {
	return tea.SetWindowTitle("Welcome to my Portfolio TUI 😄")
}
