package components

import (
	"fmt"
	"strings"

	"github.com/veliulugut/snapclean/internal/models"
)

type CleaningViewModel struct {
	Options  models.CleanOptions
	Selected int
	Message  string
}

// RenderCleaning renders the cleaning options screen
func RenderCleaning(vm CleaningViewModel) string {
	items := []struct {
		label string
		on    bool
	}{
		{"Trim whitespace (headers + cells)", vm.Options.TrimWhitespace},
		{"Normalize headers (lowercase, _)", vm.Options.NormalizeHeaders},
		{"Remove empty rows", vm.Options.RemoveEmptyRows},
		{"Remove empty columns", vm.Options.RemoveEmptyColumns},
		{"Remove dublicate rows", vm.Options.RemoveDuplicates},
	}

	var (
		b strings.Builder
	)

	b.WriteString(HeaderStyle.Render(" CLEAN DATA "))
	b.WriteString("\n\n")

	for i, it := range items {
		box := "[ ]"
		if it.on {
			box = "[x]"
		}
		line := fmt.Sprintf("%s %s", box, it.label)
		if i == vm.Selected {
			b.WriteString(TableSelectedRowStyle.Render(line))
		} else {
			b.WriteString(TableCellStyle.Render(line))
		}
		b.WriteString("\n")
	}

	if vm.Message != "" {
		b.WriteString("\n")
		b.WriteString(TableInfoStyle.Render(vm.Message))
	}

	b.WriteString("\n")
	b.WriteString(TableHelpStyle.Render("Space: Toggle | Enter: Apply | b/Esc: Back | q:Quit"))

	return TableBorderStyle.Render(b.String())
}
