// Printing the natural numbers and addding the sum of natural numbers.
package main

import "fmt"

var a int
var sum int
var i int

//main function
func main() {
	fmt.Printf("enter the number")
	fmt.Scanln(&a)
	fmt.Printf("the natural numbers are\n")
	for i = 1; i <= a; i++ {
		fmt.Printf("%d\n", i)
		sum += i
	}
	fmt.Printf("the sum of natural numbers are %d", sum)

}
