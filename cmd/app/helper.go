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
		return nil, fmt.Errorf("Wrong format for time in %s", period)
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
				return nil, errors.New("MotherFucker you'd be dead by then.")
			}
			duration *= Day * 365
		}
	}
	deadline := time.Now().Add(time.Duration(duration))
	return &deadline, nil
}

func capitalize(s string) string {
	if len(s) > 0 && s[0] >= 97 && s[0] <= 122 {
		return string(s[0]-32) + s[1:]
	}
	return s
}
func notValidId(id int) error {
	return fmt.Errorf("%d is not a valid task", id)

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
