/*
package main

import "fmt"

	func swap(px, py *int) {
		tempx := *px
		tempy := *py
		*px = tempy
		*py = tempx
	}

	func main() {
		x := int(1)
		y := int(2)
		fmt.Println("x was", x)
		fmt.Println("y was", y)
		swap(&x, &y)

		fmt.Println("x is now", x)
		fmt.Println("y is now", y)
	}
*/
package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Message struct {
	Text string `json:"text"`
}

func main() {
	http.HandleFunc("/", helloHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	message := Message{Text: "Hello, World!"}
	json.NewEncoder(w).Encode(message)
}
