// for loop example for table of a number
package main

import "fmt"

func printTable(number int) {
	fmt.Println("Table of", number, "is as following")

	for i := 1; i <= 10; i++ {
		fmt.Println(number, " x ", i, " = ", number*i)
	}
	fmt.Println()
}

func main() {
	printTable(25)
}
