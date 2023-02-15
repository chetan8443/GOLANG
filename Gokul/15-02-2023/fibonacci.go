//Program to print the fibonacci series based on number of terms given by user
package main

import "fmt"

func main() {
	var num int
	fmt.Println("Enter the no. of terms in the fibonacci sequence")
	fmt.Scan(&num)
	fmt.Print("Fibonacci series is : \n")
	fibonacci(num)
}

func fibonacci(n int){
	var a,b,c int = 0,0,1
	for i:=1 ;i<=n;i++ {
		fmt.Print(b," ")
		a = b + c
		b = c
		c = a
	}
}