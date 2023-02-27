package main

import "fmt"

func main() {

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	odds := make([]int, 0, len(numbers))
	for n := range numbers {
		if n%2 != 0 {
			odds = append(odds, n)
		}
	}

	fmt.Println(odds)
}
