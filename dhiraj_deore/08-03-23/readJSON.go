package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Person struct {
	Name       string `json:"Name"`
	Surname    string `json:"Surname"`
	Country    string `json:"Country"`
	Gender     string `json:"Gender"`
	Profession string `json:"Profession"`
}

func main() {
	//read json file
	file, err := os.Open("data.json")
	if err != nil {
		log.Fatal(err)
	}
	//var details Person
	var details map[string]string
	data, err := ioutil.ReadAll(file)
	json.Unmarshal(data, &details)
	fmt.Print(details)

}
