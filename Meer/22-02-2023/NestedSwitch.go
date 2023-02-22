package main

import "fmt"

func main() {
	var choice int

	fmt.Println("1. Food\n2.Beverages")
	fmt.Println("Select the Choice")
	fmt.Scan(&choice)
	switch choice {

	case 1:
		fmt.Println("1.North Indian Food\n2.South Indian Food")
		fmt.Println("Enter the food choice")
		var foodChoice int
		fmt.Scan(&foodChoice)
		switch foodChoice {
		case 1:
			fmt.Println("North Indian Food")
		case 2:
			fmt.Println("South Indian Food")
		default:
			fmt.Println("None Selected")
		}
	case 2:
		fmt.Println("1.Coffee\n2.Juice")
		var beverageChoice int
		fmt.Scan(&beverageChoice)
		switch beverageChoice {
		case 1:
			fmt.Println("Software Engineers Need more coffee")
		case 2:
			fmt.Println("Juice is good for health")
		default:
			fmt.Println("None seleted")
		}
	}
}
