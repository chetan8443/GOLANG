// program to calculate how many tallest candles are there in an array
package main

import (
	"fmt"
	"sort"
)

func birthdayCakeCandles(candles []int) int {
	// Write your code here
	n := len(candles)
	count := 0
	sort.Ints(candles)
	for i, _ := range candles {
		if candles[n-1] == candles[i] {
			count++
		}
	}
	return count

}
func main() {
	var candles = []int{4, 4, 1, 3}
	count := birthdayCakeCandles(candles)
	fmt.Println(count)
	var candles1 = []int{3, 2, 3, 3}
	count = birthdayCakeCandles(candles1)
	fmt.Println(count)
}
