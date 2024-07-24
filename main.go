package main

// Imports (useless comment lol)
import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
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
	keys      KeyMap
	help      help.Model
}

// Bubbletea key mapping for help component
type KeyMap struct {
	Left   key.Binding
	Right  key.Binding
	LCycle key.Binding
	RCycle key.Binding
	Enter  key.Binding
	Back   key.Binding
	Help   key.Binding
	Quit   key.Binding
}

var DefaultKeyMap = KeyMap{
	Left: key.NewBinding(
		key.WithKeys("h", "left"),
		key.WithHelp("←/h", "prev page"),
	),
	Right: key.NewBinding(
		key.WithKeys("l", "right"),
		key.WithHelp("→/l", "next page"),
	),
	LCycle: key.NewBinding(
		key.WithKeys("shift+tab"),
		key.WithHelp("^tab", "cycle page prev"),
	),
	RCycle: key.NewBinding(
		key.WithKeys("tab"),
		key.WithHelp("tab", "cycle page next"),
	),
	Enter: key.NewBinding(
		key.WithKeys("enter", " "),
		key.WithHelp("enter", "select"),
	),
	Back: key.NewBinding(
		key.WithKeys("esc", "backspace"),
		key.WithHelp("esc", "go back"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "toggle help"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
}

func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Left, k.Right, k.Help, k.Quit}
}

func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Right, k.Left},
		{k.RCycle, k.LCycle},
		{k.Enter, k.Back, k.Help, k.Quit},
	}
}

// Function to read and return markdown file data for each page
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
		switch {
		case key.Matches(msg, DefaultKeyMap.Quit):
			return m, tea.Quit
		case key.Matches(msg, DefaultKeyMap.Help):
			m.help.ShowAll = !m.help.ShowAll
		case key.Matches(msg, DefaultKeyMap.RCycle):
			return m.cyclePage("right")
		case key.Matches(msg, DefaultKeyMap.LCycle):
			return m.cyclePage("left")
		case key.Matches(msg, DefaultKeyMap.Left):
			if m.pageIndex > 0 {
				m.pageIndex--
			}
			return m, nil
		case key.Matches(msg, DefaultKeyMap.Right):
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

	helpView := m.help.View(m.keys)
	ui, err := glamour.Render(ui, "dark")
	if err != nil {
		fmt.Println("Error running program - In Glamour Render:", err)
		os.Exit(1)
	}

	return ASCIIName + nav + ui + helpView
}

// Starts the Bubbletea TUI & sets up initial state
func main() {

	// Initial model & setup when running the program
	pages := []string{"home", "about", "projects", "contact"}

	initialModel := model{
		pageIndex: 0,
		pages:     pages,
		keys:      DefaultKeyMap,
		help:      help.New(),
	}

	p := tea.NewProgram(initialModel, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
