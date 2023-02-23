// to check the largest among the three numbers.
package main

import "fmt"

func main() {
	var a = 12
	var b = 10
	var c = 14
	if a > b && a > c {
		fmt.Println("a is Greater: ", a)
	} else if b > a && b > c {
		fmt.Println("b is Greater:", b)
	} else {
		fmt.Println("c is Greater:", c)
	}
}
