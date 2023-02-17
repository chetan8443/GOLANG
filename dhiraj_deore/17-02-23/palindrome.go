// Write a code to check palindrome number

package main

import "fmt"

func main() {
	fmt.Print("Enter a number : ")
	var num int
	fmt.Scanln(&num)
	n1 := num
	n := 0

	for num > 0 {
		d := num % 10
		n = n*10 + d
		num /= 10
	}

	if n1 == n {
		fmt.Println(n, "is a palindrome number")
	} else {
		fmt.Println(n, "is not a palindrome number")
	}
}
