// Go program to create maps of structs and slice of struct.
package main

import "fmt"

//create a struct named employee
type employee struct {
	empName string
	empId   int
	age     int
}

//creating a slice of struct
func main() {
	employees := make([]employee, 5)

	//creating struct instance
	employees[1] = employee{empName: "Bhavana", empId: 101, age: 21}
	employees[2] = employee{empName: "Hemanth", empId: 102, age: 24}
	employees[3] = employee{empName: "Pavan", empId: 103, age: 30}
	employees[4] = employee{empName: "Sudheer", empId: 104, age: 31}

	for _, e := range employees {
		fmt.Println(e)
	}

	// creating a map of struct
	sample := map[employee]int{employees[1]: 1, employees[2]: 2, employees[3]: 3, employees[4]: 4}

	for str, val := range sample {
		fmt.Println(str, val)
	}

}
