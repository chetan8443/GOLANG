// program to print ‘n’ even and odd numbers using goroutines.

package main

import "fmt"

func main() {
	var sli = []int{22, 45, 12, 19, 94, 82}

	Odd := make(chan int)
	Even := make(chan int)

	go odd(Odd)
	go even(Even)

	for _, val := range sli {
		if val%2 != 0 {
			Odd <- val
		} else {
			Even <- val
		}
	}

}

func odd(ch <-chan int) {
	for i := range ch {
		fmt.Println(i, "is Odd")
	}
}

func even(ch <-chan int) {
	for i := range ch {
		fmt.Println(i, "is Even")
	}
}