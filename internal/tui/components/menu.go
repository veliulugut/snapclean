package components

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func RenderMenu(selectedItem int, options []string, statusText string) string {
	var output strings.Builder

	// Header
	header := HeaderStyle.Render(" SNAPCLEAN - DATA CLEANING TOOL ")
	output.WriteString(header)
	output.WriteString("\n\n")

	// Menu items
	for idx, option := range options {
		parts := strings.SplitN(option, "]", 2)
		tag := parts[0] + "]"
		desc := strings.TrimSpace(parts[1])

		var line string
		if idx == selectedItem {
			line = SelectedItemStyle.Render(
				"→ " + SelectedTagStyle.Render(tag) + "  " + desc,
			)
		} else {
			line = NormalItemStyle.Render(
				"  " + TagStyle.Render(tag) + "  " + desc,
			)
		}
		output.WriteString(line)
		output.WriteString("\n")
	}

	// Status bar
	if statusText != "" {
		output.WriteString("\n")
		statusBar := StatusBarStyle.Render(" STATUS ")
		output.WriteString(statusBar)
		output.WriteString("  ")
		output.WriteString(StatusTextStyle.Render(statusText))
	}

	// Help text
	output.WriteString("\n")
	helpText := HelpStyle.Render("↑/↓: Navigate  |  Enter: Select  |  ?: Help  |  q: Quit")
	output.WriteString(helpText)

	return lipgloss.Place(
		100, 30,
		lipgloss.Center, lipgloss.Center,
		ContainerStyle.Render(output.String()),
	)
}
