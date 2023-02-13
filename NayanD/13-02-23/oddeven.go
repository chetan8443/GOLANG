// write a program to find whether number is even odd

package main

import "fmt"

func main() {
	fmt.Print("Enter number : ")
	var a int
	fmt.Scanln(&a)
	if a%2 == 0 {
		fmt.Println(a, "is Even ")
	} else {
		fmt.Println(a, "is Odd ")
	}
}
