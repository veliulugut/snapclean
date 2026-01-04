package cleaner

import (
    "testing"

    "github.com/veliulugut/snapclean/internal/models"
)

func TestRemoveEmptyRows(t *testing.T) {
    dt := models.NewDataTable([]string{"Name", "Age"})
    dt.AddRow([]string{"John", "30"})
    dt.AddRow([]string{" ", "   "}) // empty after trim
    dt.AddRow([]string{"Jane", "25"})

    got := RemoveEmptyRows(dt)

    if got.RowCount() != 2 {
        t.Fatalf("expected 2 rows, got %d", got.RowCount())
    }

    row, _ := got.GetRow(1)
    if row[0] != "Jane" {
        t.Errorf("expected second row to be Jane, got %v", row)
    }
}

func TestRemoveEmptyColumns(t *testing.T) {
    dt := models.NewDataTable([]string{"Name", "Empty", "City"})
    dt.AddRow([]string{"John", "", "NYC"})
    dt.AddRow([]string{"Jane", "", "LA"})

    got := RemoveEmptyColumns(dt)

    if got.ColumnCount() != 2 {
        t.Fatalf("expected 2 columns, got %d", got.ColumnCount())
    }

    if got.Headers[0] != "Name" || got.Headers[1] != "City" {
        t.Errorf("unexpected headers: %v", got.Headers)
    }
}

func TestNormalizeHeaders(t *testing.T) {
    dt := models.NewDataTable([]string{" First Name ", "Age @#$", "City/State"})

    got := NormalizeHeaders(dt)

    want := []string{"first_name", "age", "citystate"}
    for i := range want {
        if got.Headers[i] != want[i] {
            t.Errorf("header %d: want %s, got %s", i, want[i], got.Headers[i])
        }
    }
}

func TestRemoveDuplicates(t *testing.T) {
    dt := models.NewDataTable([]string{"Name", "Age"})
    dt.AddRow([]string{"John", "30"})
    dt.AddRow([]string{"John", "30"})
    dt.AddRow([]string{"Jane", "25"})

    got := RemoveDuplicates(dt)

    if got.RowCount() != 2 {
        t.Fatalf("expected 2 unique rows, got %d", got.RowCount())
    }
}

func TestTrimWhitespace(t *testing.T) {
    dt := models.NewDataTable([]string{" Name ", " Age "})
    dt.AddRow([]string{" John ", " 30 "})

    got := TrimWhitespace(dt)

    if got.Headers[0] != "Name" || got.Headers[1] != "Age" {
        t.Errorf("unexpected headers after trim: %v", got.Headers)
    }

    row, _ := got.GetRow(0)
    if row[0] != "John" || row[1] != "30" {
        t.Errorf("unexpected row after trim: %v", row)
    }
}

func TestApplyCleaningOptions(t *testing.T) {
    dt := models.NewDataTable([]string{"  Name  ", "Empty"})
    dt.AddRow([]string{"  John  ", ""})
    dt.AddRow([]string{"  John  ", ""}) // duplicate
    dt.AddRow([]string{"", ""})         // empty row

    opts := models.CleanOptions{
        TrimWhitespace:     true,
        NormalizeHeaders:   true,
        RemoveEmptyRows:    true,
        RemoveEmptyColumns: true,
        RemoveDuplicates:   true,
    }

    got := ApplyCleaningOptions(dt, opts)

    if got.RowCount() != 1 {
        t.Fatalf("expected 1 row after cleaning, got %d", got.RowCount())
    }

    if got.ColumnCount() != 1 {
        t.Fatalf("expected 1 column after cleaning, got %d", got.ColumnCount())
    }

    if got.Headers[0] != "name" {
        t.Errorf("expected normalized header 'name', got %s", got.Headers[0])
    }

    row, _ := got.GetRow(0)
    if row[0] != "John" {
        t.Errorf("expected cleaned row value 'John', got %v", row)
    }
}

func TestApplyCleaningOptionsNil(t *testing.T) {
    var dt *models.DataTable
    opts := models.CleanOptions{TrimWhitespace: true}

    if got := ApplyCleaningOptions(dt, opts); got != nil {
        t.Errorf("expected nil result when input is nil")
    }
}