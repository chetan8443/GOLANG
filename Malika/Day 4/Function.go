// for loop example for table of a number
package main

import 

"fmt"

"Github.com/Malika/Day4/Function1"

func printTable(number int) {
	fmt.Println("Table of", number, "is as following")

	for i := 1; i <= 10; i++ {
		fmt.Println(number, " x ", i, " = ", number*i)
	}
	fmt.Println()
}

func main() {
	myphone(12345)
	printTable(2)
	printTable(8)
}
