package main

import (
	"fmt"
	v "new/demo1"
)

func main() {
	fmt.Println("it is an main function")
	fmt.Println("calling the global variable", One) //calling the global varaible
	demo()                                          // here we are calling other file demo
	v.Test()

}
