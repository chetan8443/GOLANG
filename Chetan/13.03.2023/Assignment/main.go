package main

// 1)how can you find the date that is exactly one year and two days after the current date?
// 2) Write a program that prints the current time in two different time zones.

import (
	"fmt"
	"time"
)

func main() {
	timeXone()
	additon()
}

func timeXone() {

	loc, _ := time.LoadLocation("Asia/Shanghai")
	loc1, _ := time.LoadLocation("Asia/Kolkata")
	//set timezone,
	now := time.Now().In(loc)
	now1 := time.Now().In(loc1)
	fmt.Println(now)
	fmt.Println(now1)

}

func additon() {
	// Get the current date
	currentDate := time.Now()

	// Add one year and two days to the current date
	newDate := currentDate.AddDate(1, 0, 2)

	// Format the new date as a string
	formattedDate := newDate.Format("2006-01-02")

	// Print the new date
	fmt.Println("The date one year and two days after today is:", formattedDate)
}
