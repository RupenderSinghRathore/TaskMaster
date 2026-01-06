package main

import (
	"RupenderSinghRathore/TaskMaster/internal/models"
	"fmt"
)

func (app *application) handleArgs() {
	switch {
	case len(app.args) == 0:
		app.interactiveShellMode()
	case app.args[0] == "log":
		app.log()
	case app.args[0] == "add":
	case app.args[0] == "rm":
	case app.args[0] == "clear":
	case app.args[0] == "edit":
	}
}

func (app *application) log() {
	var isDone string
	for i, task := range app.tasks {
		isDone = "❌"
		if task.Status == models.Done {
			isDone = "✔️  ☆*: .｡. o(≧▽≦)o .｡.:*☆"
		}
		fmt.Fprintf(app.writer, "    %d >\t%s\t%s\n", i+1, task.Title, isDone)
	}
	if err := app.writer.Flush(); err != nil {
		fmt.Printf("err: %s\n", err.Error())
	}
}
