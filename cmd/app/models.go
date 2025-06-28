package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	tasks  []task
	cursor int
}

func (m Model) View() string {
	s := "$$    Task Master    $$\n\n"
	for i, t := range m.tasks {
		cursor := " "
		if i == m.cursor {
			cursor = ">"
		}
		done := ""
		if t.done {
			done = "o(≧▽≦)o"
		}
		s += fmt.Sprintf("%s [%s] %s\n", cursor, t.title, done)

	}
	s += "\nPress q to exit."
	return s
}
func (m Model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "n":
			if m.cursor < len(m.tasks) {
				m.cursor++
			}
		case "e":
			if m.cursor > 0 {
				m.cursor--
			}
		}
	}
	return m, nil
}
