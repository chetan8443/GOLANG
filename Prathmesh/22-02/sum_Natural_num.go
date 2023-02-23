// Program to print numbers for natural numbers 1 + 2 + 3 + ... +n

package main

import "fmt"

func main() {
	var n int
	fmt.Println("enter the number for the sum of entered natural number")
	fmt.Scanln(&n)
	var sum = 0

	for i := 1; i <= n; i++ {
		sum += i // sum = sum + i
	}

	fmt.Println("sum =", sum)
}
