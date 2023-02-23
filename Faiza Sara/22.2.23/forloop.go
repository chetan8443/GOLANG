// To check a number is Prime or not.

package main

import "fmt"

func main() {
	var num, count int
	count = 0
	fmt.Println("Enter the number:")
	fmt.Scanln(&num)
	for i := 2; i < num/2; i++ {
		if num%i == 0 {
			count++
		}
	}
	if count == 0 && num != 1 {
		fmt.Println(num, "is a Prime Number")
	} else {
		fmt.Println(num, "is not a Prime Number")
	}
}
