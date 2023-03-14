package main

import (
	"fmt"
	"time"
)

func main() {

	location1, err := time.LoadLocation("America/New_York")

	location2, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		panic(err)
	}

	timeInNewYork := time.Now().In(location1)
	timeInIndia := time.Now().In(location2)

	fmt.Println("Time in New York : ", timeInNewYork.Format("01-02-2006 15:04:05 Monday"))
	fmt.Println("Time in India : ", timeInIndia.Format("01-02-2006 15:04:05 Monday"))

}
