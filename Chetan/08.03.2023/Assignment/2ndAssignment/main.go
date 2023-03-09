package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type info struct {
	Name    string `json:"Myname"`
	Address add    `json:"address"`
	Age     int    `json:"age"`
}
type add struct {
	City    string `json:"city"`
	State   string `json:"state"`
	Country string `json:"country"`
}

func main() {
	var std info
	dataBytes, err := ioutil.ReadFile("data.json")
	if err != nil {
		panic(fmt.Errorf("Error from readFile\n", err))
	}

	err = json.Unmarshal(dataBytes, &std)
	if err != nil {
		panic(fmt.Errorf("Error from readFile\n", err))
	}

	fmt.Printf("%#v", std)

}
