package main

import "fmt"

//Below code block initializes the value "i" starting the loop from "32" and
//goes till 126 and prints diffrent types of ASCII values
func main() {
	for i := 32; i <= 126; i++ {
		fmt.Printf("%d\t%b\t%o\t%v\t%#x\n", i, i, i, i, i)
	}
}
