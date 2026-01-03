package models

import "fmt"

// NewDataTable creates a new DataTable with given headers
func NewDataTable(headers []string) *DataTable {
	return &DataTable{
		Headers: headers,
		Rows:    make([][]string, 0),
	}
}

// DataTable represent a structered a data table with headers and rows
type DataTable struct {
	Headers  []string   // Column headers
	Rows     [][]string // Data rows
	FilePath string     // Path to the source file
	FileName string     // Name of the source file
}

// CleanOptions defines options for data cleaning operations
type CleanOptions struct {
	RemoveEmptyRows    bool // Remove rows with all empty cells
	RemoveEmptyColumns bool // Remove columns with all empty cells
	NormalizeHeaders   bool // Normalize header names(lowercase,underscores)
	RemoveDuplicates   bool // Remove duplicate rows
	TrimWhitespace     bool // Trim leading/trailing whitespace from cells
}

// ExportOptions defines options for exporting data
type ExportOptions struct {
	Format   string // "csv" or "xlsx"
	FilePath string // Destination file path
}

// RowCount returns the number of rows in the table
func (dt *DataTable) RowCount() int {
	return len(dt.Rows)
}

// ColumnCount returns the number of columns in the table
func (dt *DataTable) ColumnCount() int {
	return len(dt.Headers)
}

// AddRow adds a new row to the table
// Returns error if row length doesnt match headers
func (dt *DataTable) AddRow(row []string) error {
	if len(row) != len(dt.Headers) {
		return fmt.Errorf("row lenght %d doesnt match headers lenght %d", len(row), len(dt.Headers))
	}

	dt.Rows = append(dt.Rows, row)
	return nil
}

// GetColumn returns all values in the specified column
// Returns error if column index is out of bounds
func (dt *DataTable) GetColumn(index int) ([]string, error) {
	if index < 0 || index >= len(dt.Headers) {
		return nil, fmt.Errorf("column index %d out of bounds", index)
	}

	column := make([]string, len(dt.Rows))
	for i, row := range dt.Rows {
		if index < len(row) {
			column[i] = row[index]
		}
	}

	return column, nil
}

// IsEmpty checks if the table has no data
func (dt *DataTable) IsEmpty() bool {
	return len(dt.Rows) == 0
}

// GetRow returns the row at the given index
// Returns error if index is out of bounds
func (dt *DataTable) GetRow(index int) ([]string, error) {
	if index < 0 || index >= len(dt.Rows) {
		return nil, fmt.Errorf("row index %d out of bounds", index)
	}

	return dt.Rows[index], nil
}

// Clone creates a deep copy of the DataTable
func (dt *DataTable) Clone() *DataTable {
	newTable := &DataTable{
		Headers:  make([]string, len(dt.Headers)),
		Rows:     make([][]string, len(dt.Rows)),
		FilePath: dt.FilePath,
		FileName: dt.FileName,
	}

	copy(newTable.Headers, dt.Headers)

	for i, row := range dt.Rows {
		newTable.Rows[i] = make([]string, len(row))
		copy(newTable.Rows[i], row)
	}

	return newTable
}
