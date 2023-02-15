// Write a program to check whether given year is leap year or not
package main

import "fmt"

func leap_year() {

	var year int

	fmt.Println("Enter the year :")
	fmt.Scanf("%d", &year)

	if year%4 == 0 && year%100 != 0 || year%400 == 0 {
		fmt.Println("This is a leap year")
	} else {
		fmt.Println("This is not a leap year")
	}
}
