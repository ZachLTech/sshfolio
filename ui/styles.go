package ui

import (
	"github.com/charmbracelet/lipgloss"
)

// Ig this can be considered a "style" haha
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

// For clickable area positioning
var TermHeight int

// Lipgloss styling for view function & nav styling
var (
	NavStyle           = lipgloss.NewStyle().Margin(1, 0).Padding(0, 2)
	ListStyle          = lipgloss.NewStyle().Padding(1, 2)
	BubbleLettersStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#7aa2f7"))
	// For nav text
	ActivePageStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#4fd6be")).Bold(true).PaddingLeft(2).PaddingRight(4)
	InactivePageStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240")).PaddingLeft(4).PaddingRight(4)

	// Border styles
	BorderTitleStyle = func() lipgloss.Style {
		b := lipgloss.HiddenBorder()
		return lipgloss.NewStyle().BorderStyle(b).Padding(0, 1)
	}()
	BorderInfoStyle = func() lipgloss.Style {
		b := lipgloss.HiddenBorder()
		return BorderTitleStyle.BorderStyle(b)
	}()
)
