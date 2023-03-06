// program to print ‘n’ even and odd numbers using goroutines.
package main

import (
	"fmt"
)

func main1() {
	var intSlice = []int{53, 40, 87, 92, 69, 223, 145, 88, 44}
	chOdd := make(chan int)
	chEven := make(chan int)

	go odd(chOdd)
	go even(chEven)

	for _, value := range intSlice {
		if value%2 != 0 {
			chOdd <- value
		} else {
			chEven <- value
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
