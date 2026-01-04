package cleaner

import (
	"strings"

	"github.com/veliulugut/snapclean/internal/models"
)

// RemoveEmptyRows removes rows where all cells are empty
func RemoveEmptyRows(dt *models.DataTable) *models.DataTable {
	result := dt.Clone()
	var filteredRows [][]string

	for _, row := range result.Rows {
		if !isRowEmpty(row) {
			filteredRows = append(filteredRows, row)
		}
	}

	result.Rows = filteredRows
	return result
}

// RemoveEmptyColumns removes columns where all cells are empty
func RemoveEmptyColumns(dt *models.DataTable) *models.DataTable {
	if dt.IsEmpty() {
		return dt.Clone()
	}

	result := dt.Clone()

	// Find which columns to keep
	var keepIndices []int
	var newHeaders []string

	for colIdx := 0; colIdx < dt.ColumnCount(); colIdx++ {
		column, _ := dt.GetColumn(colIdx)

		// Check if column has any non-empty cell
		hasContent := false
		for _, cell := range column {
			if strings.TrimSpace(cell) != "" {
				hasContent = true
				break
			}
		}

		if hasContent {
			keepIndices = append(keepIndices, colIdx)
			newHeaders = append(newHeaders, dt.Headers[colIdx])
		}
	}

	// Filter rows to keep only selected columns
	var newRows [][]string
	for _, row := range result.Rows {
		var newRow []string
		for _, idx := range keepIndices {
			if idx < len(row) {
				newRow = append(newRow, row[idx])
			} else {
				newRow = append(newRow, "")
			}
		}
		newRows = append(newRows, newRow)
	}

	result.Headers = newHeaders
	result.Rows = newRows
	return result
}

// NormalizeHeaders converts headers to lowercase and replaces spaces with underscores
func NormalizeHeaders(dt *models.DataTable) *models.DataTable {
	result := dt.Clone()

	for i, header := range result.Headers {
		// Trim whitespace first
		normalized := strings.TrimSpace(header)
		// Convert to lowercase
		normalized = strings.ToLower(normalized)
		// Replace spaces with underscores
		normalized = strings.ReplaceAll(normalized, " ", "_")
		// Remove special characters (keep only alphanumeric and underscore)
		normalized = removeSpecialChars(normalized)
		// Remove leading/trailing underscores
		normalized = strings.Trim(normalized, "_")

		result.Headers[i] = normalized
	}

	return result
}

// RemoveDuplicates removes duplicate rows (keeps first occurrence)
func RemoveDuplicates(dt *models.DataTable) *models.DataTable {
	if dt.IsEmpty() {
		return dt.Clone()
	}

	result := dt.Clone()
	seen := make(map[string]bool)
	var uniqueRows [][]string

	for _, row := range result.Rows {
		rowKey := strings.Join(row, "|||")
		if !seen[rowKey] {
			seen[rowKey] = true
			uniqueRows = append(uniqueRows, row)
		}
	}

	result.Rows = uniqueRows
	return result
}

// TrimWhitespace removes leading/trailing whitespace from all cells
func TrimWhitespace(dt *models.DataTable) *models.DataTable {
	result := dt.Clone()

	// Trim headers
	for i := range result.Headers {
		result.Headers[i] = strings.TrimSpace(result.Headers[i])
	}

	// Trim rows
	for i := range result.Rows {
		for j := range result.Rows[i] {
			result.Rows[i][j] = strings.TrimSpace(result.Rows[i][j])
		}
	}

	return result
}

// ApplyCleaningOptions applies multiple cleaning operations in order
func ApplyCleaningOptions(dt *models.DataTable, opts models.CleanOptions) *models.DataTable {
	if dt == nil || dt.IsEmpty() {
		return dt
	}

	result := dt

	// Order matters: trim first, then normalize, then remove empty, then deduplicate
	if opts.TrimWhitespace {
		result = TrimWhitespace(result)
	}

	if opts.NormalizeHeaders {
		result = NormalizeHeaders(result)
	}

	if opts.RemoveEmptyRows {
		result = RemoveEmptyRows(result)
	}

	if opts.RemoveEmptyColumns {
		result = RemoveEmptyColumns(result)
	}

	if opts.RemoveDuplicates {
		result = RemoveDuplicates(result)
	}

	return result
}

// Helper functions

// isRowEmpty checks if all cells in a row are empty or whitespace
func isRowEmpty(row []string) bool {
	for _, cell := range row {
		if strings.TrimSpace(cell) != "" {
			return false
		}
	}
	return true
}

// removeSpecialChars removes special characters, keeps only alphanumeric and underscore
func removeSpecialChars(s string) string {
	var result strings.Builder
	for _, r := range s {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == '_' {
			result.WriteRune(r)
		}
	}
	return result.String()
}
