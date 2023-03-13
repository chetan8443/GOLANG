// Write a program to find a day on specific date

package main

import (
	"fmt"
	"time"
)

func main() {
	// take date
	date := time.Date(1999, time.May, 15, 0, 0, 0, 0, time.UTC)

	// printing day on given date
	day := date.Format("Monday")
	fmt.Print(day)
}
