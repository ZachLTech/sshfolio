package ui

/******************* This file correlates with sshfolio/app/methods.go & is for any helper or utility function that doesn't need to define a method on the bubbletea model ************************/

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/glamour"
	"github.com/common-nighthawk/go-figure"
	"github.com/joho/godotenv"
)

// Check err
func Check(e error, check string, fatal bool) {
	if e != nil {
		fmt.Printf("Error running program - In %v: %v", check, e)
		if fatal {
			os.Exit(1)
		}
	}
}

/******************* Projects list setup ************************/
type Item struct {
	TitleText, Desc string
}

func (i Item) Title() string       { return i.TitleText }
func (i Item) Description() string { return i.Desc }
func (i Item) FilterValue() string { return i.TitleText }

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
			Check(err, "Project Glamour Render", false)

			return projectPage
		}
	}
	return fmt.Sprintf("Could not get %s project info...", projects[selectedProject])
}

/******************* Page navigation logic ************************/
// Function to read and return markdown file data for each page
func GetMarkdown(filename string) string {
	fileData, err := os.ReadFile("./assets/markdown/" + filename + ".md")
	Check(err, "Markdown File IO", false)

	return string(fileData)
}

/******************* Help Component Defaults ************************/
// Bubbletea help component full & short displays
func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Navigate, k.RCycle, k.Enter, k.Back, k.Quit, k.Help}
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

/******************* Header setup utils ************************/
func CountLines(s string) int {
	// Split the string by the newline character
	lines := strings.Split(s, "\n")
	// Return the number of lines
	return len(lines)
}
func GetHeader() string {
	// Load dotenv
	err := godotenv.Load()
	Check(err, "Loading .env for Header", true)

	title := figure.NewFigure(strings.ToUpper(os.Getenv("HEADER")), "larry3d", true)

	return fmt.Sprintf("\n%v", title.String())
}
func GetHeaderMessage() string {
	// Load dotenv
	err := godotenv.Load()
	Check(err, "Loading .env for Header", true)

	messasge := os.Getenv("HEADER_MESSAGE")

	return fmt.Sprintf("\n%v\n", messasge)
}
