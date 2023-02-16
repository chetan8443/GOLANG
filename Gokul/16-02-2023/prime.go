//Program to check if a number is prime or composite
package main

import "fmt"

func main() {
	var num int
	var ch byte
	fmt.Println("Enter number")
	fmt.Scan(&num)
	for i :=2; i<=num/2; i++{
		if(num % i ==0){
			ch = 'c'
			break
		}
	}
	if ch == 'c' {
		fmt.Print(num," is composite ")
	}else{
		fmt.Print(num," is prime")
	}
}