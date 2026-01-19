package tui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/veliulugut/snapclean/internal/cleaner"
	"github.com/veliulugut/snapclean/internal/file"
	"github.com/veliulugut/snapclean/internal/models"
	"github.com/veliulugut/snapclean/internal/utils"
)

type fileSelectedMsg struct {
	path string
}

type fileLoadedMsg struct {
	success   bool
	message   string
	dataTable *models.DataTable
}

// Update processes all messages
func (m AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil

	case tickMsg:
		if m.currentView == splashView {
			m.splashTick++
			if m.splashTick > 40 {
				m.currentView = menuView
				m.splashDone = true
				return m, nil
			}
			return m, m.tickCmd()
		}
		return m, m.tickCmd()

	case tea.KeyMsg:
		if m.currentView == splashView && m.splashTick > 5 {
			m.currentView = menuView
			m.splashDone = true
			return m, nil
		}
		return m.handleKeyPress(msg)

	case fileSelectedMsg:
		if msg.path == "" {
			m.statusText = "✗ No file selected"
			return m, nil
		}

		m.loadedFile = msg.path
		m.statusText = "⏳ Loading file..."

		return m, func() tea.Msg {
			table, err := file.LoadFile(msg.path)
			if err != nil {
				return fileLoadedMsg{
					success:   false,
					message:   fmt.Sprintf("✗ Failed to load: %v", err),
					dataTable: nil,
				}
			}

			return fileLoadedMsg{
				success: true,
				message: fmt.Sprintf("✓ Loaded: %s (%d rows, %d columns)",
					table.FileName, table.RowCount(), table.ColumnCount()),
				dataTable: table,
			}
		}

	case fileLoadedMsg:
		m.statusText = msg.message
		if msg.success {
			m.dataTable = msg.dataTable
			m.scrollOffset = 0
			m.columnOffset = 0
		}
		return m, nil
	}

	return m, nil
}

// handleKeyPress routes key presses to appropriate handlers
func (m AppModel) handleKeyPress(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	// Cleaning view
	if m.currentView == cleaningView {
		return m.handleCleaningNavigation(msg)
	}

	// Table view with column menu
	if m.currentView == tableView && m.columnMenuMode {
		return m.handleColumnMenuNavigation(msg)
	}

	// Table view
	if m.currentView == tableView {
		return m.handleTableNavigation(msg)
	}

	// Help view
	if m.currentView == helpView {
		switch msg.String() {
		case "b", "esc":
			m.currentView = menuView
			return m, nil
		case "q", "ctrl+c":
			return m, tea.Quit
		}
		return m, nil
	}

	// Menu view
	switch msg.String() {
	case "ctrl+c", "q":
		return m, tea.Quit
	case "?":
		m.currentView = helpView
		return m, nil
	case "up", "k":
		if m.selectedItem > 0 {
			m.selectedItem--
		}
	case "down", "j":
		if m.selectedItem < len(m.options)-1 {
			m.selectedItem++
		}
	case "enter", " ":
		return m.executeSelection()
	}
	return m, nil
}

// handleTableNavigation handles navigation in table view
func (m AppModel) handleTableNavigation(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	if m.dataTable == nil {
		return m, nil
	}

	maxRowScroll := m.dataTable.RowCount() - m.pageSize
	if maxRowScroll < 0 {
		maxRowScroll = 0
	}

	switch msg.String() {
	case "b", "esc":
		m.currentView = menuView
		m.scrollOffset = 0
		m.columnOffset = 0
		return m, nil

	case "q", "ctrl+c":
		return m, tea.Quit

	// Vertical scroll
	case "up", "k":
		if m.scrollOffset > 0 {
			m.scrollOffset--
		}

	case "down", "j":
		if m.scrollOffset < maxRowScroll {
			m.scrollOffset++
		}

	case "pgup":
		m.scrollOffset -= m.pageSize
		if m.scrollOffset < 0 {
			m.scrollOffset = 0
		}

	case "pgdown":
		m.scrollOffset += m.pageSize
		if m.scrollOffset > maxRowScroll {
			m.scrollOffset = maxRowScroll
		}

	case "home":
		m.scrollOffset = 0

	case "end":
		m.scrollOffset = maxRowScroll

	// Horizontal scroll (for wide tables)
	case "left", "h":
		if m.columnOffset > 0 {
			m.columnOffset--
		}

	case "right", "l":
		maxColScroll := m.dataTable.ColumnCount() - 1
		if maxColScroll < 0 {
			maxColScroll = 0
		}
		if m.columnOffset < maxColScroll {
			m.columnOffset++
		}

	// Column menu toggle
	case "c":
		m.columnMenuMode = !m.columnMenuMode
		if m.columnMenuMode {
			m.selectedColumn = 0
			m.swapSourceCol = -1
			m.columnMessage = ""
		}
	}

	return m, nil
}

// handleColumnMenuNavigation handles navigation in column menu
func (m AppModel) handleColumnMenuNavigation(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	if m.dataTable == nil {
		return m, nil
	}

	maxCol := m.dataTable.ColumnCount() - 1

	switch msg.String() {
	case "b", "esc":
		m.columnMenuMode = false
		m.swapSourceCol = -1
		m.columnMessage = ""
		return m, nil

	case "q", "ctrl+c":
		return m, tea.Quit

	case "up", "k":
		if m.selectedColumn > 0 {
			m.selectedColumn--
		}

	case "down", "j":
		if m.selectedColumn < maxCol {
			m.selectedColumn++
		}

	case "enter":
		if m.swapSourceCol == -1 {
			// First selection - mark source column
			m.swapSourceCol = m.selectedColumn
			m.columnMessage = fmt.Sprintf("Column %d selected. Choose another column to swap with.", m.swapSourceCol+1)
		} else {
			// Second selection - perform swap
			if m.swapSourceCol == m.selectedColumn {
				m.columnMessage = "⚠ Cannot swap a column with itself."
			} else {
				m.swapColumns(m.swapSourceCol, m.selectedColumn)
				m.columnMessage = fmt.Sprintf("✓ Swapped columns %d ↔ %d", m.swapSourceCol+1, m.selectedColumn+1)
				m.swapSourceCol = -1
			}
		}
	}

	return m, nil
}

// swapColumns swaps two columns in the data table
func (m *AppModel) swapColumns(col1, col2 int) {
	if m.dataTable == nil || col1 < 0 || col2 < 0 {
		return
	}
	if col1 >= m.dataTable.ColumnCount() || col2 >= m.dataTable.ColumnCount() {
		return
	}

	// Swap headers
	m.dataTable.Headers[col1], m.dataTable.Headers[col2] = m.dataTable.Headers[col2], m.dataTable.Headers[col1]

	// Swap all row data
	for i := 0; i < m.dataTable.RowCount(); i++ {
		row, err := m.dataTable.GetRow(i)
		if err == nil && col1 < len(row) && col2 < len(row) {
			m.dataTable.Rows[i][col1], m.dataTable.Rows[i][col2] = m.dataTable.Rows[i][col2], m.dataTable.Rows[i][col1]
		}
	}
}

// handleCleaningNavigation handles navigation in cleaning view
func (m AppModel) handleCleaningNavigation(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "b", "esc":
		m.currentView = menuView
		m.cleaningMessage = ""
		return m, nil

	case "q", "ctrl+c":
		return m, tea.Quit

	case "up", "k":
		if m.cleaningSelected > 0 {
			m.cleaningSelected--
		}

	case "down", "j":
		if m.cleaningSelected < 4 {
			m.cleaningSelected++
		}

	case " ":
		// Toggle the selected option
		m.toggleCleaningOption(m.cleaningSelected)

	case "enter":
		if m.dataTable == nil {
			m.cleaningMessage = "⚠ No data loaded."
			return m, nil
		}

		beforeRows := m.dataTable.RowCount()
		beforeCols := m.dataTable.ColumnCount()

		// Apply cleaning
		cleaned := cleaner.ApplyCleaningOptions(m.dataTable, m.cleaningOptions)

		afterRows := cleaned.RowCount()
		afterCols := cleaned.ColumnCount()

		m.dataTable = cleaned

		// Show summary
		m.cleaningMessage = fmt.Sprintf(
			"✓ Cleaned! Rows: %d→%d (-%d)  Cols: %d→%d (-%d)",
			beforeRows, afterRows, beforeRows-afterRows,
			beforeCols, afterCols, beforeCols-afterCols,
		)
		m.statusText = m.cleaningMessage
	}

	return m, nil
}

// toggleCleaningOption toggles the selected cleaning option
func (m *AppModel) toggleCleaningOption(idx int) {
	switch idx {
	case 0:
		m.cleaningOptions.TrimWhitespace = !m.cleaningOptions.TrimWhitespace
	case 1:
		m.cleaningOptions.NormalizeHeaders = !m.cleaningOptions.NormalizeHeaders
	case 2:
		m.cleaningOptions.RemoveEmptyRows = !m.cleaningOptions.RemoveEmptyRows
	case 3:
		m.cleaningOptions.RemoveEmptyColumns = !m.cleaningOptions.RemoveEmptyColumns
	case 4:
		m.cleaningOptions.RemoveDuplicates = !m.cleaningOptions.RemoveDuplicates
	}
}

// executeSelection handles menu item selection
func (m AppModel) executeSelection() (tea.Model, tea.Cmd) {
	switch m.selectedItem {
	case 0: // Load CSV/Excel
		m.statusText = "⏳ Opening file picker..."
		return m, func() tea.Msg {
			path, err := utils.OpenFilePicker()
			if err != nil {
				return fileSelectedMsg{path: ""}
			}
			return fileSelectedMsg{path: path}
		}

	case 1: // View Data Table
		if m.dataTable == nil {
			m.statusText = "⚠ No data loaded. Please load a file first."
			return m, nil
		}
		m.currentView = tableView
		m.scrollOffset = 0
		return m, nil

	case 2: // Clean Data
		if m.dataTable == nil {
			m.statusText = "⚠ No data loaded. Please load a file first."
			return m, nil
		}
		m.currentView = cleaningView
		m.cleaningSelected = 0
		m.cleaningMessage = ""
		return m, nil

	case 6: // Help
		m.currentView = helpView
		return m, nil

	case 7: // Exit
		return m, tea.Quit

	default:
		m.statusText = "⚠ Feature coming soon..."
		return m, nil
	}
}
