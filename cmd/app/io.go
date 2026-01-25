package main

import (
	"RupenderSinghRathore/TaskMaster/internal/models"
	"encoding/csv"
	"errors"
	"os"
	"strconv"
	"time"
)

const FILE = "/home/kami-sama/.tasks.csv"

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
			Deadline:    deadline,
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
