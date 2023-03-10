package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Data struct {
	
	Name string 
	Age  int    
}

func main() {
	// Open the JSON file
	data, err := ioutil.ReadFile("file.json")
	if err != nil {
		panic(err)
	}

	// Parse the JSON data into a Go struct
	var jsonData Data
	err = json.Unmarshal(data, &jsonData)
	if err != nil {
		panic(err)
	}

	// Print the contents of the struct
	fmt.Println(jsonData)
}
