// write a program to swap two number

package main

import "fmt"

func swap() {
	var num1, num2 int
	num1 = 10
	num2 = 20
	fmt.Println("Numbers before swap: \n Num1 =", num1, "\n Num2 =", num2)
	num1, num2 = num2, num1
	fmt.Println("Numbers after swap:\n Num1 =", num1, "\n Num2 =", num2, "\n(Swap using a one-liner syntax)")
}

func reverseNumber(num int) int {

	res := 0
	for num > 0 {
		remainder := num % 10
		res = (res * 10) + remainder
		num /= 10
	}
	return res
}
