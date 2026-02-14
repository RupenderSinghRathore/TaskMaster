package models

import (
	"fmt"
	"strings"
	"time"
)

type Tasks []*Task
type Task struct {
	Title       string
	Description string
	Status      Status
	Deadline    time.Time
}

func (t *Tasks) Append(title string) *Task {
	task := &Task{
		Title:    title,
		Deadline: time.Now().Add(24 * time.Hour).Round(time.Second),
		Status:   Pending,
	}
	*t = append(*t, task)
	return task
}
func (t *Tasks) AppendTasks(tasks Tasks) {
	*t = append(*t, tasks...)
}
func (t *Tasks) Delete(ids map[int]bool) {
	newTasks := Tasks{}
	for j, task := range *t {
		if _, ok := ids[j]; !ok {
			newTasks = append(newTasks, task)
		}
		(*t)[j] = nil
	}
	*t = newTasks
}

func (t *Tasks) Purge() {
	newTasks := Tasks{}
	for j, task := range *t {
		if task.Status != Done {
			newTasks = append(newTasks, task)
		}
		(*t)[j] = nil
	}
	*t = newTasks
}
func (t Tasks) InsertionSort() {
	var j int
	for i, task := range t {
		curr := time.Until(task.Deadline)
		j = i - 1
		for j >= 0 && time.Until(t[j].Deadline) > curr {
			t[j+1] = t[j]
			j--
		}
		t[j+1] = task
	}
}

func (t Tasks) String() string {
	str := strings.Builder{}
	str.WriteString("{\n")
	for _, task := range t {
		if task != nil {
			str.WriteString(fmt.Sprintf("%+v", *task))
		}
	}
	str.WriteString("\n}")
	return str.String()
}
