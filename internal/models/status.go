package models

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

func (s Status) String() string {
	return statusEmum[s]
}
