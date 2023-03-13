// 1)how can you find the date that is exactly one year and two days after the current date?
// 2) Write a program that prints the current time in two different time zones.

package main

import (
	"fmt"
	"time"
)

func main() {
	timeXone() //Calling function for current time of different zones
	additon()  //Calling function for the date that is exactly one year and two days after the current date
}

func timeXone() {

	loc, _ := time.LoadLocation("Asia/Shanghai") //LoadLocation returns the Location with the given name
	loc1, _ := time.LoadLocation("Asia/Kolkata")
	
	now := time.Now().In(loc) //It returns current time of given location
	now1 := time.Now().In(loc1)
	fmt.Println("Current Time in CST Shnaghai : ", now)
	fmt.Println("Current Time in IST Kolkata : ", now1)

}

func additon() {
	// Get the current date
	todaysDate := time.Now()

	// Add one year and two days to the current date
	newDate := todaysDate.AddDate(1, 0, 2)

	// Format the new date as a string
	Date := newDate.Format("2006-01-02")

	// Print the new date
	fmt.Println("The date After one year and two days is:", Date)
}
