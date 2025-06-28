package main

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	tasks   []task
	cursor  int
	newTask textinput.Model
	styles  *Styles
	adding  bool
	width   int
	height  int
}

func (m Model) View() string {
	if m.width == 0 {
		return "loading..."
	}
	FTask := []string{"Task Master -->"}
	for i, t := range m.tasks {
		cursor := " "
		if i == m.cursor {
			cursor = ">"
		}
		done := ""
		if t.done {
			// done = "o(≧▽≦)o"
			done = "🔥o(≧▽≦)o🔥"
		}
		FTask = append(FTask, fmt.Sprintf("%s [%s] %s", cursor, t.title, done))
	}
	FTask = append(FTask, m.styles.InputField.Render(m.newTask.View()))

	return lipgloss.Place(
		m.width,
		m.height/2,
		lipgloss.Top,
		lipgloss.Top,
		lipgloss.JoinVertical(
			lipgloss.Left,
			FTask...,
		),
	)
}
func (m Model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// fmt.Printf("key:\"%v\"", msg)
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {

		// case "a":
		// 	m.adding = true
		case "ctrl+c", "q":
			m.WriteToFile()
			return m, tea.Quit
		case "down":
			if m.cursor < len(m.tasks)-1 {
				m.cursor++
			}
		case "up":
			if m.cursor > 0 {
				m.cursor--
			}

		case "ctrl+up":
			m.ToggleUp()
		case "ctrl+down":
			m.ToggleDown()

		case "ctrl+f":
			curr := &m.tasks[m.cursor]
			if curr.done {
				curr.done = false
			} else {
				curr.done = true
			}
		case "ctrl+d":
			m.RemoveTask()

		case "ctrl+shift+d":
			m.tasks = []task{}
			m.cursor = 0
		case "enter":
			newTask := m.newTask.Value()
			m.newTask.SetValue("")
			m.tasks = append(m.tasks, task{title: newTask})
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	}
	var cmd tea.Cmd
	m.newTask, cmd = m.newTask.Update(msg)
	return m, cmd
}
