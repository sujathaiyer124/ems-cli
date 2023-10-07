package update

import (
	"bufio"
	"log"
	"strconv"
	"strings"
	add "task1/add_func"
	struct_emp "task1/struct_emp"
)

func UpdateEmp(employees []struct_emp.Employee, reader *bufio.Reader, logger *log.Logger) {
	logger.Println("Enter id which you want to update:")
	id_update, _ := reader.ReadString('\n')
	id_update = strings.TrimSpace(id_update)
	update1, _ := strconv.Atoi(id_update)
	id_found := false
	for i, emp := range employees {
		if update1 == emp.ID {
			id_found = true
			updated_info := add.AddEmployeeDetails(reader, logger)
			employees[i] = updated_info
			logger.Println("Updated information of employees are:", employees[i])
		}
	}
	if !id_found {
		logger.Println("Employee with id ", update1, "not found")
	}
}
