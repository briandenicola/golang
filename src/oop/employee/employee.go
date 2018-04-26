package employee 

import (
	"fmt"
)

type employee struct {
	firstName   string
	lastName    string
	totalLeaves int
	leavesTaken int
}

func New( firstName string, lastName string, totalLeaves int, leavesTaken int) employee {
	e := employee {firstName, lastName, totalLeaves, leavesTaken}
	return e
}

func ( e employee) LeaveRemaining() {
	fmt.Printf("%s %s has %d leave remaining\n", e.firstName, e.lastName, (e.totalLeaves-e.leavesTaken))
}