package main

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

const (
	Day = int(time.Hour) * 24
)

func getDeadline(period string) (*time.Time, error) {
	count, err := strconv.Atoi(period[:len(period)-1])
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Wrong format for time in %s", period))
	}
	duration := count
	switch period[len(period)-1] {
	case 'h':
		duration *= int(time.Hour)
	case 'd':
		duration *= Day
	case 'w':
		duration *= Day * 7
	case 'm':
		duration *= Day * 30
	case 'y':
		{
			if count > 100 {
				return nil, errors.New(fmt.Sprintf("MotherFucker you'd be dead by then."))
			}
			duration *= Day * 365
		}
	}
	deadline := time.Now().Add(time.Duration(duration))
	return &deadline, nil
}
