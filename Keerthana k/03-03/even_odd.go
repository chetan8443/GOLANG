package main

import (
	"fmt"
	"sync"
)

func printEvenNumbers(n int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i <= n; i += 2 {
		fmt.Println("Even:", i)
	}
}

func printOddNumbers(n int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= n; i += 2 {
		fmt.Println("Odd:", i)
	}
}

func main() {
	var wg sync.WaitGroup
	n := 10

	wg.Add(2)
	go printEvenNumbers(n, &wg)
	go printOddNumbers(n, &wg)

	wg.Wait()
}
