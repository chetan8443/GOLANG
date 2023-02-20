package main

import "fmt"

var sum int

func main() {
	for i := 0; i < 100; i++ {
		sum = sum + i
	}
	fmt.Println(sum)
}
