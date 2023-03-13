// Write a program that prints the current time in two different time zones.

package main

import (
	"fmt"
	"time"
)

func main() {

	// Get the current date
	currenttime := time.Now()
	// Print the time in IST time zone
	fmt.Println(currenttime)

	location1, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
	}

	location2, err := time.LoadLocation("America/Chicago")
	if err != nil {
		fmt.Println(err)
	}

	shanghai := time.Now().In(location1)
	Chicago := time.Now().In(location2)

	fmt.Println("Asia/Shanghai current Time :- ", shanghai)
	fmt.Println("America/Chicago current Time :- ", Chicago)

}
