package main

import "fmt"

func main() {
	var n int = 10
	var sum int = 0
	for i := 1; i < n+1; i = i + 1 {
		sum = sum + i

	}
	fmt.Println(sum)
}
