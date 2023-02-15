package main

import "fmt"

func main() {
	var a int = 10
	for a < 20 {
		if a == 15 {
			// skip the iteration
			a++
			continue
		}
		fmt.Printf("The value of a is %d\n", a)
		a++
	}
}
