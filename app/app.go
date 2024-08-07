package app

import (
	"fmt"
	"os"
	"sshfolio/ui"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/ssh"
	"github.com/joho/godotenv"
)

func TUIConfig() (tea.Model, []tea.ProgramOption) {
	// Load dotenv
	err := godotenv.Load()
	ui.Check(err, "Loading .env for TUI Config", true)

	// Initial model & setup when running the program
	pages := []string{"home", "about", "projects", "contact"}

	// Initializes projects list and itemized list for projects page
	projects := []string{}
	itemizedProjects := []list.Item{}

	for i := 1; i >= 1; i++ {
		// Gets formatted .env title per iteration
		ProjectFileTitle := fmt.Sprintf("PROJECT_%d_MARKDOWN_FILE_TITLE", i)
		ProjectDisplayTitle := fmt.Sprintf("PROJECT_%d_DISPLAY_TITLE", i)
		ProjectDescription := fmt.Sprintf("PROJECT_%d_DESCRIPTION", i)

		// Checks if there in fact is a project for this iteration
		_, exists := os.LookupEnv(ProjectDisplayTitle)
		if !exists {
			break
		}

		// Builds the project item for bubbletea list component
		ProjectItem := ui.Item{
			TitleText: os.Getenv(ProjectDisplayTitle),
			Desc:      os.Getenv(ProjectDescription),
		}

		// Appends data to the arrays necessary to build the project list
		projects = append(projects, os.Getenv(ProjectFileTitle))
		itemizedProjects = append(itemizedProjects, ProjectItem)
	}

	initialModel := Model{
		PageIndex:   0,
		Pages:       pages,
		Projects:    projects,
		ProjectOpen: false,
		List:        list.New(itemizedProjects, list.NewDefaultDelegate(), 0, 0),
		Keys:        ui.DefaultKeyMap,
		Help:        help.New(),
	}

	initialModel.List.InfiniteScrolling = true
	initialModel.List.DisableQuitKeybindings()
	initialModel.List.SetFilteringEnabled(false)
	initialModel.List.SetShowHelp(false)
	initialModel.List.SetShowTitle(false)
	initialModel.List.Title = "Hit ENTER for more details"

	return initialModel, []tea.ProgramOption{tea.WithAltScreen(), tea.WithMouseCellMotion()}
}

func SSHTUIConfig(s ssh.Session) (tea.Model, []tea.ProgramOption) {
	initialModel, programOptions := TUIConfig()
	return initialModel, programOptions
}
