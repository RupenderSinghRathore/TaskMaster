package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"
)

const ( // no. of hours
	Day   = 24
	Week  = Day * 7
	Month = Week * 4
	Year  = Month * 12
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
		count *= Week
	case 'm':
		count *= Month
	case 'y':
		{
			if count > 100 {
				return time.Time{}, errors.New("MotherFucker you'd be dead by then.\n")
			}
			count *= Year
		}
	}
	durationNeno := count * float64(time.Hour)
	deadline := time.Now().Add(time.Duration(durationNeno)).Round(time.Second)
	return deadline, nil
}

func getTimeperiod(t time.Time) string {
	rawDuration := time.Until(t)
	hours := rawDuration.Hours()

	var period string
	switch {
	case hours < 0:
		period = "💀"
	case hours/Year >= 1:
		period += strconv.FormatFloat(hours/Year, 'f', 1, 64) + "y"
	case hours/Month >= 1:
		period += strconv.FormatFloat(hours/Month, 'f', 1, 64) + "m"
	case hours/Week >= 1:
		period += strconv.FormatFloat(hours/Week, 'f', 1, 64) + "w"
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

func (app *application) getTaskId(s string) (int, error) {
	id, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	} else if id < 1 || id > len(app.tasks) {
		return 0, fmt.Errorf("%d is not a valid task\n", id)
	}
	return id - 1, nil
}

func handleErr(err error) {
	fmt.Fprintf(os.Stderr, "err: %v\n", err)
	// fmt.Fprintf(os.Stdout, "trace: %s\n", string(debug.Stack()))
	os.Exit(1)
}
