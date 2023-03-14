package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:age`
}

func main() {
	data, err := ioutil.ReadFile("demo.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	var people []Person
	err = json.Unmarshal(data, &people)
	if err != nil {
		panic(err)
	}
	for _, p := range people {
		fmt.Printf("%s is %d years old\n", p.Name, p.Age)

	}

}
