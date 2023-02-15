package main

import (
	"fmt"
)

func greatest_num() {
	var number1, number2, number3, largest int32
	fmt.Println("Program to find the largest number among three numbers within the function")
	number1 = 50
	number2 = 75
	number3 = 45

	if number1 >= number2 && number1 >= number3 {
		largest = number1
	} else if number2 >= number1 && number2 >= number3 {
		largest = number2
	} else {
		largest = number3
	}

	fmt.Println("Number 1 =", number1, "\nNumber 2=", number2, "\nNumber 3 =", number3, "\nLargest Number =", largest)
}
