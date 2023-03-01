package main

import "fmt"

func printNumbers(n int) {
	for i := 1; i <= n; i++ {
		fmt.Println(i)
	}
}

func main() {
	fmt.Println("Printing numbers from 1 to 5:")
	printNumbers(5)
}
