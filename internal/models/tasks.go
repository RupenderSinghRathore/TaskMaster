package models

import (
	"time"
)

type Tasks []*Task
type Task struct {
	Title       string
	Description string
	Status      Status
	Deadline    *time.Time
}

func (t *Tasks) Append(title string) *Task {
	task := &Task{
		Title: title,
	}
	*t = append(*t, task)
	return task
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
