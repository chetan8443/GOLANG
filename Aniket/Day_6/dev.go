package main

import (
	"fmt"
	"math"
)

// Map Iteration
func rangeMap(m1 map[string]string) {
	for country, capital := range m1 {
		fmt.Println(capital, "is a Capital of", country)
	}
}

// Variadic  Function
func varSum(nums ...int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	return sum
}

// Closure Function/ Anonymous Function/ Inline Function
func nextInt() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// Recursion Function
func factorial(num int) int {
	fact := 1
	if num == 1 {
		return 1
	}
	fact = num * factorial(num-1)
	return fact
}

// Error Handling
func sqrt(num float64) float32 {
	if num < 0 {
		fmt.Printf("math: square root of negative number is not allowed\n")
		return 0
	} else {
		sol := math.Sqrt(num)
		return float32(sol)
	}
}
