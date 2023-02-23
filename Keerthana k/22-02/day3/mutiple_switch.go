package main

import (
	"fmt"
)

func main() {

	// string to input month from user
	var month string
	fmt.Scanln(&month)

	switch month {
	case "january", "december":
		fmt.Println("Winter.")
	case "february", "march":
		fmt.Println("Spring.")
	case "april", "may", "june":
		fmt.Println("Summer.")
	case "july", "august":
		fmt.Println("Monsoon.")
	case "september", "november":
		fmt.Println("Autumn.")
	}
}
