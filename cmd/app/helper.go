package main

func (m *Model) ToggleUp() {
	curr := m.cursor
	if curr == 0 {
		return
	}
	m.tasks[curr], m.tasks[curr-1] = m.tasks[curr-1], m.tasks[curr]
}

func (m *Model) ToggleDown() {
	curr := m.cursor
	if curr == len(m.tasks)-1 {
		return
	}
	m.tasks[curr], m.tasks[curr+1] = m.tasks[curr+1], m.tasks[curr]
}
