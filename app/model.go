package app

import (
	"sshfolio/ui"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/viewport"
)

// Bubbletea model structure
type Model struct {
	PageIndex    int
	Pages        []string
	Projects     []string
	ProjectOpen  bool
	OpenProject  int
	ProjectView  string
	ClickCounter int
	Viewport     viewport.Model
	List         list.Model
	Content      string
	Keys         ui.KeyMap
	Help         help.Model
	Ready        bool
}
