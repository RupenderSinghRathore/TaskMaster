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
func (t *Tasks) Delete(ids []int) {
	newTasks := Tasks{}
	id := ids[0]
	for j, task := range *t {
		if j != id-1 {
			newTasks = append(newTasks, task)
		}
		(*t)[j] = nil
	}
	*t = newTasks
}
