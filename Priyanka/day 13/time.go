package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	// Get the time in the "America/New_York" time zone
	loc1, err := time.LoadLocation("America/New_York")
	if err != nil {
		panic(err)
	}
	timeInNewYork := now.In(loc1)

	// Get the time in the "Asia/Tokyo" time zone
	loc2, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}
	timeInTokyo := now.In(loc2)

	// Print the times
	fmt.Println("Time in New York:", timeInNewYork.Format("15:04:05 MST"))
	fmt.Println("Time in Tokyo:", timeInTokyo.Format("15:04:05 MST"))
}
