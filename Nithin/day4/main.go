package main

import (
	"fmt"
	a "new/pack"
)

var n int = 10
var sum int

func main() {
	fmt.Println("Global Variable:", a.Num)
	fmt.Println("addition:", a.Add(10, 7))

	for i := 1; i <= n; i++ {
		sum += i // sum = sum + i
	}
	fmt.Println("sum =", sum)
}
