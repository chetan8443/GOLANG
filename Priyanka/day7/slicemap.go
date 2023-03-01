package main

import "fmt"

type Person struct {
	Name    string
	Age     int
	Country string
}

func main() {
	people := []Person{
		{Name: "Priya", Age: 28, Country: "India"},
		{Name: "Bobby", Age: 18, Country: "UK"},
		{Name: "Hyma", Age: 29, Country: "Australia"},
	}
	fmt.Println("People:", people)

	peopleMap := map[string]Person{
		"Priya": {Name: "Priya", Age: 28, Country: "India"},
		"Bobby": {Name: "Bobby", Age: 18, Country: "UK"},
		"Hyma":  {Name: "Hyma", Age: 29, Country: "Australia"},
	}
	fmt.Println("People Map:", peopleMap)
}
