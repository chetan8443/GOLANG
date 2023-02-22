package main

import (
	"fmt"
	"os"
)

type Employee struct {
	empName   string
	empEmail  string
	empSalary float64
}

func addEmp() ([]Employee, int) {
	var numberOfEmployees int
	fmt.Println("Enter the No of Employees to be Added")
	fmt.Scan(&numberOfEmployees)
	emp := make([]Employee, numberOfEmployees)
	var name, email string
	var sal float64
	for i := 0; i < numberOfEmployees; i++ {
		fmt.Println("Enter", i+1, " Employee Name")
		fmt.Scan(&name)
		fmt.Println("Enter ", i+1, " Employee Email id")
		fmt.Scan(&email)
		fmt.Println("Enter ", i+1, " Employee Salary")
		fmt.Scan(&sal)
		emp[i].empName = name
		emp[i].empEmail = email
		emp[i].empSalary = sal
	}
	//fmt.Printf("Employee Details are \n %+v:", emp)
	return emp, numberOfEmployees
}

func viewEmp(emp *[]Employee, numberOfEmployees int) {

	if numberOfEmployees != 0 {
		fmt.Println("Employee Details are :")
		for i := 0; i < numberOfEmployees; i++ {
			fmt.Println("Name :", (*emp)[i].empName)
			fmt.Println("Email :", (*emp)[i].empEmail)
			fmt.Println("Salary :", (*emp)[i].empSalary)
			fmt.Println("-----------------------------")
		}
	} else {
		fmt.Println("No records to show!!!")
	}
}
func deleteEmp(emp *[]Employee, numberOfEmployees int) ([]Employee, int) {

	for i := 0; i < numberOfEmployees; i++ {
		(*emp)[i].empName = ""
		(*emp)[i].empEmail = ""
		(*emp)[i].empSalary = 0.0

	}
	numberOfEmployees = 0
	fmt.Println("Employees Deleted Successfully")

	return (*emp), numberOfEmployees
}
func main() {
	emp := make([]Employee, 10)
	var numberOfEmployees int = 0
	for {

		fmt.Printf("\n1. Add Employee \n2. Delete Employees\n3. View Employee\n4. Exit\n")
		fmt.Println("Select the option")
		var choice int

		fmt.Scan(&choice)

		switch choice {

		case 1:
			emp, numberOfEmployees = addEmp()

		case 2:
			emp, numberOfEmployees = deleteEmp(&emp, numberOfEmployees)
		case 3:
			viewEmp(&emp, numberOfEmployees)
		case 4:
			os.Exit(1)
		}
	}
}
