package main

import (
	"fmt"
)

func main() {
	var grade int
	fmt.Println("Enter the numeric grade: ")
	fmt.Scanln(&grade)

	var letterGrade string
	switch {
	case grade >= 90:
		letterGrade = "A"
	case grade >= 80:
		letterGrade = "B"
	case grade >= 70:
		letterGrade = "C"
	case grade >= 60:
		letterGrade = "D"
	default:
		letterGrade = "F"
	}

	fmt.Printf("Letter Grade: %s\n", letterGrade)
}
