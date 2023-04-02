package util

import (
	"log"

	"github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func RunProgram(model tea.Model) {
	p := tea.NewProgram(model)
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

func SetTermColor(s string) lipgloss.Style {
	return lipgloss.NewStyle().Foreground(lipgloss.Color(s))
}
