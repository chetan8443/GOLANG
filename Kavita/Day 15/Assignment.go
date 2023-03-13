package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	oneYearTwoDays := currentTime.AddDate(1, 0, 2) // years, months, days
	
	fmt.Println("Current date:", (currentTime.Format"2006-01-02"))
	fmt.Println("One year and two days later:", oneYearTwoDays.Format("2006-01-02"))
}
