package update

import (
	"bufio"
	"log"
	"strconv"
	"strings"
	 "task1/add_func"
     "task1/struct1"
)

func UpdateEmp(employees []struct_emp.Employee, reader *bufio.Reader) {
	log.Println("Enter id which you want to update:")
	id_update, _ := reader.ReadString('\n')
	id_update = strings.TrimSpace(id_update)
	update1, _ := strconv.Atoi(id_update)
	id_found := false
	for i, emp := range employees {
		if update1 == emp.ID {
			id_found = true
			updated_info := add.AddEmployeeDetails(reader)
			employees[i] = updated_info
			log.Println("Updated information of employees are:", employees[i])
		}
	}
	if !id_found {
		log.Println("Employee with id ", update1, "not found")
	}
}