package main

import (
    "encoding/json"
    "fmt"
)

type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
	Location string `json:"location"`
}

func main() {
    // Define a Person struct and populate its fields
    person := Person{Name: "Nithin", Age: 22,Location:"banglore"}

    // Marshal the struct to JSON bitslice
    content, err := json.Marshal(person)
    if err != nil {
        fmt.Println("Error marshaling JSON:", err)
        return
    }
    jsonStr := string(content)
    fmt.Println("Marshalled JSON string:", jsonStr)

    // Unmarshal the JSON bitslice into a struct
    var person2 Person
    err = json.Unmarshal(content, &person2)
    if err != nil {
        fmt.Println("Error unmarshaling JSON:", err)
        return
    }
    fmt.Println("Unmarshaled Person struct:",person2)
}
