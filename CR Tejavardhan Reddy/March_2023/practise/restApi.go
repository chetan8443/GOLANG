package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Employee struct {
	eid   int    `json:"eid"`
	eName string `json:"eName"`
	age   int    `json:"age"`
}
type employee []Employee

func display(w http.ResponseWriter, r *http.Request) {
	e := employee{
		{eid: 100, eName: "Teja", age: 22},
	}
	fmt.Println("Starting")
	json.NewEncoder(w).Encode(e)
}
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/display", display)
	log.Fatal(http.ListenAndServe(":8082", nil))
	// fmt.Fprintf(w,"HEllo")
}

func main() {
	handleRequest()
}
