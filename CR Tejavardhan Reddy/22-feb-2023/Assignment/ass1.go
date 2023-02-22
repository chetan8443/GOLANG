package main

import "fmt"

func main() {
	var a int
	var b int
	fmt.Print("Enter the a and b value")
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	switch a {
	case 0:
		fmt.Print("0")
	case 2:
		fmt.Print("2")
	case 4:
		fmt.Print("4")
	case 6:
		fmt.Print("6")
	case 8:
		fmt.Print("8")
	default:
		switch b {
		case 1:
			fmt.Print("1")
		case 3:
			fmt.Print("3")
		case 5:
			fmt.Print("5")
		case 7:
			fmt.Print("0")
		case 9:
			fmt.Println("9")
		default:
			fmt.Print("Double Digit")
		}
	}
}
