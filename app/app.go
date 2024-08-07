package app

import (
	"sshfolio/ui"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/ssh"
)

func TUIConfig() (tea.Model, []tea.ProgramOption) {
	// Initial model & setup when running the program
	pages := []string{"home", "about", "projects", "contact"}

	// Initializes projects list and itemized list for projects page
	projects := []string{"AEVSoftware", "Devfolify", "Schedulix", "SSHfolio", "Eduquest", "MemeAPI", "WebDevCourse", "Homelab", "ZachLTechWeb"}
	itemizedProjects := []list.Item{
		ui.Item{TitleText: "Alset Solar Cybersedan Software", Desc: "A full stack ecosystem powering the FAUHS AEV solar car"},
		ui.Item{TitleText: "Devfolify", Desc: "Giving you real world problems and a chance to solve it using code in your unique way"},
		ui.Item{TitleText: "Schedulix", Desc: "A program that helps university students develop their course schedules for their upcoming semester"},
		ui.Item{TitleText: "SSHfolio", Desc: "Minimally showing off all your unique talents through a publically SSHable TUI interface written in Go"},
		ui.Item{TitleText: "Eduquest", Desc: "Enjoy a world where education is guided by YOUR passions and preferences."},
		ui.Item{TitleText: "MemeAPI", Desc: "Just a funny API with ElysiaJS because Bun is cool and I had this idea a while ago for fun"},
		ui.Item{TitleText: "WebDevCourse", Desc: "Homemade Web Dev't Course for Friends that doesn't involve 12 hour \"full course\" videos (they're bad)"},
		ui.Item{TitleText: "The Lopez Lab", Desc: "A personal homelab infrastructure I put together for complex code experimentation and hosting almost anything"},
		ui.Item{TitleText: "Personal Digital Branding", Desc: "A collective web of sites and easter eggs I laid across the internet for anyone to explore"},
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
