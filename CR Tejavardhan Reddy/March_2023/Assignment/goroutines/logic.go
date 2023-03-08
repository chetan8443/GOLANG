package main

import (
	"fmt"
	"strings"
	"sync"
)

// Write a program to print ‘n’ even and odd numbers using goroutines.

func OddEven(wg *sync.WaitGroup) {
	defer wg.Done()
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	var odd []int
	var even []int
	for _, num := range arr {
		if num%2 == 0 {
			even = append(even, num)
		} else {
			odd = append(odd, num)
		}
	}
	fmt.Println("The Even elements are:", even)
	fmt.Println("The Odd elements are:", odd)
}

// Write a program to print the number of occurrences of a character from a word using Goroutines.

func countAlpha(wg *sync.WaitGroup) {
	defer wg.Done()
	var str = "Tejavardhan"
	var str_lower = strings.ToLower(str)

	for i := 0; i < len(str_lower); i++ {
		count := strings.Count(str_lower, str_lower[i:i+1])
		fmt.Printf("%s occurred %d times\n", str_lower[i:i+1], count)
	}
}
