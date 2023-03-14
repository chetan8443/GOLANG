// how can you find the date that is exactly one year and two days after the current date ?

package main

import (
	"fmt"
	"time"
)

func main() {
	current := time.Now()

	new := current.AddDate(1, 0, 2)

	formattedDate := new.Format("2006-01-02")

	fmt.Println("Date one year and two days after today :", formattedDate)
}