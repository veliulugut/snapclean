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
	scrollOffset int // vertical scroll (rows)
	pageSize     int // number of rows per page
	columnOffset int // horizontal scroll (columns)

	// Column management state
	columnMenuMode bool   // true when column menu is active
	selectedColumn int    // currently selected column in menu
	swapSourceCol  int    // first column selected for swap (-1 if none)
	columnMessage  string // feedback message for column operations

	// Cleaning state
	cleaningOptions  models.CleanOptions
	cleaningSelected int
	cleaningMessage  string

	// UI State
	loadedFile string
	statusText string
	width      int
	height     int

	// Splash state
	splashTick int
	splashDone bool
}

func InitialModel() AppModel {
	return AppModel{
		currentView:      splashView,
		selectedItem:     0,
		scrollOffset:     0,
		columnOffset:     0,
		pageSize:         10, // show 10 rows at a time
		columnMenuMode:   false,
		selectedColumn:   0,
		swapSourceCol:    -1,
		columnMessage:    "",
		cleaningSelected: 0,
		cleaningMessage:  "",
		splashTick:       0,
		splashDone:       false,
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

		// Cleaning defaults (all enabled)
		cleaningOptions: models.CleanOptions{
			TrimWhitespace:     true,
			NormalizeHeaders:   true,
			RemoveEmptyRows:    true,
			RemoveEmptyColumns: true,
			RemoveDuplicates:   true,
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
