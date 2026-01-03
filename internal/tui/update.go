package tui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
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
			m.scrollOffset = 0 // Reset scroll
		}
		return m, nil
	}

	return m, nil
}

func (m AppModel) handleKeyPress(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	// Table view controls
	if m.currentView == tableView {
		return m.handleTableNavigation(msg)
	}

	// Help view controls
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

	// Menu controls
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

func (m AppModel) handleTableNavigation(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	if m.dataTable == nil {
		return m, nil
	}

	maxScroll := m.dataTable.RowCount() - m.pageSize
	if maxScroll < 0 {
		maxScroll = 0
	}

	maxColumnScroll := m.dataTable.ColumnCount() - m.visibleColumns
	if maxColumnScroll < 0 {
		maxColumnScroll = 0
	}

	switch msg.String() {
	case "b", "esc":
		m.currentView = menuView
		m.scrollOffset = 0
		m.columnOffset = 0
		return m, nil

	case "q", "ctrl+c":
		return m, tea.Quit

	// Vertical navigation
	case "up", "k":
		if m.scrollOffset > 0 {
			m.scrollOffset--
		}

	case "down", "j":
		if m.scrollOffset < maxScroll {
			m.scrollOffset++
		}

	// Horizontal navigation
	case "left", "h":
		if m.columnOffset > 0 {
			m.columnOffset--
		}

	case "right", "l":
		if m.columnOffset < maxColumnScroll {
			m.columnOffset++
		}

	// Page navigation
	case "pgup":
		m.scrollOffset -= m.pageSize
		if m.scrollOffset < 0 {
			m.scrollOffset = 0
		}

	case "pgdown":
		m.scrollOffset += m.pageSize
		if m.scrollOffset > maxScroll {
			m.scrollOffset = maxScroll
		}

	// Jump to start/end
	case "home":
		m.scrollOffset = 0

	case "end":
		m.scrollOffset = maxScroll

	case "ctrl+home":
		m.columnOffset = 0

	case "ctrl+end":
		m.columnOffset = maxColumnScroll
	}

	return m, nil
}

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
