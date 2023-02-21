package main

import "fmt"

//Declaration of variables
var num int
var fl float64
var b bool
var str string

//Main function 
func main(){
	fmt.Println("Default value of num is ",num)
	fmt.Println("Default value of float is ",fl)
	fmt.Println("Default value of bool is ",b)
	fmt.Println("Default value of string is ",str)

	//initialization
	num = 34
	fl  = 4.5677
	b   = true
	str = "Hello"

	fmt.Println("Value of num is ",num)
	fmt.Println("Value of float is ",fl)
	fmt.Println("Value of bool is ",b)
	fmt.Println("Value of string is ",str)

}