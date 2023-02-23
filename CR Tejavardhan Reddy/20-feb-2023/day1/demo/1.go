// go code to check whether the number is even or odd
package main

import "fmt"

func main() {
	fmt.Println("Enter a number:")
	var n int
	fmt.Scanln(&n)
	oddNumber(n)
	evenNumber(n)

}

func evenNumber(n int) {
	if n%2 == 0 {
		fmt.Println("Even number")
	}
}

func oddNumber(n int) {
	if n%2 != 0 {
		fmt.Println("Odd Number")
	}
}

//function as closures
