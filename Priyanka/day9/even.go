package main

import (
	"fmt"
	"sync"
)

func printEven(wg *sync.WaitGroup, ch chan int, n int) {
	defer wg.Done()
	for i := 0; i < n; i++ {
		ch <- 2 * i
	}
}

func printOdd(wg *sync.WaitGroup, ch chan int, n int) {
	defer wg.Done()
	for i := 0; i < n; i++ {
		ch <- 2*i + 1
	}
}

func main() {
	n := 10 // number of even and odd numbers to print

	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)

	go printEven(&wg, ch, n)
	go printOdd(&wg, ch, n)

	go func() {
		defer close(ch)
		wg.Wait()
	}()

	for num := range ch {
		fmt.Println(num)
	}
}
