package chat

import (
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbletea"
)

func (m model) Init() tea.Cmd {
	return tea.Batch(textarea.Blink, tea.ClearScreen)
}
