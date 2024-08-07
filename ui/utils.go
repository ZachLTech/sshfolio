package ui

import (
	"fmt"
	"os"

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
func SaturateContent(m model, viewportWidth int) string {
	// Checks which page the user is on and renders it accordingly
	var content string
	var err error

	rawMarkdownPageTemplate, _ := glamour.NewTermRenderer(
		glamour.WithStylePath("assets/MDStyle.json"),
		// glamour.WithAutoStyle(), - For Light/Darkmode styling except I'd rather use my custom style
		glamour.WithWordWrap(viewportWidth-20),
	)

	switch m.pageIndex {
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

// Bubbletea function to cycle each page (when tab is clicked, this function handles the update event)
func (m model) CyclePage(direction string) model {
	if m.pageIndex < len(m.pages) && direction == "right" {
		switch m.pageIndex {
		case len(m.pages) - 1:
			m.pageIndex = 0
			return m
		default:
			m.pageIndex++
			return m
		}
	} else if m.pageIndex >= 0 && direction == "left" {
		switch m.pageIndex {
		case 0:
			m.pageIndex = len(m.pages) - 1
			return m
		default:
			m.pageIndex--
			return m
		}
	} else {
		return m
	}
}

/******************* Mouse support utils ************************/
// Gets the location and size of each navigation menu button
// (this is hard coded as of now since I have no idea how to programmatically find a components location & size in the terminal)
func (m model) CalculateNavItemPosition(title string) (int, int) {
	startingPoint := m.viewport.Width/2 - 57
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
