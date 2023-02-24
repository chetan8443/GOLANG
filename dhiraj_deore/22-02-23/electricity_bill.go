//Write a program for electricity bill generation with nested if else:

package main

import "fmt"

func main() {
	fmt.Print("\nEnter units : ") // taking total unit of electricity consumed by user
	var unit int
	fmt.Scan(&unit)

	fmt.Print("\nselect user type : \n1.Domestic\n2. Industrial\n") 
	var user int // taking input for type of user 
	fmt.Scan(&user)
	fmt.Println(" \n ")

	var bill int
	if user == 1 { // condition for different types of users 

		if unit > 750 {
			bill = (unit-750)*15 + 250*10 + 500*7
			fmt.Print("Total bill in Rs : ", bill)

		} else if unit > 500 {
			bill = (unit-500)*10 + 500*7
			fmt.Print("Total bill in Rs : ", bill)

		} else {
			bill = unit * 7
			fmt.Print("Total bill in Rs : ", bill)

		}

	} else if user == 2 {

		if unit > 750 {
			bill = (unit-750)*25 + 250*15 + 500*10
			fmt.Print("Total bill in Rs : ", bill)

		} else if unit > 500 {
			bill = (unit-500)*15 + 500*10
			fmt.Print("Total bill in Rs : ", bill)

		} else {
			bill = unit * 10
			fmt.Print("Total bill in Rs : ", bill)

		}

	} else {

		fmt.Print("invalid user number !!")

	}
	
	fmt.Println(" \n ")

}
