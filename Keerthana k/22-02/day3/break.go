package main

import "fmt"

func main() {
	a := 10
	for i := a; i <= 100; i += 10 {
		if i == 60 {
			break
		}
		fmt.Println(i)
	}
}
