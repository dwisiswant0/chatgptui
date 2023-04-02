package errors

import "errors"

var (
	InvalidAPIKey      = errors.New("Invalid OpenAI API key!")
	MaxLengthRange     = errors.New("Max. length range between 1-4000!")
	GreaterFloatNumber = errors.New("Floating number cannot be greater than 1!")
)

const (
	InvalidModel       = `"%s" is not a valid model!`
	InvalidFloatNumber = `"%s" is not a valid floating number!`
)
