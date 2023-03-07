package main

import "fmt"

func main() {
	a := 28
	sum := 0
	for i := 1; i < a; i++ {
		if a%i == 0 {
			sum = sum + i
		}
	}
	if sum == a {
		fmt.Print("equal")
	} else {
		fmt.Print("Not equal")
	}
}
