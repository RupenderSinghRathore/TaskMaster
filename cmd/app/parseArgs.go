package main

import (
	"RupenderSinghRathore/TaskMaster/internal/models"
	"errors"
	"fmt"
	"os"
	"time"
)

func (app *application) handleArgs() error {
	var err error
	switch {
	case len(app.args) == 0:
		err = app.interactiveShellMode()
	case app.args[0] == "log":
	case app.args[0] == "add":
		err = app.add()
	case app.args[0] == "rm":
		err = app.remove()
	case app.args[0] == "done":
		err = app.done()
	case app.args[0] == "swap":
		err = app.swap()
	case app.args[0] == "purge":
		app.purge()
	case app.args[0] == "clear":
		app.clear()
	case app.args[0] == "edit":
		err = app.edit()
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
		fmt.Fprintf(os.Stderr, "err: %s\n", err.Error())
	}
}
func (app *application) add() error {
	n := len(app.args)
	var err error
	for i := 1; i < n; i++ {
		task := app.tasks.Append(capitalize(app.args[i]))
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
func (app *application) remove() error {
	ids := make(map[int]bool, len(app.args[1:]))
	for _, id := range app.args[1:] {
		i, err := getInt(id, len(app.tasks))
		if err != nil {
			return err
		}
		ids[i-1] = true
	}
	app.tasks.Delete(ids)
	return nil
}
func (app *application) done() error {
	for _, id := range app.args[1:] {
		idx, err := getInt(id, len(app.tasks))
		if err != nil {
			return err
		}
		app.tasks[idx-1].Status = models.Done
	}
	return nil
}
func (app *application) clear() {
	app.tasks = models.Tasks{}
}
func (app *application) edit() error {
	n := len(app.args)
	if n > 2 || n < 8 {
		id, err := getInt(app.args[1], len(app.tasks))
		if err != nil {
			return err
		}
		id--
		for i := 2; i < n; i++ {
			switch app.args[i] {
			case "-desc":
				{
					if i+1 < n {
						app.tasks[id].Description = app.args[i+1]
					} else {
						return errors.New("Not enought args")
					}
					i++
				}
			case "-time":
				{
					if i+1 < n {
						deadline, err := getDeadline(app.args[i+1])
						if err != nil {
							return fmt.Errorf("%s is not a valid time", app.args[i+1])
						}
						app.tasks[id].Deadline = deadline
					} else {
						return errors.New("Not enought args")
					}
					i++

				}
			default:
				app.tasks[id].Title = app.args[i]
			}
		}
	} else {
		return errors.New("Not enought or too many args")
	}
	return nil
}
func (app *application) swap() error {
	// swap 1 2
	if len(app.args) <= 3 {
		tasklen := len(app.tasks)
		id1, err := getInt(app.args[1], tasklen)
		if err != nil {
			return err
		}
		id1--
		id2, err := getInt(app.args[2], tasklen)
		if err != nil {
			return err
		}
		id2--
		fmt.Printf("id1: %d, id2: %d\n", id1, id2)
		app.tasks[id1], app.tasks[id2] = app.tasks[id2], app.tasks[id1]
	}
	return nil
}

func (app *application) purge() {
	app.tasks.Purge()
}
