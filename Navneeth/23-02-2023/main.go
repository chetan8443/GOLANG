package main

import (
	"fmt"
	a "new/pack"
)

var s string

func main() {
	fmt.Println("Enter a Sentence.")
	fmt.Scanf("%s", &s)
	fmt.Printf("Reversed String is : %s\n", a.Rev(s))
	fmt.Printf("Global Variable is: %d\n", a.Global())
}
