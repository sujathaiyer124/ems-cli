package view

import (
	"log"
	"task1/struct_emp"
)

func ViewEmployeeDetails(employees []struct_emp.Employee,logger *log.Logger) {

	for _, emp := range employees {
		logger.Printf("Id is: %d:\n", emp.ID)
		logger.Println("Firstname is", emp.FirstName)
		logger.Println("Lastname is", emp.LastName)
		logger.Println("Email id is", emp.Email)
		logger.Println("Phone number is", emp.PhoneNo)
		logger.Println("Role is", emp.Role)
		logger.Println("Salary is:", emp.Salary)
		logger.Println("Date of birth is:", emp.DateOfBirth.Format("2006-01-02"))
		logger.Println()
	}

}