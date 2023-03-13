// program to print ‘n’ even and odd numbers using goroutines.
package main

import (
	"fmt"
)

func main() {
	var num = []int{5, 4, 8, 12, 18, 13, 25, 23}
	Odd := make(chan int)
	Even := make(chan int)

	go odd(Odd)
	go even(Even)

	for _, value := range num {
		if value%2 != 0 {
			Odd <- value
		} else {
			Even <- value
		}
	}

}

func odd(ch <-chan int) {
	for v := range ch {
		fmt.Println("ODD :", v)
	}
}

func even(ch <-chan int) {
	for v := range ch {
		fmt.Println("EVEN:", v)
	}
}
