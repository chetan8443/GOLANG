package main

import "fmt"

func main() {
	//creating a map
	marks := map[string]int{"Python": 89, "Java": 87, "Golang": 90}
	fmt.Println(marks)
	//creating a map using var
	var subjectmarks = map[string]int{"Python": 89, "Java": 87, "Golang": 90}
	fmt.Println(subjectmarks)
	//accessing value for key Golang
	fmt.Println(marks["Golang"]) //90
	//change value
	subjectmarks["Golang"] = 92
	fmt.Println("updatedmap:", subjectmarks)
	//Adding element to map
	subjectmarks["c"] = 92
	fmt.Println(subjectmarks)
	//Removing an element from map
	delete(subjectmarks, "Python")
	fmt.Println(subjectmarks)
	//Looping through map
	for subject, marks := range subjectmarks {
		fmt.Println("marks in ", subject, "is", marks)
	}
}
