package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
	add "task1/add_func"
	dob "task1/dateofbirth"
	delete "task1/delete_func"
	display "task1/display_upcoming_birthday"
	sort "task1/sort_func"
	struct_emp "task1/struct1"
	update "task1/update_func"
	view "task1/view_func"
)

func main() {
	file, err := os.OpenFile("logFile.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	logger := log.New(io.MultiWriter(os.Stdout, file), "", log.Ldate|log.Ltime|log.Lshortfile)
	log.SetOutput(file)
	logger.Println("Employee Management System")

	employees := []struct_emp.Employee{
		{ID: 1, FirstName: "Sujatha", LastName: "Iyer", Email: "sujataiyer124@gmail.com", Password: "Password@123", PhoneNo: "2131231223", Role: "admin", Salary: 400000, DateOfBirth: dob.ParseDOB("1999-05-19")},
		{ID: 3, FirstName: "Nikita", LastName: "Patil", Email: "nikipatil123@gmail.com", Password: "niki@124", PhoneNo: "4765289763", Role: "user", Salary: 300000, DateOfBirth: dob.ParseDOB("1999-10-17")},
		{ID: 2, FirstName: "Joey", LastName: "Bin", Email: "joebin167@gmail.com", Password: "goe@8888", PhoneNo: "9394949494", Role: "user", Salary: 300000, DateOfBirth: dob.ParseDOB("1999-02-19")},
		{ID: 4, FirstName: "Shital", LastName: "Patil", Email: "shitalvyas111@gmail.com", Password: "shi123e", PhoneNo: "9994448882", Role: "admin", Salary: 400000, DateOfBirth: dob.ParseDOB("1999-10-20")},
	}

	for {
		logger.Println("Are you user or admin? Type exit to stop")
		reader := bufio.NewReader(os.Stdin)
		role, _ := reader.ReadString('\n')
		role = strings.TrimSpace(role)
		switch role {
		case "admin":
			logger.Println("Welcome admin")
			logger.Println("What task you want to perform")
			logger.Println("1.Add Employee Details")
			logger.Println("2.View Employee Details")
			logger.Println("3.Update Employee Details")
			logger.Println("4.Delete Employee Details")
			action, err := reader.ReadString('\n')
			action = strings.TrimSpace(action)

			if err != nil {
				logger.Println("Invalid action.Select numbers from 1 to 4")
				continue
			}

			switch action {
			case "1":
				//add employee details
				var empdata = add.AddEmployeeDetails(reader, logger)
				employees = append(employees, empdata)
				logger.Println("Addded employee info successfully")

			case "2":
				//View employee details
				logger.Println("View employee details:")
				view.ViewEmployeeDetails(employees,logger)

			case "3":
				//update employee details
				view.ViewEmployeeDetails(employees,logger)
				update.UpdateEmp(employees, reader, logger)
			case "4":
				//delete employee details
				view.ViewEmployeeDetails(employees,logger)
				delete.DeleteEmp(employees, reader,logger)

			case "exit":
				logger.Println("Exiting the program.")
				return

			default:
				logger.Println("Invalid action. Please enter a valid option.")
			}

		case "user":
			logger.Println("Welcome user")
			logger.Println("Display upcoming birthday of colleague")
			display.UpcomingBirthday(employees,logger)
			sort.SortbyField(employees, reader, logger)
			view.ViewEmployeeDetails(employees,logger)
		case "exit":
			return
		default:
			logger.Println("Invalid role")
		}
	}

}
