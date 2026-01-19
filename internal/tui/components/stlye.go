package components

import "github.com/charmbracelet/lipgloss"

// Header styles
var HeaderStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#FFFFFF")).
	Background(lipgloss.Color("#5F5FFF")).
	Padding(0, 2).
	MarginBottom(1).
	Align(lipgloss.Center)

// Menu item styles
var SelectedItemStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#FFFFFF")).
	Background(lipgloss.Color("#5F87FF")).
	Bold(true).
	Padding(0, 1).
	Width(50)

var NormalItemStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#BCBCBC")).
	Padding(0, 1).
	Width(50)

// Tag styles
var TagStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#87FF87")).
	Bold(true)

var SelectedTagStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#FFFF87")).
	Bold(true)

// Status bar
var StatusBarStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#FFFFFF")).
	Background(lipgloss.Color("#5F5F87")).
	Padding(0, 1).
	MarginTop(1)

var StatusTextStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#87FFAF"))

// Help section
var HelpStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#8787AF")).
	MarginTop(1).
	Align(lipgloss.Center)

// Help view styles
var HelpTitleStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#FFFFFF")).
	Background(lipgloss.Color("#5F5FFF")).
	Padding(0, 2).
	MarginBottom(2).
	Align(lipgloss.Center)

var HelpSectionStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#5F87FF")).
	Bold(true).
	MarginTop(1)

var HelpTextStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#BCBCBC")).
	MarginLeft(4)

// Container
var ContainerStyle = lipgloss.NewStyle().
	Border(lipgloss.DoubleBorder()).
	BorderForeground(lipgloss.Color("#5F87FF")).
	Padding(2, 4).
	Width(70).
	Align(lipgloss.Center)

var SplashContainerStyle = lipgloss.NewStyle().
	Border(lipgloss.DoubleBorder()).
	BorderForeground(lipgloss.Color("#5F87FF")).
	Padding(3, 4).
	Width(90).
	Align(lipgloss.Center)

// Column menu styles
var SelectedStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#FFFF87")).
	Bold(true)

var SuccessMessageStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#87FF87")).
	Bold(true).
	MarginBottom(1)
