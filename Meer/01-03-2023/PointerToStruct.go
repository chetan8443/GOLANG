// program to create a pointer to a struct and accessing its fields
package main

import "fmt"

// making a structure
type Employee struct {
	empId      int
	empName    string
	empEmailId string
	age        int
}

func main() {

	// creating the instance of the Employee struct type
	emp := Employee{1, "Meer", "meer@gmail.com", 23}

	// // Here, it is the pointer to the struct
	// ptr := &emp

	// fmt.Println(ptr)

	// fmt.Println(ptr.empName)

	// fmt.Print((*ptr).empName, " ", (*ptr).empEmailId, " ", (*ptr).age)

	var ptr1 *Employee = &emp
	fmt.Println()
	fmt.Println(ptr1)
	fmt.Println(ptr1.empName, " ", ptr1.empEmailId, " ", ptr1.age)
	ptr1.empName = "syed"
	//fmt.Println((*ptr1).empName, " ", (*ptr1).empEmailId, " ", (*ptr1).age)

}
