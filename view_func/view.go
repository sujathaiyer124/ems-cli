package view

import (
	"log"
	"task1/struct1"
)

func ViewEmployeeDetails(employees []struct_emp.Employee) {

	for _, emp := range employees {
		log.Printf("Id is: %d:\n", emp.ID)
		log.Println("Firstname is", emp.FirstName)
		log.Println("Lastname is", emp.LastName)
		log.Println("Email id is", emp.Email)
		log.Println("Phone number is", emp.PhoneNo)
		log.Println("Role is", emp.Role)
		log.Println("Salary is:", emp.Salary)
		log.Println("Date of birth is:", emp.DateOfBirth.Format("2006-01-02"))
		log.Println()
	}

}