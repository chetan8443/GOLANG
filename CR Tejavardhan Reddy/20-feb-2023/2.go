//go code to check whether a number is perfect square or not
//Input:4
//Output:true

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
