package file

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/veliulugut/snapclean/internal/models"
	"github.com/xuri/excelize/v2"
)

// LoadFile automatically detects file type and loads it
func LoadFile(filePath string) (*models.DataTable, error) {
	if filePath == "" {
		return nil, fmt.Errorf("file path is empty")
	}

	ext := strings.ToLower(filepath.Ext(filePath))

	switch ext {
	case ".csv":
		return LoadCSV(filePath)
	case ".xlsx", ".xls":
		return LoadExcel(filePath)
	default:
		return nil, fmt.Errorf("unsupported file format: %s (supported: .csv, .xlsx)", ext)
	}
}

// LoadCSV loads data from a CSV file
func LoadCSV(filePath string) (*models.DataTable, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.TrimLeadingSpace = true
	reader.FieldsPerRecord = -1 // Allow variable number of fields
	reader.LazyQuotes = true
	reader.Comment = '#'

	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read CSV: %w", err)
	}

	if len(records) == 0 {
		return nil, fmt.Errorf("CSV file is empty")
	}

	// First row as headers
	table := models.NewDataTable(records[0])
	table.FilePath = filePath
	table.FileName = filepath.Base(filePath)

	// Add data rows
	for i := 1; i < len(records); i++ {
		row := records[i]

		// Pad row if shorter than headers
		if len(row) < len(table.Headers) {
			padded := make([]string, len(table.Headers))
			copy(padded, row)
			row = padded
		} else if len(row) > len(table.Headers) {
			// Truncate if longer
			row = row[:len(table.Headers)]
		}

		table.AddRow(row)
	}

	return table, nil
}

// LoadExcel loads data from an Excel file
func LoadExcel(filePath string) (*models.DataTable, error) {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open Excel file: %w", err)
	}
	defer f.Close()

	// Get first sheet
	sheets := f.GetSheetList()
	if len(sheets) == 0 {
		return nil, fmt.Errorf("no sheets found in Excel file")
	}

	sheetName := sheets[0]
	rows, err := f.GetRows(sheetName)
	if err != nil {
		return nil, fmt.Errorf("failed to read Excel sheet: %w", err)
	}

	if len(rows) == 0 {
		return nil, fmt.Errorf("Excel sheet is empty")
	}

	// First row as headers
	table := models.NewDataTable(rows[0])
	table.FilePath = filePath
	table.FileName = filepath.Base(filePath)

	// Add data rows
	for i := 1; i < len(rows); i++ {
		row := rows[i]

		// Pad row if shorter than headers
		if len(row) < len(table.Headers) {
			padded := make([]string, len(table.Headers))
			copy(padded, row)
			row = padded
		} else if len(row) > len(table.Headers) {
			// Truncate if longer
			row = row[:len(table.Headers)]
		}

		table.AddRow(row)
	}

	return table, nil
}
