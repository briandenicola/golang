package main 

import(
	"oop/employee"
)

func main() {
	e := employee.New("Brian", "Den", 30, 20)
	e.LeaveRemaining()
}