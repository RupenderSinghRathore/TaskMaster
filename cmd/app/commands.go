package main

import (
	"RupenderSinghRathore/TaskMaster/internal/models"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"time"
)

var ErrNotEnoughArgs = errors.New("not enough args")

func (app *application) log() {
	defer func() {
		if err := app.writer.Flush(); err != nil {
			fmt.Fprintf(os.Stderr, "err: %s\n", err.Error())
		}
	}()
	if len(app.tasks) == 0 {
		fmt.Fprint(app.writer, "Your logs are empty..\n")
		return
	}
	var detailed, longTerm bool
	for _, arg := range app.args[1:] {
		switch arg {
		case "-l":
			longTerm = true
		case "-d":
			detailed = true
		case "-dl":
			detailed, longTerm = true, true
		}
	}
	if detailed {
		fmt.Fprintf(
			app.writer,
			"    %s \t%s\t%s\t%s\t%s\t%s\n",
			"id",
			"Title",
			"Description",
			"Status",
			"Deadline",
			"Time Period",
		)
	}
	var isDone string
	for i, task := range app.tasks {
		isShortTerm := time.Until(task.Deadline) < 24*time.Hour
		if longTerm && isShortTerm {
			continue
		}
		if detailed {
			fmt.Fprintf(
				app.writer,
				"    %d \t%s\t%s\t%s\t%s\t%s\n",
				i+1,
				task.Title,
				task.Description,
				task.Status,
				task.Deadline.Format(time.RFC822),
				getTimeperiod(task.Deadline),
			)
		} else if longTerm || isShortTerm {
			if task.Status == models.Done {
				isDone = "✔️  ☆*: .｡. o(≧▽≦)o .｡.:*☆"
			} else {
				isDone = "❌"
			}
			fmt.Fprintf(app.writer, "    %d >\t%s\t%s\n", i+1, task.Title, isDone)
		}
	}
}
func (app *application) undo() (string, error) {
	count := 0
	for _, id := range app.args[1:] {
		idx, err := getTaskId(id, len(app.tasks))
		if err != nil {
			return "", err
		}
		if app.tasks[idx].Status == models.Done {
			app.tasks[idx].Status = models.Pending
		}
		count++
	}
	var msg string
	if count > 1 {
		msg = "Tasks undone.."
	} else {
		msg = "Task undone.."
	}
	return msg, nil

}
func (app *application) add() (string, error) {
	n := len(app.args)
	if n < 2 {
		return "", ErrNotEnoughArgs
	}

	var err error
	tasks := models.Tasks{}

	addCmd := flag.NewFlagSet("add", flag.PanicOnError)

	desc := addCmd.String("desc", "", "new description")
	deadline := addCmd.String("time", "", "new deadline")
	status := addCmd.String("status", "", "new status")

	for i := 1; i < n; {
		*desc = ""
		*deadline = ""
		*status = ""

		title := app.args[i]
		if title == "" {
			return "", errors.New("empty title")
		}

		task := tasks.Append(capitalize(title))
		i++

		if i == n {
			break
		}

		addCmd.Parse(app.args[i:])

		if *desc != "" {
			task.Description = *desc
		}
		if *deadline != "" {
			task.Deadline, err = getDeadline(*deadline)
			if err != nil {
				return "", fmt.Errorf("%s is not a valid time", *deadline)
			}
			if task.Description == "" && time.Until(task.Deadline) > 24*time.Hour {
				task.Description = "Long term task"
			}
		}
		if *status != "" {
			err := task.Status.UpdateStatus(*status)
			if err != nil {
				return "", err
			}
		}

		i += len(app.args[i:]) - addCmd.NArg()
	}

	app.tasks.AppendTasks(tasks)
	insertionSort(app.tasks)

	msg := "Task added to your log.."
	if len(tasks) > 1 {
		msg = "Tasks added to your log.."
	}

	return msg, nil
}

func (app *application) remove() (string, error) {
	ids := make(map[int]bool, len(app.args[1:]))
	count := 0
	for _, id := range app.args[1:] {
		i, err := getTaskId(id, len(app.tasks))
		if err != nil {
			return "", err
		}
		ids[i] = true
		count++
	}
	app.tasks.Delete(ids)
	var msg string
	if count > 1 {
		msg = "Tasks removed from your log.."
	} else {
		msg = "Task removed from your log.."
	}
	return msg, nil
}
func (app *application) done() (string, error) {
	count := 0
	for _, id := range app.args[1:] {
		idx, err := getTaskId(id, len(app.tasks))
		if err != nil {
			return "", err
		}
		app.tasks[idx].Status = models.Done
		count++
	}
	var msg string
	if count > 1 {
		msg = "Tasks marked done.."
	} else {
		msg = "Task marked done.."
	}
	return msg, nil
}
func (app *application) clear() string {
	app.tasks = models.Tasks{}
	return "All your logs are cleared.."
}
func (app *application) edit() (string, error) {
	if len(app.args) < 3 {
		return "", ErrNotEnoughArgs
	}

	id, err := getTaskId(app.args[1], len(app.tasks))
	if err != nil {
		return "", err
	}

	task := app.tasks[id]

	editCmd := flag.NewFlagSet("edit", flag.PanicOnError)

	title := editCmd.String("title", "", "new title")
	desc := editCmd.String("desc", "", "new description")
	time := editCmd.String("time", "", "new deadline")
	status := editCmd.String("status", "", "new status")

	editCmd.Parse(app.args[2:])

	if *title != "" {
		task.Title = capitalize(*title)
	}
	if *desc != "" {
		task.Description = *desc
	}
	if *time != "" {
		task.Deadline, err = getDeadline(*time)
		if err != nil {
			return "", fmt.Errorf("%s is not a valid time", *time)
		}
		insertionSort(app.tasks)
	}
	if *status != "" {
		err := task.Status.UpdateStatus(*status)
		if err != nil {
			return "", err
		}
	}

	return "Task edited..", nil
}
func (app *application) swap() (string, error) {
	// swap 1 2
	if len(app.args) <= 3 {
		tasklen := len(app.tasks)
		id1, err := getTaskId(app.args[1], tasklen)
		if err != nil {
			return "", err
		}
		id2, err := getTaskId(app.args[2], tasklen)
		if err != nil {
			return "", err
		}
		app.tasks[id1], app.tasks[id2] = app.tasks[id2], app.tasks[id1]
	}
	return "Tasks swaped..", nil
}
func (app *application) purge() string {
	app.tasks.Purge()
	return "All tasks purged.."
}

func (app *application) help() {
	helpMsg := `Usage of %s:
	  -log [options]
			Prints the tasks
			options: -d(detailed), -l(long-term)
	  -add strings.. [options]
			Adds a task
			options: -desc(add description), -time(set time-period), -status(done, pending, overdue, paused)
	  -rm [idxs..]
			Removes the tasks
	  -done [idxs..]
			Marks the tasks completed
	  -undo [idxs..]
			Marks the tasks uncompleted
	  -edit [idx] [options]
			Edit the specified task
			options: -title(new title), -desc(new description), -time(new time-period), -status(done, pending, overdue, paused)
	  -swap [idx_1] [idx_2]
			Swaps the positions of the specified tasks
	  -purge
			Clear completed tasks
	  -clear
			Clear all tasks
`
	var writer io.Writer
	if app.isInteractive {
		writer = app.terminal
	} else {
		writer = os.Stdout
	}
	fmt.Fprintf(writer, helpMsg, os.Args[0])
}
