package main

import (
	"fmt"
	logic "teja/day1"
)

func main() {
	fmt.Print("Enter the number:")
	var n float64
	fmt.Scanln(&n)
	if n > 0 {
		fmt.Print(logic.PerfectSquare(n))
	}
}
