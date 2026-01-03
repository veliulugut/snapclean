// filepath: internal/tui/view.go
package tui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("205")).
			MarginBottom(1)

	selectedStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("170")).
			Bold(true)
)

func (m AppModel) View() string {
	var output strings.Builder

	output.WriteString(titleStyle.Render("🧼 SnapClean - Data Cleaning Tool"))
	output.WriteString("\n\n")

	for idx, option := range m.options {
		cursor := "  "
		if idx == m.selectedItem {
			cursor = "→ "
			output.WriteString(selectedStyle.Render(cursor + option))
		} else {
			output.WriteString(cursor + option)
		}
		output.WriteString("\n")
	}

	if m.statusText != "" {
		output.WriteString("\n" + m.statusText + "\n")
	}

	output.WriteString("\n\nPress ↑/↓ to navigate • Enter to select • q to quit\n")

	return output.String()
}
