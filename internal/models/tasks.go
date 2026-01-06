package models

import (
	"time"
)

type Tasks []*Task
type Task struct {
	Title       string    `csv:"Title"`
	Description string    `csv:"Description"`
	Status      Status    `csv:"Status"`
	Deadline    time.Time `csv:"Deadline"`
}

func (t *Tasks) Append(title string, deadline time.Time) {
	task := &Task{
		Title:    title,
		Deadline: deadline,
	}
	*t = append(*t, task)
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
