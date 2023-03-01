// Write a program to add, remove and view employees:

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
