

package main

import "fmt"

func employee(name *string, age int) {
	if age > 65 {
		fmt.Println("name : ", *name)
		panic("Age cannot be greater than retirement age :")          //using panic method for generating error if age is more than 75
	}

}

func main() {
	empName := "SWARAJ"
	age := 75

	employee(&empName, age)           //passing address of 'empname' variable and value of 'age' variable in employee function

}
