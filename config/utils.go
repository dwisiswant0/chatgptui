package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"encoding/json"

	"github.com/charmbracelet/bubbletea"

	"github.com/dwisiswant0/chatgptui/common"
	"github.com/dwisiswant0/chatgptui/errors"
)

func (m model) getPlaceholder(i int) string {
	if i > len(m.configs) {
		return ""
	}

	str := m.configs[i].label
	if m.configs[i].defaultValue != "" {
		str = fmt.Sprintf(`%s (default "%s")`, str, m.configs[i].defaultValue)
	}

	return str
}

func (m model) validateInput(i int) error {
	if i > len(m.inputs) {
		return nil
	}

	val := m.inputs[i].Value()

	switch m.configs[i].name {
	case "openai_api_key":
		if !strings.HasPrefix(val, "sk-") {
			return errors.InvalidAPIKey
		}
		m.configs[i].value = val
	case "model":
		for _, model := range common.OpenaiModels {
			if val == model {
				m.configs[i].value = val
				return nil
			}
		}
		return fmt.Errorf(errors.InvalidModel, val)
	case "max_length":
		length, err := strconv.Atoi(val)
		if err != nil {
			return errors.MaxLengthRange
		}

		if length < 1 || length > 4000 {
			return errors.MaxLengthRange
		}

		m.configs[i].value = length
	case "temperature", "top_p":
		valFloat32, err := strconv.ParseFloat(val, 32)
		if err != nil {
			return fmt.Errorf(errors.InvalidFloatNumber, val)
		}

		if valFloat32 > 1 {
			return errors.GreaterFloatNumber
		}

		m.configs[i].value = valFloat32
	}

	return nil
}

func (m model) validateInputs() error {
	for i := range m.inputs {
		err := m.validateInput(i)
		if err != nil {
			return err
		}
	}

	return nil
}

func (m model) saveConfig() error {
	file, err := os.Create(common.GetConfigPath())
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)

	cfgMap := make(map[string]any)

	for _, config := range m.configs {
		cfgMap[config.name] = config.value
	}

	if err := encoder.Encode(cfgMap); err != nil {
		return err
	}

	return nil
}

func (m *model) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.inputs))

	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}
