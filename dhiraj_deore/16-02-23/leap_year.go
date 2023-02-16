// Write a code to find given year is leap year or not

package main

import "fmt"

func main() {

	fmt.Print("Enter a year : ")
	var year int
	fmt.Scanln(&year)

	if year%4 == 0 {
		if year%100 == 0 && year%400 != 0 {
			fmt.Println("Not a leaf year")
		} else {
			fmt.Println("Leaf year")
		}

	} else {
		fmt.Println("Not a leaf year")
	}
}
