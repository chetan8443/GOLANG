package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Print("Enter a year")
	fmt.Scan(&a)
	if a%4 == 0 || (a%400 == 0 && a%100 == 0) {
		fmt.Println("It's a leap year")
	} else {
		fmt.Println("It's not a leap year")
	}
}
