package main

import (
	"RupenderSinghRathore/TaskMaster/internal/models"
	"fmt"
	"os"
	"text/tabwriter"
)

type application struct {
	tasks  models.Tasks
	args   []string
	writer *tabwriter.Writer
}

func main() {
	tasks, err := loadTasks()
	if err != nil {
		fmt.Printf("err: %s\n", err.Error())
	}
	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.DiscardEmptyColumns)
	app := application{
		tasks:  tasks,
		args:   os.Args[1:],
		writer: writer,
	}
	// app.tasks.Append("fuck you")
	// app.tasks[0].Status = models.Done
	// app.tasks.Append("make some sense")
	app.handleArgs()
	if err := saveTasks(app.tasks); err != nil {
		fmt.Printf("err: %s\n", err.Error())
	}
}
