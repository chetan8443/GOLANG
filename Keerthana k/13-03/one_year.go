package main

import (
	"fmt"
	"time"
)

func main() {
	currentDate := time.Now()                  //getting current date
	futureDate := currentDate.AddDate(1, 0, 2) //adding one year and two days to current date
	formattedDate := futureDate.Format("march 2, 2006")
	fmt.Println("The date one year and two days from today is:", formattedDate) // print result
}
