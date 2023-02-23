// Write a program to add remove and view employees:

package main

import (
	"fmt"
	"os"
)

type emp struct { //making struct of employee and fields
	empId   string
	empName string
	salary  string
}

var list []emp

func main() {
	for {
		fmt.Print("\nSelect option:\n1. Add Employee\n2. Remove Employee\n3. View Employee\n4. Exit\n")
		var opt int
		fmt.Print("selected option :")
		fmt.Scan(&opt) //selecting option from the menu

		switch opt { //run case for respective option selected
		case 1:
			fmt.Println()
			addEmp()
			fmt.Println()
		case 2:
			fmt.Println()
			rmEmp()
			fmt.Println()
		case 3:
			fmt.Println()
			viewEmp()
			fmt.Println()
		case 4:
			os.Exit(1)
		default: // if option selected is not present in menu
			fmt.Println("invalid option!!\nPlease select correct option")
		}
	}
}

func addEmp() { // adding employee to list

	var id, n, s string
	fmt.Print("Enter employee id : ")
	fmt.Scan(&id)
	fmt.Print("Enter employee name : ")
	fmt.Scan(&n)
	fmt.Print("Enter employee salary : ")
	fmt.Scan(&s)

	e := emp{id, n, s}

	list = append(list, e)
	fmt.Println("Employee Added !!")
}

func viewEmp() { // to view details of employee

	if len(list) == 0 { // if no employee present
		fmt.Println("No employee to show !!")
		return

	} else { //
		fmt.Print("Enter employee id : ")
		var id string
		fmt.Scan(&id)

		for i := 0; i < len(list); i++ {
			if id == list[i].empId {
				fmt.Print("\nEmployee Id : ", list[i].empId,
					"\nEmployee Name : ", list[i].empName,
					"\nEmployee salary : ", list[i].salary)
				return
			}
		}
		fmt.Println("Employee with this id not found !!")
	}
}

func rmEmp() { // To remove employee from list

	if len(list) == 0 {
		fmt.Println("No employee to remove !!")
		return

	} else {

		fmt.Print("Enter employee id : ")
		var id string
		fmt.Scan(&id)

		for i := 0; i < len(list); i++ {
			if id == list[i].empId {
				list = append(list[:i], list[i+1:]...)
				fmt.Println("Employee removed successfully !")
				return
			}
		}
		fmt.Println("Employee with this id not found !!")
	}
}
