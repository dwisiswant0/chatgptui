package chat

import (
	"strings"

	"github.com/dwisiswant0/chatgptui/common"
	"github.com/dwisiswant0/chatgptui/style"
)

func (m model) View() string {
	var b strings.Builder

	b.WriteString(m.viewport.View())
	b.WriteString("\n\n")
	b.WriteString(m.textarea.View())
	b.WriteString("\n\n")

	if m.err != nil {
		b.WriteString(style.Error.Render(m.err.Error()) + "\n\n")
	}

	b.WriteString(style.Help.Render(common.HelpTextProTip))
	b.WriteString("\n\n")
	b.WriteString(style.Help.Render(common.HelpText))

	return b.String()
}
