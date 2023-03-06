package main

import (
	"fmt"
	"sync"
)

func printEvenNumbers(n int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			fmt.Printf("%d ", i)
		}
	}
}

func printOddNumbers(n int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < n; i++ {
		if i%2 != 0 {
			fmt.Printf("%d ", i)
		}
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	n := 10

	go printEvenNumbers(n, &wg)
	go printOddNumbers(n, &wg)

	wg.Wait()
}
