package tui

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/veliulugut/snapclean/internal/models"
)

type viewMode int

const (
	splashView viewMode = iota
	menuView
	tableView
	cleaningView
	exportView
	helpView
)

type AppModel struct {
	currentView  viewMode
	selectedItem int
	options      []string

	// Data
	dataTable *models.DataTable

	// Table view state
	scrollOffset   int // vertical scroll (rows)
	columnOffset   int // horizontal scroll (columns)
	pageSize       int // number of rows per page
	visibleColumns int // columns visible at once

	// UI State
	loadedFile string
	statusText string
	width      int
	height     int
	splashTick int
	splashDone bool
}

func InitialModel() AppModel {
	return AppModel{
		currentView:    splashView,
		selectedItem:   0,
		splashTick:     0,
		splashDone:     false,
		dataTable:      nil,
		scrollOffset:   0,
		pageSize:       10, // show 10 rows at a time
		visibleColumns: 5,
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
	return tea.Batch(
		tea.WindowSize(),
		m.tickCmd(),
	)
}

type tickMsg time.Time

func (m AppModel) tickCmd() tea.Cmd {
	return tea.Tick(100*time.Millisecond, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}
