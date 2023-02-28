package main

import "fmt"

type Person struct {
    name string
    age  int
}

func main() {
    // Create a slice of Person structs
    people := []Person{
        {name: "Alice", age: 25},
        {name: "Bob", age: 30},
        {name: "Charlie", age: 35},
    }
    fmt.Println("People:", people)

    // Create a map of Person structs
    peopleMap := map[string]Person{
        "Alice":   {name: "Alice", age: 25},
        "Bob":     {name: "Bob", age: 30},
        "Charlie": {name: "Charlie", age: 35},
    }
    fmt.Println("People map:", peopleMap)
}
