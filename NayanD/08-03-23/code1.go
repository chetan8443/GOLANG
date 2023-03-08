//Write a program by using encoding/json package (marshal and unmarshal).

package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	Name  string
	Age   int
	email string
}

type Response struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

func main() {
	emp1 := Employee{Name: "Nayan Dhoble", Age: 25, email: "ndhoble@gmail.com"}

	//Marshalling in Golang

	empData, err := json.Marshal(emp1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Marshalling")

	fmt.Println(string(empData))

	// unmarshalling
	empJsonData := `{"Name":"Arbind Kumar","Age":90,"Address":"Newyork, USA"}`
	empBytes := []byte(empJsonData)
	var emp2 Response
	json.Unmarshal(empBytes, &emp2)
	fmt.Println("Unmarshalling")

	fmt.Println(emp2.Name)
	fmt.Println(emp2.Age)
	fmt.Println(emp2.Address)

}
