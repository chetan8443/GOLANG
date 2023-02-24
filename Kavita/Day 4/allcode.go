package main // package name

import "fmt"

func SubNums(a int, b int) int { // function written to substract two numbers
	sub := a - b
	return sub

}

func main() {
	//out := SubNums(10, 5)
	//fmt.Println(out)
	fmt.Println(SubNums(20, 10)) // calling and printing the function written in the same file in the  main fucntion

}
