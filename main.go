package main

// Imports (useless comment lol)
import (
	"context"
	"errors"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/charmbracelet/log"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	"github.com/charmbracelet/wish/activeterm"
	"github.com/charmbracelet/wish/bubbletea"
	"github.com/charmbracelet/wish/logging"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
)

const (
	host = "0.0.0.0"
	port = "22220"
)

// Starts the Bubbletea TUI & sets up initial state
func main() {

	server, err := wish.NewServer(
		wish.WithAddress(net.JoinHostPort(host, port)),
		wish.WithHostKeyPath(".ssh/id_ed25519"),
		wish.WithMiddleware(
			bubbletea.Middleware(teaHandler),
			activeterm.Middleware(), // Bubble Tea apps usually require a PTY.
			logging.Middleware(),
		),
	)
	if err != nil {
		log.Error("Could not start server", "error", err)
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	log.Info("Starting SSH server", "host", host, "port", port)
	go func() {
		if err = server.ListenAndServe(); err != nil && !errors.Is(err, ssh.ErrServerClosed) {
			log.Error("Could not start server", "error", err)
			done <- nil
		}
	}()

	<-done
	log.Info("Stopping SSH server")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer func() { cancel() }()
	if err := server.Shutdown(ctx); err != nil && !errors.Is(err, ssh.ErrServerClosed) {
		log.Error("Could not stop server", "error", err)
	}
}

func teaHandler(s ssh.Session) (tea.Model, []tea.ProgramOption) {
	// Initial model & setup when running the program
	pages := []string{"home", "about", "projects", "contact"}

	// Initializes projects list and itemized list for projects page
	projects := []string{"AEVSoftware", "Devfolify", "Schedulix", "SSHfolio", "Eduquest", "MemeAPI", "WebDevCourse", "Homelab", "ZachLTechWeb"}
	itemizedProjects := []list.Item{
		item{title: "Alset Solar Cybersedan Software", desc: "A full stack ecosystem powering the FAUHS AEV solar car"},
		item{title: "Devfolify", desc: "Giving you real world problems and a chance to solve it using code in your unique way"},
		item{title: "Schedulix", desc: "A program that helps university students develop their course schedules for their upcoming semester"},
		item{title: "SSHfolio", desc: "Minimally showing off all your unique talents through a publically SSHable TUI interface written in Go"},
		item{title: "Eduquest", desc: "Enjoy a world where education is guided by YOUR passions and preferences."},
		item{title: "MemeAPI", desc: "Just a funny API with ElysiaJS because Bun is cool and I had this idea a while ago for fun"},
		item{title: "WebDevCourse", desc: "Homemade Web Dev't Course for Friends that doesn't involve 12 hour \"full course\" videos (they're bad)"},
		item{title: "The Lopez Lab", desc: "A personal homelab infrastructure I put together for complex code experimentation and hosting almost anything"},
		item{title: "Personal Digital Branding", desc: "A collective web of sites and easter eggs I laid across the internet for anyone to explore"},
	}

	initialModel := model{
		pageIndex:   0,
		pages:       pages,
		projects:    projects,
		projectOpen: false,
		list:        list.New(itemizedProjects, list.NewDefaultDelegate(), 0, 0),
		keys:        DefaultKeyMap,
		help:        help.New(),
	}

	initialModel.list.InfiniteScrolling = true
	initialModel.list.DisableQuitKeybindings()
	initialModel.list.SetFilteringEnabled(false)
	initialModel.list.SetShowHelp(false)
	initialModel.list.SetShowTitle(false)
	initialModel.list.Title = "Hit ENTER for more details"

	return initialModel, []tea.ProgramOption{tea.WithAltScreen(), tea.WithMouseCellMotion()}
}

// ASCII text that should be displayed through every page
const ASCIIName string = `

 ________  ________  ________  ___  ___          ___       ________  ________  _______   ________     
|\_____  \|\   __  \|\   ____\|\  \|\  \        |\  \     |\   __  \|\   __  \|\  ___ \ |\_____  \    
 \|___/  /\ \  \|\  \ \  \___|\ \  \\\  \       \ \  \    \ \  \|\  \ \  \|\  \ \   __/| \|___/  /|   
     /  / /\ \   __  \ \  \    \ \   __  \       \ \  \    \ \  \\\  \ \   ____\ \  \_|/__   /  / /   
    /  /_/__\ \  \ \  \ \  \____\ \  \ \  \       \ \  \____\ \  \\\  \ \  \___|\ \  \_|\ \ /  /_/__  
   |\________\ \__\ \__\ \_______\ \__\ \__\       \ \_______\ \_______\ \__\    \ \_______\\________\
    \|_______|\|__|\|__|\|_______|\|__|\|__|        \|_______|\|_______|\|__|     \|_______|\|_______|

		"innovation isn't thinking outside the box. It's not even seeing one." - Daniel Cane
	`

// Bubbletea model structure
type model struct {
	pageIndex    int
	pages        []string
	projects     []string
	projectOpen  bool
	openProject  int
	projectView  string
	clickCounter int
	viewport     viewport.Model
	list         list.Model
	content      string
	keys         KeyMap
	help         help.Model
	ready        bool
}

// Check err
func check(e error, check string) {
	if e != nil {
		fmt.Printf("Error running program - In %v: %v", check, e)
	}
}

// For clickable area positioning
var termHeight int

// Lipgloss styling for view function & nav styling
var (
	navStyle           = lipgloss.NewStyle().Margin(1, 0).Padding(0, 2)
	listStyle          = lipgloss.NewStyle().Padding(1, 2)
	bubbleLettersStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#7aa2f7"))
	// For nav text
	activePageStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#4fd6be")).Bold(true).PaddingLeft(2).PaddingRight(4)
	inactivePageStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240")).PaddingLeft(4).PaddingRight(4)

	// Border styles
	borderTitleStyle = func() lipgloss.Style {
		b := lipgloss.HiddenBorder()
		return lipgloss.NewStyle().BorderStyle(b).Padding(0, 1)
	}()
	borderInfoStyle = func() lipgloss.Style {
		b := lipgloss.HiddenBorder()
		return borderTitleStyle.BorderStyle(b)
	}()
)

// Bubbletea key mapping (Struct + defaults)
type KeyMap struct {
	Navigate key.Binding
	Up       key.Binding
	Down     key.Binding
	Left     key.Binding
	Right    key.Binding
	LCycle   key.Binding
	RCycle   key.Binding
	Enter    key.Binding
	Back     key.Binding
	Help     key.Binding
	Quit     key.Binding
}

var DefaultKeyMap = KeyMap{
	Navigate: key.NewBinding(
		key.WithKeys("j", "k", "up", "down"),
		key.WithHelp("↑↓", "navigate"),
	),
	Up: key.NewBinding(
		key.WithKeys("k", "up"),
		key.WithHelp("↑/k", "up"),
	),
	Down: key.NewBinding(
		key.WithKeys("j", "down"),
		key.WithHelp("↓/j", "down"),
	),
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
		key.WithHelp("^tab", "prev section"),
	),
	RCycle: key.NewBinding(
		key.WithKeys("tab"),
		key.WithHelp("tab", "section"),
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

// Projects list setup
type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

func openProject(selectedProject int, projects []string, viewportWidth int) string {
	for indexedProject, project := range projects {
		if indexedProject == selectedProject {
			rawProjectPageTemplate, _ := glamour.NewTermRenderer(
				glamour.WithStylePath("assets/MDStyle.json"),
				// glamour.WithAutoStyle(), - For Light/Darkmode styling except I'd rather use my custom style
				glamour.WithWordWrap(viewportWidth-20),
			)

			projectPage, err := rawProjectPageTemplate.Render(getMarkdown("projects/" + project))
			check(err, "Project Glamour Render")

			return projectPage
		}
	}
	return fmt.Sprintf("Could not get %s project info...", projects[selectedProject])
}

// Function to read and return markdown file data for each page
func getMarkdown(filename string) string {
	fileData, err := os.ReadFile("./assets/markdown/" + filename + ".md")
	check(err, "Markdown File IO")

	return string(fileData)
}

// Function to get the proper content according to each page
func saturateContent(m model, viewportWidth int) string {
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
		content, err = rawMarkdownPageTemplate.Render(getMarkdown("homepage"))
		check(err, "Gleam Markdown Render")
	case 1: // About
		content, err = rawMarkdownPageTemplate.Render(getMarkdown("about"))
		check(err, "Gleam Markdown Render")
	case 3: // Contact
		content, err = rawMarkdownPageTemplate.Render(getMarkdown("contact"))
		check(err, "Gleam Markdown Render")
	}

	return content
}

// Bubbletea function to cycle each page (when tab is clicked, this function handles the update event)
func (m model) cyclePage(direction string) model {
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

// Gets the location and size of each navigation menu button
// (this is hard coded as of now since I have no idea how to programmatically find a components location & size in the terminal)
func (m model) calculateNavItemPosition(title string) (int, int) {
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
func calculateNavItemSize(title string) (int, int) {
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
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Page viewport header and footer render
func (m model) viewportHeader(pageTitle string) string {
	title := borderTitleStyle.Render(pageTitle)
	line := strings.Repeat("─", max(0, m.viewport.Width-lipgloss.Width(title)))
	return lipgloss.JoinHorizontal(lipgloss.Center, title, line)
}
func (m model) viewportFooter() string {
	info := borderInfoStyle.Render(fmt.Sprintf("%3.f%%", m.viewport.ScrollPercent()*100))
	line := strings.Repeat("─", max(0, m.viewport.Width-lipgloss.Width(info)))
	return lipgloss.JoinHorizontal(lipgloss.Center, line, info)
}

// Empty init for now since there's not much hard logic
func (m model) Init() tea.Cmd {
	return tea.SetWindowTitle("Welcome to my Portfolio TUI 😄")
}

// Bubbletea update/msg handling
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Commands to be returned for viewport updating
	var (
		viewportCMD     tea.Cmd
		projectsListCMD tea.Cmd
		cmds            []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.MouseMsg:
		switch tea.MouseAction(msg.Button) {
		case 1: // Mouse left click
			for i, title := range m.pages {
				x, y := m.calculateNavItemPosition(title)
				width, height := calculateNavItemSize(title)

				if msg.X >= x && msg.X <= x+width && msg.Y >= y && msg.Y <= y+height {
					m.pageIndex = i
					m.viewport.SetContent(saturateContent(m, m.viewport.Width))
					return m, nil
				} else if msg.Y >= termHeight-3 {
					m.help.ShowAll = !m.help.ShowAll
					return m, nil
				}
			}
			// This is a very lousy approach for making each item clickable but it's the only way I have time to do as of now...
			// This also causes the mouse support to break on pages past the first if pagination is necessary depending on terminal size
			if m.pageIndex == 2 && !m.projectOpen && msg.Y >= 16 && msg.Y < termHeight-3 {
				projectIndex := 0
				// BUG: for some reason after clicking down the list every once in a while it would enter the project MD even though it had only been clicked once then they all do that from that point on
				for i := 16; projectIndex <= len(m.projects)-1; i += 3 {
					if i <= msg.Y && msg.Y <= i+1 {
						if m.list.Index() == projectIndex {
							m.clickCounter++
						} else {
							m.clickCounter = 0
						}
						m.list.Select(projectIndex)
					} else {
						projectIndex++
					}
					if m.clickCounter >= 2 {
						m.clickCounter = 0
						m.projectOpen = true
						m.openProject = m.list.Index()
					}
				}
			}
		case 4: // Scroll wheel up
			if m.pageIndex == 2 && !m.projectOpen {
				if m.list.Index() == 0 {
					m.list.Select(len(m.projects))
				} else {
					m.list.Select(m.list.Index() - 1)
				}
			}
		case 5: // Scroll wheel down
			if m.pageIndex == 2 && !m.projectOpen {
				if m.list.Index() == len(m.projects)-1 {
					m.list.Select(0)
				} else {
					m.list.Select(m.list.Index() + 1)
				}
			}
		}
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, DefaultKeyMap.Quit):
			return m, tea.Quit
		case key.Matches(msg, DefaultKeyMap.Help):
			m.help.ShowAll = !m.help.ShowAll
		case key.Matches(msg, DefaultKeyMap.Navigate):
			break
		case key.Matches(msg, DefaultKeyMap.Up):
			break
		case key.Matches(msg, DefaultKeyMap.Down):
			break
		case key.Matches(msg, DefaultKeyMap.RCycle):
			cycled := m.cyclePage("right")
			cycled.viewport.SetContent(saturateContent(cycled, m.viewport.Width))
			return cycled, nil
		case key.Matches(msg, DefaultKeyMap.LCycle):
			cycled := m.cyclePage("left")
			m.viewport.SetContent(saturateContent(cycled, m.viewport.Width))
			return m.cyclePage("left"), nil
		case key.Matches(msg, DefaultKeyMap.Left):
			if m.pageIndex > 0 {
				m.pageIndex--
				m.viewport.SetContent(saturateContent(m, m.viewport.Width))
			}
			return m, nil
		case key.Matches(msg, DefaultKeyMap.Right):
			if m.pageIndex < len(m.pages)-1 {
				m.pageIndex++
				m.viewport.SetContent(saturateContent(m, m.viewport.Width))
			}
			return m, nil
		case key.Matches(msg, DefaultKeyMap.Enter):
			if m.pageIndex == 2 {
				m.projectOpen = true
				m.openProject = m.list.Index()
				m.viewport.GotoTop()
			}
		case key.Matches(msg, DefaultKeyMap.Back):
			if m.pageIndex == 2 {
				m.projectOpen = false
				m.list.Select(m.openProject)
				//m.viewport.GotoTop()
			}
		}
	case tea.WindowSizeMsg:
		// Set new terminal height for proper click areas
		termHeight = msg.Height
		// Setup for viewport sizing
		headerHeight := lipgloss.Height(m.viewportHeader(m.pages[m.pageIndex]))
		footerHeight := lipgloss.Height(m.viewportFooter())
		verticalMarginHeight := headerHeight + footerHeight
		// Project list size
		listMarginWidth, listMarginHeight := listStyle.GetFrameSize()
		m.list.SetSize(msg.Width-listMarginWidth, msg.Height-listMarginHeight-verticalMarginHeight-11)

		// Viewport creation & management
		if !m.ready {
			m.viewport = viewport.New(msg.Width, msg.Height-verticalMarginHeight-11)
			m.viewport.SetContent(saturateContent(m, m.viewport.Width))
			m.ready = true
		} else {
			m.viewport.Width = msg.Width
			m.viewport.Height = msg.Height - verticalMarginHeight - 11
		}
	}

	if m.pageIndex == 2 && m.projectOpen {
		m.viewport.SetContent(openProject(m.openProject, m.projects, m.viewport.Width))
	}
	// Handle keyboard and mouse events in the viewport
	// Gets viewport update command and map based on the message
	m.viewport, viewportCMD = m.viewport.Update(msg)
	// Update list depending on msg
	// Does the same as viewport but for projects list
	m.list, projectsListCMD = m.list.Update(msg)
	// Append all component commands to cmds
	cmds = append(cmds, viewportCMD, projectsListCMD)

	return m, tea.Batch(cmds...)
}

// Switch case with each page/TUI view
func (m model) View() string {

	// If viewport isn't ready it'll say welcome (this should only be able to happen during startup)
	if !m.ready {
		return "\n  Welcome..."
	}

	nav := `` // Empty to be saturated soon

	// Render/create nav depending on page location
	for i, title := range m.pages {
		if i == m.pageIndex {
			// Highlight the active page
			nav += activePageStyle.Render("• " + title + " ")
		} else {
			nav += inactivePageStyle.Render(title + " ")
		}
	}

	m.content = m.viewportHeader(m.pages[m.pageIndex]) + m.viewport.View() + m.viewportFooter()
	if m.pageIndex == 2 {
		if !m.projectOpen {
			m.projectView = listStyle.Render(m.list.View())
		} else if m.projectOpen {
			m.projectView = m.viewport.View()
		}
		m.content = m.viewportHeader(m.pages[m.pageIndex]) + m.projectView + m.viewportFooter()
	}

	header := lipgloss.PlaceHorizontal(m.viewport.Width, lipgloss.Center, bubbleLettersStyle.Render(ASCIIName))
	nav = lipgloss.PlaceHorizontal(m.viewport.Width, lipgloss.Center, navStyle.Render(nav))

	return header + nav + m.content + navStyle.Render(m.help.View(m.keys)) // NAME TITLE + NAVIGATION + MAIN PAGE + FOOTER/HELP
}
