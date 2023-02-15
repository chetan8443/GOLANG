package main

import "fmt"

func main() {
	var num int
	fact := 1

	fmt.Println("Enter the number for factorial")
	fmt.Scan(&num)
	for i := 1; i <= num; i++ {
		fact *= i
	}

	fmt.Printf("Factorial of %v is %v\n", num, fact)
}
