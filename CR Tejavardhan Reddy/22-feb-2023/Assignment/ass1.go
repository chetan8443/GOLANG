// The assignment on nested switch
// The action will be formed based on the values assigned to a and b
// Input:Enter the a and b value
// a=1-->goes in default case and checks for the b value
// b=1-->here b is "1" so correspond output will be displayed
//
//output:1
package main

import "fmt"

func main() {
	var a int
	var b int
	fmt.Print("Enter the a value:")
	fmt.Scanln(&a)
	fmt.Print("Enter the b value:")
	fmt.Scanln(&b)
	switch a {
	case 0:
		fmt.Print("Zero")
	case 2:
		fmt.Print("two")
	case 4:
		fmt.Print("Four")
	case 6:
		fmt.Print("Six")
	case 8:
		fmt.Print("Eight")
	default:
		switch b {
		case 1:
			fmt.Print("One")
		case 3:
			fmt.Print("Three")
		case 5:
			fmt.Print("Five")
		case 7:
			fmt.Print("Seven")
		case 9:
			fmt.Println("Nine")
		default:
			fmt.Print("Wrong cases entered")
		}
	}
}
