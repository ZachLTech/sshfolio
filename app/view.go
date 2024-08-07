package app

import (
	"sshfolio/ui"

	"github.com/charmbracelet/lipgloss"
)

// Switch case with each page/TUI view
func (m Model) View() string {

	// If viewport isn't ready it'll say welcome (this should only be able to happen during startup)
	if !m.Ready {
		return "\n  Welcome..."
	}

	nav := `` // Empty to be saturated soon

	// Render/create nav depending on page location
	for i, title := range m.Pages {
		if i == m.PageIndex {
			// Highlight the active page
			nav += ui.ActivePageStyle.Render("• " + title + " ")
		} else {
			nav += ui.InactivePageStyle.Render(title + " ")
		}
	}

	m.Content = m.ViewportHeader(m.Pages[m.PageIndex]) + m.Viewport.View() + m.ViewportFooter()
	if m.PageIndex == 2 {
		if !m.ProjectOpen {
			m.ProjectView = ui.ListStyle.Render(m.List.View())
		} else if m.ProjectOpen {
			m.ProjectView = m.Viewport.View()
		}
		m.Content = m.ViewportHeader(m.Pages[m.PageIndex]) + m.ProjectView + m.ViewportFooter()
	}

	header := lipgloss.PlaceHorizontal(m.Viewport.Width, lipgloss.Center, ui.BubbleLettersStyle.Render(ui.GetHeader()))
	headerMessage := lipgloss.PlaceHorizontal(m.Viewport.Width, lipgloss.Center, ui.BubbleLettersStyle.Render(ui.GetHeaderMessage()))
	nav = lipgloss.PlaceHorizontal(m.Viewport.Width, lipgloss.Center, ui.NavStyle.Render(nav))

	return header + headerMessage + nav + m.Content + ui.NavStyle.Render(m.Help.View(m.Keys)) // NAME TITLE + NAVIGATION + MAIN PAGE + FOOTER/HELP
}
