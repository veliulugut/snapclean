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
func RenderTable(dt *models.DataTable, scrollOffset, columnOffset, pageSize, visibleColumns int) string {
	if dt == nil || dt.IsEmpty() {
		return lipgloss.Place(
			100, 30,
			lipgloss.Center, lipgloss.Center,
			ContainerStyle.Render("No data to display"),
		)
	}

	var output strings.Builder

	// Title
	title := HeaderStyle.Render(fmt.Sprintf(" DATA VIEW - %s ", dt.FileName))
	output.WriteString(title)
	output.WriteString("\n\n")

	// Info bar
	totalCols := dt.ColumnCount()
	endCol := columnOffset + visibleColumns
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
	output.WriteString(info)
	output.WriteString("\n\n")

	// Get visible columns
	visibleHeaders := getVisibleSlice(dt.Headers, columnOffset, visibleColumns)

	// Calculate column widths for visible columns only
	colWidths := calculateColumnWidths(dt, 20, columnOffset, visibleColumns)

	// Render headers
	headerRow := renderRow(visibleHeaders, colWidths, TableHeaderStyle)
	output.WriteString(headerRow)
	output.WriteString("\n")
	output.WriteString(strings.Repeat("─", sum(colWidths)+len(colWidths)*3))
	output.WriteString("\n")

	// Render visible rows
	totalRows := dt.RowCount()
	endRow := scrollOffset + pageSize
	if endRow > totalRows {
		endRow = totalRows
	}

	for i := scrollOffset; i < endRow; i++ {
		row, _ := dt.GetRow(i)
		visibleCells := getVisibleSlice(row, columnOffset, visibleColumns)
		rowStr := renderRow(visibleCells, colWidths, TableCellStyle)
		output.WriteString(rowStr)
		output.WriteString("\n")
	}

	// Pagination info
	output.WriteString("\n")
	paginationInfo := fmt.Sprintf("Showing rows %d-%d of %d", scrollOffset+1, endRow, totalRows)
	output.WriteString(TableInfoStyle.Render(paginationInfo))

	// Help text
	output.WriteString("\n")
	helpText := TableHelpStyle.Render("↑/↓: Scroll Rows  |  ←/→: Scroll Columns  |  PgUp/PgDn: Page  |  b/Esc: Back  |  q: Quit")
	output.WriteString(helpText)

	return lipgloss.Place(
		130, 35,
		lipgloss.Center, lipgloss.Center,
		TableBorderStyle.Render(output.String()),
	)
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

// calculateColumnWidths calculates optimal width for visible columns only
func calculateColumnWidths(dt *models.DataTable, maxWidth, columnOffset, visibleColumns int) []int {
	totalCols := dt.ColumnCount()
	endCol := columnOffset + visibleColumns
	if endCol > totalCols {
		endCol = totalCols
	}

	visibleCount := endCol - columnOffset
	widths := make([]int, visibleCount)

	// Start with header widths
	for i := 0; i < visibleCount; i++ {
		colIdx := columnOffset + i
		if colIdx < len(dt.Headers) {
			widths[i] = len(dt.Headers[colIdx])
		}
	}

	// Check first 100 rows for max width
	checkRows := 100
	if dt.RowCount() < checkRows {
		checkRows = dt.RowCount()
	}

	for rowIdx := 0; rowIdx < checkRows; rowIdx++ {
		row, _ := dt.GetRow(rowIdx)
		for i := 0; i < visibleCount; i++ {
			colIdx := columnOffset + i
			if colIdx < len(row) && len(row[colIdx]) > widths[i] {
				widths[i] = len(row[colIdx])
			}
		}
	}

	// Apply max width limit
	for i := range widths {
		if widths[i] > maxWidth {
			widths[i] = maxWidth
		}
		if widths[i] < 5 {
			widths[i] = 5
		}
	}

	return widths
}

// renderRow renders a single row with given widths and style
func renderRow(cells []string, widths []int, style lipgloss.Style) string {
	var parts []string

	for i, cell := range cells {
		if i >= len(widths) {
			break
		}

		// Truncate if too long
		if len(cell) > widths[i] {
			cell = cell[:widths[i]-3] + "..."
		}

		// Pad to width
		cell = padRight(cell, widths[i])

		parts = append(parts, style.Render(cell))
	}

	return strings.Join(parts, " │ ")
}

// padRight pads a string to the right with spaces
func padRight(s string, length int) string {
	if len(s) >= length {
		return s
	}
	return s + strings.Repeat(" ", length-len(s))
}

// sum calculates the sum of integers in a slice
func sum(nums []int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}
