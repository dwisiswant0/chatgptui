package config

import (
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbletea"
)

func (m model) Init() tea.Cmd {
	return textinput.Blink
}
