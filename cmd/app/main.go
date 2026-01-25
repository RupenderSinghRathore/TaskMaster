package main

import (
	"RupenderSinghRathore/TaskMaster/internal/models"
	"fmt"
	"os"
	"text/tabwriter"

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
	insertionSort(tasks)
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
func (app *application) handleArgs() (string, error) {
	var err error
	var msg string
	switch app.args[0] {
	case "log":
		app.log()
	case "add":
		msg, err = app.add()
	case "rm":
		msg, err = app.remove()
	case "done":
		msg, err = app.done()
	case "undo":
		msg, err = app.undo()
	case "swap":
		msg, err = app.swap()
	case "purge":
		msg = app.purge()
	case "clear":
		msg = app.clear()
	case "edit":
		msg, err = app.edit()
	case "help":
		app.help()
	default:
		return "", fmt.Errorf("%s is not a valid option, try help command", app.args[0])
	}
	return msg, err
}
