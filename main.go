package main

// Imports (useless comment lol)
import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
)

// ASCII text that should be displayed through every page
const ASCIIName string = `

 ________  ________  ________  ___  ___          ___       ________  ________  _______   ________     
|\_____  \|\   __  \|\   ____\|\  \|\  \        |\  \     |\   __  \|\   __  \|\  ___ \ |\_____  \    
 \|___/  /\ \  \|\  \ \  \___|\ \  \\\  \       \ \  \    \ \  \|\  \ \  \|\  \ \   __/| \|___/  /|   
     /  / /\ \   __  \ \  \    \ \   __  \       \ \  \    \ \  \\\  \ \   ____\ \  \_|/__   /  / /   
    /  /_/__\ \  \ \  \ \  \____\ \  \ \  \       \ \  \____\ \  \\\  \ \  \___|\ \  \_|\ \ /  /_/__  
   |\________\ \__\ \__\ \_______\ \__\ \__\       \ \_______\ \_______\ \__\    \ \_______\\________\
    \|_______|\|__|\|__|\|_______|\|__|\|__|        \|_______|\|_______|\|__|     \|_______|\|_______|
	
	`

// Bubbletea model structure
type model struct {
	pageIndex int
	pages     []string
}

func getMarkdown(filename string) string {
	fileData, err := os.ReadFile("./assets/markdown/" + filename + ".md")
	check(err)

	return string(fileData)
}

// Check function for markdown file IO
func check(e error) {
	if e != nil {
		fmt.Println("Error running program - In Markdown File IO:", e)
		os.Exit(1)
	}
}

// Bubbletea function to cycle each page (when tab is clicked, this function handles the update event)
func (m model) cyclePage(direction string) (tea.Model, tea.Cmd) {
	if m.pageIndex < len(m.pages) && direction == "right" {
		switch m.pageIndex {
		case len(m.pages) - 1:
			m.pageIndex = 0
			return m, nil
		default:
			m.pageIndex++
			return m, nil
		}
	} else if m.pageIndex >= 0 && direction == "left" {
		switch m.pageIndex {
		case 0:
			m.pageIndex = len(m.pages) - 1
			return m, nil
		default:
			m.pageIndex--
			return m, nil
		}
	} else {
		return m, nil
	}
}

// Empty init for now since there's not much hard logic
func (m model) Init() tea.Cmd {
	return nil
}

// Bubbletea update/msg handling
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "tab":
			return m.cyclePage("right")
		case "shift+tab":
			return m.cyclePage("left")
		case "left":
			if m.pageIndex > 0 {
				m.pageIndex--
			}
			return m, nil
		case "right":
			if m.pageIndex < len(m.pages)-1 {
				m.pageIndex++
			}
			return m, nil
		}
	}

	return m, nil
}

// Switch case with each page/TUI view
func (m model) View() string {

	nav := ``
	ui := "\n\n"

	switch m.pageIndex {
	case 0: // Home
		nav = ``
		ui += getMarkdown("homepage")
	case 1: // About
		nav = ``
		ui += getMarkdown("about")
	case 2: // Projects
		nav = ``
		ui += "This is the projects page which is under construction... (bubble list later)"
	case 3: // Contact
		nav = ``
		ui += getMarkdown("contact")
	}

	ui, err := glamour.Render(ui, "dark")
	if err != nil {
		fmt.Println("Error running program - In Glamour Render:", err)
		os.Exit(1)
	}

	return ASCIIName + nav + ui
}

// Starts the Bubbletea TUI & sets up initial state
func main() {

	// Initial model & setup when running the program
	pages := []string{"home", "about", "projects", "contact"}

	initialModel := model{
		pageIndex: 0,
		pages:     pages,
	}

	p := tea.NewProgram(initialModel, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
