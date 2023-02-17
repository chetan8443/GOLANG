//write a Program to check Armstrong Number.

package main

import "fmt"

func main() {

	var num int
	var temp int
	var remainder int
	var result int = 0

	fmt.Println("Enter a three digit number : ")
	fmt.Scanln(&num)

	temp = num

	for {
		remainder = temp % 10
		result += remainder * remainder * remainder
		temp /= 10

		if temp == 0 {
			break
		}
	}

	if result == num {
		fmt.Printf("%d  Armstrong number.", num)
	} else {
		fmt.Printf("%d is not an Armstrong number.", num)
	}
}
