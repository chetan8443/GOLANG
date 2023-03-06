package main

import (
	"fmt"
	"time"
)

func main() {
	var n int
	fmt.Println("Enter n value.")
	fmt.Scanf("%d", &n)
	fmt.Println("odd numbers: ")
	go oddnum(n)
	time.Sleep(time.Millisecond * 100)
	fmt.Println("even numbers: ")
	go evennum(n)
	time.Sleep(time.Millisecond * 100)

}

func oddnum(n int) {
	for i := 0; i < n; i++ {
		fmt.Println(i*2 + 1)
	}
}

func evennum(n int) {
	for i := 0; i < n; i++ {
		fmt.Println(i * 2)
	}
}
