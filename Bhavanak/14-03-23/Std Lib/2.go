// Standard Library assignment
//Write a program that prints the current time in two different time zones.

package main

import (
	"fmt"
	"time"
)

func main() {
	timezone1, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		panic(err)
	}

	timezone2, err := time.LoadLocation("America/New_York")
	if err != nil {
		panic(err)
	}

	time1 := time.Now().In(timezone1)
	time2 := time.Now().In(timezone2)

	fmt.Println("Time in Kolkata : ", time1)
	fmt.Println("Time in New York : ", time2)

}
