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

func addEmp() Employee {
	var emp Employee
	var name, email string
	var sal float64
	fmt.Println("Enter Employee Name")
	fmt.Scan(&name)
	fmt.Println("Enter the Employee Email id")
	fmt.Scan(&email)
	fmt.Println("Enter the Employee Salary")
	fmt.Scan(&sal)
	emp.empName = name
	emp.empEmail = email
	emp.empSalary = sal
	//fmt.Printf("Employee Details are \n %+v:", emp)
	return emp
}

func viewEmp(emp *Employee) {

	fmt.Println("Name :", emp.empName)
	fmt.Println("Email :", emp.empEmail)
	fmt.Println("Salary :", emp.empSalary)

}
func deleteEmp(emp *Employee) Employee {
	emp.empName = ""
	emp.empEmail = ""
	emp.empSalary = 0.0
	return *emp
}
func main() {
	var emp Employee
	for {

		fmt.Printf("\n1. Add Employee \n2. Delete Employee\n3. View Employee\n4. Exit\n")
		fmt.Println("Select the option")
		var choice int
		fmt.Scan(&choice)

		switch choice {

		case 1:
			emp = addEmp()
			viewEmp(&emp)
		case 2:
			emp = deleteEmp(&emp)
		case 3:
			viewEmp(&emp)
		case 4:
			os.Exit(1)
		}
	}
}
