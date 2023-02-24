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

func addEmp(emp *[]Employee) ([]Employee, int) {
	var numberOfEmployees int
	var name, email string
	var sal float64

	fmt.Println("Enter Employee Name")
	fmt.Scan(&name)
	fmt.Println("Enter Employee Email id")
	fmt.Scan(&email)
	fmt.Println("Enter Employee Salary")
	fmt.Scan(&sal)

	e := Employee{
		name,
		email,
		sal,
	}

	(*emp) = append((*emp), e)
	numberOfEmployees = len(*emp)
	return *emp, numberOfEmployees
}

func viewEmp(emp *[]Employee, numberOfEmployees int) {

	if numberOfEmployees != 0 {
		fmt.Println("Enter the Employee Email ID :")
		var eid string
		fmt.Scan(&eid)
		for i := 0; i < numberOfEmployees; i++ {
			if eid == (*emp)[i].empEmail {
				fmt.Println("Name :", (*emp)[i].empName)
				fmt.Println("Email :", (*emp)[i].empEmail)
				fmt.Println("Salary :", (*emp)[i].empSalary)
				fmt.Println("-----------------------------")
				return
			}
		}
		fmt.Println("No Employees Found With this Email ID")
	} else {
		fmt.Println("No records to show!!!")
	}
}
func deleteEmp(emp *[]Employee, numberOfEmployees int) ([]Employee, int) {

	if numberOfEmployees != 0 {
		fmt.Println("Enter Employee Email ID :")
		var eid string
		fmt.Scan(&eid)
		for i := 0; i < numberOfEmployees; i++ {
			if eid == (*emp)[i].empEmail {
				(*emp) = append((*emp)[:i], (*emp)[i+1:]...)
				numberOfEmployees--
				fmt.Println("Employee Deleted Successfully")
				return (*emp), numberOfEmployees
			}
		}
		fmt.Println("No Employees Found With this Email ID")
	} else {
		fmt.Println("No Employees are there to delete")
	}
	return (*emp), numberOfEmployees
}
func main() {
	var emp []Employee
	var numberOfEmployees int
	for {

		fmt.Printf("\n1. Add Employee \n2. Delete Employee\n3. View Employee\n4. Exit\n")
		fmt.Println("Select the option")
		var choice int

		fmt.Scan(&choice)

		switch choice {

		case 1:
			emp, numberOfEmployees = addEmp(&emp)

		case 2:
			emp, numberOfEmployees = deleteEmp(&emp, numberOfEmployees)
		case 3:
			viewEmp(&emp, numberOfEmployees)
		case 4:
			os.Exit(0)
		}
	}
}
