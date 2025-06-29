package models

import (
	"github.com/charmbracelet/lipgloss"
)

type Styles struct {
	TitleField  lipgloss.Style
	TasksField  lipgloss.Style
	TasksColor  lipgloss.Style
	InputField  lipgloss.Style
	CursorStyle lipgloss.Style
}

func DefaultStyle() *Styles {
	s := &Styles{}

	s.CursorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#ff0000"))
	s.TasksColor = lipgloss.NewStyle().Foreground(lipgloss.Color("#82aaff"))
	s.TitleField = lipgloss.NewStyle().
		Padding(4, 1, 1, 1).
		Bold(true).
		Foreground(lipgloss.Color("#c792ea")).Width(71)

	s.TasksField = lipgloss.NewStyle().
		Padding(1).
		Bold(true).
		Width(71)

	s.InputField = lipgloss.NewStyle().
		BorderForeground(lipgloss.Color("#165e7a")).
		BorderStyle(lipgloss.NormalBorder()).
		Padding(1).
		Width(70).
		Foreground(lipgloss.Color("#00dddd")).
		UnsetBlink().
		Blink(false)
	return s
}
