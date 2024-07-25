package main

// Imports (useless comment lol)
import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
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
	viewport  viewport.Model
	content   string
	keys      KeyMap
	help      help.Model
	ready     bool
}

// Bubbletea key mapping for help component
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

// Bubbletea key mapping default behavior
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

// Lipgloss styling for view function & nav styling
var (
	navStyle           = lipgloss.NewStyle().Margin(1, 0).Padding(0, 2)
	bubbleLettersStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#7aa2f7"))
	activePageStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("#4fd6be")).Bold(true).PaddingLeft(2).PaddingRight(4)
	inactivePageStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("240")).PaddingLeft(4).PaddingRight(4)

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

// Function to read and return markdown file data for each page
func getMarkdown(filename string) string {
	fileData, err := os.ReadFile("./assets/markdown/" + filename + ".md")
	check(err, "Markdown File IO")

	return string(fileData)
}

// Function to get the proper content according to each page
func saturateContent(m model) string {
	// Checks which page the user is on and renders it accordingly
	var content string
	var err error

	switch m.pageIndex {
	case 0: // Home
		content, err = glamour.Render(getMarkdown("homepage"), "dark")
		check(err, "Gleam Markdown Render")
	case 1: // About
		content, err = glamour.Render(getMarkdown("about"), "dark")
		check(err, "Gleam Markdown Render")
	case 2: // Projects
		content += "This is the projects page which is under construction... (bubble list later)"
	case 3: // Contact
		content, err = glamour.Render(getMarkdown("contact"), "dark")
		check(err, "Gleam Markdown Render")
	case 4: // Resume
		content, err = glamour.Render(getMarkdown("resume"), "dark")
		check(err, "Gleam Markdown Render")
	}

	return content
}

// Check err
func check(e error, check string) {
	if e != nil {
		fmt.Printf("Error running program - In %v: %v", check, e)
		os.Exit(1)
	}
}

// Max function for viewport line length
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
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

// Empty init for now since there's not much hard logic
func (m model) Init() tea.Cmd {
	return nil
}

// Bubbletea update/msg handling
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Commands to be returned for viewport updating
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
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
			cycled.viewport.SetContent(saturateContent(cycled))
			return cycled, nil
		case key.Matches(msg, DefaultKeyMap.LCycle):
			cycled := m.cyclePage("left")
			m.viewport.SetContent(saturateContent(cycled))
			return m.cyclePage("left"), nil
		case key.Matches(msg, DefaultKeyMap.Left):
			if m.pageIndex > 0 {
				m.pageIndex--
				m.viewport.SetContent(saturateContent(m))
			}
			return m, nil
		case key.Matches(msg, DefaultKeyMap.Right):
			if m.pageIndex < len(m.pages)-1 {
				m.pageIndex++
				m.viewport.SetContent(saturateContent(m))
			}
			return m, nil
		}
	case tea.WindowSizeMsg:
		headerHeight := lipgloss.Height(m.viewportHeader(m.pages[m.pageIndex]))
		footerHeight := lipgloss.Height(m.viewportFooter())
		verticalMarginHeight := headerHeight + footerHeight

		if !m.ready {
			m.viewport = viewport.New(msg.Width, msg.Height-verticalMarginHeight-10)
			m.viewport.SetContent(saturateContent(m))
			m.ready = true
		} else {
			m.viewport.Width = msg.Width
			m.viewport.Height = msg.Height - verticalMarginHeight - 10
		}
	}

	// Handle keyboard and mouse events in the viewport
	m.viewport, cmd = m.viewport.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

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
	header := lipgloss.PlaceHorizontal(m.viewport.Width, lipgloss.Center, bubbleLettersStyle.Render(ASCIIName))
	nav = lipgloss.PlaceHorizontal(m.viewport.Width, lipgloss.Center, navStyle.Render(nav))

	return header + nav + m.content + navStyle.Render(m.help.View(m.keys)) // NAME TITLE + NAVIGATION + MAIN PAGE + FOOTER/HELP
}

// Starts the Bubbletea TUI & sets up initial state
func main() {

	// Initial model & setup when running the program
	pages := []string{"home", "about", "projects", "contact", "resume"}

	initialModel := model{
		pageIndex: 0,
		pages:     pages,
		keys:      DefaultKeyMap,
		help:      help.New(),
	}

	p := tea.NewProgram(initialModel, tea.WithAltScreen(), tea.WithMouseCellMotion())

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
