package components

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func RenderSplash(splashTick int) string {
	var output strings.Builder

	splash := `
  ███████╗███╗   ██╗ █████╗ ██████╗ ██████╗██╗     ███████╗ █████╗ ███╗   ██╗
  ██╔════╝████╗  ██║██╔══██╗██╔══██╗██╔════╝██║     ██╔════╝██╔══██╗████╗  ██║
  ███████╗██╔██╗ ██║███████║██████╔╝██║     ██║     █████╗  ███████║██╔██╗ ██║
  ╚════██║██║╚██╗██║██╔══██║██╔═══╝ ██║     ██║     ██╔══╝  ██╔══██║██║╚██╗██║
  ███████║██║ ╚████║██║  ██║██║     ███████╗███████╗███████╗██║  ██║██║ ╚████║
  ╚══════╝╚═╝  ╚═══╝╚═╝  ╚═╝╚═╝     ╚══════╝╚══════╝╚══════╝╚═╝  ╚═╝╚═╝  ╚═══╝
    `

	output.WriteString(lipgloss.NewStyle().
		Foreground(lipgloss.Color("#5F87FF")).
		Bold(true).
		Render(splash))

	output.WriteString("\n\n")
	output.WriteString(lipgloss.NewStyle().
		Foreground(lipgloss.Color("#87FFAF")).
		Italic(true).
		Render("Interactive Data Cleaning Tool"))

	output.WriteString("\n\n")

	frames := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
	frame := splashTick % len(frames)

	output.WriteString(lipgloss.NewStyle().
		Foreground(lipgloss.Color("#8787AF")).
		Render(frames[frame] + " Loading..."))

	output.WriteString("\n\n")
	output.WriteString(lipgloss.NewStyle().
		Foreground(lipgloss.Color("#6272A4")).
		Render("Press any key to skip"))

	return lipgloss.Place(
		100, 30,
		lipgloss.Center, lipgloss.Center,
		SplashContainerStyle.Render(output.String()),
	)
}
