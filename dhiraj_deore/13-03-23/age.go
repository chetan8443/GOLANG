// write a program to find age

package main

import (
	"fmt"
	"time"
)

func main() {
	// specify the birth date as a string
	birthDateString := "1999-05-15"

	// parse the birth date string into a time.Time value
	birthDate, err := time.Parse("2006-01-02", birthDateString)
	if err != nil {
		fmt.Println(err)
		return
	}

	// calculate the duration between the birth date and the current date
	duration := time.Since(birthDate)

	// extract the number of years from the duration
	years := duration / (365 * 24 * time.Hour)

	// print the age in years
	fmt.Printf("Age: %d years\n", years)
}
