package display

import (
	"log"
	struct_emp "task1/struct_emp"
	"time"
)

func UpcomingBirthday(employees []struct_emp.Employee,logger *log.Logger) {
	logger.Println("Employeee those who have upcoming birthday are:")
	currentmonth := time.Now().Month()
	currentday := time.Now().Day()
	for _, emp := range employees {
		birthmonth := emp.DateOfBirth.Month()
		birthday := emp.DateOfBirth.Day()
		if birthmonth == currentmonth && birthday > currentday {
			logger.Printf("%s's upcoming birthday: %s\n", emp.FirstName, emp.DateOfBirth.Format("2006-01-02"))
		}
	}

}
