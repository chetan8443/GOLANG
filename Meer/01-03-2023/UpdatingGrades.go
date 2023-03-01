package main

import "fmt"

func main() {
	var grades = []int32{73, 67, 38, 33}
	result := gradingStudents(grades)
	fmt.Println(result)

}

func gradingStudents(grades []int32) []int32 {
	// Write your code here
	var updatedGrades []int32
	for _, v := range grades {
		if v < 38 {
			updatedGrades = append(updatedGrades, v)
		} else {
			if v%5 > 2 {
				// round our number
				updatedGrades = append(updatedGrades, ((v/5)+1)*5)
			} else {
				updatedGrades = append(updatedGrades, v)
			}
		}
	}
	return updatedGrades

}
