// Switch case with multiple switch case statements
package main

import "fmt"

func getNumberOfDays(month int, year int) {
	var numDays int

	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		numDays = 31
		break
	case 4, 6, 9, 11:
		numDays = 30
		break
	case 2:
		if ((year%4 == 0) && !(year%100 == 0)) || (year%400 == 0) {
			numDays = 29
		} else {
			numDays = 28
		}
		break
	default:
		fmt.Println("Invalid month value: ", month)
		break
	}
	fmt.Println("Number of Days = ", numDays)
}

func main() {
	getNumberOfDays(1, 2023)
	getNumberOfDays(10, 2022)
	getNumberOfDays(2, 2000)
	getNumberOfDays(2, 2024)
}
