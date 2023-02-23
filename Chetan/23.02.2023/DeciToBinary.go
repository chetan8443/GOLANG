package main

import "fmt"

func main() {

	var a [10]int
	var i, n int
	fmt.Println("Enter number to convert into binary : ")
	fmt.Scan(&n)
//for loop execution for assining binary
	for i = 0; n > 0; i++ {
		a[i] = n % 2
		n = n / 2

	}
	fmt.Print("binary of given number is : ")
	for i := i-1; i>=0; i-- {                         //for loop for priniting binary
		fmt.Print(a[i])
	}
}
