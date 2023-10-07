package sort

import (
	"bufio"
	"log"
	"sort"
	"strconv"
	"strings"
	"task1/struct_emp"
)
type Field []struct_emp.Employee

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
func SortbyField(employees []struct_emp.Employee, reader *bufio.Reader,logger *log.Logger){
logger.Println("Employee sort according to specific function by \n 1.Id\n 2.FirstName\n 3.LastName \n4.Email \n 5.Phone Number \n 6.Role \n 7 Salary")
			logger.Println("Choose any option from the above")
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
				logger.Println("Invalid sorting field.")
				
			}
			sort.SliceStable(employees, func(i, j int) bool {
				return sortfunc(i, j)
			})
		}