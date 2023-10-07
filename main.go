package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	add "task1/add_func"
	dob "task1/dateofbirth"
	delete "task1/delete_func"
	display "task1/display_upcoming_birthday"
	struct_emp "task1/struct1"
	update "task1/update_func"
	view "task1/view_func"
)

func main() {
	file, err := os.OpenFile("logFile.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY,0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	log.SetOutput(file) 
	log.Println("Employee Management System")

	employees := []struct_emp.Employee{
		{ID: 1, FirstName: "Sujatha", LastName: "Iyer", Email: "sujataiyer124@gmail.com", Password: "Password@123", PhoneNo: "2131231223", Role: "admin", Salary: 400000, DateOfBirth: dob.ParseDOB("1999-05-19")},
		{ID: 3, FirstName: "Nikita", LastName: "Patil", Email: "nikipatil123@gmail.com", Password: "niki@124", PhoneNo: "4765289763", Role: "user", Salary: 300000, DateOfBirth: dob.ParseDOB("1999-10-17")},
		{ID: 2, FirstName: "Joey", LastName: "Bin", Email: "joebin167@gmail.com", Password: "goe@8888", PhoneNo: "9394949494", Role: "user", Salary: 300000, DateOfBirth: dob.ParseDOB("1999-02-19")},
		{ID: 4, FirstName: "Shital", LastName: "Patil", Email: "shitalvyas111@gmail.com", Password: "shi123e", PhoneNo: "9994448882", Role: "admin", Salary: 400000, DateOfBirth: dob.ParseDOB("1999-10-20")},
	}

	for {
		log.Println("Are you user or admin? Type exit to stop")
		reader := bufio.NewReader(os.Stdin)
		role, _ := reader.ReadString('\n')
		role = strings.TrimSpace(role)
		switch role {
		case "admin":
			log.Println("Welcome admin")
			log.Println("What task you want to perform")
			log.Println("1.Add Employee Details")
			log.Println("2.View Employee Details")
			log.Println("3.Update Employee Details")
			log.Println("4.Delete Employee Details")
			action, err := reader.ReadString('\n')
			action = strings.TrimSpace(action)

			if err != nil {
				log.Println("Invalid action.Select numbers from 1 to 4")
				continue
			}

			switch action {
			case "1":
				//add employee details
				var empdata = add.AddEmployeeDetails(reader)
				employees = append(employees, empdata)
				log.Println("Addded employee info successfully")

			case "2":
				//View employee details
				log.Println("View employee details:")
				view.ViewEmployeeDetails(employees)

			case "3":
				//update employee details
				view.ViewEmployeeDetails(employees)
				update.UpdateEmp(employees, reader)
			case "4":
				//delete employee details
				view.ViewEmployeeDetails(employees)
				delete.DeleteEmp(employees, reader)

			case "exit":
				log.Println("Exiting the program.")
				return

			default:
				log.Println("Invalid action. Please enter a valid option.")
			}

		case "user":
			log.Println("Welcome user")
			log.Println("Display upcoming birthday of colleague")
			display.UpcomingBirthday(employees)
			view.ViewEmployeeDetails(employees)
		case "exit":
			return
		default:
			log.Println("Invalid role")
		}
	}

}
