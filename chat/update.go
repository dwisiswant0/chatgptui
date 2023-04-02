package chat

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/dwisiswant0/chatgptui/style"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		tiCmd tea.Cmd
		vpCmd tea.Cmd
	)

	m.textarea, tiCmd = m.textarea.Update(msg)
	m.viewport, vpCmd = m.viewport.Update(msg)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit
		case "enter":
			val := m.textarea.Value()
			if val == "" {
				return m, nil
			}

			switch val {
			case "/c", "/clear":
				m.messages = []string{}
			default:
				m.messages = append(m.messages, fmt.Sprintf("%s %s", style.Sender.Render("👤:"), val))

				res, err := m.sendChat(val)
				if err != nil {
					m.err = err
					return m, nil
				}
				m.messages = append(m.messages, fmt.Sprintf(
					"%s %s",
					style.Response.Render("🤖:"),
					lipgloss.NewStyle().Width(78-5).Render(res)),
				)
			}

			m.viewport.SetContent(strings.Join(m.messages, "\n"))
			m.textarea.Reset()
			m.viewport.GotoBottom()
		}
	}

	return m, tea.Batch(tiCmd, vpCmd)
}
