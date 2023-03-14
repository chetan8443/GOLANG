// Standard Library Assignment
//how can you find the date that is exactly one year and two days after the current date?

package main

import (
	"fmt"
	"time"
)

func main() {
	currentDate := time.Now()
	fmt.Println(currentDate)
	newDate := currentDate.AddDate(1, 0, 2)
	fmt.Println(newDate)
	fmt.Println("The date one year and two days after the current date is:", newDate)
}
