//Program to print even and odd numbers using go routine

package main

import (
	"fmt"
)

func printEvenNumbers(ch chan int) {
	for i := 0; i <= 10; i += 2 {
		ch <- i
	}
	close(ch)
}

func printOddNumbers(ch chan int) {
	for i := 1; i <= 10; i += 2 {
		ch <- i
	}
	close(ch)
}

func main() {
	evenCh := make(chan int)
	oddCh := make(chan int)

	go printEvenNumbers(evenCh)
	go printOddNumbers(oddCh)

	for {
		select {
		case even, ok := <-evenCh:
			if !ok {
				evenCh = nil
				break
			}
			fmt.Println("Even:", even)
		case odd, ok := <-oddCh:
			if !ok {
				oddCh = nil
				break
			}
			fmt.Println("Odd:", odd)
		}
		if evenCh == nil && oddCh == nil {
			break
		}
	}
}
