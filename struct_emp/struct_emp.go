package struct_emp

import "time"

type Employee struct {
	ID          int
	FirstName   string
	LastName    string
	Email       string
	Password    string
	PhoneNo     string
	Role        string
	Salary      float64
	DateOfBirth time.Time
}
