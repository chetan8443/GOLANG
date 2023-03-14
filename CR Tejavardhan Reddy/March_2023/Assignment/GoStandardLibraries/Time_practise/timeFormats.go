package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	fmt.Println(currentTime)
	formatedDate := currentTime.Format("14-Mar-2023")
	fmt.Println(formatedDate)
	// format the future date as a string
	//formattedDate := futureDate.Format("02-Jan-2006")
}
