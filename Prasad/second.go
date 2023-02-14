package main

import "fmt"

func main() {
	var c int = 0
	var a int = 12
	var i int = 0
	for i = 1; i <= a; i++ {
		if a%i == 0 {
			c++
		}
	}
	if c == 2 {
		fmt.Println("prime number")
	} else {
		fmt.Println("not prime number")
	}

}
