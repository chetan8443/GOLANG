//Program to check whether a number is odd or even
package main

import "fmt"

func main() {
	var num int
	fmt.Print("Enter a number")
	fmt.Scan(&num)
	if(num %2 ==0){
		fmt.Print(num," is even")
	}else{
		fmt.Print(num," is odd")
	}
}