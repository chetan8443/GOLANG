// using for loop
package main

import "fmt"

func main() {
	numbers := [6]int{1, 2, 3, 4}
	var b int = 15
	var a int
	for a = 0; a < 10; a++ {
		fmt.Printf("The value of a is %d\n", a)
	}
	for a < b {
		fmt.Printf("The value of a is %d\n", a)
		a++
	}
	for i, x := range numbers {
		fmt.Printf("The value of x is %d\n on %d\n", x, i)
	}
}
