package main

import (
	"RupenderSinghRathore/TaskMaster/internal/models"
	"errors"
	"fmt"
	"time"
)

func (app *application) handleArgs() error {
	var err error
	switch {
	case len(app.args) == 0:
		err = app.interactiveShellMode()
	case app.args[0] == "log":
		app.log()
	case app.args[0] == "add":
		err = app.add()
	case app.args[0] == "rm":
	case app.args[0] == "clear":
	case app.args[0] == "edit":
	}
	app.log()
	return err
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
func (app *application) add() error {
	n := len(app.args)
	var err error
	for i := 1; i < n; i++ {
		task := app.tasks.Append(app.args[i])
		if i+1 < n && app.args[i+1] == "-desc" {
			if i+2 < n {
				task.Description = app.args[i+2]
				i++
			} else {
				return errors.New("Empty Description")
			}
			i++
		}
		if i+1 < n && app.args[i+1] == "-time" {
			if i+2 < n {
				task.Deadline, err = getDeadline(app.args[i+2])
				if err != nil {
					return err
				}
				i++
			} else {
				return errors.New("Empty Time")
			}
			i++
		} else {
			t := time.Now().Add(time.Hour * 24)
			task.Deadline = &t
		}
	}
	return nil
}
