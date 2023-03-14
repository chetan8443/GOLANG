package main

import (
	"fmt"
	"time"
)

func main() {
	// Get the current time in local time zone
	localTime := time.Now()

	// Get the time in New York time zone
	nyTimeZone, _ := time.LoadLocation("America/New_York")
	nyTime := localTime.In(nyTimeZone)

	// Get the time in Tokyo time zone
	tokyoTimeZone, _ := time.LoadLocation("Asia/Tokyo")
	tokyoTime := localTime.In(tokyoTimeZone)

	// Print the times
	fmt.Println("Local time:", localTime.Format("15:04:05"))
	fmt.Println("New York time:", nyTime.Format("15:04:05"))
	fmt.Println("Tokyo time:", tokyoTime.Format("15:04:05"))
}
