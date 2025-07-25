package main

import (
	"fmt"
	"os"

	"github.com/RupenderSinghRathore/TaskMaster/internal/models"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	// p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

func initialModel() Model {
	m := Model{
		newTask: textinput.New(),
		styles:  models.DefaultStyle(),
	}
	m.ReadFromFile()
	m.newTask.Width = 100
	m.newTask.Placeholder = "New Task.."
	m.newTask.Focus()

	return m
}
