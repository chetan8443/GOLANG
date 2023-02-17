//write a code to reverse a given number

package main

import "fmt"

func main() {
	fmt.Print("Enter a number : ")
	var num int
	fmt.Scanln(&num)
	n := 0

	for num > 0 {
		d := num % 10
		n = n*10 + d
		num /= 10
	}
	fmt.Println("reverse of the  given number is :", n)
}
