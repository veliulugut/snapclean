package cleaner

import (
	"testing"

	"github.com/veliulugut/snapclean/internal/models"
)

func TestValidateData(t *testing.T) {
	dt := models.NewDataTable([]string{"Name", "Age"})
	dt.AddRow([]string{"John", "30"})
	dt.AddRow([]string{"John", "30"}) // Duplicate
	dt.AddRow([]string{"", ""})       // Empty
	dt.AddRow([]string{"Jane", ""})   // Missing value

	result := ValidateData(dt)

	if !result.HasDuplicates {
		t.Error("Expected duplicates to be detected")
	}

	if result.DuplicateCount != 1 {
		t.Errorf("Expected 1 duplicate, got %d", result.DuplicateCount)
	}

	if result.EmptyRowCount != 1 {
		t.Errorf("Expected 1 empty row, got %d", result.EmptyRowCount)
	}

	if result.MissingValueCount < 1 {
		t.Errorf("Expected missing values, got %d", result.MissingValueCount)
	}
}

func TestValidateDataNil(t *testing.T) {
	result := ValidateData(nil)

	if result.TotalIssues != 0 {
		t.Errorf("Expected 0 issues for nil, got %d", result.TotalIssues)
	}
}

func TestValidateDataEmpty(t *testing.T) {
	dt := models.NewDataTable([]string{"Name", "Age"})

	result := ValidateData(dt)

	if result.TotalIssues != 0 {
		t.Errorf("Expected 0 issues for empty table, got %d", result.TotalIssues)
	}
}

func TestGetDuplicateRowIndices(t *testing.T) {
	dt := models.NewDataTable([]string{"Name", "Age"})
	dt.AddRow([]string{"John", "30"}) // Index 0
	dt.AddRow([]string{"Jane", "25"}) // Index 1
	dt.AddRow([]string{"John", "30"}) // Index 2 (duplicate of 0)

	indices := GetDuplicateRowIndices(dt)

	if len(indices) != 2 {
		t.Errorf("Expected 2 duplicate indices, got %d: %v", len(indices), indices)
	}

	// Should contain 0 and 2
	if !contains(indices, 0) || !contains(indices, 2) {
		t.Errorf("Expected indices 0 and 2, got %v", indices)
	}
}

func TestGetMissingValuesByColumn(t *testing.T) {
	dt := models.NewDataTable([]string{"Name", "Age", "City"})
	dt.AddRow([]string{"John", "30", ""})
	dt.AddRow([]string{"Jane", "", "LA"})
	dt.AddRow([]string{"Bob", "35", "NYC"})

	result := GetMissingValuesByColumn(dt)

	if result["Name"] != 0 {
		t.Errorf("Expected 0 missing in Name, got %d", result["Name"])
	}

	if result["Age"] != 1 {
		t.Errorf("Expected 1 missing in Age, got %d", result["Age"])
	}

	if result["City"] != 1 {
		t.Errorf("Expected 1 missing in City, got %d", result["City"])
	}
}

func TestGetMissingValuesByColumnNil(t *testing.T) {
	result := GetMissingValuesByColumn(nil)

	if len(result) != 0 {
		t.Errorf("Expected empty map for nil, got %v", result)
	}
}
