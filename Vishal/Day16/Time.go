package main

import (
	"fmt"
	"time"
)

func main() {

	location1, err := time.LoadLocation("Australia/Perth")

	location2, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		panic(err)
	}

	perth := time.Now().In(location1)
	India := time.Now().In(location2)

	fmt.Println("Time in Aus Perth : ", perth.Format("01-02-2006 15:04:05 Monday"))
	fmt.Println("Time in India : ", India.Format("01-02-2006 15:04:05 Monday"))

}
