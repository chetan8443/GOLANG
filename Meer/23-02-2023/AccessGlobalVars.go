// program to access a global variables in other golang files
package main

import (
	"fmt"
	a "golbalVars/test"
)

func main() {
	a.Add(2, 3)
	fmt.Println(a.GlobalVaribale)
}
