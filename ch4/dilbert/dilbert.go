package main

import "time"

type Employee struct {
	ID           int
	Name, Adress string
	DoB          time.Time
	Position     string
	Salary       int
	ManagerID    int
}

var dilbert Employee

func init() {
	dilbert.Salary -= 5000 // demoted, for writing too few lines of code
	position := &dilbert.Position
	*position = "Senior" + *position // promoted, for outsourcing Elbonia
}

var employeeOfTheMonth *Employee

func init() {
	employeeOfTheMonth = &dilbert
	employeeOfTheMonth.Position += " (proactive team player)"
}

// func EmployeeByID(id int) *Employee { /* ... */ }
// fmt.Println(EmployeeByID(dilbert.ManagerID).Position) // "Pointy-haired boss"
// id := dilbert.ID
// EmployeeByID(id).Salary = 0 // fired for... no real reason

//poor dilbert
