package file

import (
	"os"
	"testing"
)

func TestLoadCSV(t *testing.T) {
	// Create temporary CSV file
	content := `Name,Age,City
John,30,NYC
Jane,25,LA
Bob,35,Chicago`

	tmpfile, err := os.CreateTemp("", "test*.csv")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write([]byte(content)); err != nil {
		t.Fatal(err)
	}
	tmpfile.Close()

	// Test loading
	table, err := LoadCSV(tmpfile.Name())
	if err != nil {
		t.Fatalf("Failed to load CSV: %v", err)
	}

	if table.RowCount() != 3 {
		t.Errorf("Expected 3 rows, got %d", table.RowCount())
	}

	if table.ColumnCount() != 3 {
		t.Errorf("Expected 3 columns, got %d", table.ColumnCount())
	}

	if table.Headers[0] != "Name" {
		t.Errorf("Expected header 'Name', got '%s'", table.Headers[0])
	}

	// Check first row data
	row, _ := table.GetRow(0)
	if row[0] != "John" || row[1] != "30" || row[2] != "NYC" {
		t.Errorf("Expected ['John', '30', 'NYC'], got %v", row)
	}
}

func TestLoadEmptyCSV(t *testing.T) {
	tmpfile, err := os.CreateTemp("", "test*.csv")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())
	tmpfile.Close()

	_, err = LoadCSV(tmpfile.Name())
	if err == nil {
		t.Error("Expected error for empty CSV")
	}
}

func TestLoadFileInvalidFormat(t *testing.T) {
	_, err := LoadFile("test.txt")
	if err == nil {
		t.Error("Expected error for unsupported format")
	}
}

func TestLoadFileEmptyPath(t *testing.T) {
	_, err := LoadFile("")
	if err == nil {
		t.Error("Expected error for empty path")
	}
}

func TestLoadCSVWithVariableColumns(t *testing.T) {
	// CSV with inconsistent column counts
	content := `Name,Age,City
John,30
Jane,25,LA,Extra`

	tmpfile, err := os.CreateTemp("", "test*.csv")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	tmpfile.Write([]byte(content))
	tmpfile.Close()

	table, err := LoadCSV(tmpfile.Name())
	if err != nil {
		t.Fatalf("Failed to load CSV: %v", err)
	}

	// First row should be padded
	row1, _ := table.GetRow(0)
	if len(row1) != 3 {
		t.Errorf("Expected padded row length 3, got %d", len(row1))
	}
	if row1[2] != "" {
		t.Errorf("Expected empty string for padded cell, got '%s'", row1[2])
	}

	// Second row should be truncated
	row2, _ := table.GetRow(1)
	if len(row2) != 3 {
		t.Errorf("Expected truncated row length 3, got %d", len(row2))
	}
}
