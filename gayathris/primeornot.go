package main

import (
	"fmt"
	"math"
)

func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	num := 7
	if isPrime(num) {
		fmt.Printf("%d is prime\n", num)
	} else {
		fmt.Printf("%d is not prime\n", num)
	}
}