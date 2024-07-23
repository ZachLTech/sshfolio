package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
)

const ASCIIName string = `

 ________  ________  ________  ___  ___          ___       ________  ________  _______   ________     
|\_____  \|\   __  \|\   ____\|\  \|\  \        |\  \     |\   __  \|\   __  \|\  ___ \ |\_____  \    
 \|___/  /\ \  \|\  \ \  \___|\ \  \\\  \       \ \  \    \ \  \|\  \ \  \|\  \ \   __/| \|___/  /|   
     /  / /\ \   __  \ \  \    \ \   __  \       \ \  \    \ \  \\\  \ \   ____\ \  \_|/__   /  / /   
    /  /_/__\ \  \ \  \ \  \____\ \  \ \  \       \ \  \____\ \  \\\  \ \  \___|\ \  \_|\ \ /  /_/__  
   |\________\ \__\ \__\ \_______\ \__\ \__\       \ \_______\ \_______\ \__\    \ \_______\\________\
    \|_______|\|__|\|__|\|_______|\|__|\|__|        \|_______|\|_______|\|__|     \|_______|\|_______|
	
	`

type model struct {
	page string
}

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

func initialModel() model {
	return model{
		page: "home",
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

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

func (m model) View() string {

	ui := "\n\n"

	switch m.page {
	case "home":
		ui += `This is the home page (glamour it later)

# Hello World



This is a simple example of Markdown rendering with Glamour!
Check out the [other examples](https://github.com/charmbracelet/glamour/tree/master/examples) too.

Bye!
`
	case "about":
		ui += "This is the about page (glamour it later)"
	case "projects":
		ui += "This is the projects page (bubble list later)"
	case "contact":
		ui += "Contact page under construction..."
	}

	ui, err := glamour.Render(ui, "dark")
	if err != nil {
		fmt.Println("Error running program - In Glamour Render:", err)
		os.Exit(1)
	}

	return ASCIIName + ui
}

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
