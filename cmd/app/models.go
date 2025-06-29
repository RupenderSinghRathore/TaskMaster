package main

import (
	"fmt"
	"strconv"

	"github.com/RupenderSinghRathore/TaskMaster/internal/models"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	tasks   []models.Task
	cursor  int
	newTask textinput.Model
	styles  *models.Styles
	adding  bool
	width   int
	height  int
}

func (m Model) View() string {
	if m.width == 0 {
		return "loading..."
	}
	title := "ToDo deez nuts.."
	tasks := ""
	for i, t := range m.tasks {
		if t.Title == "" {
			continue
		}
		cursor := " "
		if i == m.cursor {
			cursor = ">"
		}
		done := ""
		if t.Done {
			done = "🔥o(≧▽≦)o🔥"
		}
		index := strconv.Itoa(i + 1)
		tasks += fmt.Sprintf(
			"%s %s %s %s\n",
			m.styles.TasksColor.Render(index),
			m.styles.CursorStyle.Render(cursor),
			m.styles.TasksColor.Render(t.Title),
			m.styles.CursorStyle.Render(done))
	}
	title = m.styles.TitleField.Render(title)
	tasks = m.styles.TasksField.Render(tasks)
	input := ""
	if m.adding {
		input = m.styles.InputField.Render(m.newTask.View())
	}
	content := lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		tasks,
		input,
	)

	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Top,
		content,
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

		case "ctrl+c":
			m.WriteToFile()
			return m, tea.Quit
		case "q":
			if !m.adding {
				m.WriteToFile()
				return m, tea.Quit
			}

		case "n", "down":
			if !m.adding && m.cursor < len(m.tasks)-1 {
				m.cursor++
			}
		case "e", "up":
			if !m.adding && m.cursor > 0 {
				m.cursor--
			}
		case "ctrl+n", "ctrl+down":
			m.ToggleDown()
		case "ctrl+e", "ctrl+up":
			m.ToggleUp()

		case "ctrl+a":
			m.adding = true
			m.newTask.SetValue("")
		case "enter":
			if m.adding {
				m.AddTask()
			} else {
				m.ToggleDone()
			}
		case "esc":
			m.adding = false

		case "d":
			if !m.adding {
				m.RemoveTask()
			}
		case "alt+r":
			m.tasks = []models.Task{}
			m.cursor = 0
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	}
	var cmd tea.Cmd
	m.newTask, cmd = m.newTask.Update(msg)
	return m, cmd
}
