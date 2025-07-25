package main

import (
	"strings"
	"unicode"

	"github.com/RupenderSinghRathore/TaskMaster/internal/models"
)

func (m *Model) ToggleUp() {
	curr := m.cursor
	if curr == 0 {
		return
	}
	m.tasks[curr], m.tasks[curr-1] = m.tasks[curr-1], m.tasks[curr]
	m.cursor--
}

func (m *Model) ToggleDown() {
	curr := m.cursor
	if curr == len(m.tasks)-1 {
		return
	}
	m.tasks[curr], m.tasks[curr+1] = m.tasks[curr+1], m.tasks[curr]
	m.cursor++
}

func (m *Model) RemoveTask() {
	if len(m.tasks) == 0 {
		return
	}
	curr := m.cursor
	if m.cursor != 0 && m.cursor == len(m.tasks)-1 {
		m.cursor--
	}
	m.tasks = append(m.tasks[0:curr], m.tasks[curr+1:]...)
}

func (m *Model) ToggleDone() {
	curr := &m.tasks[m.cursor]
	if curr.Done {
		curr.Done = false
	} else {
		curr.Done = true
	}
}

func (m *Model) AddTask() {
	newTask := m.newTask.Value()
	newTask = strings.TrimSpace(newTask)
	newTask = Title(newTask)
	if newTask != "" {
		m.tasks = append(m.tasks, models.Task{Title: newTask})
		m.newTask.SetValue("")
		m.adding = false
	}
}
func Title(s string) string {
	r := []rune(s)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}
