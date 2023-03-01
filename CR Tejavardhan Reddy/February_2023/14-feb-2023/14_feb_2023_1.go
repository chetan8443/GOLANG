package main

import "fmt"

func main() {
	var day int
	fmt.Println("Enter the number")
	fmt.Scan(&day)
	switch day {
	case 1, 5, 6:
		fmt.Println("Monday")
	case 2, 7:
		fmt.Println("Tuesday")
	case 3, 9:
		fmt.Println("Wednesday")
	case 4, 8:
		fmt.Println("Thursday")
	}
}
