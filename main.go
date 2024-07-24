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
	page string
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
func (m model) cyclePage() (tea.Model, tea.Cmd) {
	switch m.page {
	case "home":
		m.page = "about"
		return m, nil
	case "about":
		m.page = "projects"
		return m, nil
	case "projects":
		m.page = "contact"
		return m, nil
	case "contact":
		m.page = "home"
		return m, nil
	default:
		return m, nil
	}
}

// Initial model when running the program
func initialModel() model {
	return model{
		page: "home",
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
			return m.cyclePage()
		}
	}

	return m, nil
}

// Switch case with each page/TUI view
func (m model) View() string {

	ui := "\n\n"

	switch m.page {
	case "home":
		ui += getMarkdown("homepage")
	case "about":
		ui += getMarkdown("about")
	case "projects":
		ui += "This is the projects page which is under construction... (bubble list later)"
	case "contact":
		ui += getMarkdown("contact")
	}

	ui, err := glamour.Render(ui, "dark")
	if err != nil {
		fmt.Println("Error running program - In Glamour Render:", err)
		os.Exit(1)
	}

	return ASCIIName + ui
}

// Starts the Bubbletea TUI
func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
