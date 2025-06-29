package main

import (
	"bufio"
	"os"
	"strings"

	"github.com/RupenderSinghRathore/TaskMaster/internal/models"
)

func (m Model) WriteToFile() error {
	f, err := os.OpenFile("/home/kami-sama/.task.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	f.Close()
	f, err = os.OpenFile("/home/kami-sama/.task.txt", os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	for _, task := range m.tasks {
		if task.Done {
			continue
		}
		_, err := f.Write([]byte(task.Title + "\n"))
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *Model) ReadFromFile() error {
	f, err := os.Open("/home/kami-sama/.task.txt")
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line != "" {
			m.tasks = append(m.tasks, models.Task{Title: line})
		}
	}
	if err = scanner.Err(); err != nil {
		return err
	}
	return nil
}
