package delete

import (
	"bufio"
	"log"
	"strconv"
	"strings"
	"task1/struct1"
)

func DeleteEmp(employees []struct_emp.Employee, reader *bufio.Reader) {
	log.Println("Enter which id you want to delete:")
	id_to_delete, _ := reader.ReadString('\n')
	id_to_delete = strings.TrimSpace(id_to_delete)
	del, _ := strconv.Atoi(id_to_delete)
	id_found := false
	for i, emp := range employees {
		if del == emp.ID {
			id_found = true
			employees = append(employees[:i], employees[i+1:]...)
			log.Println("The employee has been deleted successfully")
		}
	}
	if !id_found {
		log.Println("Enter correct id")
	}
}