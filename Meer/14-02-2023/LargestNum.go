package main

import "fmt"

func main() {

	var num1, num2, num3 int

	fmt.Println("Enter the first number")
	fmt.Scan(&num1)

	fmt.Println("Enter the second number")
	fmt.Scan(&num2)

	fmt.Println("Enter the Third number")
	fmt.Scan(&num3)

	if num1 > num2 && num1 > num3 {
		fmt.Println(num1, "is greater number")
	} else if num2 > num3 {
		fmt.Println(num2, "is greater number")
	} else {
		fmt.Println(num3, "is greater number")
	}
}
