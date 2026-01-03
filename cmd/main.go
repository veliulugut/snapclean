package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/veliulugut/snapclean/internal/tui"
)

func main() {
	app := tui.InitialModel()
	program := tea.NewProgram(app, tea.WithAltScreen())

	if _, err := program.Run(); err != nil {
		log.Fatal(err)
	}
}
