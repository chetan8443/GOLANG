package main

import (
	"fmt"
	"sync"
)

func printEvenNumbers(n int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", i*2)
	}
}

func printOddNumbers(n int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", i*2+1)
	}
}

func main() {
	var n int
	fmt.Println("Enter the value of n:")
	fmt.Scanln(&n)

	var wg sync.WaitGroup
	wg.Add(2)

	go printEvenNumbers(n, &wg)

	go printOddNumbers(n, &wg)

	wg.Wait()
}
