package main

import (
	"fmt"
)

func main() {
	var a, t, d int
	var s int = 0
	fmt.Print("Enter number")
	fmt.Scan(&a)
	t = a
	for {
		d = a % 10
		s = s + d*d*d
		a = a / 10
		if a == 0 {
			break
		}
	}
	if s == t {
		fmt.Printf("%d is an Armstrong number", t)
	} else {
		fmt.Printf("%d is not an Armstrong number", t)
	}

}
