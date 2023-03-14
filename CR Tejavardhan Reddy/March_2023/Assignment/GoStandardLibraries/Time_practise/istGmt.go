package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	fmt.Println()
	//Two time Zones
	fmt.Println("The Local time is:", currentTime)
	fmt.Println("The Greenwich Mean time is:", currentTime.UTC())
	fmt.Println()
	Time()
}
func Time() {

	//Different formats of time
	t := time.Date(2023, 3, 14, 9, 40, 0, 0, time.UTC)
	fmt.Println("1)", t)

	fmt.Println("2)", t.Format("14-03-2023 9:42"))
	fmt.Println("3)", t.Format("3.05 14 Feb"))
}
