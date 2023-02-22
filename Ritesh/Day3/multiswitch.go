package main

import "fmt"

//Below code block verifies by comparing the string "Ritesh" against diffrent cases;
//where the case matches it prints accordingly, if it doesn't find a match it prints default statement
func main() {
	n := "Ritesh"
	switch n {
	case "Kumar", "mohanty":
		fmt.Println("Not the correct name")
	case "ritesh":
		fmt.Println("Yes thats the name")
	default:
		fmt.Println("None of the names are matching case")
	}
}
