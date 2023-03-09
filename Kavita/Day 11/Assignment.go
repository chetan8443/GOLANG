//  A program using the "encoding/json" package . Serializing and Deserializing simple json object.

package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name string 
	Age  int    
}

func main() {

	students := Student{Name: "Kavita", Age: 28}

	// Serialize the Student object to JSON () - serializing(coverting the data into the format that can be trasmitted)
	jsonBytes, err := json.Marshal(students)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	jsonString := string(jsonBytes)
	fmt.Println("Serialized JSON:", jsonString)

	// Deserialize the JSON to a Person object - Deserialize (changing the json data into the readable format)
	var student1 Student
	err = json.Unmarshal(jsonBytes, &student1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Deserialized object:", student1)
}
