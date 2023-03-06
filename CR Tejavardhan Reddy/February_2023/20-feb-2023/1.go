package main

import "fmt"

func main() {
	var low int
	var high int
	fmt.Printf("Enter the lower value:")
	fmt.Scanln(&low)
	fmt.Printf("Enter the upper value:")
	fmt.Scanln(&high)
	if low%2 == 0 && high%2 == 0 {
		fmt.Print("The odd count is:", (high-low)/2)
	} else {
		fmt.Print("The odd count is:", ((high-low)/2)+1)
	}
}

//Count of Odd numbers in the given range in mathematical method
//Input:3,9
//Output:4
