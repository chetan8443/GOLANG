package main

import "fmt"

func main() {
     a := 25
	for i := a; i < 200; i += 25 { // for loop with break statement used
		if i == 75 {
			break
		}
		fmt.Println(i)
	}
}
