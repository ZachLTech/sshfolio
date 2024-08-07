package ui

import (
	"github.com/charmbracelet/lipgloss"
)

// For clickable area positioning
var TermHeight int

// Lipgloss styling for view function & nav styling
var (
	NavStyle           = lipgloss.NewStyle().Margin(1, 0).Padding(0, 2)
	ListStyle          = lipgloss.NewStyle().Padding(1, 2)
	BubbleLettersStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#7aa2f7"))
	// For nav text
	ActivePageStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#4fd6be")).Bold(true).PaddingLeft(2).PaddingRight(4)
	InactivePageStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240")).PaddingLeft(4).PaddingRight(4)

	// Border styles
	BorderTitleStyle = func() lipgloss.Style {
		b := lipgloss.HiddenBorder()
		return lipgloss.NewStyle().BorderStyle(b).Padding(0, 1)
	}()
	BorderInfoStyle = func() lipgloss.Style {
		b := lipgloss.HiddenBorder()
		return BorderTitleStyle.BorderStyle(b)
	}()
)
