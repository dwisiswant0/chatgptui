package config

import (
	"fmt"
	"strings"

	"github.com/dwisiswant0/chatgptui/common"
	"github.com/dwisiswant0/chatgptui/style"
)

func (m model) View() string {
	var b strings.Builder

	for i := range m.inputs {
		b.WriteString(m.inputs[i].View())
		if i < len(m.inputs)-1 {
			b.WriteRune('\n')
		}
	}

	button := &style.BlurredBtn
	if m.focusIndex == len(m.inputs) {
		button = &style.FocusedBtn
	}
	fmt.Fprintf(&b, "\n\n%s\n\n", *button)

	if m.err != nil {
		b.WriteString(style.Error.Render(m.err.Error()) + "\n\n")
	}

	b.WriteString(style.Help.Render(fmt.Sprintf(
		" %s\n %s\n %s\n", common.HelpText,
		common.HelpTextTab, common.HelpTextReset,
	)))

	return b.String()
}
