package main

import (
	"fmt"
	"time"
)

func main() {
	// get the current date
	today := time.Now()

	// add one year and two days to the current date
	futureDate := today.AddDate(1, 0, 2)
	fmt.Println("One year and two days from today is:", futureDate) //formattedDate)
}
