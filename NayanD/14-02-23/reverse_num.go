// write a program  to revverse given number

package main

import "fmt"


func reverseNumber(number int) int {

	result := 0
	for number > 0 {
		remainder := number % 10
		result = (result * 10) + remainder
		number /= 10
	}
	return result
}
