package main

import "fmt"

// Nested if-else condition
func ageSwitch(age int) {

	if age >= 18 {
		fmt.Println("You are eligible for Vote.")
		if age > 21 {
			fmt.Println("You are eligible for Marriage")
		}
	} else if age < 18 && age > 12 {
		fmt.Println("You are a Teenager")
	} else {
		fmt.Println("You are a Child")
	}
}

// Multiple Switch Cases
func oddEven(num int) {

	switch num {
	case 2, 4, 6, 8, 10:
		fmt.Printf("%d is Even Number\n", num)
	case 1, 3, 5, 7, 9:
		fmt.Printf("%d is Odd Number\n", num)
	default:
		fmt.Println("The entered value is out of range")
	}
}

// Nested For Loop
func array(arr [3][3]int) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%d\t", arr[i][j])
		}
		fmt.Println()
	}
}
