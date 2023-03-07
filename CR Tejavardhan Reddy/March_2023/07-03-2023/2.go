package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var arr = make([]int, 0)
	var Range int
	fmt.Print("Enter the range:")
	fmt.Scanln(&Range)
	fmt.Print("Enter the count of numbers:")
	var n int
	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		arr = append(arr, rand.Intn(Range))
	}
	fmt.Println("The numbers are:", arr)
	fmt.Println("The perfect numbers are", isPerfect(arr))
}

func isPerfect(arr []int) []int {
	var res = make([]int, 0)
	sum := 0
	for _, num := range arr {
		sum = 0
		for i := 1; i < num; i++ {
			if num%i == 0 {
				sum = sum + i
			}
		}
		if num == sum {
			res = append(res, num)
		}

	}
	var res1 = make([]int, 0)
	if len(res) != 0 {
		return res
	} else {
		return res1
	}
}
