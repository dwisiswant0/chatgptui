package chat

import (
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	"github.com/dwisiswant0/chatgptui/common"
	"github.com/sashabaranov/go-openai"
)

type model struct {
	config   common.Config
	err      error
	messages []string
	textarea textarea.Model
	viewport viewport.Model

	openaiClient  *openai.Client
	openaiRequest openai.CompletionRequest
}
