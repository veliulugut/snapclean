package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/veliulugut/snapclean/internal/utils"
)

type fileSelectedMsg struct {
	path string
}

func (m AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		return m.handleKeyPress(msg)
	case fileSelectedMsg:
		m.loadedFile = msg.path
		m.statusText = "File loaded: " + msg.path
		return m, nil
	}

	return m, nil
}

func (m AppModel) handleKeyPress(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c", "q":
		return m, tea.Quit

	case "up":
		if m.selectedItem > 0 {
			m.selectedItem--
		}

	case "down":
		if m.selectedItem < len(m.options)-1 {
			m.selectedItem++
		}

	case "enter":
		return m.executeSelection()
	}
	return m, nil
}

func (m AppModel) executeSelection() (tea.Model, tea.Cmd) {
	switch m.selectedItem {
	case 0: // Load CSV/Excel
		return m, func() tea.Msg {
			path, err := utils.OpenFilePicker()
			if err != nil {
				return fileSelectedMsg{path: ""}
			}
			return fileSelectedMsg{path: path}
		}
	case 6: // Exit
		return m, tea.Quit
	default:
		m.statusText = "Feature coming soon..."
		return m, nil
	}
}
