package main

import (
	"fmt"
	"time"
)

func main() {
	// Define the time zones you want to print
	timezone1, err := time.LoadLocation("America/New_York")
	if err != nil {
		panic(err)
	}
	timezone2, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}

	// Get the current time in the first time zone
	time1 := time.Now().In(timezone1)

	// Get the current time in the second time zone
	time2 := time.Now().In(timezone2)

	// Print the current time in both time zones
	fmt.Printf("Current time in New York: %s\n", time1.Format("2020-01-02 15:04:05 MST"))
	fmt.Printf("Current time in Tokyo: %s\n", time2.Format("2020-01-02 15:04:05 MST"))
}
