package main

import "fmt"

func main() {
	var score int = 70
	var grade string

	switch {
	case score >= 90:
		grade = "A"
	case score >= 80:
		grade = "B"
	case score >= 70:
		grade = "C"
	case score >= 60:
		grade = "D"
	default:
		grade = "F"
	}

	fmt.Printf("Score: %d\nGrade: %s\n", score, grade)

	var dayOfWeek int = 3
	var day string

	switch dayOfWeek {
	case 1:
		day = "Monday"
	case 2:
		day = "Tuesday"
	case 3:
		day = "Wednesday"
	case 4:
		day = "Thursday"
	case 5:
		day = "Friday"
	case 6:
		day = "Saturday"
	case 7:
		day = "Sunday"
	default:
		day = "Unknown"
	}

	fmt.Printf("Day of week: %d\nDay: %s\n", dayOfWeek, day)
}
