package add

import (
	"bufio"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"task1/struct1"
	"time"
)

func AddEmployeeDetails(reader *bufio.Reader) struct_emp.Employee{

	var employee struct_emp.Employee
	rand.Seed(time.Now().UnixNano())
	id := rand.Intn(1000) + 1
	employee.ID = id

	log.Println("Enter First Name")
	employee.FirstName, _ = reader.ReadString('\n')
	employee.FirstName = strings.TrimSpace(employee.FirstName)

	log.Println("Enter Last name")
	employee.LastName, _ = reader.ReadString('\n')
	employee.LastName = strings.TrimSpace(employee.LastName)

	log.Println("Enter Email id")
	employee.Email, _ = reader.ReadString('\n')
	employee.Email = strings.TrimSpace(employee.Email)

	log.Println("Enter Password")
	employee.Password, _ = reader.ReadString('\n')
	employee.Password = strings.TrimSpace(employee.Password)

	log.Println("Enter Phone Number")
	employee.PhoneNo, _ = reader.ReadString('\n')
	employee.PhoneNo = strings.TrimSpace(employee.PhoneNo)

	log.Println("Enter Role")
	employee.Role, _ = reader.ReadString('\n')
	employee.Role = strings.TrimSpace(employee.Role)

	log.Println("Enter Salary")
	sal1, _ := reader.ReadString('\n')
	sal1 = strings.TrimSpace(sal1)
	employee.Salary, _ = strconv.ParseFloat(sal1, 64)

	log.Println("Enter date of birth in YYYY-MM-DD format")
	dobStr, _ := reader.ReadString('\n')
	dobStr = strings.TrimSpace(dobStr)
	dob, err := time.Parse("2006-01-02", dobStr)

	if err != nil {
		panic(err)
	}
    employee.DateOfBirth = dob
	return employee
}