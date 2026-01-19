package components

import (
	"fmt"
	"strings"

	"github.com/veliulugut/snapclean/internal/models"
)

// RenderColumnMenu renders the column selection and management interface
func RenderColumnMenu(dt *models.DataTable, selectedColumn int, message string) string {
	if dt == nil || dt.IsEmpty() {
		return ContainerStyle.Render("No data to display")
	}

	var output strings.Builder

	title := HeaderStyle.Render(" COLUMN MANAGEMENT ")
	output.WriteString(title + "\n\n")

	if message != "" {
		output.WriteString(SuccessMessageStyle.Render(message) + "\n\n")
	}

	output.WriteString(TableInfoStyle.Render(
		fmt.Sprintf("Total Columns: %d  |  Select a column to swap with another", dt.ColumnCount()),
	))
	output.WriteString("\n\n")

	// Display columns with selection indicator
	for i, header := range dt.Headers {
		line := fmt.Sprintf("  [%2d] %s", i+1, header)
		if i == selectedColumn {
			output.WriteString(SelectedStyle.Render("> "+line) + "\n")
		} else {
			output.WriteString(TableCellStyle.Render("  "+line) + "\n")
		}
	}

	output.WriteString("\n")
	output.WriteString(TableHelpStyle.Render(
		"↑/↓: Select Column  |  Enter: Choose Column to Swap  |  b/Esc: Back",
	))

	return TableBorderStyle.Render(output.String())
}
