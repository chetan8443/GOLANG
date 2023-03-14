//Write a program by using encoding/json package (marshal and unmarshal ).

package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string `json:"Website"`
	Password string `json:"-"`
}

func main() {
	fmt.Println("Welcome to online courses")
	encodejson()
	decodejson()
}
func encodejson() {
	Onlinecourse := []course{
		{"Golang", 299, "Geeksforgeeks", "abc123"},
		{"Reactjs", 399, "Javatpoint", "hem123"},
		{"Angular", 199, "w3schools", "chin123"},
	}
	jsondata, err := json.MarshalIndent(Onlinecourse, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", jsondata)
}

func decodejson() {
	newjsondata := []byte(`
	{
		"coursename": "Angular",
		"Price": 199,
		"Website": "w3schools"
	}
	 `)

	var newcourse course

	checkvalid := json.Valid(newjsondata)
	if checkvalid {
		fmt.Println("Json is valid")
		json.Unmarshal(newjsondata, &newcourse)
		fmt.Println(newcourse)
	} else {
		fmt.Println("Json is not valid")
	}
}
