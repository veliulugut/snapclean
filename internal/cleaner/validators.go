package cleaner

import (
	"strings"

	"github.com/veliulugut/snapclean/internal/models"
)

// ValidationResult contains data quality metrics
type ValidationResult struct {
	HasDuplicates     bool
	DuplicateCount    int
	MissingValueCount int
	EmptyRowCount     int
	EmptyColumnCount  int
	TotalIssues       int
}

// ValidateData performs comprehensive data validation
func ValidateData(dt *models.DataTable) ValidationResult {
	result := ValidationResult{}

	if dt == nil || dt.IsEmpty() {
		return result
	}

	// Count duplicates
	seen := make(map[string]bool)
	for _, row := range dt.Rows {
		key := strings.Join(row, "|||")
		if seen[key] {
			result.DuplicateCount++
		}
		seen[key] = true
	}
	result.HasDuplicates = result.DuplicateCount > 0

	// Count empty rows
	for _, row := range dt.Rows {
		if isRowEmpty(row) {
			result.EmptyRowCount++
		}
	}

	// Count empty columns
	for colIdx := 0; colIdx < dt.ColumnCount(); colIdx++ {
		column, _ := dt.GetColumn(colIdx)
		isEmpty := true
		for _, cell := range column {
			if strings.TrimSpace(cell) != "" {
				isEmpty = false
				break
			}
		}
		if isEmpty {
			result.EmptyColumnCount++
		}
	}

	// Count missing values
	for _, row := range dt.Rows {
		for _, cell := range row {
			if strings.TrimSpace(cell) == "" {
				result.MissingValueCount++
			}
		}
	}

	// Calculate total issues
	result.TotalIssues = result.DuplicateCount + result.EmptyRowCount +
		result.EmptyColumnCount + result.MissingValueCount

	return result
}

// GetDuplicateRowIndices returns indices of duplicate rows (including first occurrence)
func GetDuplicateRowIndices(dt *models.DataTable) []int {
	if dt == nil || dt.IsEmpty() {
		return []int{}
	}

	var duplicates []int
	seen := make(map[string]int)

	for i, row := range dt.Rows {
		key := strings.Join(row, "|||")
		if idx, exists := seen[key]; exists {
			// Add both original and duplicate
			if !contains(duplicates, idx) {
				duplicates = append(duplicates, idx)
			}
			duplicates = append(duplicates, i)
		}
		seen[key] = i
	}

	return duplicates
}

// GetMissingValuesByColumn returns missing value count per column
func GetMissingValuesByColumn(dt *models.DataTable) map[string]int {
	result := make(map[string]int)

	if dt == nil || dt.IsEmpty() {
		return result
	}

	for colIdx := 0; colIdx < dt.ColumnCount(); colIdx++ {
		header := dt.Headers[colIdx]
		count := 0

		for _, row := range dt.Rows {
			if colIdx >= len(row) || strings.TrimSpace(row[colIdx]) == "" {
				count++
			}
		}
		result[header] = count
	}

	return result
}

// Helper function
func contains(slice []int, val int) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}
