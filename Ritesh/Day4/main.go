package main

import (
	"fmt"

	"Github.com/Ritesh/Day4/part1"
)

// Below code block is calling diffrent functions with diffrent arguments
func main() {
	name()
	adhaar(1234567891234567)
	a1 := address("Gurgaon")
	fmt.Println(a1)
	phone()
	message := part1.HelloGreeting("Progressive Coder")
	fmt.Println(message)

}

// Below code block declare "a" with my name and runs a loop
// and prints the index value and each letter of my name
func name() {
	a := []string{"r", "i", "t", "e", "s", "h"}
	for i, v := range a {
		fmt.Println(i, v)
	}
}

// Below code block is passing parameter "a" and will print adhaar number when the function"adhaar" is called
func adhaar(a int) {
	fmt.Println("My adhaar number is :", a)
}

// Below function returns the address value from Sprint statement and stores it in above declared variable "a1"
func address(add string) string {
	return fmt.Sprint(" I stay in:", add)
}
