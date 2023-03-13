//Write a program that prints the current time in two different time zones.

package main

import (
	"fmt"
	"time"
)

func main() {
	// get the current time in the local time zone
	localTime := time.Now()

	// print the local time
	fmt.Println("Local Time:", localTime)

	// get the time zone for New York and Shanghai
	nyTimeZone, err := time.LoadLocation("America/New_York")
	if err != nil {
		fmt.Println(err)
		return
	}
	chinaTimeZone, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}
	// convert the current time to New York time and Shanghai time
	nyTime := localTime.In(nyTimeZone)
	chinaTime := localTime.In(chinaTimeZone)

	// print the New York time
	fmt.Println("New York Time:", nyTime)
	fmt.Println("Shanghai Time:", chinaTime)
}
