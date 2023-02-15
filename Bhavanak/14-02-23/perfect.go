// to check given number is perfect sum or not.

package main

import "fmt"

func main() {
	var num, sum int
	sum = 0
	fmt.Println("Enter the number :")
	fmt.Scanln(&num)
	for i := 1; i < num; i++ {
		if num%i == 0 {
			sum = sum + i
		}
	}
	if num == sum {
		fmt.Println(num, "is a Perfect Number")
	} else {
		fmt.Println(num, "is not a Perfect Number")
	}
}
