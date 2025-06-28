package main

import "github.com/charmbracelet/lipgloss"

type Styles struct {
	BorderColor lipgloss.Color
	TasksField  lipgloss.Style
	InputField  lipgloss.Style
}

func DefaultStyle() *Styles {
	s := &Styles{}
	s.BorderColor = lipgloss.Color("#09c2e3")
	s.InputField = lipgloss.NewStyle().BorderForeground(s.BorderColor).BorderStyle(lipgloss.NormalBorder()).Padding(1).Width(80)
	return s
}
