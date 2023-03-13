//how can you find the date that is exactly one year and two days after the current date?

package main

import (
	"fmt"
	"time"
)

func main() {
	// get the current date
	currentDate := time.Now()

	// add one year and two days to the current date
	newDate := currentDate.AddDate(1, 0, 2)

	// format the date as "2006-01-02"
	date := newDate.Format("2006-01-02")

	// print the formatted date
	fmt.Println(date)
}

//
