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

func AddEmployeeDetails(reader *bufio.Reader,logger *log.Logger) struct_emp.Employee{

	var employee struct_emp.Employee
	rand.Seed(time.Now().UnixNano())
	id := rand.Intn(1000) + 1
	employee.ID = id

	logger.Println("Enter First Name")
	employee.FirstName, _ = reader.ReadString('\n')
	employee.FirstName = strings.TrimSpace(employee.FirstName)

	logger.Println("Enter Last name")
	employee.LastName, _ = reader.ReadString('\n')
	employee.LastName = strings.TrimSpace(employee.LastName)

	logger.Println("Enter Email id")
	employee.Email, _ = reader.ReadString('\n')
	employee.Email = strings.TrimSpace(employee.Email)

	logger.Println("Enter Password")
	employee.Password, _ = reader.ReadString('\n')
	employee.Password = strings.TrimSpace(employee.Password)

	logger.Println("Enter Phone Number")
	employee.PhoneNo, _ = reader.ReadString('\n')
	employee.PhoneNo = strings.TrimSpace(employee.PhoneNo)

	logger.Println("Enter Role")
	employee.Role, _ = reader.ReadString('\n')
	employee.Role = strings.TrimSpace(employee.Role)

	logger.Println("Enter Salary")
	sal1, _ := reader.ReadString('\n')
	sal1 = strings.TrimSpace(sal1)
	employee.Salary, _ = strconv.ParseFloat(sal1, 64)

	logger.Println("Enter date of birth in YYYY-MM-DD format")
	dobStr, _ := reader.ReadString('\n')
	dobStr = strings.TrimSpace(dobStr)
	dob, err := time.Parse("2006-01-02", dobStr)

	if err != nil {
		panic(err)
	}
    employee.DateOfBirth = dob
	return employee
}