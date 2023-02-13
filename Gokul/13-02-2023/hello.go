package main

import "fmt"

func main() {
	fmt.Println("Enter 2 numbers")
	var a,b int
	fmt.Scan(&a,&b)
	var c int = a+b
	fmt.Printf("Sum is %d",c)
}