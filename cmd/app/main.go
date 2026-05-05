package main

import (
	"fmt"
	"os"
	"text/tabwriter"

	"RupenderSinghRathore/TaskMaster/internal/models"

	"golang.org/x/term"
)

type application struct {
	tasks         models.Tasks
	args          []string
	writer        *tabwriter.Writer
	isInteractive bool
	terminal      *term.Terminal
}

func main() {
	tasks, err := loadTasks()
	if err != nil {
		handleErr(err)
	}

	app := application{
		tasks:         tasks,
		args:          os.Args[1:],
		isInteractive: len(os.Args) == 1,
	}

	if err := app.handleModes(); err != nil {
		handleErr(err)
	}

	if err := saveTasks(app.tasks); err != nil {
		handleErr(err)
	}
}

func (app *application) handleModes() error {
	if app.isInteractive {
		if err := app.shellMode(); err != nil {
			return err
		}
	} else {
		app.writer = tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.DiscardEmptyColumns)
		msg, err := app.handleArgs()
		if err != nil {
			return err
		}
		if len(msg) != 0 {
			fmt.Printf("\n%s\n", msg)
		}
	}
	return nil
}
