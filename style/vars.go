package style

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/dwisiswant0/chatgptui/util"
)

var (
	Focused = util.SetTermColor("205")

	Cursor      = Focused.Copy()
	Error       = util.SetTermColor("11")
	Help        = util.SetTermColor("240")
	Clear       = lipgloss.NewStyle()
	Placeholder = util.SetTermColor("60")
	Response    = util.SetTermColor("#b13434")
	Sender      = util.SetTermColor("#1c74d4")
	Spinner     = Focused.Copy()
	Viewport    = lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("62")).
			PaddingRight(2)

	FocusedBtn = Focused.Copy().Render("[ Save ]")
	BlurredBtn = fmt.Sprintf("[ %s ]", Help.Render("Save"))
)
