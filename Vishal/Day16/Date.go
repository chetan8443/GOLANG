package main

import (
	"fmt"
	"time"
)

func main() {
	date := time.Now()
	fmt.Println("Current Date :\n", date.Format("01-02-2006 15:04:05 Monday"))

	newDate := date.AddDate(1, 0, 2)

	fmt.Println("After  1 Years and 2 Days :\n", newDate.Format("01-02-2006 15:04:05 Monday"))
}
