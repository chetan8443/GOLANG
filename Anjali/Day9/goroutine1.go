package main

import (
	"fmt"
	"sync"
)

func printNumbers(ch chan int, wg *sync.WaitGroup, isEven bool, n int) {
	defer wg.Done()
	num := 0
	if isEven {
		num = 2
	} else {
		num = 1
	}
	for i := 0; i < n; i++ {
		ch <- num
		num += 2
	}
}

func main() {
	n := 10
	ch := make(chan int)
	var wg sync.WaitGroup

	wg.Add(2)
	go printNumbers(ch, &wg, true, n)  // Print n even numbers
	go printNumbers(ch, &wg, false, n) // Print n odd numbers

	go func() {
		wg.Wait()
		close(ch)
	}()

	for num := range ch {
		fmt.Println(num)
	}
}
