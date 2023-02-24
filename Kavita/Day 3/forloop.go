package main

import "fmt"

func main() {
	a := 5
	for i := a; i < 100; i += 5 { // for loop with break statement used
		if i == 50 {
			break
		}
		fmt.Println(i)
	}
}
