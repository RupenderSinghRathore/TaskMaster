package main

import (
	"RupenderSinghRathore/TaskMaster/internal/models"
	"errors"
	"fmt"
	"os"
	"strconv"
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
	var isDone string
	detailed := len(app.args) > 1 && app.args[1] == "-d" ||
		(len(app.args) > 2 && app.args[2] == "-d")
	longTerm := (len(app.args) > 1 && app.args[1] == "-l") ||
		(len(app.args) > 2 && app.args[2] == "-l")
	if len(app.args) > 1 && app.args[1] == "-dl" {
		detailed, longTerm = true, true
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
	for i, task := range app.tasks {
		if longTerm && !(time.Until(*task.Deadline) > 24*time.Hour) {
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
		} else {
			isDone = "❌"
			if task.Status == models.Done {
				isDone = "✔️  ☆*: .｡. o(≧▽≦)o .｡.:*☆"
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
		msg = "Tasks removed from your log.."
	} else {
		msg = "Task removed from your log.."
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

const (
	Day = 24
)

func getDeadline(period string) (*time.Time, error) {
	count, err := strconv.Atoi(period[:len(period)-1])
	if err != nil {
		return nil, fmt.Errorf("Wrong format for time in %s\n", period)
	}
	durationHours := count
	switch period[len(period)-1] {
	case 'd':
		durationHours *= Day
	case 'w':
		durationHours *= Day * 7
	case 'm':
		durationHours *= Day * 30
	case 'y':
		{
			if count > 100 {
				return nil, errors.New("MotherFucker you'd be dead by then.\n")
			}
			durationHours *= Day * 365
		}
	}
	durationNeno := durationHours * int(time.Hour)
	deadline := time.Now().Add(time.Duration(durationNeno))
	return &deadline, nil
}
func getTimeperiod(t *time.Time) string {
	rawDuration := time.Until(*t)
	hours := int(rawDuration.Hours())

	var period string
	switch {
	case hours < 0:
		period = "💀"
	case hours/(Day*365) >= 1:
		period = strconv.Itoa(hours/(Day*365)) + "y"
	case hours/(Day*30) >= 1:
		period = strconv.Itoa(hours/(Day*30)) + "m"
	case hours/(Day*7) >= 1:
		period = strconv.Itoa(hours/(Day*7)) + "w"
	case hours/Day >= 1:
		period = strconv.Itoa(hours/Day) + "d"
	case hours >= 1:
		period = strconv.Itoa(hours) + "h"
	default:
		period = strconv.Itoa(int(rawDuration.Minutes())) + "min"
	}
	return period
}
func capitalize(s string) string {
	if len(s) > 0 && s[0] >= 97 && s[0] <= 122 {
		return string(s[0]-32) + s[1:]
	}
	return s
}
func notValidId(id int) error {
	return fmt.Errorf("%d is not a valid task\n", id)

}
func getInt(s string, tasklen int) (int, error) {
	id, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	} else if id < 1 || id > tasklen {
		return 0, notValidId(id)
	}
	return id, nil
}
func handleErr(err error) {
	fmt.Fprintf(os.Stderr, "err: %v\n", err)
	os.Exit(1)
}
