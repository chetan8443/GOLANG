package main

import (
	"fmt"
)

func main() {
	var num, rem, t int
	var rev int = 0

	fmt.Print("Enter integer : ")
	fmt.Scan(&num)

	t = num
	for {
		rem = num % 10
		rev = rev*10 + rem
		num /= 10

		if num == 0 {
			break
		}
	}

	if t == rev {
		fmt.Printf("%d is a Palindrome", t)
	} else {
		fmt.Printf("%d is not a Palindrome", t)
	}

}
