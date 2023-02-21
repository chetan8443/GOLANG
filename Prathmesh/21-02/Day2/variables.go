package main

import "fmt"

var num int = 20           // int variable
var amount float32 = 49.99 // float variable
var isValid bool = true    // bool variable
var name string = "Yash"   // string variable

func main() {

	num2 := 20 // DECLARING DIRECTLY WITH SHORTHAND

	fmt.Printf("variable num is of type %T \n", num)
	fmt.Printf("variable num is of type %T \n", amount)
	fmt.Printf("variable num is of type %T \n", isValid)
	fmt.Printf("variable num is of type %T \n", name)
	fmt.Println("value for num2 is", num2)

}
