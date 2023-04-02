package config

import (
	"github.com/charmbracelet/bubbletea"

	"github.com/dwisiswant0/chatgptui/chat"
	"github.com/dwisiswant0/chatgptui/common"
	"github.com/dwisiswant0/chatgptui/style"
	"github.com/dwisiswant0/chatgptui/util"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit
		case "ctrl+r":
			for i := 0; i <= len(m.inputs)-1; i++ {
				m.inputs[i].Reset()
			}

			return m, nil
		case "tab":
			if m.focusIndex < len(m.inputs) {
				i := m.focusIndex
				v := m.configs[i].defaultValue
				if v != "" {
					m.inputs[i].SetValue(v)
				}
			}
		case "enter", "up", "down":
			s := msg.String()
			if s == "enter" {
				if m.focusIndex == len(m.inputs) {
					if err := m.validateInputs(); err != nil {
						m.err = err
						return m, nil
					} else {
						m.err = nil
					}

					if err := m.saveConfig(); err != nil {
						m.err = err
						return m, nil
					} else {
						cfg, err := Load(common.GetConfigPath())
						if err != nil {
							m.err = err
							return m, nil
						}

						util.RunProgram(chat.New(cfg))
						m.err = nil
					}

					return m, tea.Quit
				}

				if m.inputs[m.focusIndex].Value() == "" {
					return m, nil
				}

				if err := m.validateInput(m.focusIndex); err != nil {
					m.err = err
					return m, nil
				} else {
					m.err = nil
				}
			}

			if s == "up" {
				m.focusIndex--
			} else {
				m.focusIndex++
			}

			if m.focusIndex > len(m.inputs) {
				m.focusIndex = 0
			} else if m.focusIndex < 0 {
				m.focusIndex = len(m.inputs)
			}

			cmds := make([]tea.Cmd, len(m.inputs))
			for i := 0; i <= len(m.inputs)-1; i++ {
				if i == m.focusIndex {
					cmds[i] = m.inputs[i].Focus()
					m.inputs[i].PromptStyle = style.Focused
					m.inputs[i].TextStyle = style.Focused
					continue
				}

				m.inputs[i].Blur()
				m.inputs[i].PromptStyle = style.Clear
				m.inputs[i].TextStyle = style.Clear
			}

			return m, tea.Batch(cmds...)
		}
	}

	cmd := m.updateInputs(msg)

	return m, cmd
}
