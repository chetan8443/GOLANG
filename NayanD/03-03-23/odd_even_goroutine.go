// Write a program to print ‘n’ even and odd numbers using goroutines.
package main

import (
	"fmt"
	"sync"
)

func evenNumbers(n int, g *sync.WaitGroup) {
	defer g.Done()
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			fmt.Printf("Even: %d\n", i)
		}
	}
}

func oddNumbers(n int, g *sync.WaitGroup) {
	defer g.Done()
	for i := 0; i < n; i++ {
		if i%2 != 0 {
			fmt.Printf("Odd: %d\n", i)
		}
	}
}

func main() {
	n := 11
	var g sync.WaitGroup
	g.Add(2)
	go evenNumbers(n, &g)
	go oddNumbers(n, &g)
	g.Wait()
}
