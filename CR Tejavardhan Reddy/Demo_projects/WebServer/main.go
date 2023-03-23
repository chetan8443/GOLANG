package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintln(w, "POST request successful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not Found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method id not supproted", http.StatusNotFound)
		return
	}
	fmt.Fprintln(w, "Hello")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)
	fmt.Println("Starting the connection.....")
	err := http.ListenAndServe(":8008", nil)
	CheckError(err)
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
