// how can you find the date that is exactly one year and two days after the current date ?

package main

import (
	"fmt"
	"time"
)

func main() {
	// Get the current date
	currentDate := time.Now()

	// Add one year and two days to the current date
	newDate := currentDate.AddDate(1, 0, 2)

	// Format the new date as a string
	formattedDate := newDate.Format("2006-01-02")

	// Print the new date
	fmt.Println("The date one year and two days after today is:", formattedDate)
}
