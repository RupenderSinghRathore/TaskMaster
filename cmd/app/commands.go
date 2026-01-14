package main

import (
	"RupenderSinghRathore/TaskMaster/internal/models"
	"errors"
	"fmt"
	"io"
	"os"
	"time"
)

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
		isShortTerm := time.Until(*task.Deadline) < 24*time.Hour
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
		idx, err := getInt(id, len(app.tasks))
		if err != nil {
			return "", err
		}
		idx--
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
	var err error
	count := 0
	for i := 1; i < n; i++ {
		task := app.tasks.Append(capitalize(app.args[i]))
		count++
		if i+1 < n && app.args[i+1] == "-desc" {
			if i+2 < n {
				task.Description = app.args[i+2]
				i++
			} else {
				return "", errors.New("Empty Description")
			}
			i++
		}
		if i+1 < n && app.args[i+1] == "-time" {
			if i+2 < n {
				task.Deadline, err = getDeadline(app.args[i+2])
				if err != nil {
					return "", err
				}
				i++
			} else {
				return "", errors.New("Empty Time")
			}
			i++
		} else {
			t := time.Now().Add(time.Hour * 24)
			task.Deadline = &t
		}
	}
	var msg string
	if count > 1 {
		msg = "Tasks added to your log.."
	} else {
		msg = "Task added to your log.."
	}
	return msg, nil
}
func (app *application) remove() (string, error) {
	ids := make(map[int]bool, len(app.args[1:]))
	count := 0
	for _, id := range app.args[1:] {
		i, err := getInt(id, len(app.tasks))
		if err != nil {
			return "", err
		}
		ids[i-1] = true
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
		idx, err := getInt(id, len(app.tasks))
		if err != nil {
			return "", err
		}
		app.tasks[idx-1].Status = models.Done
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
	n := len(app.args)
	if n > 2 && n < 8 {
		id, err := getInt(app.args[1], len(app.tasks))
		if err != nil {
			return "", err
		}
		id--
		for i := 2; i < n; i++ {
			switch app.args[i] {
			case "-desc":
				{
					if i+1 < n {
						app.tasks[id].Description = app.args[i+1]
					} else {
						return "", errors.New("Not enought args")
					}
					i++
				}
			case "-time":
				{
					if i+1 < n {
						deadline, err := getDeadline(app.args[i+1])
						if err != nil {
							return "", fmt.Errorf("%s is not a valid time", app.args[i+1])
						}
						app.tasks[id].Deadline = deadline
					} else {
						return "", errors.New("Not enought args")
					}
					i++

				}
			default:
				app.tasks[id].Title = app.args[i]
			}
		}
	} else {
		return "", errors.New("Not enought or too many args")
	}
	return "Task edited..", nil
}
func (app *application) swap() (string, error) {
	// swap 1 2
	if len(app.args) <= 3 {
		tasklen := len(app.tasks)
		id1, err := getInt(app.args[1], tasklen)
		if err != nil {
			return "", err
		}
		id1--
		id2, err := getInt(app.args[2], tasklen)
		if err != nil {
			return "", err
		}
		id2--
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
			options: -desc(add description), -time(time-period)
	  -rm [idxs..]
			Removes the tasks
	  -done [idxs..]
			Marks the tasks completed
	  -undo [idxs..]
			Marks the tasks uncompleted
	  -edit [idx] [new_title] [options]
			Edit the title of the specified task
			options: -desc(add description), -time(time-period)
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
