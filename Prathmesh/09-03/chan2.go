package main

import (
	"fmt"
	"time"
)

func sendNumbers(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- i
		time.Sleep(time.Millisecond * 500)
	}
	close(ch)
}

func printNumbers(ch chan int) {
	for n := range ch {
		fmt.Printf("%d ", n)
	}
}

func main() {
	ch := make(chan int)

	go sendNumbers(ch)
	go printNumbers(ch)

	time.Sleep(time.Second * 3)
	fmt.Println("Done!")
}
