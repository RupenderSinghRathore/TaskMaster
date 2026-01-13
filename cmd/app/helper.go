package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"
)

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
