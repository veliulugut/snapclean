package components

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/veliulugut/snapclean/internal/models"
)

var (
	// Table styles
	TableHeaderStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(lipgloss.Color("#FFFFFF")).
				Background(lipgloss.Color("#5F5FFF")).
				Padding(0, 1).
				Align(lipgloss.Center)

	TableCellStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#BCBCBC")).
			Padding(0, 1)

	TableSelectedRowStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#FFFFFF")).
				Background(lipgloss.Color("#5F87FF")).
				Padding(0, 1)

	TableBorderStyle = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("#5F87FF")).
				Padding(1, 2)

	TableInfoStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#87FFAF")).
			Bold(true)

	TableHelpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#8787AF")).
			MarginTop(1)
)

// RenderTable renders a data table with pagination and horizontal scroll
func RenderTable(dt *models.DataTable, scrollOffset, columnOffset, pageSize, termWidth int) string {
	if dt == nil || dt.IsEmpty() {
		return ContainerStyle.Render("No data to display")
	}

	var output strings.Builder

	title := HeaderStyle.Render(fmt.Sprintf(" DATA VIEW - %s ", dt.FileName))
	output.WriteString(title + "\n\n")

	totalCols := dt.ColumnCount()
	visibleColCount := 11 // Show 11 columns at a time
	endCol := columnOffset + visibleColCount
	if endCol > totalCols {
		endCol = totalCols
	}

	info := TableInfoStyle.Render(fmt.Sprintf(
		"Rows: %d  |  Columns: %d-%d of %d  |  File: %s",
		dt.RowCount(),
		columnOffset+1,
		endCol,
		totalCols,
		dt.FileName,
	))
	output.WriteString(info + "\n\n")

	// Get visible columns
	visibleHeaders := getVisibleSlice(dt.Headers, columnOffset, visibleColCount)
	colWidths := calculateColumnWidths(dt, columnOffset, visibleColCount)

	// Headers
	headerRow := renderRow(visibleHeaders, colWidths, TableHeaderStyle)
	output.WriteString(headerRow + "\n")
	output.WriteString(strings.Repeat("─", sum(colWidths)+len(colWidths)*3) + "\n")

	// Data rows
	totalRows := dt.RowCount()
	endRow := scrollOffset + pageSize
	if endRow > totalRows {
		endRow = totalRows
	}
	for i := scrollOffset; i < endRow; i++ {
		row, _ := dt.GetRow(i)
		visibleCells := getVisibleSlice(row, columnOffset, visibleColCount)
		output.WriteString(renderRow(visibleCells, colWidths, TableCellStyle) + "\n")
	}

	output.WriteString("\n")
	output.WriteString(TableInfoStyle.Render(
		fmt.Sprintf("Showing rows %d-%d of %d", scrollOffset+1, endRow, totalRows),
	))

	output.WriteString("\n")
	output.WriteString(TableHelpStyle.Render(
		"↑/↓: Rows  |  ←/→: Columns  |  PgUp/PgDn: Page  |  c: Column Menu  |  b/Esc: Back  |  q: Quit",
	))

	return TableBorderStyle.Render(output.String())
}

// getVisibleSlice returns a slice of visible elements based on offset and count
func getVisibleSlice(items []string, offset, count int) []string {
	if offset >= len(items) {
		return []string{}
	}
	end := offset + count
	if end > len(items) {
		end = len(items)
	}
	return items[offset:end]
}

// calculateColumnWidths calculates widths for visible columns only
func calculateColumnWidths(dt *models.DataTable, colOffset, visibleCount int) []int {
	totalCols := dt.ColumnCount()
	endCol := colOffset + visibleCount
	if endCol > totalCols {
		endCol = totalCols
	}

	visibleRange := endCol - colOffset
	if visibleRange <= 0 {
		return []int{}
	}

	widths := make([]int, visibleRange)

	// Start with header widths
	for i := 0; i < visibleRange; i++ {
		colIdx := colOffset + i
		if colIdx < len(dt.Headers) {
			widths[i] = len(dt.Headers[colIdx])
		}
	}

	// Check first 100 rows
	limit := dt.RowCount()
	if limit > 100 {
		limit = 100
	}
	for r := 0; r < limit; r++ {
		row, _ := dt.GetRow(r)
		for i := 0; i < visibleRange; i++ {
			colIdx := colOffset + i
			if colIdx < len(row) && len(row[colIdx]) > widths[i] {
				widths[i] = len(row[colIdx])
			}
		}
	}

	// Cap width at 15 chars, min 5
	for i := range widths {
		if widths[i] > 15 {
			widths[i] = 15
		}
		if widths[i] < 5 {
			widths[i] = 5
		}
	}

	return widths
}

func renderRow(cells []string, widths []int, style lipgloss.Style) string {
	var parts []string
	for i := 0; i < len(widths); i++ {
		cell := ""
		if i < len(cells) {
			cell = cells[i]
		}
		if len(cell) > widths[i] {
			cell = cell[:max(1, widths[i]-3)] + "..."
		}
		cell = padRight(cell, widths[i])
		parts = append(parts, style.Render(cell))
	}
	return strings.Join(parts, " │ ")
}

func padRight(s string, length int) string {
	if len(s) >= length {
		return s
	}
	return s + strings.Repeat(" ", length-len(s))
}

func sum(nums []int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
