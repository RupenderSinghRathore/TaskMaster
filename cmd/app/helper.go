package main

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
