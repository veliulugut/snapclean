package utils

import "github.com/ncruces/zenity"

func OpenFilePicker() (string, error) {
	filename, err := zenity.SelectFile(
		zenity.Title("Select a file to clean"),
		zenity.FileFilters{
			{Name: "CSV and Excel files", Patterns: []string{"*.csv", "*.xlsx"}},
		},
	)

	if err != nil {
		return "", err
	}

	return filename, nil
}
