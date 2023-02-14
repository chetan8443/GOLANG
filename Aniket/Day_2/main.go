package main

import "fmt"

func main() {
	fmt.Println("Enter your number: ")
	var num1 int
	fmt.Scanln(&num1)
	primeNo(num1)

	var num2 int
	fmt.Println("Enter your number: ")
	fmt.Scanln(&num2)
	pali(num2)

	var arr = [5]int{1, 2, 1, 9, 5}
	avg := avg(arr)
	fmt.Println("The average of the array is", avg)
}
