// 2)Write a program for multiple switch statement.
package main

import (
	"fmt"
)

func main() {

	// string to input month from user
	var day string
	fmt.Println("Enter the Day from week for the update")
	fmt.Scanln(&day)

	// switch case for predicting
	// office timing for weekly entered

	// each switch case has more
	// than one values
	switch day {
	case "monday", "tuesday":
		fmt.Println("office timing between 9am to 6pm")
	case "wednesday", "thursday":
		fmt.Println("office timing between 10am to 7pm")
	case "friday":
		fmt.Println("you assigned work from home")
	case "saturday", "sunday":
		fmt.Println("..........!!!!!Its weekend ENJOY...........!!!!")
	}
}
