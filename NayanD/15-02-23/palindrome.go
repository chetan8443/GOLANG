//Write a program to check whether given number is palindrome or not

package main

import "fmt"

func palindrome() {
	var number, rem, temp int
	reverse := 0

	fmt.Println("Enter a number : ")
	fmt.Scanf("%d", &number)

	temp = number
	for number > 0 {
		rem = number % 10
		reverse = reverse*10 + rem
		number /= 10
		if number == 0 {
			break
		}
	}
	if temp == reverse {
		fmt.Println("%d is a palindrome", temp)
	} else {
		fmt.Println("%d is not a palindrome", temp)
	}
}
