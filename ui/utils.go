package ui

import (
	"fmt"
	"os"

	"sshfolio/app"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/glamour"
)

// Check err
func Check(e error, check string) {
	if e != nil {
		fmt.Printf("Error running program - In %v: %v", check, e)
	}
}

/******************* Projects list setup ************************/
type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

/******************* Projects list navigation utils ************************/
func OpenProject(selectedProject int, projects []string, viewportWidth int) string {
	for indexedProject, project := range projects {
		if indexedProject == selectedProject {
			rawProjectPageTemplate, _ := glamour.NewTermRenderer(
				glamour.WithStylePath("assets/MDStyle.json"),
				// glamour.WithAutoStyle(), - For Light/Darkmode styling except I'd rather use my custom style
				glamour.WithWordWrap(viewportWidth-20),
			)

			projectPage, err := rawProjectPageTemplate.Render(GetMarkdown("projects/" + project))
			Check(err, "Project Glamour Render")

			return projectPage
		}
	}
	return fmt.Sprintf("Could not get %s project info...", projects[selectedProject])
}

/******************* Page navigation logic ************************/
// Function to read and return markdown file data for each page
func GetMarkdown(filename string) string {
	fileData, err := os.ReadFile("./assets/markdown/" + filename + ".md")
	Check(err, "Markdown File IO")

	return string(fileData)
}
func SaturateContent(m app.Model, viewportWidth int) string {
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
		content, err = rawMarkdownPageTemplate.Render(GetMarkdown("homepage"))
		Check(err, "Gleam Markdown Render")
	case 1: // About
		content, err = rawMarkdownPageTemplate.Render(GetMarkdown("about"))
		Check(err, "Gleam Markdown Render")
	case 3: // Contact
		content, err = rawMarkdownPageTemplate.Render(GetMarkdown("contact"))
		Check(err, "Gleam Markdown Render")
	}

	return content
}

/******************* Help Component Defaults ************************/
// Bubbletea help component full & short displays
func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Navigate, k.RCycle, k.Enter, k.Quit, k.Help}
}
func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.RCycle, k.Enter},
		{k.Up, k.Down},
		{k.LCycle, k.Back},
		{k.Help, k.Quit},
	}
}

/******************* Mouse support utils ************************/
// Gets the location and size of each navigation menu button
// (this is hard coded as of now since I have no idea how to programmatically find a components location & size in the terminal)
func CalculateNavItemSize(title string) (int, int) {
	switch title {
	case "home":
		return 10, 2
	case "about":
		return 10, 2
	case "projects":
		return 13, 2
	case "contact":
		return 12, 2
	default:
		return 0, 0
	}
}

// Max function for viewport line length
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
