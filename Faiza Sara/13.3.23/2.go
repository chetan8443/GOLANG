//Write a program that prints the current time in two different time zones.

package main

import (
	"fmt"
	"time"
)

func main() {

	timezone1, err := time.LoadLocation("America/New_York")
	if err != nil {
		panic(err)
	}
	timezone2, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}

	time1 := time.Now().In(timezone1)

	time2 := time.Now().In(timezone2)

	fmt.Printf("Current time in New York: %s\n", time1.Format("2020-01-02 15:04:05 MST"))
	fmt.Printf("Current time in Tokyo: %s\n", time2.Format("2020-01-02 15:04:05 MST"))
}
