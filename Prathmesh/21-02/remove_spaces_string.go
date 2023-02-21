package main

import (
	"fmt"
	"strings" //importing the package string for the replace func
)

func main() {
	str := "  This is a test example  "
	fmt.Println("Original String: ", str)
	fmt.Println("Output String: ", strings.Replace(str, " ", "", -1))
}
