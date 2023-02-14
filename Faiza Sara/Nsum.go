package main

import (
	"fmt"
)

func main() {
	var a, sum int
	fmt.Scan(&a)
	for i := 1; i <= a; i++ {
		sum += i
	}
	fmt.Print("sum :", sum)
}
