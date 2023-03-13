package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func main() {
	data, err := ioutil.ReadFile("demo.json")
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	var people []Person
	err = json.Unmarshal(data, &people)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	for _, v := range people {
		fmt.Printf(v.Name, v.Age)
	}

}
