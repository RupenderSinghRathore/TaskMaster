package main

import (
	"RupenderSinghRathore/TaskMaster/internal/models"
	"encoding/csv"
	"errors"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/google/shlex"
	"golang.org/x/term"
)

const FILE = ".tasks.csv"

func (app *application) interactiveShellMode() error {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	terminal := term.NewTerminal(os.Stdout, "> ")
	app.terminal = terminal
	app.writer = tabwriter.NewWriter(
		terminal,
		0,
		0,
		2,
		' ',
		tabwriter.DiscardEmptyColumns,
	)
	for {
		line, err := terminal.ReadLine()
		if err != nil {
			return err
		}
		if trimedLine := strings.TrimSpace(line); trimedLine != "" {
			if trimedLine == "exit" {
				terminal.Write([]byte("bye..\n"))
				break
			}
			app.args, err = shlex.Split(trimedLine)
			if err != nil {
				terminal.Write(append([]byte(err.Error()), '\n'))
			}
			msg, err := app.handleArgs()
			if err != nil {
				terminal.Write(append([]byte(err.Error()), '\n'))
			}
			terminal.Write(append([]byte(msg), '\n'))
		}
	}
	return nil
}
func loadTasks() (models.Tasks, error) {

	tasks := models.Tasks{}
	f, err := os.Open(FILE)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			os.Create(FILE)
			return tasks, nil
		}
		return nil, err
	}
	defer f.Close()

	records, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return nil, err
	}
	for _, record := range records {
		deadline, err := time.Parse(time.RFC822Z, record[3])
		if err != nil {
			return nil, err
		}
		status, err := strconv.Atoi(record[2])
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, &models.Task{
			Title:       record[0],
			Description: record[1],
			Status:      models.Status(status),
			Deadline:    &deadline,
		})
	}
	return tasks, nil
}

func saveTasks(tasks models.Tasks) error {
	f, err := os.Create(FILE)
	if err != nil {
		return err
	}
	defer f.Close()
	records := make([][]string, 0, len(tasks))
	for _, task := range tasks {
		records = append(records, []string{
			task.Title,
			task.Description,
			strconv.Itoa(int(task.Status)),
			task.Deadline.Format(time.RFC822Z),
		})
	}
	return csv.NewWriter(f).WriteAll(records)
}
