package components

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func RenderHelp() string {
	var output strings.Builder

	// Title
	title := HelpTitleStyle.Render(" HELP - KEYBOARD SHORTCUTS ")
	output.WriteString(title)
	output.WriteString("\n\n")

	// Navigation section
	output.WriteString(HelpSectionStyle.Render("NAVIGATION"))
	output.WriteString("\n")
	output.WriteString(HelpTextStyle.Render("↑ / k          Move up"))
	output.WriteString("\n")
	output.WriteString(HelpTextStyle.Render("↓ / j          Move down"))
	output.WriteString("\n")
	output.WriteString(HelpTextStyle.Render("Enter / Space  Select option"))
	output.WriteString("\n\n")

	// Features section
	output.WriteString(HelpSectionStyle.Render("FEATURES"))
	output.WriteString("\n")
	output.WriteString(HelpTextStyle.Render("LOAD    Open file picker to select CSV or Excel files"))
	output.WriteString("\n")
	output.WriteString(HelpTextStyle.Render("VIEW    Display loaded data in interactive table"))
	output.WriteString("\n")
	output.WriteString(HelpTextStyle.Render("CLEAN   Remove empty rows/columns, normalize names"))
	output.WriteString("\n")
	output.WriteString(HelpTextStyle.Render("PIVOT   Create summary tables and aggregations"))
	output.WriteString("\n")
	output.WriteString(HelpTextStyle.Render("CHECK   Run duplicate and missing value checks"))
	output.WriteString("\n")
	output.WriteString(HelpTextStyle.Render("SAVE    Export cleaned data as CSV or Excel"))
	output.WriteString("\n\n")

	// General section
	output.WriteString(HelpSectionStyle.Render("GENERAL"))
	output.WriteString("\n")
	output.WriteString(HelpTextStyle.Render("?          Show this help screen"))
	output.WriteString("\n")
	output.WriteString(HelpTextStyle.Render("b / Esc    Go back to menu"))
	output.WriteString("\n")
	output.WriteString(HelpTextStyle.Render("q / Ctrl+C Quit application"))
	output.WriteString("\n")

	return lipgloss.Place(
		100, 30,
		lipgloss.Center, lipgloss.Center,
		ContainerStyle.Render(output.String()),
	)
}
