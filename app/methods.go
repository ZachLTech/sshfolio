package app

/******************* This file correlates with sshfolio/ui/utils.go & is for any function that defines a method on the bubbletea model ************************/

import (
	"fmt"
	"strings"

	"sshfolio/ui"

	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
)

/******************* Viewport Header/Footer Render ************************/
// Page viewport header and footer render
func (m Model) ViewportHeader(pageTitle string) string {
	title := ui.BorderTitleStyle.Render(pageTitle)
	line := strings.Repeat("─", max(0, m.Viewport.Width-lipgloss.Width(title)))
	return lipgloss.JoinHorizontal(lipgloss.Center, title, line)
}
func (m Model) ViewportFooter() string {
	info := ui.BorderInfoStyle.Render(fmt.Sprintf("%3.f%%", m.Viewport.ScrollPercent()*100))
	line := strings.Repeat("─", max(0, m.Viewport.Width-lipgloss.Width(info)))
	return lipgloss.JoinHorizontal(lipgloss.Center, line, info)
}

/******************* Page navigation logic ************************/
// Bubbletea function to cycle each page (when tab is clicked, this function handles the update event)
func (m Model) CyclePage(direction string) Model {
	if m.PageIndex < len(m.Pages) && direction == "right" {
		switch m.PageIndex {
		case len(m.Pages) - 1:
			m.PageIndex = 0
			return m
		default:
			m.PageIndex++
			return m
		}
	} else if m.PageIndex >= 0 && direction == "left" {
		switch m.PageIndex {
		case 0:
			m.PageIndex = len(m.Pages) - 1
			return m
		default:
			m.PageIndex--
			return m
		}
	} else {
		return m
	}
}
func SaturateContent(m Model, viewportWidth int) string {
	// Checks which page the user is on and renders it accordingly
	var content string
	var err error

	rawMarkdownPageTemplate, _ := glamour.NewTermRenderer(
		glamour.WithStylePath("assets/MDStyle.json"),
		// glamour.WithAutoStyle(), - For Light/Darkmode styling except I'd rather use my custom style
		glamour.WithWordWrap(viewportWidth-20),
	)

	switch m.PageIndex {
	case 0: // Home
		content, err = rawMarkdownPageTemplate.Render(ui.GetMarkdown("homepage"))
		ui.Check(err, "Gleam Markdown Render", false)
	case 1: // About
		content, err = rawMarkdownPageTemplate.Render(ui.GetMarkdown("about"))
		ui.Check(err, "Gleam Markdown Render", false)
	case 3: // Contact
		content, err = rawMarkdownPageTemplate.Render(ui.GetMarkdown("contact"))
		ui.Check(err, "Gleam Markdown Render", false)
	}

	return content
}

/******************* Mouse support utils ************************/
// Gets the location and size of each navigation menu button
// (this is hard coded as of now since I have no idea how to programmatically find a components location & size in the terminal)
func (m Model) CalculateNavItemPosition(title string) (int, int) {
	startingPoint := m.Viewport.Width/2 - 57
	switch title {
	case "home":
		return startingPoint + 30, 9
	case "about":
		return startingPoint + 43, 9
	case "projects":
		return startingPoint + 58, 9
	case "contact":
		return startingPoint + 75, 9
	default:
		return 0, 0
	}
}
