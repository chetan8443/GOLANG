package main

import "fmt"

func addEmp() { // adding employee to list

	var id, n, s string
	fmt.Print("Enter employee id : ")
	fmt.Scan(&id)
	fmt.Print("Enter employee name : ")
	fmt.Scan(&n)
	fmt.Print("Enter employee salary : ")
	fmt.Scan(&s)

	e := emp{id, n, s} // add given input tofields of emp struct

	list = append(list, e)
	fmt.Println("Employee Added !!")
}

func viewEmp() { // to view details of employee

	if len(list) == 0 { // if no employee present
		fmt.Println("No employee to show !!")
		return

	} else { //take employee id and view its details
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
		fmt.Println("Employee with this id not found !!") // if given id not found in list
	}
}

func rmEmp() { // to remove employee by taking employee id

	if len(list) == 0 {
		fmt.Println("No employee to remove !!")
		return

	} else {

		fmt.Print("Enter employee id : ")
		var id string
		fmt.Scan(&id)

		for i := 0; i < len(list); i++ {
			if id == list[i].empId {
				list = append(list[:i], list[i+1:]...) // removing given index from list
				fmt.Println("Employee removed successfully !")
				return
			}
		}
		fmt.Println("Employee with this id not found !!") // if given id not found in list
	}
}
