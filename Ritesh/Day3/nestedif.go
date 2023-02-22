package main

import "fmt"

//Below code block evaluates diffrent expressions and if the result is "True",
// it prints the statements accordingly
func main() {
	if true {
		fmt.Println(1)
	}
	if false {
		fmt.Println(2)
	}
	if !true {
		fmt.Println(3)
	}
	if !false {
		fmt.Println(4)
	}
	if 2 == 2 {
		fmt.Println(5)
	}
	if 2 != 2 {
		fmt.Println(6)
	}
	if !(2 == 2) {
		fmt.Println(7)
	}
	if !(2 != 2) {
		fmt.Println(8)
	}
	x := 3
	y := 3
	if x == y {
		fmt.Println(9)
	}
	fmt.Println(x)

}
