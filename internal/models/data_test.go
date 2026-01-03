package models

import "testing"

func TestNewDataTable(t *testing.T) {
	headers := []string{"Name", "Age", "City"}
	dt := NewDataTable(headers)

	if dt.ColumnCount() != 3 {
		t.Errorf("Expected 3 columns, got %d", dt.ColumnCount())
	}

	if dt.RowCount() != 0 {
		t.Errorf("Expected 0 rows, got %d", dt.RowCount())
	}
}

func TestAddRow(t *testing.T) {
	dt := NewDataTable([]string{"Name", "Age"})

	// Valid row
	err := dt.AddRow([]string{"John", "30"})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if dt.RowCount() != 1 {
		t.Errorf("Expected 1 row, got %d", dt.RowCount())
	}

	// Invalid row (length mismatch)
	err = dt.AddRow([]string{"Jane"})
	if err == nil {
		t.Error("Expected error for mismatched row length")
	}
}

func TestGetRow(t *testing.T) {
	dt := NewDataTable([]string{"Name", "Age"})
	dt.AddRow([]string{"John", "30"})

	row, err := dt.GetRow(0)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if row[0] != "John" || row[1] != "30" {
		t.Errorf("Expected ['John', '30'], got %v", row)
	}

	// Out of bounds
	_, err = dt.GetRow(5)
	if err == nil {
		t.Error("Expected error for out of bounds index")
	}
}

func TestGetColumn(t *testing.T) {
	dt := NewDataTable([]string{"Name", "Age"})
	dt.AddRow([]string{"John", "30"})
	dt.AddRow([]string{"Jane", "25"})

	column, err := dt.GetColumn(0)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if len(column) != 2 || column[0] != "John" || column[1] != "Jane" {
		t.Errorf("Expected ['John', 'Jane'], got %v", column)
	}
}

func TestIsEmpty(t *testing.T) {
	dt := NewDataTable([]string{"Name"})

	if !dt.IsEmpty() {
		t.Error("Expected table to be empty")
	}

	dt.AddRow([]string{"John"})

	if dt.IsEmpty() {
		t.Error("Expected table to not be empty")
	}
}

func TestClone(t *testing.T) {
	dt := NewDataTable([]string{"Name", "Age"})
	dt.AddRow([]string{"John", "30"})
	dt.FilePath = "/test/path.csv"

	clone := dt.Clone()

	// Check deep copy
	if clone.RowCount() != dt.RowCount() {
		t.Error("Clone has different row count")
	}

	if clone.FilePath != dt.FilePath {
		t.Error("Clone has different file path")
	}

	// Modify original, check clone is unaffected
	dt.AddRow([]string{"Jane", "25"})

	if clone.RowCount() == dt.RowCount() {
		t.Error("Clone was affected by original modification")
	}
}
