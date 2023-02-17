//Write a code to print armstrong number under given limit

package main

import (
	"fmt"
)

func main() {

	fmt.Print("Enter the limit  : ")
	var num int
	fmt.Scanln(&num)

	fmt.Println("The armstrong numbers between 0 to", num, "are :")
	for i := 1; i < num; i++ {
		n := i
		s := 0
		for n > 0 {
			d := n % 10
			s = s + d*d*d
			n /= 10
		}
		if s == i {
			fmt.Println(i)
		}
	}
}
