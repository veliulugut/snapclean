package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

type viewMode int

const (
	menuView viewMode = iota
	tableView
	cleaningView
	exportView
)

type AppModel struct {
	currentView  viewMode
	selectedItem int
	options      []string
	dataRows     [][]string
	columnNames  []string
	loadedFile   string
	statusText   string
}

func InitialModel() AppModel {
	return AppModel{
		currentView:  menuView,
		selectedItem: 0,
		options: []string{
			"[ LOAD ]   Load CSV/Excel File",
			"[ VIEW ]   View Data Table",
			"[ CLEAN ]  Clean Data",
			"[ PIVOT ]  Summarize / Pivot",
			"[ CHECK ]  Run QA Checks",
			"[ SAVE ]   Export Data",
			"[ HELP ]   Show Help",
			"[ EXIT ]   Quit Application",
		},
	}
}

func (m AppModel) Init() tea.Cmd {
	return nil
}
