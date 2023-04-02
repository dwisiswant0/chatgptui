package config

import "github.com/charmbracelet/bubbles/textinput"

type configInput struct {
	defaultValue string
	label, name  string
	value        any
}

type model struct {
	err        error
	focusIndex int
	inputs     []textinput.Model
	configs    []configInput
}
