/* check the given number is palindrome r not
package main

import "fmt"

func main() {
	var num, remainder int
	fmt.Println("Enter the number:")
	fmt.Scanln(&num)
	reverse := 0
	for temp := num; temp > 0; temp = temp / 10 {
		remainder = temp % 10
		reverse = reverse*10 + remainder
	}
	fmt.Println("The reverse of a number=", reverse)
	if num == reverse {
		fmt.Println(num, "is a palindrome")
	} else {
		fmt.Println(num, "is not a palindrome")
	}
} */

// check the given number is palindrome r not using function
package main

import "fmt"

var reverse int = 0

func revNumber(palNum int) int {
	var remainder int

	for ; palNum > 0; palNum = palNum / 10 {
		remainder = palNum % 10
		reverse = reverse*10 + remainder
	}
	return reverse
}
func main() {

	var palNum int

	fmt.Print("Enter the Number = ")
	fmt.Scanln(&palNum)

	reverse = revNumber(palNum)
	fmt.Println("The Reverse of the Given Number = ", reverse)

	if palNum == reverse {
		fmt.Println(palNum, " is a Palindrome Number")
	} else {
		fmt.Println(palNum, " is Not a Palindrome Number")
	}
}
