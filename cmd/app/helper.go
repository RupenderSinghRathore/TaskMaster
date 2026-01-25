package main

import (
	"RupenderSinghRathore/TaskMaster/internal/models"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"
)

const (
	Day = 24
)

func getDeadline(period string) (time.Time, error) {
	count, err := strconv.ParseFloat(period[:len(period)-1], 64)
	if err != nil {
		return time.Time{}, fmt.Errorf("Wrong format for time in %s\n", period)
	}
	switch period[len(period)-1] {
	case 'd':
		count *= Day
	case 'w':
		count *= Day * 7
	case 'm':
		count *= Day * 30
	case 'y':
		{
			if count > 100 {
				return time.Time{}, errors.New("MotherFucker you'd be dead by then.\n")
			}
			count *= Day * 365
		}
	}
	durationNeno := count * float64(time.Hour)
	deadline := time.Now().Add(time.Duration(durationNeno))
	return deadline, nil
}
func getTimeperiod(t time.Time) string {
	rawDuration := time.Until(t)
	hours := rawDuration.Hours()

	var period string
	switch {
	case hours < 0:
		period = "💀"
	case hours/(Day*365) >= 1:
		period += strconv.FormatFloat(hours/(Day*365), 'f', 1, 64) + "y"
	case hours/(Day*30) >= 1:
		period += strconv.FormatFloat(hours/(Day*30), 'f', 1, 64) + "m"
	case hours/(Day*7) >= 1:
		period += strconv.FormatFloat(hours/(Day*7), 'f', 1, 64) + "w"
	case hours/Day >= 1:
		period += strconv.FormatFloat(hours/Day, 'f', 1, 64) + "d"
	case hours >= 1:
		period += strconv.FormatFloat(hours, 'f', 1, 64) + "h"
	default:
		period += strconv.FormatFloat(rawDuration.Minutes(), 'f', 0, 64) + "min"
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
	// fmt.Fprintf(os.Stdout, "trace: %s\n", string(debug.Stack()))
	os.Exit(1)
}

func insertionSort(tasks models.Tasks) {
	var  j int
	for i, task := range tasks {
		curr := time.Until(task.Deadline)
		j = i-1
		for j >= 0 && time.Until(tasks[j].Deadline) > curr {
			tasks[j+1] = tasks[j]
			j--
		}
		tasks[j+1] = task
	}
}
