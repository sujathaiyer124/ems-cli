package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

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
type Field []Employee

func (a Field) Len() int {
	return len(a)
}
func (a Field) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

// Id
func (a Field) IdToSort(i, j int) bool {
	return a[i].ID < a[j].ID
}

// firstname
func (a Field) FirstNameToSort(i, j int) bool {
	return a[i].FirstName < a[j].FirstName
}

// lastname
func (a Field) LastNameToSort(i, j int) bool {
	return a[i].LastName < a[j].LastName
}

// Email
func (a Field) EmailToSort(i, j int) bool {
	return a[i].Email < a[j].Email
}

// Phoneno
func (a Field) PhoneToSort(i, j int) bool {
	return a[i].PhoneNo < a[j].PhoneNo
}

// Role
func (a Field) RoleToSort(i, j int) bool {
	return a[i].Role < a[j].Role
}

// Salary
func (a Field) SalaryToSort(i, j int) bool {
	return a[i].Salary < a[j].Salary
}

func main() {
	fmt.Println("Employee Management System")

	employees := []Employee{
		{1, "Sujatha", "Iyer", "sujataiyer124@gmail.com", "Password@123", "2131231223", "admin", 400000, parseDOB("1999-05-19")},
		{3, "Nikita", "Patil", "nikipatil123@gmail.com", "niki@124", "4765289763", "user", 300000, parseDOB("1999-10-17")},
		{2, "Joey", "Bin", "joebin167@gmail.com", "goe@8888", "9394949494", "user", 300000, parseDOB("1999-02-19")},
		{4, "Shital", "Patil", "shitalvyas111@gmail.com", "shi123e", "9994448882", "admin", 400000, parseDOB("1999-10-20")},
	}

	for {
		fmt.Println("Are you user or admin? Type exit to stop")
		reader := bufio.NewReader(os.Stdin)
		role, _ := reader.ReadString('\n')
		role = strings.TrimSpace(role)
		switch role {
		case "admin":
			fmt.Println("Welcome admin")
			fmt.Println("What task you want to perform")
			fmt.Println("1.Add Employee Details")
			fmt.Println("2.View Employee Details")
			fmt.Println("3.Update Employee Details")
			fmt.Println("4.Delete Employee Details")
			action, err := reader.ReadString('\n')
			action = strings.TrimSpace(action)

			if err != nil {
				fmt.Println("Invalid action.Select numbers from 1 to 4")
				continue
			}

			switch action {
			case "1":
				//add employee details
				var empdata = AddEmployeeDetails(reader)
				employees = append(employees, empdata)
				fmt.Println("Addded employee info successfully")

			case "2":
				//View employee details
				fmt.Println("View employee details:")
				ViewEmployeeDetails(employees)

			case "3":
				//update employee details
				ViewEmployeeDetails(employees)
				fmt.Println("Enter id which you want to update:")
				id_update, _ := reader.ReadString('\n')
				id_update = strings.TrimSpace(id_update)
				update1, _ := strconv.Atoi(id_update)
				id_found := false
				for i, emp := range employees {
					if update1 == emp.ID {
						id_found = true
						updated_info := AddEmployeeDetails(reader)
						employees[i] = updated_info
						fmt.Println("Updated information of employees are:", employees[i])
					}
				}
				if !id_found {
					fmt.Println("Employee with id ", update1, "not found")
				}

			case "4":
				//delete employee details
				ViewEmployeeDetails(employees)
				fmt.Println("Enter which id you want to delete:")
				id_to_delete, _ := reader.ReadString('\n')
				id_to_delete = strings.TrimSpace(id_to_delete)
				del, _ := strconv.Atoi(id_to_delete)
				id_found := false
				for i, emp := range employees {
					if del == emp.ID {
						id_found = true
						employees = append(employees[:i], employees[i+1:]...)
						fmt.Println("The employee has been deleted successfully")
					}
				}
				if !id_found {
					fmt.Println("Enter correct id")
				}

			case "exit":
				fmt.Println("Exiting the program.")
				return

			default:
				fmt.Println("Invalid action. Please enter a valid option.")
			}

		case "user":
			fmt.Println("Welcome user")
			fmt.Println("Display upcoming birthday of colleague")
			upcomingBirthday(employees)
			fmt.Println("Employee sort according to specific function by \n 1.Id\n 2.FirstName\n 3.LastName \n4.Email \n 5.Phone Number \n 6.Role \n 7 Salary")
			fmt.Println("Choose any option from the above")
			sortField1, _ := reader.ReadString('\n')
			sortField1 = strings.TrimSpace(sortField1)
			sortField, _ := strconv.Atoi(sortField1)

			var sortfunc func(i, j int) bool
			switch sortField {
			case 1:
				sortfunc = Field(employees).IdToSort
			case 2:
				sortfunc = Field(employees).FirstNameToSort
			case 3:
				sortfunc = Field(employees).LastNameToSort
			case 4:
				sortfunc = Field(employees).EmailToSort
			case 5:
				sortfunc = Field(employees).PhoneToSort
			case 6:
				sortfunc = Field(employees).RoleToSort
			case 7:
				sortfunc = Field(employees).RoleToSort

			default:
				fmt.Println("Invalid sorting field.")
				continue
			}
			sort.SliceStable(employees, func(i, j int) bool {
				return sortfunc(i, j)
			})
			ViewEmployeeDetails(employees)

		case "exit":
			return
		default:
			fmt.Println("Invalid role")
		}
	}

}

func AddEmployeeDetails(reader *bufio.Reader) Employee {

	var employee Employee
	rand.Seed(time.Now().UnixNano())
	id := rand.Intn(1000) + 1
	employee.ID = id

	fmt.Println("Enter First Name")
	employee.FirstName, _ = reader.ReadString('\n')
	employee.FirstName = strings.TrimSpace(employee.FirstName)

	fmt.Println("Enter Last name")
	employee.LastName, _ = reader.ReadString('\n')
	employee.LastName = strings.TrimSpace(employee.LastName)

	fmt.Println("Enter Email id")
	employee.Email, _ = reader.ReadString('\n')
	employee.Email = strings.TrimSpace(employee.Email)

	fmt.Println("Enter Password")
	employee.Password, _ = reader.ReadString('\n')
	employee.Password = strings.TrimSpace(employee.Password)

	fmt.Println("Enter Phone Number")
	employee.PhoneNo, _ = reader.ReadString('\n')
	employee.PhoneNo = strings.TrimSpace(employee.PhoneNo)

	fmt.Println("Enter Role")
	employee.Role, _ = reader.ReadString('\n')
	employee.Role = strings.TrimSpace(employee.Role)

	fmt.Println("Enter Salary")
	sal1, _ := reader.ReadString('\n')
	sal1 = strings.TrimSpace(sal1)
	employee.Salary, _ = strconv.ParseFloat(sal1, 64)

	fmt.Println("Enter date of birth in YYYY-MM-DD format")
	dobStr, _ := reader.ReadString('\n')
	dobStr = strings.TrimSpace(dobStr)
	dob, err := time.Parse("2006-01-02", dobStr)

	if err != nil {
		panic(err)
	}
	//dob = dob.In(time.UTC).Truncate(24 * time.Hour)

	employee.DateOfBirth = dob
	return employee
}
func ViewEmployeeDetails(employees []Employee) {

	for _, emp := range employees {
		fmt.Printf("Id is: %d:\n", emp.ID)
		fmt.Println("Firstname is", emp.FirstName)
		fmt.Println("Lastname is", emp.LastName)
		fmt.Println("Email id is", emp.Email)
		fmt.Println("Phone number is", emp.PhoneNo)
		fmt.Println("Role is", emp.Role)
		fmt.Println("Salary is:", emp.Salary)
		fmt.Println("Date of birth is:", emp.DateOfBirth.Format("2006-01-02"))
		fmt.Println()
	}

}
func upcomingBirthday(employees []Employee) {
	fmt.Println("Employeee those who have upcoming birthday are:")
	currentmonth := time.Now().Month()
	currentday := time.Now().Day()
	for _, emp := range employees {
		birthmonth := emp.DateOfBirth.Month()
		birthday := emp.DateOfBirth.Day()
		if birthmonth == currentmonth && birthday > currentday {
			fmt.Printf("%s's upcoming birthday: %s\n", emp.FirstName, emp.DateOfBirth.Format("2006-01-02"))
		}
	}

}

// func to parse dateofbirth
func parseDOB(dobstring string) time.Time {
	dob, err := time.Parse("2006-01-02", dobstring)
	if err != nil {
		panic(err)
	}
	return dob
}
