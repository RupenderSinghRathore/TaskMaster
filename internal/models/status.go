package models

import (
	"fmt"
)

type Status int

const (
	Pending Status = iota
	Done
	Paused
	Overdue
)

var statusEmum = map[Status]string{
	Pending: "Pending",
	Done:    "Done",
	Paused:  "Paused",
	Overdue: "Overdue",
}
var stringToStatus = map[string]Status{
	"pending": Pending,
	"done":    Done,
	"paused":  Paused,
	"overdue": Overdue,
}

func (s Status) String() string {
	return statusEmum[s]
}

func (s *Status) UpdateStatus(newS string) error {
	if status, ok := stringToStatus[newS]; ok {
		*s = status
	} else {
		return fmt.Errorf("%s not a valud status", newS)
	}
	return nil
}
