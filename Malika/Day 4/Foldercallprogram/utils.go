package main

import "fmt"

func TablesOfNumber(number int) {
	fmt.Println("Here is the table of: ", number)

	for i := 1; i <= 10; i++ {
		fmt.Println(number, " X ", i, "=", number*i)
	}
	fmt.Println()
}

func printUserNumber(number int) {
	fmt.Println("My Phone number is ", number)
}
